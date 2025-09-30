// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSecurityInternalPolicies{}

func newResourceSecurityInternalPolicies() resource.Resource {
	return &resourceSecurityInternalPolicies{}
}

type resourceSecurityInternalPolicies struct {
	fortiClient *FortiClient
}

// resourceSecurityInternalPoliciesModel describes the resource data model.
type resourceSecurityInternalPoliciesModel struct {
	ID                  types.String                                        `tfsdk:"id"`
	PrimaryKey          types.String                                        `tfsdk:"primary_key"`
	Enabled             types.Bool                                          `tfsdk:"enabled"`
	Scope               types.String                                        `tfsdk:"scope"`
	Users               []resourceSecurityInternalPoliciesUsersModel        `tfsdk:"users"`
	Destinations        []resourceSecurityInternalPoliciesDestinationsModel `tfsdk:"destinations"`
	Services            []resourceSecurityInternalPoliciesServicesModel     `tfsdk:"services"`
	Action              types.String                                        `tfsdk:"action"`
	Schedule            *resourceSecurityInternalPoliciesScheduleModel      `tfsdk:"schedule"`
	Comments            types.String                                        `tfsdk:"comments"`
	ProfileGroup        *resourceSecurityInternalPoliciesProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic          types.String                                        `tfsdk:"log_traffic"`
	Sources             []resourceSecurityInternalPoliciesSourcesModel      `tfsdk:"sources"`
	CaptivePortalExempt types.Bool                                          `tfsdk:"captive_portal_exempt"`
}

func (r *resourceSecurityInternalPolicies) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_policies"
}

func (r *resourceSecurityInternalPolicies) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
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

func (r *resourceSecurityInternalPolicies) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityInternalPolicies) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityInternalPoliciesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityInternalPolicies(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityInternalPolicies(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityInternalPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityInternalPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityInternalPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalPolicies) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityInternalPoliciesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityInternalPoliciesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityInternalPolicies(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityInternalPolicies(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityInternalPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityInternalPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityInternalPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalPolicies) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityInternalPoliciesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalPolicies(ctx, "delete", diags))

	err := c.DeleteSecurityInternalPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceSecurityInternalPolicies) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityInternalPoliciesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityInternalPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalPolicies) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityInternalPoliciesModel) refreshSecurityInternalPolicies(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceSecurityInternalPoliciesModel) getCreateObjectSecurityInternalPolicies(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.Scope.IsNull() {
		result["scope"] = data.Scope.ValueString()
	}

	result["users"] = data.expandSecurityInternalPoliciesUsersList(ctx, data.Users, diags)

	result["destinations"] = data.expandSecurityInternalPoliciesDestinationsList(ctx, data.Destinations, diags)

	result["services"] = data.expandSecurityInternalPoliciesServicesList(ctx, data.Services, diags)

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityInternalPoliciesSchedule(ctx, diags)
	}

	if !data.Comments.IsNull() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isZeroStruct(*data.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityInternalPoliciesProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	if len(data.Sources) > 0 {
		result["sources"] = data.expandSecurityInternalPoliciesSourcesList(ctx, data.Sources, diags)
	}

	if !data.CaptivePortalExempt.IsNull() {
		result["captivePortalExempt"] = data.CaptivePortalExempt.ValueBool()
	}

	return &result
}

func (data *resourceSecurityInternalPoliciesModel) getUpdateObjectSecurityInternalPolicies(ctx context.Context, state resourceSecurityInternalPoliciesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.Scope.IsNull() {
		result["scope"] = data.Scope.ValueString()
	}

	if len(data.Users) > 0 || !isSameStruct(data.Users, state.Users) {
		result["users"] = data.expandSecurityInternalPoliciesUsersList(ctx, data.Users, diags)
	}

	if len(data.Destinations) > 0 || !isSameStruct(data.Destinations, state.Destinations) {
		result["destinations"] = data.expandSecurityInternalPoliciesDestinationsList(ctx, data.Destinations, diags)
	}

	if len(data.Services) > 0 || !isSameStruct(data.Services, state.Services) {
		result["services"] = data.expandSecurityInternalPoliciesServicesList(ctx, data.Services, diags)
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Schedule != nil && !isSameStruct(data.Schedule, state.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityInternalPoliciesSchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.Equal(state.Comments) {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isSameStruct(data.ProfileGroup, state.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityInternalPoliciesProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.Equal(state.LogTraffic) {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	if len(data.Sources) > 0 || !isSameStruct(data.Sources, state.Sources) {
		result["sources"] = data.expandSecurityInternalPoliciesSourcesList(ctx, data.Sources, diags)
	}

	if !data.CaptivePortalExempt.IsNull() && !data.CaptivePortalExempt.Equal(state.CaptivePortalExempt) {
		result["captivePortalExempt"] = data.CaptivePortalExempt.ValueBool()
	}

	return &result
}

func (data *resourceSecurityInternalPoliciesModel) getURLObjectSecurityInternalPolicies(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityInternalPoliciesUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalPoliciesDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalPoliciesServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalPoliciesScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalPoliciesProfileGroupModel struct {
	Group               *resourceSecurityInternalPoliciesProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                              `tfsdk:"force_cert_inspection"`
}

type resourceSecurityInternalPoliciesProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalPoliciesSourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityInternalPoliciesUsersModel) flattenSecurityInternalPoliciesUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPoliciesUsersModel {
	if input == nil {
		return &resourceSecurityInternalPoliciesUsersModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPoliciesUsersModel{}
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

func (s *resourceSecurityInternalPoliciesModel) flattenSecurityInternalPoliciesUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalPoliciesUsersModel {
	if o == nil {
		return []resourceSecurityInternalPoliciesUsersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []resourceSecurityInternalPoliciesUsersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalPoliciesUsersModel{}
	}

	values := make([]resourceSecurityInternalPoliciesUsersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalPoliciesUsersModel
		values[i] = *m.flattenSecurityInternalPoliciesUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalPoliciesDestinationsModel) flattenSecurityInternalPoliciesDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPoliciesDestinationsModel {
	if input == nil {
		return &resourceSecurityInternalPoliciesDestinationsModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPoliciesDestinationsModel{}
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

func (s *resourceSecurityInternalPoliciesModel) flattenSecurityInternalPoliciesDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalPoliciesDestinationsModel {
	if o == nil {
		return []resourceSecurityInternalPoliciesDestinationsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []resourceSecurityInternalPoliciesDestinationsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalPoliciesDestinationsModel{}
	}

	values := make([]resourceSecurityInternalPoliciesDestinationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalPoliciesDestinationsModel
		values[i] = *m.flattenSecurityInternalPoliciesDestinations(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalPoliciesServicesModel) flattenSecurityInternalPoliciesServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPoliciesServicesModel {
	if input == nil {
		return &resourceSecurityInternalPoliciesServicesModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPoliciesServicesModel{}
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

func (s *resourceSecurityInternalPoliciesModel) flattenSecurityInternalPoliciesServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalPoliciesServicesModel {
	if o == nil {
		return []resourceSecurityInternalPoliciesServicesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []resourceSecurityInternalPoliciesServicesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalPoliciesServicesModel{}
	}

	values := make([]resourceSecurityInternalPoliciesServicesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalPoliciesServicesModel
		values[i] = *m.flattenSecurityInternalPoliciesServices(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalPoliciesScheduleModel) flattenSecurityInternalPoliciesSchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPoliciesScheduleModel {
	if input == nil {
		return &resourceSecurityInternalPoliciesScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPoliciesScheduleModel{}
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

func (m *resourceSecurityInternalPoliciesProfileGroupModel) flattenSecurityInternalPoliciesProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPoliciesProfileGroupModel {
	if input == nil {
		return &resourceSecurityInternalPoliciesProfileGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPoliciesProfileGroupModel{}
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

func (m *resourceSecurityInternalPoliciesProfileGroupGroupModel) flattenSecurityInternalPoliciesProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPoliciesProfileGroupGroupModel {
	if input == nil {
		return &resourceSecurityInternalPoliciesProfileGroupGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPoliciesProfileGroupGroupModel{}
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

func (m *resourceSecurityInternalPoliciesSourcesModel) flattenSecurityInternalPoliciesSources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPoliciesSourcesModel {
	if input == nil {
		return &resourceSecurityInternalPoliciesSourcesModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPoliciesSourcesModel{}
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

func (s *resourceSecurityInternalPoliciesModel) flattenSecurityInternalPoliciesSourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalPoliciesSourcesModel {
	if o == nil {
		return []resourceSecurityInternalPoliciesSourcesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []resourceSecurityInternalPoliciesSourcesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalPoliciesSourcesModel{}
	}

	values := make([]resourceSecurityInternalPoliciesSourcesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalPoliciesSourcesModel
		values[i] = *m.flattenSecurityInternalPoliciesSources(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityInternalPoliciesUsersModel) expandSecurityInternalPoliciesUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalPoliciesModel) expandSecurityInternalPoliciesUsersList(ctx context.Context, l []resourceSecurityInternalPoliciesUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalPoliciesUsers(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalPoliciesDestinationsModel) expandSecurityInternalPoliciesDestinations(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalPoliciesModel) expandSecurityInternalPoliciesDestinationsList(ctx context.Context, l []resourceSecurityInternalPoliciesDestinationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalPoliciesDestinations(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalPoliciesServicesModel) expandSecurityInternalPoliciesServices(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalPoliciesModel) expandSecurityInternalPoliciesServicesList(ctx context.Context, l []resourceSecurityInternalPoliciesServicesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalPoliciesServices(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalPoliciesScheduleModel) expandSecurityInternalPoliciesSchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityInternalPoliciesProfileGroupModel) expandSecurityInternalPoliciesProfileGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandSecurityInternalPoliciesProfileGroupGroup(ctx, diags)
	}

	if !data.ForceCertInspection.IsNull() {
		result["forceCertInspection"] = data.ForceCertInspection.ValueBool()
	}

	return result
}

func (data *resourceSecurityInternalPoliciesProfileGroupGroupModel) expandSecurityInternalPoliciesProfileGroupGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityInternalPoliciesSourcesModel) expandSecurityInternalPoliciesSources(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalPoliciesModel) expandSecurityInternalPoliciesSourcesList(ctx context.Context, l []resourceSecurityInternalPoliciesSourcesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalPoliciesSources(ctx, diags)
	}
	return result
}
