// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityScheduleGroups keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityScheduleGroups{}

func newDatasourceSecurityScheduleGroups() datasource.DataSource {
	return &datasourceSecurityScheduleGroups{
		datasourceSecurityScheduleGroup: &datasourceSecurityScheduleGroup{},
	}
}

type datasourceSecurityScheduleGroups struct {
	*datasourceSecurityScheduleGroup
}

func (r *datasourceSecurityScheduleGroups) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_schedule_groups"
}

func (r *datasourceSecurityScheduleGroups) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityScheduleGroup.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_schedule_groups is deprecated. Please use fortisase_security_schedule_group instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityScheduleGroups) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityScheduleGroup.Configure(ctx, req, resp)
	r.datasourceSecurityScheduleGroup.resourceName = "fortisase_security_schedule_groups"
}
