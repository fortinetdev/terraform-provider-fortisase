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
var _ resource.Resource = &resourceSecurityOutboundPolicies{}

func newResourceSecurityOutboundPolicies() resource.Resource {
	return &resourceSecurityOutboundPolicies{}
}

type resourceSecurityOutboundPolicies struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityOutboundPoliciesModel describes the resource data model.
type resourceSecurityOutboundPoliciesModel struct {
	ID                  types.String                                        `tfsdk:"id"`
	PrimaryKey          types.String                                        `tfsdk:"primary_key"`
	Enabled             types.Bool                                          `tfsdk:"enabled"`
	Scope               types.String                                        `tfsdk:"scope"`
	Users               []resourceSecurityOutboundPoliciesUsersModel        `tfsdk:"users"`
	Destinations        []resourceSecurityOutboundPoliciesDestinationsModel `tfsdk:"destinations"`
	Services            []resourceSecurityOutboundPoliciesServicesModel     `tfsdk:"services"`
	Action              types.String                                        `tfsdk:"action"`
	Schedule            *resourceSecurityOutboundPoliciesScheduleModel      `tfsdk:"schedule"`
	Comments            types.String                                        `tfsdk:"comments"`
	ProfileGroup        *resourceSecurityOutboundPoliciesProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic          types.String                                        `tfsdk:"log_traffic"`
	Sources             []resourceSecurityOutboundPoliciesSourcesModel      `tfsdk:"sources"`
	CaptivePortalExempt types.Bool                                          `tfsdk:"captive_portal_exempt"`
}

func (r *resourceSecurityOutboundPolicies) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_outbound_policies"
}

func (r *resourceSecurityOutboundPolicies) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *resourceSecurityOutboundPolicies) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_outbound_policies"
}

func (r *resourceSecurityOutboundPolicies) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityOutboundPoliciesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityOutboundPolicies(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicies(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityOutboundPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityOutboundPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityOutboundPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityOutboundPolicies) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityOutboundPoliciesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityOutboundPoliciesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityOutboundPolicies(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicies(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityOutboundPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityOutboundPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityOutboundPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityOutboundPolicies) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityOutboundPoliciesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicies(ctx, "delete", diags))

	output, err := c.DeleteSecurityOutboundPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityOutboundPolicies) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityOutboundPoliciesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityOutboundPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityOutboundPolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityOutboundPolicies) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityOutboundPoliciesModel) refreshSecurityOutboundPolicies(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.Users = m.flattenSecurityOutboundPoliciesUsersList(ctx, v, &diags)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityOutboundPoliciesDestinationsList(ctx, v, &diags)
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

func (data *resourceSecurityOutboundPoliciesModel) getCreateObjectSecurityOutboundPolicies(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
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

	result["users"] = data.expandSecurityOutboundPoliciesUsersList(ctx, data.Users, diags)

	result["destinations"] = data.expandSecurityOutboundPoliciesDestinationsList(ctx, data.Destinations, diags)

	result["services"] = data.expandSecurityOutboundPoliciesServicesList(ctx, data.Services, diags)

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityOutboundPoliciesSchedule(ctx, diags)
	}

	if !data.Comments.IsNull() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isZeroStruct(*data.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityOutboundPoliciesProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	result["sources"] = data.expandSecurityOutboundPoliciesSourcesList(ctx, data.Sources, diags)

	if !data.CaptivePortalExempt.IsNull() {
		result["captivePortalExempt"] = data.CaptivePortalExempt.ValueBool()
	}

	return &result
}

func (data *resourceSecurityOutboundPoliciesModel) getUpdateObjectSecurityOutboundPolicies(ctx context.Context, state resourceSecurityOutboundPoliciesModel, diags *diag.Diagnostics) *map[string]interface{} {
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

	if data.Users != nil {
		result["users"] = data.expandSecurityOutboundPoliciesUsersList(ctx, data.Users, diags)
	}

	if data.Destinations != nil {
		result["destinations"] = data.expandSecurityOutboundPoliciesDestinationsList(ctx, data.Destinations, diags)
	}

	if data.Services != nil {
		result["services"] = data.expandSecurityOutboundPoliciesServicesList(ctx, data.Services, diags)
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Schedule != nil {
		result["schedule"] = data.Schedule.expandSecurityOutboundPoliciesSchedule(ctx, diags)
	}

	if !data.Comments.IsNull() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil {
		result["profileGroup"] = data.ProfileGroup.expandSecurityOutboundPoliciesProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	if data.Sources != nil {
		result["sources"] = data.expandSecurityOutboundPoliciesSourcesList(ctx, data.Sources, diags)
	}

	if !data.CaptivePortalExempt.IsNull() {
		result["captivePortalExempt"] = data.CaptivePortalExempt.ValueBool()
	}

	return &result
}

func (data *resourceSecurityOutboundPoliciesModel) getURLObjectSecurityOutboundPolicies(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityOutboundPoliciesUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityOutboundPoliciesDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityOutboundPoliciesServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityOutboundPoliciesScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityOutboundPoliciesProfileGroupModel struct {
	Group               *resourceSecurityOutboundPoliciesProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                              `tfsdk:"force_cert_inspection"`
}

type resourceSecurityOutboundPoliciesProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityOutboundPoliciesSourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityOutboundPoliciesUsersModel) flattenSecurityOutboundPoliciesUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPoliciesUsersModel {
	if input == nil {
		return &resourceSecurityOutboundPoliciesUsersModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPoliciesUsersModel{}
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

func (s *resourceSecurityOutboundPoliciesModel) flattenSecurityOutboundPoliciesUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityOutboundPoliciesUsersModel {
	if o == nil {
		return []resourceSecurityOutboundPoliciesUsersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []resourceSecurityOutboundPoliciesUsersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityOutboundPoliciesUsersModel{}
	}

	values := make([]resourceSecurityOutboundPoliciesUsersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityOutboundPoliciesUsersModel
		values[i] = *m.flattenSecurityOutboundPoliciesUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityOutboundPoliciesDestinationsModel) flattenSecurityOutboundPoliciesDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPoliciesDestinationsModel {
	if input == nil {
		return &resourceSecurityOutboundPoliciesDestinationsModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPoliciesDestinationsModel{}
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

func (s *resourceSecurityOutboundPoliciesModel) flattenSecurityOutboundPoliciesDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityOutboundPoliciesDestinationsModel {
	if o == nil {
		return []resourceSecurityOutboundPoliciesDestinationsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []resourceSecurityOutboundPoliciesDestinationsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityOutboundPoliciesDestinationsModel{}
	}

	values := make([]resourceSecurityOutboundPoliciesDestinationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityOutboundPoliciesDestinationsModel
		values[i] = *m.flattenSecurityOutboundPoliciesDestinations(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityOutboundPoliciesServicesModel) flattenSecurityOutboundPoliciesServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPoliciesServicesModel {
	if input == nil {
		return &resourceSecurityOutboundPoliciesServicesModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPoliciesServicesModel{}
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

func (s *resourceSecurityOutboundPoliciesModel) flattenSecurityOutboundPoliciesServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityOutboundPoliciesServicesModel {
	if o == nil {
		return []resourceSecurityOutboundPoliciesServicesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []resourceSecurityOutboundPoliciesServicesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityOutboundPoliciesServicesModel{}
	}

	values := make([]resourceSecurityOutboundPoliciesServicesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityOutboundPoliciesServicesModel
		values[i] = *m.flattenSecurityOutboundPoliciesServices(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityOutboundPoliciesScheduleModel) flattenSecurityOutboundPoliciesSchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPoliciesScheduleModel {
	if input == nil {
		return &resourceSecurityOutboundPoliciesScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPoliciesScheduleModel{}
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

func (m *resourceSecurityOutboundPoliciesProfileGroupModel) flattenSecurityOutboundPoliciesProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPoliciesProfileGroupModel {
	if input == nil {
		return &resourceSecurityOutboundPoliciesProfileGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPoliciesProfileGroupModel{}
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

func (m *resourceSecurityOutboundPoliciesProfileGroupGroupModel) flattenSecurityOutboundPoliciesProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPoliciesProfileGroupGroupModel {
	if input == nil {
		return &resourceSecurityOutboundPoliciesProfileGroupGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPoliciesProfileGroupGroupModel{}
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

func (m *resourceSecurityOutboundPoliciesSourcesModel) flattenSecurityOutboundPoliciesSources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPoliciesSourcesModel {
	if input == nil {
		return &resourceSecurityOutboundPoliciesSourcesModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPoliciesSourcesModel{}
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

func (s *resourceSecurityOutboundPoliciesModel) flattenSecurityOutboundPoliciesSourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityOutboundPoliciesSourcesModel {
	if o == nil {
		return []resourceSecurityOutboundPoliciesSourcesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []resourceSecurityOutboundPoliciesSourcesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityOutboundPoliciesSourcesModel{}
	}

	values := make([]resourceSecurityOutboundPoliciesSourcesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityOutboundPoliciesSourcesModel
		values[i] = *m.flattenSecurityOutboundPoliciesSources(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityOutboundPoliciesUsersModel) expandSecurityOutboundPoliciesUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityOutboundPoliciesModel) expandSecurityOutboundPoliciesUsersList(ctx context.Context, l []resourceSecurityOutboundPoliciesUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityOutboundPoliciesUsers(ctx, diags)
	}
	return result
}

func (data *resourceSecurityOutboundPoliciesDestinationsModel) expandSecurityOutboundPoliciesDestinations(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityOutboundPoliciesModel) expandSecurityOutboundPoliciesDestinationsList(ctx context.Context, l []resourceSecurityOutboundPoliciesDestinationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityOutboundPoliciesDestinations(ctx, diags)
	}
	return result
}

func (data *resourceSecurityOutboundPoliciesServicesModel) expandSecurityOutboundPoliciesServices(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityOutboundPoliciesModel) expandSecurityOutboundPoliciesServicesList(ctx context.Context, l []resourceSecurityOutboundPoliciesServicesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityOutboundPoliciesServices(ctx, diags)
	}
	return result
}

func (data *resourceSecurityOutboundPoliciesScheduleModel) expandSecurityOutboundPoliciesSchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityOutboundPoliciesProfileGroupModel) expandSecurityOutboundPoliciesProfileGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandSecurityOutboundPoliciesProfileGroupGroup(ctx, diags)
	}

	if !data.ForceCertInspection.IsNull() {
		result["forceCertInspection"] = data.ForceCertInspection.ValueBool()
	}

	return result
}

func (data *resourceSecurityOutboundPoliciesProfileGroupGroupModel) expandSecurityOutboundPoliciesProfileGroupGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityOutboundPoliciesSourcesModel) expandSecurityOutboundPoliciesSources(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityOutboundPoliciesModel) expandSecurityOutboundPoliciesSourcesList(ctx context.Context, l []resourceSecurityOutboundPoliciesSourcesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityOutboundPoliciesSources(ctx, diags)
	}
	return result
}
