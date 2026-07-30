// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/setvalidatorwarning"
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
var _ resource.Resource = &resourceSecurityRecurringSchedule{}
var _ resource.ResourceWithMoveState = &resourceSecurityRecurringSchedule{}

func newResourceSecurityRecurringSchedule() resource.Resource {
	return &resourceSecurityRecurringSchedule{}
}

type resourceSecurityRecurringSchedule struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityRecurringScheduleModel describes the resource data model.
type resourceSecurityRecurringScheduleModel struct {
	ID         types.String `tfsdk:"id"`
	PrimaryKey types.String `tfsdk:"primary_key"`
	Days       types.Set    `tfsdk:"days"`
	StartTime  types.String `tfsdk:"start_time"`
	EndTime    types.String `tfsdk:"end_time"`
}

func (r *resourceSecurityRecurringSchedule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_recurring_schedule"
}

func (r *resourceSecurityRecurringSchedule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Recurring Schedule Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 31),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"days": schema.SetAttribute{
				Validators: []validator.Set{
					setvalidatorwarning.SizeBetween(1, 7),
				},
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"start_time": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"end_time": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityRecurringSchedule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_recurring_schedule"
}
func (r *resourceSecurityRecurringSchedule) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_recurring_schedules" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityRecurringScheduleModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityRecurringSchedule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityRecurringSchedules")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityRecurringScheduleModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityRecurringSchedule(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityRecurringSchedule(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityRecurringSchedules(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityRecurringSchedule(ctx, "read", diags))

	read_output, err := c.ReadSecurityRecurringSchedules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityRecurringSchedule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityRecurringSchedule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityRecurringSchedules")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityRecurringScheduleModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityRecurringScheduleModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityRecurringSchedule(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityRecurringSchedule(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityRecurringSchedules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityRecurringSchedule(ctx, "read", diags))

	read_output, err := c.ReadSecurityRecurringSchedules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityRecurringSchedule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityRecurringSchedule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityRecurringSchedules")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityRecurringScheduleModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityRecurringSchedule(ctx, "delete", diags))

	output, err := c.DeleteSecurityRecurringSchedules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityRecurringSchedule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityRecurringScheduleModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityRecurringSchedule(ctx, "read", diags))

	read_output, err := c.ReadSecurityRecurringSchedules(&input_model)
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

	diags.Append(data.refreshSecurityRecurringSchedule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityRecurringSchedule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityRecurringScheduleModel) refreshSecurityRecurringSchedule(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["days"]; ok {
		m.Days = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Days = types.SetNull(types.StringType)
	}

	if v, ok := o["startTime"]; ok {
		m.StartTime = parseStringValue(v)
	}

	if v, ok := o["endTime"]; ok {
		m.EndTime = parseStringValue(v)
	}

	return diags
}

func (data *resourceSecurityRecurringScheduleModel) getCreateObjectSecurityRecurringSchedule(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Days.IsNull() && !data.Days.IsUnknown() {
		result["days"] = expandSetToStringList(data.Days)
	}

	if !data.StartTime.IsNull() && !data.StartTime.IsUnknown() {
		result["startTime"] = data.StartTime.ValueString()
	}

	if !data.EndTime.IsNull() && !data.EndTime.IsUnknown() {
		result["endTime"] = data.EndTime.ValueString()
	}

	return &result
}

func (data *resourceSecurityRecurringScheduleModel) getUpdateObjectSecurityRecurringSchedule(ctx context.Context, state resourceSecurityRecurringScheduleModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Days.IsNull() && !data.Days.IsUnknown() {
		result["days"] = expandSetToStringList(data.Days)
	}

	if !data.StartTime.IsNull() && !data.StartTime.IsUnknown() {
		result["startTime"] = data.StartTime.ValueString()
	}

	if !data.EndTime.IsNull() && !data.EndTime.IsUnknown() {
		result["endTime"] = data.EndTime.ValueString()
	}

	return &result
}

func (data *resourceSecurityRecurringScheduleModel) getURLObjectSecurityRecurringSchedule(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
