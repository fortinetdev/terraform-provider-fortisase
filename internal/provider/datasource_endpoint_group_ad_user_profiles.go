// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointGroupAdUserProfiles keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointGroupAdUserProfiles{}

func newDatasourceEndpointGroupAdUserProfiles() datasource.DataSource {
	return &datasourceEndpointGroupAdUserProfiles{
		datasourceEndpointGroupAdUserProfile: &datasourceEndpointGroupAdUserProfile{},
	}
}

type datasourceEndpointGroupAdUserProfiles struct {
	*datasourceEndpointGroupAdUserProfile
}

func (r *datasourceEndpointGroupAdUserProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_group_ad_user_profiles"
}

func (r *datasourceEndpointGroupAdUserProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointGroupAdUserProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_group_ad_user_profiles is deprecated. Please use fortisase_endpoint_group_ad_user_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointGroupAdUserProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointGroupAdUserProfile.Configure(ctx, req, resp)
	r.datasourceEndpointGroupAdUserProfile.resourceName = "fortisase_endpoint_group_ad_user_profiles"
}
