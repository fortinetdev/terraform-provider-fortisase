// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityInternalReversePoliciesClone keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityInternalReversePoliciesClone{}

func newResourceSecurityInternalReversePoliciesClone() resource.Resource {
	return &resourceSecurityInternalReversePoliciesClone{
		resourceSecurityInternalReversePolicyClone2Edl: &resourceSecurityInternalReversePolicyClone2Edl{},
	}
}

type resourceSecurityInternalReversePoliciesClone struct {
	*resourceSecurityInternalReversePolicyClone2Edl
}

func (r *resourceSecurityInternalReversePoliciesClone) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_reverse_policies_clone"
}

func (r *resourceSecurityInternalReversePoliciesClone) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityInternalReversePolicyClone2Edl.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_internal_reverse_policies_clone is deprecated. Please use fortisase_security_internal_reverse_policy_clone instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityInternalReversePoliciesClone) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityInternalReversePolicyClone2Edl.Configure(ctx, req, resp)
	r.resourceSecurityInternalReversePolicyClone2Edl.resourceName = "fortisase_security_internal_reverse_policies_clone"
}
func (r *resourceSecurityInternalReversePoliciesClone) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
