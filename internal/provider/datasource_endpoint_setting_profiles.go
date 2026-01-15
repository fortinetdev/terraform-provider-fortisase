// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointSettingProfiles{}

func newDatasourceEndpointSettingProfiles() datasource.DataSource {
	return &datasourceEndpointSettingProfiles{}
}

type datasourceEndpointSettingProfiles struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointSettingProfilesModel describes the datasource data model.
type datasourceEndpointSettingProfilesModel struct {
	AllowConfigBackup     types.String `tfsdk:"allow_config_backup"`
	ShowTagFortiClient    types.String `tfsdk:"show_tag_forti_client"`
	ShowNotifications     types.String `tfsdk:"show_notifications"`
	NotifyVpnIssue        types.String `tfsdk:"notify_vpn_issue"`
	UsersCanDisconnect    types.String `tfsdk:"users_can_disconnect"`
	EmsDisconnectPassword types.String `tfsdk:"ems_disconnect_password"`
	PrimaryKey            types.String `tfsdk:"primary_key"`
}

func (r *datasourceEndpointSettingProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_setting_profiles"
}

func (r *datasourceEndpointSettingProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"allow_config_backup": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"show_tag_forti_client": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"show_notifications": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"notify_vpn_issue": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"users_can_disconnect": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"ems_disconnect_password": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
			},
		},
	}
}

func (r *datasourceEndpointSettingProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_setting_profiles"
}

func (r *datasourceEndpointSettingProfiles) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointSettingProfilesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointSettingProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointSettingProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSettingProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointSettingProfilesModel) refreshEndpointSettingProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["allowConfigBackup"]; ok {
		m.AllowConfigBackup = parseStringValue(v)
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

	if v, ok := o["emsDisconnectPassword"]; ok {
		m.EmsDisconnectPassword = parseStringValue(v)
	}

	return diags
}

func (data *datasourceEndpointSettingProfilesModel) getURLObjectEndpointSettingProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
