// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityScheduleGroups keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityScheduleGroups{}

func newResourceSecurityScheduleGroups() resource.Resource {
	return &resourceSecurityScheduleGroups{
		resourceSecurityScheduleGroup: &resourceSecurityScheduleGroup{},
	}
}

type resourceSecurityScheduleGroups struct {
	*resourceSecurityScheduleGroup
}

func (r *resourceSecurityScheduleGroups) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_schedule_groups"
}

func (r *resourceSecurityScheduleGroups) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityScheduleGroup.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_schedule_groups is deprecated. Please use fortisase_security_schedule_group instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityScheduleGroups) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityScheduleGroup.Configure(ctx, req, resp)
	r.resourceSecurityScheduleGroup.resourceName = "fortisase_security_schedule_groups"
}
func (r *resourceSecurityScheduleGroups) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
