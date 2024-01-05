/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Flow.
func (mg *Flow) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DestinationFlowConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.DestinationFlowConfig[i3].DestinationConnectorProperties); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketName),
					Extract:      resource.ExtractParamPath("bucket", false),
					Reference:    mg.Spec.ForProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketNameRef,
					Selector:     mg.Spec.ForProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketNameSelector,
					To: reference.To{
						List:    &v1beta1.BucketPolicyList{},
						Managed: &v1beta1.BucketPolicy{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketName")
				}
				mg.Spec.ForProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketNameRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SourceFlowConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SourceFlowConfig[i3].SourceConnectorProperties); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketName),
					Extract:      resource.ExtractParamPath("bucket", false),
					Reference:    mg.Spec.ForProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketNameRef,
					Selector:     mg.Spec.ForProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketNameSelector,
					To: reference.To{
						List:    &v1beta1.BucketPolicyList{},
						Managed: &v1beta1.BucketPolicy{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketName")
				}
				mg.Spec.ForProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketNameRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.DestinationFlowConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.DestinationFlowConfig[i3].DestinationConnectorProperties); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketName),
					Extract:      resource.ExtractParamPath("bucket", false),
					Reference:    mg.Spec.InitProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketNameRef,
					Selector:     mg.Spec.InitProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketNameSelector,
					To: reference.To{
						List:    &v1beta1.BucketPolicyList{},
						Managed: &v1beta1.BucketPolicy{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketName")
				}
				mg.Spec.InitProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.DestinationFlowConfig[i3].DestinationConnectorProperties[i4].S3[i5].BucketNameRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.SourceFlowConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.SourceFlowConfig[i3].SourceConnectorProperties); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketName),
					Extract:      resource.ExtractParamPath("bucket", false),
					Reference:    mg.Spec.InitProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketNameRef,
					Selector:     mg.Spec.InitProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketNameSelector,
					To: reference.To{
						List:    &v1beta1.BucketPolicyList{},
						Managed: &v1beta1.BucketPolicy{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketName")
				}
				mg.Spec.InitProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.SourceFlowConfig[i3].SourceConnectorProperties[i4].S3[i5].BucketNameRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}
