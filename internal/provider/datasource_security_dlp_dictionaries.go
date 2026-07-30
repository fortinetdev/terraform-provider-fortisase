// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityDlpDictionaries keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityDlpDictionaries{}

func newDatasourceSecurityDlpDictionaries() datasource.DataSource {
	return &datasourceSecurityDlpDictionaries{
		datasourceSecurityDlpDictionary: &datasourceSecurityDlpDictionary{},
	}
}

type datasourceSecurityDlpDictionaries struct {
	*datasourceSecurityDlpDictionary
}

func (r *datasourceSecurityDlpDictionaries) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_dictionaries"
}

func (r *datasourceSecurityDlpDictionaries) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityDlpDictionary.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_dictionaries is deprecated. Please use fortisase_security_dlp_dictionary instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityDlpDictionaries) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityDlpDictionary.Configure(ctx, req, resp)
	r.datasourceSecurityDlpDictionary.resourceName = "fortisase_security_dlp_dictionaries"
}
