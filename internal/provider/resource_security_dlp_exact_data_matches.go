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
var _ resource.Resource = &resourceSecurityDlpExactDataMatches{}

func newResourceSecurityDlpExactDataMatches() resource.Resource {
	return &resourceSecurityDlpExactDataMatches{}
}

type resourceSecurityDlpExactDataMatches struct {
	fortiClient *FortiClient
}

// resourceSecurityDlpExactDataMatchesModel describes the resource data model.
type resourceSecurityDlpExactDataMatchesModel struct {
	ID                   types.String                                                  `tfsdk:"id"`
	PrimaryKey           types.String                                                  `tfsdk:"primary_key"`
	ExternalResourceData *resourceSecurityDlpExactDataMatchesExternalResourceDataModel `tfsdk:"external_resource_data"`
	Columns              []resourceSecurityDlpExactDataMatchesColumnsModel             `tfsdk:"columns"`
	OptionalCount        types.Float64                                                 `tfsdk:"optional_count"`
}

func (r *resourceSecurityDlpExactDataMatches) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_exact_data_matches"
}

func (r *resourceSecurityDlpExactDataMatches) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"optional_count": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(32),
				},
				Computed: true,
				Optional: true,
			},
			"external_resource_data": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"resource": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"refresh_rate": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.Between(1, 43200),
						},
						Computed: true,
						Optional: true,
					},
					"username": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(64),
						},
						Computed: true,
						Optional: true,
					},
					"password": schema.StringAttribute{
						Sensitive: true,
						Computed:  true,
						Optional:  true,
					},
					"update_method": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("feed", "push"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"columns": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"index": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validator.Between(1, 32),
							},
							Computed: true,
							Optional: true,
						},
						"optional": schema.BoolAttribute{
							Computed: true,
							Optional: true,
						},
						"type": schema.SingleNestedAttribute{
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

func (r *resourceSecurityDlpExactDataMatches) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityDlpExactDataMatches) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityDlpExactDataMatchesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityDlpExactDataMatches(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatches(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityDlpExactDataMatches(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatches(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpExactDataMatches(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpExactDataMatches(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpExactDataMatches) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDlpExactDataMatchesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDlpExactDataMatchesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDlpExactDataMatches(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatches(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityDlpExactDataMatches(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatches(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpExactDataMatches(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpExactDataMatches(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpExactDataMatches) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpExactDataMatchesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatches(ctx, "delete", diags))

	err := c.DeleteSecurityDlpExactDataMatches(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceSecurityDlpExactDataMatches) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpExactDataMatchesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatches(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpExactDataMatches(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpExactDataMatches(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpExactDataMatches) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityDlpExactDataMatchesModel) refreshSecurityDlpExactDataMatches(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["externalResourceData"]; ok {
		m.ExternalResourceData = m.ExternalResourceData.flattenSecurityDlpExactDataMatchesExternalResourceData(ctx, v, &diags)
	}

	if v, ok := o["columns"]; ok {
		m.Columns = m.flattenSecurityDlpExactDataMatchesColumnsList(ctx, v, &diags)
	}

	if v, ok := o["optionalCount"]; ok {
		m.OptionalCount = parseFloat64Value(v)
	}

	return diags
}

func (data *resourceSecurityDlpExactDataMatchesModel) getCreateObjectSecurityDlpExactDataMatches(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if data.ExternalResourceData != nil && !isZeroStruct(*data.ExternalResourceData) {
		result["externalResourceData"] = data.ExternalResourceData.expandSecurityDlpExactDataMatchesExternalResourceData(ctx, diags)
	}

	result["columns"] = data.expandSecurityDlpExactDataMatchesColumnsList(ctx, data.Columns, diags)

	if !data.OptionalCount.IsNull() {
		result["optionalCount"] = data.OptionalCount.ValueFloat64()
	}

	return &result
}

func (data *resourceSecurityDlpExactDataMatchesModel) getUpdateObjectSecurityDlpExactDataMatches(ctx context.Context, state resourceSecurityDlpExactDataMatchesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if data.ExternalResourceData != nil && !isSameStruct(data.ExternalResourceData, state.ExternalResourceData) {
		result["externalResourceData"] = data.ExternalResourceData.expandSecurityDlpExactDataMatchesExternalResourceData(ctx, diags)
	}

	if len(data.Columns) > 0 || !isSameStruct(data.Columns, state.Columns) {
		result["columns"] = data.expandSecurityDlpExactDataMatchesColumnsList(ctx, data.Columns, diags)
	}

	if !data.OptionalCount.IsNull() && !data.OptionalCount.Equal(state.OptionalCount) {
		result["optionalCount"] = data.OptionalCount.ValueFloat64()
	}

	return &result
}

func (data *resourceSecurityDlpExactDataMatchesModel) getURLObjectSecurityDlpExactDataMatches(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityDlpExactDataMatchesExternalResourceDataModel struct {
	Resource     types.String  `tfsdk:"resource"`
	RefreshRate  types.Float64 `tfsdk:"refresh_rate"`
	Username     types.String  `tfsdk:"username"`
	Password     types.String  `tfsdk:"password"`
	UpdateMethod types.String  `tfsdk:"update_method"`
}

type resourceSecurityDlpExactDataMatchesColumnsModel struct {
	Index    types.Float64                                        `tfsdk:"index"`
	Type     *resourceSecurityDlpExactDataMatchesColumnsTypeModel `tfsdk:"type"`
	Optional types.Bool                                           `tfsdk:"optional"`
}

type resourceSecurityDlpExactDataMatchesColumnsTypeModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityDlpExactDataMatchesExternalResourceDataModel) flattenSecurityDlpExactDataMatchesExternalResourceData(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpExactDataMatchesExternalResourceDataModel {
	if input == nil {
		return &resourceSecurityDlpExactDataMatchesExternalResourceDataModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpExactDataMatchesExternalResourceDataModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["resource"]; ok {
		m.Resource = parseStringValue(v)
	}

	if v, ok := o["refreshRate"]; ok {
		m.RefreshRate = parseFloat64Value(v)
	}

	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	if v, ok := o["updateMethod"]; ok {
		m.UpdateMethod = parseStringValue(v)
	}

	return m
}

func (m *resourceSecurityDlpExactDataMatchesColumnsModel) flattenSecurityDlpExactDataMatchesColumns(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpExactDataMatchesColumnsModel {
	if input == nil {
		return &resourceSecurityDlpExactDataMatchesColumnsModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpExactDataMatchesColumnsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["index"]; ok {
		m.Index = parseFloat64Value(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = m.Type.flattenSecurityDlpExactDataMatchesColumnsType(ctx, v, diags)
	}

	if v, ok := o["optional"]; ok {
		m.Optional = parseBoolValue(v)
	}

	return m
}

func (s *resourceSecurityDlpExactDataMatchesModel) flattenSecurityDlpExactDataMatchesColumnsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDlpExactDataMatchesColumnsModel {
	if o == nil {
		return []resourceSecurityDlpExactDataMatchesColumnsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument columns is not type of []interface{}.", "")
		return []resourceSecurityDlpExactDataMatchesColumnsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDlpExactDataMatchesColumnsModel{}
	}

	values := make([]resourceSecurityDlpExactDataMatchesColumnsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDlpExactDataMatchesColumnsModel
		values[i] = *m.flattenSecurityDlpExactDataMatchesColumns(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDlpExactDataMatchesColumnsTypeModel) flattenSecurityDlpExactDataMatchesColumnsType(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpExactDataMatchesColumnsTypeModel {
	if input == nil {
		return &resourceSecurityDlpExactDataMatchesColumnsTypeModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpExactDataMatchesColumnsTypeModel{}
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

func (data *resourceSecurityDlpExactDataMatchesExternalResourceDataModel) expandSecurityDlpExactDataMatchesExternalResourceData(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Resource.IsNull() {
		result["resource"] = data.Resource.ValueString()
	}

	if !data.RefreshRate.IsNull() {
		result["refreshRate"] = data.RefreshRate.ValueFloat64()
	}

	if !data.Username.IsNull() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	if !data.UpdateMethod.IsNull() {
		result["updateMethod"] = data.UpdateMethod.ValueString()
	}

	return result
}

func (data *resourceSecurityDlpExactDataMatchesColumnsModel) expandSecurityDlpExactDataMatchesColumns(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Index.IsNull() {
		result["index"] = data.Index.ValueFloat64()
	}

	if data.Type != nil && !isZeroStruct(*data.Type) {
		result["type"] = data.Type.expandSecurityDlpExactDataMatchesColumnsType(ctx, diags)
	}

	if !data.Optional.IsNull() {
		result["optional"] = data.Optional.ValueBool()
	}

	return result
}

func (s *resourceSecurityDlpExactDataMatchesModel) expandSecurityDlpExactDataMatchesColumnsList(ctx context.Context, l []resourceSecurityDlpExactDataMatchesColumnsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDlpExactDataMatchesColumns(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDlpExactDataMatchesColumnsTypeModel) expandSecurityDlpExactDataMatchesColumnsType(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
