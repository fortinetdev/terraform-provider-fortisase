// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityFortiguardCategories keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityFortiguardCategories{}

func newDatasourceSecurityFortiguardCategories() datasource.DataSource {
	return &datasourceSecurityFortiguardCategories{
		datasourceSecurityFortiguardCategory: &datasourceSecurityFortiguardCategory{},
	}
}

type datasourceSecurityFortiguardCategories struct {
	*datasourceSecurityFortiguardCategory
}

func (r *datasourceSecurityFortiguardCategories) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_fortiguard_categories"
}

func (r *datasourceSecurityFortiguardCategories) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityFortiguardCategory.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_fortiguard_categories is deprecated. Please use fortisase_security_fortiguard_category instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityFortiguardCategories) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityFortiguardCategory.Configure(ctx, req, resp)
	r.datasourceSecurityFortiguardCategory.resourceName = "fortisase_security_fortiguard_categories"
}
