// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourcePrivateAccessServiceConnections keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourcePrivateAccessServiceConnections{}

func newResourcePrivateAccessServiceConnections() resource.Resource {
	return &resourcePrivateAccessServiceConnections{
		resourcePrivateAccessServiceConnection: &resourcePrivateAccessServiceConnection{},
	}
}

type resourcePrivateAccessServiceConnections struct {
	*resourcePrivateAccessServiceConnection
}

func (r *resourcePrivateAccessServiceConnections) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_service_connections"
}

func (r *resourcePrivateAccessServiceConnections) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourcePrivateAccessServiceConnection.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_private_access_service_connections is deprecated. Please use fortisase_private_access_service_connection instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourcePrivateAccessServiceConnections) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourcePrivateAccessServiceConnection.Configure(ctx, req, resp)
	r.resourcePrivateAccessServiceConnection.resourceName = "fortisase_private_access_service_connections"
}
func (r *resourcePrivateAccessServiceConnections) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
