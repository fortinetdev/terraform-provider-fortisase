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
var _ resource.Resource = &resourceInfraSsids{}

func newResourceInfraSsids() resource.Resource {
	return &resourceInfraSsids{}
}

type resourceInfraSsids struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceInfraSsidsModel describes the resource data model.
type resourceInfraSsidsModel struct {
	ID             types.String                            `tfsdk:"id"`
	PrimaryKey     types.String                            `tfsdk:"primary_key"`
	WifiSsid       types.String                            `tfsdk:"wifi_ssid"`
	BroadcastSsid  types.String                            `tfsdk:"broadcast_ssid"`
	ClientLimit    types.Float64                           `tfsdk:"client_limit"`
	SecurityMode   types.String                            `tfsdk:"security_mode"`
	CaptivePortal  types.Bool                              `tfsdk:"captive_portal"`
	SecurityGroups []resourceInfraSsidsSecurityGroupsModel `tfsdk:"security_groups"`
	PreSharedKey   types.String                            `tfsdk:"pre_shared_key"`
	RadiusServer   *resourceInfraSsidsRadiusServerModel    `tfsdk:"radius_server"`
	UserGroups     []resourceInfraSsidsUserGroupsModel     `tfsdk:"user_groups"`
}

func (r *resourceInfraSsids) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_ssids"
}

func (r *resourceInfraSsids) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 10),
				},
				Required: true,
			},
			"wifi_ssid": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"broadcast_ssid": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
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
					stringvalidator.OneOf("wpa2-only-personal", "wpa2-only-enterprise", "wpa3-only-enterprise", "wpa3-sae", "open", "wpa2-only-personal+captive-portal", "captive-portal"),
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
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("auth/user-groups"),
							},
							Computed: true,
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
						Computed: true,
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("auth/radius-servers"),
						},
						Computed: true,
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
								stringvalidator.OneOf("auth/user-groups"),
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

func (r *resourceInfraSsids) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_infra_ssids"
}

func (r *resourceInfraSsids) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("InfraSsids")
	lock.Lock()
	defer lock.Unlock()
	var data resourceInfraSsidsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectInfraSsids(ctx, diags))
	input_model.URLParams = *(data.getURLObjectInfraSsids(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateInfraSsids(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectInfraSsids(ctx, "read", diags))

	read_output, err := c.ReadInfraSsids(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraSsids(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraSsids) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("InfraSsids")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceInfraSsidsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceInfraSsidsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectInfraSsids(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectInfraSsids(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectInfraSsids(ctx, "read", diags))

	read_output, err := c.ReadInfraSsids(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraSsids(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraSsids) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("InfraSsids")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceInfraSsidsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectInfraSsids(ctx, "delete", diags))

	output, err := c.DeleteInfraSsids(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceInfraSsids) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceInfraSsidsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectInfraSsids(ctx, "read", diags))

	read_output, err := c.ReadInfraSsids(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraSsids(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraSsids) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceInfraSsidsModel) refreshInfraSsids(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.SecurityGroups = m.flattenInfraSsidsSecurityGroupsList(ctx, v, &diags)
	}

	if v, ok := o["preSharedKey"]; ok {
		m.PreSharedKey = parseStringValue(v)
	}

	if v, ok := o["radiusServer"]; ok {
		m.RadiusServer = m.RadiusServer.flattenInfraSsidsRadiusServer(ctx, v, &diags)
	}

	if v, ok := o["userGroups"]; ok {
		m.UserGroups = m.flattenInfraSsidsUserGroupsList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceInfraSsidsModel) getCreateObjectInfraSsids(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.WifiSsid.IsNull() {
		result["wifiSsid"] = data.WifiSsid.ValueString()
	}

	if !data.BroadcastSsid.IsNull() {
		result["broadcastSsid"] = data.BroadcastSsid.ValueString()
	}

	if !data.ClientLimit.IsNull() {
		result["clientLimit"] = data.ClientLimit.ValueFloat64()
	}

	if !data.SecurityMode.IsNull() {
		result["securityMode"] = data.SecurityMode.ValueString()
	}

	if !data.CaptivePortal.IsNull() {
		result["captivePortal"] = data.CaptivePortal.ValueBool()
	}

	if data.SecurityGroups != nil {
		result["securityGroups"] = data.expandInfraSsidsSecurityGroupsList(ctx, data.SecurityGroups, diags)
	}

	if !data.PreSharedKey.IsNull() {
		result["preSharedKey"] = data.PreSharedKey.ValueString()
	}

	if data.RadiusServer != nil && !isZeroStruct(*data.RadiusServer) {
		result["radiusServer"] = data.RadiusServer.expandInfraSsidsRadiusServer(ctx, diags)
	}

	if data.UserGroups != nil {
		result["userGroups"] = data.expandInfraSsidsUserGroupsList(ctx, data.UserGroups, diags)
	}

	return &result
}

func (data *resourceInfraSsidsModel) getUpdateObjectInfraSsids(ctx context.Context, state resourceInfraSsidsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.WifiSsid.IsNull() {
		result["wifiSsid"] = data.WifiSsid.ValueString()
	}

	if !data.BroadcastSsid.IsNull() {
		result["broadcastSsid"] = data.BroadcastSsid.ValueString()
	}

	if !data.ClientLimit.IsNull() {
		result["clientLimit"] = data.ClientLimit.ValueFloat64()
	}

	if !data.SecurityMode.IsNull() {
		result["securityMode"] = data.SecurityMode.ValueString()
	}

	if !data.CaptivePortal.IsNull() {
		result["captivePortal"] = data.CaptivePortal.ValueBool()
	}

	if data.SecurityGroups != nil {
		result["securityGroups"] = data.expandInfraSsidsSecurityGroupsList(ctx, data.SecurityGroups, diags)
	}

	if !data.PreSharedKey.IsNull() {
		result["preSharedKey"] = data.PreSharedKey.ValueString()
	}

	if data.RadiusServer != nil {
		result["radiusServer"] = data.RadiusServer.expandInfraSsidsRadiusServer(ctx, diags)
	}

	if data.UserGroups != nil {
		result["userGroups"] = data.expandInfraSsidsUserGroupsList(ctx, data.UserGroups, diags)
	}

	return &result
}

func (data *resourceInfraSsidsModel) getURLObjectInfraSsids(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceInfraSsidsSecurityGroupsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceInfraSsidsRadiusServerModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceInfraSsidsUserGroupsModel struct {
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceInfraSsidsSecurityGroupsModel) flattenInfraSsidsSecurityGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceInfraSsidsSecurityGroupsModel {
	if input == nil {
		return &resourceInfraSsidsSecurityGroupsModel{}
	}
	if m == nil {
		m = &resourceInfraSsidsSecurityGroupsModel{}
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

func (s *resourceInfraSsidsModel) flattenInfraSsidsSecurityGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceInfraSsidsSecurityGroupsModel {
	if o == nil {
		return []resourceInfraSsidsSecurityGroupsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument security_groups is not type of []interface{}.", "")
		return []resourceInfraSsidsSecurityGroupsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceInfraSsidsSecurityGroupsModel{}
	}

	values := make([]resourceInfraSsidsSecurityGroupsModel, len(l))
	for i, ele := range l {
		var m resourceInfraSsidsSecurityGroupsModel
		values[i] = *m.flattenInfraSsidsSecurityGroups(ctx, ele, diags)
	}

	return values
}

func (m *resourceInfraSsidsRadiusServerModel) flattenInfraSsidsRadiusServer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceInfraSsidsRadiusServerModel {
	if input == nil {
		return &resourceInfraSsidsRadiusServerModel{}
	}
	if m == nil {
		m = &resourceInfraSsidsRadiusServerModel{}
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

func (m *resourceInfraSsidsUserGroupsModel) flattenInfraSsidsUserGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceInfraSsidsUserGroupsModel {
	if input == nil {
		return &resourceInfraSsidsUserGroupsModel{}
	}
	if m == nil {
		m = &resourceInfraSsidsUserGroupsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *resourceInfraSsidsModel) flattenInfraSsidsUserGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceInfraSsidsUserGroupsModel {
	if o == nil {
		return []resourceInfraSsidsUserGroupsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument user_groups is not type of []interface{}.", "")
		return []resourceInfraSsidsUserGroupsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceInfraSsidsUserGroupsModel{}
	}

	values := make([]resourceInfraSsidsUserGroupsModel, len(l))
	for i, ele := range l {
		var m resourceInfraSsidsUserGroupsModel
		values[i] = *m.flattenInfraSsidsUserGroups(ctx, ele, diags)
	}

	return values
}

func (data *resourceInfraSsidsSecurityGroupsModel) expandInfraSsidsSecurityGroups(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceInfraSsidsModel) expandInfraSsidsSecurityGroupsList(ctx context.Context, l []resourceInfraSsidsSecurityGroupsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandInfraSsidsSecurityGroups(ctx, diags)
	}
	return result
}

func (data *resourceInfraSsidsRadiusServerModel) expandInfraSsidsRadiusServer(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceInfraSsidsUserGroupsModel) expandInfraSsidsUserGroups(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceInfraSsidsModel) expandInfraSsidsUserGroupsList(ctx context.Context, l []resourceInfraSsidsUserGroupsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandInfraSsidsUserGroups(ctx, diags)
	}
	return result
}
