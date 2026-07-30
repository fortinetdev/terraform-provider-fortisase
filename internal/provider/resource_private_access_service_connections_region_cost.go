// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourcePrivateAccessServiceConnectionsRegionCost keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourcePrivateAccessServiceConnectionsRegionCost{}

func newResourcePrivateAccessServiceConnectionsRegionCost() resource.Resource {
	return &resourcePrivateAccessServiceConnectionsRegionCost{
		resourcePrivateAccessServiceConnectionRegionCost2Edl: &resourcePrivateAccessServiceConnectionRegionCost2Edl{},
	}
}

type resourcePrivateAccessServiceConnectionsRegionCost struct {
	*resourcePrivateAccessServiceConnectionRegionCost2Edl
}

func (r *resourcePrivateAccessServiceConnectionsRegionCost) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_service_connections_region_cost"
}

func (r *resourcePrivateAccessServiceConnectionsRegionCost) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourcePrivateAccessServiceConnectionRegionCost2Edl.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_private_access_service_connections_region_cost is deprecated. Please use fortisase_private_access_service_connection_region_cost instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourcePrivateAccessServiceConnectionsRegionCost) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourcePrivateAccessServiceConnectionRegionCost2Edl.Configure(ctx, req, resp)
	r.resourcePrivateAccessServiceConnectionRegionCost2Edl.resourceName = "fortisase_private_access_service_connections_region_cost"
}
func (r *resourcePrivateAccessServiceConnectionsRegionCost) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
