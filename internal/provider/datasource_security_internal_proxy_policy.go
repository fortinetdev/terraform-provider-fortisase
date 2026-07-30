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
var _ datasource.DataSource = &datasourceSecurityInternalProxyPolicy{}

func newDatasourceSecurityInternalProxyPolicy() datasource.DataSource {
	return &datasourceSecurityInternalProxyPolicy{}
}

type datasourceSecurityInternalProxyPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityInternalProxyPolicyModel describes the datasource data model.
type datasourceSecurityInternalProxyPolicyModel struct {
	PrimaryKey   types.String                                             `tfsdk:"primary_key"`
	Enabled      types.Bool                                               `tfsdk:"enabled"`
	Sources      []datasourceSecurityInternalProxyPolicySourcesModel      `tfsdk:"sources"`
	Users        []datasourceSecurityInternalProxyPolicyUsersModel        `tfsdk:"users"`
	Destinations []datasourceSecurityInternalProxyPolicyDestinationsModel `tfsdk:"destinations"`
	Action       types.String                                             `tfsdk:"action"`
	Schedule     *datasourceSecurityInternalProxyPolicyScheduleModel      `tfsdk:"schedule"`
	Comments     types.String                                             `tfsdk:"comments"`
	ProfileGroup *datasourceSecurityInternalProxyPolicyProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic   types.String                                             `tfsdk:"log_traffic"`
}

func (r *datasourceSecurityInternalProxyPolicy) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_proxy_policy"
}

func (r *datasourceSecurityInternalProxyPolicy) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Internal Proxy Policy Resource API V2 for FortiSASE.",
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
								stringvalidatorwarning.OneOf("auth/users", "auth/user-groups"),
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
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "security/ip-threat-feeds"),
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

func (r *datasourceSecurityInternalProxyPolicy) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_internal_proxy_policy"
}

func (r *datasourceSecurityInternalProxyPolicy) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityInternalProxyPolicyModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalProxyPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalProxyPolicy(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalProxyPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityInternalProxyPolicyModel) refreshSecurityInternalProxyPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityInternalProxyPolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["users"]; ok {
		m.Users = m.flattenSecurityInternalProxyPolicyUsersList(ctx, v, &diags)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityInternalProxyPolicyDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityInternalProxyPolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityInternalProxyPolicyProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	return diags
}

func (data *datasourceSecurityInternalProxyPolicyModel) getURLObjectSecurityInternalProxyPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityInternalProxyPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalProxyPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalProxyPolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalProxyPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalProxyPolicyProfileGroupModel struct {
	Group               *datasourceSecurityInternalProxyPolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                   `tfsdk:"force_cert_inspection"`
}

type datasourceSecurityInternalProxyPolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityInternalProxyPolicySourcesModel) flattenSecurityInternalProxyPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalProxyPolicySourcesModel {
	if input == nil {
		return &datasourceSecurityInternalProxyPolicySourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalProxyPolicySourcesModel{}
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

func (s *datasourceSecurityInternalProxyPolicyModel) flattenSecurityInternalProxyPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalProxyPolicySourcesModel {
	if o == nil {
		return []datasourceSecurityInternalProxyPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityInternalProxyPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalProxyPolicySourcesModel{}
	}

	values := make([]datasourceSecurityInternalProxyPolicySourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalProxyPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityInternalProxyPolicySources(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalProxyPolicyUsersModel) flattenSecurityInternalProxyPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalProxyPolicyUsersModel {
	if input == nil {
		return &datasourceSecurityInternalProxyPolicyUsersModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalProxyPolicyUsersModel{}
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

func (s *datasourceSecurityInternalProxyPolicyModel) flattenSecurityInternalProxyPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalProxyPolicyUsersModel {
	if o == nil {
		return []datasourceSecurityInternalProxyPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []datasourceSecurityInternalProxyPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalProxyPolicyUsersModel{}
	}

	values := make([]datasourceSecurityInternalProxyPolicyUsersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalProxyPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityInternalProxyPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalProxyPolicyDestinationsModel) flattenSecurityInternalProxyPolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalProxyPolicyDestinationsModel {
	if input == nil {
		return &datasourceSecurityInternalProxyPolicyDestinationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalProxyPolicyDestinationsModel{}
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

func (s *datasourceSecurityInternalProxyPolicyModel) flattenSecurityInternalProxyPolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalProxyPolicyDestinationsModel {
	if o == nil {
		return []datasourceSecurityInternalProxyPolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []datasourceSecurityInternalProxyPolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalProxyPolicyDestinationsModel{}
	}

	values := make([]datasourceSecurityInternalProxyPolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalProxyPolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityInternalProxyPolicyDestinations(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalProxyPolicyScheduleModel) flattenSecurityInternalProxyPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalProxyPolicyScheduleModel {
	if input == nil {
		return &datasourceSecurityInternalProxyPolicyScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalProxyPolicyScheduleModel{}
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

func (m *datasourceSecurityInternalProxyPolicyProfileGroupModel) flattenSecurityInternalProxyPolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalProxyPolicyProfileGroupModel {
	if input == nil {
		return &datasourceSecurityInternalProxyPolicyProfileGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalProxyPolicyProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityInternalProxyPolicyProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *datasourceSecurityInternalProxyPolicyProfileGroupGroupModel) flattenSecurityInternalProxyPolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalProxyPolicyProfileGroupGroupModel {
	if input == nil {
		return &datasourceSecurityInternalProxyPolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalProxyPolicyProfileGroupGroupModel{}
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
