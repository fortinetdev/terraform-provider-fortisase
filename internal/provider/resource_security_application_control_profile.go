// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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
var _ resource.Resource = &resourceSecurityApplicationControlProfile{}

func newResourceSecurityApplicationControlProfile() resource.Resource {
	return &resourceSecurityApplicationControlProfile{}
}

type resourceSecurityApplicationControlProfile struct {
	fortiClient *FortiClient
}

// resourceSecurityApplicationControlProfileModel describes the resource data model.
type resourceSecurityApplicationControlProfileModel struct {
	ID                              types.String                                                                `tfsdk:"id"`
	PrimaryKey                      types.String                                                                `tfsdk:"primary_key"`
	ApplicationCategoryControls     []resourceSecurityApplicationControlProfileApplicationCategoryControlsModel `tfsdk:"application_category_controls"`
	ApplicationControls             []resourceSecurityApplicationControlProfileApplicationControlsModel         `tfsdk:"application_controls"`
	UnknownApplicationAction        types.String                                                                `tfsdk:"unknown_application_action"`
	NetworkProtocolEnforcement      types.String                                                                `tfsdk:"network_protocol_enforcement"`
	NetworkProtocols                []resourceSecurityApplicationControlProfileNetworkProtocolsModel            `tfsdk:"network_protocols"`
	BlockNonDefaultPortApplications types.String                                                                `tfsdk:"block_non_default_port_applications"`
	Direction                       types.String                                                                `tfsdk:"direction"`
}

func (r *resourceSecurityApplicationControlProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_application_control_profile"
}

func (r *resourceSecurityApplicationControlProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"unknown_application_action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("block", "allow", "monitor"),
				},
				Computed: true,
				Optional: true,
			},
			"network_protocol_enforcement": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"block_non_default_port_applications": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
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
			"application_category_controls": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "monitor", "block"),
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
										stringvalidator.OneOf("security/application-categories"),
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
			"application_controls": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("monitor", "allow", "block"),
							},
							Computed: true,
							Optional: true,
						},
						"applications": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("security/applications"),
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
				},
				Computed: true,
				Optional: true,
			},
			"network_protocols": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validator.Between(1, 65535),
							},
							Computed: true,
							Optional: true,
						},
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("monitor", "pass", "block"),
							},
							Computed: true,
							Optional: true,
						},
						"services": schema.SetAttribute{
							Validators: []validator.Set{
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf("dns", "ftp", "http", "https", "imap", "nntp", "pop3", "smtp", "snmp", "ssh", "telnet"),
								),
								setvalidator.SizeAtLeast(1),
							},
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityApplicationControlProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityApplicationControlProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityApplicationControlProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecurityApplicationControlProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityApplicationControlProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecurityApplicationControlProfile(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityApplicationControlProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityApplicationControlProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityApplicationControlProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityApplicationControlProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityApplicationControlProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityApplicationControlProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityApplicationControlProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityApplicationControlProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityApplicationControlProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityApplicationControlProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityApplicationControlProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityApplicationControlProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityApplicationControlProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityApplicationControlProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityApplicationControlProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityApplicationControlProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityApplicationControlProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityApplicationControlProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityApplicationControlProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (m *resourceSecurityApplicationControlProfileModel) refreshSecurityApplicationControlProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["applicationCategoryControls"]; ok {
		m.ApplicationCategoryControls = m.flattenSecurityApplicationControlProfileApplicationCategoryControlsList(ctx, v, &diags)
	}

	if v, ok := o["applicationControls"]; ok {
		m.ApplicationControls = m.flattenSecurityApplicationControlProfileApplicationControlsList(ctx, v, &diags)
	}

	if v, ok := o["unknownApplicationAction"]; ok {
		m.UnknownApplicationAction = parseStringValue(v)
	}

	if v, ok := o["networkProtocolEnforcement"]; ok {
		m.NetworkProtocolEnforcement = parseStringValue(v)
	}

	if v, ok := o["networkProtocols"]; ok {
		m.NetworkProtocols = m.flattenSecurityApplicationControlProfileNetworkProtocolsList(ctx, v, &diags)
	}

	if v, ok := o["blockNonDefaultPortApplications"]; ok {
		m.BlockNonDefaultPortApplications = parseStringValue(v)
	}

	return diags
}

func (data *resourceSecurityApplicationControlProfileModel) getCreateObjectSecurityApplicationControlProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	result["applicationCategoryControls"] = data.expandSecurityApplicationControlProfileApplicationCategoryControlsList(ctx, data.ApplicationCategoryControls, diags)

	result["applicationControls"] = data.expandSecurityApplicationControlProfileApplicationControlsList(ctx, data.ApplicationControls, diags)

	if !data.UnknownApplicationAction.IsNull() {
		result["unknownApplicationAction"] = data.UnknownApplicationAction.ValueString()
	}

	if !data.NetworkProtocolEnforcement.IsNull() {
		result["networkProtocolEnforcement"] = data.NetworkProtocolEnforcement.ValueString()
	}

	result["networkProtocols"] = data.expandSecurityApplicationControlProfileNetworkProtocolsList(ctx, data.NetworkProtocols, diags)

	if !data.BlockNonDefaultPortApplications.IsNull() {
		result["blockNonDefaultPortApplications"] = data.BlockNonDefaultPortApplications.ValueString()
	}

	return &result
}

func (data *resourceSecurityApplicationControlProfileModel) getUpdateObjectSecurityApplicationControlProfile(ctx context.Context, state resourceSecurityApplicationControlProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if len(data.ApplicationCategoryControls) > 0 || !isSameStruct(data.ApplicationCategoryControls, state.ApplicationCategoryControls) {
		result["applicationCategoryControls"] = data.expandSecurityApplicationControlProfileApplicationCategoryControlsList(ctx, data.ApplicationCategoryControls, diags)
	}

	if len(data.ApplicationControls) > 0 || !isSameStruct(data.ApplicationControls, state.ApplicationControls) {
		result["applicationControls"] = data.expandSecurityApplicationControlProfileApplicationControlsList(ctx, data.ApplicationControls, diags)
	}

	if !data.UnknownApplicationAction.IsNull() {
		result["unknownApplicationAction"] = data.UnknownApplicationAction.ValueString()
	}

	if !data.NetworkProtocolEnforcement.IsNull() {
		result["networkProtocolEnforcement"] = data.NetworkProtocolEnforcement.ValueString()
	}

	if len(data.NetworkProtocols) > 0 || !isSameStruct(data.NetworkProtocols, state.NetworkProtocols) {
		result["networkProtocols"] = data.expandSecurityApplicationControlProfileNetworkProtocolsList(ctx, data.NetworkProtocols, diags)
	}

	if !data.BlockNonDefaultPortApplications.IsNull() {
		result["blockNonDefaultPortApplications"] = data.BlockNonDefaultPortApplications.ValueString()
	}

	return &result
}

func (data *resourceSecurityApplicationControlProfileModel) getURLObjectSecurityApplicationControlProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityApplicationControlProfileApplicationCategoryControlsModel struct {
	Action   types.String                                                                       `tfsdk:"action"`
	Category *resourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel `tfsdk:"category"`
}

type resourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityApplicationControlProfileApplicationControlsModel struct {
	Action       types.String                                                                    `tfsdk:"action"`
	Applications []resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel `tfsdk:"applications"`
}

type resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityApplicationControlProfileNetworkProtocolsModel struct {
	Port     types.Float64 `tfsdk:"port"`
	Action   types.String  `tfsdk:"action"`
	Services types.Set     `tfsdk:"services"`
}

func (m *resourceSecurityApplicationControlProfileApplicationCategoryControlsModel) flattenSecurityApplicationControlProfileApplicationCategoryControls(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityApplicationControlProfileApplicationCategoryControlsModel {
	if input == nil {
		return &resourceSecurityApplicationControlProfileApplicationCategoryControlsModel{}
	}
	if m == nil {
		m = &resourceSecurityApplicationControlProfileApplicationCategoryControlsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = m.Category.flattenSecurityApplicationControlProfileApplicationCategoryControlsCategory(ctx, v, diags)
	}

	return m
}

func (s *resourceSecurityApplicationControlProfileModel) flattenSecurityApplicationControlProfileApplicationCategoryControlsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityApplicationControlProfileApplicationCategoryControlsModel {
	if o == nil {
		return []resourceSecurityApplicationControlProfileApplicationCategoryControlsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument application_category_controls is not type of []interface{}.", "")
		return []resourceSecurityApplicationControlProfileApplicationCategoryControlsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityApplicationControlProfileApplicationCategoryControlsModel{}
	}

	values := make([]resourceSecurityApplicationControlProfileApplicationCategoryControlsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityApplicationControlProfileApplicationCategoryControlsModel
		values[i] = *m.flattenSecurityApplicationControlProfileApplicationCategoryControls(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel) flattenSecurityApplicationControlProfileApplicationCategoryControlsCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel {
	if input == nil {
		return &resourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel{}
	}
	if m == nil {
		m = &resourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel{}
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

func (m *resourceSecurityApplicationControlProfileApplicationControlsModel) flattenSecurityApplicationControlProfileApplicationControls(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityApplicationControlProfileApplicationControlsModel {
	if input == nil {
		return &resourceSecurityApplicationControlProfileApplicationControlsModel{}
	}
	if m == nil {
		m = &resourceSecurityApplicationControlProfileApplicationControlsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["applications"]; ok {
		m.Applications = m.flattenSecurityApplicationControlProfileApplicationControlsApplicationsList(ctx, v, diags)
	}

	return m
}

func (s *resourceSecurityApplicationControlProfileModel) flattenSecurityApplicationControlProfileApplicationControlsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityApplicationControlProfileApplicationControlsModel {
	if o == nil {
		return []resourceSecurityApplicationControlProfileApplicationControlsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument application_controls is not type of []interface{}.", "")
		return []resourceSecurityApplicationControlProfileApplicationControlsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityApplicationControlProfileApplicationControlsModel{}
	}

	values := make([]resourceSecurityApplicationControlProfileApplicationControlsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityApplicationControlProfileApplicationControlsModel
		values[i] = *m.flattenSecurityApplicationControlProfileApplicationControls(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel) flattenSecurityApplicationControlProfileApplicationControlsApplications(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel {
	if input == nil {
		return &resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel{}
	}
	if m == nil {
		m = &resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel{}
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

func (s *resourceSecurityApplicationControlProfileApplicationControlsModel) flattenSecurityApplicationControlProfileApplicationControlsApplicationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel {
	if o == nil {
		return []resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument applications is not type of []interface{}.", "")
		return []resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel{}
	}

	values := make([]resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel
		values[i] = *m.flattenSecurityApplicationControlProfileApplicationControlsApplications(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityApplicationControlProfileNetworkProtocolsModel) flattenSecurityApplicationControlProfileNetworkProtocols(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityApplicationControlProfileNetworkProtocolsModel {
	if input == nil {
		return &resourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}
	if m == nil {
		m = &resourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}
	o := input.(map[string]interface{})
	m.Services = types.SetNull(types.StringType)

	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["services"]; ok {
		m.Services = parseSetValue(ctx, v, types.StringType)
	}

	return m
}

func (s *resourceSecurityApplicationControlProfileModel) flattenSecurityApplicationControlProfileNetworkProtocolsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityApplicationControlProfileNetworkProtocolsModel {
	if o == nil {
		return []resourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument network_protocols is not type of []interface{}.", "")
		return []resourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	values := make([]resourceSecurityApplicationControlProfileNetworkProtocolsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityApplicationControlProfileNetworkProtocolsModel
		values[i] = *m.flattenSecurityApplicationControlProfileNetworkProtocols(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityApplicationControlProfileApplicationCategoryControlsModel) expandSecurityApplicationControlProfileApplicationCategoryControls(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Category != nil && !isZeroStruct(*data.Category) {
		result["category"] = data.Category.expandSecurityApplicationControlProfileApplicationCategoryControlsCategory(ctx, diags)
	}

	return result
}

func (s *resourceSecurityApplicationControlProfileModel) expandSecurityApplicationControlProfileApplicationCategoryControlsList(ctx context.Context, l []resourceSecurityApplicationControlProfileApplicationCategoryControlsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityApplicationControlProfileApplicationCategoryControls(ctx, diags)
	}
	return result
}

func (data *resourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel) expandSecurityApplicationControlProfileApplicationCategoryControlsCategory(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityApplicationControlProfileApplicationControlsModel) expandSecurityApplicationControlProfileApplicationControls(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	result["applications"] = data.expandSecurityApplicationControlProfileApplicationControlsApplicationsList(ctx, data.Applications, diags)

	return result
}

func (s *resourceSecurityApplicationControlProfileModel) expandSecurityApplicationControlProfileApplicationControlsList(ctx context.Context, l []resourceSecurityApplicationControlProfileApplicationControlsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityApplicationControlProfileApplicationControls(ctx, diags)
	}
	return result
}

func (data *resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel) expandSecurityApplicationControlProfileApplicationControlsApplications(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityApplicationControlProfileApplicationControlsModel) expandSecurityApplicationControlProfileApplicationControlsApplicationsList(ctx context.Context, l []resourceSecurityApplicationControlProfileApplicationControlsApplicationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityApplicationControlProfileApplicationControlsApplications(ctx, diags)
	}
	return result
}

func (data *resourceSecurityApplicationControlProfileNetworkProtocolsModel) expandSecurityApplicationControlProfileNetworkProtocols(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Port.IsNull() {
		result["port"] = data.Port.ValueFloat64()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Services.IsNull() {
		result["services"] = expandSetToStringList(data.Services)
	}

	return result
}

func (s *resourceSecurityApplicationControlProfileModel) expandSecurityApplicationControlProfileNetworkProtocolsList(ctx context.Context, l []resourceSecurityApplicationControlProfileNetworkProtocolsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityApplicationControlProfileNetworkProtocols(ctx, diags)
	}
	return result
}
