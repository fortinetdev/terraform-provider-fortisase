// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceAuthUsers keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceAuthUsers{}

func newResourceAuthUsers() resource.Resource {
	return &resourceAuthUsers{
		resourceAuthUser: &resourceAuthUser{},
	}
}

type resourceAuthUsers struct {
	*resourceAuthUser
}

func (r *resourceAuthUsers) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_users"
}

func (r *resourceAuthUsers) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceAuthUser.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_auth_users is deprecated. Please use fortisase_auth_user instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceAuthUsers) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceAuthUser.Configure(ctx, req, resp)
	r.resourceAuthUser.resourceName = "fortisase_auth_users"
}
func (r *resourceAuthUsers) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
