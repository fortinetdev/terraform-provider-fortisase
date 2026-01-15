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
var _ datasource.DataSource = &datasourceAuthRadiusServers{}

func newDatasourceAuthRadiusServers() datasource.DataSource {
	return &datasourceAuthRadiusServers{}
}

type datasourceAuthRadiusServers struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceAuthRadiusServersModel describes the datasource data model.
type datasourceAuthRadiusServersModel struct {
	PrimaryKey                 types.String `tfsdk:"primary_key"`
	AuthType                   types.String `tfsdk:"auth_type"`
	PrimaryServer              types.String `tfsdk:"primary_server"`
	IncludedInDefaultUserGroup types.Bool   `tfsdk:"included_in_default_user_group"`
	SecondaryServer            types.String `tfsdk:"secondary_server"`
}

func (r *datasourceAuthRadiusServers) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_radius_servers"
}

func (r *datasourceAuthRadiusServers) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 35),
				},
				Required: true,
			},
			"auth_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("auto", "pap", "chap", "ms_chap", "ms_chap_v2"),
				},
				Computed: true,
				Optional: true,
			},
			"primary_server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"included_in_default_user_group": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"secondary_server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceAuthRadiusServers) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_radius_servers"
}

func (r *datasourceAuthRadiusServers) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceAuthRadiusServersModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthRadiusServers(ctx, "read", diags))

	read_output, err := c.ReadAuthRadiusServers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthRadiusServers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceAuthRadiusServersModel) refreshAuthRadiusServers(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["authType"]; ok {
		m.AuthType = parseStringValue(v)
	}

	if v, ok := o["primaryServer"]; ok {
		m.PrimaryServer = parseStringValue(v)
	}

	if v, ok := o["includedInDefaultUserGroup"]; ok {
		m.IncludedInDefaultUserGroup = parseBoolValue(v)
	}

	if v, ok := o["secondaryServer"]; ok {
		m.SecondaryServer = parseStringValue(v)
	}

	return diags
}

func (data *datasourceAuthRadiusServersModel) getURLObjectAuthRadiusServers(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
