// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityInternalReversePolicies keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityInternalReversePolicies{}

func newDatasourceSecurityInternalReversePolicies() datasource.DataSource {
	return &datasourceSecurityInternalReversePolicies{
		datasourceSecurityInternalReversePolicy: &datasourceSecurityInternalReversePolicy{},
	}
}

type datasourceSecurityInternalReversePolicies struct {
	*datasourceSecurityInternalReversePolicy
}

func (r *datasourceSecurityInternalReversePolicies) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_reverse_policies"
}

func (r *datasourceSecurityInternalReversePolicies) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityInternalReversePolicy.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_internal_reverse_policies is deprecated. Please use fortisase_security_internal_reverse_policy instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityInternalReversePolicies) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityInternalReversePolicy.Configure(ctx, req, resp)
	r.datasourceSecurityInternalReversePolicy.resourceName = "fortisase_security_internal_reverse_policies"
}
