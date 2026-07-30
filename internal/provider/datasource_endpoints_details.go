// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointsDetails keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointsDetails{}

func newDatasourceEndpointsDetails() datasource.DataSource {
	return &datasourceEndpointsDetails{
		datasourceEndpointsDetail: &datasourceEndpointsDetail{},
	}
}

type datasourceEndpointsDetails struct {
	*datasourceEndpointsDetail
}

func (r *datasourceEndpointsDetails) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_details"
}

func (r *datasourceEndpointsDetails) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointsDetail.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoints_details is deprecated. Please use fortisase_endpoints_detail instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointsDetails) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointsDetail.Configure(ctx, req, resp)
	r.datasourceEndpointsDetail.resourceName = "fortisase_endpoints_details"
}
