// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityDlpDataTypes keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityDlpDataTypes{}

func newDatasourceSecurityDlpDataTypes() datasource.DataSource {
	return &datasourceSecurityDlpDataTypes{
		datasourceSecurityDlpDataType: &datasourceSecurityDlpDataType{},
	}
}

type datasourceSecurityDlpDataTypes struct {
	*datasourceSecurityDlpDataType
}

func (r *datasourceSecurityDlpDataTypes) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_data_types"
}

func (r *datasourceSecurityDlpDataTypes) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityDlpDataType.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_data_types is deprecated. Please use fortisase_security_dlp_data_type instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityDlpDataTypes) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityDlpDataType.Configure(ctx, req, resp)
	r.datasourceSecurityDlpDataType.resourceName = "fortisase_security_dlp_data_types"
}
