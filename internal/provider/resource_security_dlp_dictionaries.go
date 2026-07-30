// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityDlpDictionaries keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityDlpDictionaries{}

func newResourceSecurityDlpDictionaries() resource.Resource {
	return &resourceSecurityDlpDictionaries{
		resourceSecurityDlpDictionary: &resourceSecurityDlpDictionary{},
	}
}

type resourceSecurityDlpDictionaries struct {
	*resourceSecurityDlpDictionary
}

func (r *resourceSecurityDlpDictionaries) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_dictionaries"
}

func (r *resourceSecurityDlpDictionaries) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityDlpDictionary.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_dictionaries is deprecated. Please use fortisase_security_dlp_dictionary instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityDlpDictionaries) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityDlpDictionary.Configure(ctx, req, resp)
	r.resourceSecurityDlpDictionary.resourceName = "fortisase_security_dlp_dictionaries"
}
func (r *resourceSecurityDlpDictionaries) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
