// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceNetworkWildcardFqdnCustoms keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceNetworkWildcardFqdnCustoms{}

func newDatasourceNetworkWildcardFqdnCustoms() datasource.DataSource {
	return &datasourceNetworkWildcardFqdnCustoms{
		datasourceNetworkWildcardFqdnCustom: &datasourceNetworkWildcardFqdnCustom{},
	}
}

type datasourceNetworkWildcardFqdnCustoms struct {
	*datasourceNetworkWildcardFqdnCustom
}

func (r *datasourceNetworkWildcardFqdnCustoms) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_wildcard_fqdn_customs"
}

func (r *datasourceNetworkWildcardFqdnCustoms) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceNetworkWildcardFqdnCustom.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_network_wildcard_fqdn_customs is deprecated. Please use fortisase_network_wildcard_fqdn_custom instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceNetworkWildcardFqdnCustoms) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceNetworkWildcardFqdnCustom.Configure(ctx, req, resp)
	r.datasourceNetworkWildcardFqdnCustom.resourceName = "fortisase_network_wildcard_fqdn_customs"
}
