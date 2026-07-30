// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityEndpointToEndpointPolicies keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityEndpointToEndpointPolicies{}

func newDatasourceSecurityEndpointToEndpointPolicies() datasource.DataSource {
	return &datasourceSecurityEndpointToEndpointPolicies{
		datasourceSecurityEndpointToEndpointPolicy: &datasourceSecurityEndpointToEndpointPolicy{},
	}
}

type datasourceSecurityEndpointToEndpointPolicies struct {
	*datasourceSecurityEndpointToEndpointPolicy
}

func (r *datasourceSecurityEndpointToEndpointPolicies) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_endpoint_to_endpoint_policies"
}

func (r *datasourceSecurityEndpointToEndpointPolicies) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityEndpointToEndpointPolicy.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_endpoint_to_endpoint_policies is deprecated. Please use fortisase_security_endpoint_to_endpoint_policy instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityEndpointToEndpointPolicies) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityEndpointToEndpointPolicy.Configure(ctx, req, resp)
	r.datasourceSecurityEndpointToEndpointPolicy.resourceName = "fortisase_security_endpoint_to_endpoint_policies"
}
