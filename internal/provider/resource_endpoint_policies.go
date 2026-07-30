// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointPolicies keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointPolicies{}

func newResourceEndpointPolicies() resource.Resource {
	return &resourceEndpointPolicies{
		resourceEndpointPolicy: &resourceEndpointPolicy{},
	}
}

type resourceEndpointPolicies struct {
	*resourceEndpointPolicy
}

func (r *resourceEndpointPolicies) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_policies"
}

func (r *resourceEndpointPolicies) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointPolicy.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_policies is deprecated. Please use fortisase_endpoint_policy instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointPolicies) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointPolicy.Configure(ctx, req, resp)
	r.resourceEndpointPolicy.resourceName = "fortisase_endpoint_policies"
}
func (r *resourceEndpointPolicies) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
