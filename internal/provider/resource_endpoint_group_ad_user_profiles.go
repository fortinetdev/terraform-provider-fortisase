// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointGroupAdUserProfiles keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointGroupAdUserProfiles{}

func newResourceEndpointGroupAdUserProfiles() resource.Resource {
	return &resourceEndpointGroupAdUserProfiles{
		resourceEndpointGroupAdUserProfile: &resourceEndpointGroupAdUserProfile{},
	}
}

type resourceEndpointGroupAdUserProfiles struct {
	*resourceEndpointGroupAdUserProfile
}

func (r *resourceEndpointGroupAdUserProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_group_ad_user_profiles"
}

func (r *resourceEndpointGroupAdUserProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointGroupAdUserProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_group_ad_user_profiles is deprecated. Please use fortisase_endpoint_group_ad_user_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointGroupAdUserProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointGroupAdUserProfile.Configure(ctx, req, resp)
	r.resourceEndpointGroupAdUserProfile.resourceName = "fortisase_endpoint_group_ad_user_profiles"
}
func (r *resourceEndpointGroupAdUserProfiles) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
