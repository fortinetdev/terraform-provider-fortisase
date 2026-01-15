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
var _ datasource.DataSource = &datasourceSecurityInternalPolicies{}

func newDatasourceSecurityInternalPolicies() datasource.DataSource {
	return &datasourceSecurityInternalPolicies{}
}

type datasourceSecurityInternalPolicies struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityInternalPoliciesModel describes the datasource data model.
type datasourceSecurityInternalPoliciesModel struct {
	PrimaryKey          types.String                                          `tfsdk:"primary_key"`
	Enabled             types.Bool                                            `tfsdk:"enabled"`
	Scope               types.String                                          `tfsdk:"scope"`
	Users               []datasourceSecurityInternalPoliciesUsersModel        `tfsdk:"users"`
	Destinations        []datasourceSecurityInternalPoliciesDestinationsModel `tfsdk:"destinations"`
	Services            []datasourceSecurityInternalPoliciesServicesModel     `tfsdk:"services"`
	Action              types.String                                          `tfsdk:"action"`
	Schedule            *datasourceSecurityInternalPoliciesScheduleModel      `tfsdk:"schedule"`
	Comments            types.String                                          `tfsdk:"comments"`
	ProfileGroup        *datasourceSecurityInternalPoliciesProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic          types.String                                          `tfsdk:"log_traffic"`
	Sources             []datasourceSecurityInternalPoliciesSourcesModel      `tfsdk:"sources"`
	CaptivePortalExempt types.Bool                                            `tfsdk:"captive_portal_exempt"`
}

func (r *datasourceSecurityInternalPolicies) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_policies"
}

func (r *datasourceSecurityInternalPolicies) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
								stringvalidator.OneOf("auth/users", "auth/user-groups"),
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
								stringvalidator.OneOf("network/hosts", "network/host-groups", "security/ip-threat-feeds"),
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

func (r *datasourceSecurityInternalPolicies) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_internal_policies"
}

func (r *datasourceSecurityInternalPolicies) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityInternalPoliciesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityInternalPoliciesModel) refreshSecurityInternalPolicies(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.Users = m.flattenSecurityInternalPoliciesUsersList(ctx, v, &diags)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityInternalPoliciesDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityInternalPoliciesServicesList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityInternalPoliciesSchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityInternalPoliciesProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityInternalPoliciesSourcesList(ctx, v, &diags)
	}

	if v, ok := o["captivePortalExempt"]; ok {
		m.CaptivePortalExempt = parseBoolValue(v)
	}

	return diags
}

func (data *datasourceSecurityInternalPoliciesModel) getURLObjectSecurityInternalPolicies(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityInternalPoliciesUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalPoliciesDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalPoliciesServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalPoliciesScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalPoliciesProfileGroupModel struct {
	Group               *datasourceSecurityInternalPoliciesProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                `tfsdk:"force_cert_inspection"`
}

type datasourceSecurityInternalPoliciesProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalPoliciesSourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityInternalPoliciesUsersModel) flattenSecurityInternalPoliciesUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPoliciesUsersModel {
	if input == nil {
		return &datasourceSecurityInternalPoliciesUsersModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPoliciesUsersModel{}
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

func (s *datasourceSecurityInternalPoliciesModel) flattenSecurityInternalPoliciesUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalPoliciesUsersModel {
	if o == nil {
		return []datasourceSecurityInternalPoliciesUsersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []datasourceSecurityInternalPoliciesUsersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalPoliciesUsersModel{}
	}

	values := make([]datasourceSecurityInternalPoliciesUsersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalPoliciesUsersModel
		values[i] = *m.flattenSecurityInternalPoliciesUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalPoliciesDestinationsModel) flattenSecurityInternalPoliciesDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPoliciesDestinationsModel {
	if input == nil {
		return &datasourceSecurityInternalPoliciesDestinationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPoliciesDestinationsModel{}
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

func (s *datasourceSecurityInternalPoliciesModel) flattenSecurityInternalPoliciesDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalPoliciesDestinationsModel {
	if o == nil {
		return []datasourceSecurityInternalPoliciesDestinationsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []datasourceSecurityInternalPoliciesDestinationsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalPoliciesDestinationsModel{}
	}

	values := make([]datasourceSecurityInternalPoliciesDestinationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalPoliciesDestinationsModel
		values[i] = *m.flattenSecurityInternalPoliciesDestinations(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalPoliciesServicesModel) flattenSecurityInternalPoliciesServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPoliciesServicesModel {
	if input == nil {
		return &datasourceSecurityInternalPoliciesServicesModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPoliciesServicesModel{}
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

func (s *datasourceSecurityInternalPoliciesModel) flattenSecurityInternalPoliciesServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalPoliciesServicesModel {
	if o == nil {
		return []datasourceSecurityInternalPoliciesServicesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []datasourceSecurityInternalPoliciesServicesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalPoliciesServicesModel{}
	}

	values := make([]datasourceSecurityInternalPoliciesServicesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalPoliciesServicesModel
		values[i] = *m.flattenSecurityInternalPoliciesServices(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalPoliciesScheduleModel) flattenSecurityInternalPoliciesSchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPoliciesScheduleModel {
	if input == nil {
		return &datasourceSecurityInternalPoliciesScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPoliciesScheduleModel{}
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

func (m *datasourceSecurityInternalPoliciesProfileGroupModel) flattenSecurityInternalPoliciesProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPoliciesProfileGroupModel {
	if input == nil {
		return &datasourceSecurityInternalPoliciesProfileGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPoliciesProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityInternalPoliciesProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *datasourceSecurityInternalPoliciesProfileGroupGroupModel) flattenSecurityInternalPoliciesProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPoliciesProfileGroupGroupModel {
	if input == nil {
		return &datasourceSecurityInternalPoliciesProfileGroupGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPoliciesProfileGroupGroupModel{}
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

func (m *datasourceSecurityInternalPoliciesSourcesModel) flattenSecurityInternalPoliciesSources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalPoliciesSourcesModel {
	if input == nil {
		return &datasourceSecurityInternalPoliciesSourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalPoliciesSourcesModel{}
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

func (s *datasourceSecurityInternalPoliciesModel) flattenSecurityInternalPoliciesSourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalPoliciesSourcesModel {
	if o == nil {
		return []datasourceSecurityInternalPoliciesSourcesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityInternalPoliciesSourcesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalPoliciesSourcesModel{}
	}

	values := make([]datasourceSecurityInternalPoliciesSourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalPoliciesSourcesModel
		values[i] = *m.flattenSecurityInternalPoliciesSources(ctx, ele, diags)
	}

	return values
}
