// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointPoliciesClone keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointPoliciesClone{}

func newResourceEndpointPoliciesClone() resource.Resource {
	return &resourceEndpointPoliciesClone{
		resourceEndpointPolicyClone2Edl: &resourceEndpointPolicyClone2Edl{},
	}
}

type resourceEndpointPoliciesClone struct {
	*resourceEndpointPolicyClone2Edl
}

func (r *resourceEndpointPoliciesClone) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_policies_clone"
}

func (r *resourceEndpointPoliciesClone) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointPolicyClone2Edl.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_policies_clone is deprecated. Please use fortisase_endpoint_policy_clone instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointPoliciesClone) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointPolicyClone2Edl.Configure(ctx, req, resp)
	r.resourceEndpointPolicyClone2Edl.resourceName = "fortisase_endpoint_policies_clone"
}
func (r *resourceEndpointPoliciesClone) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
