// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityRecurringSchedules keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityRecurringSchedules{}

func newResourceSecurityRecurringSchedules() resource.Resource {
	return &resourceSecurityRecurringSchedules{
		resourceSecurityRecurringSchedule: &resourceSecurityRecurringSchedule{},
	}
}

type resourceSecurityRecurringSchedules struct {
	*resourceSecurityRecurringSchedule
}

func (r *resourceSecurityRecurringSchedules) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_recurring_schedules"
}

func (r *resourceSecurityRecurringSchedules) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityRecurringSchedule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_recurring_schedules is deprecated. Please use fortisase_security_recurring_schedule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityRecurringSchedules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityRecurringSchedule.Configure(ctx, req, resp)
	r.resourceSecurityRecurringSchedule.resourceName = "fortisase_security_recurring_schedules"
}
func (r *resourceSecurityRecurringSchedules) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
