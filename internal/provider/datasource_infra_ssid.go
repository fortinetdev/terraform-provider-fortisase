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
var _ datasource.DataSource = &datasourceInfraSsid{}

func newDatasourceInfraSsid() datasource.DataSource {
	return &datasourceInfraSsid{}
}

type datasourceInfraSsid struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceInfraSsidModel describes the datasource data model.
type datasourceInfraSsidModel struct {
	PrimaryKey     types.String                             `tfsdk:"primary_key"`
	WifiSsid       types.String                             `tfsdk:"wifi_ssid"`
	BroadcastSsid  types.String                             `tfsdk:"broadcast_ssid"`
	ClientLimit    types.Float64                            `tfsdk:"client_limit"`
	SecurityMode   types.String                             `tfsdk:"security_mode"`
	CaptivePortal  types.Bool                               `tfsdk:"captive_portal"`
	SecurityGroups []datasourceInfraSsidSecurityGroupsModel `tfsdk:"security_groups"`
	PreSharedKey   types.String                             `tfsdk:"pre_shared_key"`
	RadiusServer   *datasourceInfraSsidRadiusServerModel    `tfsdk:"radius_server"`
	UserGroups     []datasourceInfraSsidUserGroupsModel     `tfsdk:"user_groups"`
}

func (r *datasourceInfraSsid) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_ssid"
}

func (r *datasourceInfraSsid) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "FortiAP SSID Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 15),
				},
				Required: true,
			},
			"wifi_ssid": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 32),
				},
				Computed: true,
			},
			"broadcast_ssid": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"client_limit": schema.Float64Attribute{
				Computed: true,
			},
			"security_mode": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("wpa2-only-personal", "wpa2-only-enterprise", "wpa3-only-enterprise", "wpa3-sae", "open", "wpa2-only-personal+captive-portal", "captive-portal"),
				},
				Computed: true,
			},
			"captive_portal": schema.BoolAttribute{
				Computed: true,
			},
			"pre_shared_key": schema.StringAttribute{
				Computed: true,
			},
			"security_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("auth/user-groups"),
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"radius_server": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("auth/radius-servers"),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"user_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("auth/user-groups"),
							},
							Computed: true,
						},
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceInfraSsid) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceInfraSsid) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceInfraSsidModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectInfraSsid(ctx, "read", diags))

	read_output, err := c.ReadInfraSsids(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
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

func (m *datasourceInfraSsidModel) refreshInfraSsid(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

	if v, ok := o["preSharedKey"]; ok {
		m.PreSharedKey = parseStringValue(v)
	}

	if v, ok := o["radiusServer"]; ok {
		m.RadiusServer = m.RadiusServer.flattenInfraSsidRadiusServer(ctx, v, &diags)
	}

	if v, ok := o["userGroups"]; ok {
		m.UserGroups = m.flattenInfraSsidUserGroupsList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceInfraSsidModel) getURLObjectInfraSsid(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceInfraSsidSecurityGroupsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceInfraSsidRadiusServerModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceInfraSsidUserGroupsModel struct {
	Datasource types.String `tfsdk:"datasource"`
	PrimaryKey types.String `tfsdk:"primary_key"`
}

func (m *datasourceInfraSsidSecurityGroupsModel) flattenInfraSsidSecurityGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceInfraSsidSecurityGroupsModel {
	if input == nil {
		return &datasourceInfraSsidSecurityGroupsModel{}
	}
	if m == nil {
		m = &datasourceInfraSsidSecurityGroupsModel{}
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

func (s *datasourceInfraSsidModel) flattenInfraSsidSecurityGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceInfraSsidSecurityGroupsModel {
	if o == nil {
		return []datasourceInfraSsidSecurityGroupsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument security_groups is not type of []interface{}.", "")
		return []datasourceInfraSsidSecurityGroupsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceInfraSsidSecurityGroupsModel{}
	}

	values := make([]datasourceInfraSsidSecurityGroupsModel, len(l))
	for i, ele := range l {
		var m datasourceInfraSsidSecurityGroupsModel
		if i < len(s.SecurityGroups) {
			m = s.SecurityGroups[i]
		}
		values[i] = *m.flattenInfraSsidSecurityGroups(ctx, ele, diags)
	}

	return values
}

func (m *datasourceInfraSsidRadiusServerModel) flattenInfraSsidRadiusServer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceInfraSsidRadiusServerModel {
	if input == nil {
		return &datasourceInfraSsidRadiusServerModel{}
	}
	if m == nil {
		m = &datasourceInfraSsidRadiusServerModel{}
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

func (m *datasourceInfraSsidUserGroupsModel) flattenInfraSsidUserGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceInfraSsidUserGroupsModel {
	if input == nil {
		return &datasourceInfraSsidUserGroupsModel{}
	}
	if m == nil {
		m = &datasourceInfraSsidUserGroupsModel{}
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

func (s *datasourceInfraSsidModel) flattenInfraSsidUserGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceInfraSsidUserGroupsModel {
	if o == nil {
		return []datasourceInfraSsidUserGroupsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument user_groups is not type of []interface{}.", "")
		return []datasourceInfraSsidUserGroupsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceInfraSsidUserGroupsModel{}
	}

	values := make([]datasourceInfraSsidUserGroupsModel, len(l))
	for i, ele := range l {
		var m datasourceInfraSsidUserGroupsModel
		if i < len(s.UserGroups) {
			m = s.UserGroups[i]
		}
		values[i] = *m.flattenInfraSsidUserGroups(ctx, ele, diags)
	}

	return values
}
