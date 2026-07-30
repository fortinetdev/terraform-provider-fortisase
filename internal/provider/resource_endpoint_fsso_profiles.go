// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointFssoProfiles keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointFssoProfiles{}

func newResourceEndpointFssoProfiles() resource.Resource {
	return &resourceEndpointFssoProfiles{
		resourceEndpointFssoProfile: &resourceEndpointFssoProfile{},
	}
}

type resourceEndpointFssoProfiles struct {
	*resourceEndpointFssoProfile
}

func (r *resourceEndpointFssoProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_fsso_profiles"
}

func (r *resourceEndpointFssoProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointFssoProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_fsso_profiles is deprecated. Please use fortisase_endpoint_fsso_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointFssoProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointFssoProfile.Configure(ctx, req, resp)
	r.resourceEndpointFssoProfile.resourceName = "fortisase_endpoint_fsso_profiles"
}
func (r *resourceEndpointFssoProfiles) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
