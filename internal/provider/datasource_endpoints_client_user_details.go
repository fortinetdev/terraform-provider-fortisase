// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointsClientUserDetails keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointsClientUserDetails{}

func newDatasourceEndpointsClientUserDetails() datasource.DataSource {
	return &datasourceEndpointsClientUserDetails{
		datasourceEndpointsClientUserDetail: &datasourceEndpointsClientUserDetail{},
	}
}

type datasourceEndpointsClientUserDetails struct {
	*datasourceEndpointsClientUserDetail
}

func (r *datasourceEndpointsClientUserDetails) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_client_user_details"
}

func (r *datasourceEndpointsClientUserDetails) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointsClientUserDetail.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoints_client_user_details is deprecated. Please use fortisase_endpoints_client_user_detail instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointsClientUserDetails) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointsClientUserDetail.Configure(ctx, req, resp)
	r.datasourceEndpointsClientUserDetail.resourceName = "fortisase_endpoints_client_user_details"
}
