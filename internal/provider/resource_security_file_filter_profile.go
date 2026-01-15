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
var _ resource.Resource = &resourceSecurityFileFilterProfile{}

func newResourceSecurityFileFilterProfile() resource.Resource {
	return &resourceSecurityFileFilterProfile{}
}

type resourceSecurityFileFilterProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityFileFilterProfileModel describes the resource data model.
type resourceSecurityFileFilterProfileModel struct {
	ID                          types.String                                    `tfsdk:"id"`
	PrimaryKey                  types.String                                    `tfsdk:"primary_key"`
	Block                       []resourceSecurityFileFilterProfileBlockModel   `tfsdk:"block"`
	Monitor                     []resourceSecurityFileFilterProfileMonitorModel `tfsdk:"monitor"`
	BlockPasswordProtectedFiles types.Bool                                      `tfsdk:"block_password_protected_files"`
	Direction                   types.String                                    `tfsdk:"direction"`
}

func (r *resourceSecurityFileFilterProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_file_filter_profile"
}

func (r *resourceSecurityFileFilterProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"block_password_protected_files": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				MarkdownDescription: "The direction of the target resource.\nSupported values: internal-profiles, outbound-profiles.",
				Computed:            true,
				Optional:            true,
			},
			"block": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("security/antivirus-filetypes"),
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"monitor": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("security/antivirus-filetypes"),
							},
							Computed: true,
							Optional: true,
						},
						"primary_key": schema.StringAttribute{
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

func (r *resourceSecurityFileFilterProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_file_filter_profile"
}

func (r *resourceSecurityFileFilterProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityFileFilterProfile")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityFileFilterProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecurityFileFilterProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityFileFilterProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecurityFileFilterProfile(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityFileFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityFileFilterProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityFileFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityFileFilterProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityFileFilterProfile")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityFileFilterProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityFileFilterProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityFileFilterProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityFileFilterProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityFileFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityFileFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityFileFilterProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityFileFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityFileFilterProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityFileFilterProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityFileFilterProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityFileFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityFileFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityFileFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityFileFilterProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityFileFilterProfileModel) refreshSecurityFileFilterProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["block"]; ok {
		m.Block = m.flattenSecurityFileFilterProfileBlockList(ctx, v, &diags)
	}

	if v, ok := o["monitor"]; ok {
		m.Monitor = m.flattenSecurityFileFilterProfileMonitorList(ctx, v, &diags)
	}

	if v, ok := o["blockPasswordProtectedFiles"]; ok {
		m.BlockPasswordProtectedFiles = parseBoolValue(v)
	}

	return diags
}

func (data *resourceSecurityFileFilterProfileModel) getCreateObjectSecurityFileFilterProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	result["block"] = data.expandSecurityFileFilterProfileBlockList(ctx, data.Block, diags)

	result["monitor"] = data.expandSecurityFileFilterProfileMonitorList(ctx, data.Monitor, diags)

	if !data.BlockPasswordProtectedFiles.IsNull() {
		result["blockPasswordProtectedFiles"] = data.BlockPasswordProtectedFiles.ValueBool()
	}

	return &result
}

func (data *resourceSecurityFileFilterProfileModel) getUpdateObjectSecurityFileFilterProfile(ctx context.Context, state resourceSecurityFileFilterProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if data.Block != nil {
		result["block"] = data.expandSecurityFileFilterProfileBlockList(ctx, data.Block, diags)
	}

	if data.Monitor != nil {
		result["monitor"] = data.expandSecurityFileFilterProfileMonitorList(ctx, data.Monitor, diags)
	}

	if !data.BlockPasswordProtectedFiles.IsNull() {
		result["blockPasswordProtectedFiles"] = data.BlockPasswordProtectedFiles.ValueBool()
	}

	return &result
}

func (data *resourceSecurityFileFilterProfileModel) getURLObjectSecurityFileFilterProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		diags.AddWarning("\"direction\" is deprecated and may be removed in future.",
			"It is recommended to recreate the resource without \"direction\" to avoid unexpected behavior in future.",
		)
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityFileFilterProfileBlockModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityFileFilterProfileMonitorModel struct {
	Datasource types.String `tfsdk:"datasource"`
	PrimaryKey types.String `tfsdk:"primary_key"`
}

func (m *resourceSecurityFileFilterProfileBlockModel) flattenSecurityFileFilterProfileBlock(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityFileFilterProfileBlockModel {
	if input == nil {
		return &resourceSecurityFileFilterProfileBlockModel{}
	}
	if m == nil {
		m = &resourceSecurityFileFilterProfileBlockModel{}
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

func (s *resourceSecurityFileFilterProfileModel) flattenSecurityFileFilterProfileBlockList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityFileFilterProfileBlockModel {
	if o == nil {
		return []resourceSecurityFileFilterProfileBlockModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument block is not type of []interface{}.", "")
		return []resourceSecurityFileFilterProfileBlockModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityFileFilterProfileBlockModel{}
	}

	values := make([]resourceSecurityFileFilterProfileBlockModel, len(l))
	for i, ele := range l {
		var m resourceSecurityFileFilterProfileBlockModel
		values[i] = *m.flattenSecurityFileFilterProfileBlock(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityFileFilterProfileMonitorModel) flattenSecurityFileFilterProfileMonitor(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityFileFilterProfileMonitorModel {
	if input == nil {
		return &resourceSecurityFileFilterProfileMonitorModel{}
	}
	if m == nil {
		m = &resourceSecurityFileFilterProfileMonitorModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	return m
}

func (s *resourceSecurityFileFilterProfileModel) flattenSecurityFileFilterProfileMonitorList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityFileFilterProfileMonitorModel {
	if o == nil {
		return []resourceSecurityFileFilterProfileMonitorModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument monitor is not type of []interface{}.", "")
		return []resourceSecurityFileFilterProfileMonitorModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityFileFilterProfileMonitorModel{}
	}

	values := make([]resourceSecurityFileFilterProfileMonitorModel, len(l))
	for i, ele := range l {
		var m resourceSecurityFileFilterProfileMonitorModel
		values[i] = *m.flattenSecurityFileFilterProfileMonitor(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityFileFilterProfileBlockModel) expandSecurityFileFilterProfileBlock(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityFileFilterProfileModel) expandSecurityFileFilterProfileBlockList(ctx context.Context, l []resourceSecurityFileFilterProfileBlockModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityFileFilterProfileBlock(ctx, diags)
	}
	return result
}

func (data *resourceSecurityFileFilterProfileMonitorModel) expandSecurityFileFilterProfileMonitor(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return result
}

func (s *resourceSecurityFileFilterProfileModel) expandSecurityFileFilterProfileMonitorList(ctx context.Context, l []resourceSecurityFileFilterProfileMonitorModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityFileFilterProfileMonitor(ctx, diags)
	}
	return result
}
