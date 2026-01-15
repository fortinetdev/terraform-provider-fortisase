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
var _ resource.Resource = &resourceSecurityDlpFingerprintDatabases{}

func newResourceSecurityDlpFingerprintDatabases() resource.Resource {
	return &resourceSecurityDlpFingerprintDatabases{}
}

type resourceSecurityDlpFingerprintDatabases struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityDlpFingerprintDatabasesModel describes the resource data model.
type resourceSecurityDlpFingerprintDatabasesModel struct {
	ID                            types.String                                                `tfsdk:"id"`
	PrimaryKey                    types.String                                                `tfsdk:"primary_key"`
	Server                        types.String                                                `tfsdk:"server"`
	Sensitivity                   types.String                                                `tfsdk:"sensitivity"`
	IncludeSubdirectories         types.String                                                `tfsdk:"include_subdirectories"`
	ServerDirectory               types.String                                                `tfsdk:"server_directory"`
	FilePattern                   types.String                                                `tfsdk:"file_pattern"`
	Schedule                      *resourceSecurityDlpFingerprintDatabasesScheduleModel       `tfsdk:"schedule"`
	RemoveDeletedFileFingerprints types.String                                                `tfsdk:"remove_deleted_file_fingerprints"`
	KeepModified                  types.String                                                `tfsdk:"keep_modified"`
	ScanOnCreation                types.String                                                `tfsdk:"scan_on_creation"`
	Authentication                *resourceSecurityDlpFingerprintDatabasesAuthenticationModel `tfsdk:"authentication"`
}

func (r *resourceSecurityDlpFingerprintDatabases) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_fingerprint_databases"
}

func (r *resourceSecurityDlpFingerprintDatabases) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(35),
				},
				Computed: true,
				Optional: true,
			},
			"sensitivity": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("Warning", "Private", "Critical"),
				},
				Computed: true,
				Optional: true,
			},
			"include_subdirectories": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"server_directory": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(119),
				},
				Computed: true,
				Optional: true,
			},
			"file_pattern": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(35),
				},
				Computed: true,
				Optional: true,
			},
			"remove_deleted_file_fingerprints": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"keep_modified": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"scan_on_creation": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"schedule": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"period": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("daily", "weekly", "monthly"),
						},
						Computed: true,
						Optional: true,
					},
					"sync_hour": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.AtMost(23),
						},
						Computed: true,
						Optional: true,
					},
					"sync_minute": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.AtMost(59),
						},
						Computed: true,
						Optional: true,
					},
					"weekday": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"),
						},
						Computed: true,
						Optional: true,
					},
					"sync_day_of_the_month": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.Between(1, 31),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"authentication": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"username": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(35),
						},
						Computed: true,
						Optional: true,
					},
					"password": schema.StringAttribute{
						Sensitive: true,
						Computed:  true,
						Optional:  true,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityDlpFingerprintDatabases) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_fingerprint_databases"
}

func (r *resourceSecurityDlpFingerprintDatabases) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpFingerprintDatabases")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityDlpFingerprintDatabasesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityDlpFingerprintDatabases(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpFingerprintDatabases(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityDlpFingerprintDatabases(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpFingerprintDatabases(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFingerprintDatabases(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFingerprintDatabases(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpFingerprintDatabases) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpFingerprintDatabases")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDlpFingerprintDatabasesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDlpFingerprintDatabasesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDlpFingerprintDatabases(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpFingerprintDatabases(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityDlpFingerprintDatabases(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpFingerprintDatabases(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFingerprintDatabases(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFingerprintDatabases(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpFingerprintDatabases) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpFingerprintDatabases")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityDlpFingerprintDatabasesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpFingerprintDatabases(ctx, "delete", diags))

	output, err := c.DeleteSecurityDlpFingerprintDatabases(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityDlpFingerprintDatabases) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpFingerprintDatabasesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpFingerprintDatabases(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFingerprintDatabases(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFingerprintDatabases(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpFingerprintDatabases) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityDlpFingerprintDatabasesModel) refreshSecurityDlpFingerprintDatabases(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["server"]; ok {
		m.Server = parseStringValue(v)
	}

	if v, ok := o["sensitivity"]; ok {
		m.Sensitivity = parseStringValue(v)
	}

	if v, ok := o["includeSubdirectories"]; ok {
		m.IncludeSubdirectories = parseStringValue(v)
	}

	if v, ok := o["serverDirectory"]; ok {
		m.ServerDirectory = parseStringValue(v)
	}

	if v, ok := o["filePattern"]; ok {
		m.FilePattern = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityDlpFingerprintDatabasesSchedule(ctx, v, &diags)
	}

	if v, ok := o["removeDeletedFileFingerprints"]; ok {
		m.RemoveDeletedFileFingerprints = parseStringValue(v)
	}

	if v, ok := o["keepModified"]; ok {
		m.KeepModified = parseStringValue(v)
	}

	if v, ok := o["scanOnCreation"]; ok {
		m.ScanOnCreation = parseStringValue(v)
	}

	if v, ok := o["authentication"]; ok {
		m.Authentication = m.Authentication.flattenSecurityDlpFingerprintDatabasesAuthentication(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityDlpFingerprintDatabasesModel) getCreateObjectSecurityDlpFingerprintDatabases(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Server.IsNull() {
		result["server"] = data.Server.ValueString()
	}

	if !data.Sensitivity.IsNull() {
		result["sensitivity"] = data.Sensitivity.ValueString()
	}

	if !data.IncludeSubdirectories.IsNull() {
		result["includeSubdirectories"] = data.IncludeSubdirectories.ValueString()
	}

	if !data.ServerDirectory.IsNull() {
		result["serverDirectory"] = data.ServerDirectory.ValueString()
	}

	if !data.FilePattern.IsNull() {
		result["filePattern"] = data.FilePattern.ValueString()
	}

	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityDlpFingerprintDatabasesSchedule(ctx, diags)
	}

	if !data.RemoveDeletedFileFingerprints.IsNull() {
		result["removeDeletedFileFingerprints"] = data.RemoveDeletedFileFingerprints.ValueString()
	}

	if !data.KeepModified.IsNull() {
		result["keepModified"] = data.KeepModified.ValueString()
	}

	if !data.ScanOnCreation.IsNull() {
		result["scanOnCreation"] = data.ScanOnCreation.ValueString()
	}

	if data.Authentication != nil && !isZeroStruct(*data.Authentication) {
		result["authentication"] = data.Authentication.expandSecurityDlpFingerprintDatabasesAuthentication(ctx, diags)
	}

	return &result
}

func (data *resourceSecurityDlpFingerprintDatabasesModel) getUpdateObjectSecurityDlpFingerprintDatabases(ctx context.Context, state resourceSecurityDlpFingerprintDatabasesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Server.IsNull() {
		result["server"] = data.Server.ValueString()
	}

	if !data.Sensitivity.IsNull() {
		result["sensitivity"] = data.Sensitivity.ValueString()
	}

	if !data.IncludeSubdirectories.IsNull() {
		result["includeSubdirectories"] = data.IncludeSubdirectories.ValueString()
	}

	if !data.ServerDirectory.IsNull() {
		result["serverDirectory"] = data.ServerDirectory.ValueString()
	}

	if !data.FilePattern.IsNull() {
		result["filePattern"] = data.FilePattern.ValueString()
	}

	if data.Schedule != nil {
		result["schedule"] = data.Schedule.expandSecurityDlpFingerprintDatabasesSchedule(ctx, diags)
	}

	if !data.RemoveDeletedFileFingerprints.IsNull() {
		result["removeDeletedFileFingerprints"] = data.RemoveDeletedFileFingerprints.ValueString()
	}

	if !data.KeepModified.IsNull() {
		result["keepModified"] = data.KeepModified.ValueString()
	}

	if !data.ScanOnCreation.IsNull() {
		result["scanOnCreation"] = data.ScanOnCreation.ValueString()
	}

	if data.Authentication != nil {
		result["authentication"] = data.Authentication.expandSecurityDlpFingerprintDatabasesAuthentication(ctx, diags)
	}

	return &result
}

func (data *resourceSecurityDlpFingerprintDatabasesModel) getURLObjectSecurityDlpFingerprintDatabases(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityDlpFingerprintDatabasesScheduleModel struct {
	Period            types.String  `tfsdk:"period"`
	SyncHour          types.Float64 `tfsdk:"sync_hour"`
	SyncMinute        types.Float64 `tfsdk:"sync_minute"`
	Weekday           types.String  `tfsdk:"weekday"`
	SyncDayOfTheMonth types.Float64 `tfsdk:"sync_day_of_the_month"`
}

type resourceSecurityDlpFingerprintDatabasesAuthenticationModel struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (m *resourceSecurityDlpFingerprintDatabasesScheduleModel) flattenSecurityDlpFingerprintDatabasesSchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpFingerprintDatabasesScheduleModel {
	if input == nil {
		return &resourceSecurityDlpFingerprintDatabasesScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpFingerprintDatabasesScheduleModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["period"]; ok {
		m.Period = parseStringValue(v)
	}

	if v, ok := o["syncHour"]; ok {
		m.SyncHour = parseFloat64Value(v)
	}

	if v, ok := o["syncMinute"]; ok {
		m.SyncMinute = parseFloat64Value(v)
	}

	if v, ok := o["weekday"]; ok {
		m.Weekday = parseStringValue(v)
	}

	if v, ok := o["syncDayOfTheMonth"]; ok {
		m.SyncDayOfTheMonth = parseFloat64Value(v)
	}

	return m
}

func (m *resourceSecurityDlpFingerprintDatabasesAuthenticationModel) flattenSecurityDlpFingerprintDatabasesAuthentication(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpFingerprintDatabasesAuthenticationModel {
	if input == nil {
		return &resourceSecurityDlpFingerprintDatabasesAuthenticationModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpFingerprintDatabasesAuthenticationModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	return m
}

func (data *resourceSecurityDlpFingerprintDatabasesScheduleModel) expandSecurityDlpFingerprintDatabasesSchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Period.IsNull() {
		result["period"] = data.Period.ValueString()
	}

	if !data.SyncHour.IsNull() {
		result["syncHour"] = data.SyncHour.ValueFloat64()
	}

	if !data.SyncMinute.IsNull() {
		result["syncMinute"] = data.SyncMinute.ValueFloat64()
	}

	if !data.Weekday.IsNull() {
		result["weekday"] = data.Weekday.ValueString()
	}

	if !data.SyncDayOfTheMonth.IsNull() {
		result["syncDayOfTheMonth"] = data.SyncDayOfTheMonth.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityDlpFingerprintDatabasesAuthenticationModel) expandSecurityDlpFingerprintDatabasesAuthentication(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Username.IsNull() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	return result
}
