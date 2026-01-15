// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceAuthUserGroups{}

func newDatasourceAuthUserGroups() datasource.DataSource {
	return &datasourceAuthUserGroups{}
}

type datasourceAuthUserGroups struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceAuthUserGroupsModel describes the datasource data model.
type datasourceAuthUserGroupsModel struct {
	PrimaryKey       types.String                                    `tfsdk:"primary_key"`
	GroupType        types.String                                    `tfsdk:"group_type"`
	LocalUsers       []datasourceAuthUserGroupsLocalUsersModel       `tfsdk:"local_users"`
	RemoteUserGroups []datasourceAuthUserGroupsRemoteUserGroupsModel `tfsdk:"remote_user_groups"`
}

func (r *datasourceAuthUserGroups) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_user_groups"
}

func (r *datasourceAuthUserGroups) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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

func (r *datasourceAuthUserGroups) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceAuthUserGroups) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceAuthUserGroupsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthUserGroups(ctx, "read", diags))

	read_output, err := c.ReadAuthUserGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
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

func (m *datasourceAuthUserGroupsModel) refreshAuthUserGroups(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *datasourceAuthUserGroupsModel) getURLObjectAuthUserGroups(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceAuthUserGroupsLocalUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceAuthUserGroupsRemoteUserGroupsModel struct {
	Server  *datasourceAuthUserGroupsRemoteUserGroupsServerModel `tfsdk:"server"`
	Matches types.Set                                            `tfsdk:"matches"`
}

type datasourceAuthUserGroupsRemoteUserGroupsServerModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceAuthUserGroupsLocalUsersModel) flattenAuthUserGroupsLocalUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthUserGroupsLocalUsersModel {
	if input == nil {
		return &datasourceAuthUserGroupsLocalUsersModel{}
	}
	if m == nil {
		m = &datasourceAuthUserGroupsLocalUsersModel{}
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

func (s *datasourceAuthUserGroupsModel) flattenAuthUserGroupsLocalUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceAuthUserGroupsLocalUsersModel {
	if o == nil {
		return []datasourceAuthUserGroupsLocalUsersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument local_users is not type of []interface{}.", "")
		return []datasourceAuthUserGroupsLocalUsersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceAuthUserGroupsLocalUsersModel{}
	}

	values := make([]datasourceAuthUserGroupsLocalUsersModel, len(l))
	for i, ele := range l {
		var m datasourceAuthUserGroupsLocalUsersModel
		values[i] = *m.flattenAuthUserGroupsLocalUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceAuthUserGroupsRemoteUserGroupsModel) flattenAuthUserGroupsRemoteUserGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthUserGroupsRemoteUserGroupsModel {
	if input == nil {
		return &datasourceAuthUserGroupsRemoteUserGroupsModel{}
	}
	if m == nil {
		m = &datasourceAuthUserGroupsRemoteUserGroupsModel{}
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

func (s *datasourceAuthUserGroupsModel) flattenAuthUserGroupsRemoteUserGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceAuthUserGroupsRemoteUserGroupsModel {
	if o == nil {
		return []datasourceAuthUserGroupsRemoteUserGroupsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument remote_user_groups is not type of []interface{}.", "")
		return []datasourceAuthUserGroupsRemoteUserGroupsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceAuthUserGroupsRemoteUserGroupsModel{}
	}

	values := make([]datasourceAuthUserGroupsRemoteUserGroupsModel, len(l))
	for i, ele := range l {
		var m datasourceAuthUserGroupsRemoteUserGroupsModel
		values[i] = *m.flattenAuthUserGroupsRemoteUserGroups(ctx, ele, diags)
	}

	return values
}

func (m *datasourceAuthUserGroupsRemoteUserGroupsServerModel) flattenAuthUserGroupsRemoteUserGroupsServer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthUserGroupsRemoteUserGroupsServerModel {
	if input == nil {
		return &datasourceAuthUserGroupsRemoteUserGroupsServerModel{}
	}
	if m == nil {
		m = &datasourceAuthUserGroupsRemoteUserGroupsServerModel{}
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
