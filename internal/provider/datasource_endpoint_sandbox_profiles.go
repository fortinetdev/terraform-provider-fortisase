// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointSandboxProfiles keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointSandboxProfiles{}

func newDatasourceEndpointSandboxProfiles() datasource.DataSource {
	return &datasourceEndpointSandboxProfiles{
		datasourceEndpointSandboxProfile: &datasourceEndpointSandboxProfile{},
	}
}

type datasourceEndpointSandboxProfiles struct {
	*datasourceEndpointSandboxProfile
}

func (r *datasourceEndpointSandboxProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_sandbox_profiles"
}

func (r *datasourceEndpointSandboxProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointSandboxProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_sandbox_profiles is deprecated. Please use fortisase_endpoint_sandbox_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointSandboxProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointSandboxProfile.Configure(ctx, req, resp)
	r.datasourceEndpointSandboxProfile.resourceName = "fortisase_endpoint_sandbox_profiles"
}
