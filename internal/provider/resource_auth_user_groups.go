// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceAuthUserGroups keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceAuthUserGroups{}

func newResourceAuthUserGroups() resource.Resource {
	return &resourceAuthUserGroups{
		resourceAuthUserGroup: &resourceAuthUserGroup{},
	}
}

type resourceAuthUserGroups struct {
	*resourceAuthUserGroup
}

func (r *resourceAuthUserGroups) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_user_groups"
}

func (r *resourceAuthUserGroups) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceAuthUserGroup.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_auth_user_groups is deprecated. Please use fortisase_auth_user_group instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceAuthUserGroups) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceAuthUserGroup.Configure(ctx, req, resp)
	r.resourceAuthUserGroup.resourceName = "fortisase_auth_user_groups"
}
func (r *resourceAuthUserGroups) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
