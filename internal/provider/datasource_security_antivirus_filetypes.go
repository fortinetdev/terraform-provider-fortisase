// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityAntivirusFiletypes keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityAntivirusFiletypes{}

func newDatasourceSecurityAntivirusFiletypes() datasource.DataSource {
	return &datasourceSecurityAntivirusFiletypes{
		datasourceSecurityAntivirusFiletype: &datasourceSecurityAntivirusFiletype{},
	}
}

type datasourceSecurityAntivirusFiletypes struct {
	*datasourceSecurityAntivirusFiletype
}

func (r *datasourceSecurityAntivirusFiletypes) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_antivirus_filetypes"
}

func (r *datasourceSecurityAntivirusFiletypes) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityAntivirusFiletype.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_antivirus_filetypes is deprecated. Please use fortisase_security_antivirus_filetype instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityAntivirusFiletypes) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityAntivirusFiletype.Configure(ctx, req, resp)
	r.datasourceSecurityAntivirusFiletype.resourceName = "fortisase_security_antivirus_filetypes"
}
