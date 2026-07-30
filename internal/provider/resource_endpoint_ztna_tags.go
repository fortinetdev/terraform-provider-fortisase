// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointZtnaTags keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointZtnaTags{}

func newResourceEndpointZtnaTags() resource.Resource {
	return &resourceEndpointZtnaTags{
		resourceEndpointZtnaTag: &resourceEndpointZtnaTag{},
	}
}

type resourceEndpointZtnaTags struct {
	*resourceEndpointZtnaTag
}

func (r *resourceEndpointZtnaTags) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_tags"
}

func (r *resourceEndpointZtnaTags) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointZtnaTag.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_ztna_tags is deprecated. Please use fortisase_endpoint_ztna_tag instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointZtnaTags) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointZtnaTag.Configure(ctx, req, resp)
	r.resourceEndpointZtnaTag.resourceName = "fortisase_endpoint_ztna_tags"
}
func (r *resourceEndpointZtnaTags) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
