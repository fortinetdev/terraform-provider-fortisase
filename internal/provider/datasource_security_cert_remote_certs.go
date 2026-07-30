// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityCertRemoteCerts keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityCertRemoteCerts{}

func newDatasourceSecurityCertRemoteCerts() datasource.DataSource {
	return &datasourceSecurityCertRemoteCerts{
		datasourceSecurityCertRemoteCert: &datasourceSecurityCertRemoteCert{},
	}
}

type datasourceSecurityCertRemoteCerts struct {
	*datasourceSecurityCertRemoteCert
}

func (r *datasourceSecurityCertRemoteCerts) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_remote_certs"
}

func (r *datasourceSecurityCertRemoteCerts) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityCertRemoteCert.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_cert_remote_certs is deprecated. Please use fortisase_security_cert_remote_cert instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityCertRemoteCerts) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityCertRemoteCert.Configure(ctx, req, resp)
	r.datasourceSecurityCertRemoteCert.resourceName = "fortisase_security_cert_remote_certs"
}
