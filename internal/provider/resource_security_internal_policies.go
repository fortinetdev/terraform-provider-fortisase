// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityInternalPolicies keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityInternalPolicies{}

func newResourceSecurityInternalPolicies() resource.Resource {
	return &resourceSecurityInternalPolicies{
		resourceSecurityInternalPolicy: &resourceSecurityInternalPolicy{},
	}
}

type resourceSecurityInternalPolicies struct {
	*resourceSecurityInternalPolicy
}

func (r *resourceSecurityInternalPolicies) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_policies"
}

func (r *resourceSecurityInternalPolicies) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityInternalPolicy.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_internal_policies is deprecated. Please use fortisase_security_internal_policy instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityInternalPolicies) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityInternalPolicy.Configure(ctx, req, resp)
	r.resourceSecurityInternalPolicy.resourceName = "fortisase_security_internal_policies"
}
func (r *resourceSecurityInternalPolicies) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
