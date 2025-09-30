// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceDemSpaApplications{}

func newDatasourceDemSpaApplications() datasource.DataSource {
	return &datasourceDemSpaApplications{}
}

type datasourceDemSpaApplications struct {
	fortiClient *FortiClient
}

// datasourceDemSpaApplicationsModel describes the datasource data model.
type datasourceDemSpaApplicationsModel struct {
	PrimaryKey          types.String  `tfsdk:"primary_key"`
	Server              types.String  `tfsdk:"server"`
	LatencyThreshold    types.Float64 `tfsdk:"latency_threshold"`
	JitterThreshold     types.Float64 `tfsdk:"jitter_threshold"`
	PacketlossThreshold types.Float64 `tfsdk:"packetloss_threshold"`
	Interval            types.Float64 `tfsdk:"interval"`
	FailTime            types.Float64 `tfsdk:"fail_time"`
	RecoveryTime        types.Float64 `tfsdk:"recovery_time"`
}

func (r *datasourceDemSpaApplications) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dem_spa_applications"
}

func (r *datasourceDemSpaApplications) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 35),
				},
				Required: true,
			},
			"server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 79),
				},
				Computed: true,
				Optional: true,
			},
			"latency_threshold": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(10000000),
				},
				Computed: true,
				Optional: true,
			},
			"jitter_threshold": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(10000000),
				},
				Computed: true,
				Optional: true,
			},
			"packetloss_threshold": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(100),
				},
				Computed: true,
				Optional: true,
			},
			"interval": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(20, 3600000),
				},
				Computed: true,
				Optional: true,
			},
			"fail_time": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(1, 3600),
				},
				Computed: true,
				Optional: true,
			},
			"recovery_time": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(1, 3600),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceDemSpaApplications) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
}

func (r *datasourceDemSpaApplications) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceDemSpaApplicationsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectDemSpaApplications(ctx, "read", diags))

	read_output, err := c.ReadDemSpaApplications(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshDemSpaApplications(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceDemSpaApplicationsModel) refreshDemSpaApplications(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
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

func (data *datasourceDemSpaApplicationsModel) getURLObjectDemSpaApplications(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
