// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityDlpFilePatterns keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityDlpFilePatterns{}

func newResourceSecurityDlpFilePatterns() resource.Resource {
	return &resourceSecurityDlpFilePatterns{
		resourceSecurityDlpFilePattern: &resourceSecurityDlpFilePattern{},
	}
}

type resourceSecurityDlpFilePatterns struct {
	*resourceSecurityDlpFilePattern
}

func (r *resourceSecurityDlpFilePatterns) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_file_patterns"
}

func (r *resourceSecurityDlpFilePatterns) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityDlpFilePattern.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_file_patterns is deprecated. Please use fortisase_security_dlp_file_pattern instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityDlpFilePatterns) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityDlpFilePattern.Configure(ctx, req, resp)
	r.resourceSecurityDlpFilePattern.resourceName = "fortisase_security_dlp_file_patterns"
}
func (r *resourceSecurityDlpFilePatterns) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
