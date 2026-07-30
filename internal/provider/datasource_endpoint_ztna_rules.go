// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointZtnaRules keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointZtnaRules{}

func newDatasourceEndpointZtnaRules() datasource.DataSource {
	return &datasourceEndpointZtnaRules{
		datasourceEndpointZtnaRule: &datasourceEndpointZtnaRule{},
	}
}

type datasourceEndpointZtnaRules struct {
	*datasourceEndpointZtnaRule
}

func (r *datasourceEndpointZtnaRules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_rules"
}

func (r *datasourceEndpointZtnaRules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointZtnaRule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_ztna_rules is deprecated. Please use fortisase_endpoint_ztna_rule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointZtnaRules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointZtnaRule.Configure(ctx, req, resp)
	r.datasourceEndpointZtnaRule.resourceName = "fortisase_endpoint_ztna_rules"
}
