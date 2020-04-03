// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hyperfoil) DeepCopyInto(out *Hyperfoil) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hyperfoil.
func (in *Hyperfoil) DeepCopy() *Hyperfoil {
	if in == nil {
		return nil
	}
	out := new(Hyperfoil)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hyperfoil) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperfoilList) DeepCopyInto(out *HyperfoilList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hyperfoil, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperfoilList.
func (in *HyperfoilList) DeepCopy() *HyperfoilList {
	if in == nil {
		return nil
	}
	out := new(HyperfoilList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperfoilList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperfoilSpec) DeepCopyInto(out *HyperfoilSpec) {
	*out = *in
	if in.SecretEnvVars != nil {
		in, out := &in.SecretEnvVars, &out.SecretEnvVars
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperfoilSpec.
func (in *HyperfoilSpec) DeepCopy() *HyperfoilSpec {
	if in == nil {
		return nil
	}
	out := new(HyperfoilSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperfoilStatus) DeepCopyInto(out *HyperfoilStatus) {
	*out = *in
	in.LastUpdate.DeepCopyInto(&out.LastUpdate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperfoilStatus.
func (in *HyperfoilStatus) DeepCopy() *HyperfoilStatus {
	if in == nil {
		return nil
	}
	out := new(HyperfoilStatus)
	in.DeepCopyInto(out)
	return out
}
