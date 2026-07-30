// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/setvalidatorwarning"
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
var _ resource.Resource = &resourceSecurityApplicationControlProfile{}

func newResourceSecurityApplicationControlProfile() resource.Resource {
	return &resourceSecurityApplicationControlProfile{}
}

type resourceSecurityApplicationControlProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityApplicationControlProfileModel describes the resource data model.
type resourceSecurityApplicationControlProfileModel struct {
	ID                              types.String                                                     `tfsdk:"id"`
	PrimaryKey                      types.String                                                     `tfsdk:"primary_key"`
	Controls                        []resourceSecurityApplicationControlProfileControlsModel         `tfsdk:"controls"`
	UnknownApplicationAction        types.String                                                     `tfsdk:"unknown_application_action"`
	NetworkProtocolEnforcement      types.String                                                     `tfsdk:"network_protocol_enforcement"`
	NetworkProtocols                []resourceSecurityApplicationControlProfileNetworkProtocolsModel `tfsdk:"network_protocols"`
	BlockNonDefaultPortApplications types.String                                                     `tfsdk:"block_non_default_port_applications"`
	Direction                       types.String                                                     `tfsdk:"direction"`
}

func (r *resourceSecurityApplicationControlProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_application_control_profile"
}

func (r *resourceSecurityApplicationControlProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Application Control Profile Resource API V2 for FortiSASE.",
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"unknown_application_action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("block", "allow", "monitor"),
				},
				Computed: true,
				Optional: true,
			},
			"network_protocol_enforcement": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"block_non_default_port_applications": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("internal-profiles", "outbound-profiles"),
				},
				MarkdownDescription: "The direction of the target resource.\nSupported values: internal-profiles, outbound-profiles.",
				Computed:            true,
				Optional:            true,
			},
			"controls": schema.ListNestedAttribute{
				MarkdownDescription: "Generic controls defining actions for applications, filters, and overrides. Array order matters. Entries are evaluated first-to-last. Overrides must be placed ahead of application category controls for correct evaluation.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("monitor", "allow", "block"),
							},
							Computed: true,
							Optional: true,
						},
						"risk": schema.SetAttribute{
							MarkdownDescription: "Risk level(s) with 0 being lowest and 4 being highest",
							Computed:            true,
							Optional:            true,
							ElementType:         types.Int64Type,
						},
						"popularity": schema.SetAttribute{
							MarkdownDescription: "Popularity level(s) with 1 being lowest and 5 being highest",
							Computed:            true,
							Optional:            true,
							ElementType:         types.Int64Type,
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
						"categories": schema.ListNestedAttribute{
							MarkdownDescription: "Set the control action for a given application category. For the 'Proxy' category, the action cannot be set to 'block' if the linked profile group is referenced by a proxy policy.",
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
						"ips_attributes": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Optional: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidatorwarning.OneOf("security/ips-attr-map"),
										},
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
								float64validatorwarning.Between(1, 65535),
							},
							Computed: true,
							Optional: true,
						},
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("monitor", "pass", "block"),
							},
							Computed: true,
							Optional: true,
						},
						"services": schema.SetAttribute{
							Validators: []validator.Set{
								setvalidatorwarning.ValueStringsAre(
									stringvalidatorwarning.OneOf("dns", "ftp", "http", "https", "imap", "nntp", "pop3", "smtp", "snmp", "ssh", "telnet"),
								),
								setvalidatorwarning.SizeAtLeast(1),
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
	r.resourceName = "fortisase_security_application_control_profile"
}

func (r *resourceSecurityApplicationControlProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityApplicationControlProfile")
	lock.Lock()
	defer lock.Unlock()
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
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := data.PrimaryKey.ValueString()
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityApplicationControlProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityApplicationControlProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
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
	lock := r.fortiClient.GetResourceLock("SecurityApplicationControlProfile")
	lock.Lock()
	defer lock.Unlock()
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

	output, err := c.UpdateSecurityApplicationControlProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityApplicationControlProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityApplicationControlProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
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

	diags.Append(data.refreshSecurityApplicationControlProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityApplicationControlProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityApplicationControlProfileModel) refreshSecurityApplicationControlProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["controls"]; ok {
		m.Controls = m.flattenSecurityApplicationControlProfileControlsList(ctx, v, &diags)
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
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if data.Controls != nil {
		result["controls"] = data.expandSecurityApplicationControlProfileControlsList(ctx, data.Controls, diags)
	}

	if !data.UnknownApplicationAction.IsNull() && !data.UnknownApplicationAction.IsUnknown() {
		result["unknownApplicationAction"] = data.UnknownApplicationAction.ValueString()
	}

	if !data.NetworkProtocolEnforcement.IsNull() && !data.NetworkProtocolEnforcement.IsUnknown() {
		result["networkProtocolEnforcement"] = data.NetworkProtocolEnforcement.ValueString()
	}

	if data.NetworkProtocols != nil {
		result["networkProtocols"] = data.expandSecurityApplicationControlProfileNetworkProtocolsList(ctx, data.NetworkProtocols, diags)
	}

	if !data.BlockNonDefaultPortApplications.IsNull() && !data.BlockNonDefaultPortApplications.IsUnknown() {
		result["blockNonDefaultPortApplications"] = data.BlockNonDefaultPortApplications.ValueString()
	}

	return &result
}

func (data *resourceSecurityApplicationControlProfileModel) getUpdateObjectSecurityApplicationControlProfile(ctx context.Context, state resourceSecurityApplicationControlProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if data.Controls != nil {
		result["controls"] = data.expandSecurityApplicationControlProfileControlsList(ctx, data.Controls, diags)
	}

	if !data.UnknownApplicationAction.IsNull() && !data.UnknownApplicationAction.IsUnknown() {
		result["unknownApplicationAction"] = data.UnknownApplicationAction.ValueString()
	}

	if !data.NetworkProtocolEnforcement.IsNull() && !data.NetworkProtocolEnforcement.IsUnknown() {
		result["networkProtocolEnforcement"] = data.NetworkProtocolEnforcement.ValueString()
	}

	if data.NetworkProtocols != nil {
		result["networkProtocols"] = data.expandSecurityApplicationControlProfileNetworkProtocolsList(ctx, data.NetworkProtocols, diags)
	}

	if !data.BlockNonDefaultPortApplications.IsNull() && !data.BlockNonDefaultPortApplications.IsUnknown() {
		result["blockNonDefaultPortApplications"] = data.BlockNonDefaultPortApplications.ValueString()
	}

	return &result
}

func (data *resourceSecurityApplicationControlProfileModel) getURLObjectSecurityApplicationControlProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() && !data.Direction.IsUnknown() {
		diags.AddWarning("\"direction\" is deprecated and may be removed in future.",
			"It is recommended to recreate the resource without \"direction\" to avoid unexpected behavior in future.",
		)
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityApplicationControlProfileControlsModel struct {
	Action        types.String                                                          `tfsdk:"action"`
	Applications  []resourceSecurityApplicationControlProfileControlsApplicationsModel  `tfsdk:"applications"`
	Categories    []resourceSecurityApplicationControlProfileControlsCategoriesModel    `tfsdk:"categories"`
	Risk          types.Set                                                             `tfsdk:"risk"`
	Popularity    types.Set                                                             `tfsdk:"popularity"`
	IpsAttributes []resourceSecurityApplicationControlProfileControlsIpsAttributesModel `tfsdk:"ips_attributes"`
}

type resourceSecurityApplicationControlProfileControlsApplicationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityApplicationControlProfileControlsCategoriesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityApplicationControlProfileControlsIpsAttributesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityApplicationControlProfileNetworkProtocolsModel struct {
	Port     types.Float64 `tfsdk:"port"`
	Action   types.String  `tfsdk:"action"`
	Services types.Set     `tfsdk:"services"`
}

func (m *resourceSecurityApplicationControlProfileControlsModel) flattenSecurityApplicationControlProfileControls(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityApplicationControlProfileControlsModel {
	if input == nil {
		return &resourceSecurityApplicationControlProfileControlsModel{}
	}
	if m == nil {
		m = &resourceSecurityApplicationControlProfileControlsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["applications"]; ok {
		m.Applications = m.flattenSecurityApplicationControlProfileControlsApplicationsList(ctx, v, diags)
	}

	if v, ok := o["categories"]; ok {
		m.Categories = m.flattenSecurityApplicationControlProfileControlsCategoriesList(ctx, v, diags)
	}

	if v, ok := o["risk"]; ok {
		m.Risk = parseSetValue(ctx, v, types.Int64Type)
	} else {
		m.Risk = types.SetNull(types.Int64Type)
	}

	if v, ok := o["popularity"]; ok {
		m.Popularity = parseSetValue(ctx, v, types.Int64Type)
	} else {
		m.Popularity = types.SetNull(types.Int64Type)
	}

	if v, ok := o["ipsAttributes"]; ok {
		m.IpsAttributes = m.flattenSecurityApplicationControlProfileControlsIpsAttributesList(ctx, v, diags)
	}

	return m
}

func (s *resourceSecurityApplicationControlProfileModel) flattenSecurityApplicationControlProfileControlsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityApplicationControlProfileControlsModel {
	if o == nil {
		return []resourceSecurityApplicationControlProfileControlsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument controls is not type of []interface{}.", "")
		return []resourceSecurityApplicationControlProfileControlsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityApplicationControlProfileControlsModel{}
	}

	values := make([]resourceSecurityApplicationControlProfileControlsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityApplicationControlProfileControlsModel
		if i < len(s.Controls) {
			m = s.Controls[i]
		}
		values[i] = *m.flattenSecurityApplicationControlProfileControls(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityApplicationControlProfileControlsApplicationsModel) flattenSecurityApplicationControlProfileControlsApplications(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityApplicationControlProfileControlsApplicationsModel {
	if input == nil {
		return &resourceSecurityApplicationControlProfileControlsApplicationsModel{}
	}
	if m == nil {
		m = &resourceSecurityApplicationControlProfileControlsApplicationsModel{}
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

func (s *resourceSecurityApplicationControlProfileControlsModel) flattenSecurityApplicationControlProfileControlsApplicationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityApplicationControlProfileControlsApplicationsModel {
	if o == nil {
		return []resourceSecurityApplicationControlProfileControlsApplicationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument applications is not type of []interface{}.", "")
		return []resourceSecurityApplicationControlProfileControlsApplicationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityApplicationControlProfileControlsApplicationsModel{}
	}

	values := make([]resourceSecurityApplicationControlProfileControlsApplicationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityApplicationControlProfileControlsApplicationsModel
		if i < len(s.Applications) {
			m = s.Applications[i]
		}
		values[i] = *m.flattenSecurityApplicationControlProfileControlsApplications(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityApplicationControlProfileControlsCategoriesModel) flattenSecurityApplicationControlProfileControlsCategories(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityApplicationControlProfileControlsCategoriesModel {
	if input == nil {
		return &resourceSecurityApplicationControlProfileControlsCategoriesModel{}
	}
	if m == nil {
		m = &resourceSecurityApplicationControlProfileControlsCategoriesModel{}
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

func (s *resourceSecurityApplicationControlProfileControlsModel) flattenSecurityApplicationControlProfileControlsCategoriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityApplicationControlProfileControlsCategoriesModel {
	if o == nil {
		return []resourceSecurityApplicationControlProfileControlsCategoriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument categories is not type of []interface{}.", "")
		return []resourceSecurityApplicationControlProfileControlsCategoriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityApplicationControlProfileControlsCategoriesModel{}
	}

	values := make([]resourceSecurityApplicationControlProfileControlsCategoriesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityApplicationControlProfileControlsCategoriesModel
		if i < len(s.Categories) {
			m = s.Categories[i]
		}
		values[i] = *m.flattenSecurityApplicationControlProfileControlsCategories(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityApplicationControlProfileControlsIpsAttributesModel) flattenSecurityApplicationControlProfileControlsIpsAttributes(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityApplicationControlProfileControlsIpsAttributesModel {
	if input == nil {
		return &resourceSecurityApplicationControlProfileControlsIpsAttributesModel{}
	}
	if m == nil {
		m = &resourceSecurityApplicationControlProfileControlsIpsAttributesModel{}
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

func (s *resourceSecurityApplicationControlProfileControlsModel) flattenSecurityApplicationControlProfileControlsIpsAttributesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityApplicationControlProfileControlsIpsAttributesModel {
	if o == nil {
		return []resourceSecurityApplicationControlProfileControlsIpsAttributesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument ips_attributes is not type of []interface{}.", "")
		return []resourceSecurityApplicationControlProfileControlsIpsAttributesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityApplicationControlProfileControlsIpsAttributesModel{}
	}

	values := make([]resourceSecurityApplicationControlProfileControlsIpsAttributesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityApplicationControlProfileControlsIpsAttributesModel
		if i < len(s.IpsAttributes) {
			m = s.IpsAttributes[i]
		}
		values[i] = *m.flattenSecurityApplicationControlProfileControlsIpsAttributes(ctx, ele, diags)
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
	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["services"]; ok {
		m.Services = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Services = types.SetNull(types.StringType)
	}

	return m
}

func (s *resourceSecurityApplicationControlProfileModel) flattenSecurityApplicationControlProfileNetworkProtocolsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityApplicationControlProfileNetworkProtocolsModel {
	if o == nil {
		return []resourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument network_protocols is not type of []interface{}.", "")
		return []resourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	values := make([]resourceSecurityApplicationControlProfileNetworkProtocolsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityApplicationControlProfileNetworkProtocolsModel
		if i < len(s.NetworkProtocols) {
			m = s.NetworkProtocols[i]
		}
		values[i] = *m.flattenSecurityApplicationControlProfileNetworkProtocols(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityApplicationControlProfileControlsModel) expandSecurityApplicationControlProfileControls(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	result["applications"] = data.expandSecurityApplicationControlProfileControlsApplicationsList(ctx, data.Applications, diags)

	result["categories"] = data.expandSecurityApplicationControlProfileControlsCategoriesList(ctx, data.Categories, diags)

	if !data.Risk.IsNull() && !data.Risk.IsUnknown() {
		result["risk"] = expandSetToInt64List(data.Risk)
	}

	if !data.Popularity.IsNull() && !data.Popularity.IsUnknown() {
		result["popularity"] = expandSetToInt64List(data.Popularity)
	}

	result["ipsAttributes"] = data.expandSecurityApplicationControlProfileControlsIpsAttributesList(ctx, data.IpsAttributes, diags)

	return result
}

func (s *resourceSecurityApplicationControlProfileModel) expandSecurityApplicationControlProfileControlsList(ctx context.Context, l []resourceSecurityApplicationControlProfileControlsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityApplicationControlProfileControls(ctx, diags)
	}
	return result
}

func (data *resourceSecurityApplicationControlProfileControlsApplicationsModel) expandSecurityApplicationControlProfileControlsApplications(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityApplicationControlProfileControlsModel) expandSecurityApplicationControlProfileControlsApplicationsList(ctx context.Context, l []resourceSecurityApplicationControlProfileControlsApplicationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityApplicationControlProfileControlsApplications(ctx, diags)
	}
	return result
}

func (data *resourceSecurityApplicationControlProfileControlsCategoriesModel) expandSecurityApplicationControlProfileControlsCategories(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityApplicationControlProfileControlsModel) expandSecurityApplicationControlProfileControlsCategoriesList(ctx context.Context, l []resourceSecurityApplicationControlProfileControlsCategoriesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityApplicationControlProfileControlsCategories(ctx, diags)
	}
	return result
}

func (data *resourceSecurityApplicationControlProfileControlsIpsAttributesModel) expandSecurityApplicationControlProfileControlsIpsAttributes(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityApplicationControlProfileControlsModel) expandSecurityApplicationControlProfileControlsIpsAttributesList(ctx context.Context, l []resourceSecurityApplicationControlProfileControlsIpsAttributesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityApplicationControlProfileControlsIpsAttributes(ctx, diags)
	}
	return result
}

func (data *resourceSecurityApplicationControlProfileNetworkProtocolsModel) expandSecurityApplicationControlProfileNetworkProtocols(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		result["port"] = data.Port.ValueFloat64()
	}

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Services.IsNull() && !data.Services.IsUnknown() {
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
