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
var _ datasource.DataSource = &datasourceSecurityInternalReversePolicies{}

func newDatasourceSecurityInternalReversePolicies() datasource.DataSource {
	return &datasourceSecurityInternalReversePolicies{}
}

type datasourceSecurityInternalReversePolicies struct {
	fortiClient *FortiClient
}

// datasourceSecurityInternalReversePoliciesModel describes the datasource data model.
type datasourceSecurityInternalReversePoliciesModel struct {
	PrimaryKey   types.String                                                 `tfsdk:"primary_key"`
	Enabled      types.Bool                                                   `tfsdk:"enabled"`
	Scope        types.String                                                 `tfsdk:"scope"`
	Sources      []datasourceSecurityInternalReversePoliciesSourcesModel      `tfsdk:"sources"`
	Services     []datasourceSecurityInternalReversePoliciesServicesModel     `tfsdk:"services"`
	Action       types.String                                                 `tfsdk:"action"`
	Schedule     *datasourceSecurityInternalReversePoliciesScheduleModel      `tfsdk:"schedule"`
	Comments     types.String                                                 `tfsdk:"comments"`
	ProfileGroup *datasourceSecurityInternalReversePoliciesProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic   types.String                                                 `tfsdk:"log_traffic"`
	Destinations []datasourceSecurityInternalReversePoliciesDestinationsModel `tfsdk:"destinations"`
}

func (r *datasourceSecurityInternalReversePolicies) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_reverse_policies"
}

func (r *datasourceSecurityInternalReversePolicies) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"sources": schema.ListNestedAttribute{
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
			"destinations": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("network/hosts", "network/host-groups", "infra/ssids", "infra/fortigates", "infra/extenders"),
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

func (r *datasourceSecurityInternalReversePolicies) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityInternalReversePolicies) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityInternalReversePoliciesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalReversePolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityInternalReversePolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityInternalReversePoliciesModel) refreshSecurityInternalReversePolicies(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityInternalReversePoliciesSourcesList(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityInternalReversePoliciesServicesList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityInternalReversePoliciesSchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityInternalReversePoliciesProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityInternalReversePoliciesDestinationsList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityInternalReversePoliciesModel) getURLObjectSecurityInternalReversePolicies(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityInternalReversePoliciesSourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalReversePoliciesServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalReversePoliciesScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalReversePoliciesProfileGroupModel struct {
	Group               *datasourceSecurityInternalReversePoliciesProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                       `tfsdk:"force_cert_inspection"`
}

type datasourceSecurityInternalReversePoliciesProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalReversePoliciesDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityInternalReversePoliciesSourcesModel) flattenSecurityInternalReversePoliciesSources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePoliciesSourcesModel {
	if input == nil {
		return &datasourceSecurityInternalReversePoliciesSourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePoliciesSourcesModel{}
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

func (s *datasourceSecurityInternalReversePoliciesModel) flattenSecurityInternalReversePoliciesSourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalReversePoliciesSourcesModel {
	if o == nil {
		return []datasourceSecurityInternalReversePoliciesSourcesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityInternalReversePoliciesSourcesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalReversePoliciesSourcesModel{}
	}

	values := make([]datasourceSecurityInternalReversePoliciesSourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalReversePoliciesSourcesModel
		values[i] = *m.flattenSecurityInternalReversePoliciesSources(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalReversePoliciesServicesModel) flattenSecurityInternalReversePoliciesServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePoliciesServicesModel {
	if input == nil {
		return &datasourceSecurityInternalReversePoliciesServicesModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePoliciesServicesModel{}
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

func (s *datasourceSecurityInternalReversePoliciesModel) flattenSecurityInternalReversePoliciesServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalReversePoliciesServicesModel {
	if o == nil {
		return []datasourceSecurityInternalReversePoliciesServicesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []datasourceSecurityInternalReversePoliciesServicesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalReversePoliciesServicesModel{}
	}

	values := make([]datasourceSecurityInternalReversePoliciesServicesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalReversePoliciesServicesModel
		values[i] = *m.flattenSecurityInternalReversePoliciesServices(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalReversePoliciesScheduleModel) flattenSecurityInternalReversePoliciesSchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePoliciesScheduleModel {
	if input == nil {
		return &datasourceSecurityInternalReversePoliciesScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePoliciesScheduleModel{}
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

func (m *datasourceSecurityInternalReversePoliciesProfileGroupModel) flattenSecurityInternalReversePoliciesProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePoliciesProfileGroupModel {
	if input == nil {
		return &datasourceSecurityInternalReversePoliciesProfileGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePoliciesProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityInternalReversePoliciesProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *datasourceSecurityInternalReversePoliciesProfileGroupGroupModel) flattenSecurityInternalReversePoliciesProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePoliciesProfileGroupGroupModel {
	if input == nil {
		return &datasourceSecurityInternalReversePoliciesProfileGroupGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePoliciesProfileGroupGroupModel{}
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

func (m *datasourceSecurityInternalReversePoliciesDestinationsModel) flattenSecurityInternalReversePoliciesDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePoliciesDestinationsModel {
	if input == nil {
		return &datasourceSecurityInternalReversePoliciesDestinationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePoliciesDestinationsModel{}
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

func (s *datasourceSecurityInternalReversePoliciesModel) flattenSecurityInternalReversePoliciesDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalReversePoliciesDestinationsModel {
	if o == nil {
		return []datasourceSecurityInternalReversePoliciesDestinationsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []datasourceSecurityInternalReversePoliciesDestinationsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalReversePoliciesDestinationsModel{}
	}

	values := make([]datasourceSecurityInternalReversePoliciesDestinationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalReversePoliciesDestinationsModel
		values[i] = *m.flattenSecurityInternalReversePoliciesDestinations(ctx, ele, diags)
	}

	return values
}
