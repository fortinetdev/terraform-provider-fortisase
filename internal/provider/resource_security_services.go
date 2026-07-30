// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityServices keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityServices{}

func newResourceSecurityServices() resource.Resource {
	return &resourceSecurityServices{
		resourceSecurityService: &resourceSecurityService{},
	}
}

type resourceSecurityServices struct {
	*resourceSecurityService
}

func (r *resourceSecurityServices) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_services"
}

func (r *resourceSecurityServices) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityService.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_services is deprecated. Please use fortisase_security_service instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityServices) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityService.Configure(ctx, req, resp)
	r.resourceSecurityService.resourceName = "fortisase_security_services"
}
func (r *resourceSecurityServices) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
