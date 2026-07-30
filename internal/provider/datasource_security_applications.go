// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityApplications keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityApplications{}

func newDatasourceSecurityApplications() datasource.DataSource {
	return &datasourceSecurityApplications{
		datasourceSecurityApplication: &datasourceSecurityApplication{},
	}
}

type datasourceSecurityApplications struct {
	*datasourceSecurityApplication
}

func (r *datasourceSecurityApplications) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_applications"
}

func (r *datasourceSecurityApplications) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityApplication.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_applications is deprecated. Please use fortisase_security_application instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityApplications) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityApplication.Configure(ctx, req, resp)
	r.datasourceSecurityApplication.resourceName = "fortisase_security_applications"
}
