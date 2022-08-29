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

type ApplicationSnapshotObservation struct {

	// The current application version ID when the snapshot was created.
	ApplicationVersionID *float64 `json:"applicationVersionId,omitempty" tf:"application_version_id,omitempty"`

	// The application snapshot identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The timestamp of the application snapshot.
	SnapshotCreationTimestamp *string `json:"snapshotCreationTimestamp,omitempty" tf:"snapshot_creation_timestamp,omitempty"`
}

type ApplicationSnapshotParameters struct {

	// The name of an existing  Kinesis Analytics v2 Application. Note that the application must be running for a snapshot to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/kinesisanalyticsv2/v1beta1.Application
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	ApplicationName *string `json:"applicationName,omitempty" tf:"application_name,omitempty"`

	// Reference to a Application in kinesisanalyticsv2 to populate applicationName.
	// +kubebuilder:validation:Optional
	ApplicationNameRef *v1.Reference `json:"applicationNameRef,omitempty" tf:"-"`

	// Selector for a Application in kinesisanalyticsv2 to populate applicationName.
	// +kubebuilder:validation:Optional
	ApplicationNameSelector *v1.Selector `json:"applicationNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ApplicationSnapshotSpec defines the desired state of ApplicationSnapshot
type ApplicationSnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationSnapshotParameters `json:"forProvider"`
}

// ApplicationSnapshotStatus defines the observed state of ApplicationSnapshot.
type ApplicationSnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationSnapshot is the Schema for the ApplicationSnapshots API. Manages a Kinesis Analytics v2 Application Snapshot.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ApplicationSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSnapshotSpec   `json:"spec"`
	Status            ApplicationSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationSnapshotList contains a list of ApplicationSnapshots
type ApplicationSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationSnapshot `json:"items"`
}

// Repository type metadata.
var (
	ApplicationSnapshot_Kind             = "ApplicationSnapshot"
	ApplicationSnapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationSnapshot_Kind}.String()
	ApplicationSnapshot_KindAPIVersion   = ApplicationSnapshot_Kind + "." + CRDGroupVersion.String()
	ApplicationSnapshot_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationSnapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationSnapshot{}, &ApplicationSnapshotList{})
}
