// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityOutboundPolicies keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityOutboundPolicies{}

func newResourceSecurityOutboundPolicies() resource.Resource {
	return &resourceSecurityOutboundPolicies{
		resourceSecurityOutboundPolicy: &resourceSecurityOutboundPolicy{},
	}
}

type resourceSecurityOutboundPolicies struct {
	*resourceSecurityOutboundPolicy
}

func (r *resourceSecurityOutboundPolicies) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_outbound_policies"
}

func (r *resourceSecurityOutboundPolicies) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityOutboundPolicy.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_outbound_policies is deprecated. Please use fortisase_security_outbound_policy instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityOutboundPolicies) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityOutboundPolicy.Configure(ctx, req, resp)
	r.resourceSecurityOutboundPolicy.resourceName = "fortisase_security_outbound_policies"
}
func (r *resourceSecurityOutboundPolicies) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
