// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityCertRemoteCaCerts keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityCertRemoteCaCerts{}

func newResourceSecurityCertRemoteCaCerts() resource.Resource {
	return &resourceSecurityCertRemoteCaCerts{
		resourceSecurityCertRemoteCaCert: &resourceSecurityCertRemoteCaCert{},
	}
}

type resourceSecurityCertRemoteCaCerts struct {
	*resourceSecurityCertRemoteCaCert
}

func (r *resourceSecurityCertRemoteCaCerts) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_remote_ca_certs"
}

func (r *resourceSecurityCertRemoteCaCerts) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityCertRemoteCaCert.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_cert_remote_ca_certs is deprecated. Please use fortisase_security_cert_remote_ca_cert instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityCertRemoteCaCerts) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityCertRemoteCaCert.Configure(ctx, req, resp)
	r.resourceSecurityCertRemoteCaCert.resourceName = "fortisase_security_cert_remote_ca_certs"
}
func (r *resourceSecurityCertRemoteCaCerts) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
