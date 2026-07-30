// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityCertRemoteCaCerts keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityCertRemoteCaCerts{}

func newDatasourceSecurityCertRemoteCaCerts() datasource.DataSource {
	return &datasourceSecurityCertRemoteCaCerts{
		datasourceSecurityCertRemoteCaCert: &datasourceSecurityCertRemoteCaCert{},
	}
}

type datasourceSecurityCertRemoteCaCerts struct {
	*datasourceSecurityCertRemoteCaCert
}

func (r *datasourceSecurityCertRemoteCaCerts) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_remote_ca_certs"
}

func (r *datasourceSecurityCertRemoteCaCerts) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityCertRemoteCaCert.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_cert_remote_ca_certs is deprecated. Please use fortisase_security_cert_remote_ca_cert instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityCertRemoteCaCerts) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityCertRemoteCaCert.Configure(ctx, req, resp)
	r.datasourceSecurityCertRemoteCaCert.resourceName = "fortisase_security_cert_remote_ca_certs"
}
