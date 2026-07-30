// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityPkiUsers keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityPkiUsers{}

func newDatasourceSecurityPkiUsers() datasource.DataSource {
	return &datasourceSecurityPkiUsers{
		datasourceSecurityPkiUser: &datasourceSecurityPkiUser{},
	}
}

type datasourceSecurityPkiUsers struct {
	*datasourceSecurityPkiUser
}

func (r *datasourceSecurityPkiUsers) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_pki_users"
}

func (r *datasourceSecurityPkiUsers) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityPkiUser.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_pki_users is deprecated. Please use fortisase_security_pki_user instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityPkiUsers) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityPkiUser.Configure(ctx, req, resp)
	r.datasourceSecurityPkiUser.resourceName = "fortisase_security_pki_users"
}
