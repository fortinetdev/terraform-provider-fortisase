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
var _ datasource.DataSource = &datasourceSecurityOutboundProxyPolicy{}

func newDatasourceSecurityOutboundProxyPolicy() datasource.DataSource {
	return &datasourceSecurityOutboundProxyPolicy{}
}

type datasourceSecurityOutboundProxyPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityOutboundProxyPolicyModel describes the datasource data model.
type datasourceSecurityOutboundProxyPolicyModel struct {
	PrimaryKey   types.String                                             `tfsdk:"primary_key"`
	Enabled      types.Bool                                               `tfsdk:"enabled"`
	Sources      []datasourceSecurityOutboundProxyPolicySourcesModel      `tfsdk:"sources"`
	Users        []datasourceSecurityOutboundProxyPolicyUsersModel        `tfsdk:"users"`
	Destinations []datasourceSecurityOutboundProxyPolicyDestinationsModel `tfsdk:"destinations"`
	Action       types.String                                             `tfsdk:"action"`
	Schedule     *datasourceSecurityOutboundProxyPolicyScheduleModel      `tfsdk:"schedule"`
	Comments     types.String                                             `tfsdk:"comments"`
	ProfileGroup *datasourceSecurityOutboundProxyPolicyProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic   types.String                                             `tfsdk:"log_traffic"`
}

func (r *datasourceSecurityOutboundProxyPolicy) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_outbound_proxy_policy"
}

func (r *datasourceSecurityOutboundProxyPolicy) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Outbound Proxy Policy Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 35),
				},
				Required: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("accept", "deny"),
				},
				Computed: true,
			},
			"comments": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(1023),
				},
				Computed: true,
			},
			"log_traffic": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("all", "utm", "disable"),
				},
				Computed: true,
			},
			"sources": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "endpoint/ztna-tags", "endpoint/ztna-tag-rules", "security/ip-threat-feeds"),
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"users": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("auth/users", "auth/user-groups", "auth/ad-groups"),
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"destinations": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "security/ip-threat-feeds", "network/internet-services"),
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"schedule": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("security/onetime-schedules", "security/recurring-schedules", "security/schedule-groups"),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"profile_group": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"force_cert_inspection": schema.BoolAttribute{
						Computed: true,
					},
					"group": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.OneOf("security/profile-groups"),
								},
								Computed: true,
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceSecurityOutboundProxyPolicy) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_outbound_proxy_policy"
}

func (r *datasourceSecurityOutboundProxyPolicy) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityOutboundProxyPolicyModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityOutboundProxyPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityOutboundProxyPolicy(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityOutboundProxyPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityOutboundProxyPolicyModel) refreshSecurityOutboundProxyPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityOutboundProxyPolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["users"]; ok {
		m.Users = m.flattenSecurityOutboundProxyPolicyUsersList(ctx, v, &diags)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityOutboundProxyPolicyDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityOutboundProxyPolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityOutboundProxyPolicyProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	return diags
}

func (data *datasourceSecurityOutboundProxyPolicyModel) getURLObjectSecurityOutboundProxyPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityOutboundProxyPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundProxyPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundProxyPolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundProxyPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundProxyPolicyProfileGroupModel struct {
	Group               *datasourceSecurityOutboundProxyPolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                   `tfsdk:"force_cert_inspection"`
}

type datasourceSecurityOutboundProxyPolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityOutboundProxyPolicySourcesModel) flattenSecurityOutboundProxyPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundProxyPolicySourcesModel {
	if input == nil {
		return &datasourceSecurityOutboundProxyPolicySourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundProxyPolicySourcesModel{}
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

func (s *datasourceSecurityOutboundProxyPolicyModel) flattenSecurityOutboundProxyPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundProxyPolicySourcesModel {
	if o == nil {
		return []datasourceSecurityOutboundProxyPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityOutboundProxyPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundProxyPolicySourcesModel{}
	}

	values := make([]datasourceSecurityOutboundProxyPolicySourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundProxyPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityOutboundProxyPolicySources(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityOutboundProxyPolicyUsersModel) flattenSecurityOutboundProxyPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundProxyPolicyUsersModel {
	if input == nil {
		return &datasourceSecurityOutboundProxyPolicyUsersModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundProxyPolicyUsersModel{}
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

func (s *datasourceSecurityOutboundProxyPolicyModel) flattenSecurityOutboundProxyPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundProxyPolicyUsersModel {
	if o == nil {
		return []datasourceSecurityOutboundProxyPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []datasourceSecurityOutboundProxyPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundProxyPolicyUsersModel{}
	}

	values := make([]datasourceSecurityOutboundProxyPolicyUsersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundProxyPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityOutboundProxyPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityOutboundProxyPolicyDestinationsModel) flattenSecurityOutboundProxyPolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundProxyPolicyDestinationsModel {
	if input == nil {
		return &datasourceSecurityOutboundProxyPolicyDestinationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundProxyPolicyDestinationsModel{}
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

func (s *datasourceSecurityOutboundProxyPolicyModel) flattenSecurityOutboundProxyPolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundProxyPolicyDestinationsModel {
	if o == nil {
		return []datasourceSecurityOutboundProxyPolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []datasourceSecurityOutboundProxyPolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundProxyPolicyDestinationsModel{}
	}

	values := make([]datasourceSecurityOutboundProxyPolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundProxyPolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityOutboundProxyPolicyDestinations(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityOutboundProxyPolicyScheduleModel) flattenSecurityOutboundProxyPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundProxyPolicyScheduleModel {
	if input == nil {
		return &datasourceSecurityOutboundProxyPolicyScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundProxyPolicyScheduleModel{}
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

func (m *datasourceSecurityOutboundProxyPolicyProfileGroupModel) flattenSecurityOutboundProxyPolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundProxyPolicyProfileGroupModel {
	if input == nil {
		return &datasourceSecurityOutboundProxyPolicyProfileGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundProxyPolicyProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityOutboundProxyPolicyProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *datasourceSecurityOutboundProxyPolicyProfileGroupGroupModel) flattenSecurityOutboundProxyPolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundProxyPolicyProfileGroupGroupModel {
	if input == nil {
		return &datasourceSecurityOutboundProxyPolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundProxyPolicyProfileGroupGroupModel{}
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
