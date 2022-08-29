/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EncryptionConfigurationObservation struct {
}

type EncryptionConfigurationParameters struct {

	// The encryption type to use for the repository. Valid values are AES256 or KMS. Defaults to AES256.
	// +kubebuilder:validation:Optional
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// The ARN of the KMS key to use when encryption_type is KMS. If not specified, uses the default AWS managed key for ECR.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// Reference to a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeyRef *v1.Reference `json:"kmsKeyRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeySelector *v1.Selector `json:"kmsKeySelector,omitempty" tf:"-"`
}

type ImageScanningConfigurationObservation struct {
}

type ImageScanningConfigurationParameters struct {

	// Indicates whether images are scanned after being pushed to the repository  or not scanned .
	// +kubebuilder:validation:Required
	ScanOnPush *bool `json:"scanOnPush" tf:"scan_on_push,omitempty"`
}

type RepositoryObservation struct {

	// Full ARN of the repository.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The registry ID where the repository was created.
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// The URL of the repository .
	RepositoryURL *string `json:"repositoryUrl,omitempty" tf:"repository_url,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RepositoryParameters struct {

	// Encryption configuration for the repository. See below for schema.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the ECR User Guide for more information about image scanning.
	// +kubebuilder:validation:Optional
	ImageScanningConfiguration []ImageScanningConfigurationParameters `json:"imageScanningConfiguration,omitempty" tf:"image_scanning_configuration,omitempty"`

	// The tag mutability setting for the repository. Must be one of: MUTABLE or IMMUTABLE. Defaults to MUTABLE.
	// +kubebuilder:validation:Optional
	ImageTagMutability *string `json:"imageTagMutability,omitempty" tf:"image_tag_mutability,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Repository is the Schema for the Repositorys API. Provides an Elastic Container Registry Repository.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositorySpec   `json:"spec"`
	Status            RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	Repository_Kind             = "Repository"
	Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Repository_Kind}.String()
	Repository_KindAPIVersion   = Repository_Kind + "." + CRDGroupVersion.String()
	Repository_GroupVersionKind = CRDGroupVersion.WithKind(Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
