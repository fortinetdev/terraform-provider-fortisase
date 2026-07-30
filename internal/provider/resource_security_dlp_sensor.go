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
var _ resource.Resource = &resourceSecurityDlpSensor{}
var _ resource.ResourceWithMoveState = &resourceSecurityDlpSensor{}

func newResourceSecurityDlpSensor() resource.Resource {
	return &resourceSecurityDlpSensor{}
}

type resourceSecurityDlpSensor struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityDlpSensorModel describes the resource data model.
type resourceSecurityDlpSensorModel struct {
	ID                          types.String                                       `tfsdk:"id"`
	PrimaryKey                  types.String                                       `tfsdk:"primary_key"`
	EntryMatchesToTriggerSensor types.String                                       `tfsdk:"entry_matches_to_trigger_sensor"`
	SensorDictionaries          []resourceSecurityDlpSensorSensorDictionariesModel `tfsdk:"sensor_dictionaries"`
	Description                 types.String                                       `tfsdk:"description"`
}

func (r *resourceSecurityDlpSensor) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_sensor"
}

func (r *resourceSecurityDlpSensor) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DLP Sensor Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 64),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"entry_matches_to_trigger_sensor": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("any", "all"),
				},
				Computed: true,
				Optional: true,
			},
			"description": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(255),
				},
				Computed: true,
				Optional: true,
			},
			"sensor_dictionaries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dictionary_id": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(1, 32),
							},
							Computed: true,
							Optional: true,
						},
						"dictionary_matches_to_consider_risk": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtMost(255),
							},
							Computed: true,
							Optional: true,
						},
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"dictionary": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("security/dlp-dictionaries", "security/dlp-exact-data-matches"),
									},
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

func (r *resourceSecurityDlpSensor) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_sensor"
}
func (r *resourceSecurityDlpSensor) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_dlp_sensors" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityDlpSensorModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityDlpSensor) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpSensors")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityDlpSensorModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityDlpSensor(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpSensor(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityDlpSensors(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpSensor(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpSensors(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpSensor(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpSensor) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpSensors")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDlpSensorModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDlpSensorModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDlpSensor(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpSensor(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityDlpSensors(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpSensor(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpSensors(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpSensor(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpSensor) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpSensors")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityDlpSensorModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpSensor(ctx, "delete", diags))

	output, err := c.DeleteSecurityDlpSensors(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityDlpSensor) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpSensorModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpSensor(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpSensors(&input_model)
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

	diags.Append(data.refreshSecurityDlpSensor(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpSensor) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityDlpSensorModel) refreshSecurityDlpSensor(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["entryMatchesToTriggerSensor"]; ok {
		m.EntryMatchesToTriggerSensor = parseStringValue(v)
	}

	if v, ok := o["sensorDictionaries"]; ok {
		m.SensorDictionaries = m.flattenSecurityDlpSensorSensorDictionariesList(ctx, v, &diags)
	}

	if v, ok := o["description"]; ok {
		m.Description = parseStringValue(v)
	}

	return diags
}

func (data *resourceSecurityDlpSensorModel) getCreateObjectSecurityDlpSensor(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.EntryMatchesToTriggerSensor.IsNull() && !data.EntryMatchesToTriggerSensor.IsUnknown() {
		result["entryMatchesToTriggerSensor"] = data.EntryMatchesToTriggerSensor.ValueString()
	}

	result["sensorDictionaries"] = data.expandSecurityDlpSensorSensorDictionariesList(ctx, data.SensorDictionaries, diags)

	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		result["description"] = data.Description.ValueString()
	}

	return &result
}

func (data *resourceSecurityDlpSensorModel) getUpdateObjectSecurityDlpSensor(ctx context.Context, state resourceSecurityDlpSensorModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.EntryMatchesToTriggerSensor.IsNull() && !data.EntryMatchesToTriggerSensor.IsUnknown() {
		result["entryMatchesToTriggerSensor"] = data.EntryMatchesToTriggerSensor.ValueString()
	}

	if data.SensorDictionaries != nil {
		result["sensorDictionaries"] = data.expandSecurityDlpSensorSensorDictionariesList(ctx, data.SensorDictionaries, diags)
	}

	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		result["description"] = data.Description.ValueString()
	}

	return &result
}

func (data *resourceSecurityDlpSensorModel) getURLObjectSecurityDlpSensor(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityDlpSensorSensorDictionariesModel struct {
	DictionaryId                    types.Float64                                               `tfsdk:"dictionary_id"`
	Dictionary                      *resourceSecurityDlpSensorSensorDictionariesDictionaryModel `tfsdk:"dictionary"`
	DictionaryMatchesToConsiderRisk types.Float64                                               `tfsdk:"dictionary_matches_to_consider_risk"`
	Status                          types.String                                                `tfsdk:"status"`
}

type resourceSecurityDlpSensorSensorDictionariesDictionaryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityDlpSensorSensorDictionariesModel) flattenSecurityDlpSensorSensorDictionaries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpSensorSensorDictionariesModel {
	if input == nil {
		return &resourceSecurityDlpSensorSensorDictionariesModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpSensorSensorDictionariesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["dictionaryId"]; ok {
		m.DictionaryId = parseFloat64Value(v)
	}

	if v, ok := o["dictionary"]; ok {
		m.Dictionary = m.Dictionary.flattenSecurityDlpSensorSensorDictionariesDictionary(ctx, v, diags)
	}

	if v, ok := o["dictionaryMatchesToConsiderRisk"]; ok {
		m.DictionaryMatchesToConsiderRisk = parseFloat64Value(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	return m
}

func (s *resourceSecurityDlpSensorModel) flattenSecurityDlpSensorSensorDictionariesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDlpSensorSensorDictionariesModel {
	if o == nil {
		return []resourceSecurityDlpSensorSensorDictionariesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sensor_dictionaries is not type of []interface{}.", "")
		return []resourceSecurityDlpSensorSensorDictionariesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDlpSensorSensorDictionariesModel{}
	}

	values := make([]resourceSecurityDlpSensorSensorDictionariesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDlpSensorSensorDictionariesModel
		if i < len(s.SensorDictionaries) {
			m = s.SensorDictionaries[i]
		}
		values[i] = *m.flattenSecurityDlpSensorSensorDictionaries(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDlpSensorSensorDictionariesDictionaryModel) flattenSecurityDlpSensorSensorDictionariesDictionary(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpSensorSensorDictionariesDictionaryModel {
	if input == nil {
		return &resourceSecurityDlpSensorSensorDictionariesDictionaryModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpSensorSensorDictionariesDictionaryModel{}
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

func (data *resourceSecurityDlpSensorSensorDictionariesModel) expandSecurityDlpSensorSensorDictionaries(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.DictionaryId.IsNull() && !data.DictionaryId.IsUnknown() {
		result["dictionaryId"] = data.DictionaryId.ValueFloat64()
	}

	result["dictionary"] = nil
	if data.Dictionary != nil && !isZeroStruct(*data.Dictionary) {
		result["dictionary"] = data.Dictionary.expandSecurityDlpSensorSensorDictionariesDictionary(ctx, diags)
	}

	if !data.DictionaryMatchesToConsiderRisk.IsNull() && !data.DictionaryMatchesToConsiderRisk.IsUnknown() {
		result["dictionaryMatchesToConsiderRisk"] = data.DictionaryMatchesToConsiderRisk.ValueFloat64()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	return result
}

func (s *resourceSecurityDlpSensorModel) expandSecurityDlpSensorSensorDictionariesList(ctx context.Context, l []resourceSecurityDlpSensorSensorDictionariesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDlpSensorSensorDictionaries(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDlpSensorSensorDictionariesDictionaryModel) expandSecurityDlpSensorSensorDictionariesDictionary(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
