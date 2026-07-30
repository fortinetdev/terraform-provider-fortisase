// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceNetworkImplicitDnsRules keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceNetworkImplicitDnsRules{}

func newResourceNetworkImplicitDnsRules() resource.Resource {
	return &resourceNetworkImplicitDnsRules{
		resourceNetworkImplicitDnsRule: &resourceNetworkImplicitDnsRule{},
	}
}

type resourceNetworkImplicitDnsRules struct {
	*resourceNetworkImplicitDnsRule
}

func (r *resourceNetworkImplicitDnsRules) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_implicit_dns_rules"
}

func (r *resourceNetworkImplicitDnsRules) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceNetworkImplicitDnsRule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_network_implicit_dns_rules is deprecated. Please use fortisase_network_implicit_dns_rule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceNetworkImplicitDnsRules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceNetworkImplicitDnsRule.Configure(ctx, req, resp)
	r.resourceNetworkImplicitDnsRule.resourceName = "fortisase_network_implicit_dns_rules"
}
func (r *resourceNetworkImplicitDnsRules) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
