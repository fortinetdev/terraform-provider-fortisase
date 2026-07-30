// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointsGroups keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointsGroups{}

func newDatasourceEndpointsGroups() datasource.DataSource {
	return &datasourceEndpointsGroups{
		datasourceEndpointsGroup: &datasourceEndpointsGroup{},
	}
}

type datasourceEndpointsGroups struct {
	*datasourceEndpointsGroup
}

func (r *datasourceEndpointsGroups) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_groups"
}

func (r *datasourceEndpointsGroups) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointsGroup.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoints_groups is deprecated. Please use fortisase_endpoints_group instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointsGroups) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointsGroup.Configure(ctx, req, resp)
	r.datasourceEndpointsGroup.resourceName = "fortisase_endpoints_groups"
}
