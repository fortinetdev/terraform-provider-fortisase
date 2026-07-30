// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityIpsCustomSignatures keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityIpsCustomSignatures{}

func newResourceSecurityIpsCustomSignatures() resource.Resource {
	return &resourceSecurityIpsCustomSignatures{
		resourceSecurityIpsCustomSignature: &resourceSecurityIpsCustomSignature{},
	}
}

type resourceSecurityIpsCustomSignatures struct {
	*resourceSecurityIpsCustomSignature
}

func (r *resourceSecurityIpsCustomSignatures) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ips_custom_signatures"
}

func (r *resourceSecurityIpsCustomSignatures) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityIpsCustomSignature.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_ips_custom_signatures is deprecated. Please use fortisase_security_ips_custom_signature instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityIpsCustomSignatures) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityIpsCustomSignature.Configure(ctx, req, resp)
	r.resourceSecurityIpsCustomSignature.resourceName = "fortisase_security_ips_custom_signatures"
}
func (r *resourceSecurityIpsCustomSignatures) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
