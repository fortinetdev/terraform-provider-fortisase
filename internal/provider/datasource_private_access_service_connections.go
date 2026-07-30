// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourcePrivateAccessServiceConnections keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourcePrivateAccessServiceConnections{}

func newDatasourcePrivateAccessServiceConnections() datasource.DataSource {
	return &datasourcePrivateAccessServiceConnections{
		datasourcePrivateAccessServiceConnection: &datasourcePrivateAccessServiceConnection{},
	}
}

type datasourcePrivateAccessServiceConnections struct {
	*datasourcePrivateAccessServiceConnection
}

func (r *datasourcePrivateAccessServiceConnections) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_service_connections"
}

func (r *datasourcePrivateAccessServiceConnections) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourcePrivateAccessServiceConnection.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_private_access_service_connections is deprecated. Please use fortisase_private_access_service_connection instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourcePrivateAccessServiceConnections) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourcePrivateAccessServiceConnection.Configure(ctx, req, resp)
	r.datasourcePrivateAccessServiceConnection.resourceName = "fortisase_private_access_service_connections"
}
