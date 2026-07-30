// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityPkiUsers keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityPkiUsers{}

func newResourceSecurityPkiUsers() resource.Resource {
	return &resourceSecurityPkiUsers{
		resourceSecurityPkiUser: &resourceSecurityPkiUser{},
	}
}

type resourceSecurityPkiUsers struct {
	*resourceSecurityPkiUser
}

func (r *resourceSecurityPkiUsers) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_pki_users"
}

func (r *resourceSecurityPkiUsers) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityPkiUser.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_pki_users is deprecated. Please use fortisase_security_pki_user instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityPkiUsers) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityPkiUser.Configure(ctx, req, resp)
	r.resourceSecurityPkiUser.resourceName = "fortisase_security_pki_users"
}
func (r *resourceSecurityPkiUsers) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
