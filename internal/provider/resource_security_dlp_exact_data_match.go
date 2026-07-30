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
var _ resource.Resource = &resourceSecurityDlpExactDataMatch{}
var _ resource.ResourceWithMoveState = &resourceSecurityDlpExactDataMatch{}

func newResourceSecurityDlpExactDataMatch() resource.Resource {
	return &resourceSecurityDlpExactDataMatch{}
}

type resourceSecurityDlpExactDataMatch struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityDlpExactDataMatchModel describes the resource data model.
type resourceSecurityDlpExactDataMatchModel struct {
	ID                   types.String                                                `tfsdk:"id"`
	PrimaryKey           types.String                                                `tfsdk:"primary_key"`
	ExternalResourceData *resourceSecurityDlpExactDataMatchExternalResourceDataModel `tfsdk:"external_resource_data"`
	Columns              []resourceSecurityDlpExactDataMatchColumnsModel             `tfsdk:"columns"`
	OptionalCount        types.Float64                                               `tfsdk:"optional_count"`
}

func (r *resourceSecurityDlpExactDataMatch) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_exact_data_match"
}

func (r *resourceSecurityDlpExactDataMatch) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DLP Exact Data Match Resource API V2 for FortiSASE.",
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"optional_count": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(32),
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
							float64validatorwarning.Between(1, 43200),
						},
						Computed: true,
						Optional: true,
					},
					"username": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.LengthAtMost(64),
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
							stringvalidatorwarning.OneOf("feed", "push"),
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
								float64validatorwarning.Between(1, 32),
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

func (r *resourceSecurityDlpExactDataMatch) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_exact_data_match"
}
func (r *resourceSecurityDlpExactDataMatch) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_dlp_exact_data_matches" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityDlpExactDataMatchModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityDlpExactDataMatch) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpExactDataMatches")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityDlpExactDataMatchModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityDlpExactDataMatch(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatch(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityDlpExactDataMatches(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatch(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpExactDataMatches(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpExactDataMatch(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpExactDataMatch) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpExactDataMatches")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDlpExactDataMatchModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDlpExactDataMatchModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDlpExactDataMatch(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatch(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityDlpExactDataMatches(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatch(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpExactDataMatches(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpExactDataMatch(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpExactDataMatch) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpExactDataMatches")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityDlpExactDataMatchModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatch(ctx, "delete", diags))

	output, err := c.DeleteSecurityDlpExactDataMatches(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityDlpExactDataMatch) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpExactDataMatchModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatch(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpExactDataMatches(&input_model)
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

	diags.Append(data.refreshSecurityDlpExactDataMatch(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpExactDataMatch) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityDlpExactDataMatchModel) refreshSecurityDlpExactDataMatch(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["externalResourceData"]; ok {
		m.ExternalResourceData = m.ExternalResourceData.flattenSecurityDlpExactDataMatchExternalResourceData(ctx, v, &diags)
	}

	if v, ok := o["columns"]; ok {
		m.Columns = m.flattenSecurityDlpExactDataMatchColumnsList(ctx, v, &diags)
	}

	if v, ok := o["optionalCount"]; ok {
		m.OptionalCount = parseFloat64Value(v)
	}

	return diags
}

func (data *resourceSecurityDlpExactDataMatchModel) getCreateObjectSecurityDlpExactDataMatch(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if data.ExternalResourceData != nil && !isZeroStruct(*data.ExternalResourceData) {
		result["externalResourceData"] = data.ExternalResourceData.expandSecurityDlpExactDataMatchExternalResourceData(ctx, diags)
	}

	result["columns"] = data.expandSecurityDlpExactDataMatchColumnsList(ctx, data.Columns, diags)

	if !data.OptionalCount.IsNull() && !data.OptionalCount.IsUnknown() {
		result["optionalCount"] = data.OptionalCount.ValueFloat64()
	}

	return &result
}

func (data *resourceSecurityDlpExactDataMatchModel) getUpdateObjectSecurityDlpExactDataMatch(ctx context.Context, state resourceSecurityDlpExactDataMatchModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if data.ExternalResourceData != nil {
		result["externalResourceData"] = data.ExternalResourceData.expandSecurityDlpExactDataMatchExternalResourceData(ctx, diags)
	}

	if data.Columns != nil {
		result["columns"] = data.expandSecurityDlpExactDataMatchColumnsList(ctx, data.Columns, diags)
	}

	if !data.OptionalCount.IsNull() && !data.OptionalCount.IsUnknown() {
		result["optionalCount"] = data.OptionalCount.ValueFloat64()
	}

	return &result
}

func (data *resourceSecurityDlpExactDataMatchModel) getURLObjectSecurityDlpExactDataMatch(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityDlpExactDataMatchExternalResourceDataModel struct {
	Resource     types.String  `tfsdk:"resource"`
	RefreshRate  types.Float64 `tfsdk:"refresh_rate"`
	Username     types.String  `tfsdk:"username"`
	Password     types.String  `tfsdk:"password"`
	UpdateMethod types.String  `tfsdk:"update_method"`
}

type resourceSecurityDlpExactDataMatchColumnsModel struct {
	Index    types.Float64                                      `tfsdk:"index"`
	Type     *resourceSecurityDlpExactDataMatchColumnsTypeModel `tfsdk:"type"`
	Optional types.Bool                                         `tfsdk:"optional"`
}

type resourceSecurityDlpExactDataMatchColumnsTypeModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityDlpExactDataMatchExternalResourceDataModel) flattenSecurityDlpExactDataMatchExternalResourceData(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpExactDataMatchExternalResourceDataModel {
	if input == nil {
		return &resourceSecurityDlpExactDataMatchExternalResourceDataModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpExactDataMatchExternalResourceDataModel{}
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

func (m *resourceSecurityDlpExactDataMatchColumnsModel) flattenSecurityDlpExactDataMatchColumns(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpExactDataMatchColumnsModel {
	if input == nil {
		return &resourceSecurityDlpExactDataMatchColumnsModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpExactDataMatchColumnsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["index"]; ok {
		m.Index = parseFloat64Value(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = m.Type.flattenSecurityDlpExactDataMatchColumnsType(ctx, v, diags)
	}

	if v, ok := o["optional"]; ok {
		m.Optional = parseBoolValue(v)
	}

	return m
}

func (s *resourceSecurityDlpExactDataMatchModel) flattenSecurityDlpExactDataMatchColumnsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDlpExactDataMatchColumnsModel {
	if o == nil {
		return []resourceSecurityDlpExactDataMatchColumnsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument columns is not type of []interface{}.", "")
		return []resourceSecurityDlpExactDataMatchColumnsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDlpExactDataMatchColumnsModel{}
	}

	values := make([]resourceSecurityDlpExactDataMatchColumnsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDlpExactDataMatchColumnsModel
		if i < len(s.Columns) {
			m = s.Columns[i]
		}
		values[i] = *m.flattenSecurityDlpExactDataMatchColumns(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDlpExactDataMatchColumnsTypeModel) flattenSecurityDlpExactDataMatchColumnsType(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpExactDataMatchColumnsTypeModel {
	if input == nil {
		return &resourceSecurityDlpExactDataMatchColumnsTypeModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpExactDataMatchColumnsTypeModel{}
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

func (data *resourceSecurityDlpExactDataMatchExternalResourceDataModel) expandSecurityDlpExactDataMatchExternalResourceData(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Resource.IsNull() && !data.Resource.IsUnknown() {
		result["resource"] = data.Resource.ValueString()
	}

	if !data.RefreshRate.IsNull() && !data.RefreshRate.IsUnknown() {
		result["refreshRate"] = data.RefreshRate.ValueFloat64()
	}

	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		result["password"] = data.Password.ValueString()
	}

	if !data.UpdateMethod.IsNull() && !data.UpdateMethod.IsUnknown() {
		result["updateMethod"] = data.UpdateMethod.ValueString()
	}

	return result
}

func (data *resourceSecurityDlpExactDataMatchColumnsModel) expandSecurityDlpExactDataMatchColumns(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Index.IsNull() && !data.Index.IsUnknown() {
		result["index"] = data.Index.ValueFloat64()
	}

	result["type"] = nil
	if data.Type != nil && !isZeroStruct(*data.Type) {
		result["type"] = data.Type.expandSecurityDlpExactDataMatchColumnsType(ctx, diags)
	}

	if !data.Optional.IsNull() && !data.Optional.IsUnknown() {
		result["optional"] = data.Optional.ValueBool()
	}

	return result
}

func (s *resourceSecurityDlpExactDataMatchModel) expandSecurityDlpExactDataMatchColumnsList(ctx context.Context, l []resourceSecurityDlpExactDataMatchColumnsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDlpExactDataMatchColumns(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDlpExactDataMatchColumnsTypeModel) expandSecurityDlpExactDataMatchColumnsType(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
