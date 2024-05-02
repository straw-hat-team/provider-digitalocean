// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AttachmentInitParameters struct {
}

type AttachmentObservation struct {

	// ID of the Droplet to attach the volume to.
	DropletID *float64 `json:"dropletId,omitempty" tf:"droplet_id,omitempty"`

	// The unique identifier for the volume attachment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the Volume to be attached to the Droplet.
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type AttachmentParameters struct {

	// ID of the Droplet to attach the volume to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-digitalocean/apis/droplet/v1alpha1.Droplet
	// +kubebuilder:validation:Optional
	DropletID *float64 `json:"dropletId,omitempty" tf:"droplet_id,omitempty"`

	// Reference to a Droplet in droplet to populate dropletId.
	// +kubebuilder:validation:Optional
	DropletIDRef *v1.Reference `json:"dropletIdRef,omitempty" tf:"-"`

	// Selector for a Droplet in droplet to populate dropletId.
	// +kubebuilder:validation:Optional
	DropletIDSelector *v1.Selector `json:"dropletIdSelector,omitempty" tf:"-"`

	// ID of the Volume to be attached to the Droplet.
	// +crossplane:generate:reference:type=Volume
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// Reference to a Volume to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// Selector for a Volume to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`
}

// AttachmentSpec defines the desired state of Attachment
type AttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttachmentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AttachmentInitParameters `json:"initProvider,omitempty"`
}

// AttachmentStatus defines the observed state of Attachment.
type AttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Attachment is the Schema for the Attachments API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type Attachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AttachmentSpec   `json:"spec"`
	Status            AttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachmentList contains a list of Attachments
type AttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Attachment `json:"items"`
}

// Repository type metadata.
var (
	Attachment_Kind             = "Attachment"
	Attachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Attachment_Kind}.String()
	Attachment_KindAPIVersion   = Attachment_Kind + "." + CRDGroupVersion.String()
	Attachment_GroupVersionKind = CRDGroupVersion.WithKind(Attachment_Kind)
)

func init() {
	SchemeBuilder.Register(&Attachment{}, &AttachmentList{})
}
