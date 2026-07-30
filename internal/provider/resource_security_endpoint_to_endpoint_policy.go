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
var _ resource.Resource = &resourceSecurityEndpointToEndpointPolicy{}
var _ resource.ResourceWithMoveState = &resourceSecurityEndpointToEndpointPolicy{}

func newResourceSecurityEndpointToEndpointPolicy() resource.Resource {
	return &resourceSecurityEndpointToEndpointPolicy{}
}

type resourceSecurityEndpointToEndpointPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityEndpointToEndpointPolicyModel describes the resource data model.
type resourceSecurityEndpointToEndpointPolicyModel struct {
	ID           types.String                                               `tfsdk:"id"`
	PrimaryKey   types.String                                               `tfsdk:"primary_key"`
	Enabled      types.Bool                                                 `tfsdk:"enabled"`
	Users        []resourceSecurityEndpointToEndpointPolicyUsersModel       `tfsdk:"users"`
	Sources      []resourceSecurityEndpointToEndpointPolicySourcesModel     `tfsdk:"sources"`
	Services     []resourceSecurityEndpointToEndpointPolicyServicesModel    `tfsdk:"services"`
	Action       types.String                                               `tfsdk:"action"`
	Schedule     *resourceSecurityEndpointToEndpointPolicyScheduleModel     `tfsdk:"schedule"`
	Comments     types.String                                               `tfsdk:"comments"`
	ProfileGroup *resourceSecurityEndpointToEndpointPolicyProfileGroupModel `tfsdk:"profile_group"`
	LogTraffic   types.String                                               `tfsdk:"log_traffic"`
}

func (r *resourceSecurityEndpointToEndpointPolicy) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_endpoint_to_endpoint_policy"
}

func (r *resourceSecurityEndpointToEndpointPolicy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Endpoint to Endpoint Policy Resource API V2 for FortiSASE.",
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
			"users": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("auth/users", "auth/user-groups", "auth/ad-groups"),
							},
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
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("endpoint/ztna-tags", "endpoint/ztna-tag-rules"),
							},
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
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("security/services", "security/service-groups"),
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

func (r *resourceSecurityEndpointToEndpointPolicy) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *resourceSecurityEndpointToEndpointPolicy) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_endpoint_to_endpoint_policies" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityEndpointToEndpointPolicyModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityEndpointToEndpointPolicy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityEndpointToEndpointPolicyModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityEndpointToEndpointPolicy(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicy(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityEndpointToEndpointPolicies(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityEndpointToEndpointPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityEndpointToEndpointPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityEndpointToEndpointPolicy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityEndpointToEndpointPolicyModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityEndpointToEndpointPolicyModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityEndpointToEndpointPolicy(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicy(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityEndpointToEndpointPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityEndpointToEndpointPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityEndpointToEndpointPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityEndpointToEndpointPolicy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityEndpointToEndpointPolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicy(ctx, "delete", diags))

	output, err := c.DeleteSecurityEndpointToEndpointPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityEndpointToEndpointPolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityEndpointToEndpointPolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityEndpointToEndpointPolicies(&input_model)
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

	diags.Append(data.refreshSecurityEndpointToEndpointPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityEndpointToEndpointPolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityEndpointToEndpointPolicyModel) refreshSecurityEndpointToEndpointPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceSecurityEndpointToEndpointPolicyModel) getCreateObjectSecurityEndpointToEndpointPolicy(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	result["users"] = data.expandSecurityEndpointToEndpointPolicyUsersList(ctx, data.Users, diags)

	if data.Sources != nil {
		result["sources"] = data.expandSecurityEndpointToEndpointPolicySourcesList(ctx, data.Sources, diags)
	}

	result["services"] = data.expandSecurityEndpointToEndpointPolicyServicesList(ctx, data.Services, diags)

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityEndpointToEndpointPolicySchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isZeroStruct(*data.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityEndpointToEndpointPolicyProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.IsUnknown() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	return &result
}

func (data *resourceSecurityEndpointToEndpointPolicyModel) getUpdateObjectSecurityEndpointToEndpointPolicy(ctx context.Context, state resourceSecurityEndpointToEndpointPolicyModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if data.Users != nil {
		result["users"] = data.expandSecurityEndpointToEndpointPolicyUsersList(ctx, data.Users, diags)
	}

	if data.Sources != nil {
		result["sources"] = data.expandSecurityEndpointToEndpointPolicySourcesList(ctx, data.Sources, diags)
	}

	if data.Services != nil {
		result["services"] = data.expandSecurityEndpointToEndpointPolicyServicesList(ctx, data.Services, diags)
	}

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityEndpointToEndpointPolicySchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil {
		result["profileGroup"] = data.ProfileGroup.expandSecurityEndpointToEndpointPolicyProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.IsUnknown() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	return &result
}

func (data *resourceSecurityEndpointToEndpointPolicyModel) getURLObjectSecurityEndpointToEndpointPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityEndpointToEndpointPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityEndpointToEndpointPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityEndpointToEndpointPolicyServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityEndpointToEndpointPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityEndpointToEndpointPolicyProfileGroupModel struct {
	Group               *resourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                      `tfsdk:"force_cert_inspection"`
}

type resourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityEndpointToEndpointPolicyUsersModel) flattenSecurityEndpointToEndpointPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPolicyUsersModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPolicyUsersModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPolicyUsersModel{}
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

func (s *resourceSecurityEndpointToEndpointPolicyModel) flattenSecurityEndpointToEndpointPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityEndpointToEndpointPolicyUsersModel {
	if o == nil {
		return []resourceSecurityEndpointToEndpointPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []resourceSecurityEndpointToEndpointPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityEndpointToEndpointPolicyUsersModel{}
	}

	values := make([]resourceSecurityEndpointToEndpointPolicyUsersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityEndpointToEndpointPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityEndpointToEndpointPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityEndpointToEndpointPolicySourcesModel) flattenSecurityEndpointToEndpointPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPolicySourcesModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPolicySourcesModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPolicySourcesModel{}
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

func (s *resourceSecurityEndpointToEndpointPolicyModel) flattenSecurityEndpointToEndpointPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityEndpointToEndpointPolicySourcesModel {
	if o == nil {
		return []resourceSecurityEndpointToEndpointPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []resourceSecurityEndpointToEndpointPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityEndpointToEndpointPolicySourcesModel{}
	}

	values := make([]resourceSecurityEndpointToEndpointPolicySourcesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityEndpointToEndpointPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityEndpointToEndpointPolicySources(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityEndpointToEndpointPolicyServicesModel) flattenSecurityEndpointToEndpointPolicyServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPolicyServicesModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPolicyServicesModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPolicyServicesModel{}
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

func (s *resourceSecurityEndpointToEndpointPolicyModel) flattenSecurityEndpointToEndpointPolicyServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityEndpointToEndpointPolicyServicesModel {
	if o == nil {
		return []resourceSecurityEndpointToEndpointPolicyServicesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []resourceSecurityEndpointToEndpointPolicyServicesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityEndpointToEndpointPolicyServicesModel{}
	}

	values := make([]resourceSecurityEndpointToEndpointPolicyServicesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityEndpointToEndpointPolicyServicesModel
		if i < len(s.Services) {
			m = s.Services[i]
		}
		values[i] = *m.flattenSecurityEndpointToEndpointPolicyServices(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityEndpointToEndpointPolicyScheduleModel) flattenSecurityEndpointToEndpointPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPolicyScheduleModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPolicyScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPolicyScheduleModel{}
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

func (m *resourceSecurityEndpointToEndpointPolicyProfileGroupModel) flattenSecurityEndpointToEndpointPolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPolicyProfileGroupModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPolicyProfileGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPolicyProfileGroupModel{}
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

func (m *resourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel) flattenSecurityEndpointToEndpointPolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel {
	if input == nil {
		return &resourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel{}
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

func (data *resourceSecurityEndpointToEndpointPolicyUsersModel) expandSecurityEndpointToEndpointPolicyUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityEndpointToEndpointPolicyModel) expandSecurityEndpointToEndpointPolicyUsersList(ctx context.Context, l []resourceSecurityEndpointToEndpointPolicyUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityEndpointToEndpointPolicyUsers(ctx, diags)
	}
	return result
}

func (data *resourceSecurityEndpointToEndpointPolicySourcesModel) expandSecurityEndpointToEndpointPolicySources(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityEndpointToEndpointPolicyModel) expandSecurityEndpointToEndpointPolicySourcesList(ctx context.Context, l []resourceSecurityEndpointToEndpointPolicySourcesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityEndpointToEndpointPolicySources(ctx, diags)
	}
	return result
}

func (data *resourceSecurityEndpointToEndpointPolicyServicesModel) expandSecurityEndpointToEndpointPolicyServices(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityEndpointToEndpointPolicyModel) expandSecurityEndpointToEndpointPolicyServicesList(ctx context.Context, l []resourceSecurityEndpointToEndpointPolicyServicesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityEndpointToEndpointPolicyServices(ctx, diags)
	}
	return result
}

func (data *resourceSecurityEndpointToEndpointPolicyScheduleModel) expandSecurityEndpointToEndpointPolicySchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityEndpointToEndpointPolicyProfileGroupModel) expandSecurityEndpointToEndpointPolicyProfileGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	result["group"] = nil
	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandSecurityEndpointToEndpointPolicyProfileGroupGroup(ctx, diags)
	}

	if !data.ForceCertInspection.IsNull() && !data.ForceCertInspection.IsUnknown() {
		result["forceCertInspection"] = data.ForceCertInspection.ValueBool()
	}

	return result
}

func (data *resourceSecurityEndpointToEndpointPolicyProfileGroupGroupModel) expandSecurityEndpointToEndpointPolicyProfileGroupGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
