// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityFortiguardLocalCategories keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityFortiguardLocalCategories{}

func newDatasourceSecurityFortiguardLocalCategories() datasource.DataSource {
	return &datasourceSecurityFortiguardLocalCategories{
		datasourceSecurityFortiguardLocalCategory: &datasourceSecurityFortiguardLocalCategory{},
	}
}

type datasourceSecurityFortiguardLocalCategories struct {
	*datasourceSecurityFortiguardLocalCategory
}

func (r *datasourceSecurityFortiguardLocalCategories) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_fortiguard_local_categories"
}

func (r *datasourceSecurityFortiguardLocalCategories) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityFortiguardLocalCategory.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_fortiguard_local_categories is deprecated. Please use fortisase_security_fortiguard_local_category instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityFortiguardLocalCategories) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityFortiguardLocalCategory.Configure(ctx, req, resp)
	r.datasourceSecurityFortiguardLocalCategory.resourceName = "fortisase_security_fortiguard_local_categories"
}
