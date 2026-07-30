// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceNetworkImplicitDnsRules keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceNetworkImplicitDnsRules{}

func newDatasourceNetworkImplicitDnsRules() datasource.DataSource {
	return &datasourceNetworkImplicitDnsRules{
		datasourceNetworkImplicitDnsRule: &datasourceNetworkImplicitDnsRule{},
	}
}

type datasourceNetworkImplicitDnsRules struct {
	*datasourceNetworkImplicitDnsRule
}

func (r *datasourceNetworkImplicitDnsRules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_implicit_dns_rules"
}

func (r *datasourceNetworkImplicitDnsRules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceNetworkImplicitDnsRule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_network_implicit_dns_rules is deprecated. Please use fortisase_network_implicit_dns_rule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceNetworkImplicitDnsRules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceNetworkImplicitDnsRule.Configure(ctx, req, resp)
	r.datasourceNetworkImplicitDnsRule.resourceName = "fortisase_network_implicit_dns_rules"
}
