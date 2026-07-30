// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityAppCustomSignatures keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityAppCustomSignatures{}

func newResourceSecurityAppCustomSignatures() resource.Resource {
	return &resourceSecurityAppCustomSignatures{
		resourceSecurityAppCustomSignature: &resourceSecurityAppCustomSignature{},
	}
}

type resourceSecurityAppCustomSignatures struct {
	*resourceSecurityAppCustomSignature
}

func (r *resourceSecurityAppCustomSignatures) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_app_custom_signatures"
}

func (r *resourceSecurityAppCustomSignatures) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityAppCustomSignature.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_app_custom_signatures is deprecated. Please use fortisase_security_app_custom_signature instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityAppCustomSignatures) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityAppCustomSignature.Configure(ctx, req, resp)
	r.resourceSecurityAppCustomSignature.resourceName = "fortisase_security_app_custom_signatures"
}
func (r *resourceSecurityAppCustomSignatures) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
