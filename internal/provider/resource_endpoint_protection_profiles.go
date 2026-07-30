// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointProtectionProfiles keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointProtectionProfiles{}

func newResourceEndpointProtectionProfiles() resource.Resource {
	return &resourceEndpointProtectionProfiles{
		resourceEndpointProtectionProfile: &resourceEndpointProtectionProfile{},
	}
}

type resourceEndpointProtectionProfiles struct {
	*resourceEndpointProtectionProfile
}

func (r *resourceEndpointProtectionProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_protection_profiles"
}

func (r *resourceEndpointProtectionProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointProtectionProfile.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_protection_profiles is deprecated. Please use fortisase_endpoint_protection_profile instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointProtectionProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointProtectionProfile.Configure(ctx, req, resp)
	r.resourceEndpointProtectionProfile.resourceName = "fortisase_endpoint_protection_profiles"
}
func (r *resourceEndpointProtectionProfiles) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
