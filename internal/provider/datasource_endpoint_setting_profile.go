// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointSettingProfile{}

func newDatasourceEndpointSettingProfile() datasource.DataSource {
	return &datasourceEndpointSettingProfile{}
}

type datasourceEndpointSettingProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointSettingProfileModel describes the datasource data model.
type datasourceEndpointSettingProfileModel struct {
	AllowConfigBackup               types.String                                 `tfsdk:"allow_config_backup"`
	AllowDebugLogGeneration         types.String                                 `tfsdk:"allow_debug_log_generation"`
	ShowTagFortiClient              types.String                                 `tfsdk:"show_tag_forti_client"`
	ShowNotifications               types.String                                 `tfsdk:"show_notifications"`
	NotifyVpnIssue                  types.String                                 `tfsdk:"notify_vpn_issue"`
	UsersCanDisconnect              types.String                                 `tfsdk:"users_can_disconnect"`
	TriggerVulnScanOnSoftwareChange types.String                                 `tfsdk:"trigger_vuln_scan_on_software_change"`
	FctGui                          *datasourceEndpointSettingProfileFctGuiModel `tfsdk:"fct_gui"`
	EmsDisconnectPassword           types.String                                 `tfsdk:"ems_disconnect_password"`
	PrimaryKey                      types.String                                 `tfsdk:"primary_key"`
}

func (r *datasourceEndpointSettingProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_setting_profile"
}

func (r *datasourceEndpointSettingProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Settings Profile Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"allow_config_backup": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"allow_debug_log_generation": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"show_tag_forti_client": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"show_notifications": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"notify_vpn_issue": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"users_can_disconnect": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"trigger_vuln_scan_on_software_change": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"ems_disconnect_password": schema.StringAttribute{
				Computed: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
			},
			"fct_gui": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"default_tab": schema.StringAttribute{
						Computed: true,
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceEndpointSettingProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_setting_profile"
}

func (r *datasourceEndpointSettingProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointSettingProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointSettingProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointSettingProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSettingProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointSettingProfileModel) refreshEndpointSettingProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["allowConfigBackup"]; ok {
		m.AllowConfigBackup = parseStringValue(v)
	}

	if v, ok := o["allowDebugLogGeneration"]; ok {
		m.AllowDebugLogGeneration = parseStringValue(v)
	}

	if v, ok := o["showTagFortiClient"]; ok {
		m.ShowTagFortiClient = parseStringValue(v)
	}

	if v, ok := o["showNotifications"]; ok {
		m.ShowNotifications = parseStringValue(v)
	}

	if v, ok := o["notifyVpnIssue"]; ok {
		m.NotifyVpnIssue = parseStringValue(v)
	}

	if v, ok := o["usersCanDisconnect"]; ok {
		m.UsersCanDisconnect = parseStringValue(v)
	}

	if v, ok := o["triggerVulnScanOnSoftwareChange"]; ok {
		m.TriggerVulnScanOnSoftwareChange = parseStringValue(v)
	}

	if v, ok := o["fctGui"]; ok {
		m.FctGui = m.FctGui.flattenEndpointSettingProfileFctGui(ctx, v, &diags)
	}

	if v, ok := o["emsDisconnectPassword"]; ok {
		m.EmsDisconnectPassword = parseStringValue(v)
	}

	return diags
}

func (data *datasourceEndpointSettingProfileModel) getURLObjectEndpointSettingProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointSettingProfileFctGuiModel struct {
	DefaultTab types.String `tfsdk:"default_tab"`
}

func (m *datasourceEndpointSettingProfileFctGuiModel) flattenEndpointSettingProfileFctGui(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointSettingProfileFctGuiModel {
	if input == nil {
		return &datasourceEndpointSettingProfileFctGuiModel{}
	}
	if m == nil {
		m = &datasourceEndpointSettingProfileFctGuiModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["defaultTab"]; ok {
		m.DefaultTab = parseStringValue(v)
	}

	return m
}
