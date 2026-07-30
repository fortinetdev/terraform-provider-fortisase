// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointFssoProfiles keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointFssoProfiles{}

func newDatasourceEndpointFssoProfiles() datasource.DataSource {
	return &datasourceEndpointFssoProfiles{
		datasourceEndpointFssoProfile: &datasourceEndpointFssoProfile{},
	}
}

type datasourceEndpointFssoProfiles struct {
	*datasourceEndpointFssoProfile
}

func (r *datasourceEndpointFssoProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_fsso_profiles"
}

func (r *datasourceEndpointFssoProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointFssoProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_fsso_profiles is deprecated. Please use fortisase_endpoint_fsso_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointFssoProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointFssoProfile.Configure(ctx, req, resp)
	r.datasourceEndpointFssoProfile.resourceName = "fortisase_endpoint_fsso_profiles"
}
