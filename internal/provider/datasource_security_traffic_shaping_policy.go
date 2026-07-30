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
var _ datasource.DataSource = &datasourceSecurityTrafficShapingPolicy{}

func newDatasourceSecurityTrafficShapingPolicy() datasource.DataSource {
	return &datasourceSecurityTrafficShapingPolicy{}
}

type datasourceSecurityTrafficShapingPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityTrafficShapingPolicyModel describes the datasource data model.
type datasourceSecurityTrafficShapingPolicyModel struct {
	PrimaryKey           types.String                                                     `tfsdk:"primary_key"`
	Comment              types.String                                                     `tfsdk:"comment"`
	Status               types.String                                                     `tfsdk:"status"`
	TrafficDirection     types.String                                                     `tfsdk:"traffic_direction"`
	Scope                types.String                                                     `tfsdk:"scope"`
	Sources              []datasourceSecurityTrafficShapingPolicySourcesModel             `tfsdk:"sources"`
	DestinationScope     types.String                                                     `tfsdk:"destination_scope"`
	Destinations         []datasourceSecurityTrafficShapingPolicyDestinationsModel        `tfsdk:"destinations"`
	Users                []datasourceSecurityTrafficShapingPolicyUsersModel               `tfsdk:"users"`
	Schedule             *datasourceSecurityTrafficShapingPolicyScheduleModel             `tfsdk:"schedule"`
	Services             []datasourceSecurityTrafficShapingPolicyServicesModel            `tfsdk:"services"`
	Applications         []datasourceSecurityTrafficShapingPolicyApplicationsModel        `tfsdk:"applications"`
	AppCategories        []datasourceSecurityTrafficShapingPolicyAppCategoriesModel       `tfsdk:"app_categories"`
	UrlCategories        []datasourceSecurityTrafficShapingPolicyUrlCategoriesModel       `tfsdk:"url_categories"`
	TrafficShaper        *datasourceSecurityTrafficShapingPolicyTrafficShaperModel        `tfsdk:"traffic_shaper"`
	TrafficShaperReverse *datasourceSecurityTrafficShapingPolicyTrafficShaperReverseModel `tfsdk:"traffic_shaper_reverse"`
	PerIpShaper          *datasourceSecurityTrafficShapingPolicyPerIpShaperModel          `tfsdk:"per_ip_shaper"`
}

func (r *datasourceSecurityTrafficShapingPolicy) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_traffic_shaping_policy"
}

func (r *datasourceSecurityTrafficShapingPolicy) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Traffic Shaping Policy Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 35),
				},
				Required: true,
			},
			"comment": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(255),
				},
				Computed: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"traffic_direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("outbound", "internal", "internal-reverse"),
				},
				Computed: true,
			},
			"scope": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("all", "vpn-user", "specify"),
				},
				Computed: true,
			},
			"destination_scope": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("all", "vpn-user", "specify"),
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
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "endpoint/ztna-tags", "security/ip-threat-feeds", "infra/ssids", "infra/fortigates", "infra/extenders"),
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
			"applications": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("security/applications"),
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"app_categories": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("security/application-categories"),
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"url_categories": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Float64Attribute{
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"traffic_shaper": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("security/traffic-shapers"),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"traffic_shaper_reverse": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("security/traffic-shapers"),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"per_ip_shaper": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("security/per-ip-traffic-shapers"),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceSecurityTrafficShapingPolicy) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_traffic_shaping_policy"
}

func (r *datasourceSecurityTrafficShapingPolicy) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityTrafficShapingPolicyModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityTrafficShapingPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityTrafficShapingPolicy(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityTrafficShapingPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityTrafficShapingPolicyModel) refreshSecurityTrafficShapingPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["comment"]; ok {
		m.Comment = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["trafficDirection"]; ok {
		m.TrafficDirection = parseStringValue(v)
	}

	if v, ok := o["scope"]; ok {
		m.Scope = parseStringValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityTrafficShapingPolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["destinationScope"]; ok {
		m.DestinationScope = parseStringValue(v)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityTrafficShapingPolicyDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["users"]; ok {
		m.Users = m.flattenSecurityTrafficShapingPolicyUsersList(ctx, v, &diags)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityTrafficShapingPolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityTrafficShapingPolicyServicesList(ctx, v, &diags)
	}

	if v, ok := o["applications"]; ok {
		m.Applications = m.flattenSecurityTrafficShapingPolicyApplicationsList(ctx, v, &diags)
	}

	if v, ok := o["appCategories"]; ok {
		m.AppCategories = m.flattenSecurityTrafficShapingPolicyAppCategoriesList(ctx, v, &diags)
	}

	if v, ok := o["urlCategories"]; ok {
		m.UrlCategories = m.flattenSecurityTrafficShapingPolicyUrlCategoriesList(ctx, v, &diags)
	}

	if v, ok := o["trafficShaper"]; ok {
		m.TrafficShaper = m.TrafficShaper.flattenSecurityTrafficShapingPolicyTrafficShaper(ctx, v, &diags)
	}

	if v, ok := o["trafficShaperReverse"]; ok {
		m.TrafficShaperReverse = m.TrafficShaperReverse.flattenSecurityTrafficShapingPolicyTrafficShaperReverse(ctx, v, &diags)
	}

	if v, ok := o["perIpShaper"]; ok {
		m.PerIpShaper = m.PerIpShaper.flattenSecurityTrafficShapingPolicyPerIpShaper(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityTrafficShapingPolicyModel) getURLObjectSecurityTrafficShapingPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityTrafficShapingPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityTrafficShapingPolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityTrafficShapingPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityTrafficShapingPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityTrafficShapingPolicyServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityTrafficShapingPolicyApplicationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityTrafficShapingPolicyAppCategoriesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityTrafficShapingPolicyUrlCategoriesModel struct {
	Id types.Float64 `tfsdk:"id"`
}

type datasourceSecurityTrafficShapingPolicyTrafficShaperModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityTrafficShapingPolicyTrafficShaperReverseModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityTrafficShapingPolicyPerIpShaperModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityTrafficShapingPolicySourcesModel) flattenSecurityTrafficShapingPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicySourcesModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicySourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicySourcesModel{}
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

func (s *datasourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityTrafficShapingPolicySourcesModel {
	if o == nil {
		return []datasourceSecurityTrafficShapingPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityTrafficShapingPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityTrafficShapingPolicySourcesModel{}
	}

	values := make([]datasourceSecurityTrafficShapingPolicySourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityTrafficShapingPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicySources(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityTrafficShapingPolicyDestinationsModel) flattenSecurityTrafficShapingPolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicyDestinationsModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicyDestinationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicyDestinationsModel{}
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

func (s *datasourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityTrafficShapingPolicyDestinationsModel {
	if o == nil {
		return []datasourceSecurityTrafficShapingPolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []datasourceSecurityTrafficShapingPolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityTrafficShapingPolicyDestinationsModel{}
	}

	values := make([]datasourceSecurityTrafficShapingPolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityTrafficShapingPolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyDestinations(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityTrafficShapingPolicyUsersModel) flattenSecurityTrafficShapingPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicyUsersModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicyUsersModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicyUsersModel{}
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

func (s *datasourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityTrafficShapingPolicyUsersModel {
	if o == nil {
		return []datasourceSecurityTrafficShapingPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []datasourceSecurityTrafficShapingPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityTrafficShapingPolicyUsersModel{}
	}

	values := make([]datasourceSecurityTrafficShapingPolicyUsersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityTrafficShapingPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityTrafficShapingPolicyScheduleModel) flattenSecurityTrafficShapingPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicyScheduleModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicyScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicyScheduleModel{}
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

func (m *datasourceSecurityTrafficShapingPolicyServicesModel) flattenSecurityTrafficShapingPolicyServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicyServicesModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicyServicesModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicyServicesModel{}
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

func (s *datasourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityTrafficShapingPolicyServicesModel {
	if o == nil {
		return []datasourceSecurityTrafficShapingPolicyServicesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []datasourceSecurityTrafficShapingPolicyServicesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityTrafficShapingPolicyServicesModel{}
	}

	values := make([]datasourceSecurityTrafficShapingPolicyServicesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityTrafficShapingPolicyServicesModel
		if i < len(s.Services) {
			m = s.Services[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyServices(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityTrafficShapingPolicyApplicationsModel) flattenSecurityTrafficShapingPolicyApplications(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicyApplicationsModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicyApplicationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicyApplicationsModel{}
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

func (s *datasourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyApplicationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityTrafficShapingPolicyApplicationsModel {
	if o == nil {
		return []datasourceSecurityTrafficShapingPolicyApplicationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument applications is not type of []interface{}.", "")
		return []datasourceSecurityTrafficShapingPolicyApplicationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityTrafficShapingPolicyApplicationsModel{}
	}

	values := make([]datasourceSecurityTrafficShapingPolicyApplicationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityTrafficShapingPolicyApplicationsModel
		if i < len(s.Applications) {
			m = s.Applications[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyApplications(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityTrafficShapingPolicyAppCategoriesModel) flattenSecurityTrafficShapingPolicyAppCategories(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicyAppCategoriesModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicyAppCategoriesModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicyAppCategoriesModel{}
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

func (s *datasourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyAppCategoriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityTrafficShapingPolicyAppCategoriesModel {
	if o == nil {
		return []datasourceSecurityTrafficShapingPolicyAppCategoriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument app_categories is not type of []interface{}.", "")
		return []datasourceSecurityTrafficShapingPolicyAppCategoriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityTrafficShapingPolicyAppCategoriesModel{}
	}

	values := make([]datasourceSecurityTrafficShapingPolicyAppCategoriesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityTrafficShapingPolicyAppCategoriesModel
		if i < len(s.AppCategories) {
			m = s.AppCategories[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyAppCategories(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityTrafficShapingPolicyUrlCategoriesModel) flattenSecurityTrafficShapingPolicyUrlCategories(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicyUrlCategoriesModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicyUrlCategoriesModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicyUrlCategoriesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	return m
}

func (s *datasourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyUrlCategoriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityTrafficShapingPolicyUrlCategoriesModel {
	if o == nil {
		return []datasourceSecurityTrafficShapingPolicyUrlCategoriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument url_categories is not type of []interface{}.", "")
		return []datasourceSecurityTrafficShapingPolicyUrlCategoriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityTrafficShapingPolicyUrlCategoriesModel{}
	}

	values := make([]datasourceSecurityTrafficShapingPolicyUrlCategoriesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityTrafficShapingPolicyUrlCategoriesModel
		if i < len(s.UrlCategories) {
			m = s.UrlCategories[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyUrlCategories(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityTrafficShapingPolicyTrafficShaperModel) flattenSecurityTrafficShapingPolicyTrafficShaper(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicyTrafficShaperModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicyTrafficShaperModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicyTrafficShaperModel{}
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

func (m *datasourceSecurityTrafficShapingPolicyTrafficShaperReverseModel) flattenSecurityTrafficShapingPolicyTrafficShaperReverse(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicyTrafficShaperReverseModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicyTrafficShaperReverseModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicyTrafficShaperReverseModel{}
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

func (m *datasourceSecurityTrafficShapingPolicyPerIpShaperModel) flattenSecurityTrafficShapingPolicyPerIpShaper(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityTrafficShapingPolicyPerIpShaperModel {
	if input == nil {
		return &datasourceSecurityTrafficShapingPolicyPerIpShaperModel{}
	}
	if m == nil {
		m = &datasourceSecurityTrafficShapingPolicyPerIpShaperModel{}
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
