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

type RequestParameterInitParameters struct {

	// Request parameter key. This is a request data mapping parameter.
	RequestParameterKey *string `json:"requestParameterKey,omitempty" tf:"request_parameter_key,omitempty"`

	// Boolean whether or not the parameter is required.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

type RequestParameterObservation struct {

	// Request parameter key. This is a request data mapping parameter.
	RequestParameterKey *string `json:"requestParameterKey,omitempty" tf:"request_parameter_key,omitempty"`

	// Boolean whether or not the parameter is required.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

type RequestParameterParameters struct {

	// Request parameter key. This is a request data mapping parameter.
	// +kubebuilder:validation:Optional
	RequestParameterKey *string `json:"requestParameterKey" tf:"request_parameter_key,omitempty"`

	// Boolean whether or not the parameter is required.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required" tf:"required,omitempty"`
}

type RouteInitParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=API
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// Boolean whether an API key is required for the route. Defaults to false. Supported only for WebSocket APIs.
	APIKeyRequired *bool `json:"apiKeyRequired,omitempty" tf:"api_key_required,omitempty"`

	// Authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
	// +listType=set
	AuthorizationScopes []*string `json:"authorizationScopes,omitempty" tf:"authorization_scopes,omitempty"`

	// Authorization type for the route.
	// For WebSocket APIs, valid values are NONE for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer.
	// For HTTP APIs, valid values are NONE for open access, JWT for using JSON Web Tokens, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer.
	// Defaults to NONE.
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// Identifier of the aws_apigatewayv2_authorizer resource to be associated with this route.
	// +crossplane:generate:reference:type=Authorizer
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Reference to a Authorizer to populate authorizerId.
	// +kubebuilder:validation:Optional
	AuthorizerIDRef *v1.Reference `json:"authorizerIdRef,omitempty" tf:"-"`

	// Selector for a Authorizer to populate authorizerId.
	// +kubebuilder:validation:Optional
	AuthorizerIDSelector *v1.Selector `json:"authorizerIdSelector,omitempty" tf:"-"`

	// The model selection expression for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty" tf:"model_selection_expression,omitempty"`

	// Operation name for the route. Must be between 1 and 64 characters in length.
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// Request models for the route. Supported only for WebSocket APIs.
	// +mapType=granular
	RequestModels map[string]*string `json:"requestModels,omitempty" tf:"request_models,omitempty"`

	// Request parameters for the route. Supported only for WebSocket APIs.
	RequestParameter []RequestParameterInitParameters `json:"requestParameter,omitempty" tf:"request_parameter,omitempty"`

	// Route key for the route. For HTTP APIs, the route key can be either $default, or a combination of an HTTP method and resource path, for example, GET /pets.
	RouteKey *string `json:"routeKey,omitempty" tf:"route_key,omitempty"`

	// The route response selection expression for the route. Supported only for WebSocket APIs.
	RouteResponseSelectionExpression *string `json:"routeResponseSelectionExpression,omitempty" tf:"route_response_selection_expression,omitempty"`

	// Target for the route, of the form integrations/IntegrationID, where IntegrationID is the identifier of an aws_apigatewayv2_integration resource.
	// +crossplane:generate:reference:type=Integration
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/apis/apigatewayv2/v1beta1.IntegrationIDPrefixed()
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Reference to a Integration to populate target.
	// +kubebuilder:validation:Optional
	TargetRef *v1.Reference `json:"targetRef,omitempty" tf:"-"`

	// Selector for a Integration to populate target.
	// +kubebuilder:validation:Optional
	TargetSelector *v1.Selector `json:"targetSelector,omitempty" tf:"-"`
}

type RouteObservation struct {

	// API identifier.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Boolean whether an API key is required for the route. Defaults to false. Supported only for WebSocket APIs.
	APIKeyRequired *bool `json:"apiKeyRequired,omitempty" tf:"api_key_required,omitempty"`

	// Authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
	// +listType=set
	AuthorizationScopes []*string `json:"authorizationScopes,omitempty" tf:"authorization_scopes,omitempty"`

	// Authorization type for the route.
	// For WebSocket APIs, valid values are NONE for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer.
	// For HTTP APIs, valid values are NONE for open access, JWT for using JSON Web Tokens, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer.
	// Defaults to NONE.
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// Identifier of the aws_apigatewayv2_authorizer resource to be associated with this route.
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Route identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The model selection expression for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty" tf:"model_selection_expression,omitempty"`

	// Operation name for the route. Must be between 1 and 64 characters in length.
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// Request models for the route. Supported only for WebSocket APIs.
	// +mapType=granular
	RequestModels map[string]*string `json:"requestModels,omitempty" tf:"request_models,omitempty"`

	// Request parameters for the route. Supported only for WebSocket APIs.
	RequestParameter []RequestParameterObservation `json:"requestParameter,omitempty" tf:"request_parameter,omitempty"`

	// Route key for the route. For HTTP APIs, the route key can be either $default, or a combination of an HTTP method and resource path, for example, GET /pets.
	RouteKey *string `json:"routeKey,omitempty" tf:"route_key,omitempty"`

	// The route response selection expression for the route. Supported only for WebSocket APIs.
	RouteResponseSelectionExpression *string `json:"routeResponseSelectionExpression,omitempty" tf:"route_response_selection_expression,omitempty"`

	// Target for the route, of the form integrations/IntegrationID, where IntegrationID is the identifier of an aws_apigatewayv2_integration resource.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type RouteParameters struct {

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

	// Boolean whether an API key is required for the route. Defaults to false. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	APIKeyRequired *bool `json:"apiKeyRequired,omitempty" tf:"api_key_required,omitempty"`

	// Authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
	// +kubebuilder:validation:Optional
	// +listType=set
	AuthorizationScopes []*string `json:"authorizationScopes,omitempty" tf:"authorization_scopes,omitempty"`

	// Authorization type for the route.
	// For WebSocket APIs, valid values are NONE for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer.
	// For HTTP APIs, valid values are NONE for open access, JWT for using JSON Web Tokens, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer.
	// Defaults to NONE.
	// +kubebuilder:validation:Optional
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// Identifier of the aws_apigatewayv2_authorizer resource to be associated with this route.
	// +crossplane:generate:reference:type=Authorizer
	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Reference to a Authorizer to populate authorizerId.
	// +kubebuilder:validation:Optional
	AuthorizerIDRef *v1.Reference `json:"authorizerIdRef,omitempty" tf:"-"`

	// Selector for a Authorizer to populate authorizerId.
	// +kubebuilder:validation:Optional
	AuthorizerIDSelector *v1.Selector `json:"authorizerIdSelector,omitempty" tf:"-"`

	// The model selection expression for the route. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty" tf:"model_selection_expression,omitempty"`

	// Operation name for the route. Must be between 1 and 64 characters in length.
	// +kubebuilder:validation:Optional
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Request models for the route. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	RequestModels map[string]*string `json:"requestModels,omitempty" tf:"request_models,omitempty"`

	// Request parameters for the route. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	RequestParameter []RequestParameterParameters `json:"requestParameter,omitempty" tf:"request_parameter,omitempty"`

	// Route key for the route. For HTTP APIs, the route key can be either $default, or a combination of an HTTP method and resource path, for example, GET /pets.
	// +kubebuilder:validation:Optional
	RouteKey *string `json:"routeKey,omitempty" tf:"route_key,omitempty"`

	// The route response selection expression for the route. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Optional
	RouteResponseSelectionExpression *string `json:"routeResponseSelectionExpression,omitempty" tf:"route_response_selection_expression,omitempty"`

	// Target for the route, of the form integrations/IntegrationID, where IntegrationID is the identifier of an aws_apigatewayv2_integration resource.
	// +crossplane:generate:reference:type=Integration
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/apis/apigatewayv2/v1beta1.IntegrationIDPrefixed()
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Reference to a Integration to populate target.
	// +kubebuilder:validation:Optional
	TargetRef *v1.Reference `json:"targetRef,omitempty" tf:"-"`

	// Selector for a Integration to populate target.
	// +kubebuilder:validation:Optional
	TargetSelector *v1.Selector `json:"targetSelector,omitempty" tf:"-"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteParameters `json:"forProvider"`
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
	InitProvider RouteInitParameters `json:"initProvider,omitempty"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API. Manages an Amazon API Gateway Version 2 route.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routeKey) || (has(self.initProvider) && has(self.initProvider.routeKey))",message="spec.forProvider.routeKey is a required parameter"
	Spec   RouteSpec   `json:"spec"`
	Status RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
