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
var _ resource.Resource = &resourceInfraSsid{}
var _ resource.ResourceWithMoveState = &resourceInfraSsid{}

func newResourceInfraSsid() resource.Resource {
	return &resourceInfraSsid{}
}

type resourceInfraSsid struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceInfraSsidModel describes the resource data model.
type resourceInfraSsidModel struct {
	ID             types.String                           `tfsdk:"id"`
	PrimaryKey     types.String                           `tfsdk:"primary_key"`
	WifiSsid       types.String                           `tfsdk:"wifi_ssid"`
	BroadcastSsid  types.String                           `tfsdk:"broadcast_ssid"`
	ClientLimit    types.Float64                          `tfsdk:"client_limit"`
	SecurityMode   types.String                           `tfsdk:"security_mode"`
	CaptivePortal  types.Bool                             `tfsdk:"captive_portal"`
	SecurityGroups []resourceInfraSsidSecurityGroupsModel `tfsdk:"security_groups"`
	PreSharedKey   types.String                           `tfsdk:"pre_shared_key"`
	RadiusServer   *resourceInfraSsidRadiusServerModel    `tfsdk:"radius_server"`
	UserGroups     []resourceInfraSsidUserGroupsModel     `tfsdk:"user_groups"`
}

func (r *resourceInfraSsid) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_ssid"
}

func (r *resourceInfraSsid) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "FortiAP SSID Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 15),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"wifi_ssid": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 32),
				},
				Computed: true,
				Optional: true,
			},
			"broadcast_ssid": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"client_limit": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"security_mode": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("wpa2-only-personal", "wpa2-only-enterprise", "wpa3-only-enterprise", "wpa3-sae", "open", "wpa2-only-personal+captive-portal", "captive-portal"),
				},
				Computed: true,
				Optional: true,
			},
			"captive_portal": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"pre_shared_key": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("auth/user-groups"),
							},
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"radius_server": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("auth/radius-servers"),
						},
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"user_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("auth/user-groups"),
							},
							Optional: true,
						},
						"primary_key": schema.StringAttribute{
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

func (r *resourceInfraSsid) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_infra_ssid"
}
func (r *resourceInfraSsid) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_infra_ssids" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceInfraSsidModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceInfraSsid) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("InfraSsids")
	lock.Lock()
	defer lock.Unlock()
	var data resourceInfraSsidModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectInfraSsid(ctx, diags))
	input_model.URLParams = *(data.getURLObjectInfraSsid(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateInfraSsids(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectInfraSsid(ctx, "read", diags))

	read_output, err := c.ReadInfraSsids(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraSsid(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraSsid) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("InfraSsids")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceInfraSsidModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceInfraSsidModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectInfraSsid(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectInfraSsid(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateInfraSsids(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectInfraSsid(ctx, "read", diags))

	read_output, err := c.ReadInfraSsids(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraSsid(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraSsid) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("InfraSsids")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceInfraSsidModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectInfraSsid(ctx, "delete", diags))

	output, err := c.DeleteInfraSsids(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceInfraSsid) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceInfraSsidModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectInfraSsid(ctx, "read", diags))

	read_output, err := c.ReadInfraSsids(&input_model)
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

	diags.Append(data.refreshInfraSsid(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraSsid) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceInfraSsidModel) refreshInfraSsid(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["wifiSsid"]; ok {
		m.WifiSsid = parseStringValue(v)
	}

	if v, ok := o["broadcastSsid"]; ok {
		m.BroadcastSsid = parseStringValue(v)
	}

	if v, ok := o["clientLimit"]; ok {
		m.ClientLimit = parseFloat64Value(v)
	}

	if v, ok := o["securityMode"]; ok {
		m.SecurityMode = parseStringValue(v)
	}

	if v, ok := o["captivePortal"]; ok {
		m.CaptivePortal = parseBoolValue(v)
	}

	if v, ok := o["securityGroups"]; ok {
		m.SecurityGroups = m.flattenInfraSsidSecurityGroupsList(ctx, v, &diags)
	}

	if v, ok := o["radiusServer"]; ok {
		m.RadiusServer = m.RadiusServer.flattenInfraSsidRadiusServer(ctx, v, &diags)
	}

	if v, ok := o["userGroups"]; ok {
		m.UserGroups = m.flattenInfraSsidUserGroupsList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceInfraSsidModel) getCreateObjectInfraSsid(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.WifiSsid.IsNull() && !data.WifiSsid.IsUnknown() {
		result["wifiSsid"] = data.WifiSsid.ValueString()
	}

	if !data.BroadcastSsid.IsNull() && !data.BroadcastSsid.IsUnknown() {
		result["broadcastSsid"] = data.BroadcastSsid.ValueString()
	}

	if !data.ClientLimit.IsNull() && !data.ClientLimit.IsUnknown() {
		result["clientLimit"] = data.ClientLimit.ValueFloat64()
	}

	if !data.SecurityMode.IsNull() && !data.SecurityMode.IsUnknown() {
		result["securityMode"] = data.SecurityMode.ValueString()
	}

	if !data.CaptivePortal.IsNull() && !data.CaptivePortal.IsUnknown() {
		result["captivePortal"] = data.CaptivePortal.ValueBool()
	}

	if data.SecurityGroups != nil {
		result["securityGroups"] = data.expandInfraSsidSecurityGroupsList(ctx, data.SecurityGroups, diags)
	}

	if !data.PreSharedKey.IsNull() && !data.PreSharedKey.IsUnknown() {
		result["preSharedKey"] = data.PreSharedKey.ValueString()
	}

	result["radiusServer"] = nil
	if data.RadiusServer != nil && !isZeroStruct(*data.RadiusServer) {
		result["radiusServer"] = data.RadiusServer.expandInfraSsidRadiusServer(ctx, diags)
	}

	if data.UserGroups != nil {
		result["userGroups"] = data.expandInfraSsidUserGroupsList(ctx, data.UserGroups, diags)
	}

	return &result
}

func (data *resourceInfraSsidModel) getUpdateObjectInfraSsid(ctx context.Context, state resourceInfraSsidModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.WifiSsid.IsNull() && !data.WifiSsid.IsUnknown() {
		result["wifiSsid"] = data.WifiSsid.ValueString()
	}

	if !data.BroadcastSsid.IsNull() && !data.BroadcastSsid.IsUnknown() {
		result["broadcastSsid"] = data.BroadcastSsid.ValueString()
	}

	if !data.ClientLimit.IsNull() && !data.ClientLimit.IsUnknown() {
		result["clientLimit"] = data.ClientLimit.ValueFloat64()
	}

	if !data.SecurityMode.IsNull() && !data.SecurityMode.IsUnknown() {
		result["securityMode"] = data.SecurityMode.ValueString()
	}

	if !data.CaptivePortal.IsNull() && !data.CaptivePortal.IsUnknown() {
		result["captivePortal"] = data.CaptivePortal.ValueBool()
	}

	if data.SecurityGroups != nil {
		result["securityGroups"] = data.expandInfraSsidSecurityGroupsList(ctx, data.SecurityGroups, diags)
	}

	if !data.PreSharedKey.IsNull() && !data.PreSharedKey.IsUnknown() {
		result["preSharedKey"] = data.PreSharedKey.ValueString()
	}

	result["radiusServer"] = nil
	if data.RadiusServer != nil && !isZeroStruct(*data.RadiusServer) {
		result["radiusServer"] = data.RadiusServer.expandInfraSsidRadiusServer(ctx, diags)
	}

	if data.UserGroups != nil {
		result["userGroups"] = data.expandInfraSsidUserGroupsList(ctx, data.UserGroups, diags)
	}

	return &result
}

func (data *resourceInfraSsidModel) getURLObjectInfraSsid(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceInfraSsidSecurityGroupsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceInfraSsidRadiusServerModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceInfraSsidUserGroupsModel struct {
	Datasource types.String `tfsdk:"datasource"`
	PrimaryKey types.String `tfsdk:"primary_key"`
}

func (m *resourceInfraSsidSecurityGroupsModel) flattenInfraSsidSecurityGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceInfraSsidSecurityGroupsModel {
	if input == nil {
		return &resourceInfraSsidSecurityGroupsModel{}
	}
	if m == nil {
		m = &resourceInfraSsidSecurityGroupsModel{}
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

func (s *resourceInfraSsidModel) flattenInfraSsidSecurityGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceInfraSsidSecurityGroupsModel {
	if o == nil {
		return []resourceInfraSsidSecurityGroupsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument security_groups is not type of []interface{}.", "")
		return []resourceInfraSsidSecurityGroupsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceInfraSsidSecurityGroupsModel{}
	}

	values := make([]resourceInfraSsidSecurityGroupsModel, len(l))
	for i, ele := range l {
		var m resourceInfraSsidSecurityGroupsModel
		if i < len(s.SecurityGroups) {
			m = s.SecurityGroups[i]
		}
		values[i] = *m.flattenInfraSsidSecurityGroups(ctx, ele, diags)
	}

	return values
}

func (m *resourceInfraSsidRadiusServerModel) flattenInfraSsidRadiusServer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceInfraSsidRadiusServerModel {
	if input == nil {
		return &resourceInfraSsidRadiusServerModel{}
	}
	if m == nil {
		m = &resourceInfraSsidRadiusServerModel{}
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

func (m *resourceInfraSsidUserGroupsModel) flattenInfraSsidUserGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceInfraSsidUserGroupsModel {
	if input == nil {
		return &resourceInfraSsidUserGroupsModel{}
	}
	if m == nil {
		m = &resourceInfraSsidUserGroupsModel{}
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

func (s *resourceInfraSsidModel) flattenInfraSsidUserGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceInfraSsidUserGroupsModel {
	if o == nil {
		return []resourceInfraSsidUserGroupsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument user_groups is not type of []interface{}.", "")
		return []resourceInfraSsidUserGroupsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceInfraSsidUserGroupsModel{}
	}

	values := make([]resourceInfraSsidUserGroupsModel, len(l))
	for i, ele := range l {
		var m resourceInfraSsidUserGroupsModel
		if i < len(s.UserGroups) {
			m = s.UserGroups[i]
		}
		values[i] = *m.flattenInfraSsidUserGroups(ctx, ele, diags)
	}

	return values
}

func (data *resourceInfraSsidSecurityGroupsModel) expandInfraSsidSecurityGroups(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceInfraSsidModel) expandInfraSsidSecurityGroupsList(ctx context.Context, l []resourceInfraSsidSecurityGroupsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandInfraSsidSecurityGroups(ctx, diags)
	}
	return result
}

func (data *resourceInfraSsidRadiusServerModel) expandInfraSsidRadiusServer(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceInfraSsidUserGroupsModel) expandInfraSsidUserGroups(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return result
}

func (s *resourceInfraSsidModel) expandInfraSsidUserGroupsList(ctx context.Context, l []resourceInfraSsidUserGroupsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandInfraSsidUserGroups(ctx, diags)
	}
	return result
}
