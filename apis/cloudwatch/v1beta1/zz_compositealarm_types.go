// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type CompositeAlarmInitParameters struct {

	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to true.
	ActionsEnabled *bool `json:"actionsEnabled,omitempty" tf:"actions_enabled,omitempty"`

	// The set of actions to execute when this alarm transitions to the ALARM state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +listType=set
	AlarmActions []*string `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// References to Topic in sns to populate alarmActions.
	// +kubebuilder:validation:Optional
	AlarmActionsRefs []v1.Reference `json:"alarmActionsRefs,omitempty" tf:"-"`

	// Selector for a list of Topic in sns to populate alarmActions.
	// +kubebuilder:validation:Optional
	AlarmActionsSelector *v1.Selector `json:"alarmActionsSelector,omitempty" tf:"-"`

	// The description for the composite alarm.
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see Creating a Composite Alarm. The maximum length is 10240 characters.
	AlarmRule *string `json:"alarmRule,omitempty" tf:"alarm_rule,omitempty"`

	// The set of actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	// +listType=set
	InsufficientDataActions []*string `json:"insufficientDataActions,omitempty" tf:"insufficient_data_actions,omitempty"`

	// The set of actions to execute when this alarm transitions to an OK state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +listType=set
	OkActions []*string `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// References to Topic in sns to populate okActions.
	// +kubebuilder:validation:Optional
	OkActionsRefs []v1.Reference `json:"okActionsRefs,omitempty" tf:"-"`

	// Selector for a list of Topic in sns to populate okActions.
	// +kubebuilder:validation:Optional
	OkActionsSelector *v1.Selector `json:"okActionsSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CompositeAlarmObservation struct {

	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to true.
	ActionsEnabled *bool `json:"actionsEnabled,omitempty" tf:"actions_enabled,omitempty"`

	// The set of actions to execute when this alarm transitions to the ALARM state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	// +listType=set
	AlarmActions []*string `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// The description for the composite alarm.
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see Creating a Composite Alarm. The maximum length is 10240 characters.
	AlarmRule *string `json:"alarmRule,omitempty" tf:"alarm_rule,omitempty"`

	// The ARN of the composite alarm.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the composite alarm resource, which is equivalent to its alarm_name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The set of actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	// +listType=set
	InsufficientDataActions []*string `json:"insufficientDataActions,omitempty" tf:"insufficient_data_actions,omitempty"`

	// The set of actions to execute when this alarm transitions to an OK state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	// +listType=set
	OkActions []*string `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CompositeAlarmParameters struct {

	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. Defaults to true.
	// +kubebuilder:validation:Optional
	ActionsEnabled *bool `json:"actionsEnabled,omitempty" tf:"actions_enabled,omitempty"`

	// The set of actions to execute when this alarm transitions to the ALARM state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	AlarmActions []*string `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// References to Topic in sns to populate alarmActions.
	// +kubebuilder:validation:Optional
	AlarmActionsRefs []v1.Reference `json:"alarmActionsRefs,omitempty" tf:"-"`

	// Selector for a list of Topic in sns to populate alarmActions.
	// +kubebuilder:validation:Optional
	AlarmActionsSelector *v1.Selector `json:"alarmActionsSelector,omitempty" tf:"-"`

	// The description for the composite alarm.
	// +kubebuilder:validation:Optional
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For syntax, see Creating a Composite Alarm. The maximum length is 10240 characters.
	// +kubebuilder:validation:Optional
	AlarmRule *string `json:"alarmRule,omitempty" tf:"alarm_rule,omitempty"`

	// The set of actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	// +kubebuilder:validation:Optional
	// +listType=set
	InsufficientDataActions []*string `json:"insufficientDataActions,omitempty" tf:"insufficient_data_actions,omitempty"`

	// The set of actions to execute when this alarm transitions to an OK state from any other state. Each action is specified as an ARN. Up to 5 actions are allowed.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	OkActions []*string `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// References to Topic in sns to populate okActions.
	// +kubebuilder:validation:Optional
	OkActionsRefs []v1.Reference `json:"okActionsRefs,omitempty" tf:"-"`

	// Selector for a list of Topic in sns to populate okActions.
	// +kubebuilder:validation:Optional
	OkActionsSelector *v1.Selector `json:"okActionsSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// CompositeAlarmSpec defines the desired state of CompositeAlarm
type CompositeAlarmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CompositeAlarmParameters `json:"forProvider"`
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
	InitProvider CompositeAlarmInitParameters `json:"initProvider,omitempty"`
}

// CompositeAlarmStatus defines the observed state of CompositeAlarm.
type CompositeAlarmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CompositeAlarmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CompositeAlarm is the Schema for the CompositeAlarms API. Provides a CloudWatch Composite Alarm resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CompositeAlarm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.alarmRule) || (has(self.initProvider) && has(self.initProvider.alarmRule))",message="spec.forProvider.alarmRule is a required parameter"
	Spec   CompositeAlarmSpec   `json:"spec"`
	Status CompositeAlarmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CompositeAlarmList contains a list of CompositeAlarms
type CompositeAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CompositeAlarm `json:"items"`
}

// Repository type metadata.
var (
	CompositeAlarm_Kind             = "CompositeAlarm"
	CompositeAlarm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CompositeAlarm_Kind}.String()
	CompositeAlarm_KindAPIVersion   = CompositeAlarm_Kind + "." + CRDGroupVersion.String()
	CompositeAlarm_GroupVersionKind = CRDGroupVersion.WithKind(CompositeAlarm_Kind)
)

func init() {
	SchemeBuilder.Register(&CompositeAlarm{}, &CompositeAlarmList{})
}
