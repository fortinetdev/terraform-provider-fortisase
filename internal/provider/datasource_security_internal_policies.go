// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityInternalPolicies keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityInternalPolicies{}

func newDatasourceSecurityInternalPolicies() datasource.DataSource {
	return &datasourceSecurityInternalPolicies{
		datasourceSecurityInternalPolicy: &datasourceSecurityInternalPolicy{},
	}
}

type datasourceSecurityInternalPolicies struct {
	*datasourceSecurityInternalPolicy
}

func (r *datasourceSecurityInternalPolicies) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_policies"
}

func (r *datasourceSecurityInternalPolicies) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityInternalPolicy.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_internal_policies is deprecated. Please use fortisase_security_internal_policy instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityInternalPolicies) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityInternalPolicy.Configure(ctx, req, resp)
	r.datasourceSecurityInternalPolicy.resourceName = "fortisase_security_internal_policies"
}
