// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityCertLocalCaCerts keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityCertLocalCaCerts{}

func newResourceSecurityCertLocalCaCerts() resource.Resource {
	return &resourceSecurityCertLocalCaCerts{
		resourceSecurityCertLocalCaCert: &resourceSecurityCertLocalCaCert{},
	}
}

type resourceSecurityCertLocalCaCerts struct {
	*resourceSecurityCertLocalCaCert
}

func (r *resourceSecurityCertLocalCaCerts) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_local_ca_certs"
}

func (r *resourceSecurityCertLocalCaCerts) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityCertLocalCaCert.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_cert_local_ca_certs is deprecated. Please use fortisase_security_cert_local_ca_cert instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityCertLocalCaCerts) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityCertLocalCaCert.Configure(ctx, req, resp)
	r.resourceSecurityCertLocalCaCert.resourceName = "fortisase_security_cert_local_ca_certs"
}
func (r *resourceSecurityCertLocalCaCerts) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
