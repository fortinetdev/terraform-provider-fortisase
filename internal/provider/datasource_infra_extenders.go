// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceInfraExtenders keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceInfraExtenders{}

func newDatasourceInfraExtenders() datasource.DataSource {
	return &datasourceInfraExtenders{
		datasourceInfraExtender: &datasourceInfraExtender{},
	}
}

type datasourceInfraExtenders struct {
	*datasourceInfraExtender
}

func (r *datasourceInfraExtenders) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_extenders"
}

func (r *datasourceInfraExtenders) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceInfraExtender.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_infra_extenders is deprecated. Please use fortisase_infra_extender instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceInfraExtenders) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceInfraExtender.Configure(ctx, req, resp)
	r.datasourceInfraExtender.resourceName = "fortisase_infra_extenders"
}
