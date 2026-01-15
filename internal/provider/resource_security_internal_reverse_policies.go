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
var _ resource.Resource = &resourceSecurityInternalReversePolicies{}

func newResourceSecurityInternalReversePolicies() resource.Resource {
	return &resourceSecurityInternalReversePolicies{}
}

type resourceSecurityInternalReversePolicies struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityInternalReversePoliciesModel describes the resource data model.
type resourceSecurityInternalReversePoliciesModel struct {
	ID           types.String                                               `tfsdk:"id"`
	PrimaryKey   types.String                                               `tfsdk:"primary_key"`
	Enabled      types.Bool                                                 `tfsdk:"enabled"`
	Scope        types.String                                               `tfsdk:"scope"`
	Sources      []resourceSecurityInternalReversePoliciesSourcesModel      `tfsdk:"sources"`
	Services     []resourceSecurityInternalReversePoliciesServicesModel     `tfsdk:"services"`
	Action       types.String                                               `tfsdk:"action"`
	Schedule     *resourceSecurityInternalReversePoliciesScheduleModel      `tfsdk:"schedule"`
	Comments     types.String                                               `tfsdk:"comments"`
	ProfileGroup *resourceSecurityInternalReversePoliciesProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic   types.String                                               `tfsdk:"log_traffic"`
	Destinations []resourceSecurityInternalReversePoliciesDestinationsModel `tfsdk:"destinations"`
}

func (r *resourceSecurityInternalReversePolicies) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_reverse_policies"
}

func (r *resourceSecurityInternalReversePolicies) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *resourceSecurityInternalReversePolicies) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_internal_reverse_policies"
}

func (r *resourceSecurityInternalReversePolicies) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityInternalReversePoliciesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityInternalReversePolicies(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicies(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityInternalReversePolicies(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalReversePolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalReversePolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalReversePolicies) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityInternalReversePoliciesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityInternalReversePoliciesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityInternalReversePolicies(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicies(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityInternalReversePolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalReversePolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalReversePolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalReversePolicies) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityInternalReversePoliciesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicies(ctx, "delete", diags))

	output, err := c.DeleteSecurityInternalReversePolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityInternalReversePolicies) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityInternalReversePoliciesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicies(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalReversePolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalReversePolicies(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalReversePolicies) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityInternalReversePoliciesModel) refreshSecurityInternalReversePolicies(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceSecurityInternalReversePoliciesModel) getCreateObjectSecurityInternalReversePolicies(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
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

	result["sources"] = data.expandSecurityInternalReversePoliciesSourcesList(ctx, data.Sources, diags)

	result["services"] = data.expandSecurityInternalReversePoliciesServicesList(ctx, data.Services, diags)

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityInternalReversePoliciesSchedule(ctx, diags)
	}

	if !data.Comments.IsNull() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isZeroStruct(*data.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityInternalReversePoliciesProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	if data.Destinations != nil {
		result["destinations"] = data.expandSecurityInternalReversePoliciesDestinationsList(ctx, data.Destinations, diags)
	}

	return &result
}

func (data *resourceSecurityInternalReversePoliciesModel) getUpdateObjectSecurityInternalReversePolicies(ctx context.Context, state resourceSecurityInternalReversePoliciesModel, diags *diag.Diagnostics) *map[string]interface{} {
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

	if data.Sources != nil {
		result["sources"] = data.expandSecurityInternalReversePoliciesSourcesList(ctx, data.Sources, diags)
	}

	if data.Services != nil {
		result["services"] = data.expandSecurityInternalReversePoliciesServicesList(ctx, data.Services, diags)
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Schedule != nil {
		result["schedule"] = data.Schedule.expandSecurityInternalReversePoliciesSchedule(ctx, diags)
	}

	if !data.Comments.IsNull() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil {
		result["profileGroup"] = data.ProfileGroup.expandSecurityInternalReversePoliciesProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	if data.Destinations != nil {
		result["destinations"] = data.expandSecurityInternalReversePoliciesDestinationsList(ctx, data.Destinations, diags)
	}

	return &result
}

func (data *resourceSecurityInternalReversePoliciesModel) getURLObjectSecurityInternalReversePolicies(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityInternalReversePoliciesSourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalReversePoliciesServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalReversePoliciesScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalReversePoliciesProfileGroupModel struct {
	Group               *resourceSecurityInternalReversePoliciesProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                     `tfsdk:"force_cert_inspection"`
}

type resourceSecurityInternalReversePoliciesProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalReversePoliciesDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityInternalReversePoliciesSourcesModel) flattenSecurityInternalReversePoliciesSources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePoliciesSourcesModel {
	if input == nil {
		return &resourceSecurityInternalReversePoliciesSourcesModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePoliciesSourcesModel{}
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

func (s *resourceSecurityInternalReversePoliciesModel) flattenSecurityInternalReversePoliciesSourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalReversePoliciesSourcesModel {
	if o == nil {
		return []resourceSecurityInternalReversePoliciesSourcesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []resourceSecurityInternalReversePoliciesSourcesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalReversePoliciesSourcesModel{}
	}

	values := make([]resourceSecurityInternalReversePoliciesSourcesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalReversePoliciesSourcesModel
		values[i] = *m.flattenSecurityInternalReversePoliciesSources(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalReversePoliciesServicesModel) flattenSecurityInternalReversePoliciesServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePoliciesServicesModel {
	if input == nil {
		return &resourceSecurityInternalReversePoliciesServicesModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePoliciesServicesModel{}
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

func (s *resourceSecurityInternalReversePoliciesModel) flattenSecurityInternalReversePoliciesServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalReversePoliciesServicesModel {
	if o == nil {
		return []resourceSecurityInternalReversePoliciesServicesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []resourceSecurityInternalReversePoliciesServicesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalReversePoliciesServicesModel{}
	}

	values := make([]resourceSecurityInternalReversePoliciesServicesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalReversePoliciesServicesModel
		values[i] = *m.flattenSecurityInternalReversePoliciesServices(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalReversePoliciesScheduleModel) flattenSecurityInternalReversePoliciesSchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePoliciesScheduleModel {
	if input == nil {
		return &resourceSecurityInternalReversePoliciesScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePoliciesScheduleModel{}
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

func (m *resourceSecurityInternalReversePoliciesProfileGroupModel) flattenSecurityInternalReversePoliciesProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePoliciesProfileGroupModel {
	if input == nil {
		return &resourceSecurityInternalReversePoliciesProfileGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePoliciesProfileGroupModel{}
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

func (m *resourceSecurityInternalReversePoliciesProfileGroupGroupModel) flattenSecurityInternalReversePoliciesProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePoliciesProfileGroupGroupModel {
	if input == nil {
		return &resourceSecurityInternalReversePoliciesProfileGroupGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePoliciesProfileGroupGroupModel{}
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

func (m *resourceSecurityInternalReversePoliciesDestinationsModel) flattenSecurityInternalReversePoliciesDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePoliciesDestinationsModel {
	if input == nil {
		return &resourceSecurityInternalReversePoliciesDestinationsModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePoliciesDestinationsModel{}
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

func (s *resourceSecurityInternalReversePoliciesModel) flattenSecurityInternalReversePoliciesDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalReversePoliciesDestinationsModel {
	if o == nil {
		return []resourceSecurityInternalReversePoliciesDestinationsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []resourceSecurityInternalReversePoliciesDestinationsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalReversePoliciesDestinationsModel{}
	}

	values := make([]resourceSecurityInternalReversePoliciesDestinationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalReversePoliciesDestinationsModel
		values[i] = *m.flattenSecurityInternalReversePoliciesDestinations(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityInternalReversePoliciesSourcesModel) expandSecurityInternalReversePoliciesSources(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalReversePoliciesModel) expandSecurityInternalReversePoliciesSourcesList(ctx context.Context, l []resourceSecurityInternalReversePoliciesSourcesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalReversePoliciesSources(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalReversePoliciesServicesModel) expandSecurityInternalReversePoliciesServices(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalReversePoliciesModel) expandSecurityInternalReversePoliciesServicesList(ctx context.Context, l []resourceSecurityInternalReversePoliciesServicesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalReversePoliciesServices(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalReversePoliciesScheduleModel) expandSecurityInternalReversePoliciesSchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityInternalReversePoliciesProfileGroupModel) expandSecurityInternalReversePoliciesProfileGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandSecurityInternalReversePoliciesProfileGroupGroup(ctx, diags)
	}

	if !data.ForceCertInspection.IsNull() {
		result["forceCertInspection"] = data.ForceCertInspection.ValueBool()
	}

	return result
}

func (data *resourceSecurityInternalReversePoliciesProfileGroupGroupModel) expandSecurityInternalReversePoliciesProfileGroupGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityInternalReversePoliciesDestinationsModel) expandSecurityInternalReversePoliciesDestinations(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalReversePoliciesModel) expandSecurityInternalReversePoliciesDestinationsList(ctx context.Context, l []resourceSecurityInternalReversePoliciesDestinationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalReversePoliciesDestinations(ctx, diags)
	}
	return result
}
