// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointConnectionProfiles keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointConnectionProfiles{}

func newDatasourceEndpointConnectionProfiles() datasource.DataSource {
	return &datasourceEndpointConnectionProfiles{
		datasourceEndpointConnectionProfile: &datasourceEndpointConnectionProfile{},
	}
}

type datasourceEndpointConnectionProfiles struct {
	*datasourceEndpointConnectionProfile
}

func (r *datasourceEndpointConnectionProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_connection_profiles"
}

func (r *datasourceEndpointConnectionProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointConnectionProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_connection_profiles is deprecated. Please use fortisase_endpoint_connection_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointConnectionProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointConnectionProfile.Configure(ctx, req, resp)
	r.datasourceEndpointConnectionProfile.resourceName = "fortisase_endpoint_connection_profiles"
}
