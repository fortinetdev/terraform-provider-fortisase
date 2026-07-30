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
var _ resource.Resource = &resourceSecurityInternalPolicy{}
var _ resource.ResourceWithMoveState = &resourceSecurityInternalPolicy{}

func newResourceSecurityInternalPolicy() resource.Resource {
	return &resourceSecurityInternalPolicy{}
}

type resourceSecurityInternalPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityInternalPolicyModel describes the resource data model.
type resourceSecurityInternalPolicyModel struct {
	ID                  types.String                                      `tfsdk:"id"`
	PrimaryKey          types.String                                      `tfsdk:"primary_key"`
	Enabled             types.Bool                                        `tfsdk:"enabled"`
	Scope               types.String                                      `tfsdk:"scope"`
	Users               []resourceSecurityInternalPolicyUsersModel        `tfsdk:"users"`
	Destinations        []resourceSecurityInternalPolicyDestinationsModel `tfsdk:"destinations"`
	Services            []resourceSecurityInternalPolicyServicesModel     `tfsdk:"services"`
	Action              types.String                                      `tfsdk:"action"`
	Schedule            *resourceSecurityInternalPolicyScheduleModel      `tfsdk:"schedule"`
	Comments            types.String                                      `tfsdk:"comments"`
	ProfileGroup        *resourceSecurityInternalPolicyProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic          types.String                                      `tfsdk:"log_traffic"`
	Sources             []resourceSecurityInternalPolicySourcesModel      `tfsdk:"sources"`
	CaptivePortalExempt types.Bool                                        `tfsdk:"captive_portal_exempt"`
}

func (r *resourceSecurityInternalPolicy) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_policy"
}

func (r *resourceSecurityInternalPolicy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Internal Policy Resource API V2 for FortiSASE.",
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

func (r *resourceSecurityInternalPolicy) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_internal_policy"
}
func (r *resourceSecurityInternalPolicy) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_internal_policies" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityInternalPolicyModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityInternalPolicy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityInternalPolicyModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityInternalPolicy(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityInternalPolicy(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityInternalPolicies(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityInternalPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}
	ignore_refresh := false
	if messageVal, hasMessage := read_output["message"]; hasMessage {
		if message, ok := messageVal.(string); ok {
			if message == "Warning: one or more Security PoP(s) are down. Returned available data." {
				ignore_refresh = true
				diags.AddWarning(
					fmt.Sprintf("Warning: one or more Security PoP(s) are down."),
					"Please go to GUI to check the status of the Security PoP(s).",
				)
				is_resource_created := false
				if dataList, ok := read_output["data"]; ok {
					if dataList, ok := dataList.([]interface{}); ok {
						for _, item := range dataList {
							if dataItem, ok := item.(map[string]interface{}); ok {
								if dataName, ok := dataItem["name"].(string); ok && dataName == data.PrimaryKey.ValueString() {
									is_resource_created = true
									break
								}
							}
						}
					}
				}
				if !is_resource_created {
					ignore_refresh = true
					diags.AddError(
						fmt.Sprintf("Error: resource %s not created", r.resourceName),
						getErrorDetail(&read_input_model, read_output),
					)
				}
			}
		}
	}
	if !ignore_refresh {
		diags.Append(data.refreshSecurityInternalPolicy(ctx, read_output)...)
		if diags.HasError() {
			return
		}
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalPolicy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityInternalPolicyModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityInternalPolicyModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityInternalPolicy(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityInternalPolicy(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityInternalPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityInternalPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalPolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalPolicy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityInternalPolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalPolicy(ctx, "delete", diags))

	output, err := c.DeleteSecurityInternalPolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityInternalPolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityInternalPolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalPolicies(&input_model)
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

	diags.Append(data.refreshSecurityInternalPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalPolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityInternalPolicyModel) refreshSecurityInternalPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.Users = m.flattenSecurityInternalPolicyUsersList(ctx, v, &diags)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityInternalPolicyDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityInternalPolicyServicesList(ctx, v, &diags)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityInternalPolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["profileGroup"]; ok {
		m.ProfileGroup = m.ProfileGroup.flattenSecurityInternalPolicyProfileGroup(ctx, v, &diags)
	}

	if v, ok := o["logTraffic"]; ok {
		m.LogTraffic = parseStringValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityInternalPolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["captivePortalExempt"]; ok {
		m.CaptivePortalExempt = parseBoolValue(v)
	}

	return diags
}

func (data *resourceSecurityInternalPolicyModel) getCreateObjectSecurityInternalPolicy(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
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

	result["users"] = data.expandSecurityInternalPolicyUsersList(ctx, data.Users, diags)

	result["destinations"] = data.expandSecurityInternalPolicyDestinationsList(ctx, data.Destinations, diags)

	result["services"] = data.expandSecurityInternalPolicyServicesList(ctx, data.Services, diags)

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityInternalPolicySchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isZeroStruct(*data.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityInternalPolicyProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.IsUnknown() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	if data.Sources != nil {
		result["sources"] = data.expandSecurityInternalPolicySourcesList(ctx, data.Sources, diags)
	}

	if !data.CaptivePortalExempt.IsNull() && !data.CaptivePortalExempt.IsUnknown() {
		result["captivePortalExempt"] = data.CaptivePortalExempt.ValueBool()
	}

	return &result
}

func (data *resourceSecurityInternalPolicyModel) getUpdateObjectSecurityInternalPolicy(ctx context.Context, state resourceSecurityInternalPolicyModel, diags *diag.Diagnostics) *map[string]interface{} {
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
		result["users"] = data.expandSecurityInternalPolicyUsersList(ctx, data.Users, diags)
	}

	if data.Destinations != nil {
		result["destinations"] = data.expandSecurityInternalPolicyDestinationsList(ctx, data.Destinations, diags)
	}

	if data.Services != nil {
		result["services"] = data.expandSecurityInternalPolicyServicesList(ctx, data.Services, diags)
	}

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityInternalPolicySchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil {
		result["profileGroup"] = data.ProfileGroup.expandSecurityInternalPolicyProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.IsUnknown() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	if data.Sources != nil {
		result["sources"] = data.expandSecurityInternalPolicySourcesList(ctx, data.Sources, diags)
	}

	if !data.CaptivePortalExempt.IsNull() && !data.CaptivePortalExempt.IsUnknown() {
		result["captivePortalExempt"] = data.CaptivePortalExempt.ValueBool()
	}

	return &result
}

func (data *resourceSecurityInternalPolicyModel) getURLObjectSecurityInternalPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityInternalPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalPolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalPolicyServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalPolicyProfileGroupModel struct {
	Group               *resourceSecurityInternalPolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                            `tfsdk:"force_cert_inspection"`
}

type resourceSecurityInternalPolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityInternalPolicyUsersModel) flattenSecurityInternalPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPolicyUsersModel {
	if input == nil {
		return &resourceSecurityInternalPolicyUsersModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPolicyUsersModel{}
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

func (s *resourceSecurityInternalPolicyModel) flattenSecurityInternalPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalPolicyUsersModel {
	if o == nil {
		return []resourceSecurityInternalPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []resourceSecurityInternalPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalPolicyUsersModel{}
	}

	values := make([]resourceSecurityInternalPolicyUsersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityInternalPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalPolicyDestinationsModel) flattenSecurityInternalPolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPolicyDestinationsModel {
	if input == nil {
		return &resourceSecurityInternalPolicyDestinationsModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPolicyDestinationsModel{}
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

func (s *resourceSecurityInternalPolicyModel) flattenSecurityInternalPolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalPolicyDestinationsModel {
	if o == nil {
		return []resourceSecurityInternalPolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []resourceSecurityInternalPolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalPolicyDestinationsModel{}
	}

	values := make([]resourceSecurityInternalPolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalPolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityInternalPolicyDestinations(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalPolicyServicesModel) flattenSecurityInternalPolicyServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPolicyServicesModel {
	if input == nil {
		return &resourceSecurityInternalPolicyServicesModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPolicyServicesModel{}
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

func (s *resourceSecurityInternalPolicyModel) flattenSecurityInternalPolicyServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalPolicyServicesModel {
	if o == nil {
		return []resourceSecurityInternalPolicyServicesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []resourceSecurityInternalPolicyServicesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalPolicyServicesModel{}
	}

	values := make([]resourceSecurityInternalPolicyServicesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalPolicyServicesModel
		if i < len(s.Services) {
			m = s.Services[i]
		}
		values[i] = *m.flattenSecurityInternalPolicyServices(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalPolicyScheduleModel) flattenSecurityInternalPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPolicyScheduleModel {
	if input == nil {
		return &resourceSecurityInternalPolicyScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPolicyScheduleModel{}
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

func (m *resourceSecurityInternalPolicyProfileGroupModel) flattenSecurityInternalPolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPolicyProfileGroupModel {
	if input == nil {
		return &resourceSecurityInternalPolicyProfileGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPolicyProfileGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenSecurityInternalPolicyProfileGroupGroup(ctx, v, diags)
	}

	if v, ok := o["forceCertInspection"]; ok {
		m.ForceCertInspection = parseBoolValue(v)
	}

	return m
}

func (m *resourceSecurityInternalPolicyProfileGroupGroupModel) flattenSecurityInternalPolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPolicyProfileGroupGroupModel {
	if input == nil {
		return &resourceSecurityInternalPolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPolicyProfileGroupGroupModel{}
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

func (m *resourceSecurityInternalPolicySourcesModel) flattenSecurityInternalPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalPolicySourcesModel {
	if input == nil {
		return &resourceSecurityInternalPolicySourcesModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalPolicySourcesModel{}
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

func (s *resourceSecurityInternalPolicyModel) flattenSecurityInternalPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalPolicySourcesModel {
	if o == nil {
		return []resourceSecurityInternalPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []resourceSecurityInternalPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalPolicySourcesModel{}
	}

	values := make([]resourceSecurityInternalPolicySourcesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityInternalPolicySources(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityInternalPolicyUsersModel) expandSecurityInternalPolicyUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalPolicyModel) expandSecurityInternalPolicyUsersList(ctx context.Context, l []resourceSecurityInternalPolicyUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalPolicyUsers(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalPolicyDestinationsModel) expandSecurityInternalPolicyDestinations(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalPolicyModel) expandSecurityInternalPolicyDestinationsList(ctx context.Context, l []resourceSecurityInternalPolicyDestinationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalPolicyDestinations(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalPolicyServicesModel) expandSecurityInternalPolicyServices(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalPolicyModel) expandSecurityInternalPolicyServicesList(ctx context.Context, l []resourceSecurityInternalPolicyServicesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalPolicyServices(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalPolicyScheduleModel) expandSecurityInternalPolicySchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityInternalPolicyProfileGroupModel) expandSecurityInternalPolicyProfileGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	result["group"] = nil
	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandSecurityInternalPolicyProfileGroupGroup(ctx, diags)
	}

	if !data.ForceCertInspection.IsNull() && !data.ForceCertInspection.IsUnknown() {
		result["forceCertInspection"] = data.ForceCertInspection.ValueBool()
	}

	return result
}

func (data *resourceSecurityInternalPolicyProfileGroupGroupModel) expandSecurityInternalPolicyProfileGroupGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityInternalPolicySourcesModel) expandSecurityInternalPolicySources(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalPolicyModel) expandSecurityInternalPolicySourcesList(ctx context.Context, l []resourceSecurityInternalPolicySourcesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalPolicySources(ctx, diags)
	}
	return result
}
