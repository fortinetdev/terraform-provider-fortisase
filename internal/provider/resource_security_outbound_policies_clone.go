// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityOutboundPoliciesClone keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityOutboundPoliciesClone{}

func newResourceSecurityOutboundPoliciesClone() resource.Resource {
	return &resourceSecurityOutboundPoliciesClone{
		resourceSecurityOutboundPolicyClone2Edl: &resourceSecurityOutboundPolicyClone2Edl{},
	}
}

type resourceSecurityOutboundPoliciesClone struct {
	*resourceSecurityOutboundPolicyClone2Edl
}

func (r *resourceSecurityOutboundPoliciesClone) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_outbound_policies_clone"
}

func (r *resourceSecurityOutboundPoliciesClone) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityOutboundPolicyClone2Edl.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_outbound_policies_clone is deprecated. Please use fortisase_security_outbound_policy_clone instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityOutboundPoliciesClone) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityOutboundPolicyClone2Edl.Configure(ctx, req, resp)
	r.resourceSecurityOutboundPolicyClone2Edl.resourceName = "fortisase_security_outbound_policies_clone"
}
func (r *resourceSecurityOutboundPoliciesClone) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
