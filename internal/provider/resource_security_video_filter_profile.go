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
	"strings"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSecurityVideoFilterProfile{}

func newResourceSecurityVideoFilterProfile() resource.Resource {
	return &resourceSecurityVideoFilterProfile{}
}

type resourceSecurityVideoFilterProfile struct {
	fortiClient *FortiClient
}

// resourceSecurityVideoFilterProfileModel describes the resource data model.
type resourceSecurityVideoFilterProfileModel struct {
	ID                types.String                                               `tfsdk:"id"`
	PrimaryKey        types.String                                               `tfsdk:"primary_key"`
	FortiguardFilters []resourceSecurityVideoFilterProfileFortiguardFiltersModel `tfsdk:"fortiguard_filters"`
	DefaultAction     types.String                                               `tfsdk:"default_action"`
	Channels          []resourceSecurityVideoFilterProfileChannelsModel          `tfsdk:"channels"`
	Direction         types.String                                               `tfsdk:"direction"`
}

func (r *resourceSecurityVideoFilterProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_video_filter_profile"
}

func (r *resourceSecurityVideoFilterProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Required: true,
			},
			"default_action": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"direction": schema.StringAttribute{
				Description: "The direction of the target resource.",
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				Computed: true,
				Optional: true,
			},
			"fortiguard_filters": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "monitor", "block", "warning", "default"),
							},
							Computed: true,
							Optional: true,
						},
						"category": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("security/video-filter-fortiguard-categories"),
									},
									Computed: true,
									Optional: true,
								},
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"channels": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "monitor", "block"),
							},
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
							Computed: true,
							Optional: true,
						},
						"channel_id": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 64),
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

func (r *resourceSecurityVideoFilterProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityVideoFilterProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityVideoFilterProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecurityVideoFilterProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityVideoFilterProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecurityVideoFilterProfile(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityVideoFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityVideoFilterProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityVideoFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityVideoFilterProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityVideoFilterProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityVideoFilterProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityVideoFilterProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityVideoFilterProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityVideoFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityVideoFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityVideoFilterProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityVideoFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityVideoFilterProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityVideoFilterProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityVideoFilterProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityVideoFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityVideoFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityVideoFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityVideoFilterProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected format: direction/primary_key, got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("direction"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), parts[1])...)
}

func (m *resourceSecurityVideoFilterProfileModel) refreshSecurityVideoFilterProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["fortiguardFilters"]; ok {
		m.FortiguardFilters = m.flattenSecurityVideoFilterProfileFortiguardFiltersList(ctx, v, &diags)
	}

	if v, ok := o["defaultAction"]; ok {
		m.DefaultAction = parseStringValue(v)
	}

	if v, ok := o["channels"]; ok {
		m.Channels = m.flattenSecurityVideoFilterProfileChannelsList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityVideoFilterProfileModel) getCreateObjectSecurityVideoFilterProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	result["fortiguardFilters"] = data.expandSecurityVideoFilterProfileFortiguardFiltersList(ctx, data.FortiguardFilters, diags)

	if !data.DefaultAction.IsNull() {
		result["defaultAction"] = data.DefaultAction.ValueString()
	}

	result["channels"] = data.expandSecurityVideoFilterProfileChannelsList(ctx, data.Channels, diags)

	return &result
}

func (data *resourceSecurityVideoFilterProfileModel) getUpdateObjectSecurityVideoFilterProfile(ctx context.Context, state resourceSecurityVideoFilterProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if len(data.FortiguardFilters) > 0 || !isSameStruct(data.FortiguardFilters, state.FortiguardFilters) {
		result["fortiguardFilters"] = data.expandSecurityVideoFilterProfileFortiguardFiltersList(ctx, data.FortiguardFilters, diags)
	}

	if !data.DefaultAction.IsNull() {
		result["defaultAction"] = data.DefaultAction.ValueString()
	}

	if len(data.Channels) > 0 || !isSameStruct(data.Channels, state.Channels) {
		result["channels"] = data.expandSecurityVideoFilterProfileChannelsList(ctx, data.Channels, diags)
	}

	return &result
}

func (data *resourceSecurityVideoFilterProfileModel) getURLObjectSecurityVideoFilterProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityVideoFilterProfileFortiguardFiltersModel struct {
	Action   types.String                                                      `tfsdk:"action"`
	Category *resourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel `tfsdk:"category"`
}

type resourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityVideoFilterProfileChannelsModel struct {
	Action    types.String `tfsdk:"action"`
	Name      types.String `tfsdk:"name"`
	ChannelId types.String `tfsdk:"channel_id"`
}

func (m *resourceSecurityVideoFilterProfileFortiguardFiltersModel) flattenSecurityVideoFilterProfileFortiguardFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityVideoFilterProfileFortiguardFiltersModel {
	if input == nil {
		return &resourceSecurityVideoFilterProfileFortiguardFiltersModel{}
	}
	if m == nil {
		m = &resourceSecurityVideoFilterProfileFortiguardFiltersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = m.Category.flattenSecurityVideoFilterProfileFortiguardFiltersCategory(ctx, v, diags)
	}

	return m
}

func (s *resourceSecurityVideoFilterProfileModel) flattenSecurityVideoFilterProfileFortiguardFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityVideoFilterProfileFortiguardFiltersModel {
	if o == nil {
		return []resourceSecurityVideoFilterProfileFortiguardFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument fortiguard_filters is not type of []interface{}.", "")
		return []resourceSecurityVideoFilterProfileFortiguardFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityVideoFilterProfileFortiguardFiltersModel{}
	}

	values := make([]resourceSecurityVideoFilterProfileFortiguardFiltersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityVideoFilterProfileFortiguardFiltersModel
		values[i] = *m.flattenSecurityVideoFilterProfileFortiguardFilters(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel) flattenSecurityVideoFilterProfileFortiguardFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel {
	if input == nil {
		return &resourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel{}
	}
	if m == nil {
		m = &resourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel{}
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

func (m *resourceSecurityVideoFilterProfileChannelsModel) flattenSecurityVideoFilterProfileChannels(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityVideoFilterProfileChannelsModel {
	if input == nil {
		return &resourceSecurityVideoFilterProfileChannelsModel{}
	}
	if m == nil {
		m = &resourceSecurityVideoFilterProfileChannelsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["channelId"]; ok {
		m.ChannelId = parseStringValue(v)
	}

	return m
}

func (s *resourceSecurityVideoFilterProfileModel) flattenSecurityVideoFilterProfileChannelsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityVideoFilterProfileChannelsModel {
	if o == nil {
		return []resourceSecurityVideoFilterProfileChannelsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument channels is not type of []interface{}.", "")
		return []resourceSecurityVideoFilterProfileChannelsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityVideoFilterProfileChannelsModel{}
	}

	values := make([]resourceSecurityVideoFilterProfileChannelsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityVideoFilterProfileChannelsModel
		values[i] = *m.flattenSecurityVideoFilterProfileChannels(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityVideoFilterProfileFortiguardFiltersModel) expandSecurityVideoFilterProfileFortiguardFilters(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Category != nil && !isZeroStruct(*data.Category) {
		result["category"] = data.Category.expandSecurityVideoFilterProfileFortiguardFiltersCategory(ctx, diags)
	}

	return result
}

func (s *resourceSecurityVideoFilterProfileModel) expandSecurityVideoFilterProfileFortiguardFiltersList(ctx context.Context, l []resourceSecurityVideoFilterProfileFortiguardFiltersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityVideoFilterProfileFortiguardFilters(ctx, diags)
	}
	return result
}

func (data *resourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel) expandSecurityVideoFilterProfileFortiguardFiltersCategory(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityVideoFilterProfileChannelsModel) expandSecurityVideoFilterProfileChannels(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Name.IsNull() {
		result["name"] = data.Name.ValueString()
	}

	if !data.ChannelId.IsNull() {
		result["channelId"] = data.ChannelId.ValueString()
	}

	return result
}

func (s *resourceSecurityVideoFilterProfileModel) expandSecurityVideoFilterProfileChannelsList(ctx context.Context, l []resourceSecurityVideoFilterProfileChannelsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityVideoFilterProfileChannels(ctx, diags)
	}
	return result
}
