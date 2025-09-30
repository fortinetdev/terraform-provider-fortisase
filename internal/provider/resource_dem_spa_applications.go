// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceDemSpaApplications{}

func newResourceDemSpaApplications() resource.Resource {
	return &resourceDemSpaApplications{}
}

type resourceDemSpaApplications struct {
	fortiClient *FortiClient
}

// resourceDemSpaApplicationsModel describes the resource data model.
type resourceDemSpaApplicationsModel struct {
	ID                  types.String  `tfsdk:"id"`
	PrimaryKey          types.String  `tfsdk:"primary_key"`
	Server              types.String  `tfsdk:"server"`
	LatencyThreshold    types.Float64 `tfsdk:"latency_threshold"`
	JitterThreshold     types.Float64 `tfsdk:"jitter_threshold"`
	PacketlossThreshold types.Float64 `tfsdk:"packetloss_threshold"`
	Interval            types.Float64 `tfsdk:"interval"`
	FailTime            types.Float64 `tfsdk:"fail_time"`
	RecoveryTime        types.Float64 `tfsdk:"recovery_time"`
}

func (r *resourceDemSpaApplications) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dem_spa_applications"
}

func (r *resourceDemSpaApplications) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
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

func (r *resourceDemSpaApplications) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceDemSpaApplications) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceDemSpaApplicationsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectDemSpaApplications(ctx, diags))
	input_model.URLParams = *(data.getURLObjectDemSpaApplications(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateDemSpaApplications(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectDemSpaApplications(ctx, "read", diags))

	read_output, err := c.ReadDemSpaApplications(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
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

func (r *resourceDemSpaApplications) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceDemSpaApplicationsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceDemSpaApplicationsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectDemSpaApplications(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectDemSpaApplications(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateDemSpaApplications(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectDemSpaApplications(ctx, "read", diags))

	read_output, err := c.ReadDemSpaApplications(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
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

func (r *resourceDemSpaApplications) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceDemSpaApplicationsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectDemSpaApplications(ctx, "delete", diags))

	err := c.DeleteDemSpaApplications(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceDemSpaApplications) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceDemSpaApplicationsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectDemSpaApplications(ctx, "read", diags))

	read_output, err := c.ReadDemSpaApplications(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
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

func (r *resourceDemSpaApplications) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceDemSpaApplicationsModel) refreshDemSpaApplications(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceDemSpaApplicationsModel) getCreateObjectDemSpaApplications(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Server.IsNull() {
		result["server"] = data.Server.ValueString()
	}

	if !data.LatencyThreshold.IsNull() {
		result["latencyThreshold"] = data.LatencyThreshold.ValueFloat64()
	}

	if !data.JitterThreshold.IsNull() {
		result["jitterThreshold"] = data.JitterThreshold.ValueFloat64()
	}

	if !data.PacketlossThreshold.IsNull() {
		result["packetlossThreshold"] = data.PacketlossThreshold.ValueFloat64()
	}

	if !data.Interval.IsNull() {
		result["interval"] = data.Interval.ValueFloat64()
	}

	if !data.FailTime.IsNull() {
		result["failTime"] = data.FailTime.ValueFloat64()
	}

	if !data.RecoveryTime.IsNull() {
		result["recoveryTime"] = data.RecoveryTime.ValueFloat64()
	}

	return &result
}

func (data *resourceDemSpaApplicationsModel) getUpdateObjectDemSpaApplications(ctx context.Context, state resourceDemSpaApplicationsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Server.IsNull() {
		result["server"] = data.Server.ValueString()
	}

	if !data.LatencyThreshold.IsNull() && !data.LatencyThreshold.Equal(state.LatencyThreshold) {
		result["latencyThreshold"] = data.LatencyThreshold.ValueFloat64()
	}

	if !data.JitterThreshold.IsNull() && !data.JitterThreshold.Equal(state.JitterThreshold) {
		result["jitterThreshold"] = data.JitterThreshold.ValueFloat64()
	}

	if !data.PacketlossThreshold.IsNull() && !data.PacketlossThreshold.Equal(state.PacketlossThreshold) {
		result["packetlossThreshold"] = data.PacketlossThreshold.ValueFloat64()
	}

	if !data.Interval.IsNull() && !data.Interval.Equal(state.Interval) {
		result["interval"] = data.Interval.ValueFloat64()
	}

	if !data.FailTime.IsNull() && !data.FailTime.Equal(state.FailTime) {
		result["failTime"] = data.FailTime.ValueFloat64()
	}

	if !data.RecoveryTime.IsNull() && !data.RecoveryTime.Equal(state.RecoveryTime) {
		result["recoveryTime"] = data.RecoveryTime.ValueFloat64()
	}

	return &result
}

func (data *resourceDemSpaApplicationsModel) getURLObjectDemSpaApplications(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
