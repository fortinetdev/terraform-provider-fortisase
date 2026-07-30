// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceAuthLdapServers keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceAuthLdapServers{}

func newResourceAuthLdapServers() resource.Resource {
	return &resourceAuthLdapServers{
		resourceAuthLdapServer: &resourceAuthLdapServer{},
	}
}

type resourceAuthLdapServers struct {
	*resourceAuthLdapServer
}

func (r *resourceAuthLdapServers) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_ldap_servers"
}

func (r *resourceAuthLdapServers) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceAuthLdapServer.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_auth_ldap_servers is deprecated. Please use fortisase_auth_ldap_server instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceAuthLdapServers) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceAuthLdapServer.Configure(ctx, req, resp)
	r.resourceAuthLdapServer.resourceName = "fortisase_auth_ldap_servers"
}
func (r *resourceAuthLdapServers) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
