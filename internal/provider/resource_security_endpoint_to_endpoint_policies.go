// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityEndpointToEndpointPolicies keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityEndpointToEndpointPolicies{}

func newResourceSecurityEndpointToEndpointPolicies() resource.Resource {
	return &resourceSecurityEndpointToEndpointPolicies{
		resourceSecurityEndpointToEndpointPolicy: &resourceSecurityEndpointToEndpointPolicy{},
	}
}

type resourceSecurityEndpointToEndpointPolicies struct {
	*resourceSecurityEndpointToEndpointPolicy
}

func (r *resourceSecurityEndpointToEndpointPolicies) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_endpoint_to_endpoint_policies"
}

func (r *resourceSecurityEndpointToEndpointPolicies) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityEndpointToEndpointPolicy.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_endpoint_to_endpoint_policies is deprecated. Please use fortisase_security_endpoint_to_endpoint_policy instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityEndpointToEndpointPolicies) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityEndpointToEndpointPolicy.Configure(ctx, req, resp)
	r.resourceSecurityEndpointToEndpointPolicy.resourceName = "fortisase_security_endpoint_to_endpoint_policies"
}
func (r *resourceSecurityEndpointToEndpointPolicies) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
