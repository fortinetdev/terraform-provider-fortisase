// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceInfraFortigates keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceInfraFortigates{}

func newDatasourceInfraFortigates() datasource.DataSource {
	return &datasourceInfraFortigates{
		datasourceInfraFortigate: &datasourceInfraFortigate{},
	}
}

type datasourceInfraFortigates struct {
	*datasourceInfraFortigate
}

func (r *datasourceInfraFortigates) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_fortigates"
}

func (r *datasourceInfraFortigates) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceInfraFortigate.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_infra_fortigates is deprecated. Please use fortisase_infra_fortigate instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceInfraFortigates) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceInfraFortigate.Configure(ctx, req, resp)
	r.datasourceInfraFortigate.resourceName = "fortisase_infra_fortigates"
}
