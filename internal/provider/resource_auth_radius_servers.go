// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceAuthRadiusServers keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceAuthRadiusServers{}

func newResourceAuthRadiusServers() resource.Resource {
	return &resourceAuthRadiusServers{
		resourceAuthRadiusServer: &resourceAuthRadiusServer{},
	}
}

type resourceAuthRadiusServers struct {
	*resourceAuthRadiusServer
}

func (r *resourceAuthRadiusServers) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_radius_servers"
}

func (r *resourceAuthRadiusServers) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceAuthRadiusServer.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_auth_radius_servers is deprecated. Please use fortisase_auth_radius_server instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceAuthRadiusServers) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceAuthRadiusServer.Configure(ctx, req, resp)
	r.resourceAuthRadiusServer.resourceName = "fortisase_auth_radius_servers"
}
func (r *resourceAuthRadiusServers) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
