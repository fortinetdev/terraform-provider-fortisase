// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityCertLocalCerts keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityCertLocalCerts{}

func newResourceSecurityCertLocalCerts() resource.Resource {
	return &resourceSecurityCertLocalCerts{
		resourceSecurityCertLocalCert: &resourceSecurityCertLocalCert{},
	}
}

type resourceSecurityCertLocalCerts struct {
	*resourceSecurityCertLocalCert
}

func (r *resourceSecurityCertLocalCerts) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_local_certs"
}

func (r *resourceSecurityCertLocalCerts) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityCertLocalCert.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_cert_local_certs is deprecated. Please use fortisase_security_cert_local_cert instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityCertLocalCerts) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityCertLocalCert.Configure(ctx, req, resp)
	r.resourceSecurityCertLocalCert.resourceName = "fortisase_security_cert_local_certs"
}
func (r *resourceSecurityCertLocalCerts) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
