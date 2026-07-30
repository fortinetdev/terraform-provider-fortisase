// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/setvalidatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityApplicationControlProfile{}

func newDatasourceSecurityApplicationControlProfile() datasource.DataSource {
	return &datasourceSecurityApplicationControlProfile{}
}

type datasourceSecurityApplicationControlProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityApplicationControlProfileModel describes the datasource data model.
type datasourceSecurityApplicationControlProfileModel struct {
	PrimaryKey                      types.String                                                       `tfsdk:"primary_key"`
	Controls                        []datasourceSecurityApplicationControlProfileControlsModel         `tfsdk:"controls"`
	UnknownApplicationAction        types.String                                                       `tfsdk:"unknown_application_action"`
	NetworkProtocolEnforcement      types.String                                                       `tfsdk:"network_protocol_enforcement"`
	NetworkProtocols                []datasourceSecurityApplicationControlProfileNetworkProtocolsModel `tfsdk:"network_protocols"`
	BlockNonDefaultPortApplications types.String                                                       `tfsdk:"block_non_default_port_applications"`
	Direction                       types.String                                                       `tfsdk:"direction"`
}

func (r *datasourceSecurityApplicationControlProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_application_control_profile"
}

func (r *datasourceSecurityApplicationControlProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Application Control Profile Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"unknown_application_action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("block", "allow", "monitor"),
				},
				Computed: true,
			},
			"network_protocol_enforcement": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"block_non_default_port_applications": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("internal-profiles", "outbound-profiles"),
				},
				MarkdownDescription: "The direction of the target resource.\nSupported values: internal-profiles, outbound-profiles.",
				Required:            true,
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
						},
						"risk": schema.SetAttribute{
							MarkdownDescription: "Risk level(s) with 0 being lowest and 4 being highest",
							Computed:            true,
							ElementType:         types.Int64Type,
						},
						"popularity": schema.SetAttribute{
							MarkdownDescription: "Popularity level(s) with 1 being lowest and 5 being highest",
							Computed:            true,
							ElementType:         types.Int64Type,
						},
						"applications": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Computed: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidatorwarning.OneOf("security/applications"),
										},
										Computed: true,
									},
								},
							},
							Computed: true,
						},
						"categories": schema.ListNestedAttribute{
							MarkdownDescription: "Set the control action for a given application category. For the 'Proxy' category, the action cannot be set to 'block' if the linked profile group is referenced by a proxy policy.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Computed: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidatorwarning.OneOf("security/application-categories"),
										},
										Computed: true,
									},
								},
							},
							Computed: true,
						},
						"ips_attributes": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Computed: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidatorwarning.OneOf("security/ips-attr-map"),
										},
										Computed: true,
									},
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"network_protocols": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(1, 65535),
							},
							Computed: true,
						},
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("monitor", "pass", "block"),
							},
							Computed: true,
						},
						"services": schema.SetAttribute{
							Validators: []validator.Set{
								setvalidatorwarning.ValueStringsAre(
									stringvalidatorwarning.OneOf("dns", "ftp", "http", "https", "imap", "nntp", "pop3", "smtp", "snmp", "ssh", "telnet"),
								),
								setvalidatorwarning.SizeAtLeast(1),
							},
							Computed:    true,
							ElementType: types.StringType,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceSecurityApplicationControlProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityApplicationControlProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityApplicationControlProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityApplicationControlProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityApplicationControlProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
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

func (m *datasourceSecurityApplicationControlProfileModel) refreshSecurityApplicationControlProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *datasourceSecurityApplicationControlProfileModel) getURLObjectSecurityApplicationControlProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
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

type datasourceSecurityApplicationControlProfileControlsModel struct {
	Action        types.String                                                            `tfsdk:"action"`
	Applications  []datasourceSecurityApplicationControlProfileControlsApplicationsModel  `tfsdk:"applications"`
	Categories    []datasourceSecurityApplicationControlProfileControlsCategoriesModel    `tfsdk:"categories"`
	Risk          types.Set                                                               `tfsdk:"risk"`
	Popularity    types.Set                                                               `tfsdk:"popularity"`
	IpsAttributes []datasourceSecurityApplicationControlProfileControlsIpsAttributesModel `tfsdk:"ips_attributes"`
}

type datasourceSecurityApplicationControlProfileControlsApplicationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityApplicationControlProfileControlsCategoriesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityApplicationControlProfileControlsIpsAttributesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityApplicationControlProfileNetworkProtocolsModel struct {
	Port     types.Float64 `tfsdk:"port"`
	Action   types.String  `tfsdk:"action"`
	Services types.Set     `tfsdk:"services"`
}

func (m *datasourceSecurityApplicationControlProfileControlsModel) flattenSecurityApplicationControlProfileControls(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityApplicationControlProfileControlsModel {
	if input == nil {
		return &datasourceSecurityApplicationControlProfileControlsModel{}
	}
	if m == nil {
		m = &datasourceSecurityApplicationControlProfileControlsModel{}
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

func (s *datasourceSecurityApplicationControlProfileModel) flattenSecurityApplicationControlProfileControlsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityApplicationControlProfileControlsModel {
	if o == nil {
		return []datasourceSecurityApplicationControlProfileControlsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument controls is not type of []interface{}.", "")
		return []datasourceSecurityApplicationControlProfileControlsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityApplicationControlProfileControlsModel{}
	}

	values := make([]datasourceSecurityApplicationControlProfileControlsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityApplicationControlProfileControlsModel
		if i < len(s.Controls) {
			m = s.Controls[i]
		}
		values[i] = *m.flattenSecurityApplicationControlProfileControls(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityApplicationControlProfileControlsApplicationsModel) flattenSecurityApplicationControlProfileControlsApplications(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityApplicationControlProfileControlsApplicationsModel {
	if input == nil {
		return &datasourceSecurityApplicationControlProfileControlsApplicationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityApplicationControlProfileControlsApplicationsModel{}
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

func (s *datasourceSecurityApplicationControlProfileControlsModel) flattenSecurityApplicationControlProfileControlsApplicationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityApplicationControlProfileControlsApplicationsModel {
	if o == nil {
		return []datasourceSecurityApplicationControlProfileControlsApplicationsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument applications is not type of []interface{}.", "")
		return []datasourceSecurityApplicationControlProfileControlsApplicationsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityApplicationControlProfileControlsApplicationsModel{}
	}

	values := make([]datasourceSecurityApplicationControlProfileControlsApplicationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityApplicationControlProfileControlsApplicationsModel
		if i < len(s.Applications) {
			m = s.Applications[i]
		}
		values[i] = *m.flattenSecurityApplicationControlProfileControlsApplications(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityApplicationControlProfileControlsCategoriesModel) flattenSecurityApplicationControlProfileControlsCategories(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityApplicationControlProfileControlsCategoriesModel {
	if input == nil {
		return &datasourceSecurityApplicationControlProfileControlsCategoriesModel{}
	}
	if m == nil {
		m = &datasourceSecurityApplicationControlProfileControlsCategoriesModel{}
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

func (s *datasourceSecurityApplicationControlProfileControlsModel) flattenSecurityApplicationControlProfileControlsCategoriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityApplicationControlProfileControlsCategoriesModel {
	if o == nil {
		return []datasourceSecurityApplicationControlProfileControlsCategoriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument categories is not type of []interface{}.", "")
		return []datasourceSecurityApplicationControlProfileControlsCategoriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityApplicationControlProfileControlsCategoriesModel{}
	}

	values := make([]datasourceSecurityApplicationControlProfileControlsCategoriesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityApplicationControlProfileControlsCategoriesModel
		if i < len(s.Categories) {
			m = s.Categories[i]
		}
		values[i] = *m.flattenSecurityApplicationControlProfileControlsCategories(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityApplicationControlProfileControlsIpsAttributesModel) flattenSecurityApplicationControlProfileControlsIpsAttributes(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityApplicationControlProfileControlsIpsAttributesModel {
	if input == nil {
		return &datasourceSecurityApplicationControlProfileControlsIpsAttributesModel{}
	}
	if m == nil {
		m = &datasourceSecurityApplicationControlProfileControlsIpsAttributesModel{}
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

func (s *datasourceSecurityApplicationControlProfileControlsModel) flattenSecurityApplicationControlProfileControlsIpsAttributesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityApplicationControlProfileControlsIpsAttributesModel {
	if o == nil {
		return []datasourceSecurityApplicationControlProfileControlsIpsAttributesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument ips_attributes is not type of []interface{}.", "")
		return []datasourceSecurityApplicationControlProfileControlsIpsAttributesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityApplicationControlProfileControlsIpsAttributesModel{}
	}

	values := make([]datasourceSecurityApplicationControlProfileControlsIpsAttributesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityApplicationControlProfileControlsIpsAttributesModel
		if i < len(s.IpsAttributes) {
			m = s.IpsAttributes[i]
		}
		values[i] = *m.flattenSecurityApplicationControlProfileControlsIpsAttributes(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityApplicationControlProfileNetworkProtocolsModel) flattenSecurityApplicationControlProfileNetworkProtocols(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityApplicationControlProfileNetworkProtocolsModel {
	if input == nil {
		return &datasourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}
	if m == nil {
		m = &datasourceSecurityApplicationControlProfileNetworkProtocolsModel{}
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

func (s *datasourceSecurityApplicationControlProfileModel) flattenSecurityApplicationControlProfileNetworkProtocolsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityApplicationControlProfileNetworkProtocolsModel {
	if o == nil {
		return []datasourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument network_protocols is not type of []interface{}.", "")
		return []datasourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	values := make([]datasourceSecurityApplicationControlProfileNetworkProtocolsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityApplicationControlProfileNetworkProtocolsModel
		if i < len(s.NetworkProtocols) {
			m = s.NetworkProtocols[i]
		}
		values[i] = *m.flattenSecurityApplicationControlProfileNetworkProtocols(ctx, ele, diags)
	}

	return values
}
