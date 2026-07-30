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
var _ datasource.DataSource = &datasourceSecurityInternalReversePolicy{}

func newDatasourceSecurityInternalReversePolicy() datasource.DataSource {
	return &datasourceSecurityInternalReversePolicy{}
}

type datasourceSecurityInternalReversePolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityInternalReversePolicyModel describes the datasource data model.
type datasourceSecurityInternalReversePolicyModel struct {
	PrimaryKey   types.String                                               `tfsdk:"primary_key"`
	Enabled      types.Bool                                                 `tfsdk:"enabled"`
	Scope        types.String                                               `tfsdk:"scope"`
	Sources      []datasourceSecurityInternalReversePolicySourcesModel      `tfsdk:"sources"`
	Services     []datasourceSecurityInternalReversePolicyServicesModel     `tfsdk:"services"`
	Action       types.String                                               `tfsdk:"action"`
	Schedule     *datasourceSecurityInternalReversePolicyScheduleModel      `tfsdk:"schedule"`
	Comments     types.String                                               `tfsdk:"comments"`
	ProfileGroup *datasourceSecurityInternalReversePolicyProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic   types.String                                               `tfsdk:"log_traffic"`
	Destinations []datasourceSecurityInternalReversePolicyDestinationsModel `tfsdk:"destinations"`
}

func (r *datasourceSecurityInternalReversePolicy) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_reverse_policy"
}

func (r *datasourceSecurityInternalReversePolicy) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Internal Reverse Policy Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.OneOf("all", "vpn-user", "thin-edge", "specify"),
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
			"sources": schema.ListNestedAttribute{
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
			"destinations": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "infra/ssids", "infra/fortigates", "infra/extenders"),
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

func (r *datasourceSecurityInternalReversePolicy) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_internal_reverse_policy"
}

func (r *datasourceSecurityInternalReversePolicy) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityInternalReversePolicyModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalReversePolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalReversePolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityInternalReversePolicyModel) refreshSecurityInternalReversePolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityInternalReversePolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityInternalReversePolicyServicesList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityInternalReversePolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityInternalReversePolicyProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityInternalReversePolicyDestinationsList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityInternalReversePolicyModel) getURLObjectSecurityInternalReversePolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityInternalReversePolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalReversePolicyServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalReversePolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalReversePolicyProfileGroupModel struct {
	Group               *datasourceSecurityInternalReversePolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                     `tfsdk:"force_cert_inspection"`
}

type datasourceSecurityInternalReversePolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityInternalReversePolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityInternalReversePolicySourcesModel) flattenSecurityInternalReversePolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePolicySourcesModel {
	if input == nil {
		return &datasourceSecurityInternalReversePolicySourcesModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePolicySourcesModel{}
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

func (s *datasourceSecurityInternalReversePolicyModel) flattenSecurityInternalReversePolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalReversePolicySourcesModel {
	if o == nil {
		return []datasourceSecurityInternalReversePolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []datasourceSecurityInternalReversePolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalReversePolicySourcesModel{}
	}

	values := make([]datasourceSecurityInternalReversePolicySourcesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalReversePolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityInternalReversePolicySources(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalReversePolicyServicesModel) flattenSecurityInternalReversePolicyServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePolicyServicesModel {
	if input == nil {
		return &datasourceSecurityInternalReversePolicyServicesModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePolicyServicesModel{}
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

func (s *datasourceSecurityInternalReversePolicyModel) flattenSecurityInternalReversePolicyServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalReversePolicyServicesModel {
	if o == nil {
		return []datasourceSecurityInternalReversePolicyServicesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []datasourceSecurityInternalReversePolicyServicesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalReversePolicyServicesModel{}
	}

	values := make([]datasourceSecurityInternalReversePolicyServicesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalReversePolicyServicesModel
		if i < len(s.Services) {
			m = s.Services[i]
		}
		values[i] = *m.flattenSecurityInternalReversePolicyServices(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityInternalReversePolicyScheduleModel) flattenSecurityInternalReversePolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePolicyScheduleModel {
	if input == nil {
		return &datasourceSecurityInternalReversePolicyScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePolicyScheduleModel{}
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

func (m *datasourceSecurityInternalReversePolicyProfileGroupModel) flattenSecurityInternalReversePolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePolicyProfileGroupModel {
	if input == nil {
		return &datasourceSecurityInternalReversePolicyProfileGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePolicyProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityInternalReversePolicyProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *datasourceSecurityInternalReversePolicyProfileGroupGroupModel) flattenSecurityInternalReversePolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePolicyProfileGroupGroupModel {
	if input == nil {
		return &datasourceSecurityInternalReversePolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePolicyProfileGroupGroupModel{}
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

func (m *datasourceSecurityInternalReversePolicyDestinationsModel) flattenSecurityInternalReversePolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityInternalReversePolicyDestinationsModel {
	if input == nil {
		return &datasourceSecurityInternalReversePolicyDestinationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityInternalReversePolicyDestinationsModel{}
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

func (s *datasourceSecurityInternalReversePolicyModel) flattenSecurityInternalReversePolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityInternalReversePolicyDestinationsModel {
	if o == nil {
		return []datasourceSecurityInternalReversePolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []datasourceSecurityInternalReversePolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityInternalReversePolicyDestinationsModel{}
	}

	values := make([]datasourceSecurityInternalReversePolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityInternalReversePolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityInternalReversePolicyDestinations(ctx, ele, diags)
	}

	return values
}
