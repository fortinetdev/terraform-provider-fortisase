// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceDemSpaApplications keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceDemSpaApplications{}

func newDatasourceDemSpaApplications() datasource.DataSource {
	return &datasourceDemSpaApplications{
		datasourceDemSpaApplication: &datasourceDemSpaApplication{},
	}
}

type datasourceDemSpaApplications struct {
	*datasourceDemSpaApplication
}

func (r *datasourceDemSpaApplications) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dem_spa_applications"
}

func (r *datasourceDemSpaApplications) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceDemSpaApplication.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_dem_spa_applications is deprecated. Please use fortisase_dem_spa_application instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceDemSpaApplications) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceDemSpaApplication.Configure(ctx, req, resp)
	r.datasourceDemSpaApplication.resourceName = "fortisase_dem_spa_applications"
}
