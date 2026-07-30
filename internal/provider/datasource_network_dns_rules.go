// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceNetworkDnsRules keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceNetworkDnsRules{}

func newDatasourceNetworkDnsRules() datasource.DataSource {
	return &datasourceNetworkDnsRules{
		datasourceNetworkDnsRule: &datasourceNetworkDnsRule{},
	}
}

type datasourceNetworkDnsRules struct {
	*datasourceNetworkDnsRule
}

func (r *datasourceNetworkDnsRules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_dns_rules"
}

func (r *datasourceNetworkDnsRules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceNetworkDnsRule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_network_dns_rules is deprecated. Please use fortisase_network_dns_rule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceNetworkDnsRules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceNetworkDnsRule.Configure(ctx, req, resp)
	r.datasourceNetworkDnsRule.resourceName = "fortisase_network_dns_rules"
}
