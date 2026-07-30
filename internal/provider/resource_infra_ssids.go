// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceInfraSsids keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceInfraSsids{}

func newResourceInfraSsids() resource.Resource {
	return &resourceInfraSsids{
		resourceInfraSsid: &resourceInfraSsid{},
	}
}

type resourceInfraSsids struct {
	*resourceInfraSsid
}

func (r *resourceInfraSsids) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_ssids"
}

func (r *resourceInfraSsids) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceInfraSsid.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_infra_ssids is deprecated. Please use fortisase_infra_ssid instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceInfraSsids) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceInfraSsid.Configure(ctx, req, resp)
	r.resourceInfraSsid.resourceName = "fortisase_infra_ssids"
}
func (r *resourceInfraSsids) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
