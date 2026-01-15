// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
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
var _ resource.Resource = &resourceSecurityDlpDictionaries{}

func newResourceSecurityDlpDictionaries() resource.Resource {
	return &resourceSecurityDlpDictionaries{}
}

type resourceSecurityDlpDictionaries struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityDlpDictionariesModel describes the resource data model.
type resourceSecurityDlpDictionariesModel struct {
	ID                   types.String                                  `tfsdk:"id"`
	PrimaryKey           types.String                                  `tfsdk:"primary_key"`
	DictionaryType       types.String                                  `tfsdk:"dictionary_type"`
	SensitivityLabelGuid types.String                                  `tfsdk:"sensitivity_label_guid"`
	EntriesToEvaluate    types.String                                  `tfsdk:"entries_to_evaluate"`
	Entries              []resourceSecurityDlpDictionariesEntriesModel `tfsdk:"entries"`
}

func (r *resourceSecurityDlpDictionaries) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_dictionaries"
}

func (r *resourceSecurityDlpDictionaries) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Required: true,
			},
			"dictionary_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("sensor", "mip-label"),
				},
				Computed: true,
				Optional: true,
			},
			"sensitivity_label_guid": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"entries_to_evaluate": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("all", "any"),
				},
				Computed: true,
				Optional: true,
			},
			"entries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"repeat": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
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
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"dlp_data_type": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("security/dlp-data-types"),
									},
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

func (r *resourceSecurityDlpDictionaries) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_dictionaries"
}

func (r *resourceSecurityDlpDictionaries) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpDictionaries")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityDlpDictionariesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityDlpDictionaries(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpDictionaries(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityDlpDictionaries(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpDictionaries(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpDictionaries(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpDictionaries(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpDictionaries) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpDictionaries")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDlpDictionariesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDlpDictionariesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDlpDictionaries(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpDictionaries(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpDictionaries(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpDictionaries(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpDictionaries(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpDictionaries) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpDictionaries")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityDlpDictionariesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpDictionaries(ctx, "delete", diags))

	output, err := c.DeleteSecurityDlpDictionaries(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityDlpDictionaries) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpDictionariesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpDictionaries(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpDictionaries(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpDictionaries(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpDictionaries) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityDlpDictionariesModel) refreshSecurityDlpDictionaries(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["dictionaryType"]; ok {
		m.DictionaryType = parseStringValue(v)
	}

	if v, ok := o["sensitivityLabelGuid"]; ok {
		m.SensitivityLabelGuid = parseStringValue(v)
	}

	if v, ok := o["entriesToEvaluate"]; ok {
		m.EntriesToEvaluate = parseStringValue(v)
	}

	if v, ok := o["entries"]; ok {
		m.Entries = m.flattenSecurityDlpDictionariesEntriesList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityDlpDictionariesModel) getCreateObjectSecurityDlpDictionaries(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.DictionaryType.IsNull() {
		result["dictionaryType"] = data.DictionaryType.ValueString()
	}

	if !data.SensitivityLabelGuid.IsNull() {
		result["sensitivityLabelGuid"] = data.SensitivityLabelGuid.ValueString()
	}

	if !data.EntriesToEvaluate.IsNull() {
		result["entriesToEvaluate"] = data.EntriesToEvaluate.ValueString()
	}

	if data.Entries != nil {
		result["entries"] = data.expandSecurityDlpDictionariesEntriesList(ctx, data.Entries, diags)
	}

	return &result
}

func (data *resourceSecurityDlpDictionariesModel) getUpdateObjectSecurityDlpDictionaries(ctx context.Context, state resourceSecurityDlpDictionariesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.DictionaryType.IsNull() {
		result["dictionaryType"] = data.DictionaryType.ValueString()
	}

	if !data.SensitivityLabelGuid.IsNull() {
		result["sensitivityLabelGuid"] = data.SensitivityLabelGuid.ValueString()
	}

	if !data.EntriesToEvaluate.IsNull() {
		result["entriesToEvaluate"] = data.EntriesToEvaluate.ValueString()
	}

	if data.Entries != nil {
		result["entries"] = data.expandSecurityDlpDictionariesEntriesList(ctx, data.Entries, diags)
	}

	return &result
}

func (data *resourceSecurityDlpDictionariesModel) getURLObjectSecurityDlpDictionaries(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityDlpDictionariesEntriesModel struct {
	DlpDataType   *resourceSecurityDlpDictionariesEntriesDlpDataTypeModel `tfsdk:"dlp_data_type"`
	Status        types.String                                            `tfsdk:"status"`
	Repeat        types.String                                            `tfsdk:"repeat"`
	Pattern       types.String                                            `tfsdk:"pattern"`
	CaseSensitive types.String                                            `tfsdk:"case_sensitive"`
}

type resourceSecurityDlpDictionariesEntriesDlpDataTypeModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityDlpDictionariesEntriesModel) flattenSecurityDlpDictionariesEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpDictionariesEntriesModel {
	if input == nil {
		return &resourceSecurityDlpDictionariesEntriesModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpDictionariesEntriesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["dlpDataType"]; ok {
		m.DlpDataType = m.DlpDataType.flattenSecurityDlpDictionariesEntriesDlpDataType(ctx, v, diags)
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

func (s *resourceSecurityDlpDictionariesModel) flattenSecurityDlpDictionariesEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDlpDictionariesEntriesModel {
	if o == nil {
		return []resourceSecurityDlpDictionariesEntriesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument entries is not type of []interface{}.", "")
		return []resourceSecurityDlpDictionariesEntriesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDlpDictionariesEntriesModel{}
	}

	values := make([]resourceSecurityDlpDictionariesEntriesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDlpDictionariesEntriesModel
		values[i] = *m.flattenSecurityDlpDictionariesEntries(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDlpDictionariesEntriesDlpDataTypeModel) flattenSecurityDlpDictionariesEntriesDlpDataType(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpDictionariesEntriesDlpDataTypeModel {
	if input == nil {
		return &resourceSecurityDlpDictionariesEntriesDlpDataTypeModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpDictionariesEntriesDlpDataTypeModel{}
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

func (data *resourceSecurityDlpDictionariesEntriesModel) expandSecurityDlpDictionariesEntries(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.DlpDataType != nil && !isZeroStruct(*data.DlpDataType) {
		result["dlpDataType"] = data.DlpDataType.expandSecurityDlpDictionariesEntriesDlpDataType(ctx, diags)
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Repeat.IsNull() {
		result["repeat"] = data.Repeat.ValueString()
	}

	if !data.Pattern.IsNull() {
		result["pattern"] = data.Pattern.ValueString()
	}

	if !data.CaseSensitive.IsNull() {
		result["caseSensitive"] = data.CaseSensitive.ValueString()
	}

	return result
}

func (s *resourceSecurityDlpDictionariesModel) expandSecurityDlpDictionariesEntriesList(ctx context.Context, l []resourceSecurityDlpDictionariesEntriesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDlpDictionariesEntries(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDlpDictionariesEntriesDlpDataTypeModel) expandSecurityDlpDictionariesEntriesDlpDataType(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
