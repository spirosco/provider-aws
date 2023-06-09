/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DeploymentGroup.
func (mg *DeploymentGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AppName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AppNameRef,
		Selector:     mg.Spec.ForProvider.AppNameSelector,
		To: reference.To{
			List:    &AppList{},
			Managed: &App{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AppName")
	}
	mg.Spec.ForProvider.AppName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AppNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ServiceRoleArnRef,
		Selector:     mg.Spec.ForProvider.ServiceRoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceRoleArn")
	}
	mg.Spec.ForProvider.ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceRoleArnRef = rsp.ResolvedReference

	return nil
}
