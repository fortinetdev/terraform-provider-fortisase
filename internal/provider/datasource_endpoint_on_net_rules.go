// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointOnNetRules keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointOnNetRules{}

func newDatasourceEndpointOnNetRules() datasource.DataSource {
	return &datasourceEndpointOnNetRules{
		datasourceEndpointOnNetRule: &datasourceEndpointOnNetRule{},
	}
}

type datasourceEndpointOnNetRules struct {
	*datasourceEndpointOnNetRule
}

func (r *datasourceEndpointOnNetRules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_on_net_rules"
}

func (r *datasourceEndpointOnNetRules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointOnNetRule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_on_net_rules is deprecated. Please use fortisase_endpoint_on_net_rule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointOnNetRules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointOnNetRule.Configure(ctx, req, resp)
	r.datasourceEndpointOnNetRule.resourceName = "fortisase_endpoint_on_net_rules"
}
