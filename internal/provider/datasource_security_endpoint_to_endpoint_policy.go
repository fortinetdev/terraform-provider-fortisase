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
var _ datasource.DataSource = &datasourceSecurityEndpointToEndpointPolicy{}

func newDatasourceSecurityEndpointToEndpointPolicy() datasource.DataSource {
	return &datasourceSecurityEndpointToEndpointPolicy{}
}

type datasourceSecurityEndpointToEndpointPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityEndpointToEndpointPolicyModel describes the datasource data model.
type datasourceSecurityEndpointToEndpointPolicyModel struct {
	PrimaryKey   types.String                                                 `tfsdk:"primary_key"`
	Enabled      types.Bool                                                   `tfsdk:"enabled"`
	Users        []datasourceSecurityEndpointToEndpointPolicyUsersModel       `tfsdk:"users"`
	Sources      []datasourceSecurityEndpointToEndpointPolicySourcesModel     `tfsdk:"sources"`
	Services     []datasourceSecurityEndpointToEndpointPolicyServicesModel    `tfsdk:"services"`
	Action       types.String                                                 `tfsdk:"action"`
	Schedule     *datasourceSecurityEndpointToEndpointPolicyScheduleModel     `tfsdk:"schedule"`
	Comments     types.String                                                 `tfsdk:"comments"`
	ProfileGroup *datasourceSecurityEndpointToEndpointPolicyProfileGroupModel `tfsdk:"profile_group"`
	LogTraffic   types.String                                                 `tfsdk:"log_traffic"`
}

func (r *datasourceSecurityEndpointToEndpointPolicy) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_endpoint_to_endpoint_policy"
}

func (r *datasourceSecurityEndpointToEndpointPolicy) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Endpoint to Endpoint Policy Resource API V2 for FortiSASE.",
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
			"sources": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("endpoint/ztna-tags", "endpoint/ztna-tag-rules"),
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
		},
	}
}

func (r *datasourceSecurityEndpointToEndpointPolicy) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_endpoint_to_endpoint_policy"
}

func (r *datasourceSecurityEndpointToEndpointPolicy) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityEndpointToEndpointPolicyModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityEndpointToEndpointPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityEndpointToEndpointPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityEndpointToEndpointPolicyModel) refreshSecurityEndpointToEndpointPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["users"]; ok {
		m.Users = m.flattenSecurityEndpointToEndpointPolicyUsersList(ctx, v, &diags)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityEndpointToEndpointPolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityEndpointToEndpointPolicyServicesList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityEndpointToEndpointPolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityEndpointToEndpointPolicyProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	return diags
}

func (data *datasourceSecurityEndpointToEndpointPolicyModel) getURLObjectSecurityEndpointToEndpointPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityEndpointToEndpointPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityEndpointToEndpointPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityEndpointToEndpointPolicyServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityEndpointToEndpointPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityEndpointToEndpointPolicyProfileGroupModel struct {
	Group               *datasourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                        `tfsdk:"force_cert_inspection"`
}

type datasourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityEndpointToEndpointPolicyUsersModel) flattenSecurityEndpointToEndpointPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPolicyUsersModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPolicyUsersModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPolicyUsersModel{}
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

func (s *datasourceSecurityEndpointToEndpointPolicyModel) flattenSecurityEndpointToEndpointPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityEndpointToEndpointPolicyUsersModel {
	if o == nil {
		return []datasourceSecurityEndpointToEndpointPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []datasourceSecurityEndpointToEndpointPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityEndpointToEndpointPolicyUsersModel{}
	}

	values := make([]datasourceSecurityEndpointToEndpointPolicyUsersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityEndpointToEndpointPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityEndpointToEndpointPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityEndpointToEndpointPolicySourcesModel) flattenSecurityEndpointToEndpointPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPolicySourcesModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPolicySourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPolicySourcesModel{}
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

func (s *datasourceSecurityEndpointToEndpointPolicyModel) flattenSecurityEndpointToEndpointPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityEndpointToEndpointPolicySourcesModel {
	if o == nil {
		return []datasourceSecurityEndpointToEndpointPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityEndpointToEndpointPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityEndpointToEndpointPolicySourcesModel{}
	}

	values := make([]datasourceSecurityEndpointToEndpointPolicySourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityEndpointToEndpointPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityEndpointToEndpointPolicySources(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityEndpointToEndpointPolicyServicesModel) flattenSecurityEndpointToEndpointPolicyServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPolicyServicesModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPolicyServicesModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPolicyServicesModel{}
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

func (s *datasourceSecurityEndpointToEndpointPolicyModel) flattenSecurityEndpointToEndpointPolicyServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityEndpointToEndpointPolicyServicesModel {
	if o == nil {
		return []datasourceSecurityEndpointToEndpointPolicyServicesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []datasourceSecurityEndpointToEndpointPolicyServicesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityEndpointToEndpointPolicyServicesModel{}
	}

	values := make([]datasourceSecurityEndpointToEndpointPolicyServicesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityEndpointToEndpointPolicyServicesModel
		if i < len(s.Services) {
			m = s.Services[i]
		}
		values[i] = *m.flattenSecurityEndpointToEndpointPolicyServices(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityEndpointToEndpointPolicyScheduleModel) flattenSecurityEndpointToEndpointPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPolicyScheduleModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPolicyScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPolicyScheduleModel{}
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

func (m *datasourceSecurityEndpointToEndpointPolicyProfileGroupModel) flattenSecurityEndpointToEndpointPolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPolicyProfileGroupModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPolicyProfileGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPolicyProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityEndpointToEndpointPolicyProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *datasourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel) flattenSecurityEndpointToEndpointPolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel {
	if input == nil {
		return &datasourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel{}
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
