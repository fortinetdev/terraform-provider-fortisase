// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceAuthUserGroup{}

func newDatasourceAuthUserGroup() datasource.DataSource {
	return &datasourceAuthUserGroup{}
}

type datasourceAuthUserGroup struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceAuthUserGroupModel describes the datasource data model.
type datasourceAuthUserGroupModel struct {
	PrimaryKey       types.String                                   `tfsdk:"primary_key"`
	GroupType        types.String                                   `tfsdk:"group_type"`
	LocalUsers       []datasourceAuthUserGroupLocalUsersModel       `tfsdk:"local_users"`
	RemoteUserGroups []datasourceAuthUserGroupRemoteUserGroupsModel `tfsdk:"remote_user_groups"`
}

func (r *datasourceAuthUserGroup) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_user_group"
}

func (r *datasourceAuthUserGroup) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "User Group Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(35),
				},
				Required: true,
			},
			"group_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("fsso", "firewall", "scim"),
				},
				Computed: true,
			},
			"local_users": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"remote_user_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"matches": schema.SetAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"server": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("auth/ldap-servers", "auth/radius-servers", "auth/swg-saml-server", "auth/vpn-saml-server", "auth/sslvpn-saml-server"),
									},
									Computed: true,
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceAuthUserGroup) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceAuthUserGroup) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceAuthUserGroupModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthUserGroup(ctx, "read", diags))

	read_output, err := c.ReadAuthUserGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
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

func (m *datasourceAuthUserGroupModel) refreshAuthUserGroup(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *datasourceAuthUserGroupModel) getURLObjectAuthUserGroup(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceAuthUserGroupLocalUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceAuthUserGroupRemoteUserGroupsModel struct {
	Server  *datasourceAuthUserGroupRemoteUserGroupsServerModel `tfsdk:"server"`
	Matches types.Set                                           `tfsdk:"matches"`
}

type datasourceAuthUserGroupRemoteUserGroupsServerModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceAuthUserGroupLocalUsersModel) flattenAuthUserGroupLocalUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthUserGroupLocalUsersModel {
	if input == nil {
		return &datasourceAuthUserGroupLocalUsersModel{}
	}
	if m == nil {
		m = &datasourceAuthUserGroupLocalUsersModel{}
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

func (s *datasourceAuthUserGroupModel) flattenAuthUserGroupLocalUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceAuthUserGroupLocalUsersModel {
	if o == nil {
		return []datasourceAuthUserGroupLocalUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument local_users is not type of []interface{}.", "")
		return []datasourceAuthUserGroupLocalUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceAuthUserGroupLocalUsersModel{}
	}

	values := make([]datasourceAuthUserGroupLocalUsersModel, len(l))
	for i, ele := range l {
		var m datasourceAuthUserGroupLocalUsersModel
		if i < len(s.LocalUsers) {
			m = s.LocalUsers[i]
		}
		values[i] = *m.flattenAuthUserGroupLocalUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceAuthUserGroupRemoteUserGroupsModel) flattenAuthUserGroupRemoteUserGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthUserGroupRemoteUserGroupsModel {
	if input == nil {
		return &datasourceAuthUserGroupRemoteUserGroupsModel{}
	}
	if m == nil {
		m = &datasourceAuthUserGroupRemoteUserGroupsModel{}
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

func (s *datasourceAuthUserGroupModel) flattenAuthUserGroupRemoteUserGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceAuthUserGroupRemoteUserGroupsModel {
	if o == nil {
		return []datasourceAuthUserGroupRemoteUserGroupsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument remote_user_groups is not type of []interface{}.", "")
		return []datasourceAuthUserGroupRemoteUserGroupsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceAuthUserGroupRemoteUserGroupsModel{}
	}

	values := make([]datasourceAuthUserGroupRemoteUserGroupsModel, len(l))
	for i, ele := range l {
		var m datasourceAuthUserGroupRemoteUserGroupsModel
		if i < len(s.RemoteUserGroups) {
			m = s.RemoteUserGroups[i]
		}
		values[i] = *m.flattenAuthUserGroupRemoteUserGroups(ctx, ele, diags)
	}

	return values
}

func (m *datasourceAuthUserGroupRemoteUserGroupsServerModel) flattenAuthUserGroupRemoteUserGroupsServer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthUserGroupRemoteUserGroupsServerModel {
	if input == nil {
		return &datasourceAuthUserGroupRemoteUserGroupsServerModel{}
	}
	if m == nil {
		m = &datasourceAuthUserGroupRemoteUserGroupsServerModel{}
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
