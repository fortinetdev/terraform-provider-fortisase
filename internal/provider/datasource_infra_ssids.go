// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceInfraSsids keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceInfraSsids{}

func newDatasourceInfraSsids() datasource.DataSource {
	return &datasourceInfraSsids{
		datasourceInfraSsid: &datasourceInfraSsid{},
	}
}

type datasourceInfraSsids struct {
	*datasourceInfraSsid
}

func (r *datasourceInfraSsids) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_ssids"
}

func (r *datasourceInfraSsids) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceInfraSsid.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_infra_ssids is deprecated. Please use fortisase_infra_ssid instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceInfraSsids) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceInfraSsid.Configure(ctx, req, resp)
	r.datasourceInfraSsid.resourceName = "fortisase_infra_ssids"
}
