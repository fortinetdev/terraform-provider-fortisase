// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityOnetimeSchedules keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityOnetimeSchedules{}

func newResourceSecurityOnetimeSchedules() resource.Resource {
	return &resourceSecurityOnetimeSchedules{
		resourceSecurityOnetimeSchedule: &resourceSecurityOnetimeSchedule{},
	}
}

type resourceSecurityOnetimeSchedules struct {
	*resourceSecurityOnetimeSchedule
}

func (r *resourceSecurityOnetimeSchedules) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_onetime_schedules"
}

func (r *resourceSecurityOnetimeSchedules) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityOnetimeSchedule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_onetime_schedules is deprecated. Please use fortisase_security_onetime_schedule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityOnetimeSchedules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityOnetimeSchedule.Configure(ctx, req, resp)
	r.resourceSecurityOnetimeSchedule.resourceName = "fortisase_security_onetime_schedules"
}
func (r *resourceSecurityOnetimeSchedules) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
