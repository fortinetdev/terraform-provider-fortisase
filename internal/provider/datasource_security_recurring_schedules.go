// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityRecurringSchedules keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityRecurringSchedules{}

func newDatasourceSecurityRecurringSchedules() datasource.DataSource {
	return &datasourceSecurityRecurringSchedules{
		datasourceSecurityRecurringSchedule: &datasourceSecurityRecurringSchedule{},
	}
}

type datasourceSecurityRecurringSchedules struct {
	*datasourceSecurityRecurringSchedule
}

func (r *datasourceSecurityRecurringSchedules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_recurring_schedules"
}

func (r *datasourceSecurityRecurringSchedules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityRecurringSchedule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_recurring_schedules is deprecated. Please use fortisase_security_recurring_schedule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityRecurringSchedules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityRecurringSchedule.Configure(ctx, req, resp)
	r.datasourceSecurityRecurringSchedule.resourceName = "fortisase_security_recurring_schedules"
}
