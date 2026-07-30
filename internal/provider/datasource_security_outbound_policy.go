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
var _ datasource.DataSource = &datasourceSecurityOutboundPolicy{}

func newDatasourceSecurityOutboundPolicy() datasource.DataSource {
	return &datasourceSecurityOutboundPolicy{}
}

type datasourceSecurityOutboundPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityOutboundPolicyModel describes the datasource data model.
type datasourceSecurityOutboundPolicyModel struct {
	PrimaryKey          types.String                                        `tfsdk:"primary_key"`
	Enabled             types.Bool                                          `tfsdk:"enabled"`
	Scope               types.String                                        `tfsdk:"scope"`
	Users               []datasourceSecurityOutboundPolicyUsersModel        `tfsdk:"users"`
	Destinations        []datasourceSecurityOutboundPolicyDestinationsModel `tfsdk:"destinations"`
	Services            []datasourceSecurityOutboundPolicyServicesModel     `tfsdk:"services"`
	Action              types.String                                        `tfsdk:"action"`
	Schedule            *datasourceSecurityOutboundPolicyScheduleModel      `tfsdk:"schedule"`
	Comments            types.String                                        `tfsdk:"comments"`
	ProfileGroup        *datasourceSecurityOutboundPolicyProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic          types.String                                        `tfsdk:"log_traffic"`
	Sources             []datasourceSecurityOutboundPolicySourcesModel      `tfsdk:"sources"`
	CaptivePortalExempt types.Bool                                          `tfsdk:"captive_portal_exempt"`
}

func (r *datasourceSecurityOutboundPolicy) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_outbound_policy"
}

func (r *datasourceSecurityOutboundPolicy) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Outbound Policy Resource API V2 for FortiSASE.",
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
			"scope": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("all", "vpn-user", "thin-edge", "all-pre-logon-users", "specify"),
				},
				Computed: true,
			},
			"action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("accept", "deny", "isolate"),
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
			"captive_portal_exempt": schema.BoolAttribute{
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
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "security/ip-threat-feeds", "security/proxy-address-rbi", "network/internet-services"),
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"services": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("security/services", "security/service-groups"),
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
			"sources": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "endpoint/ztna-tags", "endpoint/ztna-tag-rules", "security/ip-threat-feeds", "infra/ssids", "infra/fortigates", "infra/extenders"),
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

func (r *datasourceSecurityOutboundPolicy) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_outbound_policy"
}

func (r *datasourceSecurityOutboundPolicy) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityOutboundPolicyModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityOutboundPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityOutboundPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityOutboundPolicyModel) refreshSecurityOutboundPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["scope"]; ok {
		m.Scope = parseStringValue(v)
	}

	if v, ok := o["users"]; ok {
		m.Users = m.flattenSecurityOutboundPolicyUsersList(ctx, v, &diags)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityOutboundPolicyDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityOutboundPolicyServicesList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityOutboundPolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityOutboundPolicyProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityOutboundPolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["captivePortalExempt"]; ok {
		m.CaptivePortalExempt = parseBoolValue(v)
	}

	return diags
}

func (data *datasourceSecurityOutboundPolicyModel) getURLObjectSecurityOutboundPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityOutboundPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundPolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundPolicyServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundPolicyProfileGroupModel struct {
	Group               *datasourceSecurityOutboundPolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                              `tfsdk:"force_cert_inspection"`
}

type datasourceSecurityOutboundPolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityOutboundPolicyUsersModel) flattenSecurityOutboundPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPolicyUsersModel {
	if input == nil {
		return &datasourceSecurityOutboundPolicyUsersModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPolicyUsersModel{}
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

func (s *datasourceSecurityOutboundPolicyModel) flattenSecurityOutboundPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundPolicyUsersModel {
	if o == nil {
		return []datasourceSecurityOutboundPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []datasourceSecurityOutboundPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundPolicyUsersModel{}
	}

	values := make([]datasourceSecurityOutboundPolicyUsersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityOutboundPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityOutboundPolicyDestinationsModel) flattenSecurityOutboundPolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPolicyDestinationsModel {
	if input == nil {
		return &datasourceSecurityOutboundPolicyDestinationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPolicyDestinationsModel{}
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

func (s *datasourceSecurityOutboundPolicyModel) flattenSecurityOutboundPolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundPolicyDestinationsModel {
	if o == nil {
		return []datasourceSecurityOutboundPolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []datasourceSecurityOutboundPolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundPolicyDestinationsModel{}
	}

	values := make([]datasourceSecurityOutboundPolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundPolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityOutboundPolicyDestinations(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityOutboundPolicyServicesModel) flattenSecurityOutboundPolicyServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPolicyServicesModel {
	if input == nil {
		return &datasourceSecurityOutboundPolicyServicesModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPolicyServicesModel{}
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

func (s *datasourceSecurityOutboundPolicyModel) flattenSecurityOutboundPolicyServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundPolicyServicesModel {
	if o == nil {
		return []datasourceSecurityOutboundPolicyServicesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []datasourceSecurityOutboundPolicyServicesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundPolicyServicesModel{}
	}

	values := make([]datasourceSecurityOutboundPolicyServicesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundPolicyServicesModel
		if i < len(s.Services) {
			m = s.Services[i]
		}
		values[i] = *m.flattenSecurityOutboundPolicyServices(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityOutboundPolicyScheduleModel) flattenSecurityOutboundPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPolicyScheduleModel {
	if input == nil {
		return &datasourceSecurityOutboundPolicyScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPolicyScheduleModel{}
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

func (m *datasourceSecurityOutboundPolicyProfileGroupModel) flattenSecurityOutboundPolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPolicyProfileGroupModel {
	if input == nil {
		return &datasourceSecurityOutboundPolicyProfileGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPolicyProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityOutboundPolicyProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *datasourceSecurityOutboundPolicyProfileGroupGroupModel) flattenSecurityOutboundPolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPolicyProfileGroupGroupModel {
	if input == nil {
		return &datasourceSecurityOutboundPolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPolicyProfileGroupGroupModel{}
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

func (m *datasourceSecurityOutboundPolicySourcesModel) flattenSecurityOutboundPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPolicySourcesModel {
	if input == nil {
		return &datasourceSecurityOutboundPolicySourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPolicySourcesModel{}
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

func (s *datasourceSecurityOutboundPolicyModel) flattenSecurityOutboundPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundPolicySourcesModel {
	if o == nil {
		return []datasourceSecurityOutboundPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityOutboundPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundPolicySourcesModel{}
	}

	values := make([]datasourceSecurityOutboundPolicySourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityOutboundPolicySources(ctx, ele, diags)
	}

	return values
}
