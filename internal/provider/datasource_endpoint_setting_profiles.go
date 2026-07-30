// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointSettingProfiles keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointSettingProfiles{}

func newDatasourceEndpointSettingProfiles() datasource.DataSource {
	return &datasourceEndpointSettingProfiles{
		datasourceEndpointSettingProfile: &datasourceEndpointSettingProfile{},
	}
}

type datasourceEndpointSettingProfiles struct {
	*datasourceEndpointSettingProfile
}

func (r *datasourceEndpointSettingProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_setting_profiles"
}

func (r *datasourceEndpointSettingProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointSettingProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_setting_profiles is deprecated. Please use fortisase_endpoint_setting_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointSettingProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointSettingProfile.Configure(ctx, req, resp)
	r.datasourceEndpointSettingProfile.resourceName = "fortisase_endpoint_setting_profiles"
}
