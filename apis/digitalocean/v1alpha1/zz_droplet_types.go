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

type DropletInitParameters struct {

	// Boolean controlling if backups are made. Defaults to
	// false.
	Backups *bool `json:"backups,omitempty" tf:"backups,omitempty"`

	// A boolean indicating whether to install the
	// DigitalOcean agent used for providing access to the Droplet web console in
	// the control panel. By default, the agent is installed on new Droplets but
	// installation errors (i.e. OS not supported) are ignored. To prevent it from
	// being installed, set to false. To make installation errors fatal, explicitly
	// set it to true.
	DropletAgent *bool `json:"dropletAgent,omitempty" tf:"droplet_agent,omitempty"`

	// A boolean indicating whether the droplet
	// should be gracefully shut down before it is deleted.
	GracefulShutdown *bool `json:"gracefulShutdown,omitempty" tf:"graceful_shutdown,omitempty"`

	// Boolean controlling if IPv6 is enabled. Defaults to false.
	// Once enabled for a Droplet, IPv6 can not be disabled. When enabling IPv6 on
	// an existing Droplet, additional OS-level configuration
	// is required.
	IPv6 *bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// The IPv6 address
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// Boolean controlling whether monitoring agent is installed.
	// Defaults to false. If set to true, you can configure monitor alert policies
	// monitor alert resource
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// The Droplet name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Deprecated Boolean controlling if private networking
	// is enabled. This parameter has been deprecated. Use vpc_uuid instead to specify a VPC network for the Droplet. If no vpc_uuid is provided, the Droplet will be placed in your account's default VPC for the region.
	PrivateNetworking *bool `json:"privateNetworking,omitempty" tf:"private_networking,omitempty"`

	// The region where the Droplet will be created.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Boolean controlling whether to increase the disk
	// size when resizing a Droplet. It defaults to true. When set to false,
	// only the Droplet's RAM and CPU will be resized. Increasing a Droplet's disk
	// size is a permanent change. Increasing only RAM and CPU is reversible.
	ResizeDisk *bool `json:"resizeDisk,omitempty" tf:"resize_disk,omitempty"`

	// The unique slug that indentifies the type of Droplet. You can find a list of available slugs on DigitalOcean API documentation.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// A string of the desired User Data for the Droplet.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// A list of the IDs of each block storage volume to be attached to the Droplet.
	VolumeIds []*string `json:"volumeIds,omitempty" tf:"volume_ids,omitempty"`
}

type DropletObservation struct {

	// Boolean controlling if backups are made. Defaults to
	// false.
	Backups *bool `json:"backups,omitempty" tf:"backups,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The size of the instance's disk in GB
	Disk *float64 `json:"disk,omitempty" tf:"disk,omitempty"`

	// A boolean indicating whether to install the
	// DigitalOcean agent used for providing access to the Droplet web console in
	// the control panel. By default, the agent is installed on new Droplets but
	// installation errors (i.e. OS not supported) are ignored. To prevent it from
	// being installed, set to false. To make installation errors fatal, explicitly
	// set it to true.
	DropletAgent *bool `json:"dropletAgent,omitempty" tf:"droplet_agent,omitempty"`

	// A boolean indicating whether the droplet
	// should be gracefully shut down before it is deleted.
	GracefulShutdown *bool `json:"gracefulShutdown,omitempty" tf:"graceful_shutdown,omitempty"`

	// The ID of the Droplet
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IPv4 address
	IPv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address,omitempty"`

	// The private networking IPv4 address
	IPv4AddressPrivate *string `json:"ipv4AddressPrivate,omitempty" tf:"ipv4_address_private,omitempty"`

	// Boolean controlling if IPv6 is enabled. Defaults to false.
	// Once enabled for a Droplet, IPv6 can not be disabled. When enabling IPv6 on
	// an existing Droplet, additional OS-level configuration
	// is required.
	IPv6 *bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// The IPv6 address
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// The Droplet image ID or slug. This could be either image ID or droplet snapshot ID.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Is the Droplet locked
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`

	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// Boolean controlling whether monitoring agent is installed.
	// Defaults to false. If set to true, you can configure monitor alert policies
	// monitor alert resource
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// The Droplet name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Droplet hourly price
	PriceHourly *float64 `json:"priceHourly,omitempty" tf:"price_hourly,omitempty"`

	// Droplet monthly price
	PriceMonthly *float64 `json:"priceMonthly,omitempty" tf:"price_monthly,omitempty"`

	// Deprecated Boolean controlling if private networking
	// is enabled. This parameter has been deprecated. Use vpc_uuid instead to specify a VPC network for the Droplet. If no vpc_uuid is provided, the Droplet will be placed in your account's default VPC for the region.
	PrivateNetworking *bool `json:"privateNetworking,omitempty" tf:"private_networking,omitempty"`

	// The region where the Droplet will be created.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Boolean controlling whether to increase the disk
	// size when resizing a Droplet. It defaults to true. When set to false,
	// only the Droplet's RAM and CPU will be resized. Increasing a Droplet's disk
	// size is a permanent change. Increasing only RAM and CPU is reversible.
	ResizeDisk *bool `json:"resizeDisk,omitempty" tf:"resize_disk,omitempty"`

	// A list of SSH key IDs or fingerprints to enable in
	// the format [12345, 123456]. To retrieve this info, use the
	// DigitalOcean API
	// or CLI (doctl compute ssh-key list). Once a Droplet is created keys can not
	// be added or removed via this provider. Modifying this field will prompt you
	// to destroy and recreate the Droplet.
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// The unique slug that indentifies the type of Droplet. You can find a list of available slugs on DigitalOcean API documentation.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// The status of the Droplet
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A list of the tags to be applied to this Droplet.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The uniform resource name of the Droplet
	Urn *string `json:"urn,omitempty" tf:"urn,omitempty"`

	// A string of the desired User Data for the Droplet.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The ID of the VPC where the Droplet will be located.
	VPCUUID *string `json:"vpcUuid,omitempty" tf:"vpc_uuid,omitempty"`

	// The number of the instance's virtual CPUs
	Vcpus *float64 `json:"vcpus,omitempty" tf:"vcpus,omitempty"`

	// A list of the IDs of each block storage volume to be attached to the Droplet.
	VolumeIds []*string `json:"volumeIds,omitempty" tf:"volume_ids,omitempty"`
}

type DropletParameters struct {

	// Boolean controlling if backups are made. Defaults to
	// false.
	// +kubebuilder:validation:Optional
	Backups *bool `json:"backups,omitempty" tf:"backups,omitempty"`

	// A boolean indicating whether to install the
	// DigitalOcean agent used for providing access to the Droplet web console in
	// the control panel. By default, the agent is installed on new Droplets but
	// installation errors (i.e. OS not supported) are ignored. To prevent it from
	// being installed, set to false. To make installation errors fatal, explicitly
	// set it to true.
	// +kubebuilder:validation:Optional
	DropletAgent *bool `json:"dropletAgent,omitempty" tf:"droplet_agent,omitempty"`

	// A boolean indicating whether the droplet
	// should be gracefully shut down before it is deleted.
	// +kubebuilder:validation:Optional
	GracefulShutdown *bool `json:"gracefulShutdown,omitempty" tf:"graceful_shutdown,omitempty"`

	// Boolean controlling if IPv6 is enabled. Defaults to false.
	// Once enabled for a Droplet, IPv6 can not be disabled. When enabling IPv6 on
	// an existing Droplet, additional OS-level configuration
	// is required.
	// +kubebuilder:validation:Optional
	IPv6 *bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// The IPv6 address
	// +kubebuilder:validation:Optional
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// The Droplet image ID or slug. This could be either image ID or droplet snapshot ID.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-digitalocean/apis/custom/v1alpha1.Image
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Reference to a Image in custom to populate image.
	// +kubebuilder:validation:Optional
	ImageRef *v1.Reference `json:"imageRef,omitempty" tf:"-"`

	// Selector for a Image in custom to populate image.
	// +kubebuilder:validation:Optional
	ImageSelector *v1.Selector `json:"imageSelector,omitempty" tf:"-"`

	// Boolean controlling whether monitoring agent is installed.
	// Defaults to false. If set to true, you can configure monitor alert policies
	// monitor alert resource
	// +kubebuilder:validation:Optional
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// The Droplet name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Deprecated Boolean controlling if private networking
	// is enabled. This parameter has been deprecated. Use vpc_uuid instead to specify a VPC network for the Droplet. If no vpc_uuid is provided, the Droplet will be placed in your account's default VPC for the region.
	// +kubebuilder:validation:Optional
	PrivateNetworking *bool `json:"privateNetworking,omitempty" tf:"private_networking,omitempty"`

	// The region where the Droplet will be created.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Boolean controlling whether to increase the disk
	// size when resizing a Droplet. It defaults to true. When set to false,
	// only the Droplet's RAM and CPU will be resized. Increasing a Droplet's disk
	// size is a permanent change. Increasing only RAM and CPU is reversible.
	// +kubebuilder:validation:Optional
	ResizeDisk *bool `json:"resizeDisk,omitempty" tf:"resize_disk,omitempty"`

	// A list of SSH key IDs or fingerprints to enable in
	// the format [12345, 123456]. To retrieve this info, use the
	// DigitalOcean API
	// or CLI (doctl compute ssh-key list). Once a Droplet is created keys can not
	// be added or removed via this provider. Modifying this field will prompt you
	// to destroy and recreate the Droplet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-digitalocean/apis/ssh/v1alpha1.Key
	// +kubebuilder:validation:Optional
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// References to Key in ssh to populate sshKeys.
	// +kubebuilder:validation:Optional
	SSHKeysRefs []v1.Reference `json:"sshKeysRefs,omitempty" tf:"-"`

	// Selector for a list of Key in ssh to populate sshKeys.
	// +kubebuilder:validation:Optional
	SSHKeysSelector *v1.Selector `json:"sshKeysSelector,omitempty" tf:"-"`

	// The unique slug that indentifies the type of Droplet. You can find a list of available slugs on DigitalOcean API documentation.
	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// A list of the tags to be applied to this Droplet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-digitalocean/apis/digitalocean/v1alpha1.Tag
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// References to Tag in digitalocean to populate tags.
	// +kubebuilder:validation:Optional
	TagsRefs []v1.Reference `json:"tagsRefs,omitempty" tf:"-"`

	// Selector for a list of Tag in digitalocean to populate tags.
	// +kubebuilder:validation:Optional
	TagsSelector *v1.Selector `json:"tagsSelector,omitempty" tf:"-"`

	// A string of the desired User Data for the Droplet.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The ID of the VPC where the Droplet will be located.
	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	VPCUUID *string `json:"vpcUuid,omitempty" tf:"vpc_uuid,omitempty"`

	// Reference to a VPC to populate vpcUuid.
	// +kubebuilder:validation:Optional
	VPCUUIDRef *v1.Reference `json:"vpcUuidRef,omitempty" tf:"-"`

	// Selector for a VPC to populate vpcUuid.
	// +kubebuilder:validation:Optional
	VPCUUIDSelector *v1.Selector `json:"vpcUuidSelector,omitempty" tf:"-"`

	// A list of the IDs of each block storage volume to be attached to the Droplet.
	// +kubebuilder:validation:Optional
	VolumeIds []*string `json:"volumeIds,omitempty" tf:"volume_ids,omitempty"`
}

// DropletSpec defines the desired state of Droplet
type DropletSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DropletParameters `json:"forProvider"`
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
	InitProvider DropletInitParameters `json:"initProvider,omitempty"`
}

// DropletStatus defines the observed state of Droplet.
type DropletStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DropletObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Droplet is the Schema for the Droplets API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type Droplet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.size) || (has(self.initProvider) && has(self.initProvider.size))",message="spec.forProvider.size is a required parameter"
	Spec   DropletSpec   `json:"spec"`
	Status DropletStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DropletList contains a list of Droplets
type DropletList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Droplet `json:"items"`
}

// Repository type metadata.
var (
	Droplet_Kind             = "Droplet"
	Droplet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Droplet_Kind}.String()
	Droplet_KindAPIVersion   = Droplet_Kind + "." + CRDGroupVersion.String()
	Droplet_GroupVersionKind = CRDGroupVersion.WithKind(Droplet_Kind)
)

func init() {
	SchemeBuilder.Register(&Droplet{}, &DropletList{})
}
