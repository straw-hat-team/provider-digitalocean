//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Registry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryDockerCredentials) DeepCopyInto(out *RegistryDockerCredentials) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryDockerCredentials.
func (in *RegistryDockerCredentials) DeepCopy() *RegistryDockerCredentials {
	if in == nil {
		return nil
	}
	out := new(RegistryDockerCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryDockerCredentials) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryDockerCredentialsInitParameters) DeepCopyInto(out *RegistryDockerCredentialsInitParameters) {
	*out = *in
	if in.ExpirySeconds != nil {
		in, out := &in.ExpirySeconds, &out.ExpirySeconds
		*out = new(float64)
		**out = **in
	}
	if in.RegistryName != nil {
		in, out := &in.RegistryName, &out.RegistryName
		*out = new(string)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryDockerCredentialsInitParameters.
func (in *RegistryDockerCredentialsInitParameters) DeepCopy() *RegistryDockerCredentialsInitParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryDockerCredentialsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryDockerCredentialsList) DeepCopyInto(out *RegistryDockerCredentialsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RegistryDockerCredentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryDockerCredentialsList.
func (in *RegistryDockerCredentialsList) DeepCopy() *RegistryDockerCredentialsList {
	if in == nil {
		return nil
	}
	out := new(RegistryDockerCredentialsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryDockerCredentialsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryDockerCredentialsObservation) DeepCopyInto(out *RegistryDockerCredentialsObservation) {
	*out = *in
	if in.CredentialExpirationTime != nil {
		in, out := &in.CredentialExpirationTime, &out.CredentialExpirationTime
		*out = new(string)
		**out = **in
	}
	if in.ExpirySeconds != nil {
		in, out := &in.ExpirySeconds, &out.ExpirySeconds
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RegistryName != nil {
		in, out := &in.RegistryName, &out.RegistryName
		*out = new(string)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryDockerCredentialsObservation.
func (in *RegistryDockerCredentialsObservation) DeepCopy() *RegistryDockerCredentialsObservation {
	if in == nil {
		return nil
	}
	out := new(RegistryDockerCredentialsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryDockerCredentialsParameters) DeepCopyInto(out *RegistryDockerCredentialsParameters) {
	*out = *in
	if in.ExpirySeconds != nil {
		in, out := &in.ExpirySeconds, &out.ExpirySeconds
		*out = new(float64)
		**out = **in
	}
	if in.RegistryName != nil {
		in, out := &in.RegistryName, &out.RegistryName
		*out = new(string)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryDockerCredentialsParameters.
func (in *RegistryDockerCredentialsParameters) DeepCopy() *RegistryDockerCredentialsParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryDockerCredentialsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryDockerCredentialsSpec) DeepCopyInto(out *RegistryDockerCredentialsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryDockerCredentialsSpec.
func (in *RegistryDockerCredentialsSpec) DeepCopy() *RegistryDockerCredentialsSpec {
	if in == nil {
		return nil
	}
	out := new(RegistryDockerCredentialsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryDockerCredentialsStatus) DeepCopyInto(out *RegistryDockerCredentialsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryDockerCredentialsStatus.
func (in *RegistryDockerCredentialsStatus) DeepCopy() *RegistryDockerCredentialsStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryDockerCredentialsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryInitParameters) DeepCopyInto(out *RegistryInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionTierSlug != nil {
		in, out := &in.SubscriptionTierSlug, &out.SubscriptionTierSlug
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryInitParameters.
func (in *RegistryInitParameters) DeepCopy() *RegistryInitParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryList) DeepCopyInto(out *RegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Registry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryList.
func (in *RegistryList) DeepCopy() *RegistryList {
	if in == nil {
		return nil
	}
	out := new(RegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryObservation) DeepCopyInto(out *RegistryObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ServerURL != nil {
		in, out := &in.ServerURL, &out.ServerURL
		*out = new(string)
		**out = **in
	}
	if in.StorageUsageBytes != nil {
		in, out := &in.StorageUsageBytes, &out.StorageUsageBytes
		*out = new(float64)
		**out = **in
	}
	if in.SubscriptionTierSlug != nil {
		in, out := &in.SubscriptionTierSlug, &out.SubscriptionTierSlug
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryObservation.
func (in *RegistryObservation) DeepCopy() *RegistryObservation {
	if in == nil {
		return nil
	}
	out := new(RegistryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryParameters) DeepCopyInto(out *RegistryParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionTierSlug != nil {
		in, out := &in.SubscriptionTierSlug, &out.SubscriptionTierSlug
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryParameters.
func (in *RegistryParameters) DeepCopy() *RegistryParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpec) DeepCopyInto(out *RegistrySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpec.
func (in *RegistrySpec) DeepCopy() *RegistrySpec {
	if in == nil {
		return nil
	}
	out := new(RegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryStatus) DeepCopyInto(out *RegistryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryStatus.
func (in *RegistryStatus) DeepCopy() *RegistryStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryStatus)
	in.DeepCopyInto(out)
	return out
}
