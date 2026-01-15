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
var _ datasource.DataSource = &datasourceInfraSsids{}

func newDatasourceInfraSsids() datasource.DataSource {
	return &datasourceInfraSsids{}
}

type datasourceInfraSsids struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceInfraSsidsModel describes the datasource data model.
type datasourceInfraSsidsModel struct {
	PrimaryKey     types.String                              `tfsdk:"primary_key"`
	WifiSsid       types.String                              `tfsdk:"wifi_ssid"`
	BroadcastSsid  types.String                              `tfsdk:"broadcast_ssid"`
	ClientLimit    types.Float64                             `tfsdk:"client_limit"`
	SecurityMode   types.String                              `tfsdk:"security_mode"`
	CaptivePortal  types.Bool                                `tfsdk:"captive_portal"`
	SecurityGroups []datasourceInfraSsidsSecurityGroupsModel `tfsdk:"security_groups"`
	PreSharedKey   types.String                              `tfsdk:"pre_shared_key"`
	RadiusServer   *datasourceInfraSsidsRadiusServerModel    `tfsdk:"radius_server"`
	UserGroups     []datasourceInfraSsidsUserGroupsModel     `tfsdk:"user_groups"`
}

func (r *datasourceInfraSsids) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_ssids"
}

func (r *datasourceInfraSsids) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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

func (r *datasourceInfraSsids) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceInfraSsids) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceInfraSsidsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectInfraSsids(ctx, "read", diags))

	read_output, err := c.ReadInfraSsids(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
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

func (m *datasourceInfraSsidsModel) refreshInfraSsids(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *datasourceInfraSsidsModel) getURLObjectInfraSsids(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceInfraSsidsSecurityGroupsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceInfraSsidsRadiusServerModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceInfraSsidsUserGroupsModel struct {
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceInfraSsidsSecurityGroupsModel) flattenInfraSsidsSecurityGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceInfraSsidsSecurityGroupsModel {
	if input == nil {
		return &datasourceInfraSsidsSecurityGroupsModel{}
	}
	if m == nil {
		m = &datasourceInfraSsidsSecurityGroupsModel{}
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

func (s *datasourceInfraSsidsModel) flattenInfraSsidsSecurityGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceInfraSsidsSecurityGroupsModel {
	if o == nil {
		return []datasourceInfraSsidsSecurityGroupsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument security_groups is not type of []interface{}.", "")
		return []datasourceInfraSsidsSecurityGroupsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceInfraSsidsSecurityGroupsModel{}
	}

	values := make([]datasourceInfraSsidsSecurityGroupsModel, len(l))
	for i, ele := range l {
		var m datasourceInfraSsidsSecurityGroupsModel
		values[i] = *m.flattenInfraSsidsSecurityGroups(ctx, ele, diags)
	}

	return values
}

func (m *datasourceInfraSsidsRadiusServerModel) flattenInfraSsidsRadiusServer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceInfraSsidsRadiusServerModel {
	if input == nil {
		return &datasourceInfraSsidsRadiusServerModel{}
	}
	if m == nil {
		m = &datasourceInfraSsidsRadiusServerModel{}
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

func (m *datasourceInfraSsidsUserGroupsModel) flattenInfraSsidsUserGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceInfraSsidsUserGroupsModel {
	if input == nil {
		return &datasourceInfraSsidsUserGroupsModel{}
	}
	if m == nil {
		m = &datasourceInfraSsidsUserGroupsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *datasourceInfraSsidsModel) flattenInfraSsidsUserGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceInfraSsidsUserGroupsModel {
	if o == nil {
		return []datasourceInfraSsidsUserGroupsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument user_groups is not type of []interface{}.", "")
		return []datasourceInfraSsidsUserGroupsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceInfraSsidsUserGroupsModel{}
	}

	values := make([]datasourceInfraSsidsUserGroupsModel, len(l))
	for i, ele := range l {
		var m datasourceInfraSsidsUserGroupsModel
		values[i] = *m.flattenInfraSsidsUserGroups(ctx, ele, diags)
	}

	return values
}
