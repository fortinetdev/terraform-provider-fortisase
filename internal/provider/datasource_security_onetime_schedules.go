// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityOnetimeSchedules keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityOnetimeSchedules{}

func newDatasourceSecurityOnetimeSchedules() datasource.DataSource {
	return &datasourceSecurityOnetimeSchedules{
		datasourceSecurityOnetimeSchedule: &datasourceSecurityOnetimeSchedule{},
	}
}

type datasourceSecurityOnetimeSchedules struct {
	*datasourceSecurityOnetimeSchedule
}

func (r *datasourceSecurityOnetimeSchedules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_onetime_schedules"
}

func (r *datasourceSecurityOnetimeSchedules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityOnetimeSchedule.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_onetime_schedules is deprecated. Please use fortisase_security_onetime_schedule instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityOnetimeSchedules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityOnetimeSchedule.Configure(ctx, req, resp)
	r.datasourceSecurityOnetimeSchedule.resourceName = "fortisase_security_onetime_schedules"
}
