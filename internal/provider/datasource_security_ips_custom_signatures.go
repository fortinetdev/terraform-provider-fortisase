// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityIpsCustomSignatures keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityIpsCustomSignatures{}

func newDatasourceSecurityIpsCustomSignatures() datasource.DataSource {
	return &datasourceSecurityIpsCustomSignatures{
		datasourceSecurityIpsCustomSignature: &datasourceSecurityIpsCustomSignature{},
	}
}

type datasourceSecurityIpsCustomSignatures struct {
	*datasourceSecurityIpsCustomSignature
}

func (r *datasourceSecurityIpsCustomSignatures) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ips_custom_signatures"
}

func (r *datasourceSecurityIpsCustomSignatures) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityIpsCustomSignature.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_ips_custom_signatures is deprecated. Please use fortisase_security_ips_custom_signature instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityIpsCustomSignatures) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityIpsCustomSignature.Configure(ctx, req, resp)
	r.datasourceSecurityIpsCustomSignature.resourceName = "fortisase_security_ips_custom_signatures"
}
