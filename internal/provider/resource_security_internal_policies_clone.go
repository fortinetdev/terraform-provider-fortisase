// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityInternalPoliciesClone keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityInternalPoliciesClone{}

func newResourceSecurityInternalPoliciesClone() resource.Resource {
	return &resourceSecurityInternalPoliciesClone{
		resourceSecurityInternalPolicyClone2Edl: &resourceSecurityInternalPolicyClone2Edl{},
	}
}

type resourceSecurityInternalPoliciesClone struct {
	*resourceSecurityInternalPolicyClone2Edl
}

func (r *resourceSecurityInternalPoliciesClone) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_policies_clone"
}

func (r *resourceSecurityInternalPoliciesClone) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityInternalPolicyClone2Edl.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_internal_policies_clone is deprecated. Please use fortisase_security_internal_policy_clone instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityInternalPoliciesClone) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityInternalPolicyClone2Edl.Configure(ctx, req, resp)
	r.resourceSecurityInternalPolicyClone2Edl.resourceName = "fortisase_security_internal_policies_clone"
}
func (r *resourceSecurityInternalPoliciesClone) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
