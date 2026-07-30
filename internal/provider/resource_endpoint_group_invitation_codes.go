// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceEndpointGroupInvitationCodes keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceEndpointGroupInvitationCodes{}

func newResourceEndpointGroupInvitationCodes() resource.Resource {
	return &resourceEndpointGroupInvitationCodes{
		resourceEndpointGroupInvitationCode: &resourceEndpointGroupInvitationCode{},
	}
}

type resourceEndpointGroupInvitationCodes struct {
	*resourceEndpointGroupInvitationCode
}

func (r *resourceEndpointGroupInvitationCodes) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_group_invitation_codes"
}

func (r *resourceEndpointGroupInvitationCodes) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceEndpointGroupInvitationCode.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_group_invitation_codes is deprecated. Please use fortisase_endpoint_group_invitation_code instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceEndpointGroupInvitationCodes) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceEndpointGroupInvitationCode.Configure(ctx, req, resp)
	r.resourceEndpointGroupInvitationCode.resourceName = "fortisase_endpoint_group_invitation_codes"
}
func (r *resourceEndpointGroupInvitationCodes) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
