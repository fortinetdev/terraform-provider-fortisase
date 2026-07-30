// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityCertLocalCaCerts keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityCertLocalCaCerts{}

func newDatasourceSecurityCertLocalCaCerts() datasource.DataSource {
	return &datasourceSecurityCertLocalCaCerts{
		datasourceSecurityCertLocalCaCert: &datasourceSecurityCertLocalCaCert{},
	}
}

type datasourceSecurityCertLocalCaCerts struct {
	*datasourceSecurityCertLocalCaCert
}

func (r *datasourceSecurityCertLocalCaCerts) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_local_ca_certs"
}

func (r *datasourceSecurityCertLocalCaCerts) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityCertLocalCaCert.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_cert_local_ca_certs is deprecated. Please use fortisase_security_cert_local_ca_cert instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityCertLocalCaCerts) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityCertLocalCaCert.Configure(ctx, req, resp)
	r.datasourceSecurityCertLocalCaCert.resourceName = "fortisase_security_cert_local_ca_certs"
}
