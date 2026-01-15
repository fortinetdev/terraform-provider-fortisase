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
var _ datasource.DataSource = &datasourceSecurityEndpointToEndpointPolicies{}

func newDatasourceSecurityEndpointToEndpointPolicies() datasource.DataSource {
	return &datasourceSecurityEndpointToEndpointPolicies{}
}

type datasourceSecurityEndpointToEndpointPolicies struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityEndpointToEndpointPoliciesModel describes the datasource data model.
type datasourceSecurityEndpointToEndpointPoliciesModel struct {
	PrimaryKey   types.String                                                   `tfsdk:"primary_key"`
	Enabled      types.Bool                                                     `tfsdk:"enabled"`
	Users        []datasourceSecurityEndpointToEndpointPoliciesUsersModel       `tfsdk:"users"`
	Sources      []datasourceSecurityEndpointToEndpointPoliciesSourcesModel     `tfsdk:"sources"`
	Services     []datasourceSecurityEndpointToEndpointPoliciesServicesModel    `tfsdk:"services"`
	Action       types.String                                                   `tfsdk:"action"`
	Schedule     *datasourceSecurityEndpointToEndpointPoliciesScheduleModel     `tfsdk:"schedule"`
	Comments     types.String                                                   `tfsdk:"comments"`
	ProfileGroup *datasourceSecurityEndpointToEndpointPoliciesProfileGroupModel `tfsdk:"profile_group"`
	LogTraffic   types.String                                                   `tfsdk:"log_traffic"`
}

func (r *datasourceSecurityEndpointToEndpointPolicies) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_endpoint_to_endpoint_policies"
}

func (r *datasourceSecurityEndpointToEndpointPolicies) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"sources": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("endpoint/ztna-tags"),
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
		},
	}
}

func (r *datasourceSecurityEndpointToEndpointPolicies) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_endpoint_to_endpoint_policies"
}

func (r *datasourceSecurityEndpointToEndpointPolicies) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityEndpointToEndpointPoliciesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityEndpointToEndpointPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityEndpointToEndpointPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityEndpointToEndpointPoliciesModel) refreshSecurityEndpointToEndpointPolicies(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["users"]; ok {
		m.Users = m.flattenSecurityEndpointToEndpointPoliciesUsersList(ctx, v, &diags)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityEndpointToEndpointPoliciesSourcesList(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityEndpointToEndpointPoliciesServicesList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityEndpointToEndpointPoliciesSchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityEndpointToEndpointPoliciesProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	return diags
}

func (data *datasourceSecurityEndpointToEndpointPoliciesModel) getURLObjectSecurityEndpointToEndpointPolicies(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityEndpointToEndpointPoliciesUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityEndpointToEndpointPoliciesSourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityEndpointToEndpointPoliciesServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityEndpointToEndpointPoliciesScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityEndpointToEndpointPoliciesProfileGroupModel struct {
	Group               *datasourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                          `tfsdk:"force_cert_inspection"`
}

type datasourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityEndpointToEndpointPoliciesUsersModel) flattenSecurityEndpointToEndpointPoliciesUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPoliciesUsersModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPoliciesUsersModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPoliciesUsersModel{}
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

func (s *datasourceSecurityEndpointToEndpointPoliciesModel) flattenSecurityEndpointToEndpointPoliciesUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityEndpointToEndpointPoliciesUsersModel {
	if o == nil {
		return []datasourceSecurityEndpointToEndpointPoliciesUsersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []datasourceSecurityEndpointToEndpointPoliciesUsersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityEndpointToEndpointPoliciesUsersModel{}
	}

	values := make([]datasourceSecurityEndpointToEndpointPoliciesUsersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityEndpointToEndpointPoliciesUsersModel
		values[i] = *m.flattenSecurityEndpointToEndpointPoliciesUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityEndpointToEndpointPoliciesSourcesModel) flattenSecurityEndpointToEndpointPoliciesSources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPoliciesSourcesModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPoliciesSourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPoliciesSourcesModel{}
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

func (s *datasourceSecurityEndpointToEndpointPoliciesModel) flattenSecurityEndpointToEndpointPoliciesSourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityEndpointToEndpointPoliciesSourcesModel {
	if o == nil {
		return []datasourceSecurityEndpointToEndpointPoliciesSourcesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityEndpointToEndpointPoliciesSourcesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityEndpointToEndpointPoliciesSourcesModel{}
	}

	values := make([]datasourceSecurityEndpointToEndpointPoliciesSourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityEndpointToEndpointPoliciesSourcesModel
		values[i] = *m.flattenSecurityEndpointToEndpointPoliciesSources(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityEndpointToEndpointPoliciesServicesModel) flattenSecurityEndpointToEndpointPoliciesServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPoliciesServicesModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPoliciesServicesModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPoliciesServicesModel{}
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

func (s *datasourceSecurityEndpointToEndpointPoliciesModel) flattenSecurityEndpointToEndpointPoliciesServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityEndpointToEndpointPoliciesServicesModel {
	if o == nil {
		return []datasourceSecurityEndpointToEndpointPoliciesServicesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []datasourceSecurityEndpointToEndpointPoliciesServicesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityEndpointToEndpointPoliciesServicesModel{}
	}

	values := make([]datasourceSecurityEndpointToEndpointPoliciesServicesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityEndpointToEndpointPoliciesServicesModel
		values[i] = *m.flattenSecurityEndpointToEndpointPoliciesServices(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityEndpointToEndpointPoliciesScheduleModel) flattenSecurityEndpointToEndpointPoliciesSchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPoliciesScheduleModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPoliciesScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPoliciesScheduleModel{}
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

func (m *datasourceSecurityEndpointToEndpointPoliciesProfileGroupModel) flattenSecurityEndpointToEndpointPoliciesProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPoliciesProfileGroupModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPoliciesProfileGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPoliciesProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityEndpointToEndpointPoliciesProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *datasourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel) flattenSecurityEndpointToEndpointPoliciesProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel{}
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
