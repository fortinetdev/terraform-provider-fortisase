// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityServiceGroups keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityServiceGroups{}

func newResourceSecurityServiceGroups() resource.Resource {
	return &resourceSecurityServiceGroups{
		resourceSecurityServiceGroup: &resourceSecurityServiceGroup{},
	}
}

type resourceSecurityServiceGroups struct {
	*resourceSecurityServiceGroup
}

func (r *resourceSecurityServiceGroups) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_service_groups"
}

func (r *resourceSecurityServiceGroups) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityServiceGroup.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_service_groups is deprecated. Please use fortisase_security_service_group instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityServiceGroups) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityServiceGroup.Configure(ctx, req, resp)
	r.resourceSecurityServiceGroup.resourceName = "fortisase_security_service_groups"
}
func (r *resourceSecurityServiceGroups) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
