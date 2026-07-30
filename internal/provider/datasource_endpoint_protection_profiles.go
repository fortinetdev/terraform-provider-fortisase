// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointProtectionProfiles keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointProtectionProfiles{}

func newDatasourceEndpointProtectionProfiles() datasource.DataSource {
	return &datasourceEndpointProtectionProfiles{
		datasourceEndpointProtectionProfile: &datasourceEndpointProtectionProfile{},
	}
}

type datasourceEndpointProtectionProfiles struct {
	*datasourceEndpointProtectionProfile
}

func (r *datasourceEndpointProtectionProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_protection_profiles"
}

func (r *datasourceEndpointProtectionProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointProtectionProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_protection_profiles is deprecated. Please use fortisase_endpoint_protection_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointProtectionProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointProtectionProfile.Configure(ctx, req, resp)
	r.datasourceEndpointProtectionProfile.resourceName = "fortisase_endpoint_protection_profiles"
}
