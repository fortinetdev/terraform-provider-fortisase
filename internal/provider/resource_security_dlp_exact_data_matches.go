// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityDlpExactDataMatches keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityDlpExactDataMatches{}

func newResourceSecurityDlpExactDataMatches() resource.Resource {
	return &resourceSecurityDlpExactDataMatches{
		resourceSecurityDlpExactDataMatch: &resourceSecurityDlpExactDataMatch{},
	}
}

type resourceSecurityDlpExactDataMatches struct {
	*resourceSecurityDlpExactDataMatch
}

func (r *resourceSecurityDlpExactDataMatches) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_exact_data_matches"
}

func (r *resourceSecurityDlpExactDataMatches) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityDlpExactDataMatch.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_exact_data_matches is deprecated. Please use fortisase_security_dlp_exact_data_match instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityDlpExactDataMatches) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityDlpExactDataMatch.Configure(ctx, req, resp)
	r.resourceSecurityDlpExactDataMatch.resourceName = "fortisase_security_dlp_exact_data_matches"
}
func (r *resourceSecurityDlpExactDataMatches) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
