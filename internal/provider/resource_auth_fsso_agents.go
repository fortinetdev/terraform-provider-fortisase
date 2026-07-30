// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceAuthFssoAgents keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceAuthFssoAgents{}

func newResourceAuthFssoAgents() resource.Resource {
	return &resourceAuthFssoAgents{
		resourceAuthFssoAgent: &resourceAuthFssoAgent{},
	}
}

type resourceAuthFssoAgents struct {
	*resourceAuthFssoAgent
}

func (r *resourceAuthFssoAgents) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_fsso_agents"
}

func (r *resourceAuthFssoAgents) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceAuthFssoAgent.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_auth_fsso_agents is deprecated. Please use fortisase_auth_fsso_agent instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceAuthFssoAgents) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceAuthFssoAgent.Configure(ctx, req, resp)
	r.resourceAuthFssoAgent.resourceName = "fortisase_auth_fsso_agents"
}
func (r *resourceAuthFssoAgents) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
