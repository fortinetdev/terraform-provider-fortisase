// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityInternalReversePolicies keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityInternalReversePolicies{}

func newResourceSecurityInternalReversePolicies() resource.Resource {
	return &resourceSecurityInternalReversePolicies{
		resourceSecurityInternalReversePolicy: &resourceSecurityInternalReversePolicy{},
	}
}

type resourceSecurityInternalReversePolicies struct {
	*resourceSecurityInternalReversePolicy
}

func (r *resourceSecurityInternalReversePolicies) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_reverse_policies"
}

func (r *resourceSecurityInternalReversePolicies) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityInternalReversePolicy.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_internal_reverse_policies is deprecated. Please use fortisase_security_internal_reverse_policy instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityInternalReversePolicies) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityInternalReversePolicy.Configure(ctx, req, resp)
	r.resourceSecurityInternalReversePolicy.resourceName = "fortisase_security_internal_reverse_policies"
}
func (r *resourceSecurityInternalReversePolicies) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
