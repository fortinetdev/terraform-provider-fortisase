// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityDlpFilePatterns keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityDlpFilePatterns{}

func newDatasourceSecurityDlpFilePatterns() datasource.DataSource {
	return &datasourceSecurityDlpFilePatterns{
		datasourceSecurityDlpFilePattern: &datasourceSecurityDlpFilePattern{},
	}
}

type datasourceSecurityDlpFilePatterns struct {
	*datasourceSecurityDlpFilePattern
}

func (r *datasourceSecurityDlpFilePatterns) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_file_patterns"
}

func (r *datasourceSecurityDlpFilePatterns) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityDlpFilePattern.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_file_patterns is deprecated. Please use fortisase_security_dlp_file_pattern instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityDlpFilePatterns) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityDlpFilePattern.Configure(ctx, req, resp)
	r.datasourceSecurityDlpFilePattern.resourceName = "fortisase_security_dlp_file_patterns"
}
