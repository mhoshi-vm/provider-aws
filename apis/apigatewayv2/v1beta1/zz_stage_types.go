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

type AccessLogSettingsInitParameters struct {

	// ARN of the CloudWatch Logs log group to receive access logs. Any trailing :* is trimmed from the ARN.
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// Single line format of the access logs of data. Refer to log settings for HTTP or Websocket.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`
}

type AccessLogSettingsObservation struct {

	// ARN of the CloudWatch Logs log group to receive access logs. Any trailing :* is trimmed from the ARN.
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// Single line format of the access logs of data. Refer to log settings for HTTP or Websocket.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`
}

type AccessLogSettingsParameters struct {

	// ARN of the CloudWatch Logs log group to receive access logs. Any trailing :* is trimmed from the ARN.
	// +kubebuilder:validation:Optional
	DestinationArn *string `json:"destinationArn" tf:"destination_arn,omitempty"`

	// Single line format of the access logs of data. Refer to log settings for HTTP or Websocket.
	// +kubebuilder:validation:Optional
	Format *string `json:"format" tf:"format,omitempty"`
}

type DefaultRouteSettingsInitParameters struct {

	// Whether data trace logging is enabled for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Defaults to false. Supported only for WebSocket APIs.
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// Whether detailed metrics are enabled for the default route. Defaults to false.
	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty" tf:"detailed_metrics_enabled,omitempty"`

	// Logging level for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Valid values: ERROR, INFO, OFF. Defaults to OFF. Supported only for WebSocket APIs.
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// Throttling burst limit for the default route.
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// Throttling rate limit for the default route.
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`
}

type DefaultRouteSettingsObservation struct {

	// Whether data trace logging is enabled for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Defaults to false. Supported only for WebSocket APIs.
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// Whether detailed metrics are enabled for the default route. Defaults to false.
	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty" tf:"detailed_metrics_enabled,omitempty"`

	// Logging level for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Valid values: ERROR, INFO, OFF. Defaults to OFF. Supported only for WebSocket APIs.
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// Throttling burst limit for the default route.
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// Throttling rate limit for the default route.
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`
}

type DefaultRouteSettingsParameters struct {

	// Whether data trace logging is enabled for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Defaults to false. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// Whether detailed metrics are enabled for the default route. Defaults to false.
	// +kubebuilder:validation:Optional
	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty" tf:"detailed_metrics_enabled,omitempty"`

	// Logging level for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Valid values: ERROR, INFO, OFF. Defaults to OFF. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// Throttling burst limit for the default route.
	// +kubebuilder:validation:Optional
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// Throttling rate limit for the default route.
	// +kubebuilder:validation:Optional
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`
}

type RouteSettingsInitParameters struct {

	// Whether data trace logging is enabled for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Defaults to false. Supported only for WebSocket APIs.
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// Whether detailed metrics are enabled for the route. Defaults to false.
	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty" tf:"detailed_metrics_enabled,omitempty"`

	// Logging level for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Valid values: ERROR, INFO, OFF. Defaults to OFF. Supported only for WebSocket APIs.
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// Route key.
	RouteKey *string `json:"routeKey,omitempty" tf:"route_key,omitempty"`

	// Throttling burst limit for the route.
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// Throttling rate limit for the route.
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`
}

type RouteSettingsObservation struct {

	// Whether data trace logging is enabled for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Defaults to false. Supported only for WebSocket APIs.
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// Whether detailed metrics are enabled for the route. Defaults to false.
	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty" tf:"detailed_metrics_enabled,omitempty"`

	// Logging level for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Valid values: ERROR, INFO, OFF. Defaults to OFF. Supported only for WebSocket APIs.
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// Route key.
	RouteKey *string `json:"routeKey,omitempty" tf:"route_key,omitempty"`

	// Throttling burst limit for the route.
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// Throttling rate limit for the route.
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`
}

type RouteSettingsParameters struct {

	// Whether data trace logging is enabled for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Defaults to false. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`

	// Whether detailed metrics are enabled for the route. Defaults to false.
	// +kubebuilder:validation:Optional
	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty" tf:"detailed_metrics_enabled,omitempty"`

	// Logging level for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
	// Valid values: ERROR, INFO, OFF. Defaults to OFF. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	LoggingLevel *string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`

	// Route key.
	// +kubebuilder:validation:Optional
	RouteKey *string `json:"routeKey" tf:"route_key,omitempty"`

	// Throttling burst limit for the route.
	// +kubebuilder:validation:Optional
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`

	// Throttling rate limit for the route.
	// +kubebuilder:validation:Optional
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`
}

type StageInitParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=API
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// Settings for logging access in this stage.
	// Use the aws_api_gateway_account resource to configure permissions for CloudWatch Logging.
	AccessLogSettings []AccessLogSettingsInitParameters `json:"accessLogSettings,omitempty" tf:"access_log_settings,omitempty"`

	// Whether updates to an API automatically trigger a new deployment. Defaults to false. Applicable for HTTP APIs.
	AutoDeploy *bool `json:"autoDeploy,omitempty" tf:"auto_deploy,omitempty"`

	// Identifier of a client certificate for the stage. Use the aws_api_gateway_client_certificate resource to configure a client certificate.
	// Supported only for WebSocket APIs.
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// Default route settings for the stage.
	DefaultRouteSettings []DefaultRouteSettingsInitParameters `json:"defaultRouteSettings,omitempty" tf:"default_route_settings,omitempty"`

	// Deployment identifier of the stage. Use the aws_apigatewayv2_deployment resource to configure a deployment.
	// +crossplane:generate:reference:type=Deployment
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Reference to a Deployment to populate deploymentId.
	// +kubebuilder:validation:Optional
	DeploymentIDRef *v1.Reference `json:"deploymentIdRef,omitempty" tf:"-"`

	// Selector for a Deployment to populate deploymentId.
	// +kubebuilder:validation:Optional
	DeploymentIDSelector *v1.Selector `json:"deploymentIdSelector,omitempty" tf:"-"`

	// Description for the stage. Must be less than or equal to 1024 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Route settings for the stage.
	RouteSettings []RouteSettingsInitParameters `json:"routeSettings,omitempty" tf:"route_settings,omitempty"`

	// Map that defines the stage variables for the stage.
	// +mapType=granular
	StageVariables map[string]*string `json:"stageVariables,omitempty" tf:"stage_variables,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StageObservation struct {

	// API identifier.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Settings for logging access in this stage.
	// Use the aws_api_gateway_account resource to configure permissions for CloudWatch Logging.
	AccessLogSettings []AccessLogSettingsObservation `json:"accessLogSettings,omitempty" tf:"access_log_settings,omitempty"`

	// ARN of the stage.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Whether updates to an API automatically trigger a new deployment. Defaults to false. Applicable for HTTP APIs.
	AutoDeploy *bool `json:"autoDeploy,omitempty" tf:"auto_deploy,omitempty"`

	// Identifier of a client certificate for the stage. Use the aws_api_gateway_client_certificate resource to configure a client certificate.
	// Supported only for WebSocket APIs.
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// Default route settings for the stage.
	DefaultRouteSettings []DefaultRouteSettingsObservation `json:"defaultRouteSettings,omitempty" tf:"default_route_settings,omitempty"`

	// Deployment identifier of the stage. Use the aws_apigatewayv2_deployment resource to configure a deployment.
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Description for the stage. Must be less than or equal to 1024 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ARN prefix to be used in an aws_lambda_permission's source_arn attribute.
	// For WebSocket APIs this attribute can additionally be used in an aws_iam_policy to authorize access to the @connections API.
	// See the Amazon API Gateway Developer Guide for details.
	ExecutionArn *string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`

	// Stage identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// URL to invoke the API pointing to the stage,
	// e.g., wss://z4675bid1j.execute-api.eu-west-2.amazonaws.com/example-stage, or https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/
	InvokeURL *string `json:"invokeUrl,omitempty" tf:"invoke_url,omitempty"`

	// Route settings for the stage.
	RouteSettings []RouteSettingsObservation `json:"routeSettings,omitempty" tf:"route_settings,omitempty"`

	// Map that defines the stage variables for the stage.
	// +mapType=granular
	StageVariables map[string]*string `json:"stageVariables,omitempty" tf:"stage_variables,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type StageParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=API
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// Settings for logging access in this stage.
	// Use the aws_api_gateway_account resource to configure permissions for CloudWatch Logging.
	// +kubebuilder:validation:Optional
	AccessLogSettings []AccessLogSettingsParameters `json:"accessLogSettings,omitempty" tf:"access_log_settings,omitempty"`

	// Whether updates to an API automatically trigger a new deployment. Defaults to false. Applicable for HTTP APIs.
	// +kubebuilder:validation:Optional
	AutoDeploy *bool `json:"autoDeploy,omitempty" tf:"auto_deploy,omitempty"`

	// Identifier of a client certificate for the stage. Use the aws_api_gateway_client_certificate resource to configure a client certificate.
	// Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// Default route settings for the stage.
	// +kubebuilder:validation:Optional
	DefaultRouteSettings []DefaultRouteSettingsParameters `json:"defaultRouteSettings,omitempty" tf:"default_route_settings,omitempty"`

	// Deployment identifier of the stage. Use the aws_apigatewayv2_deployment resource to configure a deployment.
	// +crossplane:generate:reference:type=Deployment
	// +kubebuilder:validation:Optional
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Reference to a Deployment to populate deploymentId.
	// +kubebuilder:validation:Optional
	DeploymentIDRef *v1.Reference `json:"deploymentIdRef,omitempty" tf:"-"`

	// Selector for a Deployment to populate deploymentId.
	// +kubebuilder:validation:Optional
	DeploymentIDSelector *v1.Selector `json:"deploymentIdSelector,omitempty" tf:"-"`

	// Description for the stage. Must be less than or equal to 1024 characters in length.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Route settings for the stage.
	// +kubebuilder:validation:Optional
	RouteSettings []RouteSettingsParameters `json:"routeSettings,omitempty" tf:"route_settings,omitempty"`

	// Map that defines the stage variables for the stage.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	StageVariables map[string]*string `json:"stageVariables,omitempty" tf:"stage_variables,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StageSpec defines the desired state of Stage
type StageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StageParameters `json:"forProvider"`
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
	InitProvider StageInitParameters `json:"initProvider,omitempty"`
}

// StageStatus defines the observed state of Stage.
type StageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stage is the Schema for the Stages API. Manages an Amazon API Gateway Version 2 stage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Stage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StageSpec   `json:"spec"`
	Status            StageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StageList contains a list of Stages
type StageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stage `json:"items"`
}

// Repository type metadata.
var (
	Stage_Kind             = "Stage"
	Stage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stage_Kind}.String()
	Stage_KindAPIVersion   = Stage_Kind + "." + CRDGroupVersion.String()
	Stage_GroupVersionKind = CRDGroupVersion.WithKind(Stage_Kind)
)

func init() {
	SchemeBuilder.Register(&Stage{}, &StageList{})
}
