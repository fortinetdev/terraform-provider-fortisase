// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointSandboxProfiles keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointSandboxProfiles{}

func newResourceEndpointSandboxProfiles() resource.Resource {
	return &resourceEndpointSandboxProfiles{
		resourceEndpointSandboxProfile: &resourceEndpointSandboxProfile{},
	}
}

type resourceEndpointSandboxProfiles struct {
	*resourceEndpointSandboxProfile
}

func (r *resourceEndpointSandboxProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_sandbox_profiles"
}

func (r *resourceEndpointSandboxProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointSandboxProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_sandbox_profiles is deprecated. Please use fortisase_endpoint_sandbox_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointSandboxProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointSandboxProfile.Configure(ctx, req, resp)
	r.resourceEndpointSandboxProfile.resourceName = "fortisase_endpoint_sandbox_profiles"
}
func (r *resourceEndpointSandboxProfiles) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
