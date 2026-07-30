// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointZtnaRules keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointZtnaRules{}

func newResourceEndpointZtnaRules() resource.Resource {
	return &resourceEndpointZtnaRules{
		resourceEndpointZtnaRule: &resourceEndpointZtnaRule{},
	}
}

type resourceEndpointZtnaRules struct {
	*resourceEndpointZtnaRule
}

func (r *resourceEndpointZtnaRules) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_rules"
}

func (r *resourceEndpointZtnaRules) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointZtnaRule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_ztna_rules is deprecated. Please use fortisase_endpoint_ztna_rule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointZtnaRules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointZtnaRule.Configure(ctx, req, resp)
	r.resourceEndpointZtnaRule.resourceName = "fortisase_endpoint_ztna_rules"
}
func (r *resourceEndpointZtnaRules) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
