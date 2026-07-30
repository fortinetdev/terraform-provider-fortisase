// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceAuthUserGroups keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceAuthUserGroups{}

func newDatasourceAuthUserGroups() datasource.DataSource {
	return &datasourceAuthUserGroups{
		datasourceAuthUserGroup: &datasourceAuthUserGroup{},
	}
}

type datasourceAuthUserGroups struct {
	*datasourceAuthUserGroup
}

func (r *datasourceAuthUserGroups) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_user_groups"
}

func (r *datasourceAuthUserGroups) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceAuthUserGroup.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_auth_user_groups is deprecated. Please use fortisase_auth_user_group instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceAuthUserGroups) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceAuthUserGroup.Configure(ctx, req, resp)
	r.datasourceAuthUserGroup.resourceName = "fortisase_auth_user_groups"
}
