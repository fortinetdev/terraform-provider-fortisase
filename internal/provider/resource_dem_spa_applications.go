// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceDemSpaApplications keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceDemSpaApplications{}

func newResourceDemSpaApplications() resource.Resource {
	return &resourceDemSpaApplications{
		resourceDemSpaApplication: &resourceDemSpaApplication{},
	}
}

type resourceDemSpaApplications struct {
	*resourceDemSpaApplication
}

func (r *resourceDemSpaApplications) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dem_spa_applications"
}

func (r *resourceDemSpaApplications) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceDemSpaApplication.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_dem_spa_applications is deprecated. Please use fortisase_dem_spa_application instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceDemSpaApplications) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceDemSpaApplication.Configure(ctx, req, resp)
	r.resourceDemSpaApplication.resourceName = "fortisase_dem_spa_applications"
}
func (r *resourceDemSpaApplications) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
