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
var _ resource.Resource = &resourceEndpointSettingProfiles{}

func newResourceEndpointSettingProfiles() resource.Resource {
	return &resourceEndpointSettingProfiles{}
}

type resourceEndpointSettingProfiles struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointSettingProfilesModel describes the resource data model.
type resourceEndpointSettingProfilesModel struct {
	ID                    types.String `tfsdk:"id"`
	AllowConfigBackup     types.String `tfsdk:"allow_config_backup"`
	ShowTagFortiClient    types.String `tfsdk:"show_tag_forti_client"`
	ShowNotifications     types.String `tfsdk:"show_notifications"`
	NotifyVpnIssue        types.String `tfsdk:"notify_vpn_issue"`
	UsersCanDisconnect    types.String `tfsdk:"users_can_disconnect"`
	EmsDisconnectPassword types.String `tfsdk:"ems_disconnect_password"`
	PrimaryKey            types.String `tfsdk:"primary_key"`
}

func (r *resourceEndpointSettingProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_setting_profiles"
}

func (r *resourceEndpointSettingProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"show_tag_forti_client": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"show_notifications": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"notify_vpn_issue": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"users_can_disconnect": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
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
		},
	}
}

func (r *resourceEndpointSettingProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_setting_profiles"
}

func (r *resourceEndpointSettingProfiles) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointSettingProfiles")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointSettingProfilesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointSettingProfiles(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointSettingProfiles(ctx, "create", diags))

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

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointSettingProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointSettingProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSettingProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSettingProfiles) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointSettingProfiles")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointSettingProfilesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointSettingProfilesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointSettingProfiles(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointSettingProfiles(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectEndpointSettingProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointSettingProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSettingProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSettingProfiles) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointSettingProfiles) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointSettingProfilesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointSettingProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointSettingProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSettingProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSettingProfiles) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointSettingProfilesModel) refreshEndpointSettingProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["allowConfigBackup"]; ok {
		m.AllowConfigBackup = parseStringValue(v)
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

	if v, ok := o["emsDisconnectPassword"]; ok {
		m.EmsDisconnectPassword = parseStringValue(v)
	}

	return diags
}

func (data *resourceEndpointSettingProfilesModel) getCreateObjectEndpointSettingProfiles(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AllowConfigBackup.IsNull() {
		result["allowConfigBackup"] = data.AllowConfigBackup.ValueString()
	}

	if !data.ShowTagFortiClient.IsNull() {
		result["showTagFortiClient"] = data.ShowTagFortiClient.ValueString()
	}

	if !data.ShowNotifications.IsNull() {
		result["showNotifications"] = data.ShowNotifications.ValueString()
	}

	if !data.NotifyVpnIssue.IsNull() {
		result["notifyVpnIssue"] = data.NotifyVpnIssue.ValueString()
	}

	if !data.UsersCanDisconnect.IsNull() {
		result["usersCanDisconnect"] = data.UsersCanDisconnect.ValueString()
	}

	if !data.EmsDisconnectPassword.IsNull() {
		result["emsDisconnectPassword"] = data.EmsDisconnectPassword.ValueString()
	}

	return &result
}

func (data *resourceEndpointSettingProfilesModel) getUpdateObjectEndpointSettingProfiles(ctx context.Context, state resourceEndpointSettingProfilesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AllowConfigBackup.IsNull() {
		result["allowConfigBackup"] = data.AllowConfigBackup.ValueString()
	}

	if !data.ShowTagFortiClient.IsNull() {
		result["showTagFortiClient"] = data.ShowTagFortiClient.ValueString()
	}

	if !data.ShowNotifications.IsNull() {
		result["showNotifications"] = data.ShowNotifications.ValueString()
	}

	if !data.NotifyVpnIssue.IsNull() {
		result["notifyVpnIssue"] = data.NotifyVpnIssue.ValueString()
	}

	if !data.UsersCanDisconnect.IsNull() {
		result["usersCanDisconnect"] = data.UsersCanDisconnect.ValueString()
	}

	if !data.EmsDisconnectPassword.IsNull() {
		result["emsDisconnectPassword"] = data.EmsDisconnectPassword.ValueString()
	}

	return &result
}

func (data *resourceEndpointSettingProfilesModel) getURLObjectEndpointSettingProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
