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

type ActionInitParameters struct {

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. Valid values for action are ALLOW, BLOCK or COUNT. Valid values for override_action are COUNT and NONE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ActionObservation struct {

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. Valid values for action are ALLOW, BLOCK or COUNT. Valid values for override_action are COUNT and NONE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ActionParameters struct {

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. Valid values for action are ALLOW, BLOCK or COUNT. Valid values for override_action are COUNT and NONE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type DefaultActionInitParameters struct {

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a ruleE.g., ALLOW, BLOCK or COUNT
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DefaultActionObservation struct {

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a ruleE.g., ALLOW, BLOCK or COUNT
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DefaultActionParameters struct {

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a ruleE.g., ALLOW, BLOCK or COUNT
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type LoggingConfigurationInitParameters struct {

	// Amazon Resource Name (ARN) of Kinesis Firehose Delivery Stream
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta1.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",false)
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// Reference to a DeliveryStream in firehose to populate logDestination.
	// +kubebuilder:validation:Optional
	LogDestinationRef *v1.Reference `json:"logDestinationRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate logDestination.
	// +kubebuilder:validation:Optional
	LogDestinationSelector *v1.Selector `json:"logDestinationSelector,omitempty" tf:"-"`

	// Configuration block containing parts of the request that you want redacted from the logs. Detailed below.
	RedactedFields []RedactedFieldsInitParameters `json:"redactedFields,omitempty" tf:"redacted_fields,omitempty"`
}

type LoggingConfigurationObservation struct {

	// Amazon Resource Name (ARN) of Kinesis Firehose Delivery Stream
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// Configuration block containing parts of the request that you want redacted from the logs. Detailed below.
	RedactedFields []RedactedFieldsObservation `json:"redactedFields,omitempty" tf:"redacted_fields,omitempty"`
}

type LoggingConfigurationParameters struct {

	// Amazon Resource Name (ARN) of Kinesis Firehose Delivery Stream
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta1.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// Reference to a DeliveryStream in firehose to populate logDestination.
	// +kubebuilder:validation:Optional
	LogDestinationRef *v1.Reference `json:"logDestinationRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate logDestination.
	// +kubebuilder:validation:Optional
	LogDestinationSelector *v1.Selector `json:"logDestinationSelector,omitempty" tf:"-"`

	// Configuration block containing parts of the request that you want redacted from the logs. Detailed below.
	// +kubebuilder:validation:Optional
	RedactedFields []RedactedFieldsParameters `json:"redactedFields,omitempty" tf:"redacted_fields,omitempty"`
}

type OverrideActionInitParameters struct {

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. Valid values for action are ALLOW, BLOCK or COUNT. Valid values for override_action are COUNT and NONE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OverrideActionObservation struct {

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. Valid values for action are ALLOW, BLOCK or COUNT. Valid values for override_action are COUNT and NONE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OverrideActionParameters struct {

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. Valid values for action are ALLOW, BLOCK or COUNT. Valid values for override_action are COUNT and NONE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type RedactedFieldsFieldToMatchInitParameters struct {

	// When the value of type is HEADER, enter the name of the header that you want the WAF to search, for example, User-Agent or Referer. If the value of type is any other value, omit data.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. Valid values for action are ALLOW, BLOCK or COUNT. Valid values for override_action are COUNT and NONE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RedactedFieldsFieldToMatchObservation struct {

	// When the value of type is HEADER, enter the name of the header that you want the WAF to search, for example, User-Agent or Referer. If the value of type is any other value, omit data.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. Valid values for action are ALLOW, BLOCK or COUNT. Valid values for override_action are COUNT and NONE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RedactedFieldsFieldToMatchParameters struct {

	// When the value of type is HEADER, enter the name of the header that you want the WAF to search, for example, User-Agent or Referer. If the value of type is any other value, omit data.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. Valid values for action are ALLOW, BLOCK or COUNT. Valid values for override_action are COUNT and NONE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type RedactedFieldsInitParameters struct {

	// Set of configuration blocks for fields to redact. Detailed below.
	FieldToMatch []RedactedFieldsFieldToMatchInitParameters `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`
}

type RedactedFieldsObservation struct {

	// Set of configuration blocks for fields to redact. Detailed below.
	FieldToMatch []RedactedFieldsFieldToMatchObservation `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`
}

type RedactedFieldsParameters struct {

	// Set of configuration blocks for fields to redact. Detailed below.
	// +kubebuilder:validation:Optional
	FieldToMatch []RedactedFieldsFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`
}

type WebACLInitParameters struct {

	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction []DefaultActionInitParameters `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration []LoggingConfigurationInitParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the web ACL.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rule []WebACLRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WebACLObservation struct {

	// Amazon Resource Name (ARN) of the WAF Regional WebACL.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction []DefaultActionObservation `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// The ID of the WAF Regional WebACL.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration []LoggingConfigurationObservation `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the web ACL.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rule []WebACLRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WebACLParameters struct {

	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	// +kubebuilder:validation:Optional
	DefaultAction []DefaultActionParameters `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// Configuration block to enable WAF logging. Detailed below.
	// +kubebuilder:validation:Optional
	LoggingConfiguration []LoggingConfigurationParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// The name or description for the Amazon CloudWatch metric of this web ACL.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the web ACL.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	// +kubebuilder:validation:Optional
	Rule []WebACLRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WebACLRuleInitParameters struct {

	// Configuration block of the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if type is GROUP. Detailed below.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// Configuration block of the override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if type is GROUP. Detailed below.
	OverrideAction []OverrideActionInitParameters `json:"overrideAction,omitempty" tf:"override_action,omitempty"`

	// Specifies the order in which the rules in a WebACL are evaluated.
	// Rules with a lower value are evaluated before rules with a higher value.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// ID of the associated WAF (Regional) rule (e.g., aws_wafregional_rule). WAF (Global) rules cannot be used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/wafregional/v1beta1.Rule
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Reference to a Rule in wafregional to populate ruleId.
	// +kubebuilder:validation:Optional
	RuleIDRef *v1.Reference `json:"ruleIdRef,omitempty" tf:"-"`

	// Selector for a Rule in wafregional to populate ruleId.
	// +kubebuilder:validation:Optional
	RuleIDSelector *v1.Selector `json:"ruleIdSelector,omitempty" tf:"-"`

	// The rule type, either REGULAR, as defined by Rule, RATE_BASED, as defined by RateBasedRule, or GROUP, as defined by RuleGroup. The default is REGULAR. If you add a RATE_BASED rule, you need to set type as RATE_BASED. If you add a GROUP rule, you need to set type as GROUP.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type WebACLRuleObservation struct {

	// Configuration block of the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if type is GROUP. Detailed below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// Configuration block of the override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if type is GROUP. Detailed below.
	OverrideAction []OverrideActionObservation `json:"overrideAction,omitempty" tf:"override_action,omitempty"`

	// Specifies the order in which the rules in a WebACL are evaluated.
	// Rules with a lower value are evaluated before rules with a higher value.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// ID of the associated WAF (Regional) rule (e.g., aws_wafregional_rule). WAF (Global) rules cannot be used.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// The rule type, either REGULAR, as defined by Rule, RATE_BASED, as defined by RateBasedRule, or GROUP, as defined by RuleGroup. The default is REGULAR. If you add a RATE_BASED rule, you need to set type as RATE_BASED. If you add a GROUP rule, you need to set type as GROUP.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type WebACLRuleParameters struct {

	// Configuration block of the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if type is GROUP. Detailed below.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// Configuration block of the override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if type is GROUP. Detailed below.
	// +kubebuilder:validation:Optional
	OverrideAction []OverrideActionParameters `json:"overrideAction,omitempty" tf:"override_action,omitempty"`

	// Specifies the order in which the rules in a WebACL are evaluated.
	// Rules with a lower value are evaluated before rules with a higher value.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// ID of the associated WAF (Regional) rule (e.g., aws_wafregional_rule). WAF (Global) rules cannot be used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/wafregional/v1beta1.Rule
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Reference to a Rule in wafregional to populate ruleId.
	// +kubebuilder:validation:Optional
	RuleIDRef *v1.Reference `json:"ruleIdRef,omitempty" tf:"-"`

	// Selector for a Rule in wafregional to populate ruleId.
	// +kubebuilder:validation:Optional
	RuleIDSelector *v1.Selector `json:"ruleIdSelector,omitempty" tf:"-"`

	// The rule type, either REGULAR, as defined by Rule, RATE_BASED, as defined by RateBasedRule, or GROUP, as defined by RuleGroup. The default is REGULAR. If you add a RATE_BASED rule, you need to set type as RATE_BASED. If you add a GROUP rule, you need to set type as GROUP.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// WebACLSpec defines the desired state of WebACL
type WebACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebACLParameters `json:"forProvider"`
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
	InitProvider WebACLInitParameters `json:"initProvider,omitempty"`
}

// WebACLStatus defines the observed state of WebACL.
type WebACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WebACL is the Schema for the WebACLs API. Provides a AWS WAF Regional web access control group (ACL) resource for use with ALB.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WebACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultAction) || (has(self.initProvider) && has(self.initProvider.defaultAction))",message="spec.forProvider.defaultAction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metricName) || (has(self.initProvider) && has(self.initProvider.metricName))",message="spec.forProvider.metricName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   WebACLSpec   `json:"spec"`
	Status WebACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebACLList contains a list of WebACLs
type WebACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebACL `json:"items"`
}

// Repository type metadata.
var (
	WebACL_Kind             = "WebACL"
	WebACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebACL_Kind}.String()
	WebACL_KindAPIVersion   = WebACL_Kind + "." + CRDGroupVersion.String()
	WebACL_GroupVersionKind = CRDGroupVersion.WithKind(WebACL_Kind)
)

func init() {
	SchemeBuilder.Register(&WebACL{}, &WebACLList{})
}
