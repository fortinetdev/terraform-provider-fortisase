// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointConnectionProfiles keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointConnectionProfiles{}

func newResourceEndpointConnectionProfiles() resource.Resource {
	return &resourceEndpointConnectionProfiles{
		resourceEndpointConnectionProfile: &resourceEndpointConnectionProfile{},
	}
}

type resourceEndpointConnectionProfiles struct {
	*resourceEndpointConnectionProfile
}

func (r *resourceEndpointConnectionProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_connection_profiles"
}

func (r *resourceEndpointConnectionProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointConnectionProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_connection_profiles is deprecated. Please use fortisase_endpoint_connection_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointConnectionProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointConnectionProfile.Configure(ctx, req, resp)
	r.resourceEndpointConnectionProfile.resourceName = "fortisase_endpoint_connection_profiles"
}
func (r *resourceEndpointConnectionProfiles) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
