// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceDemCustomSaasApps keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceDemCustomSaasApps{}

func newDatasourceDemCustomSaasApps() datasource.DataSource {
	return &datasourceDemCustomSaasApps{
		datasourceDemCustomSaasApp: &datasourceDemCustomSaasApp{},
	}
}

type datasourceDemCustomSaasApps struct {
	*datasourceDemCustomSaasApp
}

func (r *datasourceDemCustomSaasApps) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dem_custom_saas_apps"
}

func (r *datasourceDemCustomSaasApps) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceDemCustomSaasApp.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_dem_custom_saas_apps is deprecated. Please use fortisase_dem_custom_saas_app instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceDemCustomSaasApps) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceDemCustomSaasApp.Configure(ctx, req, resp)
	r.datasourceDemCustomSaasApp.resourceName = "fortisase_dem_custom_saas_apps"
}
