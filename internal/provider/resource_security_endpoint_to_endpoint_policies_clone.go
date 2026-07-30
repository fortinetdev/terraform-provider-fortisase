// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityEndpointToEndpointPoliciesClone keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityEndpointToEndpointPoliciesClone{}

func newResourceSecurityEndpointToEndpointPoliciesClone() resource.Resource {
	return &resourceSecurityEndpointToEndpointPoliciesClone{
		resourceSecurityEndpointToEndpointPolicyClone2Edl: &resourceSecurityEndpointToEndpointPolicyClone2Edl{},
	}
}

type resourceSecurityEndpointToEndpointPoliciesClone struct {
	*resourceSecurityEndpointToEndpointPolicyClone2Edl
}

func (r *resourceSecurityEndpointToEndpointPoliciesClone) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_endpoint_to_endpoint_policies_clone"
}

func (r *resourceSecurityEndpointToEndpointPoliciesClone) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityEndpointToEndpointPolicyClone2Edl.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_endpoint_to_endpoint_policies_clone is deprecated. Please use fortisase_security_endpoint_to_endpoint_policy_clone instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityEndpointToEndpointPoliciesClone) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityEndpointToEndpointPolicyClone2Edl.Configure(ctx, req, resp)
	r.resourceSecurityEndpointToEndpointPolicyClone2Edl.resourceName = "fortisase_security_endpoint_to_endpoint_policies_clone"
}
func (r *resourceSecurityEndpointToEndpointPoliciesClone) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
