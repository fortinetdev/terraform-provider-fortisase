// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointZtnaTags keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointZtnaTags{}

func newDatasourceEndpointZtnaTags() datasource.DataSource {
	return &datasourceEndpointZtnaTags{
		datasourceEndpointZtnaTag: &datasourceEndpointZtnaTag{},
	}
}

type datasourceEndpointZtnaTags struct {
	*datasourceEndpointZtnaTag
}

func (r *datasourceEndpointZtnaTags) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_tags"
}

func (r *datasourceEndpointZtnaTags) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointZtnaTag.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_ztna_tags is deprecated. Please use fortisase_endpoint_ztna_tag instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointZtnaTags) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointZtnaTag.Configure(ctx, req, resp)
	r.datasourceEndpointZtnaTag.resourceName = "fortisase_endpoint_ztna_tags"
}
