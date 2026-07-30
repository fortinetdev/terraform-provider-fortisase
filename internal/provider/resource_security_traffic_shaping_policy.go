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
var _ resource.Resource = &resourceSecurityTrafficShapingPolicy{}

func newResourceSecurityTrafficShapingPolicy() resource.Resource {
	return &resourceSecurityTrafficShapingPolicy{}
}

type resourceSecurityTrafficShapingPolicy struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityTrafficShapingPolicyModel describes the resource data model.
type resourceSecurityTrafficShapingPolicyModel struct {
	ID                   types.String                                                   `tfsdk:"id"`
	PrimaryKey           types.String                                                   `tfsdk:"primary_key"`
	Comment              types.String                                                   `tfsdk:"comment"`
	Status               types.String                                                   `tfsdk:"status"`
	TrafficDirection     types.String                                                   `tfsdk:"traffic_direction"`
	Scope                types.String                                                   `tfsdk:"scope"`
	Sources              []resourceSecurityTrafficShapingPolicySourcesModel             `tfsdk:"sources"`
	DestinationScope     types.String                                                   `tfsdk:"destination_scope"`
	Destinations         []resourceSecurityTrafficShapingPolicyDestinationsModel        `tfsdk:"destinations"`
	Users                []resourceSecurityTrafficShapingPolicyUsersModel               `tfsdk:"users"`
	Schedule             *resourceSecurityTrafficShapingPolicyScheduleModel             `tfsdk:"schedule"`
	Services             []resourceSecurityTrafficShapingPolicyServicesModel            `tfsdk:"services"`
	Applications         []resourceSecurityTrafficShapingPolicyApplicationsModel        `tfsdk:"applications"`
	AppCategories        []resourceSecurityTrafficShapingPolicyAppCategoriesModel       `tfsdk:"app_categories"`
	UrlCategories        []resourceSecurityTrafficShapingPolicyUrlCategoriesModel       `tfsdk:"url_categories"`
	TrafficShaper        *resourceSecurityTrafficShapingPolicyTrafficShaperModel        `tfsdk:"traffic_shaper"`
	TrafficShaperReverse *resourceSecurityTrafficShapingPolicyTrafficShaperReverseModel `tfsdk:"traffic_shaper_reverse"`
	PerIpShaper          *resourceSecurityTrafficShapingPolicyPerIpShaperModel          `tfsdk:"per_ip_shaper"`
}

func (r *resourceSecurityTrafficShapingPolicy) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_traffic_shaping_policy"
}

func (r *resourceSecurityTrafficShapingPolicy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Traffic Shaping Policy Resource API V2 for FortiSASE.",
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
			"comment": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(255),
				},
				Computed: true,
				Optional: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"traffic_direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("outbound", "internal", "internal-reverse"),
				},
				Computed: true,
				Optional: true,
			},
			"scope": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("all", "vpn-user", "specify"),
				},
				Computed: true,
				Optional: true,
			},
			"destination_scope": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("all", "vpn-user", "specify"),
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
								stringvalidatorwarning.OneOf("network/hosts", "network/host-groups", "endpoint/ztna-tags", "security/ip-threat-feeds", "infra/ssids", "infra/fortigates", "infra/extenders"),
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
			"applications": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("security/applications"),
							},
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"app_categories": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("security/application-categories"),
							},
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"url_categories": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"traffic_shaper": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("security/traffic-shapers"),
						},
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"traffic_shaper_reverse": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("security/traffic-shapers"),
						},
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"per_ip_shaper": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("security/per-ip-traffic-shapers"),
						},
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityTrafficShapingPolicy) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_traffic_shaping_policy"
}

func (r *resourceSecurityTrafficShapingPolicy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityTrafficShapingPolicy")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityTrafficShapingPolicyModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityTrafficShapingPolicy(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityTrafficShapingPolicy(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityTrafficShapingPolicy(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityTrafficShapingPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityTrafficShapingPolicy(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityTrafficShapingPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityTrafficShapingPolicy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityTrafficShapingPolicy")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityTrafficShapingPolicyModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityTrafficShapingPolicyModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityTrafficShapingPolicy(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityTrafficShapingPolicy(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityTrafficShapingPolicy(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityTrafficShapingPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityTrafficShapingPolicy(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityTrafficShapingPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityTrafficShapingPolicy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityTrafficShapingPolicy")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityTrafficShapingPolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityTrafficShapingPolicy(ctx, "delete", diags))

	output, err := c.DeleteSecurityTrafficShapingPolicy(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityTrafficShapingPolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityTrafficShapingPolicyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityTrafficShapingPolicy(ctx, "read", diags))

	read_output, err := c.ReadSecurityTrafficShapingPolicy(&input_model)
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

	diags.Append(data.refreshSecurityTrafficShapingPolicy(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityTrafficShapingPolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityTrafficShapingPolicyModel) refreshSecurityTrafficShapingPolicy(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["comment"]; ok {
		m.Comment = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["trafficDirection"]; ok {
		m.TrafficDirection = parseStringValue(v)
	}

	if v, ok := o["scope"]; ok {
		m.Scope = parseStringValue(v)
	}

	if v, ok := o["sources"]; ok {
		m.Sources = m.flattenSecurityTrafficShapingPolicySourcesList(ctx, v, &diags)
	}

	if v, ok := o["destinationScope"]; ok {
		m.DestinationScope = parseStringValue(v)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityTrafficShapingPolicyDestinationsList(ctx, v, &diags)
	}

	if v, ok := o["users"]; ok {
		m.Users = m.flattenSecurityTrafficShapingPolicyUsersList(ctx, v, &diags)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityTrafficShapingPolicySchedule(ctx, v, &diags)
	}

	if v, ok := o["services"]; ok {
		m.Services = m.flattenSecurityTrafficShapingPolicyServicesList(ctx, v, &diags)
	}

	if v, ok := o["applications"]; ok {
		m.Applications = m.flattenSecurityTrafficShapingPolicyApplicationsList(ctx, v, &diags)
	}

	if v, ok := o["appCategories"]; ok {
		m.AppCategories = m.flattenSecurityTrafficShapingPolicyAppCategoriesList(ctx, v, &diags)
	}

	if v, ok := o["urlCategories"]; ok {
		m.UrlCategories = m.flattenSecurityTrafficShapingPolicyUrlCategoriesList(ctx, v, &diags)
	}

	if v, ok := o["trafficShaper"]; ok {
		m.TrafficShaper = m.TrafficShaper.flattenSecurityTrafficShapingPolicyTrafficShaper(ctx, v, &diags)
	}

	if v, ok := o["trafficShaperReverse"]; ok {
		m.TrafficShaperReverse = m.TrafficShaperReverse.flattenSecurityTrafficShapingPolicyTrafficShaperReverse(ctx, v, &diags)
	}

	if v, ok := o["perIpShaper"]; ok {
		m.PerIpShaper = m.PerIpShaper.flattenSecurityTrafficShapingPolicyPerIpShaper(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityTrafficShapingPolicyModel) getCreateObjectSecurityTrafficShapingPolicy(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		result["comment"] = data.Comment.ValueString()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if !data.TrafficDirection.IsNull() && !data.TrafficDirection.IsUnknown() {
		result["trafficDirection"] = data.TrafficDirection.ValueString()
	}

	if !data.Scope.IsNull() && !data.Scope.IsUnknown() {
		result["scope"] = data.Scope.ValueString()
	}

	result["sources"] = data.expandSecurityTrafficShapingPolicySourcesList(ctx, data.Sources, diags)

	if !data.DestinationScope.IsNull() && !data.DestinationScope.IsUnknown() {
		result["destinationScope"] = data.DestinationScope.ValueString()
	}

	result["destinations"] = data.expandSecurityTrafficShapingPolicyDestinationsList(ctx, data.Destinations, diags)

	result["users"] = data.expandSecurityTrafficShapingPolicyUsersList(ctx, data.Users, diags)

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityTrafficShapingPolicySchedule(ctx, diags)
	}

	result["services"] = data.expandSecurityTrafficShapingPolicyServicesList(ctx, data.Services, diags)

	result["applications"] = data.expandSecurityTrafficShapingPolicyApplicationsList(ctx, data.Applications, diags)

	result["appCategories"] = data.expandSecurityTrafficShapingPolicyAppCategoriesList(ctx, data.AppCategories, diags)

	result["urlCategories"] = data.expandSecurityTrafficShapingPolicyUrlCategoriesList(ctx, data.UrlCategories, diags)

	result["trafficShaper"] = nil
	if data.TrafficShaper != nil && !isZeroStruct(*data.TrafficShaper) {
		result["trafficShaper"] = data.TrafficShaper.expandSecurityTrafficShapingPolicyTrafficShaper(ctx, diags)
	}

	result["trafficShaperReverse"] = nil
	if data.TrafficShaperReverse != nil && !isZeroStruct(*data.TrafficShaperReverse) {
		result["trafficShaperReverse"] = data.TrafficShaperReverse.expandSecurityTrafficShapingPolicyTrafficShaperReverse(ctx, diags)
	}

	result["perIpShaper"] = nil
	if data.PerIpShaper != nil && !isZeroStruct(*data.PerIpShaper) {
		result["perIpShaper"] = data.PerIpShaper.expandSecurityTrafficShapingPolicyPerIpShaper(ctx, diags)
	}

	return &result
}

func (data *resourceSecurityTrafficShapingPolicyModel) getUpdateObjectSecurityTrafficShapingPolicy(ctx context.Context, state resourceSecurityTrafficShapingPolicyModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		result["comment"] = data.Comment.ValueString()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if !data.TrafficDirection.IsNull() && !data.TrafficDirection.IsUnknown() {
		result["trafficDirection"] = data.TrafficDirection.ValueString()
	}

	if !data.Scope.IsNull() && !data.Scope.IsUnknown() {
		result["scope"] = data.Scope.ValueString()
	}

	if data.Sources != nil {
		result["sources"] = data.expandSecurityTrafficShapingPolicySourcesList(ctx, data.Sources, diags)
	}

	if !data.DestinationScope.IsNull() && !data.DestinationScope.IsUnknown() {
		result["destinationScope"] = data.DestinationScope.ValueString()
	}

	if data.Destinations != nil {
		result["destinations"] = data.expandSecurityTrafficShapingPolicyDestinationsList(ctx, data.Destinations, diags)
	}

	if data.Users != nil {
		result["users"] = data.expandSecurityTrafficShapingPolicyUsersList(ctx, data.Users, diags)
	}

	result["schedule"] = nil
	if data.Schedule != nil && !isZeroStruct(*data.Schedule) {
		result["schedule"] = data.Schedule.expandSecurityTrafficShapingPolicySchedule(ctx, diags)
	}

	if data.Services != nil {
		result["services"] = data.expandSecurityTrafficShapingPolicyServicesList(ctx, data.Services, diags)
	}

	if data.Applications != nil {
		result["applications"] = data.expandSecurityTrafficShapingPolicyApplicationsList(ctx, data.Applications, diags)
	}

	if data.AppCategories != nil {
		result["appCategories"] = data.expandSecurityTrafficShapingPolicyAppCategoriesList(ctx, data.AppCategories, diags)
	}

	if data.UrlCategories != nil {
		result["urlCategories"] = data.expandSecurityTrafficShapingPolicyUrlCategoriesList(ctx, data.UrlCategories, diags)
	}

	result["trafficShaper"] = nil
	if data.TrafficShaper != nil && !isZeroStruct(*data.TrafficShaper) {
		result["trafficShaper"] = data.TrafficShaper.expandSecurityTrafficShapingPolicyTrafficShaper(ctx, diags)
	}

	result["trafficShaperReverse"] = nil
	if data.TrafficShaperReverse != nil && !isZeroStruct(*data.TrafficShaperReverse) {
		result["trafficShaperReverse"] = data.TrafficShaperReverse.expandSecurityTrafficShapingPolicyTrafficShaperReverse(ctx, diags)
	}

	result["perIpShaper"] = nil
	if data.PerIpShaper != nil && !isZeroStruct(*data.PerIpShaper) {
		result["perIpShaper"] = data.PerIpShaper.expandSecurityTrafficShapingPolicyPerIpShaper(ctx, diags)
	}

	return &result
}

func (data *resourceSecurityTrafficShapingPolicyModel) getURLObjectSecurityTrafficShapingPolicy(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityTrafficShapingPolicySourcesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityTrafficShapingPolicyDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityTrafficShapingPolicyUsersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityTrafficShapingPolicyScheduleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityTrafficShapingPolicyServicesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityTrafficShapingPolicyApplicationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityTrafficShapingPolicyAppCategoriesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityTrafficShapingPolicyUrlCategoriesModel struct {
	Id types.Float64 `tfsdk:"id"`
}

type resourceSecurityTrafficShapingPolicyTrafficShaperModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityTrafficShapingPolicyTrafficShaperReverseModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityTrafficShapingPolicyPerIpShaperModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityTrafficShapingPolicySourcesModel) flattenSecurityTrafficShapingPolicySources(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicySourcesModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicySourcesModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicySourcesModel{}
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

func (s *resourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicySourcesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityTrafficShapingPolicySourcesModel {
	if o == nil {
		return []resourceSecurityTrafficShapingPolicySourcesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sources is not type of []interface{}.", "")
		return []resourceSecurityTrafficShapingPolicySourcesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityTrafficShapingPolicySourcesModel{}
	}

	values := make([]resourceSecurityTrafficShapingPolicySourcesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityTrafficShapingPolicySourcesModel
		if i < len(s.Sources) {
			m = s.Sources[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicySources(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityTrafficShapingPolicyDestinationsModel) flattenSecurityTrafficShapingPolicyDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicyDestinationsModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicyDestinationsModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicyDestinationsModel{}
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

func (s *resourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityTrafficShapingPolicyDestinationsModel {
	if o == nil {
		return []resourceSecurityTrafficShapingPolicyDestinationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []resourceSecurityTrafficShapingPolicyDestinationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityTrafficShapingPolicyDestinationsModel{}
	}

	values := make([]resourceSecurityTrafficShapingPolicyDestinationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityTrafficShapingPolicyDestinationsModel
		if i < len(s.Destinations) {
			m = s.Destinations[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyDestinations(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityTrafficShapingPolicyUsersModel) flattenSecurityTrafficShapingPolicyUsers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicyUsersModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicyUsersModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicyUsersModel{}
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

func (s *resourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyUsersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityTrafficShapingPolicyUsersModel {
	if o == nil {
		return []resourceSecurityTrafficShapingPolicyUsersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument users is not type of []interface{}.", "")
		return []resourceSecurityTrafficShapingPolicyUsersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityTrafficShapingPolicyUsersModel{}
	}

	values := make([]resourceSecurityTrafficShapingPolicyUsersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityTrafficShapingPolicyUsersModel
		if i < len(s.Users) {
			m = s.Users[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyUsers(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityTrafficShapingPolicyScheduleModel) flattenSecurityTrafficShapingPolicySchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicyScheduleModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicyScheduleModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicyScheduleModel{}
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

func (m *resourceSecurityTrafficShapingPolicyServicesModel) flattenSecurityTrafficShapingPolicyServices(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicyServicesModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicyServicesModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicyServicesModel{}
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

func (s *resourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyServicesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityTrafficShapingPolicyServicesModel {
	if o == nil {
		return []resourceSecurityTrafficShapingPolicyServicesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument services is not type of []interface{}.", "")
		return []resourceSecurityTrafficShapingPolicyServicesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityTrafficShapingPolicyServicesModel{}
	}

	values := make([]resourceSecurityTrafficShapingPolicyServicesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityTrafficShapingPolicyServicesModel
		if i < len(s.Services) {
			m = s.Services[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyServices(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityTrafficShapingPolicyApplicationsModel) flattenSecurityTrafficShapingPolicyApplications(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicyApplicationsModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicyApplicationsModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicyApplicationsModel{}
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

func (s *resourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyApplicationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityTrafficShapingPolicyApplicationsModel {
	if o == nil {
		return []resourceSecurityTrafficShapingPolicyApplicationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument applications is not type of []interface{}.", "")
		return []resourceSecurityTrafficShapingPolicyApplicationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityTrafficShapingPolicyApplicationsModel{}
	}

	values := make([]resourceSecurityTrafficShapingPolicyApplicationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityTrafficShapingPolicyApplicationsModel
		if i < len(s.Applications) {
			m = s.Applications[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyApplications(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityTrafficShapingPolicyAppCategoriesModel) flattenSecurityTrafficShapingPolicyAppCategories(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicyAppCategoriesModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicyAppCategoriesModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicyAppCategoriesModel{}
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

func (s *resourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyAppCategoriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityTrafficShapingPolicyAppCategoriesModel {
	if o == nil {
		return []resourceSecurityTrafficShapingPolicyAppCategoriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument app_categories is not type of []interface{}.", "")
		return []resourceSecurityTrafficShapingPolicyAppCategoriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityTrafficShapingPolicyAppCategoriesModel{}
	}

	values := make([]resourceSecurityTrafficShapingPolicyAppCategoriesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityTrafficShapingPolicyAppCategoriesModel
		if i < len(s.AppCategories) {
			m = s.AppCategories[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyAppCategories(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityTrafficShapingPolicyUrlCategoriesModel) flattenSecurityTrafficShapingPolicyUrlCategories(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicyUrlCategoriesModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicyUrlCategoriesModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicyUrlCategoriesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	return m
}

func (s *resourceSecurityTrafficShapingPolicyModel) flattenSecurityTrafficShapingPolicyUrlCategoriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityTrafficShapingPolicyUrlCategoriesModel {
	if o == nil {
		return []resourceSecurityTrafficShapingPolicyUrlCategoriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument url_categories is not type of []interface{}.", "")
		return []resourceSecurityTrafficShapingPolicyUrlCategoriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityTrafficShapingPolicyUrlCategoriesModel{}
	}

	values := make([]resourceSecurityTrafficShapingPolicyUrlCategoriesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityTrafficShapingPolicyUrlCategoriesModel
		if i < len(s.UrlCategories) {
			m = s.UrlCategories[i]
		}
		values[i] = *m.flattenSecurityTrafficShapingPolicyUrlCategories(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityTrafficShapingPolicyTrafficShaperModel) flattenSecurityTrafficShapingPolicyTrafficShaper(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicyTrafficShaperModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicyTrafficShaperModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicyTrafficShaperModel{}
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

func (m *resourceSecurityTrafficShapingPolicyTrafficShaperReverseModel) flattenSecurityTrafficShapingPolicyTrafficShaperReverse(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicyTrafficShaperReverseModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicyTrafficShaperReverseModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicyTrafficShaperReverseModel{}
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

func (m *resourceSecurityTrafficShapingPolicyPerIpShaperModel) flattenSecurityTrafficShapingPolicyPerIpShaper(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityTrafficShapingPolicyPerIpShaperModel {
	if input == nil {
		return &resourceSecurityTrafficShapingPolicyPerIpShaperModel{}
	}
	if m == nil {
		m = &resourceSecurityTrafficShapingPolicyPerIpShaperModel{}
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

func (data *resourceSecurityTrafficShapingPolicySourcesModel) expandSecurityTrafficShapingPolicySources(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityTrafficShapingPolicyModel) expandSecurityTrafficShapingPolicySourcesList(ctx context.Context, l []resourceSecurityTrafficShapingPolicySourcesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityTrafficShapingPolicySources(ctx, diags)
	}
	return result
}

func (data *resourceSecurityTrafficShapingPolicyDestinationsModel) expandSecurityTrafficShapingPolicyDestinations(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityTrafficShapingPolicyModel) expandSecurityTrafficShapingPolicyDestinationsList(ctx context.Context, l []resourceSecurityTrafficShapingPolicyDestinationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityTrafficShapingPolicyDestinations(ctx, diags)
	}
	return result
}

func (data *resourceSecurityTrafficShapingPolicyUsersModel) expandSecurityTrafficShapingPolicyUsers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityTrafficShapingPolicyModel) expandSecurityTrafficShapingPolicyUsersList(ctx context.Context, l []resourceSecurityTrafficShapingPolicyUsersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityTrafficShapingPolicyUsers(ctx, diags)
	}
	return result
}

func (data *resourceSecurityTrafficShapingPolicyScheduleModel) expandSecurityTrafficShapingPolicySchedule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityTrafficShapingPolicyServicesModel) expandSecurityTrafficShapingPolicyServices(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityTrafficShapingPolicyModel) expandSecurityTrafficShapingPolicyServicesList(ctx context.Context, l []resourceSecurityTrafficShapingPolicyServicesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityTrafficShapingPolicyServices(ctx, diags)
	}
	return result
}

func (data *resourceSecurityTrafficShapingPolicyApplicationsModel) expandSecurityTrafficShapingPolicyApplications(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityTrafficShapingPolicyModel) expandSecurityTrafficShapingPolicyApplicationsList(ctx context.Context, l []resourceSecurityTrafficShapingPolicyApplicationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityTrafficShapingPolicyApplications(ctx, diags)
	}
	return result
}

func (data *resourceSecurityTrafficShapingPolicyAppCategoriesModel) expandSecurityTrafficShapingPolicyAppCategories(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityTrafficShapingPolicyModel) expandSecurityTrafficShapingPolicyAppCategoriesList(ctx context.Context, l []resourceSecurityTrafficShapingPolicyAppCategoriesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityTrafficShapingPolicyAppCategories(ctx, diags)
	}
	return result
}

func (data *resourceSecurityTrafficShapingPolicyUrlCategoriesModel) expandSecurityTrafficShapingPolicyUrlCategories(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		result["id"] = data.Id.ValueFloat64()
	}

	return result
}

func (s *resourceSecurityTrafficShapingPolicyModel) expandSecurityTrafficShapingPolicyUrlCategoriesList(ctx context.Context, l []resourceSecurityTrafficShapingPolicyUrlCategoriesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityTrafficShapingPolicyUrlCategories(ctx, diags)
	}
	return result
}

func (data *resourceSecurityTrafficShapingPolicyTrafficShaperModel) expandSecurityTrafficShapingPolicyTrafficShaper(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityTrafficShapingPolicyTrafficShaperReverseModel) expandSecurityTrafficShapingPolicyTrafficShaperReverse(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityTrafficShapingPolicyPerIpShaperModel) expandSecurityTrafficShapingPolicyPerIpShaper(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
