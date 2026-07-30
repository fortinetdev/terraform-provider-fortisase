// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityFortiguardLocalCategories keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityFortiguardLocalCategories{}

func newResourceSecurityFortiguardLocalCategories() resource.Resource {
	return &resourceSecurityFortiguardLocalCategories{
		resourceSecurityFortiguardLocalCategory: &resourceSecurityFortiguardLocalCategory{},
	}
}

type resourceSecurityFortiguardLocalCategories struct {
	*resourceSecurityFortiguardLocalCategory
}

func (r *resourceSecurityFortiguardLocalCategories) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_fortiguard_local_categories"
}

func (r *resourceSecurityFortiguardLocalCategories) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityFortiguardLocalCategory.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_fortiguard_local_categories is deprecated. Please use fortisase_security_fortiguard_local_category instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityFortiguardLocalCategories) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityFortiguardLocalCategory.Configure(ctx, req, resp)
	r.resourceSecurityFortiguardLocalCategory.resourceName = "fortisase_security_fortiguard_local_categories"
}
func (r *resourceSecurityFortiguardLocalCategories) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
