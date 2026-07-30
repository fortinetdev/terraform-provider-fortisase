// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceAuthUsers keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceAuthUsers{}

func newDatasourceAuthUsers() datasource.DataSource {
	return &datasourceAuthUsers{
		datasourceAuthUser: &datasourceAuthUser{},
	}
}

type datasourceAuthUsers struct {
	*datasourceAuthUser
}

func (r *datasourceAuthUsers) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_users"
}

func (r *datasourceAuthUsers) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceAuthUser.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_auth_users is deprecated. Please use fortisase_auth_user instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceAuthUsers) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceAuthUser.Configure(ctx, req, resp)
	r.datasourceAuthUser.resourceName = "fortisase_auth_users"
}
