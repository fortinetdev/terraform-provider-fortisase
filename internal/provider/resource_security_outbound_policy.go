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
var _ resource.Resource = &resourceSecurityOutboundPolicy{}
var _ resource.ResourceWithMoveState = &resourceSecurityOutboundPolicy{}

func newResourceSecurityOutboundPolicy() resource.Resource {
	return &resourceSecurityOutboundPolicy{}
}

type resourceSecurityOutboundPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityOutboundPolicyModel describes the resource data model.
type resourceSecurityOutboundPolicyModel struct {
	ID                  types.String                                      `tfsdk:"id"`
	PrimaryKey          types.String                                      `tfsdk:"primary_key"`
	Enabled             types.Bool                                        `tfsdk:"enabled"`
	Scope               types.String                                      `tfsdk:"scope"`
	Users               []resourceSecurityOutboundPolicyUsersModel        `tfsdk:"users"`
	Destinations        []resourceSecurityOutboundPolicyDestinationsModel `tfsdk:"destinations"`
	Services            []resourceSecurityOutboundPolicyServicesModel     `tfsdk:"services"`
	Action              types.String                                      `tfsdk:"action"`
	Schedule            *resourceSecurityOutboundPolicyScheduleModel      `tfsdk:"schedule"`
	Comments            types.String                                      `tfsdk:"comments"`
	ProfileGroup        *resourceSecurityOutboundPolicyProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic          types.String                                      `tfsdk:"log_traffic"`
	Sources             []resourceSecurityOutboundPolicySourcesModel      `tfsdk:"sources"`
	CaptivePortalExempt types.Bool                                        `tfsdk:"captive_portal_exempt"`
}

func (r *resourceSecurityOutboundPolicy) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_outbound_policy"
}

func (r *resourceSecurityOutboundPolicy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Outbound Policy Resource API V2 for FortiSASE.",
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
			"scope": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("all", "vpn-user", "thin-edge", "all-pre-logon-users", "specify"),
				},
				Computed: true,
				Optional: true,
			},
			"action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("accept", "deny", "isolate"),
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
			"captive_portal_exempt": schema.BoolAttribute{
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
			"destinations": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "security/ip-threat-feeds", "security/proxy-address-rbi", "network/internet-services"),
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
			"sources": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "endpoint/ztna-tags", "endpoint/ztna-tag-rules", "security/ip-threat-feeds", "infra/ssids", "infra/fortigates", "infra/extenders"),
							},
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

func (r *resourceSecurityOutboundPolicy) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_outbound_policy"
}
func (r *resourceSecurityOutboundPolicy) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_outbound_policies" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityOutboundPolicyModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityOutboundPolicy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityOutboundPolicyModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityOutboundPolicy(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicy(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityOutboundPolicies(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityOutboundPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityOutboundPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityOutboundPolicy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityOutboundPolicyModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityOutboundPolicyModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityOutboundPolicy(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicy(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityOutboundPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityOutboundPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityOutboundPolicy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityOutboundPolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicy(ctx, "delete", diags))

	output, err := c.DeleteSecurityOutboundPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityOutboundPolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityOutboundPolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityOutboundPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityOutboundPolicies(&input_model)
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

	diags.Append(data.refreshSecurityOutboundPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityOutboundPolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityOutboundPolicyModel) refreshSecurityOutboundPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.Users = m.flattenSecurityOutboundPolicyUsersList(ctx, v, &diags)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityOutboundPolicyDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityOutboundPolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityOutboundPolicyProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityOutboundPolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["captivePortalExempt"]; ok {
		m.CaptivePortalExempt = parseBoolValue(v)
	}

	return diags
}

func (data *resourceSecurityOutboundPolicyModel) getCreateObjectSecurityOutboundPolicy(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.Scope.IsNull() && !data.Scope.IsUnknown() {
		result["scope"] = data.Scope.ValueString()
	}

	result["users"] = data.expandSecurityOutboundPolicyUsersList(ctx, data.Users, diags)

	result["destinations"] = data.expandSecurityOutboundPolicyDestinationsList(ctx, data.Destinations, diags)

	result["services"] = data.expandSecurityOutboundPolicyServicesList(ctx, data.Services, diags)

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityOutboundPolicySchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isZeroStruct(*data.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityOutboundPolicyProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.IsUnknown() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	result["sources"] = data.expandSecurityOutboundPolicySourcesList(ctx, data.Sources, diags)

	if !data.CaptivePortalExempt.IsNull() && !data.CaptivePortalExempt.IsUnknown() {
		result["captivePortalExempt"] = data.CaptivePortalExempt.ValueBool()
	}

	return &result
}

func (data *resourceSecurityOutboundPolicyModel) getUpdateObjectSecurityOutboundPolicy(ctx context.Context, state resourceSecurityOutboundPolicyModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.Scope.IsNull() && !data.Scope.IsUnknown() {
		result["scope"] = data.Scope.ValueString()
	}

	if data.Users != nil {
		result["users"] = data.expandSecurityOutboundPolicyUsersList(ctx, data.Users, diags)
	}

	if data.Destinations != nil {
		result["destinations"] = data.expandSecurityOutboundPolicyDestinationsList(ctx, data.Destinations, diags)
	}

	if data.Services != nil {
		result["services"] = data.expandSecurityOutboundPolicyServicesList(ctx, data.Services, diags)
	}

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityOutboundPolicySchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil {
		result["profileGroup"] = data.ProfileGroup.expandSecurityOutboundPolicyProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.IsUnknown() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	if data.Sources != nil {
		result["sources"] = data.expandSecurityOutboundPolicySourcesList(ctx, data.Sources, diags)
	}

	if !data.CaptivePortalExempt.IsNull() && !data.CaptivePortalExempt.IsUnknown() {
		result["captivePortalExempt"] = data.CaptivePortalExempt.ValueBool()
	}

	return &result
}

func (data *resourceSecurityOutboundPolicyModel) getURLObjectSecurityOutboundPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityOutboundPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityOutboundPolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityOutboundPolicyServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityOutboundPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityOutboundPolicyProfileGroupModel struct {
	Group               *resourceSecurityOutboundPolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                            `tfsdk:"force_cert_inspection"`
}

type resourceSecurityOutboundPolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityOutboundPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityOutboundPolicyUsersModel) flattenSecurityOutboundPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPolicyUsersModel {
	if input == nil {
		return &resourceSecurityOutboundPolicyUsersModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPolicyUsersModel{}
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

func (s *resourceSecurityOutboundPolicyModel) flattenSecurityOutboundPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityOutboundPolicyUsersModel {
	if o == nil {
		return []resourceSecurityOutboundPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []resourceSecurityOutboundPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityOutboundPolicyUsersModel{}
	}

	values := make([]resourceSecurityOutboundPolicyUsersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityOutboundPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityOutboundPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityOutboundPolicyDestinationsModel) flattenSecurityOutboundPolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPolicyDestinationsModel {
	if input == nil {
		return &resourceSecurityOutboundPolicyDestinationsModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPolicyDestinationsModel{}
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

func (s *resourceSecurityOutboundPolicyModel) flattenSecurityOutboundPolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityOutboundPolicyDestinationsModel {
	if o == nil {
		return []resourceSecurityOutboundPolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []resourceSecurityOutboundPolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityOutboundPolicyDestinationsModel{}
	}

	values := make([]resourceSecurityOutboundPolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityOutboundPolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityOutboundPolicyDestinations(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityOutboundPolicyServicesModel) flattenSecurityOutboundPolicyServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPolicyServicesModel {
	if input == nil {
		return &resourceSecurityOutboundPolicyServicesModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPolicyServicesModel{}
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

func (s *resourceSecurityOutboundPolicyModel) flattenSecurityOutboundPolicyServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityOutboundPolicyServicesModel {
	if o == nil {
		return []resourceSecurityOutboundPolicyServicesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []resourceSecurityOutboundPolicyServicesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityOutboundPolicyServicesModel{}
	}

	values := make([]resourceSecurityOutboundPolicyServicesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityOutboundPolicyServicesModel
		if i < len(s.Services) {
			m = s.Services[i]
		}
		values[i] = *m.flattenSecurityOutboundPolicyServices(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityOutboundPolicyScheduleModel) flattenSecurityOutboundPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPolicyScheduleModel {
	if input == nil {
		return &resourceSecurityOutboundPolicyScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPolicyScheduleModel{}
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

func (m *resourceSecurityOutboundPolicyProfileGroupModel) flattenSecurityOutboundPolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPolicyProfileGroupModel {
	if input == nil {
		return &resourceSecurityOutboundPolicyProfileGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPolicyProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityOutboundPolicyProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *resourceSecurityOutboundPolicyProfileGroupGroupModel) flattenSecurityOutboundPolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPolicyProfileGroupGroupModel {
	if input == nil {
		return &resourceSecurityOutboundPolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPolicyProfileGroupGroupModel{}
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

func (m *resourceSecurityOutboundPolicySourcesModel) flattenSecurityOutboundPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityOutboundPolicySourcesModel {
	if input == nil {
		return &resourceSecurityOutboundPolicySourcesModel{}
	}
	if m == nil {
		m = &resourceSecurityOutboundPolicySourcesModel{}
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

func (s *resourceSecurityOutboundPolicyModel) flattenSecurityOutboundPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityOutboundPolicySourcesModel {
	if o == nil {
		return []resourceSecurityOutboundPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []resourceSecurityOutboundPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityOutboundPolicySourcesModel{}
	}

	values := make([]resourceSecurityOutboundPolicySourcesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityOutboundPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityOutboundPolicySources(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityOutboundPolicyUsersModel) expandSecurityOutboundPolicyUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityOutboundPolicyModel) expandSecurityOutboundPolicyUsersList(ctx context.Context, l []resourceSecurityOutboundPolicyUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityOutboundPolicyUsers(ctx, diags)
	}
	return result
}

func (data *resourceSecurityOutboundPolicyDestinationsModel) expandSecurityOutboundPolicyDestinations(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityOutboundPolicyModel) expandSecurityOutboundPolicyDestinationsList(ctx context.Context, l []resourceSecurityOutboundPolicyDestinationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityOutboundPolicyDestinations(ctx, diags)
	}
	return result
}

func (data *resourceSecurityOutboundPolicyServicesModel) expandSecurityOutboundPolicyServices(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityOutboundPolicyModel) expandSecurityOutboundPolicyServicesList(ctx context.Context, l []resourceSecurityOutboundPolicyServicesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityOutboundPolicyServices(ctx, diags)
	}
	return result
}

func (data *resourceSecurityOutboundPolicyScheduleModel) expandSecurityOutboundPolicySchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityOutboundPolicyProfileGroupModel) expandSecurityOutboundPolicyProfileGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	result["group"] = nil
	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandSecurityOutboundPolicyProfileGroupGroup(ctx, diags)
	}

	if !data.ForceCertInspection.IsNull() && !data.ForceCertInspection.IsUnknown() {
		result["forceCertInspection"] = data.ForceCertInspection.ValueBool()
	}

	return result
}

func (data *resourceSecurityOutboundPolicyProfileGroupGroupModel) expandSecurityOutboundPolicyProfileGroupGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityOutboundPolicySourcesModel) expandSecurityOutboundPolicySources(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityOutboundPolicyModel) expandSecurityOutboundPolicySourcesList(ctx context.Context, l []resourceSecurityOutboundPolicySourcesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityOutboundPolicySources(ctx, diags)
	}
	return result
}
