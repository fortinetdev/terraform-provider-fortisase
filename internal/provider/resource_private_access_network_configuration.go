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
	"time"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourcePrivateAccessNetworkConfiguration{}

func newResourcePrivateAccessNetworkConfiguration() resource.Resource {
	return &resourcePrivateAccessNetworkConfiguration{}
}

type resourcePrivateAccessNetworkConfiguration struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourcePrivateAccessNetworkConfigurationModel describes the resource data model.
type resourcePrivateAccessNetworkConfigurationModel struct {
	ID                 types.String `tfsdk:"id"`
	BgpRouterIdsSubnet types.String `tfsdk:"bgp_router_ids_subnet"`
	AsNumber           types.String `tfsdk:"as_number"`
	RecursiveNextHop   types.Bool   `tfsdk:"recursive_next_hop"`
	SdwanRuleEnable    types.Bool   `tfsdk:"sdwan_rule_enable"`
	SdwanHealthCheckVm types.String `tfsdk:"sdwan_health_check_vm"`
	ConfigState        types.String `tfsdk:"config_state"`
	BgpDesign          types.String `tfsdk:"bgp_design"`
}

func (r *resourcePrivateAccessNetworkConfiguration) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_network_configuration"
}

func (r *resourcePrivateAccessNetworkConfiguration) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
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

func (r *resourcePrivateAccessNetworkConfiguration) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourcePrivateAccessNetworkConfiguration) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourcePrivateAccessNetworkConfigurationModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectPrivateAccessNetworkConfiguration(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreatePrivateAccessNetworkConfiguration(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := "PrivateAccessNetworkConfiguration"
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output := make(map[string]interface{})
	for i := 0; i < 20; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadPrivateAccessNetworkConfiguration(&read_input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
				getErrorDetail(&read_input_model, read_output),
			)
			return
		}
		if v, ok := read_output["config_state"]; ok {
			if v != "success" {
				continue
			}
		}
		break
	}

	diags.Append(data.refreshPrivateAccessNetworkConfiguration(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessNetworkConfiguration) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourcePrivateAccessNetworkConfigurationModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourcePrivateAccessNetworkConfigurationModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectPrivateAccessNetworkConfiguration(ctx, state, diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdatePrivateAccessNetworkConfiguration(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output := make(map[string]interface{})
	for i := 0; i < 20; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadPrivateAccessNetworkConfiguration(&read_input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
				getErrorDetail(&read_input_model, read_output),
			)
			return
		}
		if v, ok := read_output["config_state"]; ok {
			if v != "success" {
				continue
			}
		}
		break
	}

	diags.Append(data.refreshPrivateAccessNetworkConfiguration(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessNetworkConfiguration) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourcePrivateAccessNetworkConfigurationModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	output, err := c.DeletePrivateAccessNetworkConfiguration(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	read_output := make(map[string]interface{})
	for i := 0; i < 20; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadPrivateAccessNetworkConfiguration(&input_model)
		if err != nil || len(read_output) == 0 {
			// Delete success
			return
		}
	}
	diags.AddError(
		fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
		fmt.Sprintf("The resource still exists %s: %v", r.resourceName, read_output),
	)
}

func (r *resourcePrivateAccessNetworkConfiguration) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourcePrivateAccessNetworkConfigurationModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadPrivateAccessNetworkConfiguration(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
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

func (r *resourcePrivateAccessNetworkConfiguration) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourcePrivateAccessNetworkConfigurationModel) refreshPrivateAccessNetworkConfiguration(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourcePrivateAccessNetworkConfigurationModel) getCreateObjectPrivateAccessNetworkConfiguration(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.BgpRouterIdsSubnet.IsNull() {
		result["bgp_router_ids_subnet"] = data.BgpRouterIdsSubnet.ValueString()
	}

	if !data.AsNumber.IsNull() {
		result["as_number"] = data.AsNumber.ValueString()
	}

	if !data.RecursiveNextHop.IsNull() {
		result["recursive_next_hop"] = data.RecursiveNextHop.ValueBool()
	}

	if !data.SdwanRuleEnable.IsNull() {
		result["sdwan_rule_enable"] = data.SdwanRuleEnable.ValueBool()
	}

	if !data.SdwanHealthCheckVm.IsNull() {
		result["sdwan_health_check_vm"] = data.SdwanHealthCheckVm.ValueString()
	}

	if !data.BgpDesign.IsNull() {
		result["bgp_design"] = data.BgpDesign.ValueString()
	}

	return &result
}

func (data *resourcePrivateAccessNetworkConfigurationModel) getUpdateObjectPrivateAccessNetworkConfiguration(ctx context.Context, state resourcePrivateAccessNetworkConfigurationModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.BgpRouterIdsSubnet.IsNull() {
		result["bgp_router_ids_subnet"] = data.BgpRouterIdsSubnet.ValueString()
	}

	if !data.AsNumber.IsNull() {
		result["as_number"] = data.AsNumber.ValueString()
	}

	if !data.RecursiveNextHop.IsNull() {
		result["recursive_next_hop"] = data.RecursiveNextHop.ValueBool()
	}

	if !data.SdwanRuleEnable.IsNull() {
		result["sdwan_rule_enable"] = data.SdwanRuleEnable.ValueBool()
	}

	if !data.SdwanHealthCheckVm.IsNull() {
		result["sdwan_health_check_vm"] = data.SdwanHealthCheckVm.ValueString()
	}

	return &result
}
