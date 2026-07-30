// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityServiceGroups keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityServiceGroups{}

func newDatasourceSecurityServiceGroups() datasource.DataSource {
	return &datasourceSecurityServiceGroups{
		datasourceSecurityServiceGroup: &datasourceSecurityServiceGroup{},
	}
}

type datasourceSecurityServiceGroups struct {
	*datasourceSecurityServiceGroup
}

func (r *datasourceSecurityServiceGroups) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_service_groups"
}

func (r *datasourceSecurityServiceGroups) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityServiceGroup.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_service_groups is deprecated. Please use fortisase_security_service_group instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityServiceGroups) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityServiceGroup.Configure(ctx, req, resp)
	r.datasourceSecurityServiceGroup.resourceName = "fortisase_security_service_groups"
}
