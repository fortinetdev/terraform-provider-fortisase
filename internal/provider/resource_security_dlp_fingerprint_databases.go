// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityDlpFingerprintDatabases keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityDlpFingerprintDatabases{}

func newResourceSecurityDlpFingerprintDatabases() resource.Resource {
	return &resourceSecurityDlpFingerprintDatabases{
		resourceSecurityDlpFingerprintDatabase: &resourceSecurityDlpFingerprintDatabase{},
	}
}

type resourceSecurityDlpFingerprintDatabases struct {
	*resourceSecurityDlpFingerprintDatabase
}

func (r *resourceSecurityDlpFingerprintDatabases) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_fingerprint_databases"
}

func (r *resourceSecurityDlpFingerprintDatabases) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityDlpFingerprintDatabase.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_fingerprint_databases is deprecated. Please use fortisase_security_dlp_fingerprint_database instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityDlpFingerprintDatabases) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityDlpFingerprintDatabase.Configure(ctx, req, resp)
	r.resourceSecurityDlpFingerprintDatabase.resourceName = "fortisase_security_dlp_fingerprint_databases"
}
func (r *resourceSecurityDlpFingerprintDatabases) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
