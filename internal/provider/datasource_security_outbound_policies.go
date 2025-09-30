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
var _ datasource.DataSource = &datasourceSecurityOutboundPolicies{}

func newDatasourceSecurityOutboundPolicies() datasource.DataSource {
	return &datasourceSecurityOutboundPolicies{}
}

type datasourceSecurityOutboundPolicies struct {
	fortiClient *FortiClient
}

// datasourceSecurityOutboundPoliciesModel describes the datasource data model.
type datasourceSecurityOutboundPoliciesModel struct {
	PrimaryKey          types.String                                          `tfsdk:"primary_key"`
	Enabled             types.Bool                                            `tfsdk:"enabled"`
	Scope               types.String                                          `tfsdk:"scope"`
	Users               []datasourceSecurityOutboundPoliciesUsersModel        `tfsdk:"users"`
	Destinations        []datasourceSecurityOutboundPoliciesDestinationsModel `tfsdk:"destinations"`
	Services            []datasourceSecurityOutboundPoliciesServicesModel     `tfsdk:"services"`
	Action              types.String                                          `tfsdk:"action"`
	Schedule            *datasourceSecurityOutboundPoliciesScheduleModel      `tfsdk:"schedule"`
	Comments            types.String                                          `tfsdk:"comments"`
	ProfileGroup        *datasourceSecurityOutboundPoliciesProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic          types.String                                          `tfsdk:"log_traffic"`
	Sources             []datasourceSecurityOutboundPoliciesSourcesModel      `tfsdk:"sources"`
	CaptivePortalExempt types.Bool                                            `tfsdk:"captive_portal_exempt"`
}

func (r *datasourceSecurityOutboundPolicies) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_outbound_policies"
}

func (r *datasourceSecurityOutboundPolicies) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 35),
				},
				Required: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"scope": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("all", "vpn-user", "thin-edge", "specify"),
				},
				Computed: true,
				Optional: true,
			},
			"action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("accept", "deny"),
				},
				Computed: true,
				Optional: true,
			},
			"comments": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(1023),
				},
				Computed: true,
				Optional: true,
			},
			"log_traffic": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("all", "utm", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"captive_portal_exempt": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"users": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("auth/users", "auth/user-groups", "auth/ad-groups"),
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"destinations": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("network/hosts", "network/host-groups", "security/ip-threat-feeds", "network/internet-services"),
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"services": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("security/services", "security/service-groups"),
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"schedule": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("security/onetime-schedules", "security/recurring-schedules", "security/schedule-groups"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"profile_group": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"force_cert_inspection": schema.BoolAttribute{
						Computed: true,
						Optional: true,
					},
					"group": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("security/profile-groups"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"sources": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("network/hosts", "network/host-groups", "endpoint/ztna-tags", "security/ip-threat-feeds", "infra/ssids", "infra/fortigates", "infra/extenders"),
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

func (r *datasourceSecurityOutboundPolicies) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
}

func (r *datasourceSecurityOutboundPolicies) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityOutboundPoliciesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityOutboundPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityOutboundPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityOutboundPoliciesModel) refreshSecurityOutboundPolicies(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["scope"]; ok {
		m.Scope = parseStringValue(v)
	}

	if v, ok := o["users"]; ok {
		m.Users = m.flattenSecurityOutboundPoliciesUsersList(ctx, v, &diags)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityOutboundPoliciesDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityOutboundPoliciesServicesList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityOutboundPoliciesSchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityOutboundPoliciesProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityOutboundPoliciesSourcesList(ctx, v, &diags)
	}

	if v, ok := o["captivePortalExempt"]; ok {
		m.CaptivePortalExempt = parseBoolValue(v)
	}

	return diags
}

func (data *datasourceSecurityOutboundPoliciesModel) getURLObjectSecurityOutboundPolicies(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityOutboundPoliciesUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundPoliciesDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundPoliciesServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundPoliciesScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundPoliciesProfileGroupModel struct {
	Group               *datasourceSecurityOutboundPoliciesProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                `tfsdk:"force_cert_inspection"`
}

type datasourceSecurityOutboundPoliciesProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityOutboundPoliciesSourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityOutboundPoliciesUsersModel) flattenSecurityOutboundPoliciesUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPoliciesUsersModel {
	if input == nil {
		return &datasourceSecurityOutboundPoliciesUsersModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPoliciesUsersModel{}
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

func (s *datasourceSecurityOutboundPoliciesModel) flattenSecurityOutboundPoliciesUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundPoliciesUsersModel {
	if o == nil {
		return []datasourceSecurityOutboundPoliciesUsersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []datasourceSecurityOutboundPoliciesUsersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundPoliciesUsersModel{}
	}

	values := make([]datasourceSecurityOutboundPoliciesUsersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundPoliciesUsersModel
		values[i] = *m.flattenSecurityOutboundPoliciesUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityOutboundPoliciesDestinationsModel) flattenSecurityOutboundPoliciesDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPoliciesDestinationsModel {
	if input == nil {
		return &datasourceSecurityOutboundPoliciesDestinationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPoliciesDestinationsModel{}
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

func (s *datasourceSecurityOutboundPoliciesModel) flattenSecurityOutboundPoliciesDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundPoliciesDestinationsModel {
	if o == nil {
		return []datasourceSecurityOutboundPoliciesDestinationsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []datasourceSecurityOutboundPoliciesDestinationsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundPoliciesDestinationsModel{}
	}

	values := make([]datasourceSecurityOutboundPoliciesDestinationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundPoliciesDestinationsModel
		values[i] = *m.flattenSecurityOutboundPoliciesDestinations(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityOutboundPoliciesServicesModel) flattenSecurityOutboundPoliciesServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPoliciesServicesModel {
	if input == nil {
		return &datasourceSecurityOutboundPoliciesServicesModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPoliciesServicesModel{}
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

func (s *datasourceSecurityOutboundPoliciesModel) flattenSecurityOutboundPoliciesServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundPoliciesServicesModel {
	if o == nil {
		return []datasourceSecurityOutboundPoliciesServicesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []datasourceSecurityOutboundPoliciesServicesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundPoliciesServicesModel{}
	}

	values := make([]datasourceSecurityOutboundPoliciesServicesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundPoliciesServicesModel
		values[i] = *m.flattenSecurityOutboundPoliciesServices(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityOutboundPoliciesScheduleModel) flattenSecurityOutboundPoliciesSchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPoliciesScheduleModel {
	if input == nil {
		return &datasourceSecurityOutboundPoliciesScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPoliciesScheduleModel{}
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

func (m *datasourceSecurityOutboundPoliciesProfileGroupModel) flattenSecurityOutboundPoliciesProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPoliciesProfileGroupModel {
	if input == nil {
		return &datasourceSecurityOutboundPoliciesProfileGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPoliciesProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityOutboundPoliciesProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *datasourceSecurityOutboundPoliciesProfileGroupGroupModel) flattenSecurityOutboundPoliciesProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPoliciesProfileGroupGroupModel {
	if input == nil {
		return &datasourceSecurityOutboundPoliciesProfileGroupGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPoliciesProfileGroupGroupModel{}
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

func (m *datasourceSecurityOutboundPoliciesSourcesModel) flattenSecurityOutboundPoliciesSources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityOutboundPoliciesSourcesModel {
	if input == nil {
		return &datasourceSecurityOutboundPoliciesSourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityOutboundPoliciesSourcesModel{}
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

func (s *datasourceSecurityOutboundPoliciesModel) flattenSecurityOutboundPoliciesSourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityOutboundPoliciesSourcesModel {
	if o == nil {
		return []datasourceSecurityOutboundPoliciesSourcesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityOutboundPoliciesSourcesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityOutboundPoliciesSourcesModel{}
	}

	values := make([]datasourceSecurityOutboundPoliciesSourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityOutboundPoliciesSourcesModel
		values[i] = *m.flattenSecurityOutboundPoliciesSources(ctx, ele, diags)
	}

	return values
}
