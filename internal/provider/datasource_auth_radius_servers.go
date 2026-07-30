// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceAuthRadiusServers keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceAuthRadiusServers{}

func newDatasourceAuthRadiusServers() datasource.DataSource {
	return &datasourceAuthRadiusServers{
		datasourceAuthRadiusServer: &datasourceAuthRadiusServer{},
	}
}

type datasourceAuthRadiusServers struct {
	*datasourceAuthRadiusServer
}

func (r *datasourceAuthRadiusServers) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_radius_servers"
}

func (r *datasourceAuthRadiusServers) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceAuthRadiusServer.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_auth_radius_servers is deprecated. Please use fortisase_auth_radius_server instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceAuthRadiusServers) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceAuthRadiusServer.Configure(ctx, req, resp)
	r.datasourceAuthRadiusServer.resourceName = "fortisase_auth_radius_servers"
}
