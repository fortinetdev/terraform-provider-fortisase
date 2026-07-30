// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointZtnaProfiles keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointZtnaProfiles{}

func newDatasourceEndpointZtnaProfiles() datasource.DataSource {
	return &datasourceEndpointZtnaProfiles{
		datasourceEndpointZtnaProfile: &datasourceEndpointZtnaProfile{},
	}
}

type datasourceEndpointZtnaProfiles struct {
	*datasourceEndpointZtnaProfile
}

func (r *datasourceEndpointZtnaProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_profiles"
}

func (r *datasourceEndpointZtnaProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointZtnaProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_ztna_profiles is deprecated. Please use fortisase_endpoint_ztna_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointZtnaProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointZtnaProfile.Configure(ctx, req, resp)
	r.datasourceEndpointZtnaProfile.resourceName = "fortisase_endpoint_ztna_profiles"
}
