//go:build !ignore_autogenerated

/*
Copyright 2024.

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
	"github.com/nukleros/operator-builder-tools/pkg/status"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformConfig) DeepCopyInto(out *PlatformConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformConfig.
func (in *PlatformConfig) DeepCopy() *PlatformConfig {
	if in == nil {
		return nil
	}
	out := new(PlatformConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformConfigList) DeepCopyInto(out *PlatformConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PlatformConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformConfigList.
func (in *PlatformConfigList) DeepCopy() *PlatformConfigList {
	if in == nil {
		return nil
	}
	out := new(PlatformConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformConfigSpec) DeepCopyInto(out *PlatformConfigSpec) {
	*out = *in
	out.Platform = in.Platform
	out.Cloud = in.Cloud
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformConfigSpec.
func (in *PlatformConfigSpec) DeepCopy() *PlatformConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PlatformConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformConfigSpecCloud) DeepCopyInto(out *PlatformConfigSpecCloud) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformConfigSpecCloud.
func (in *PlatformConfigSpecCloud) DeepCopy() *PlatformConfigSpecCloud {
	if in == nil {
		return nil
	}
	out := new(PlatformConfigSpecCloud)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformConfigSpecPlatform) DeepCopyInto(out *PlatformConfigSpecPlatform) {
	*out = *in
	out.Certificates = in.Certificates
	out.Identity = in.Identity
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformConfigSpecPlatform.
func (in *PlatformConfigSpecPlatform) DeepCopy() *PlatformConfigSpecPlatform {
	if in == nil {
		return nil
	}
	out := new(PlatformConfigSpecPlatform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformConfigSpecPlatformCertificates) DeepCopyInto(out *PlatformConfigSpecPlatformCertificates) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformConfigSpecPlatformCertificates.
func (in *PlatformConfigSpecPlatformCertificates) DeepCopy() *PlatformConfigSpecPlatformCertificates {
	if in == nil {
		return nil
	}
	out := new(PlatformConfigSpecPlatformCertificates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformConfigSpecPlatformIdentity) DeepCopyInto(out *PlatformConfigSpecPlatformIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformConfigSpecPlatformIdentity.
func (in *PlatformConfigSpecPlatformIdentity) DeepCopy() *PlatformConfigSpecPlatformIdentity {
	if in == nil {
		return nil
	}
	out := new(PlatformConfigSpecPlatformIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformConfigStatus) DeepCopyInto(out *PlatformConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*status.PhaseCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.PhaseCondition)
				**out = **in
			}
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*status.ChildResource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.ChildResource)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformConfigStatus.
func (in *PlatformConfigStatus) DeepCopy() *PlatformConfigStatus {
	if in == nil {
		return nil
	}
	out := new(PlatformConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformOperators) DeepCopyInto(out *PlatformOperators) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformOperators.
func (in *PlatformOperators) DeepCopy() *PlatformOperators {
	if in == nil {
		return nil
	}
	out := new(PlatformOperators)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformOperators) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformOperatorsList) DeepCopyInto(out *PlatformOperatorsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PlatformOperators, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformOperatorsList.
func (in *PlatformOperatorsList) DeepCopy() *PlatformOperatorsList {
	if in == nil {
		return nil
	}
	out := new(PlatformOperatorsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformOperatorsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformOperatorsSpec) DeepCopyInto(out *PlatformOperatorsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformOperatorsSpec.
func (in *PlatformOperatorsSpec) DeepCopy() *PlatformOperatorsSpec {
	if in == nil {
		return nil
	}
	out := new(PlatformOperatorsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformOperatorsStatus) DeepCopyInto(out *PlatformOperatorsStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*status.PhaseCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.PhaseCondition)
				**out = **in
			}
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*status.ChildResource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.ChildResource)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformOperatorsStatus.
func (in *PlatformOperatorsStatus) DeepCopy() *PlatformOperatorsStatus {
	if in == nil {
		return nil
	}
	out := new(PlatformOperatorsStatus)
	in.DeepCopyInto(out)
	return out
}
