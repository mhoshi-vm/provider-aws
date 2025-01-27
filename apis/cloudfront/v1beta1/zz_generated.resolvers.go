/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Distribution) ResolveReferences( // ResolveReferences of this Distribution.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.OrderedCacheBehavior); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.OrderedCacheBehavior[i3].FunctionAssociation); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "Function", "FunctionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArnRef,
					Selector:     mg.Spec.ForProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArn")
			}
			mg.Spec.ForProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.OrderedCacheBehavior); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta1", "Function", "FunctionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArn),
					Extract:      resource.ExtractParamPath("qualified_arn", true),
					Reference:    mg.Spec.ForProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArnRef,
					Selector:     mg.Spec.ForProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArn")
			}
			mg.Spec.ForProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Origin); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "OriginAccessControl", "OriginAccessControlList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Origin[i3].OriginAccessControlID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Origin[i3].OriginAccessControlIDRef,
				Selector:     mg.Spec.ForProvider.Origin[i3].OriginAccessControlIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Origin[i3].OriginAccessControlID")
		}
		mg.Spec.ForProvider.Origin[i3].OriginAccessControlID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Origin[i3].OriginAccessControlIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Origin); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Origin[i3].S3OriginConfig); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "OriginAccessIdentity", "OriginAccessIdentityList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentity),
					Extract:      resource.ExtractParamPath("cloudfront_access_identity_path", true),
					Reference:    mg.Spec.ForProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentityRef,
					Selector:     mg.Spec.ForProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentitySelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentity")
			}
			mg.Spec.ForProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentity = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentityRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.OrderedCacheBehavior); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.OrderedCacheBehavior[i3].FunctionAssociation); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "Function", "FunctionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArnRef,
					Selector:     mg.Spec.InitProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArn")
			}
			mg.Spec.InitProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.OrderedCacheBehavior[i3].FunctionAssociation[i4].FunctionArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.OrderedCacheBehavior); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta1", "Function", "FunctionList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArn),
					Extract:      resource.ExtractParamPath("qualified_arn", true),
					Reference:    mg.Spec.InitProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArnRef,
					Selector:     mg.Spec.InitProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArn")
			}
			mg.Spec.InitProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.OrderedCacheBehavior[i3].LambdaFunctionAssociation[i4].LambdaArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Origin); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "OriginAccessControl", "OriginAccessControlList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Origin[i3].OriginAccessControlID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Origin[i3].OriginAccessControlIDRef,
				Selector:     mg.Spec.InitProvider.Origin[i3].OriginAccessControlIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Origin[i3].OriginAccessControlID")
		}
		mg.Spec.InitProvider.Origin[i3].OriginAccessControlID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Origin[i3].OriginAccessControlIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Origin); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Origin[i3].S3OriginConfig); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "OriginAccessIdentity", "OriginAccessIdentityList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentity),
					Extract:      resource.ExtractParamPath("cloudfront_access_identity_path", true),
					Reference:    mg.Spec.InitProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentityRef,
					Selector:     mg.Spec.InitProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentitySelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentity")
			}
			mg.Spec.InitProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentity = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Origin[i3].S3OriginConfig[i4].OriginAccessIdentityRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this FieldLevelEncryptionConfig.
func (mg *FieldLevelEncryptionConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.QueryArgProfileConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.QueryArgProfileConfig[i3].QueryArgProfiles); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "FieldLevelEncryptionProfile", "FieldLevelEncryptionProfileList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileID),
						Extract:      resource.ExtractResourceID(),
						Reference:    mg.Spec.ForProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileIDRef,
						Selector:     mg.Spec.ForProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileIDSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileID")
				}
				mg.Spec.ForProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.QueryArgProfileConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.QueryArgProfileConfig[i3].QueryArgProfiles); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "FieldLevelEncryptionProfile", "FieldLevelEncryptionProfileList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileID),
						Extract:      resource.ExtractResourceID(),
						Reference:    mg.Spec.InitProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileIDRef,
						Selector:     mg.Spec.InitProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileIDSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileID")
				}
				mg.Spec.InitProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.QueryArgProfileConfig[i3].QueryArgProfiles[i4].Items[i5].ProfileIDRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}

// ResolveReferences of this FieldLevelEncryptionProfile.
func (mg *FieldLevelEncryptionProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EncryptionEntities); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.EncryptionEntities[i3].Items); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "PublicKey", "PublicKeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionEntities[i3].Items[i4].PublicKeyID),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.EncryptionEntities[i3].Items[i4].PublicKeyIDRef,
					Selector:     mg.Spec.ForProvider.EncryptionEntities[i3].Items[i4].PublicKeyIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionEntities[i3].Items[i4].PublicKeyID")
			}
			mg.Spec.ForProvider.EncryptionEntities[i3].Items[i4].PublicKeyID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EncryptionEntities[i3].Items[i4].PublicKeyIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.EncryptionEntities); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.EncryptionEntities[i3].Items); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "PublicKey", "PublicKeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EncryptionEntities[i3].Items[i4].PublicKeyID),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.EncryptionEntities[i3].Items[i4].PublicKeyIDRef,
					Selector:     mg.Spec.InitProvider.EncryptionEntities[i3].Items[i4].PublicKeyIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EncryptionEntities[i3].Items[i4].PublicKeyID")
			}
			mg.Spec.InitProvider.EncryptionEntities[i3].Items[i4].PublicKeyID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EncryptionEntities[i3].Items[i4].PublicKeyIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this KeyGroup.
func (mg *KeyGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "PublicKey", "PublicKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Items),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.ItemRefs,
			Selector:      mg.Spec.ForProvider.ItemSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Items")
	}
	mg.Spec.ForProvider.Items = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ItemRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "PublicKey", "PublicKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Items),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.ItemRefs,
			Selector:      mg.Spec.InitProvider.ItemSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Items")
	}
	mg.Spec.InitProvider.Items = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.ItemRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this MonitoringSubscription.
func (mg *MonitoringSubscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "Distribution", "DistributionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DistributionID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DistributionIDRef,
			Selector:     mg.Spec.ForProvider.DistributionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DistributionID")
	}
	mg.Spec.ForProvider.DistributionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DistributionIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudfront.aws.upbound.io", "v1beta1", "Distribution", "DistributionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DistributionID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DistributionIDRef,
			Selector:     mg.Spec.InitProvider.DistributionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DistributionID")
	}
	mg.Spec.InitProvider.DistributionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DistributionIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RealtimeLogConfig.
func (mg *RealtimeLogConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Endpoint); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArnRef,
					Selector:     mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArn")
			}
			mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Endpoint); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("kinesis.aws.upbound.io", "v1beta1", "Stream", "StreamList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArn),
					Extract:      common.TerraformID(),
					Reference:    mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArnRef,
					Selector:     mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArn")
			}
			mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Endpoint); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArnRef,
					Selector:     mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArn")
			}
			mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Endpoint); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("kinesis.aws.upbound.io", "v1beta1", "Stream", "StreamList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArn),
					Extract:      common.TerraformID(),
					Reference:    mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArnRef,
					Selector:     mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArn")
			}
			mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Endpoint[i3].KinesisStreamConfig[i4].StreamArnRef = rsp.ResolvedReference

		}
	}

	return nil
}
