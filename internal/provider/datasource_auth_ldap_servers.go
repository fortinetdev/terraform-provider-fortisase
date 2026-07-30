// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceAuthLdapServers keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceAuthLdapServers{}

func newDatasourceAuthLdapServers() datasource.DataSource {
	return &datasourceAuthLdapServers{
		datasourceAuthLdapServer: &datasourceAuthLdapServer{},
	}
}

type datasourceAuthLdapServers struct {
	*datasourceAuthLdapServer
}

func (r *datasourceAuthLdapServers) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_ldap_servers"
}

func (r *datasourceAuthLdapServers) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceAuthLdapServer.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_auth_ldap_servers is deprecated. Please use fortisase_auth_ldap_server instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceAuthLdapServers) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceAuthLdapServer.Configure(ctx, req, resp)
	r.datasourceAuthLdapServer.resourceName = "fortisase_auth_ldap_servers"
}
