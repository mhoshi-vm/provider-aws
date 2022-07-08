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

type RegionSettingsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RegionSettingsParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceTypeManagementPreference map[string]*bool `json:"resourceTypeManagementPreference,omitempty" tf:"resource_type_management_preference,omitempty"`

	// +kubebuilder:validation:Required
	ResourceTypeOptInPreference map[string]*bool `json:"resourceTypeOptInPreference" tf:"resource_type_opt_in_preference,omitempty"`
}

// RegionSettingsSpec defines the desired state of RegionSettings
type RegionSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionSettingsParameters `json:"forProvider"`
}

// RegionSettingsStatus defines the observed state of RegionSettings.
type RegionSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionSettings is the Schema for the RegionSettingss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RegionSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionSettingsSpec   `json:"spec"`
	Status            RegionSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionSettingsList contains a list of RegionSettingss
type RegionSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionSettings `json:"items"`
}

// Repository type metadata.
var (
	RegionSettings_Kind             = "RegionSettings"
	RegionSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionSettings_Kind}.String()
	RegionSettings_KindAPIVersion   = RegionSettings_Kind + "." + CRDGroupVersion.String()
	RegionSettings_GroupVersionKind = CRDGroupVersion.WithKind(RegionSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionSettings{}, &RegionSettingsList{})
}
