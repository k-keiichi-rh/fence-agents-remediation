//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FenceAgentsRemediation) DeepCopyInto(out *FenceAgentsRemediation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FenceAgentsRemediation.
func (in *FenceAgentsRemediation) DeepCopy() *FenceAgentsRemediation {
	if in == nil {
		return nil
	}
	out := new(FenceAgentsRemediation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FenceAgentsRemediation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FenceAgentsRemediationList) DeepCopyInto(out *FenceAgentsRemediationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FenceAgentsRemediation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FenceAgentsRemediationList.
func (in *FenceAgentsRemediationList) DeepCopy() *FenceAgentsRemediationList {
	if in == nil {
		return nil
	}
	out := new(FenceAgentsRemediationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FenceAgentsRemediationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FenceAgentsRemediationSpec) DeepCopyInto(out *FenceAgentsRemediationSpec) {
	*out = *in
	if in.SharedParameters != nil {
		in, out := &in.SharedParameters, &out.SharedParameters
		*out = make(map[ParameterName]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeParameters != nil {
		in, out := &in.NodeParameters, &out.NodeParameters
		*out = make(map[ParameterName]map[NodeName]string, len(*in))
		for key, val := range *in {
			var outVal map[NodeName]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[NodeName]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FenceAgentsRemediationSpec.
func (in *FenceAgentsRemediationSpec) DeepCopy() *FenceAgentsRemediationSpec {
	if in == nil {
		return nil
	}
	out := new(FenceAgentsRemediationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FenceAgentsRemediationStatus) DeepCopyInto(out *FenceAgentsRemediationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FenceAgentsRemediationStatus.
func (in *FenceAgentsRemediationStatus) DeepCopy() *FenceAgentsRemediationStatus {
	if in == nil {
		return nil
	}
	out := new(FenceAgentsRemediationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FenceAgentsRemediationTemplate) DeepCopyInto(out *FenceAgentsRemediationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FenceAgentsRemediationTemplate.
func (in *FenceAgentsRemediationTemplate) DeepCopy() *FenceAgentsRemediationTemplate {
	if in == nil {
		return nil
	}
	out := new(FenceAgentsRemediationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FenceAgentsRemediationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FenceAgentsRemediationTemplateList) DeepCopyInto(out *FenceAgentsRemediationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FenceAgentsRemediationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FenceAgentsRemediationTemplateList.
func (in *FenceAgentsRemediationTemplateList) DeepCopy() *FenceAgentsRemediationTemplateList {
	if in == nil {
		return nil
	}
	out := new(FenceAgentsRemediationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FenceAgentsRemediationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FenceAgentsRemediationTemplateResource) DeepCopyInto(out *FenceAgentsRemediationTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FenceAgentsRemediationTemplateResource.
func (in *FenceAgentsRemediationTemplateResource) DeepCopy() *FenceAgentsRemediationTemplateResource {
	if in == nil {
		return nil
	}
	out := new(FenceAgentsRemediationTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FenceAgentsRemediationTemplateSpec) DeepCopyInto(out *FenceAgentsRemediationTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FenceAgentsRemediationTemplateSpec.
func (in *FenceAgentsRemediationTemplateSpec) DeepCopy() *FenceAgentsRemediationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(FenceAgentsRemediationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FenceAgentsRemediationTemplateStatus) DeepCopyInto(out *FenceAgentsRemediationTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FenceAgentsRemediationTemplateStatus.
func (in *FenceAgentsRemediationTemplateStatus) DeepCopy() *FenceAgentsRemediationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(FenceAgentsRemediationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}
