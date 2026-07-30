// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityDlpSensors keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityDlpSensors{}

func newDatasourceSecurityDlpSensors() datasource.DataSource {
	return &datasourceSecurityDlpSensors{
		datasourceSecurityDlpSensor: &datasourceSecurityDlpSensor{},
	}
}

type datasourceSecurityDlpSensors struct {
	*datasourceSecurityDlpSensor
}

func (r *datasourceSecurityDlpSensors) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_sensors"
}

func (r *datasourceSecurityDlpSensors) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityDlpSensor.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_sensors is deprecated. Please use fortisase_security_dlp_sensor instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityDlpSensors) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityDlpSensor.Configure(ctx, req, resp)
	r.datasourceSecurityDlpSensor.resourceName = "fortisase_security_dlp_sensors"
}
