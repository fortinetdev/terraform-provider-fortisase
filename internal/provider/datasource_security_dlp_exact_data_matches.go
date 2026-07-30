// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityDlpExactDataMatches keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityDlpExactDataMatches{}

func newDatasourceSecurityDlpExactDataMatches() datasource.DataSource {
	return &datasourceSecurityDlpExactDataMatches{
		datasourceSecurityDlpExactDataMatch: &datasourceSecurityDlpExactDataMatch{},
	}
}

type datasourceSecurityDlpExactDataMatches struct {
	*datasourceSecurityDlpExactDataMatch
}

func (r *datasourceSecurityDlpExactDataMatches) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_exact_data_matches"
}

func (r *datasourceSecurityDlpExactDataMatches) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityDlpExactDataMatch.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_exact_data_matches is deprecated. Please use fortisase_security_dlp_exact_data_match instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityDlpExactDataMatches) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityDlpExactDataMatch.Configure(ctx, req, resp)
	r.datasourceSecurityDlpExactDataMatch.resourceName = "fortisase_security_dlp_exact_data_matches"
}
