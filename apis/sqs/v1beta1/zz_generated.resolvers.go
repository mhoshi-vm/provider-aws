/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	errors "github.com/pkg/errors"
	common "github.com/upbound/provider-aws/config/common"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *QueuePolicy) ResolveReferences( // ResolveReferences of this QueuePolicy.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sqs.aws.upbound.io", "v1beta1", "Queue", "QueueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QueueURL),
			Extract:      common.TerraformID(),
			Reference:    mg.Spec.ForProvider.QueueURLRef,
			Selector:     mg.Spec.ForProvider.QueueURLSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.QueueURL")
	}
	mg.Spec.ForProvider.QueueURL = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QueueURLRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("sqs.aws.upbound.io", "v1beta1", "Queue", "QueueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.QueueURL),
			Extract:      common.TerraformID(),
			Reference:    mg.Spec.InitProvider.QueueURLRef,
			Selector:     mg.Spec.InitProvider.QueueURLSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.QueueURL")
	}
	mg.Spec.InitProvider.QueueURL = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.QueueURLRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QueueRedriveAllowPolicy.
func (mg *QueueRedriveAllowPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sqs.aws.upbound.io", "v1beta1", "Queue", "QueueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QueueURL),
			Extract:      common.TerraformID(),
			Reference:    mg.Spec.ForProvider.QueueURLRef,
			Selector:     mg.Spec.ForProvider.QueueURLSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.QueueURL")
	}
	mg.Spec.ForProvider.QueueURL = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QueueURLRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("sqs.aws.upbound.io", "v1beta1", "Queue", "QueueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.QueueURL),
			Extract:      common.TerraformID(),
			Reference:    mg.Spec.InitProvider.QueueURLRef,
			Selector:     mg.Spec.InitProvider.QueueURLSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.QueueURL")
	}
	mg.Spec.InitProvider.QueueURL = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.QueueURLRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QueueRedrivePolicy.
func (mg *QueueRedrivePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sqs.aws.upbound.io", "v1beta1", "Queue", "QueueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QueueURL),
			Extract:      common.TerraformID(),
			Reference:    mg.Spec.ForProvider.QueueURLRef,
			Selector:     mg.Spec.ForProvider.QueueURLSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.QueueURL")
	}
	mg.Spec.ForProvider.QueueURL = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QueueURLRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("sqs.aws.upbound.io", "v1beta1", "Queue", "QueueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.QueueURL),
			Extract:      common.TerraformID(),
			Reference:    mg.Spec.InitProvider.QueueURLRef,
			Selector:     mg.Spec.InitProvider.QueueURLSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.QueueURL")
	}
	mg.Spec.InitProvider.QueueURL = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.QueueURLRef = rsp.ResolvedReference

	return nil
}
