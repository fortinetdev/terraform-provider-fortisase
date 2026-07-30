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
var _ datasource.DataSource = &datasourceSecurityInternalPolicy{}

func newDatasourceSecurityInternalPolicy() datasource.DataSource {
	return &datasourceSecurityInternalPolicy{}
}

type datasourceSecurityInternalPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityInternalPolicyModel describes the datasource data model.
type datasourceSecurityInternalPolicyModel struct {
	PrimaryKey          types.String                                        `tfsdk:"primary_key"`
	Enabled             types.Bool                                          `tfsdk:"enabled"`
	Scope               types.String                                        `tfsdk:"scope"`
	Users               []datasourceSecurityInternalPolicyUsersModel        `tfsdk:"users"`
	Destinations        []datasourceSecurityInternalPolicyDestinationsModel `tfsdk:"destinations"`
	Services            []datasourceSecurityInternalPolicyServicesModel     `tfsdk:"services"`
	Action              types.String                                        `tfsdk:"action"`
	Schedule            *datasourceSecurityInternalPolicyScheduleModel      `tfsdk:"schedule"`
	Comments            types.String                                        `tfsdk:"comments"`
	ProfileGroup        *datasourceSecurityInternalPolicyProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic          types.String                                        `tfsdk:"log_traffic"`
	Sources             []datasourceSecurityInternalPolicySourcesModel      `tfsdk:"sources"`
	CaptivePortalExempt types.Bool                                          `tfsdk:"captive_portal_exempt"`
}

func (r *datasourceSecurityInternalPolicy) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_policy"
}

func (r *datasourceSecurityInternalPolicy) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Internal Policy Resource API V2 for FortiSASE.",
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

func (r *datasourceSecurityInternalPolicy) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_internal_policy"
}

func (r *datasourceSecurityInternalPolicy) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityInternalPolicyModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityInternalPolicyModel) refreshSecurityInternalPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.Users = m.flattenSecurityInternalPolicyUsersList(ctx, v, &diags)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityInternalPolicyDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityInternalPolicyServicesList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityInternalPolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityInternalPolicyProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityInternalPolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["captivePortalExempt"]; ok {
		m.CaptivePortalExempt = parseBoolValue(v)
	}

	return diags
}

func (data *datasourceSecurityInternalPolicyModel) getURLObjectSecurityInternalPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityInternalPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalPolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalPolicyServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalPolicyProfileGroupModel struct {
	Group               *datasourceSecurityInternalPolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                              `tfsdk:"force_cert_inspection"`
}

type datasourceSecurityInternalPolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityInternalPolicyUsersModel) flattenSecurityInternalPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPolicyUsersModel {
	if input == nil {
		return &datasourceSecurityInternalPolicyUsersModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPolicyUsersModel{}
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

func (s *datasourceSecurityInternalPolicyModel) flattenSecurityInternalPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalPolicyUsersModel {
	if o == nil {
		return []datasourceSecurityInternalPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []datasourceSecurityInternalPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalPolicyUsersModel{}
	}

	values := make([]datasourceSecurityInternalPolicyUsersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityInternalPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalPolicyDestinationsModel) flattenSecurityInternalPolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPolicyDestinationsModel {
	if input == nil {
		return &datasourceSecurityInternalPolicyDestinationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPolicyDestinationsModel{}
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

func (s *datasourceSecurityInternalPolicyModel) flattenSecurityInternalPolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalPolicyDestinationsModel {
	if o == nil {
		return []datasourceSecurityInternalPolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []datasourceSecurityInternalPolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalPolicyDestinationsModel{}
	}

	values := make([]datasourceSecurityInternalPolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalPolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityInternalPolicyDestinations(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalPolicyServicesModel) flattenSecurityInternalPolicyServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPolicyServicesModel {
	if input == nil {
		return &datasourceSecurityInternalPolicyServicesModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPolicyServicesModel{}
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

func (s *datasourceSecurityInternalPolicyModel) flattenSecurityInternalPolicyServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalPolicyServicesModel {
	if o == nil {
		return []datasourceSecurityInternalPolicyServicesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []datasourceSecurityInternalPolicyServicesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalPolicyServicesModel{}
	}

	values := make([]datasourceSecurityInternalPolicyServicesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalPolicyServicesModel
		if i < len(s.Services) {
			m = s.Services[i]
		}
		values[i] = *m.flattenSecurityInternalPolicyServices(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalPolicyScheduleModel) flattenSecurityInternalPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPolicyScheduleModel {
	if input == nil {
		return &datasourceSecurityInternalPolicyScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPolicyScheduleModel{}
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

func (m *datasourceSecurityInternalPolicyProfileGroupModel) flattenSecurityInternalPolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPolicyProfileGroupModel {
	if input == nil {
		return &datasourceSecurityInternalPolicyProfileGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPolicyProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityInternalPolicyProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *datasourceSecurityInternalPolicyProfileGroupGroupModel) flattenSecurityInternalPolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPolicyProfileGroupGroupModel {
	if input == nil {
		return &datasourceSecurityInternalPolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPolicyProfileGroupGroupModel{}
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

func (m *datasourceSecurityInternalPolicySourcesModel) flattenSecurityInternalPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPolicySourcesModel {
	if input == nil {
		return &datasourceSecurityInternalPolicySourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPolicySourcesModel{}
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

func (s *datasourceSecurityInternalPolicyModel) flattenSecurityInternalPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalPolicySourcesModel {
	if o == nil {
		return []datasourceSecurityInternalPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityInternalPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalPolicySourcesModel{}
	}

	values := make([]datasourceSecurityInternalPolicySourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityInternalPolicySources(ctx, ele, diags)
	}

	return values
}
