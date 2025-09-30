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
var _ resource.Resource = &resourceSecurityEndpointToEndpointPolicies{}

func newResourceSecurityEndpointToEndpointPolicies() resource.Resource {
	return &resourceSecurityEndpointToEndpointPolicies{}
}

type resourceSecurityEndpointToEndpointPolicies struct {
	fortiClient *FortiClient
}

// resourceSecurityEndpointToEndpointPoliciesModel describes the resource data model.
type resourceSecurityEndpointToEndpointPoliciesModel struct {
	ID           types.String                                                 `tfsdk:"id"`
	PrimaryKey   types.String                                                 `tfsdk:"primary_key"`
	Enabled      types.Bool                                                   `tfsdk:"enabled"`
	Users        []resourceSecurityEndpointToEndpointPoliciesUsersModel       `tfsdk:"users"`
	Sources      []resourceSecurityEndpointToEndpointPoliciesSourcesModel     `tfsdk:"sources"`
	Services     []resourceSecurityEndpointToEndpointPoliciesServicesModel    `tfsdk:"services"`
	Action       types.String                                                 `tfsdk:"action"`
	Schedule     *resourceSecurityEndpointToEndpointPoliciesScheduleModel     `tfsdk:"schedule"`
	Comments     types.String                                                 `tfsdk:"comments"`
	ProfileGroup *resourceSecurityEndpointToEndpointPoliciesProfileGroupModel `tfsdk:"profile_group"`
	LogTraffic   types.String                                                 `tfsdk:"log_traffic"`
}

func (r *resourceSecurityEndpointToEndpointPolicies) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_endpoint_to_endpoint_policies"
}

func (r *resourceSecurityEndpointToEndpointPolicies) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *resourceSecurityEndpointToEndpointPolicies) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityEndpointToEndpointPolicies) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityEndpointToEndpointPoliciesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityEndpointToEndpointPolicies(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicies(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityEndpointToEndpointPolicies(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityEndpointToEndpointPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityEndpointToEndpointPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityEndpointToEndpointPolicies) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityEndpointToEndpointPoliciesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityEndpointToEndpointPoliciesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityEndpointToEndpointPolicies(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicies(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityEndpointToEndpointPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityEndpointToEndpointPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityEndpointToEndpointPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityEndpointToEndpointPolicies) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityEndpointToEndpointPoliciesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicies(ctx, "delete", diags))

	err := c.DeleteSecurityEndpointToEndpointPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceSecurityEndpointToEndpointPolicies) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityEndpointToEndpointPoliciesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityEndpointToEndpointPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityEndpointToEndpointPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityEndpointToEndpointPolicies) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityEndpointToEndpointPoliciesModel) refreshSecurityEndpointToEndpointPolicies(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceSecurityEndpointToEndpointPoliciesModel) getCreateObjectSecurityEndpointToEndpointPolicies(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	result["users"] = data.expandSecurityEndpointToEndpointPoliciesUsersList(ctx, data.Users, diags)

	if len(data.Sources) > 0 {
		result["sources"] = data.expandSecurityEndpointToEndpointPoliciesSourcesList(ctx, data.Sources, diags)
	}

	result["services"] = data.expandSecurityEndpointToEndpointPoliciesServicesList(ctx, data.Services, diags)

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityEndpointToEndpointPoliciesSchedule(ctx, diags)
	}

	if !data.Comments.IsNull() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isZeroStruct(*data.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityEndpointToEndpointPoliciesProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	return &result
}

func (data *resourceSecurityEndpointToEndpointPoliciesModel) getUpdateObjectSecurityEndpointToEndpointPolicies(ctx context.Context, state resourceSecurityEndpointToEndpointPoliciesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if len(data.Users) > 0 || !isSameStruct(data.Users, state.Users) {
		result["users"] = data.expandSecurityEndpointToEndpointPoliciesUsersList(ctx, data.Users, diags)
	}

	if len(data.Sources) > 0 || !isSameStruct(data.Sources, state.Sources) {
		result["sources"] = data.expandSecurityEndpointToEndpointPoliciesSourcesList(ctx, data.Sources, diags)
	}

	if len(data.Services) > 0 || !isSameStruct(data.Services, state.Services) {
		result["services"] = data.expandSecurityEndpointToEndpointPoliciesServicesList(ctx, data.Services, diags)
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Schedule != nil && !isSameStruct(data.Schedule, state.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityEndpointToEndpointPoliciesSchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.Equal(state.Comments) {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isSameStruct(data.ProfileGroup, state.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityEndpointToEndpointPoliciesProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.Equal(state.LogTraffic) {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	return &result
}

func (data *resourceSecurityEndpointToEndpointPoliciesModel) getURLObjectSecurityEndpointToEndpointPolicies(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityEndpointToEndpointPoliciesUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityEndpointToEndpointPoliciesSourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityEndpointToEndpointPoliciesServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityEndpointToEndpointPoliciesScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityEndpointToEndpointPoliciesProfileGroupModel struct {
	Group               *resourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                        `tfsdk:"force_cert_inspection"`
}

type resourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityEndpointToEndpointPoliciesUsersModel) flattenSecurityEndpointToEndpointPoliciesUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPoliciesUsersModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPoliciesUsersModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPoliciesUsersModel{}
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

func (s *resourceSecurityEndpointToEndpointPoliciesModel) flattenSecurityEndpointToEndpointPoliciesUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityEndpointToEndpointPoliciesUsersModel {
	if o == nil {
		return []resourceSecurityEndpointToEndpointPoliciesUsersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []resourceSecurityEndpointToEndpointPoliciesUsersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityEndpointToEndpointPoliciesUsersModel{}
	}

	values := make([]resourceSecurityEndpointToEndpointPoliciesUsersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityEndpointToEndpointPoliciesUsersModel
		values[i] = *m.flattenSecurityEndpointToEndpointPoliciesUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityEndpointToEndpointPoliciesSourcesModel) flattenSecurityEndpointToEndpointPoliciesSources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPoliciesSourcesModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPoliciesSourcesModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPoliciesSourcesModel{}
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

func (s *resourceSecurityEndpointToEndpointPoliciesModel) flattenSecurityEndpointToEndpointPoliciesSourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityEndpointToEndpointPoliciesSourcesModel {
	if o == nil {
		return []resourceSecurityEndpointToEndpointPoliciesSourcesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []resourceSecurityEndpointToEndpointPoliciesSourcesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityEndpointToEndpointPoliciesSourcesModel{}
	}

	values := make([]resourceSecurityEndpointToEndpointPoliciesSourcesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityEndpointToEndpointPoliciesSourcesModel
		values[i] = *m.flattenSecurityEndpointToEndpointPoliciesSources(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityEndpointToEndpointPoliciesServicesModel) flattenSecurityEndpointToEndpointPoliciesServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPoliciesServicesModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPoliciesServicesModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPoliciesServicesModel{}
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

func (s *resourceSecurityEndpointToEndpointPoliciesModel) flattenSecurityEndpointToEndpointPoliciesServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityEndpointToEndpointPoliciesServicesModel {
	if o == nil {
		return []resourceSecurityEndpointToEndpointPoliciesServicesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []resourceSecurityEndpointToEndpointPoliciesServicesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityEndpointToEndpointPoliciesServicesModel{}
	}

	values := make([]resourceSecurityEndpointToEndpointPoliciesServicesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityEndpointToEndpointPoliciesServicesModel
		values[i] = *m.flattenSecurityEndpointToEndpointPoliciesServices(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityEndpointToEndpointPoliciesScheduleModel) flattenSecurityEndpointToEndpointPoliciesSchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPoliciesScheduleModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPoliciesScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPoliciesScheduleModel{}
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

func (m *resourceSecurityEndpointToEndpointPoliciesProfileGroupModel) flattenSecurityEndpointToEndpointPoliciesProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPoliciesProfileGroupModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPoliciesProfileGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPoliciesProfileGroupModel{}
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

func (m *resourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel) flattenSecurityEndpointToEndpointPoliciesProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel{}
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

func (data *resourceSecurityEndpointToEndpointPoliciesUsersModel) expandSecurityEndpointToEndpointPoliciesUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityEndpointToEndpointPoliciesModel) expandSecurityEndpointToEndpointPoliciesUsersList(ctx context.Context, l []resourceSecurityEndpointToEndpointPoliciesUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityEndpointToEndpointPoliciesUsers(ctx, diags)
	}
	return result
}

func (data *resourceSecurityEndpointToEndpointPoliciesSourcesModel) expandSecurityEndpointToEndpointPoliciesSources(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityEndpointToEndpointPoliciesModel) expandSecurityEndpointToEndpointPoliciesSourcesList(ctx context.Context, l []resourceSecurityEndpointToEndpointPoliciesSourcesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityEndpointToEndpointPoliciesSources(ctx, diags)
	}
	return result
}

func (data *resourceSecurityEndpointToEndpointPoliciesServicesModel) expandSecurityEndpointToEndpointPoliciesServices(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityEndpointToEndpointPoliciesModel) expandSecurityEndpointToEndpointPoliciesServicesList(ctx context.Context, l []resourceSecurityEndpointToEndpointPoliciesServicesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityEndpointToEndpointPoliciesServices(ctx, diags)
	}
	return result
}

func (data *resourceSecurityEndpointToEndpointPoliciesScheduleModel) expandSecurityEndpointToEndpointPoliciesSchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityEndpointToEndpointPoliciesProfileGroupModel) expandSecurityEndpointToEndpointPoliciesProfileGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandSecurityEndpointToEndpointPoliciesProfileGroupGroup(ctx, diags)
	}

	if !data.ForceCertInspection.IsNull() {
		result["forceCertInspection"] = data.ForceCertInspection.ValueBool()
	}

	return result
}

func (data *resourceSecurityEndpointToEndpointPoliciesProfileGroupGroupModel) expandSecurityEndpointToEndpointPoliciesProfileGroupGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
