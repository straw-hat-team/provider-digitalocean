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

type BucketObjectInitParameters struct {

	// The canned ACL to apply. DigitalOcean supports "private" and "public-read". (Defaults to "private".)
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The name of the bucket to put the file in.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Specifies caching behavior along the request/reply chain Read w3c cache_control for further details.
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the gzipbase64 function with small text strings. For larger objects, use source to stream the content from a disk file.
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// Specifies presentational information for the object. Read w3c content_disposition for further information.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Used to trigger updates.11.11.11 or earlier).
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is false. This value should be set to true only if the bucket has S3 object lock enabled.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The name of the object once it is in the bucket.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A mapping of keys/values to provision metadata (will be automatically prefixed by x-amz-meta-, note that only lowercase label are currently supported by the AWS Go API).
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The region where the bucket resides (Defaults to nyc3)
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies a target URL for website redirect.
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type BucketObjectObservation struct {

	// The canned ACL to apply. DigitalOcean supports "private" and "public-read". (Defaults to "private".)
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The name of the bucket to put the file in.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Specifies caching behavior along the request/reply chain Read w3c cache_control for further details.
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the gzipbase64 function with small text strings. For larger objects, use source to stream the content from a disk file.
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// Specifies presentational information for the object. Read w3c content_disposition for further information.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Used to trigger updates.11.11.11 or earlier).
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is false. This value should be set to true only if the bucket has S3 object lock enabled.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the object once it is in the bucket.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A mapping of keys/values to provision metadata (will be automatically prefixed by x-amz-meta-, note that only lowercase label are currently supported by the AWS Go API).
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The region where the bucket resides (Defaults to nyc3)
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`

	// Specifies a target URL for website redirect.
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type BucketObjectParameters struct {

	// The canned ACL to apply. DigitalOcean supports "private" and "public-read". (Defaults to "private".)
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The name of the bucket to put the file in.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Specifies caching behavior along the request/reply chain Read w3c cache_control for further details.
	// +kubebuilder:validation:Optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the gzipbase64 function with small text strings. For larger objects, use source to stream the content from a disk file.
	// +kubebuilder:validation:Optional
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// Specifies presentational information for the object. Read w3c content_disposition for further information.
	// +kubebuilder:validation:Optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	// +kubebuilder:validation:Optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// The language the content is in e.g. en-US or en-GB.
	// +kubebuilder:validation:Optional
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Used to trigger updates.11.11.11 or earlier).
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is false. This value should be set to true only if the bucket has S3 object lock enabled.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The name of the object once it is in the bucket.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A mapping of keys/values to provision metadata (will be automatically prefixed by x-amz-meta-, note that only lowercase label are currently supported by the AWS Go API).
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The region where the bucket resides (Defaults to nyc3)
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The path to a file that will be read and uploaded as raw bytes for the object content.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies a target URL for website redirect.
	// +kubebuilder:validation:Optional
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

// BucketObjectSpec defines the desired state of BucketObject
type BucketObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketObjectParameters `json:"forProvider"`
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
	InitProvider BucketObjectInitParameters `json:"initProvider,omitempty"`
}

// BucketObjectStatus defines the observed state of BucketObject.
type BucketObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketObject is the Schema for the BucketObjects API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type BucketObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucket) || (has(self.initProvider) && has(self.initProvider.bucket))",message="spec.forProvider.bucket is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	Spec   BucketObjectSpec   `json:"spec"`
	Status BucketObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketObjectList contains a list of BucketObjects
type BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketObject `json:"items"`
}

// Repository type metadata.
var (
	BucketObject_Kind             = "BucketObject"
	BucketObject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketObject_Kind}.String()
	BucketObject_KindAPIVersion   = BucketObject_Kind + "." + CRDGroupVersion.String()
	BucketObject_GroupVersionKind = CRDGroupVersion.WithKind(BucketObject_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketObject{}, &BucketObjectList{})
}
