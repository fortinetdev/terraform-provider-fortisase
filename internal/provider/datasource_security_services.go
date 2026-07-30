// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityServices keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityServices{}

func newDatasourceSecurityServices() datasource.DataSource {
	return &datasourceSecurityServices{
		datasourceSecurityService: &datasourceSecurityService{},
	}
}

type datasourceSecurityServices struct {
	*datasourceSecurityService
}

func (r *datasourceSecurityServices) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_services"
}

func (r *datasourceSecurityServices) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityService.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_services is deprecated. Please use fortisase_security_service instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityServices) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityService.Configure(ctx, req, resp)
	r.datasourceSecurityService.resourceName = "fortisase_security_services"
}
