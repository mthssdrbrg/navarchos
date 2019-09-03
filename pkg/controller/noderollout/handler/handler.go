package handler

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	navarchosv1alpha1 "github.com/pusher/navarchos/pkg/apis/navarchos/v1alpha1"
	"github.com/pusher/navarchos/pkg/controller/noderollout/status"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metalabels "k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Options are used to configure the NodeRolloutHandler
type Options struct {
	// MaxAge determines the maximum age a NodeRollout should be before it is
	// garbage collected
	MaxAge *time.Duration
}

// Complete defaults any values that are not explicitly set
func (o *Options) Complete() {
	if o.MaxAge == nil {
		maxAge := 48 * time.Hour
		o.MaxAge = &maxAge
	}
}

// NodeRolloutHandler handles the business logic within the NodeRollout controller.
type NodeRolloutHandler struct {
	client client.Client
	maxAge time.Duration
}

// nodeReplacementSpec is a container to allow easier construction of
// NodeReplacements
type nodeReplacementSpec struct {
	node            corev1.Node
	replacementSpec navarchosv1alpha1.NodeReplacementSpec
}

// replacementCreationResult is a container struct used for returning errors and
// names when creating NodeReplacements
type replacementCreationResult struct {
	err                error
	replacementCreated string
}

// NewNodeRolloutHandler creates a new NodeRolloutHandler
func NewNodeRolloutHandler(c client.Client, opts *Options) *NodeRolloutHandler {
	opts.Complete()
	return &NodeRolloutHandler{
		client: c,
		maxAge: *opts.MaxAge,
	}
}

// Handle performs the business logic of the NodeRollout and returns information
// in a Result
func (h *NodeRolloutHandler) Handle(instance *navarchosv1alpha1.NodeRollout) *status.Result {
	switch instance.Status.Phase {
	case navarchosv1alpha1.RolloutPhaseNew:
		return h.handleNew(instance)
	case navarchosv1alpha1.RolloutPhaseInProgress:
		return h.handleInProgress(instance)
	case navarchosv1alpha1.RolloutPhaseCompleted:
		return h.handleCompleted(instance)
	default:
		return h.handleNew(instance)
	}
}

// handleNew handles a NodeRollout in the 'New' phase. It creates
// NodeReplacements from the provided NodeRollout instance and updates the phase
// to in progress if it does not fail
func (h *NodeRolloutHandler) handleNew(instance *navarchosv1alpha1.NodeRollout) *status.Result {
	result := &status.Result{}
	nodes := &corev1.NodeList{}
	err := h.client.List(context.Background(), nodes)
	if err != nil {
		result.ReplacementsCreatedError = fmt.Errorf("failed to list nodes: %v", err)
		return result
	}

	nodeReplacementMap := make(map[string]nodeReplacementSpec)
	nodeReplacementMap, err = filterNodeSelectors(nodes, instance.Spec.NodeSelectors, nodeReplacementMap)
	if err != nil {
		result.ReplacementsCreatedError = fmt.Errorf("failed to filter nodes: %v", err)
		return result
	}
	nodeReplacementMap = filterNodeNames(nodes, instance.Spec.NodeNames, nodeReplacementMap)

	outputChannel, err := h.createNodeReplacements(nodeReplacementMap, instance)
	if err != nil {
		result.ReplacementsCreatedError = fmt.Errorf("failed to create node replacements: %v", err)
		return result
	}

	// retrieve any errors.
	errMap := make(map[error]int)
	for output := range outputChannel {
		if output.err != nil {
			errMap[output.err]++
		} else {
			result.ReplacementsCreated = append(result.ReplacementsCreated, output.replacementCreated)
		}
	}

	// if there are any errors concatenate and return them, don't update the
	// phase
	if len(errMap) > 0 {
		errSlice := []string{}
		for errName, count := range errMap {
			errSlice = append(errSlice, fmt.Sprintf("Error: \"%s\" has occurred \"%d\" time(s)", errName.Error(), count))
		}
		result.ReplacementsCreatedError = fmt.Errorf(strings.Join(errSlice, ",\n"))
		return result
	}

	inProgress := navarchosv1alpha1.RolloutPhaseInProgress
	result.Phase = &inProgress

	return result
}

// filterNodeSelectors filters the list of all nodes.  If a nodes labels match
// it adds the node to the nodeMap
func filterNodeSelectors(nodes *corev1.NodeList, selectors []navarchosv1alpha1.NodeLabelSelector, nodeMap map[string]nodeReplacementSpec) (map[string]nodeReplacementSpec, error) {
	for _, nls := range selectors {
		selector, err := metav1.LabelSelectorAsSelector(&nls.LabelSelector)
		if err != nil {
			return nil, err
		}
		// check which nodes match the LabelSelector
		for _, node := range nodes.Items {
			labels := metalabels.Set(node.GetLabels())
			if selector.Matches(labels) {
				nodeMap[node.GetName()] = newNodeReplacementSpec(node, nls.ReplacementSpec)
			}

		}
	}
	return nodeMap, nil
}

// newNodeReplacementSpec takes a node and a ReplacementSpec and returns a
// nodeReplacementSpec
func newNodeReplacementSpec(node corev1.Node, replacementSpec navarchosv1alpha1.ReplacementSpec) nodeReplacementSpec {
	return nodeReplacementSpec{
		node: node,
		replacementSpec: navarchosv1alpha1.NodeReplacementSpec{
			ReplacementSpec: replacementSpec,
			NodeName:        node.GetName(),
			NodeUID:         node.GetUID(),
		},
	}
}

// filterNodeNames filters the list of all nodes. If a nodes name matches one
// provided it adds the node to the nodeMap
func filterNodeNames(nodes *corev1.NodeList, nodeNames []navarchosv1alpha1.NodeName, nodeMap map[string]nodeReplacementSpec) map[string]nodeReplacementSpec {
	for _, selectedName := range nodeNames {
		for _, node := range nodes.Items {
			if node.GetName() == selectedName.Name {
				nodeMap[node.GetName()] = newNodeReplacementSpec(node, selectedName.ReplacementSpec)
				break
			}
		}
	}
	return nodeMap
}

// createNodeReplacementFromSpec takes a NodeReplacementSpec, NodeRollout and
// node and returns a NodeReplacement with the correct owners
func createNodeReplacementFromSpec(spec navarchosv1alpha1.NodeReplacementSpec, rolloutOwner *navarchosv1alpha1.NodeRollout, nodeOwner *corev1.Node) *navarchosv1alpha1.NodeReplacement {
	nodeReplacement := &navarchosv1alpha1.NodeReplacement{
		TypeMeta: metav1.TypeMeta{
			APIVersion: navarchosv1alpha1.SchemeGroupVersion.String(),
			Kind:       "NodeReplacement",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-", spec.NodeName),
			OwnerReferences: []metav1.OwnerReference{
				newOwnerRef(rolloutOwner, rolloutOwner.GroupVersionKind(), true, true),
				newOwnerRef(nodeOwner, nodeOwner.GroupVersionKind(), false, false),
			},
		},
		Spec:   spec,
		Status: navarchosv1alpha1.NodeReplacementStatus{},
	}
	return nodeReplacement.DeepCopy()
}

// newOwnerRef creates an OwnerReference pointing to the given owner
func newOwnerRef(owner metav1.Object, gvk schema.GroupVersionKind, isController bool, blockOwnerDeletion bool) metav1.OwnerReference {
	return metav1.OwnerReference{
		APIVersion:         gvk.GroupVersion().String(),
		Kind:               gvk.Kind,
		Name:               owner.GetName(),
		UID:                owner.GetUID(),
		BlockOwnerDeletion: &blockOwnerDeletion,
		Controller:         &isController,
	}
}

// createNodeRollouts uses a new goroutine to create each NodeReplacment in
// parallel. If a node already has a replacement that is owned by the same
// rollout, it is skipped
func (h *NodeRolloutHandler) createNodeReplacements(nodeReplacementMap map[string]nodeReplacementSpec, instance *navarchosv1alpha1.NodeRollout) (<-chan replacementCreationResult, error) {
	existingNodeReplacements := &navarchosv1alpha1.NodeReplacementList{}
	err := h.client.List(context.Background(), existingNodeReplacements)
	if err != nil {
		return nil, fmt.Errorf("error listing NodeReplacements: %v", err)
	}

	filteredNodeReplacements := filterReplacementsByOwner(existingNodeReplacements, instance)

	outputChannel := make(chan replacementCreationResult, len(nodeReplacementMap))
	var wg sync.WaitGroup

	wg.Add(len(nodeReplacementMap))
	for _, spec := range nodeReplacementMap {
		go func(spec nodeReplacementSpec, instance *navarchosv1alpha1.NodeRollout, filteredNr []navarchosv1alpha1.NodeReplacement, client client.Client) {
			defer wg.Done()
			nodeReplacement := createNodeReplacementFromSpec(spec.replacementSpec, instance, &spec.node)

			if !replacementAlreadyExists(filteredNr, nodeReplacement) {
				err := client.Create(context.Background(), nodeReplacement)
				if err != nil {
					outputChannel <- replacementCreationResult{err: fmt.Errorf("failed to create NodeReplacement: %v", err), replacementCreated: ""}
				}
			}

			outputChannel <- replacementCreationResult{err: nil, replacementCreated: spec.replacementSpec.NodeName}
		}(spec, instance, filteredNodeReplacements, h.client)
	}

	wg.Wait()
	close(outputChannel)
	return outputChannel, nil
}

// filterReplacementsByOwner takes a list of NodeReplacements and a NodeRollout.
// Any NodeReplacements that are owned by the NodeRollout are returned as a list
func filterReplacementsByOwner(nodeReplacementList *navarchosv1alpha1.NodeReplacementList, instance *navarchosv1alpha1.NodeRollout) []navarchosv1alpha1.NodeReplacement {
	rolloutOwnerRef := newOwnerRef(instance, instance.GroupVersionKind(), true, true)
	nodeReplacements := []navarchosv1alpha1.NodeReplacement{}

	for _, nr := range nodeReplacementList.Items {
		if containsOwnerReference(nr.ObjectMeta.OwnerReferences, rolloutOwnerRef) {
			nodeReplacements = append(nodeReplacements, nr)
		}
	}
	return nodeReplacements
}

// containsOwnerReference returns true if ownerRef exists in the list ownerRefs
func containsOwnerReference(ownerRefs []metav1.OwnerReference, ownerRef metav1.OwnerReference) bool {
	for _, or := range ownerRefs {
		if reflect.DeepEqual(or, ownerRef) {
			return true
		}
	}
	return false
}

// replacementAlreadyExists returns true if nodeReplacement exists in the list nodeReplacements
func replacementAlreadyExists(nodeReplacements []navarchosv1alpha1.NodeReplacement, nodeReplacement *navarchosv1alpha1.NodeReplacement) bool {
	for _, nr := range nodeReplacements {
		if (nr.Spec.NodeUID == nodeReplacement.Spec.NodeUID) && (nr.Spec.NodeName == nodeReplacement.Spec.NodeName) {
			return true
		}
	}
	return false
}

func (h *NodeRolloutHandler) handleInProgress(instance *navarchosv1alpha1.NodeRollout) *status.Result {
	return &status.Result{}
}

func (h *NodeRolloutHandler) handleCompleted(instance *navarchosv1alpha1.NodeRollout) *status.Result {
	return &status.Result{}
}
