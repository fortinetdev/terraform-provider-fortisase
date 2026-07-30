// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceNetworkHostGroups keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceNetworkHostGroups{}

func newDatasourceNetworkHostGroups() datasource.DataSource {
	return &datasourceNetworkHostGroups{
		datasourceNetworkHostGroup: &datasourceNetworkHostGroup{},
	}
}

type datasourceNetworkHostGroups struct {
	*datasourceNetworkHostGroup
}

func (r *datasourceNetworkHostGroups) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_host_groups"
}

func (r *datasourceNetworkHostGroups) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceNetworkHostGroup.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_network_host_groups is deprecated. Please use fortisase_network_host_group instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceNetworkHostGroups) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceNetworkHostGroup.Configure(ctx, req, resp)
	r.datasourceNetworkHostGroup.resourceName = "fortisase_network_host_groups"
}
