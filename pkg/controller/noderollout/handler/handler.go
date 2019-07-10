package handler

import (
	"time"

	navarchosv1alpha1 "github.com/pusher/navarchos/pkg/apis/navarchos/v1alpha1"
	"github.com/pusher/navarchos/pkg/controller/noderollout/status"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// This is the maximum age that Completed/Failed NodeRollouts are allowed to
// be before they are cleaned up by the Handler.
const maxNodeRolloutAge = time.Hour * 48

// NodeRolloutHandler handles the business logic within the NodeRollout controller.
type NodeRolloutHandler struct {
	client client.Client
}

// NewNodeRolloutHandler creates a new NodeRolloutHandler
func NewNodeRolloutHandler(c client.Client) *NodeRolloutHandler {
	return &NodeRolloutHandler{client: c}
}

// Handle performs the business logic of the NodeRollout and returns information
// in a Result
func (h *NodeRolloutHandler) Handle(instance *navarchosv1alpha1.NodeRollout) *status.Result {
	return &status.Result{}
}
