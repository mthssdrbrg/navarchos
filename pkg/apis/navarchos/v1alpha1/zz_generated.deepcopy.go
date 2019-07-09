// +build !ignore_autogenerated

/*
Copyright 2019 Pusher Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeReplacement) DeepCopyInto(out *NodeReplacement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeReplacement.
func (in *NodeReplacement) DeepCopy() *NodeReplacement {
	if in == nil {
		return nil
	}
	out := new(NodeReplacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeReplacement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeReplacementList) DeepCopyInto(out *NodeReplacementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeReplacement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeReplacementList.
func (in *NodeReplacementList) DeepCopy() *NodeReplacementList {
	if in == nil {
		return nil
	}
	out := new(NodeReplacementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeReplacementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeReplacementSpec) DeepCopyInto(out *NodeReplacementSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeReplacementSpec.
func (in *NodeReplacementSpec) DeepCopy() *NodeReplacementSpec {
	if in == nil {
		return nil
	}
	out := new(NodeReplacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeReplacementStatus) DeepCopyInto(out *NodeReplacementStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeReplacementStatus.
func (in *NodeReplacementStatus) DeepCopy() *NodeReplacementStatus {
	if in == nil {
		return nil
	}
	out := new(NodeReplacementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRollout) DeepCopyInto(out *NodeRollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRollout.
func (in *NodeRollout) DeepCopy() *NodeRollout {
	if in == nil {
		return nil
	}
	out := new(NodeRollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeRollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRolloutCondition) DeepCopyInto(out *NodeRolloutCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRolloutCondition.
func (in *NodeRolloutCondition) DeepCopy() *NodeRolloutCondition {
	if in == nil {
		return nil
	}
	out := new(NodeRolloutCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRolloutList) DeepCopyInto(out *NodeRolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeRollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRolloutList.
func (in *NodeRolloutList) DeepCopy() *NodeRolloutList {
	if in == nil {
		return nil
	}
	out := new(NodeRolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeRolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRolloutSpec) DeepCopyInto(out *NodeRolloutSpec) {
	*out = *in
	if in.NodeSelectors != nil {
		in, out := &in.NodeSelectors, &out.NodeSelectors
		*out = make([]PriorityLabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = make([]PriorityName, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRolloutSpec.
func (in *NodeRolloutSpec) DeepCopy() *NodeRolloutSpec {
	if in == nil {
		return nil
	}
	out := new(NodeRolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRolloutStatus) DeepCopyInto(out *NodeRolloutStatus) {
	*out = *in
	if in.ReplacementsCreated != nil {
		in, out := &in.ReplacementsCreated, &out.ReplacementsCreated
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReplacementsCompleted != nil {
		in, out := &in.ReplacementsCompleted, &out.ReplacementsCompleted
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReplacementsFailed != nil {
		in, out := &in.ReplacementsFailed, &out.ReplacementsFailed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRolloutStatus.
func (in *NodeRolloutStatus) DeepCopy() *NodeRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(NodeRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLabelSelector) DeepCopyInto(out *PriorityLabelSelector) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLabelSelector.
func (in *PriorityLabelSelector) DeepCopy() *PriorityLabelSelector {
	if in == nil {
		return nil
	}
	out := new(PriorityLabelSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityName) DeepCopyInto(out *PriorityName) {
	*out = *in
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityName.
func (in *PriorityName) DeepCopy() *PriorityName {
	if in == nil {
		return nil
	}
	out := new(PriorityName)
	in.DeepCopyInto(out)
	return out
}
