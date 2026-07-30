// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointZtnaProfiles keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointZtnaProfiles{}

func newResourceEndpointZtnaProfiles() resource.Resource {
	return &resourceEndpointZtnaProfiles{
		resourceEndpointZtnaProfile: &resourceEndpointZtnaProfile{},
	}
}

type resourceEndpointZtnaProfiles struct {
	*resourceEndpointZtnaProfile
}

func (r *resourceEndpointZtnaProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_profiles"
}

func (r *resourceEndpointZtnaProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointZtnaProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_ztna_profiles is deprecated. Please use fortisase_endpoint_ztna_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointZtnaProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointZtnaProfile.Configure(ctx, req, resp)
	r.resourceEndpointZtnaProfile.resourceName = "fortisase_endpoint_ztna_profiles"
}
func (r *resourceEndpointZtnaProfiles) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
