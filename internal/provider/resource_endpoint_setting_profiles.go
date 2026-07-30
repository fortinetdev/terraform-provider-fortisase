// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointSettingProfiles keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointSettingProfiles{}

func newResourceEndpointSettingProfiles() resource.Resource {
	return &resourceEndpointSettingProfiles{
		resourceEndpointSettingProfile: &resourceEndpointSettingProfile{},
	}
}

type resourceEndpointSettingProfiles struct {
	*resourceEndpointSettingProfile
}

func (r *resourceEndpointSettingProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_setting_profiles"
}

func (r *resourceEndpointSettingProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointSettingProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_setting_profiles is deprecated. Please use fortisase_endpoint_setting_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointSettingProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointSettingProfile.Configure(ctx, req, resp)
	r.resourceEndpointSettingProfile.resourceName = "fortisase_endpoint_setting_profiles"
}
func (r *resourceEndpointSettingProfiles) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
