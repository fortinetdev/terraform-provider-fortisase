// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityCertRemoteCerts keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityCertRemoteCerts{}

func newResourceSecurityCertRemoteCerts() resource.Resource {
	return &resourceSecurityCertRemoteCerts{
		resourceSecurityCertRemoteCert: &resourceSecurityCertRemoteCert{},
	}
}

type resourceSecurityCertRemoteCerts struct {
	*resourceSecurityCertRemoteCert
}

func (r *resourceSecurityCertRemoteCerts) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_remote_certs"
}

func (r *resourceSecurityCertRemoteCerts) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityCertRemoteCert.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_cert_remote_certs is deprecated. Please use fortisase_security_cert_remote_cert instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityCertRemoteCerts) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityCertRemoteCert.Configure(ctx, req, resp)
	r.resourceSecurityCertRemoteCert.resourceName = "fortisase_security_cert_remote_certs"
}
func (r *resourceSecurityCertRemoteCerts) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
