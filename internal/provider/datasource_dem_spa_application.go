// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceDemSpaApplication{}

func newDatasourceDemSpaApplication() datasource.DataSource {
	return &datasourceDemSpaApplication{}
}

type datasourceDemSpaApplication struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceDemSpaApplicationModel describes the datasource data model.
type datasourceDemSpaApplicationModel struct {
	PrimaryKey          types.String  `tfsdk:"primary_key"`
	Server              types.String  `tfsdk:"server"`
	LatencyThreshold    types.Float64 `tfsdk:"latency_threshold"`
	JitterThreshold     types.Float64 `tfsdk:"jitter_threshold"`
	PacketlossThreshold types.Float64 `tfsdk:"packetloss_threshold"`
	Interval            types.Float64 `tfsdk:"interval"`
	FailTime            types.Float64 `tfsdk:"fail_time"`
	RecoveryTime        types.Float64 `tfsdk:"recovery_time"`
}

func (r *datasourceDemSpaApplication) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dem_spa_application"
}

func (r *datasourceDemSpaApplication) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DEM SPA Application Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 35),
				},
				Required: true,
			},
			"server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 79),
				},
				Computed: true,
			},
			"latency_threshold": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(10000000),
				},
				Computed: true,
			},
			"jitter_threshold": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(10000000),
				},
				Computed: true,
			},
			"packetloss_threshold": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(100),
				},
				Computed: true,
			},
			"interval": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.Between(20, 3600000),
				},
				Computed: true,
			},
			"fail_time": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.Between(1, 3600),
				},
				Computed: true,
			},
			"recovery_time": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.Between(1, 3600),
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceDemSpaApplication) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_dem_spa_application"
}

func (r *datasourceDemSpaApplication) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceDemSpaApplicationModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectDemSpaApplication(ctx, "read", diags))

	read_output, err := c.ReadDemSpaApplications(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshDemSpaApplication(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceDemSpaApplicationModel) refreshDemSpaApplication(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["server"]; ok {
		m.Server = parseStringValue(v)
	}

	if v, ok := o["latencyThreshold"]; ok {
		m.LatencyThreshold = parseFloat64Value(v)
	}

	if v, ok := o["jitterThreshold"]; ok {
		m.JitterThreshold = parseFloat64Value(v)
	}

	if v, ok := o["packetlossThreshold"]; ok {
		m.PacketlossThreshold = parseFloat64Value(v)
	}

	if v, ok := o["interval"]; ok {
		m.Interval = parseFloat64Value(v)
	}

	if v, ok := o["failTime"]; ok {
		m.FailTime = parseFloat64Value(v)
	}

	if v, ok := o["recoveryTime"]; ok {
		m.RecoveryTime = parseFloat64Value(v)
	}

	return diags
}

func (data *datasourceDemSpaApplicationModel) getURLObjectDemSpaApplication(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
