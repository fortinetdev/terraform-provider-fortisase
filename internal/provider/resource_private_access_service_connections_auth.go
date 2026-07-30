// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourcePrivateAccessServiceConnectionsAuth keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourcePrivateAccessServiceConnectionsAuth{}

func newResourcePrivateAccessServiceConnectionsAuth() resource.Resource {
	return &resourcePrivateAccessServiceConnectionsAuth{
		resourcePrivateAccessServiceConnectionAuth2Edl: &resourcePrivateAccessServiceConnectionAuth2Edl{},
	}
}

type resourcePrivateAccessServiceConnectionsAuth struct {
	*resourcePrivateAccessServiceConnectionAuth2Edl
}

func (r *resourcePrivateAccessServiceConnectionsAuth) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_service_connections_auth"
}

func (r *resourcePrivateAccessServiceConnectionsAuth) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourcePrivateAccessServiceConnectionAuth2Edl.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_private_access_service_connections_auth is deprecated. Please use fortisase_private_access_service_connection_auth instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourcePrivateAccessServiceConnectionsAuth) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourcePrivateAccessServiceConnectionAuth2Edl.Configure(ctx, req, resp)
	r.resourcePrivateAccessServiceConnectionAuth2Edl.resourceName = "fortisase_private_access_service_connections_auth"
}
func (r *resourcePrivateAccessServiceConnectionsAuth) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
