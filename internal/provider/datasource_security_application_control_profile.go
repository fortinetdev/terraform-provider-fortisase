// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
	fortiClient *FortiClient
}

// datasourceSecurityApplicationControlProfileModel describes the datasource data model.
type datasourceSecurityApplicationControlProfileModel struct {
	PrimaryKey                      types.String                                                                  `tfsdk:"primary_key"`
	ApplicationCategoryControls     []datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel `tfsdk:"application_category_controls"`
	ApplicationControls             []datasourceSecurityApplicationControlProfileApplicationControlsModel         `tfsdk:"application_controls"`
	UnknownApplicationAction        types.String                                                                  `tfsdk:"unknown_application_action"`
	NetworkProtocolEnforcement      types.String                                                                  `tfsdk:"network_protocol_enforcement"`
	NetworkProtocols                []datasourceSecurityApplicationControlProfileNetworkProtocolsModel            `tfsdk:"network_protocols"`
	BlockNonDefaultPortApplications types.String                                                                  `tfsdk:"block_non_default_port_applications"`
	Direction                       types.String                                                                  `tfsdk:"direction"`
}

func (r *datasourceSecurityApplicationControlProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_application_control_profile"
}

func (r *datasourceSecurityApplicationControlProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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
			fmt.Sprintf("Error to read data source: %v", err),
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

func (m *datasourceSecurityApplicationControlProfileModel) refreshSecurityApplicationControlProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *datasourceSecurityApplicationControlProfileModel) getURLObjectSecurityApplicationControlProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel struct {
	Action   types.String                                                                         `tfsdk:"action"`
	Category *datasourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel `tfsdk:"category"`
}

type datasourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityApplicationControlProfileApplicationControlsModel struct {
	Action       types.String                                                                      `tfsdk:"action"`
	Applications []datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel `tfsdk:"applications"`
}

type datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityApplicationControlProfileNetworkProtocolsModel struct {
	Port     types.Float64 `tfsdk:"port"`
	Action   types.String  `tfsdk:"action"`
	Services types.Set     `tfsdk:"services"`
}

func (m *datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel) flattenSecurityApplicationControlProfileApplicationCategoryControls(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel {
	if input == nil {
		return &datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel{}
	}
	if m == nil {
		m = &datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel{}
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

func (s *datasourceSecurityApplicationControlProfileModel) flattenSecurityApplicationControlProfileApplicationCategoryControlsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel {
	if o == nil {
		return []datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument application_category_controls is not type of []interface{}.", "")
		return []datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel{}
	}

	values := make([]datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityApplicationControlProfileApplicationCategoryControlsModel
		values[i] = *m.flattenSecurityApplicationControlProfileApplicationCategoryControls(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel) flattenSecurityApplicationControlProfileApplicationCategoryControlsCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel {
	if input == nil {
		return &datasourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel{}
	}
	if m == nil {
		m = &datasourceSecurityApplicationControlProfileApplicationCategoryControlsCategoryModel{}
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

func (m *datasourceSecurityApplicationControlProfileApplicationControlsModel) flattenSecurityApplicationControlProfileApplicationControls(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityApplicationControlProfileApplicationControlsModel {
	if input == nil {
		return &datasourceSecurityApplicationControlProfileApplicationControlsModel{}
	}
	if m == nil {
		m = &datasourceSecurityApplicationControlProfileApplicationControlsModel{}
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

func (s *datasourceSecurityApplicationControlProfileModel) flattenSecurityApplicationControlProfileApplicationControlsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityApplicationControlProfileApplicationControlsModel {
	if o == nil {
		return []datasourceSecurityApplicationControlProfileApplicationControlsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument application_controls is not type of []interface{}.", "")
		return []datasourceSecurityApplicationControlProfileApplicationControlsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityApplicationControlProfileApplicationControlsModel{}
	}

	values := make([]datasourceSecurityApplicationControlProfileApplicationControlsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityApplicationControlProfileApplicationControlsModel
		values[i] = *m.flattenSecurityApplicationControlProfileApplicationControls(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel) flattenSecurityApplicationControlProfileApplicationControlsApplications(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel {
	if input == nil {
		return &datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel{}
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

func (s *datasourceSecurityApplicationControlProfileApplicationControlsModel) flattenSecurityApplicationControlProfileApplicationControlsApplicationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel {
	if o == nil {
		return []datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument applications is not type of []interface{}.", "")
		return []datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel{}
	}

	values := make([]datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityApplicationControlProfileApplicationControlsApplicationsModel
		values[i] = *m.flattenSecurityApplicationControlProfileApplicationControlsApplications(ctx, ele, diags)
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

func (s *datasourceSecurityApplicationControlProfileModel) flattenSecurityApplicationControlProfileNetworkProtocolsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityApplicationControlProfileNetworkProtocolsModel {
	if o == nil {
		return []datasourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument network_protocols is not type of []interface{}.", "")
		return []datasourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityApplicationControlProfileNetworkProtocolsModel{}
	}

	values := make([]datasourceSecurityApplicationControlProfileNetworkProtocolsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityApplicationControlProfileNetworkProtocolsModel
		values[i] = *m.flattenSecurityApplicationControlProfileNetworkProtocols(ctx, ele, diags)
	}

	return values
}
