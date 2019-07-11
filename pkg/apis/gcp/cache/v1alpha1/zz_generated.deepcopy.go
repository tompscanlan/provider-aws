// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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
func (in *CloudMemorystoreInstance) DeepCopyInto(out *CloudMemorystoreInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMemorystoreInstance.
func (in *CloudMemorystoreInstance) DeepCopy() *CloudMemorystoreInstance {
	if in == nil {
		return nil
	}
	out := new(CloudMemorystoreInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudMemorystoreInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudMemorystoreInstanceList) DeepCopyInto(out *CloudMemorystoreInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudMemorystoreInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMemorystoreInstanceList.
func (in *CloudMemorystoreInstanceList) DeepCopy() *CloudMemorystoreInstanceList {
	if in == nil {
		return nil
	}
	out := new(CloudMemorystoreInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudMemorystoreInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudMemorystoreInstanceSpec) DeepCopyInto(out *CloudMemorystoreInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.RedisConfigs != nil {
		in, out := &in.RedisConfigs, &out.RedisConfigs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMemorystoreInstanceSpec.
func (in *CloudMemorystoreInstanceSpec) DeepCopy() *CloudMemorystoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudMemorystoreInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudMemorystoreInstanceStatus) DeepCopyInto(out *CloudMemorystoreInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMemorystoreInstanceStatus.
func (in *CloudMemorystoreInstanceStatus) DeepCopy() *CloudMemorystoreInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudMemorystoreInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
