// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
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
var _ resource.Resource = &resourceSecurityDlpDictionary{}
var _ resource.ResourceWithMoveState = &resourceSecurityDlpDictionary{}

func newResourceSecurityDlpDictionary() resource.Resource {
	return &resourceSecurityDlpDictionary{}
}

type resourceSecurityDlpDictionary struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityDlpDictionaryModel describes the resource data model.
type resourceSecurityDlpDictionaryModel struct {
	ID                types.String                                `tfsdk:"id"`
	PrimaryKey        types.String                                `tfsdk:"primary_key"`
	EntriesToEvaluate types.String                                `tfsdk:"entries_to_evaluate"`
	DictionaryType    types.String                                `tfsdk:"dictionary_type"`
	Entries           []resourceSecurityDlpDictionaryEntriesModel `tfsdk:"entries"`
}

func (r *resourceSecurityDlpDictionary) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_dictionary"
}

func (r *resourceSecurityDlpDictionary) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DLP Dictionary Resource API V2 for FortiSASE.",
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
			"entries_to_evaluate": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("all", "any"),
				},
				Computed: true,
				Optional: true,
			},
			"dictionary_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("mip-label", "sensor"),
				},
				MarkdownDescription: "This property is used to classify DLP Dictionaries. It is server-generated and cannot be modified.\nSupported values: mip-label, sensor.",
				Computed:            true,
				Optional:            true,
			},
			"entries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"repeat": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"pattern": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"case_sensitive": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"dlp_data_type": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("security/dlp-data-types"),
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

func (r *resourceSecurityDlpDictionary) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_dictionary"
}
func (r *resourceSecurityDlpDictionary) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_dlp_dictionaries" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityDlpDictionaryModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityDlpDictionary) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpDictionaries")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityDlpDictionaryModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityDlpDictionary(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpDictionary(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityDlpDictionaries(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpDictionary(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpDictionaries(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpDictionary(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpDictionary) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpDictionaries")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDlpDictionaryModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDlpDictionaryModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDlpDictionary(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpDictionary(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityDlpDictionaries(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpDictionary(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpDictionaries(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpDictionary(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpDictionary) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpDictionaries")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityDlpDictionaryModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpDictionary(ctx, "delete", diags))

	output, err := c.DeleteSecurityDlpDictionaries(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityDlpDictionary) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpDictionaryModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpDictionary(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpDictionaries(&input_model)
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

	diags.Append(data.refreshSecurityDlpDictionary(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpDictionary) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityDlpDictionaryModel) refreshSecurityDlpDictionary(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["entriesToEvaluate"]; ok {
		m.EntriesToEvaluate = parseStringValue(v)
	}

	if v, ok := o["dictionaryType"]; ok {
		m.DictionaryType = parseStringValue(v)
	}

	if v, ok := o["entries"]; ok {
		m.Entries = m.flattenSecurityDlpDictionaryEntriesList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityDlpDictionaryModel) getCreateObjectSecurityDlpDictionary(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.EntriesToEvaluate.IsNull() && !data.EntriesToEvaluate.IsUnknown() {
		result["entriesToEvaluate"] = data.EntriesToEvaluate.ValueString()
	}

	if !data.DictionaryType.IsNull() && !data.DictionaryType.IsUnknown() {
		diags.AddWarning("\"dictionary_type\" is deprecated and may be removed in future.",
			"It is recommended to recreate the resource without \"dictionary_type\" to avoid unexpected behavior in future.",
		)
		result["dictionaryType"] = data.DictionaryType.ValueString()
	}

	if data.Entries != nil {
		result["entries"] = data.expandSecurityDlpDictionaryEntriesList(ctx, data.Entries, diags)
	}

	return &result
}

func (data *resourceSecurityDlpDictionaryModel) getUpdateObjectSecurityDlpDictionary(ctx context.Context, state resourceSecurityDlpDictionaryModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.EntriesToEvaluate.IsNull() && !data.EntriesToEvaluate.IsUnknown() {
		result["entriesToEvaluate"] = data.EntriesToEvaluate.ValueString()
	}

	if !data.DictionaryType.IsNull() && !data.DictionaryType.IsUnknown() {
		result["dictionaryType"] = data.DictionaryType.ValueString()
	}

	if data.Entries != nil {
		result["entries"] = data.expandSecurityDlpDictionaryEntriesList(ctx, data.Entries, diags)
	}

	return &result
}

func (data *resourceSecurityDlpDictionaryModel) getURLObjectSecurityDlpDictionary(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityDlpDictionaryEntriesModel struct {
	DlpDataType   *resourceSecurityDlpDictionaryEntriesDlpDataTypeModel `tfsdk:"dlp_data_type"`
	Status        types.String                                          `tfsdk:"status"`
	Repeat        types.String                                          `tfsdk:"repeat"`
	Pattern       types.String                                          `tfsdk:"pattern"`
	CaseSensitive types.String                                          `tfsdk:"case_sensitive"`
}

type resourceSecurityDlpDictionaryEntriesDlpDataTypeModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityDlpDictionaryEntriesModel) flattenSecurityDlpDictionaryEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpDictionaryEntriesModel {
	if input == nil {
		return &resourceSecurityDlpDictionaryEntriesModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpDictionaryEntriesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["dlpDataType"]; ok {
		m.DlpDataType = m.DlpDataType.flattenSecurityDlpDictionaryEntriesDlpDataType(ctx, v, diags)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["repeat"]; ok {
		m.Repeat = parseStringValue(v)
	}

	if v, ok := o["pattern"]; ok {
		m.Pattern = parseStringValue(v)
	}

	if v, ok := o["caseSensitive"]; ok {
		m.CaseSensitive = parseStringValue(v)
	}

	return m
}

func (s *resourceSecurityDlpDictionaryModel) flattenSecurityDlpDictionaryEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDlpDictionaryEntriesModel {
	if o == nil {
		return []resourceSecurityDlpDictionaryEntriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument entries is not type of []interface{}.", "")
		return []resourceSecurityDlpDictionaryEntriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDlpDictionaryEntriesModel{}
	}

	values := make([]resourceSecurityDlpDictionaryEntriesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDlpDictionaryEntriesModel
		if i < len(s.Entries) {
			m = s.Entries[i]
		}
		values[i] = *m.flattenSecurityDlpDictionaryEntries(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDlpDictionaryEntriesDlpDataTypeModel) flattenSecurityDlpDictionaryEntriesDlpDataType(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpDictionaryEntriesDlpDataTypeModel {
	if input == nil {
		return &resourceSecurityDlpDictionaryEntriesDlpDataTypeModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpDictionaryEntriesDlpDataTypeModel{}
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

func (data *resourceSecurityDlpDictionaryEntriesModel) expandSecurityDlpDictionaryEntries(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	result["dlpDataType"] = nil
	if data.DlpDataType != nil && !isZeroStruct(*data.DlpDataType) {
		result["dlpDataType"] = data.DlpDataType.expandSecurityDlpDictionaryEntriesDlpDataType(ctx, diags)
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Repeat.IsNull() && !data.Repeat.IsUnknown() {
		result["repeat"] = data.Repeat.ValueString()
	}

	if !data.Pattern.IsNull() && !data.Pattern.IsUnknown() {
		result["pattern"] = data.Pattern.ValueString()
	}

	if !data.CaseSensitive.IsNull() && !data.CaseSensitive.IsUnknown() {
		result["caseSensitive"] = data.CaseSensitive.ValueString()
	}

	return result
}

func (s *resourceSecurityDlpDictionaryModel) expandSecurityDlpDictionaryEntriesList(ctx context.Context, l []resourceSecurityDlpDictionaryEntriesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDlpDictionaryEntries(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDlpDictionaryEntriesDlpDataTypeModel) expandSecurityDlpDictionaryEntriesDlpDataType(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
