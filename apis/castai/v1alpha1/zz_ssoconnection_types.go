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

type AadObservation struct {

	// Azure AD domain
	AdDomain *string `json:"adDomain,omitempty" tf:"ad_domain,omitempty"`

	// Azure AD client ID
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`
}

type AadParameters struct {

	// Azure AD domain
	// +kubebuilder:validation:Required
	AdDomain *string `json:"adDomain" tf:"ad_domain,omitempty"`

	// Azure AD client ID
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// Azure AD client secret
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`
}

type OktaObservation struct {

	// Okta client ID
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Okta domain
	OktaDomain *string `json:"oktaDomain,omitempty" tf:"okta_domain,omitempty"`
}

type OktaParameters struct {

	// Okta client ID
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// Okta client secret
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// Okta domain
	// +kubebuilder:validation:Required
	OktaDomain *string `json:"oktaDomain" tf:"okta_domain,omitempty"`
}

type SSOConnectionObservation struct {

	// Azure AD connector
	Aad []AadObservation `json:"aad,omitempty" tf:"aad,omitempty"`

	// Email domain of the connection
	EmailDomain *string `json:"emailDomain,omitempty" tf:"email_domain,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Connection name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Okta connector
	Okta []OktaObservation `json:"okta,omitempty" tf:"okta,omitempty"`
}

type SSOConnectionParameters struct {

	// Azure AD connector
	// +kubebuilder:validation:Optional
	Aad []AadParameters `json:"aad,omitempty" tf:"aad,omitempty"`

	// Email domain of the connection
	// +kubebuilder:validation:Optional
	EmailDomain *string `json:"emailDomain,omitempty" tf:"email_domain,omitempty"`

	// Connection name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Okta connector
	// +kubebuilder:validation:Optional
	Okta []OktaParameters `json:"okta,omitempty" tf:"okta,omitempty"`
}

// SSOConnectionSpec defines the desired state of SSOConnection
type SSOConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSOConnectionParameters `json:"forProvider"`
}

// SSOConnectionStatus defines the observed state of SSOConnection.
type SSOConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSOConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SSOConnection is the Schema for the SSOConnections API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type SSOConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.emailDomain)",message="emailDomain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   SSOConnectionSpec   `json:"spec"`
	Status SSOConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSOConnectionList contains a list of SSOConnections
type SSOConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSOConnection `json:"items"`
}

// Repository type metadata.
var (
	SSOConnection_Kind             = "SSOConnection"
	SSOConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSOConnection_Kind}.String()
	SSOConnection_KindAPIVersion   = SSOConnection_Kind + "." + CRDGroupVersion.String()
	SSOConnection_GroupVersionKind = CRDGroupVersion.WithKind(SSOConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&SSOConnection{}, &SSOConnectionList{})
}