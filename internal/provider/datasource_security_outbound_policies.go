// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityOutboundPolicies keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityOutboundPolicies{}

func newDatasourceSecurityOutboundPolicies() datasource.DataSource {
	return &datasourceSecurityOutboundPolicies{
		datasourceSecurityOutboundPolicy: &datasourceSecurityOutboundPolicy{},
	}
}

type datasourceSecurityOutboundPolicies struct {
	*datasourceSecurityOutboundPolicy
}

func (r *datasourceSecurityOutboundPolicies) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_outbound_policies"
}

func (r *datasourceSecurityOutboundPolicies) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityOutboundPolicy.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_outbound_policies is deprecated. Please use fortisase_security_outbound_policy instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityOutboundPolicies) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityOutboundPolicy.Configure(ctx, req, resp)
	r.datasourceSecurityOutboundPolicy.resourceName = "fortisase_security_outbound_policies"
}
