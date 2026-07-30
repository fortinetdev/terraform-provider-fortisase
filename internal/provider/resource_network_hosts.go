// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceNetworkHosts keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceNetworkHosts{}

func newResourceNetworkHosts() resource.Resource {
	return &resourceNetworkHosts{
		resourceNetworkHost: &resourceNetworkHost{},
	}
}

type resourceNetworkHosts struct {
	*resourceNetworkHost
}

func (r *resourceNetworkHosts) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hosts"
}

func (r *resourceNetworkHosts) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceNetworkHost.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_network_hosts is deprecated. Please use fortisase_network_host instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceNetworkHosts) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceNetworkHost.Configure(ctx, req, resp)
	r.resourceNetworkHost.resourceName = "fortisase_network_hosts"
}
func (r *resourceNetworkHosts) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
