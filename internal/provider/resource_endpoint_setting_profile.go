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
var _ resource.Resource = &resourceEndpointSettingProfile{}
var _ resource.ResourceWithMoveState = &resourceEndpointSettingProfile{}

func newResourceEndpointSettingProfile() resource.Resource {
	return &resourceEndpointSettingProfile{}
}

type resourceEndpointSettingProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointSettingProfileModel describes the resource data model.
type resourceEndpointSettingProfileModel struct {
	ID                              types.String                               `tfsdk:"id"`
	AllowConfigBackup               types.String                               `tfsdk:"allow_config_backup"`
	AllowDebugLogGeneration         types.String                               `tfsdk:"allow_debug_log_generation"`
	ShowTagFortiClient              types.String                               `tfsdk:"show_tag_forti_client"`
	ShowNotifications               types.String                               `tfsdk:"show_notifications"`
	NotifyVpnIssue                  types.String                               `tfsdk:"notify_vpn_issue"`
	UsersCanDisconnect              types.String                               `tfsdk:"users_can_disconnect"`
	TriggerVulnScanOnSoftwareChange types.String                               `tfsdk:"trigger_vuln_scan_on_software_change"`
	FctGui                          *resourceEndpointSettingProfileFctGuiModel `tfsdk:"fct_gui"`
	EmsDisconnectPassword           types.String                               `tfsdk:"ems_disconnect_password"`
	PrimaryKey                      types.String                               `tfsdk:"primary_key"`
}

func (r *resourceEndpointSettingProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_setting_profile"
}

func (r *resourceEndpointSettingProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Settings Profile Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"allow_config_backup": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"allow_debug_log_generation": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"show_tag_forti_client": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"show_notifications": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"notify_vpn_issue": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"users_can_disconnect": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"trigger_vuln_scan_on_software_change": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"ems_disconnect_password": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"fct_gui": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"default_tab": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceEndpointSettingProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_setting_profile"
}
func (r *resourceEndpointSettingProfile) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_endpoint_setting_profiles" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceEndpointSettingProfileModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceEndpointSettingProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointSettingProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointSettingProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointSettingProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateEndpointSettingProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := data.PrimaryKey.ValueString()
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointSettingProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointSettingProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSettingProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSettingProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointSettingProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointSettingProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointSettingProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointSettingProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointSettingProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointSettingProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointSettingProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSettingProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSettingProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointSettingProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointSettingProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointSettingProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointSettingProfiles(&input_model)
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

	diags.Append(data.refreshEndpointSettingProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSettingProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceEndpointSettingProfileModel) refreshEndpointSettingProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["allowConfigBackup"]; ok {
		m.AllowConfigBackup = parseStringValue(v)
	}

	if v, ok := o["allowDebugLogGeneration"]; ok {
		m.AllowDebugLogGeneration = parseStringValue(v)
	}

	if v, ok := o["showTagFortiClient"]; ok {
		m.ShowTagFortiClient = parseStringValue(v)
	}

	if v, ok := o["showNotifications"]; ok {
		m.ShowNotifications = parseStringValue(v)
	}

	if v, ok := o["notifyVpnIssue"]; ok {
		m.NotifyVpnIssue = parseStringValue(v)
	}

	if v, ok := o["usersCanDisconnect"]; ok {
		m.UsersCanDisconnect = parseStringValue(v)
	}

	if v, ok := o["triggerVulnScanOnSoftwareChange"]; ok {
		m.TriggerVulnScanOnSoftwareChange = parseStringValue(v)
	}

	if v, ok := o["fctGui"]; ok {
		m.FctGui = m.FctGui.flattenEndpointSettingProfileFctGui(ctx, v, &diags)
	}

	if v, ok := o["emsDisconnectPassword"]; ok {
		m.EmsDisconnectPassword = parseStringValue(v)
	}

	return diags
}

func (data *resourceEndpointSettingProfileModel) getCreateObjectEndpointSettingProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AllowConfigBackup.IsNull() && !data.AllowConfigBackup.IsUnknown() {
		result["allowConfigBackup"] = data.AllowConfigBackup.ValueString()
	}

	if !data.AllowDebugLogGeneration.IsNull() && !data.AllowDebugLogGeneration.IsUnknown() {
		result["allowDebugLogGeneration"] = data.AllowDebugLogGeneration.ValueString()
	}

	if !data.ShowTagFortiClient.IsNull() && !data.ShowTagFortiClient.IsUnknown() {
		result["showTagFortiClient"] = data.ShowTagFortiClient.ValueString()
	}

	if !data.ShowNotifications.IsNull() && !data.ShowNotifications.IsUnknown() {
		result["showNotifications"] = data.ShowNotifications.ValueString()
	}

	if !data.NotifyVpnIssue.IsNull() && !data.NotifyVpnIssue.IsUnknown() {
		result["notifyVpnIssue"] = data.NotifyVpnIssue.ValueString()
	}

	if !data.UsersCanDisconnect.IsNull() && !data.UsersCanDisconnect.IsUnknown() {
		result["usersCanDisconnect"] = data.UsersCanDisconnect.ValueString()
	}

	if !data.TriggerVulnScanOnSoftwareChange.IsNull() && !data.TriggerVulnScanOnSoftwareChange.IsUnknown() {
		result["triggerVulnScanOnSoftwareChange"] = data.TriggerVulnScanOnSoftwareChange.ValueString()
	}

	if data.FctGui != nil && !isZeroStruct(*data.FctGui) {
		result["fctGui"] = data.FctGui.expandEndpointSettingProfileFctGui(ctx, diags)
	}

	if !data.EmsDisconnectPassword.IsNull() && !data.EmsDisconnectPassword.IsUnknown() {
		result["emsDisconnectPassword"] = data.EmsDisconnectPassword.ValueString()
	}

	return &result
}

func (data *resourceEndpointSettingProfileModel) getUpdateObjectEndpointSettingProfile(ctx context.Context, state resourceEndpointSettingProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AllowConfigBackup.IsNull() && !data.AllowConfigBackup.IsUnknown() {
		result["allowConfigBackup"] = data.AllowConfigBackup.ValueString()
	}

	if !data.AllowDebugLogGeneration.IsNull() && !data.AllowDebugLogGeneration.IsUnknown() {
		result["allowDebugLogGeneration"] = data.AllowDebugLogGeneration.ValueString()
	}

	if !data.ShowTagFortiClient.IsNull() && !data.ShowTagFortiClient.IsUnknown() {
		result["showTagFortiClient"] = data.ShowTagFortiClient.ValueString()
	}

	if !data.ShowNotifications.IsNull() && !data.ShowNotifications.IsUnknown() {
		result["showNotifications"] = data.ShowNotifications.ValueString()
	}

	if !data.NotifyVpnIssue.IsNull() && !data.NotifyVpnIssue.IsUnknown() {
		result["notifyVpnIssue"] = data.NotifyVpnIssue.ValueString()
	}

	if !data.UsersCanDisconnect.IsNull() && !data.UsersCanDisconnect.IsUnknown() {
		result["usersCanDisconnect"] = data.UsersCanDisconnect.ValueString()
	}

	if !data.TriggerVulnScanOnSoftwareChange.IsNull() && !data.TriggerVulnScanOnSoftwareChange.IsUnknown() {
		result["triggerVulnScanOnSoftwareChange"] = data.TriggerVulnScanOnSoftwareChange.ValueString()
	}

	if data.FctGui != nil {
		result["fctGui"] = data.FctGui.expandEndpointSettingProfileFctGui(ctx, diags)
	}

	if !data.EmsDisconnectPassword.IsNull() && !data.EmsDisconnectPassword.IsUnknown() {
		result["emsDisconnectPassword"] = data.EmsDisconnectPassword.ValueString()
	}

	return &result
}

func (data *resourceEndpointSettingProfileModel) getURLObjectEndpointSettingProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointSettingProfileFctGuiModel struct {
	DefaultTab types.String `tfsdk:"default_tab"`
}

func (m *resourceEndpointSettingProfileFctGuiModel) flattenEndpointSettingProfileFctGui(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointSettingProfileFctGuiModel {
	if input == nil {
		return &resourceEndpointSettingProfileFctGuiModel{}
	}
	if m == nil {
		m = &resourceEndpointSettingProfileFctGuiModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["defaultTab"]; ok {
		m.DefaultTab = parseStringValue(v)
	}

	return m
}

func (data *resourceEndpointSettingProfileFctGuiModel) expandEndpointSettingProfileFctGui(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.DefaultTab.IsNull() && !data.DefaultTab.IsUnknown() {
		result["defaultTab"] = data.DefaultTab.ValueString()
	}

	return result
}
