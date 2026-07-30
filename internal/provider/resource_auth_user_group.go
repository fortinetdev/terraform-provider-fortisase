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
var _ resource.Resource = &resourceAuthUserGroup{}
var _ resource.ResourceWithMoveState = &resourceAuthUserGroup{}

func newResourceAuthUserGroup() resource.Resource {
	return &resourceAuthUserGroup{}
}

type resourceAuthUserGroup struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceAuthUserGroupModel describes the resource data model.
type resourceAuthUserGroupModel struct {
	ID               types.String                                 `tfsdk:"id"`
	PrimaryKey       types.String                                 `tfsdk:"primary_key"`
	GroupType        types.String                                 `tfsdk:"group_type"`
	LocalUsers       []resourceAuthUserGroupLocalUsersModel       `tfsdk:"local_users"`
	RemoteUserGroups []resourceAuthUserGroupRemoteUserGroupsModel `tfsdk:"remote_user_groups"`
	ScimUsers        []resourceAuthUserGroupScimUsersModel        `tfsdk:"scim_users"`
	ScimGroups       []resourceAuthUserGroupScimGroupsModel       `tfsdk:"scim_groups"`
	LocalUser        types.String                                 `tfsdk:"local_user"`
}

func (r *resourceAuthUserGroup) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_user_group"
}

func (r *resourceAuthUserGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "User Group Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthAtMost(35),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"group_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("fsso", "firewall", "scim"),
				},
				Computed: true,
				Optional: true,
			},
			"local_user": schema.StringAttribute{
				Optional: true,
			},
			"local_users": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"remote_user_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"matches": schema.SetAttribute{
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
						"server": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("auth/ldap-servers", "auth/radius-servers", "auth/swg-saml-server", "auth/vpn-saml-server", "auth/sslvpn-saml-server"),
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
			"scim_users": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
						},
					},
				},
				Optional: true,
			},
			"scim_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
						},
					},
				},
				Optional: true,
			},
		},
	}
}

func (r *resourceAuthUserGroup) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_user_group"
}
func (r *resourceAuthUserGroup) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_auth_user_groups" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceAuthUserGroupModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceAuthUserGroup) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthUserGroups")
	lock.Lock()
	defer lock.Unlock()
	var data resourceAuthUserGroupModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectAuthUserGroup(ctx, diags))
	input_model.URLParams = *(data.getURLObjectAuthUserGroup(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateAuthUserGroups(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectAuthUserGroup(ctx, "read", diags))

	read_output, err := c.ReadAuthUserGroups(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthUserGroup(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthUserGroup) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthUserGroups")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceAuthUserGroupModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceAuthUserGroupModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectAuthUserGroup(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectAuthUserGroup(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateAuthUserGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectAuthUserGroup(ctx, "read", diags))

	read_output, err := c.ReadAuthUserGroups(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthUserGroup(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthUserGroup) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("AuthUserGroups")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceAuthUserGroupModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthUserGroup(ctx, "delete", diags))

	output, err := c.DeleteAuthUserGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceAuthUserGroup) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthUserGroupModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthUserGroup(ctx, "read", diags))

	read_output, err := c.ReadAuthUserGroups(&input_model)
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

	diags.Append(data.refreshAuthUserGroup(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthUserGroup) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceAuthUserGroupModel) refreshAuthUserGroup(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["groupType"]; ok {
		m.GroupType = parseStringValue(v)
	}

	if v, ok := o["localUsers"]; ok {
		m.LocalUsers = m.flattenAuthUserGroupLocalUsersList(ctx, v, &diags)
	}

	if v, ok := o["remoteUserGroups"]; ok {
		m.RemoteUserGroups = m.flattenAuthUserGroupRemoteUserGroupsList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceAuthUserGroupModel) getCreateObjectAuthUserGroup(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.GroupType.IsNull() && !data.GroupType.IsUnknown() {
		result["groupType"] = data.GroupType.ValueString()
	}

	if data.LocalUsers != nil {
		result["localUsers"] = data.expandAuthUserGroupLocalUsersList(ctx, data.LocalUsers, diags)
	}

	if data.RemoteUserGroups != nil {
		result["remoteUserGroups"] = data.expandAuthUserGroupRemoteUserGroupsList(ctx, data.RemoteUserGroups, diags)
	}

	if data.ScimUsers != nil {
		result["scimUsers"] = data.expandAuthUserGroupScimUsersList(ctx, data.ScimUsers, diags)
	}

	if data.ScimGroups != nil {
		result["scimGroups"] = data.expandAuthUserGroupScimGroupsList(ctx, data.ScimGroups, diags)
	}

	if !data.LocalUser.IsNull() && !data.LocalUser.IsUnknown() {
		result["localUser"] = data.LocalUser.ValueString()
	}

	return &result
}

func (data *resourceAuthUserGroupModel) getUpdateObjectAuthUserGroup(ctx context.Context, state resourceAuthUserGroupModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.GroupType.IsNull() && !data.GroupType.IsUnknown() {
		result["groupType"] = data.GroupType.ValueString()
	}

	if data.LocalUsers != nil {
		result["localUsers"] = data.expandAuthUserGroupLocalUsersList(ctx, data.LocalUsers, diags)
	}

	if data.RemoteUserGroups != nil {
		result["remoteUserGroups"] = data.expandAuthUserGroupRemoteUserGroupsList(ctx, data.RemoteUserGroups, diags)
	}

	if data.ScimUsers != nil {
		result["scimUsers"] = data.expandAuthUserGroupScimUsersList(ctx, data.ScimUsers, diags)
	}

	if data.ScimGroups != nil {
		result["scimGroups"] = data.expandAuthUserGroupScimGroupsList(ctx, data.ScimGroups, diags)
	}

	if !data.LocalUser.IsNull() && !data.LocalUser.IsUnknown() {
		result["localUser"] = data.LocalUser.ValueString()
	}

	return &result
}

func (data *resourceAuthUserGroupModel) getURLObjectAuthUserGroup(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceAuthUserGroupLocalUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceAuthUserGroupRemoteUserGroupsModel struct {
	Server  *resourceAuthUserGroupRemoteUserGroupsServerModel `tfsdk:"server"`
	Matches types.Set                                         `tfsdk:"matches"`
}

type resourceAuthUserGroupRemoteUserGroupsServerModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceAuthUserGroupScimUsersModel struct {
	Name types.String `tfsdk:"name"`
}

type resourceAuthUserGroupScimGroupsModel struct {
	Name types.String `tfsdk:"name"`
}

func (m *resourceAuthUserGroupLocalUsersModel) flattenAuthUserGroupLocalUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthUserGroupLocalUsersModel {
	if input == nil {
		return &resourceAuthUserGroupLocalUsersModel{}
	}
	if m == nil {
		m = &resourceAuthUserGroupLocalUsersModel{}
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

func (s *resourceAuthUserGroupModel) flattenAuthUserGroupLocalUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceAuthUserGroupLocalUsersModel {
	if o == nil {
		return []resourceAuthUserGroupLocalUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument local_users is not type of []interface{}.", "")
		return []resourceAuthUserGroupLocalUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceAuthUserGroupLocalUsersModel{}
	}

	values := make([]resourceAuthUserGroupLocalUsersModel, len(l))
	for i, ele := range l {
		var m resourceAuthUserGroupLocalUsersModel
		if i < len(s.LocalUsers) {
			m = s.LocalUsers[i]
		}
		values[i] = *m.flattenAuthUserGroupLocalUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceAuthUserGroupRemoteUserGroupsModel) flattenAuthUserGroupRemoteUserGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthUserGroupRemoteUserGroupsModel {
	if input == nil {
		return &resourceAuthUserGroupRemoteUserGroupsModel{}
	}
	if m == nil {
		m = &resourceAuthUserGroupRemoteUserGroupsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["server"]; ok {
		m.Server = m.Server.flattenAuthUserGroupRemoteUserGroupsServer(ctx, v, diags)
	}

	if v, ok := o["matches"]; ok {
		m.Matches = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Matches = types.SetNull(types.StringType)
	}

	return m
}

func (s *resourceAuthUserGroupModel) flattenAuthUserGroupRemoteUserGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceAuthUserGroupRemoteUserGroupsModel {
	if o == nil {
		return []resourceAuthUserGroupRemoteUserGroupsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument remote_user_groups is not type of []interface{}.", "")
		return []resourceAuthUserGroupRemoteUserGroupsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceAuthUserGroupRemoteUserGroupsModel{}
	}

	values := make([]resourceAuthUserGroupRemoteUserGroupsModel, len(l))
	for i, ele := range l {
		var m resourceAuthUserGroupRemoteUserGroupsModel
		if i < len(s.RemoteUserGroups) {
			m = s.RemoteUserGroups[i]
		}
		values[i] = *m.flattenAuthUserGroupRemoteUserGroups(ctx, ele, diags)
	}

	return values
}

func (m *resourceAuthUserGroupRemoteUserGroupsServerModel) flattenAuthUserGroupRemoteUserGroupsServer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthUserGroupRemoteUserGroupsServerModel {
	if input == nil {
		return &resourceAuthUserGroupRemoteUserGroupsServerModel{}
	}
	if m == nil {
		m = &resourceAuthUserGroupRemoteUserGroupsServerModel{}
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

func (m *resourceAuthUserGroupScimUsersModel) flattenAuthUserGroupScimUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthUserGroupScimUsersModel {
	if input == nil {
		return &resourceAuthUserGroupScimUsersModel{}
	}
	if m == nil {
		m = &resourceAuthUserGroupScimUsersModel{}
	}

	return m
}

func (s *resourceAuthUserGroupModel) flattenAuthUserGroupScimUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceAuthUserGroupScimUsersModel {
	if o == nil {
		return []resourceAuthUserGroupScimUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument scim_users is not type of []interface{}.", "")
		return []resourceAuthUserGroupScimUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceAuthUserGroupScimUsersModel{}
	}

	values := make([]resourceAuthUserGroupScimUsersModel, len(l))
	for i, ele := range l {
		var m resourceAuthUserGroupScimUsersModel
		if i < len(s.ScimUsers) {
			m = s.ScimUsers[i]
		}
		values[i] = *m.flattenAuthUserGroupScimUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceAuthUserGroupScimGroupsModel) flattenAuthUserGroupScimGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthUserGroupScimGroupsModel {
	if input == nil {
		return &resourceAuthUserGroupScimGroupsModel{}
	}
	if m == nil {
		m = &resourceAuthUserGroupScimGroupsModel{}
	}

	return m
}

func (s *resourceAuthUserGroupModel) flattenAuthUserGroupScimGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceAuthUserGroupScimGroupsModel {
	if o == nil {
		return []resourceAuthUserGroupScimGroupsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument scim_groups is not type of []interface{}.", "")
		return []resourceAuthUserGroupScimGroupsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceAuthUserGroupScimGroupsModel{}
	}

	values := make([]resourceAuthUserGroupScimGroupsModel, len(l))
	for i, ele := range l {
		var m resourceAuthUserGroupScimGroupsModel
		if i < len(s.ScimGroups) {
			m = s.ScimGroups[i]
		}
		values[i] = *m.flattenAuthUserGroupScimGroups(ctx, ele, diags)
	}

	return values
}

func (data *resourceAuthUserGroupLocalUsersModel) expandAuthUserGroupLocalUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceAuthUserGroupModel) expandAuthUserGroupLocalUsersList(ctx context.Context, l []resourceAuthUserGroupLocalUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandAuthUserGroupLocalUsers(ctx, diags)
	}
	return result
}

func (data *resourceAuthUserGroupRemoteUserGroupsModel) expandAuthUserGroupRemoteUserGroups(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	result["server"] = nil
	if data.Server != nil && !isZeroStruct(*data.Server) {
		result["server"] = data.Server.expandAuthUserGroupRemoteUserGroupsServer(ctx, diags)
	}

	if !data.Matches.IsNull() && !data.Matches.IsUnknown() {
		result["matches"] = expandSetToStringList(data.Matches)
	}

	return result
}

func (s *resourceAuthUserGroupModel) expandAuthUserGroupRemoteUserGroupsList(ctx context.Context, l []resourceAuthUserGroupRemoteUserGroupsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandAuthUserGroupRemoteUserGroups(ctx, diags)
	}
	return result
}

func (data *resourceAuthUserGroupRemoteUserGroupsServerModel) expandAuthUserGroupRemoteUserGroupsServer(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceAuthUserGroupScimUsersModel) expandAuthUserGroupScimUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		result["name"] = data.Name.ValueString()
	}

	return result
}

func (s *resourceAuthUserGroupModel) expandAuthUserGroupScimUsersList(ctx context.Context, l []resourceAuthUserGroupScimUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandAuthUserGroupScimUsers(ctx, diags)
	}
	return result
}

func (data *resourceAuthUserGroupScimGroupsModel) expandAuthUserGroupScimGroups(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		result["name"] = data.Name.ValueString()
	}

	return result
}

func (s *resourceAuthUserGroupModel) expandAuthUserGroupScimGroupsList(ctx context.Context, l []resourceAuthUserGroupScimGroupsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandAuthUserGroupScimGroups(ctx, diags)
	}
	return result
}
