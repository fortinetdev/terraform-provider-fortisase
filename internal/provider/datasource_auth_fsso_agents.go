// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceAuthFssoAgents keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceAuthFssoAgents{}

func newDatasourceAuthFssoAgents() datasource.DataSource {
	return &datasourceAuthFssoAgents{
		datasourceAuthFssoAgent: &datasourceAuthFssoAgent{},
	}
}

type datasourceAuthFssoAgents struct {
	*datasourceAuthFssoAgent
}

func (r *datasourceAuthFssoAgents) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_fsso_agents"
}

func (r *datasourceAuthFssoAgents) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceAuthFssoAgent.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_auth_fsso_agents is deprecated. Please use fortisase_auth_fsso_agent instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceAuthFssoAgents) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceAuthFssoAgent.Configure(ctx, req, resp)
	r.datasourceAuthFssoAgent.resourceName = "fortisase_auth_fsso_agents"
}
