// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
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
var _ resource.Resource = &resourceSecurityInternalProxyPolicy{}

func newResourceSecurityInternalProxyPolicy() resource.Resource {
	return &resourceSecurityInternalProxyPolicy{}
}

type resourceSecurityInternalProxyPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityInternalProxyPolicyModel describes the resource data model.
type resourceSecurityInternalProxyPolicyModel struct {
	ID           types.String                                           `tfsdk:"id"`
	PrimaryKey   types.String                                           `tfsdk:"primary_key"`
	Enabled      types.Bool                                             `tfsdk:"enabled"`
	Sources      []resourceSecurityInternalProxyPolicySourcesModel      `tfsdk:"sources"`
	Users        []resourceSecurityInternalProxyPolicyUsersModel        `tfsdk:"users"`
	Destinations []resourceSecurityInternalProxyPolicyDestinationsModel `tfsdk:"destinations"`
	Action       types.String                                           `tfsdk:"action"`
	Schedule     *resourceSecurityInternalProxyPolicyScheduleModel      `tfsdk:"schedule"`
	Comments     types.String                                           `tfsdk:"comments"`
	ProfileGroup *resourceSecurityInternalProxyPolicyProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic   types.String                                           `tfsdk:"log_traffic"`
}

func (r *resourceSecurityInternalProxyPolicy) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_proxy_policy"
}

func (r *resourceSecurityInternalProxyPolicy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Internal Proxy Policy Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 35),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("accept", "deny"),
				},
				Computed: true,
				Optional: true,
			},
			"comments": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(1023),
				},
				Computed: true,
				Optional: true,
			},
			"log_traffic": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("all", "utm", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"sources": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "endpoint/ztna-tags", "endpoint/ztna-tag-rules", "security/ip-threat-feeds"),
							},
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"users": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("auth/users", "auth/user-groups"),
							},
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
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "security/ip-threat-feeds"),
							},
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
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("security/onetime-schedules", "security/recurring-schedules", "security/schedule-groups"),
						},
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
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.OneOf("security/profile-groups"),
								},
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

func (r *resourceSecurityInternalProxyPolicy) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_internal_proxy_policy"
}

func (r *resourceSecurityInternalProxyPolicy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityInternalProxyPolicyModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityInternalProxyPolicy(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityInternalProxyPolicy(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityInternalProxyPolicy(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	if responseMkey, ok := getCreateResponseMkey(output, "primaryKey"); ok {
		mkey = responseMkey
	}
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityInternalProxyPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalProxyPolicy(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalProxyPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalProxyPolicy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityInternalProxyPolicyModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityInternalProxyPolicyModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityInternalProxyPolicy(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityInternalProxyPolicy(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityInternalProxyPolicy(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityInternalProxyPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalProxyPolicy(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalProxyPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalProxyPolicy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityInternalProxyPolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalProxyPolicy(ctx, "delete", diags))

	output, err := c.DeleteSecurityInternalProxyPolicy(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityInternalProxyPolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityInternalProxyPolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalProxyPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalProxyPolicy(&input_model)
	if err != nil {
		if isNotFoundResponse(read_output) {
			resp.State.RemoveResource(ctx)
			return
		}
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalProxyPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalProxyPolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityInternalProxyPolicyModel) refreshSecurityInternalProxyPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityInternalProxyPolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["users"]; ok {
		m.Users = m.flattenSecurityInternalProxyPolicyUsersList(ctx, v, &diags)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityInternalProxyPolicyDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityInternalProxyPolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityInternalProxyPolicyProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	return diags
}

func (data *resourceSecurityInternalProxyPolicyModel) getCreateObjectSecurityInternalProxyPolicy(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	result["sources"] = data.expandSecurityInternalProxyPolicySourcesList(ctx, data.Sources, diags)

	result["users"] = data.expandSecurityInternalProxyPolicyUsersList(ctx, data.Users, diags)

	result["destinations"] = data.expandSecurityInternalProxyPolicyDestinationsList(ctx, data.Destinations, diags)

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityInternalProxyPolicySchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isZeroStruct(*data.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityInternalProxyPolicyProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.IsUnknown() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	return &result
}

func (data *resourceSecurityInternalProxyPolicyModel) getUpdateObjectSecurityInternalProxyPolicy(ctx context.Context, state resourceSecurityInternalProxyPolicyModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if data.Sources != nil {
		result["sources"] = data.expandSecurityInternalProxyPolicySourcesList(ctx, data.Sources, diags)
	}

	if data.Users != nil {
		result["users"] = data.expandSecurityInternalProxyPolicyUsersList(ctx, data.Users, diags)
	}

	if data.Destinations != nil {
		result["destinations"] = data.expandSecurityInternalProxyPolicyDestinationsList(ctx, data.Destinations, diags)
	}

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityInternalProxyPolicySchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil {
		result["profileGroup"] = data.ProfileGroup.expandSecurityInternalProxyPolicyProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.IsUnknown() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	return &result
}

func (data *resourceSecurityInternalProxyPolicyModel) getURLObjectSecurityInternalProxyPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityInternalProxyPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalProxyPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalProxyPolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalProxyPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalProxyPolicyProfileGroupModel struct {
	Group               *resourceSecurityInternalProxyPolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                 `tfsdk:"force_cert_inspection"`
}

type resourceSecurityInternalProxyPolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityInternalProxyPolicySourcesModel) flattenSecurityInternalProxyPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalProxyPolicySourcesModel {
	if input == nil {
		return &resourceSecurityInternalProxyPolicySourcesModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalProxyPolicySourcesModel{}
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

func (s *resourceSecurityInternalProxyPolicyModel) flattenSecurityInternalProxyPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalProxyPolicySourcesModel {
	if o == nil {
		return []resourceSecurityInternalProxyPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []resourceSecurityInternalProxyPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalProxyPolicySourcesModel{}
	}

	values := make([]resourceSecurityInternalProxyPolicySourcesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalProxyPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityInternalProxyPolicySources(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalProxyPolicyUsersModel) flattenSecurityInternalProxyPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalProxyPolicyUsersModel {
	if input == nil {
		return &resourceSecurityInternalProxyPolicyUsersModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalProxyPolicyUsersModel{}
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

func (s *resourceSecurityInternalProxyPolicyModel) flattenSecurityInternalProxyPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalProxyPolicyUsersModel {
	if o == nil {
		return []resourceSecurityInternalProxyPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []resourceSecurityInternalProxyPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalProxyPolicyUsersModel{}
	}

	values := make([]resourceSecurityInternalProxyPolicyUsersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalProxyPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityInternalProxyPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalProxyPolicyDestinationsModel) flattenSecurityInternalProxyPolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalProxyPolicyDestinationsModel {
	if input == nil {
		return &resourceSecurityInternalProxyPolicyDestinationsModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalProxyPolicyDestinationsModel{}
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

func (s *resourceSecurityInternalProxyPolicyModel) flattenSecurityInternalProxyPolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalProxyPolicyDestinationsModel {
	if o == nil {
		return []resourceSecurityInternalProxyPolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []resourceSecurityInternalProxyPolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalProxyPolicyDestinationsModel{}
	}

	values := make([]resourceSecurityInternalProxyPolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalProxyPolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityInternalProxyPolicyDestinations(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalProxyPolicyScheduleModel) flattenSecurityInternalProxyPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalProxyPolicyScheduleModel {
	if input == nil {
		return &resourceSecurityInternalProxyPolicyScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalProxyPolicyScheduleModel{}
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

func (m *resourceSecurityInternalProxyPolicyProfileGroupModel) flattenSecurityInternalProxyPolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalProxyPolicyProfileGroupModel {
	if input == nil {
		return &resourceSecurityInternalProxyPolicyProfileGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalProxyPolicyProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityInternalProxyPolicyProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *resourceSecurityInternalProxyPolicyProfileGroupGroupModel) flattenSecurityInternalProxyPolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalProxyPolicyProfileGroupGroupModel {
	if input == nil {
		return &resourceSecurityInternalProxyPolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalProxyPolicyProfileGroupGroupModel{}
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

func (data *resourceSecurityInternalProxyPolicySourcesModel) expandSecurityInternalProxyPolicySources(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalProxyPolicyModel) expandSecurityInternalProxyPolicySourcesList(ctx context.Context, l []resourceSecurityInternalProxyPolicySourcesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalProxyPolicySources(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalProxyPolicyUsersModel) expandSecurityInternalProxyPolicyUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalProxyPolicyModel) expandSecurityInternalProxyPolicyUsersList(ctx context.Context, l []resourceSecurityInternalProxyPolicyUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalProxyPolicyUsers(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalProxyPolicyDestinationsModel) expandSecurityInternalProxyPolicyDestinations(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalProxyPolicyModel) expandSecurityInternalProxyPolicyDestinationsList(ctx context.Context, l []resourceSecurityInternalProxyPolicyDestinationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalProxyPolicyDestinations(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalProxyPolicyScheduleModel) expandSecurityInternalProxyPolicySchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityInternalProxyPolicyProfileGroupModel) expandSecurityInternalProxyPolicyProfileGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	result["group"] = nil
	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandSecurityInternalProxyPolicyProfileGroupGroup(ctx, diags)
	}

	if !data.ForceCertInspection.IsNull() && !data.ForceCertInspection.IsUnknown() {
		result["forceCertInspection"] = data.ForceCertInspection.ValueBool()
	}

	return result
}

func (data *resourceSecurityInternalProxyPolicyProfileGroupGroupModel) expandSecurityInternalProxyPolicyProfileGroupGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
