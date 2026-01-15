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
var _ datasource.DataSource = &datasourcePrivateAccessNetworkConfiguration{}

func newDatasourcePrivateAccessNetworkConfiguration() datasource.DataSource {
	return &datasourcePrivateAccessNetworkConfiguration{}
}

type datasourcePrivateAccessNetworkConfiguration struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourcePrivateAccessNetworkConfigurationModel describes the datasource data model.
type datasourcePrivateAccessNetworkConfigurationModel struct {
	BgpRouterIdsSubnet types.String `tfsdk:"bgp_router_ids_subnet"`
	AsNumber           types.String `tfsdk:"as_number"`
	RecursiveNextHop   types.Bool   `tfsdk:"recursive_next_hop"`
	SdwanRuleEnable    types.Bool   `tfsdk:"sdwan_rule_enable"`
	SdwanHealthCheckVm types.String `tfsdk:"sdwan_health_check_vm"`
	ConfigState        types.String `tfsdk:"config_state"`
	BgpDesign          types.String `tfsdk:"bgp_design"`
}

func (r *datasourcePrivateAccessNetworkConfiguration) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_network_configuration"
}

func (r *datasourcePrivateAccessNetworkConfiguration) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bgp_router_ids_subnet": schema.StringAttribute{
				MarkdownDescription: "Available/unused subnet that can be used to assign loopback interface IP addresses used for BGP router IDs parameter on the FortiSASE security PoPs. /28 is the minimum subnet size.",
				Computed:            true,
				Optional:            true,
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: "Autonomous System Number (ASN).",
				Computed:            true,
				Optional:            true,
			},
			"recursive_next_hop": schema.BoolAttribute{
				MarkdownDescription: "BGP Recursive Routing. Enabling this setting allows for interhub connectivity. When use BGP design on-loopback this has to be enabled.",
				Computed:            true,
				Optional:            true,
			},
			"sdwan_rule_enable": schema.BoolAttribute{
				MarkdownDescription: "Hub Selection Method. Enabling this setting the highest priority service connection that meets minimum SLA requirements is selected. Otherwise BGP MED (Multi-Exit Discriminator) will be used",
				Computed:            true,
				Optional:            true,
			},
			"sdwan_health_check_vm": schema.StringAttribute{
				MarkdownDescription: "Health Check IP. Must be provided when enable sdwan rule which used to obtain Jitter, latency and packet loss measurements.",
				Computed:            true,
				Optional:            true,
			},
			"config_state": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("success", "failed", "creating", "updating", "deleting"),
				},
				MarkdownDescription: "Configuration state of network configuration.\nSupported values: success, failed, creating, updating, deleting.",
				Computed:            true,
			},
			"bgp_design": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("overlay", "loopback"),
				},
				MarkdownDescription: "BGP Routing Design.\nSupported values: overlay, loopback.",
				Computed:            true,
				Optional:            true,
			},
		},
	}
}

func (r *datasourcePrivateAccessNetworkConfiguration) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_private_access_network_configuration"
}

func (r *datasourcePrivateAccessNetworkConfiguration) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourcePrivateAccessNetworkConfigurationModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := "PrivateAccessNetworkConfiguration"

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadPrivateAccessNetworkConfiguration(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshPrivateAccessNetworkConfiguration(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourcePrivateAccessNetworkConfigurationModel) refreshPrivateAccessNetworkConfiguration(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["bgp_router_ids_subnet"]; ok {
		m.BgpRouterIdsSubnet = parseStringValue(v)
	}

	if v, ok := o["as_number"]; ok {
		m.AsNumber = parseStringValue(v)
	}

	if v, ok := o["recursive_next_hop"]; ok {
		m.RecursiveNextHop = parseBoolValue(v)
	}

	if v, ok := o["sdwan_rule_enable"]; ok {
		m.SdwanRuleEnable = parseBoolValue(v)
	}

	if v, ok := o["sdwan_health_check_vm"]; ok {
		m.SdwanHealthCheckVm = parseStringValue(v)
	}

	if v, ok := o["config_state"]; ok {
		m.ConfigState = parseStringValue(v)
	}

	if v, ok := o["bgp_design"]; ok {
		m.BgpDesign = parseStringValue(v)
	}

	return diags
}
