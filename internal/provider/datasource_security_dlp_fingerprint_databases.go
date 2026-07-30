// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityDlpFingerprintDatabases keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityDlpFingerprintDatabases{}

func newDatasourceSecurityDlpFingerprintDatabases() datasource.DataSource {
	return &datasourceSecurityDlpFingerprintDatabases{
		datasourceSecurityDlpFingerprintDatabase: &datasourceSecurityDlpFingerprintDatabase{},
	}
}

type datasourceSecurityDlpFingerprintDatabases struct {
	*datasourceSecurityDlpFingerprintDatabase
}

func (r *datasourceSecurityDlpFingerprintDatabases) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_fingerprint_databases"
}

func (r *datasourceSecurityDlpFingerprintDatabases) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityDlpFingerprintDatabase.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_fingerprint_databases is deprecated. Please use fortisase_security_dlp_fingerprint_database instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityDlpFingerprintDatabases) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityDlpFingerprintDatabase.Configure(ctx, req, resp)
	r.datasourceSecurityDlpFingerprintDatabase.resourceName = "fortisase_security_dlp_fingerprint_databases"
}
