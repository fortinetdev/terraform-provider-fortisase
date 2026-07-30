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
var _ resource.Resource = &resourceSecurityInternalReversePolicy{}
var _ resource.ResourceWithMoveState = &resourceSecurityInternalReversePolicy{}

func newResourceSecurityInternalReversePolicy() resource.Resource {
	return &resourceSecurityInternalReversePolicy{}
}

type resourceSecurityInternalReversePolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityInternalReversePolicyModel describes the resource data model.
type resourceSecurityInternalReversePolicyModel struct {
	ID           types.String                                             `tfsdk:"id"`
	PrimaryKey   types.String                                             `tfsdk:"primary_key"`
	Enabled      types.Bool                                               `tfsdk:"enabled"`
	Scope        types.String                                             `tfsdk:"scope"`
	Sources      []resourceSecurityInternalReversePolicySourcesModel      `tfsdk:"sources"`
	Services     []resourceSecurityInternalReversePolicyServicesModel     `tfsdk:"services"`
	Action       types.String                                             `tfsdk:"action"`
	Schedule     *resourceSecurityInternalReversePolicyScheduleModel      `tfsdk:"schedule"`
	Comments     types.String                                             `tfsdk:"comments"`
	ProfileGroup *resourceSecurityInternalReversePolicyProfileGroupModel  `tfsdk:"profile_group"`
	LogTraffic   types.String                                             `tfsdk:"log_traffic"`
	Destinations []resourceSecurityInternalReversePolicyDestinationsModel `tfsdk:"destinations"`
}

func (r *resourceSecurityInternalReversePolicy) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_internal_reverse_policy"
}

func (r *resourceSecurityInternalReversePolicy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Internal Reverse Policy Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.OneOf("all", "vpn-user", "thin-edge", "specify"),
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
			"sources": schema.ListNestedAttribute{
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
			"destinations": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "infra/ssids", "infra/fortigates", "infra/extenders"),
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

func (r *resourceSecurityInternalReversePolicy) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *resourceSecurityInternalReversePolicy) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_internal_reverse_policies" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityInternalReversePolicyModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityInternalReversePolicy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityInternalReversePolicyModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityInternalReversePolicy(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicy(ctx, "create", diags))

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
	read_input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalReversePolicies(&read_input_model)
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
		diags.Append(data.refreshSecurityInternalReversePolicy(ctx, read_output)...)
		if diags.HasError() {
			return
		}
	}
	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalReversePolicy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityInternalReversePolicyModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityInternalReversePolicyModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityInternalReversePolicy(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicy(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalReversePolicies(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityInternalReversePolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalReversePolicy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("profile-group")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityInternalReversePolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicy(ctx, "delete", diags))

	output, err := c.DeleteSecurityInternalReversePolicies(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityInternalReversePolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityInternalReversePolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityInternalReversePolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityInternalReversePolicies(&input_model)
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

	diags.Append(data.refreshSecurityInternalReversePolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityInternalReversePolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityInternalReversePolicyModel) refreshSecurityInternalReversePolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceSecurityInternalReversePolicyModel) getCreateObjectSecurityInternalReversePolicy(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
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

	result["sources"] = data.expandSecurityInternalReversePolicySourcesList(ctx, data.Sources, diags)

	result["services"] = data.expandSecurityInternalReversePolicyServicesList(ctx, data.Services, diags)

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityInternalReversePolicySchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil && !isZeroStruct(*data.ProfileGroup) {
		result["profileGroup"] = data.ProfileGroup.expandSecurityInternalReversePolicyProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.IsUnknown() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	if data.Destinations != nil {
		result["destinations"] = data.expandSecurityInternalReversePolicyDestinationsList(ctx, data.Destinations, diags)
	}

	return &result
}

func (data *resourceSecurityInternalReversePolicyModel) getUpdateObjectSecurityInternalReversePolicy(ctx context.Context, state resourceSecurityInternalReversePolicyModel, diags *diag.Diagnostics) *map[string]interface{} {
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

	if data.Sources != nil {
		result["sources"] = data.expandSecurityInternalReversePolicySourcesList(ctx, data.Sources, diags)
	}

	if data.Services != nil {
		result["services"] = data.expandSecurityInternalReversePolicyServicesList(ctx, data.Services, diags)
	}

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityInternalReversePolicySchedule(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.ProfileGroup != nil {
		result["profileGroup"] = data.ProfileGroup.expandSecurityInternalReversePolicyProfileGroup(ctx, diags)
	}

	if !data.LogTraffic.IsNull() && !data.LogTraffic.IsUnknown() {
		result["logTraffic"] = data.LogTraffic.ValueString()
	}

	if data.Destinations != nil {
		result["destinations"] = data.expandSecurityInternalReversePolicyDestinationsList(ctx, data.Destinations, diags)
	}

	return &result
}

func (data *resourceSecurityInternalReversePolicyModel) getURLObjectSecurityInternalReversePolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityInternalReversePolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalReversePolicyServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalReversePolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalReversePolicyProfileGroupModel struct {
	Group               *resourceSecurityInternalReversePolicyProfileGroupGroupModel `tfsdk:"group"`
	ForceCertInspection types.Bool                                                   `tfsdk:"force_cert_inspection"`
}

type resourceSecurityInternalReversePolicyProfileGroupGroupModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityInternalReversePolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityInternalReversePolicySourcesModel) flattenSecurityInternalReversePolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePolicySourcesModel {
	if input == nil {
		return &resourceSecurityInternalReversePolicySourcesModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePolicySourcesModel{}
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

func (s *resourceSecurityInternalReversePolicyModel) flattenSecurityInternalReversePolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalReversePolicySourcesModel {
	if o == nil {
		return []resourceSecurityInternalReversePolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []resourceSecurityInternalReversePolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalReversePolicySourcesModel{}
	}

	values := make([]resourceSecurityInternalReversePolicySourcesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalReversePolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityInternalReversePolicySources(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalReversePolicyServicesModel) flattenSecurityInternalReversePolicyServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePolicyServicesModel {
	if input == nil {
		return &resourceSecurityInternalReversePolicyServicesModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePolicyServicesModel{}
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

func (s *resourceSecurityInternalReversePolicyModel) flattenSecurityInternalReversePolicyServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalReversePolicyServicesModel {
	if o == nil {
		return []resourceSecurityInternalReversePolicyServicesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []resourceSecurityInternalReversePolicyServicesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalReversePolicyServicesModel{}
	}

	values := make([]resourceSecurityInternalReversePolicyServicesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalReversePolicyServicesModel
		if i < len(s.Services) {
			m = s.Services[i]
		}
		values[i] = *m.flattenSecurityInternalReversePolicyServices(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityInternalReversePolicyScheduleModel) flattenSecurityInternalReversePolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePolicyScheduleModel {
	if input == nil {
		return &resourceSecurityInternalReversePolicyScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePolicyScheduleModel{}
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

func (m *resourceSecurityInternalReversePolicyProfileGroupModel) flattenSecurityInternalReversePolicyProfileGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePolicyProfileGroupModel {
	if input == nil {
		return &resourceSecurityInternalReversePolicyProfileGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePolicyProfileGroupModel{}
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

func (m *resourceSecurityInternalReversePolicyProfileGroupGroupModel) flattenSecurityInternalReversePolicyProfileGroupGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePolicyProfileGroupGroupModel {
	if input == nil {
		return &resourceSecurityInternalReversePolicyProfileGroupGroupModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePolicyProfileGroupGroupModel{}
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

func (m *resourceSecurityInternalReversePolicyDestinationsModel) flattenSecurityInternalReversePolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityInternalReversePolicyDestinationsModel {
	if input == nil {
		return &resourceSecurityInternalReversePolicyDestinationsModel{}
	}
	if m == nil {
		m = &resourceSecurityInternalReversePolicyDestinationsModel{}
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

func (s *resourceSecurityInternalReversePolicyModel) flattenSecurityInternalReversePolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityInternalReversePolicyDestinationsModel {
	if o == nil {
		return []resourceSecurityInternalReversePolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []resourceSecurityInternalReversePolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityInternalReversePolicyDestinationsModel{}
	}

	values := make([]resourceSecurityInternalReversePolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityInternalReversePolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityInternalReversePolicyDestinations(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityInternalReversePolicySourcesModel) expandSecurityInternalReversePolicySources(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalReversePolicyModel) expandSecurityInternalReversePolicySourcesList(ctx context.Context, l []resourceSecurityInternalReversePolicySourcesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalReversePolicySources(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalReversePolicyServicesModel) expandSecurityInternalReversePolicyServices(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalReversePolicyModel) expandSecurityInternalReversePolicyServicesList(ctx context.Context, l []resourceSecurityInternalReversePolicyServicesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalReversePolicyServices(ctx, diags)
	}
	return result
}

func (data *resourceSecurityInternalReversePolicyScheduleModel) expandSecurityInternalReversePolicySchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityInternalReversePolicyProfileGroupModel) expandSecurityInternalReversePolicyProfileGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	result["group"] = nil
	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandSecurityInternalReversePolicyProfileGroupGroup(ctx, diags)
	}

	if !data.ForceCertInspection.IsNull() && !data.ForceCertInspection.IsUnknown() {
		result["forceCertInspection"] = data.ForceCertInspection.ValueBool()
	}

	return result
}

func (data *resourceSecurityInternalReversePolicyProfileGroupGroupModel) expandSecurityInternalReversePolicyProfileGroupGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityInternalReversePolicyDestinationsModel) expandSecurityInternalReversePolicyDestinations(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityInternalReversePolicyModel) expandSecurityInternalReversePolicyDestinationsList(ctx context.Context, l []resourceSecurityInternalReversePolicyDestinationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityInternalReversePolicyDestinations(ctx, diags)
	}
	return result
}
