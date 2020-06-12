// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArticlesSpec) DeepCopyInto(out *ArticlesSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArticlesSpec.
func (in *ArticlesSpec) DeepCopy() *ArticlesSpec {
	if in == nil {
		return nil
	}
	out := new(ArticlesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Weekly) DeepCopyInto(out *Weekly) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Weekly.
func (in *Weekly) DeepCopy() *Weekly {
	if in == nil {
		return nil
	}
	out := new(Weekly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Weekly) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyList) DeepCopyInto(out *WeeklyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Weekly, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyList.
func (in *WeeklyList) DeepCopy() *WeeklyList {
	if in == nil {
		return nil
	}
	out := new(WeeklyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WeeklyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklySpec) DeepCopyInto(out *WeeklySpec) {
	*out = *in
	if in.Articles != nil {
		in, out := &in.Articles, &out.Articles
		*out = make([]ArticlesSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklySpec.
func (in *WeeklySpec) DeepCopy() *WeeklySpec {
	if in == nil {
		return nil
	}
	out := new(WeeklySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyStatus) DeepCopyInto(out *WeeklyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyStatus.
func (in *WeeklyStatus) DeepCopy() *WeeklyStatus {
	if in == nil {
		return nil
	}
	out := new(WeeklyStatus)
	in.DeepCopyInto(out)
	return out
}