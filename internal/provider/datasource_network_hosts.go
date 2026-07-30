// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceNetworkHosts keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceNetworkHosts{}

func newDatasourceNetworkHosts() datasource.DataSource {
	return &datasourceNetworkHosts{
		datasourceNetworkHost: &datasourceNetworkHost{},
	}
}

type datasourceNetworkHosts struct {
	*datasourceNetworkHost
}

func (r *datasourceNetworkHosts) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hosts"
}

func (r *datasourceNetworkHosts) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceNetworkHost.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_network_hosts is deprecated. Please use fortisase_network_host instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceNetworkHosts) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceNetworkHost.Configure(ctx, req, resp)
	r.datasourceNetworkHost.resourceName = "fortisase_network_hosts"
}
