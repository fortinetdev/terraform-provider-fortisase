// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
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
var _ resource.Resource = &resourceDemSpaApplication{}
var _ resource.ResourceWithMoveState = &resourceDemSpaApplication{}

func newResourceDemSpaApplication() resource.Resource {
	return &resourceDemSpaApplication{}
}

type resourceDemSpaApplication struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceDemSpaApplicationModel describes the resource data model.
type resourceDemSpaApplicationModel struct {
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

func (r *resourceDemSpaApplication) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dem_spa_application"
}

func (r *resourceDemSpaApplication) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DEM SPA Application Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 35),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 79),
				},
				Computed: true,
				Optional: true,
			},
			"latency_threshold": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(10000000),
				},
				Computed: true,
				Optional: true,
			},
			"jitter_threshold": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(10000000),
				},
				Computed: true,
				Optional: true,
			},
			"packetloss_threshold": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(100),
				},
				Computed: true,
				Optional: true,
			},
			"interval": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.Between(20, 3600000),
				},
				Computed: true,
				Optional: true,
			},
			"fail_time": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.Between(1, 3600),
				},
				Computed: true,
				Optional: true,
			},
			"recovery_time": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.Between(1, 3600),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceDemSpaApplication) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *resourceDemSpaApplication) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_dem_spa_applications" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceDemSpaApplicationModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceDemSpaApplication) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("DemSpaApplications")
	lock.Lock()
	defer lock.Unlock()
	var data resourceDemSpaApplicationModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectDemSpaApplication(ctx, diags))
	input_model.URLParams = *(data.getURLObjectDemSpaApplication(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateDemSpaApplications(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	if responseMkey, ok := getCreateResponseMkey(output, "primaryKey"); ok {
		mkey = responseMkey
	}
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectDemSpaApplication(ctx, "read", diags))

	read_output, err := c.ReadDemSpaApplications(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshDemSpaApplication(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceDemSpaApplication) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("DemSpaApplications")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceDemSpaApplicationModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceDemSpaApplicationModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectDemSpaApplication(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectDemSpaApplication(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateDemSpaApplications(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectDemSpaApplication(ctx, "read", diags))

	read_output, err := c.ReadDemSpaApplications(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshDemSpaApplication(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceDemSpaApplication) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("DemSpaApplications")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceDemSpaApplicationModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectDemSpaApplication(ctx, "delete", diags))

	output, err := c.DeleteDemSpaApplications(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceDemSpaApplication) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceDemSpaApplicationModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectDemSpaApplication(ctx, "read", diags))

	read_output, err := c.ReadDemSpaApplications(&input_model)
	if err != nil {
		if isNotFoundResponse(read_output) {
			resp.State.RemoveResource(ctx)
			return
		}
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
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

func (r *resourceDemSpaApplication) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceDemSpaApplicationModel) refreshDemSpaApplication(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceDemSpaApplicationModel) getCreateObjectDemSpaApplication(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Server.IsNull() && !data.Server.IsUnknown() {
		result["server"] = data.Server.ValueString()
	}

	if !data.LatencyThreshold.IsNull() && !data.LatencyThreshold.IsUnknown() {
		result["latencyThreshold"] = data.LatencyThreshold.ValueFloat64()
	}

	if !data.JitterThreshold.IsNull() && !data.JitterThreshold.IsUnknown() {
		result["jitterThreshold"] = data.JitterThreshold.ValueFloat64()
	}

	if !data.PacketlossThreshold.IsNull() && !data.PacketlossThreshold.IsUnknown() {
		result["packetlossThreshold"] = data.PacketlossThreshold.ValueFloat64()
	}

	if !data.Interval.IsNull() && !data.Interval.IsUnknown() {
		result["interval"] = data.Interval.ValueFloat64()
	}

	if !data.FailTime.IsNull() && !data.FailTime.IsUnknown() {
		result["failTime"] = data.FailTime.ValueFloat64()
	}

	if !data.RecoveryTime.IsNull() && !data.RecoveryTime.IsUnknown() {
		result["recoveryTime"] = data.RecoveryTime.ValueFloat64()
	}

	return &result
}

func (data *resourceDemSpaApplicationModel) getUpdateObjectDemSpaApplication(ctx context.Context, state resourceDemSpaApplicationModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Server.IsNull() && !data.Server.IsUnknown() {
		result["server"] = data.Server.ValueString()
	}

	if !data.LatencyThreshold.IsNull() && !data.LatencyThreshold.IsUnknown() {
		result["latencyThreshold"] = data.LatencyThreshold.ValueFloat64()
	}

	if !data.JitterThreshold.IsNull() && !data.JitterThreshold.IsUnknown() {
		result["jitterThreshold"] = data.JitterThreshold.ValueFloat64()
	}

	if !data.PacketlossThreshold.IsNull() && !data.PacketlossThreshold.IsUnknown() {
		result["packetlossThreshold"] = data.PacketlossThreshold.ValueFloat64()
	}

	if !data.Interval.IsNull() && !data.Interval.IsUnknown() {
		result["interval"] = data.Interval.ValueFloat64()
	}

	if !data.FailTime.IsNull() && !data.FailTime.IsUnknown() {
		result["failTime"] = data.FailTime.ValueFloat64()
	}

	if !data.RecoveryTime.IsNull() && !data.RecoveryTime.IsUnknown() {
		result["recoveryTime"] = data.RecoveryTime.ValueFloat64()
	}

	return &result
}

func (data *resourceDemSpaApplicationModel) getURLObjectDemSpaApplication(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
