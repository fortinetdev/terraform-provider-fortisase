// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// resourceSecurityDlpSensors keeps the deprecated Terraform type available while
// reusing the canonical resource implementation.
var _ resource.Resource = &resourceSecurityDlpSensors{}

func newResourceSecurityDlpSensors() resource.Resource {
	return &resourceSecurityDlpSensors{
		resourceSecurityDlpSensor: &resourceSecurityDlpSensor{},
	}
}

type resourceSecurityDlpSensors struct {
	*resourceSecurityDlpSensor
}

func (r *resourceSecurityDlpSensors) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_sensors"
}

func (r *resourceSecurityDlpSensors) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.resourceSecurityDlpSensor.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_dlp_sensors is deprecated. Please use fortisase_security_dlp_sensor instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *resourceSecurityDlpSensors) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.resourceSecurityDlpSensor.Configure(ctx, req, resp)
	r.resourceSecurityDlpSensor.resourceName = "fortisase_security_dlp_sensors"
}
func (r *resourceSecurityDlpSensors) MoveState(ctx context.Context) []resource.StateMover {
	return nil
}
