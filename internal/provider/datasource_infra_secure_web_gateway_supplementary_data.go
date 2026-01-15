// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceInfraSecureWebGatewaySupplementaryData{}

func newDatasourceInfraSecureWebGatewaySupplementaryData() datasource.DataSource {
	return &datasourceInfraSecureWebGatewaySupplementaryData{}
}

type datasourceInfraSecureWebGatewaySupplementaryData struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceInfraSecureWebGatewaySupplementaryDataModel describes the datasource data model.
type datasourceInfraSecureWebGatewaySupplementaryDataModel struct {
	PrimaryKey           types.String  `tfsdk:"primary_key"`
	SessionDurationHours types.Float64 `tfsdk:"session_duration_hours"`
	EndSessionAfterMins  types.Float64 `tfsdk:"end_session_after_mins"`
}

func (r *datasourceInfraSecureWebGatewaySupplementaryData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_secure_web_gateway_supplementary_data"
}

func (r *datasourceInfraSecureWebGatewaySupplementaryData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("$sase-global"),
				},
				Required: true,
			},
			"session_duration_hours": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"end_session_after_mins": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceInfraSecureWebGatewaySupplementaryData) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Always perform a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*FortiClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *FortiClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.fortiClient = client
	r.resourceName = "fortisase_infra_secure_web_gateway_supplementary_data"
}

func (r *datasourceInfraSecureWebGatewaySupplementaryData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceInfraSecureWebGatewaySupplementaryDataModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadInfraSecureWebGatewaySupplementaryData(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraSecureWebGatewaySupplementaryData(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceInfraSecureWebGatewaySupplementaryDataModel) refreshInfraSecureWebGatewaySupplementaryData(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["sessionDurationHours"]; ok {
		m.SessionDurationHours = parseFloat64Value(v)
	}

	if v, ok := o["endSessionAfterMins"]; ok {
		m.EndSessionAfterMins = parseFloat64Value(v)
	}

	return diags
}
