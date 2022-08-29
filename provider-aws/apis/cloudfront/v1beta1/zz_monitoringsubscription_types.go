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

type MonitoringSubscriptionMonitoringSubscriptionObservation struct {
}

type MonitoringSubscriptionMonitoringSubscriptionParameters struct {

	// A subscription configuration for additional CloudWatch metrics. See below.
	// +kubebuilder:validation:Required
	RealtimeMetricsSubscriptionConfig []RealtimeMetricsSubscriptionConfigParameters `json:"realtimeMetricsSubscriptionConfig" tf:"realtime_metrics_subscription_config,omitempty"`
}

type MonitoringSubscriptionObservation struct {

	// The ID of the CloudFront monitoring subscription, which corresponds to the distribution_id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitoringSubscriptionParameters struct {

	// The ID of the distribution that you are enabling metrics for.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/cloudfront/v1beta1.Distribution
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DistributionID *string `json:"distributionId,omitempty" tf:"distribution_id,omitempty"`

	// Reference to a Distribution in cloudfront to populate distributionId.
	// +kubebuilder:validation:Optional
	DistributionIDRef *v1.Reference `json:"distributionIdRef,omitempty" tf:"-"`

	// Selector for a Distribution in cloudfront to populate distributionId.
	// +kubebuilder:validation:Optional
	DistributionIDSelector *v1.Selector `json:"distributionIdSelector,omitempty" tf:"-"`

	// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	// +kubebuilder:validation:Required
	MonitoringSubscription []MonitoringSubscriptionMonitoringSubscriptionParameters `json:"monitoringSubscription" tf:"monitoring_subscription,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type RealtimeMetricsSubscriptionConfigObservation struct {
}

type RealtimeMetricsSubscriptionConfigParameters struct {

	// A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution. Valid values are Enabled and Disabled. See below.
	// +kubebuilder:validation:Required
	RealtimeMetricsSubscriptionStatus *string `json:"realtimeMetricsSubscriptionStatus" tf:"realtime_metrics_subscription_status,omitempty"`
}

// MonitoringSubscriptionSpec defines the desired state of MonitoringSubscription
type MonitoringSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitoringSubscriptionParameters `json:"forProvider"`
}

// MonitoringSubscriptionStatus defines the observed state of MonitoringSubscription.
type MonitoringSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitoringSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitoringSubscription is the Schema for the MonitoringSubscriptions API. Provides a CloudFront monitoring subscription resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MonitoringSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringSubscriptionSpec   `json:"spec"`
	Status            MonitoringSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitoringSubscriptionList contains a list of MonitoringSubscriptions
type MonitoringSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringSubscription `json:"items"`
}

// Repository type metadata.
var (
	MonitoringSubscription_Kind             = "MonitoringSubscription"
	MonitoringSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitoringSubscription_Kind}.String()
	MonitoringSubscription_KindAPIVersion   = MonitoringSubscription_Kind + "." + CRDGroupVersion.String()
	MonitoringSubscription_GroupVersionKind = CRDGroupVersion.WithKind(MonitoringSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitoringSubscription{}, &MonitoringSubscriptionList{})
}
