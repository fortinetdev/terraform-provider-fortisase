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
var _ resource.Resource = &resourceSecurityDlpSensors{}

func newResourceSecurityDlpSensors() resource.Resource {
	return &resourceSecurityDlpSensors{}
}

type resourceSecurityDlpSensors struct {
	fortiClient *FortiClient
}

// resourceSecurityDlpSensorsModel describes the resource data model.
type resourceSecurityDlpSensorsModel struct {
	ID                          types.String                                        `tfsdk:"id"`
	PrimaryKey                  types.String                                        `tfsdk:"primary_key"`
	EntryMatchesToTriggerSensor types.String                                        `tfsdk:"entry_matches_to_trigger_sensor"`
	SensorDictionaries          []resourceSecurityDlpSensorsSensorDictionariesModel `tfsdk:"sensor_dictionaries"`
}

func (r *resourceSecurityDlpSensors) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_sensors"
}

func (r *resourceSecurityDlpSensors) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthBetween(1, 64),
				},
				Required: true,
			},
			"entry_matches_to_trigger_sensor": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("any", "all"),
				},
				Computed: true,
				Optional: true,
			},
			"sensor_dictionaries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dictionary_id": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validator.Between(1, 32),
							},
							Computed: true,
							Optional: true,
						},
						"dictionary_matches_to_consider_risk": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validator.AtMost(255),
							},
							Computed: true,
							Optional: true,
						},
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"dictionary": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityDlpSensors) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityDlpSensors) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityDlpSensorsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityDlpSensors(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpSensors(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityDlpSensors(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpSensors(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpSensors(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpSensors(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpSensors) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDlpSensorsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDlpSensorsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDlpSensors(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpSensors(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityDlpSensors(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpSensors(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpSensors(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpSensors(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpSensors) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpSensorsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpSensors(ctx, "delete", diags))

	err := c.DeleteSecurityDlpSensors(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceSecurityDlpSensors) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpSensorsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpSensors(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpSensors(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpSensors(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpSensors) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityDlpSensorsModel) refreshSecurityDlpSensors(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["entryMatchesToTriggerSensor"]; ok {
		m.EntryMatchesToTriggerSensor = parseStringValue(v)
	}

	if v, ok := o["sensorDictionaries"]; ok {
		m.SensorDictionaries = m.flattenSecurityDlpSensorsSensorDictionariesList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityDlpSensorsModel) getCreateObjectSecurityDlpSensors(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.EntryMatchesToTriggerSensor.IsNull() {
		result["entryMatchesToTriggerSensor"] = data.EntryMatchesToTriggerSensor.ValueString()
	}

	result["sensorDictionaries"] = data.expandSecurityDlpSensorsSensorDictionariesList(ctx, data.SensorDictionaries, diags)

	return &result
}

func (data *resourceSecurityDlpSensorsModel) getUpdateObjectSecurityDlpSensors(ctx context.Context, state resourceSecurityDlpSensorsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.EntryMatchesToTriggerSensor.IsNull() {
		result["entryMatchesToTriggerSensor"] = data.EntryMatchesToTriggerSensor.ValueString()
	}

	if len(data.SensorDictionaries) > 0 || !isSameStruct(data.SensorDictionaries, state.SensorDictionaries) {
		result["sensorDictionaries"] = data.expandSecurityDlpSensorsSensorDictionariesList(ctx, data.SensorDictionaries, diags)
	}

	return &result
}

func (data *resourceSecurityDlpSensorsModel) getURLObjectSecurityDlpSensors(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityDlpSensorsSensorDictionariesModel struct {
	DictionaryId                    types.Float64                                                `tfsdk:"dictionary_id"`
	Dictionary                      *resourceSecurityDlpSensorsSensorDictionariesDictionaryModel `tfsdk:"dictionary"`
	DictionaryMatchesToConsiderRisk types.Float64                                                `tfsdk:"dictionary_matches_to_consider_risk"`
	Status                          types.String                                                 `tfsdk:"status"`
}

type resourceSecurityDlpSensorsSensorDictionariesDictionaryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityDlpSensorsSensorDictionariesModel) flattenSecurityDlpSensorsSensorDictionaries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpSensorsSensorDictionariesModel {
	if input == nil {
		return &resourceSecurityDlpSensorsSensorDictionariesModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpSensorsSensorDictionariesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["dictionaryId"]; ok {
		m.DictionaryId = parseFloat64Value(v)
	}

	if v, ok := o["dictionary"]; ok {
		m.Dictionary = m.Dictionary.flattenSecurityDlpSensorsSensorDictionariesDictionary(ctx, v, diags)
	}

	if v, ok := o["dictionaryMatchesToConsiderRisk"]; ok {
		m.DictionaryMatchesToConsiderRisk = parseFloat64Value(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	return m
}

func (s *resourceSecurityDlpSensorsModel) flattenSecurityDlpSensorsSensorDictionariesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDlpSensorsSensorDictionariesModel {
	if o == nil {
		return []resourceSecurityDlpSensorsSensorDictionariesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sensor_dictionaries is not type of []interface{}.", "")
		return []resourceSecurityDlpSensorsSensorDictionariesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDlpSensorsSensorDictionariesModel{}
	}

	values := make([]resourceSecurityDlpSensorsSensorDictionariesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDlpSensorsSensorDictionariesModel
		values[i] = *m.flattenSecurityDlpSensorsSensorDictionaries(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDlpSensorsSensorDictionariesDictionaryModel) flattenSecurityDlpSensorsSensorDictionariesDictionary(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpSensorsSensorDictionariesDictionaryModel {
	if input == nil {
		return &resourceSecurityDlpSensorsSensorDictionariesDictionaryModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpSensorsSensorDictionariesDictionaryModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (data *resourceSecurityDlpSensorsSensorDictionariesModel) expandSecurityDlpSensorsSensorDictionaries(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.DictionaryId.IsNull() {
		result["dictionaryId"] = data.DictionaryId.ValueFloat64()
	}

	if data.Dictionary != nil && !isZeroStruct(*data.Dictionary) {
		result["dictionary"] = data.Dictionary.expandSecurityDlpSensorsSensorDictionariesDictionary(ctx, diags)
	}

	if !data.DictionaryMatchesToConsiderRisk.IsNull() {
		result["dictionaryMatchesToConsiderRisk"] = data.DictionaryMatchesToConsiderRisk.ValueFloat64()
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	return result
}

func (s *resourceSecurityDlpSensorsModel) expandSecurityDlpSensorsSensorDictionariesList(ctx context.Context, l []resourceSecurityDlpSensorsSensorDictionariesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDlpSensorsSensorDictionaries(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDlpSensorsSensorDictionariesDictionaryModel) expandSecurityDlpSensorsSensorDictionariesDictionary(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
