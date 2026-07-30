// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceNetworkDnsRules keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceNetworkDnsRules{}

func newResourceNetworkDnsRules() resource.Resource {
	return &resourceNetworkDnsRules{
		resourceNetworkDnsRule: &resourceNetworkDnsRule{},
	}
}

type resourceNetworkDnsRules struct {
	*resourceNetworkDnsRule
}

func (r *resourceNetworkDnsRules) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_dns_rules"
}

func (r *resourceNetworkDnsRules) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceNetworkDnsRule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_network_dns_rules is deprecated. Please use fortisase_network_dns_rule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceNetworkDnsRules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceNetworkDnsRule.Configure(ctx, req, resp)
	r.resourceNetworkDnsRule.resourceName = "fortisase_network_dns_rules"
}
func (r *resourceNetworkDnsRules) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
