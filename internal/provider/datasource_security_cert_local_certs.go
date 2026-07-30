// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityCertLocalCerts keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityCertLocalCerts{}

func newDatasourceSecurityCertLocalCerts() datasource.DataSource {
	return &datasourceSecurityCertLocalCerts{
		datasourceSecurityCertLocalCert: &datasourceSecurityCertLocalCert{},
	}
}

type datasourceSecurityCertLocalCerts struct {
	*datasourceSecurityCertLocalCert
}

func (r *datasourceSecurityCertLocalCerts) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_local_certs"
}

func (r *datasourceSecurityCertLocalCerts) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityCertLocalCert.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_cert_local_certs is deprecated. Please use fortisase_security_cert_local_cert instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityCertLocalCerts) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityCertLocalCert.Configure(ctx, req, resp)
	r.datasourceSecurityCertLocalCert.resourceName = "fortisase_security_cert_local_certs"
}
