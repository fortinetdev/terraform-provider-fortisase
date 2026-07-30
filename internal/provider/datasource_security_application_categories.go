// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityApplicationCategories keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityApplicationCategories{}

func newDatasourceSecurityApplicationCategories() datasource.DataSource {
	return &datasourceSecurityApplicationCategories{
		datasourceSecurityApplicationCategory: &datasourceSecurityApplicationCategory{},
	}
}

type datasourceSecurityApplicationCategories struct {
	*datasourceSecurityApplicationCategory
}

func (r *datasourceSecurityApplicationCategories) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_application_categories"
}

func (r *datasourceSecurityApplicationCategories) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityApplicationCategory.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_application_categories is deprecated. Please use fortisase_security_application_category instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityApplicationCategories) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityApplicationCategory.Configure(ctx, req, resp)
	r.datasourceSecurityApplicationCategory.resourceName = "fortisase_security_application_categories"
}
