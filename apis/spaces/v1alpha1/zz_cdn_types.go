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

type CdnInitParameters struct {

	// Deprecated The ID of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.
	// ID of a DigitalOcean managed TLS certificate for use with custom domains
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The time to live for the CDN Endpoint, in seconds. Default is 3600 seconds.
	// The amount of time the content is cached in the CDN
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type CdnObservation struct {

	// Deprecated The ID of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.
	// ID of a DigitalOcean managed TLS certificate for use with custom domains
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The unique name of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// The date and time when the CDN Endpoint was created.
	// The date and time (ISO8601) of when the CDN endpoint was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint.
	// fully qualified domain name (FQDN) for custom subdomain, (requires certificate_id)
	CustomDomain *string `json:"customDomain,omitempty" tf:"custom_domain,omitempty"`

	// The fully qualified domain name (FQDN) from which the CDN-backed content is served.
	// fully qualified domain name (FQDN) to serve the CDN content
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// A unique ID that can be used to identify and reference a CDN Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The fully qualified domain name, (FQDN) for a Space.
	// fully qualified domain name (FQDN) for the origin server
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

	// The time to live for the CDN Endpoint, in seconds. Default is 3600 seconds.
	// The amount of time the content is cached in the CDN
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type CdnParameters struct {

	// Selector for a Bucket to populate origin.
	// +kubebuilder:validation:Optional
	BucketDomainName *v1.Selector `json:"bucketDomainName,omitempty" tf:"-"`

	// Deprecated The ID of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.
	// ID of a DigitalOcean managed TLS certificate for use with custom domains
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The unique name of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-digitalocean/apis/networking/v1alpha1.Certificate
	// +crossplane:generate:reference:selectorFieldName=Name
	// +kubebuilder:validation:Optional
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// Reference to a Certificate in networking to populate certificateName.
	// +kubebuilder:validation:Optional
	CertificateNameRef *v1.Reference `json:"certificateNameRef,omitempty" tf:"-"`

	// The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint.
	// fully qualified domain name (FQDN) for custom subdomain, (requires certificate_id)
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-digitalocean/apis/dns/v1alpha1.Domain
	// +crossplane:generate:reference:selectorFieldName=ID
	// +kubebuilder:validation:Optional
	CustomDomain *string `json:"customDomain,omitempty" tf:"custom_domain,omitempty"`

	// Reference to a Domain in dns to populate customDomain.
	// +kubebuilder:validation:Optional
	CustomDomainRef *v1.Reference `json:"customDomainRef,omitempty" tf:"-"`

	// A unique ID that can be used to identify and reference a CDN Endpoint.
	// +kubebuilder:validation:Optional
	ID *v1.Selector `json:"id,omitempty" tf:"-"`

	// Selector for a Certificate in networking to populate certificateName.
	// +kubebuilder:validation:Optional
	Name *v1.Selector `json:"name,omitempty" tf:"-"`

	// The fully qualified domain name, (FQDN) for a Space.
	// fully qualified domain name (FQDN) for the origin server
	// +crossplane:generate:reference:type=Bucket
	// +crossplane:generate:reference:selectorFieldName=BucketDomainName
	// +kubebuilder:validation:Optional
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

	// Reference to a Bucket to populate origin.
	// +kubebuilder:validation:Optional
	OriginRef *v1.Reference `json:"originRef,omitempty" tf:"-"`

	// The time to live for the CDN Endpoint, in seconds. Default is 3600 seconds.
	// The amount of time the content is cached in the CDN
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// CdnSpec defines the desired state of Cdn
type CdnSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CdnParameters `json:"forProvider"`
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
	InitProvider CdnInitParameters `json:"initProvider,omitempty"`
}

// CdnStatus defines the observed state of Cdn.
type CdnStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CdnObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cdn is the Schema for the Cdns API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type Cdn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CdnSpec   `json:"spec"`
	Status            CdnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CdnList contains a list of Cdns
type CdnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cdn `json:"items"`
}

// Repository type metadata.
var (
	Cdn_Kind             = "Cdn"
	Cdn_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cdn_Kind}.String()
	Cdn_KindAPIVersion   = Cdn_Kind + "." + CRDGroupVersion.String()
	Cdn_GroupVersionKind = CRDGroupVersion.WithKind(Cdn_Kind)
)

func init() {
	SchemeBuilder.Register(&Cdn{}, &CdnList{})
}