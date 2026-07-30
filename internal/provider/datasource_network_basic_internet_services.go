// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceNetworkBasicInternetServices keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceNetworkBasicInternetServices{}

func newDatasourceNetworkBasicInternetServices() datasource.DataSource {
	return &datasourceNetworkBasicInternetServices{
		datasourceNetworkBasicInternetService: &datasourceNetworkBasicInternetService{},
	}
}

type datasourceNetworkBasicInternetServices struct {
	*datasourceNetworkBasicInternetService
}

func (r *datasourceNetworkBasicInternetServices) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_basic_internet_services"
}

func (r *datasourceNetworkBasicInternetServices) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceNetworkBasicInternetService.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_network_basic_internet_services is deprecated. Please use fortisase_network_basic_internet_service instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceNetworkBasicInternetServices) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceNetworkBasicInternetService.Configure(ctx, req, resp)
	r.datasourceNetworkBasicInternetService.resourceName = "fortisase_network_basic_internet_services"
}
