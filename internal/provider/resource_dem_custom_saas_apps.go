// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceDemCustomSaasApps keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceDemCustomSaasApps{}

func newResourceDemCustomSaasApps() resource.Resource {
	return &resourceDemCustomSaasApps{
		resourceDemCustomSaasApp: &resourceDemCustomSaasApp{},
	}
}

type resourceDemCustomSaasApps struct {
	*resourceDemCustomSaasApp
}

func (r *resourceDemCustomSaasApps) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dem_custom_saas_apps"
}

func (r *resourceDemCustomSaasApps) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceDemCustomSaasApp.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_dem_custom_saas_apps is deprecated. Please use fortisase_dem_custom_saas_app instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceDemCustomSaasApps) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceDemCustomSaasApp.Configure(ctx, req, resp)
	r.resourceDemCustomSaasApp.resourceName = "fortisase_dem_custom_saas_apps"
}
func (r *resourceDemCustomSaasApps) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
