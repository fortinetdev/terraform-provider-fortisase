// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointPolicies keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointPolicies{}

func newDatasourceEndpointPolicies() datasource.DataSource {
	return &datasourceEndpointPolicies{
		datasourceEndpointPolicy: &datasourceEndpointPolicy{},
	}
}

type datasourceEndpointPolicies struct {
	*datasourceEndpointPolicy
}

func (r *datasourceEndpointPolicies) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_policies"
}

func (r *datasourceEndpointPolicies) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointPolicy.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_policies is deprecated. Please use fortisase_endpoint_policy instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointPolicies) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointPolicy.Configure(ctx, req, resp)
	r.datasourceEndpointPolicy.resourceName = "fortisase_endpoint_policies"
}
