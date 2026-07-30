// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityAppCustomSignatures keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityAppCustomSignatures{}

func newDatasourceSecurityAppCustomSignatures() datasource.DataSource {
	return &datasourceSecurityAppCustomSignatures{
		datasourceSecurityAppCustomSignature: &datasourceSecurityAppCustomSignature{},
	}
}

type datasourceSecurityAppCustomSignatures struct {
	*datasourceSecurityAppCustomSignature
}

func (r *datasourceSecurityAppCustomSignatures) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_app_custom_signatures"
}

func (r *datasourceSecurityAppCustomSignatures) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityAppCustomSignature.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_app_custom_signatures is deprecated. Please use fortisase_security_app_custom_signature instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityAppCustomSignatures) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityAppCustomSignature.Configure(ctx, req, resp)
	r.datasourceSecurityAppCustomSignature.resourceName = "fortisase_security_app_custom_signatures"
}
