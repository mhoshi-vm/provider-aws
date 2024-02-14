/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Stack.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Stack) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IAMRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.IAMRoleArnRef,
			Selector:     mg.Spec.ForProvider.IAMRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRoleArn")
	}
	mg.Spec.ForProvider.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IAMRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IAMRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.IAMRoleArnRef,
			Selector:     mg.Spec.InitProvider.IAMRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IAMRoleArn")
	}
	mg.Spec.InitProvider.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IAMRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StackSet.
func (mg *StackSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AdministrationRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.AdministrationRoleArnRef,
			Selector:     mg.Spec.ForProvider.AdministrationRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AdministrationRoleArn")
	}
	mg.Spec.ForProvider.AdministrationRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AdministrationRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AdministrationRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.AdministrationRoleArnRef,
			Selector:     mg.Spec.InitProvider.AdministrationRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AdministrationRoleArn")
	}
	mg.Spec.InitProvider.AdministrationRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AdministrationRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StackSetInstance.
func (mg *StackSetInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudformation.aws.upbound.io", "v1beta1", "StackSet", "StackSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StackSetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.StackSetNameRef,
			Selector:     mg.Spec.ForProvider.StackSetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StackSetName")
	}
	mg.Spec.ForProvider.StackSetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StackSetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudformation.aws.upbound.io", "v1beta1", "StackSet", "StackSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StackSetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.StackSetNameRef,
			Selector:     mg.Spec.InitProvider.StackSetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StackSetName")
	}
	mg.Spec.InitProvider.StackSetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StackSetNameRef = rsp.ResolvedReference

	return nil
}
