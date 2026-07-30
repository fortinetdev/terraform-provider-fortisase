// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointOnNetRules keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointOnNetRules{}

func newResourceEndpointOnNetRules() resource.Resource {
	return &resourceEndpointOnNetRules{
		resourceEndpointOnNetRule: &resourceEndpointOnNetRule{},
	}
}

type resourceEndpointOnNetRules struct {
	*resourceEndpointOnNetRule
}

func (r *resourceEndpointOnNetRules) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_on_net_rules"
}

func (r *resourceEndpointOnNetRules) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointOnNetRule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_on_net_rules is deprecated. Please use fortisase_endpoint_on_net_rule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointOnNetRules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointOnNetRule.Configure(ctx, req, resp)
	r.resourceEndpointOnNetRule.resourceName = "fortisase_endpoint_on_net_rules"
}
func (r *resourceEndpointOnNetRules) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
