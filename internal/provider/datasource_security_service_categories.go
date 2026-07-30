// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityServiceCategories keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityServiceCategories{}

func newDatasourceSecurityServiceCategories() datasource.DataSource {
	return &datasourceSecurityServiceCategories{
		datasourceSecurityServiceCategory: &datasourceSecurityServiceCategory{},
	}
}

type datasourceSecurityServiceCategories struct {
	*datasourceSecurityServiceCategory
}

func (r *datasourceSecurityServiceCategories) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_service_categories"
}

func (r *datasourceSecurityServiceCategories) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityServiceCategory.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_service_categories is deprecated. Please use fortisase_security_service_category instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityServiceCategories) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityServiceCategory.Configure(ctx, req, resp)
	r.datasourceSecurityServiceCategory.resourceName = "fortisase_security_service_categories"
}
