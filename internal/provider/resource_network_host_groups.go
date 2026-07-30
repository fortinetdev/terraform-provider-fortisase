// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceNetworkHostGroups keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceNetworkHostGroups{}

func newResourceNetworkHostGroups() resource.Resource {
	return &resourceNetworkHostGroups{
		resourceNetworkHostGroup: &resourceNetworkHostGroup{},
	}
}

type resourceNetworkHostGroups struct {
	*resourceNetworkHostGroup
}

func (r *resourceNetworkHostGroups) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_host_groups"
}

func (r *resourceNetworkHostGroups) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceNetworkHostGroup.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_network_host_groups is deprecated. Please use fortisase_network_host_group instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceNetworkHostGroups) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceNetworkHostGroup.Configure(ctx, req, resp)
	r.resourceNetworkHostGroup.resourceName = "fortisase_network_host_groups"
}
func (r *resourceNetworkHostGroups) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
