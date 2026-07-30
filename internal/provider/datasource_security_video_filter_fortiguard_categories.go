// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityVideoFilterFortiguardCategories keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityVideoFilterFortiguardCategories{}

func newDatasourceSecurityVideoFilterFortiguardCategories() datasource.DataSource {
	return &datasourceSecurityVideoFilterFortiguardCategories{
		datasourceSecurityVideoFilterFortiguardCategory: &datasourceSecurityVideoFilterFortiguardCategory{},
	}
}

type datasourceSecurityVideoFilterFortiguardCategories struct {
	*datasourceSecurityVideoFilterFortiguardCategory
}

func (r *datasourceSecurityVideoFilterFortiguardCategories) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_video_filter_fortiguard_categories"
}

func (r *datasourceSecurityVideoFilterFortiguardCategories) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityVideoFilterFortiguardCategory.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_video_filter_fortiguard_categories is deprecated. Please use fortisase_security_video_filter_fortiguard_category instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityVideoFilterFortiguardCategories) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityVideoFilterFortiguardCategory.Configure(ctx, req, resp)
	r.datasourceSecurityVideoFilterFortiguardCategory.resourceName = "fortisase_security_video_filter_fortiguard_categories"
}
