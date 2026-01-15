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
var _ resource.Resource = &resourceAuthUserGroups{}

func newResourceAuthUserGroups() resource.Resource {
	return &resourceAuthUserGroups{}
}

type resourceAuthUserGroups struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceAuthUserGroupsModel describes the resource data model.
type resourceAuthUserGroupsModel struct {
	ID               types.String                                  `tfsdk:"id"`
	PrimaryKey       types.String                                  `tfsdk:"primary_key"`
	GroupType        types.String                                  `tfsdk:"group_type"`
	LocalUsers       []resourceAuthUserGroupsLocalUsersModel       `tfsdk:"local_users"`
	RemoteUserGroups []resourceAuthUserGroupsRemoteUserGroupsModel `tfsdk:"remote_user_groups"`
}

func (r *resourceAuthUserGroups) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_user_groups"
}

func (r *resourceAuthUserGroups) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthAtMost(35),
				},
				Required: true,
			},
			"group_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("fsso", "firewall"),
				},
				Computed: true,
				Optional: true,
			},
			"local_users": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("auth/users"),
							},
							Computed: true,
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
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("auth/ldap-servers", "auth/radius-servers", "auth/swg-saml-server", "auth/vpn-saml-server"),
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

func (r *resourceAuthUserGroups) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_user_groups"
}

func (r *resourceAuthUserGroups) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthUserGroups")
	lock.Lock()
	defer lock.Unlock()
	var data resourceAuthUserGroupsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectAuthUserGroups(ctx, diags))
	input_model.URLParams = *(data.getURLObjectAuthUserGroups(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateAuthUserGroups(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectAuthUserGroups(ctx, "read", diags))

	read_output, err := c.ReadAuthUserGroups(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthUserGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthUserGroups) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthUserGroups")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceAuthUserGroupsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceAuthUserGroupsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectAuthUserGroups(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectAuthUserGroups(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectAuthUserGroups(ctx, "read", diags))

	read_output, err := c.ReadAuthUserGroups(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthUserGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthUserGroups) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("AuthUserGroups")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceAuthUserGroupsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthUserGroups(ctx, "delete", diags))

	output, err := c.DeleteAuthUserGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceAuthUserGroups) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthUserGroupsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthUserGroups(ctx, "read", diags))

	read_output, err := c.ReadAuthUserGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthUserGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthUserGroups) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceAuthUserGroupsModel) refreshAuthUserGroups(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["groupType"]; ok {
		m.GroupType = parseStringValue(v)
	}

	if v, ok := o["localUsers"]; ok {
		m.LocalUsers = m.flattenAuthUserGroupsLocalUsersList(ctx, v, &diags)
	}

	if v, ok := o["remoteUserGroups"]; ok {
		m.RemoteUserGroups = m.flattenAuthUserGroupsRemoteUserGroupsList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceAuthUserGroupsModel) getCreateObjectAuthUserGroups(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.GroupType.IsNull() {
		result["groupType"] = data.GroupType.ValueString()
	}

	result["localUsers"] = data.expandAuthUserGroupsLocalUsersList(ctx, data.LocalUsers, diags)

	result["remoteUserGroups"] = data.expandAuthUserGroupsRemoteUserGroupsList(ctx, data.RemoteUserGroups, diags)

	return &result
}

func (data *resourceAuthUserGroupsModel) getUpdateObjectAuthUserGroups(ctx context.Context, state resourceAuthUserGroupsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.GroupType.IsNull() {
		result["groupType"] = data.GroupType.ValueString()
	}

	if data.LocalUsers != nil {
		result["localUsers"] = data.expandAuthUserGroupsLocalUsersList(ctx, data.LocalUsers, diags)
	}

	if data.RemoteUserGroups != nil {
		result["remoteUserGroups"] = data.expandAuthUserGroupsRemoteUserGroupsList(ctx, data.RemoteUserGroups, diags)
	}

	return &result
}

func (data *resourceAuthUserGroupsModel) getURLObjectAuthUserGroups(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceAuthUserGroupsLocalUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceAuthUserGroupsRemoteUserGroupsModel struct {
	Server  *resourceAuthUserGroupsRemoteUserGroupsServerModel `tfsdk:"server"`
	Matches types.Set                                          `tfsdk:"matches"`
}

type resourceAuthUserGroupsRemoteUserGroupsServerModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceAuthUserGroupsLocalUsersModel) flattenAuthUserGroupsLocalUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthUserGroupsLocalUsersModel {
	if input == nil {
		return &resourceAuthUserGroupsLocalUsersModel{}
	}
	if m == nil {
		m = &resourceAuthUserGroupsLocalUsersModel{}
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

func (s *resourceAuthUserGroupsModel) flattenAuthUserGroupsLocalUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceAuthUserGroupsLocalUsersModel {
	if o == nil {
		return []resourceAuthUserGroupsLocalUsersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument local_users is not type of []interface{}.", "")
		return []resourceAuthUserGroupsLocalUsersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceAuthUserGroupsLocalUsersModel{}
	}

	values := make([]resourceAuthUserGroupsLocalUsersModel, len(l))
	for i, ele := range l {
		var m resourceAuthUserGroupsLocalUsersModel
		values[i] = *m.flattenAuthUserGroupsLocalUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceAuthUserGroupsRemoteUserGroupsModel) flattenAuthUserGroupsRemoteUserGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthUserGroupsRemoteUserGroupsModel {
	if input == nil {
		return &resourceAuthUserGroupsRemoteUserGroupsModel{}
	}
	if m == nil {
		m = &resourceAuthUserGroupsRemoteUserGroupsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["server"]; ok {
		m.Server = m.Server.flattenAuthUserGroupsRemoteUserGroupsServer(ctx, v, diags)
	}

	if v, ok := o["matches"]; ok {
		m.Matches = parseSetValue(ctx, v, types.StringType)
	}

	return m
}

func (s *resourceAuthUserGroupsModel) flattenAuthUserGroupsRemoteUserGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceAuthUserGroupsRemoteUserGroupsModel {
	if o == nil {
		return []resourceAuthUserGroupsRemoteUserGroupsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument remote_user_groups is not type of []interface{}.", "")
		return []resourceAuthUserGroupsRemoteUserGroupsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceAuthUserGroupsRemoteUserGroupsModel{}
	}

	values := make([]resourceAuthUserGroupsRemoteUserGroupsModel, len(l))
	for i, ele := range l {
		var m resourceAuthUserGroupsRemoteUserGroupsModel
		values[i] = *m.flattenAuthUserGroupsRemoteUserGroups(ctx, ele, diags)
	}

	return values
}

func (m *resourceAuthUserGroupsRemoteUserGroupsServerModel) flattenAuthUserGroupsRemoteUserGroupsServer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthUserGroupsRemoteUserGroupsServerModel {
	if input == nil {
		return &resourceAuthUserGroupsRemoteUserGroupsServerModel{}
	}
	if m == nil {
		m = &resourceAuthUserGroupsRemoteUserGroupsServerModel{}
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

func (data *resourceAuthUserGroupsLocalUsersModel) expandAuthUserGroupsLocalUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceAuthUserGroupsModel) expandAuthUserGroupsLocalUsersList(ctx context.Context, l []resourceAuthUserGroupsLocalUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandAuthUserGroupsLocalUsers(ctx, diags)
	}
	return result
}

func (data *resourceAuthUserGroupsRemoteUserGroupsModel) expandAuthUserGroupsRemoteUserGroups(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Server != nil && !isZeroStruct(*data.Server) {
		result["server"] = data.Server.expandAuthUserGroupsRemoteUserGroupsServer(ctx, diags)
	}

	if !data.Matches.IsNull() {
		result["matches"] = expandSetToStringList(data.Matches)
	}

	return result
}

func (s *resourceAuthUserGroupsModel) expandAuthUserGroupsRemoteUserGroupsList(ctx context.Context, l []resourceAuthUserGroupsRemoteUserGroupsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandAuthUserGroupsRemoteUserGroups(ctx, diags)
	}
	return result
}

func (data *resourceAuthUserGroupsRemoteUserGroupsServerModel) expandAuthUserGroupsRemoteUserGroupsServer(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
