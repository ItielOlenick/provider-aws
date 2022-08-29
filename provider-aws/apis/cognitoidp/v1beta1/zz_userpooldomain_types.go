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

type UserPoolDomainObservation struct {

	// The AWS account ID for the user pool owner.
	AwsAccountID *string `json:"awsAccountId,omitempty" tf:"aws_account_id,omitempty"`

	// The URL of the CloudFront distribution. This is required to generate the ALIAS aws_route53_record
	CloudfrontDistributionArn *string `json:"cloudfrontDistributionArn,omitempty" tf:"cloudfront_distribution_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The S3 bucket where the static files for this domain are stored.
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// The app version.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type UserPoolDomainParameters struct {

	// The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/acm/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Reference to a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnRef *v1.Reference `json:"certificateArnRef,omitempty" tf:"-"`

	// Selector for a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnSelector *v1.Selector `json:"certificateArnSelector,omitempty" tf:"-"`

	// For custom domains, this is the fully-qualified domain name, such as auth.example.com. For Amazon Cognito prefix domains, this is the prefix alone, such as auth.
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The user pool ID.
	// +crossplane:generate:reference:type=UserPool
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`
}

// UserPoolDomainSpec defines the desired state of UserPoolDomain
type UserPoolDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPoolDomainParameters `json:"forProvider"`
}

// UserPoolDomainStatus defines the observed state of UserPoolDomain.
type UserPoolDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPoolDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolDomain is the Schema for the UserPoolDomains API. Provides a Cognito User Pool Domain resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserPoolDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPoolDomainSpec   `json:"spec"`
	Status            UserPoolDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPoolDomainList contains a list of UserPoolDomains
type UserPoolDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPoolDomain `json:"items"`
}

// Repository type metadata.
var (
	UserPoolDomain_Kind             = "UserPoolDomain"
	UserPoolDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPoolDomain_Kind}.String()
	UserPoolDomain_KindAPIVersion   = UserPoolDomain_Kind + "." + CRDGroupVersion.String()
	UserPoolDomain_GroupVersionKind = CRDGroupVersion.WithKind(UserPoolDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&UserPoolDomain{}, &UserPoolDomainList{})
}
