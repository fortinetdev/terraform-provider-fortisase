// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceNetworkImplicitDnsRules{}

func newResourceNetworkImplicitDnsRules() resource.Resource {
	return &resourceNetworkImplicitDnsRules{}
}

type resourceNetworkImplicitDnsRules struct {
	fortiClient *FortiClient
}

// resourceNetworkImplicitDnsRulesModel describes the resource data model.
type resourceNetworkImplicitDnsRulesModel struct {
	ID         types.String `tfsdk:"id"`
	PrimaryKey types.String `tfsdk:"primary_key"`
	DnsServer  types.String `tfsdk:"dns_server"`
	DnsServer1 types.String `tfsdk:"dns_server1"`
	DnsServer2 types.String `tfsdk:"dns_server2"`
	Protocols  types.Set    `tfsdk:"protocols"`
	ForPrivate types.Bool   `tfsdk:"for_private"`
}

func (r *resourceNetworkImplicitDnsRules) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_implicit_dns_rules"
}

func (r *resourceNetworkImplicitDnsRules) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Validators: []validator.String{
					stringvalidator.OneOf("vpn", "other", "implicit_all"),
				},
				Required: true,
			},
			"dns_server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("fortiguard", "google", "quad9", "cloudflare", "endpoint", "custom"),
				},
				Computed: true,
				Optional: true,
			},
			"dns_server1": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"dns_server2": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"protocols": schema.SetAttribute{
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
				},
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"for_private": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceNetworkImplicitDnsRules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceNetworkImplicitDnsRules) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceNetworkImplicitDnsRulesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectNetworkImplicitDnsRules(ctx, diags))
	input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRules(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateNetworkImplicitDnsRules(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRules(ctx, "read", diags))

	read_output, err := c.ReadNetworkImplicitDnsRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkImplicitDnsRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkImplicitDnsRules) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceNetworkImplicitDnsRulesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceNetworkImplicitDnsRulesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectNetworkImplicitDnsRules(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRules(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateNetworkImplicitDnsRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRules(ctx, "read", diags))

	read_output, err := c.ReadNetworkImplicitDnsRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkImplicitDnsRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkImplicitDnsRules) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceNetworkImplicitDnsRules) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceNetworkImplicitDnsRulesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRules(ctx, "read", diags))

	read_output, err := c.ReadNetworkImplicitDnsRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkImplicitDnsRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkImplicitDnsRules) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceNetworkImplicitDnsRulesModel) refreshNetworkImplicitDnsRules(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["dnsServer"]; ok {
		m.DnsServer = parseStringValue(v)
	}

	if v, ok := o["dnsServer1"]; ok {
		m.DnsServer1 = parseStringValue(v)
	}

	if v, ok := o["dnsServer2"]; ok {
		m.DnsServer2 = parseStringValue(v)
	}

	if v, ok := o["protocols"]; ok {
		m.Protocols = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["forPrivate"]; ok {
		m.ForPrivate = parseBoolValue(v)
	}

	return diags
}

func (data *resourceNetworkImplicitDnsRulesModel) getCreateObjectNetworkImplicitDnsRules(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.DnsServer.IsNull() {
		result["dnsServer"] = data.DnsServer.ValueString()
	}

	if !data.DnsServer1.IsNull() {
		result["dnsServer1"] = data.DnsServer1.ValueString()
	}

	if !data.DnsServer2.IsNull() {
		result["dnsServer2"] = data.DnsServer2.ValueString()
	}

	if !data.Protocols.IsNull() {
		result["protocols"] = expandSetToStringList(data.Protocols)
	}

	if !data.ForPrivate.IsNull() {
		result["forPrivate"] = data.ForPrivate.ValueBool()
	}

	return &result
}

func (data *resourceNetworkImplicitDnsRulesModel) getUpdateObjectNetworkImplicitDnsRules(ctx context.Context, state resourceNetworkImplicitDnsRulesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.DnsServer.IsNull() {
		result["dnsServer"] = data.DnsServer.ValueString()
	}

	if !data.DnsServer1.IsNull() && !data.DnsServer1.Equal(state.DnsServer1) {
		result["dnsServer1"] = data.DnsServer1.ValueString()
	}

	if !data.DnsServer2.IsNull() && !data.DnsServer2.Equal(state.DnsServer2) {
		result["dnsServer2"] = data.DnsServer2.ValueString()
	}

	if !data.Protocols.IsNull() && !data.Protocols.Equal(state.Protocols) {
		result["protocols"] = expandSetToStringList(data.Protocols)
	}

	if !data.ForPrivate.IsNull() && !data.ForPrivate.Equal(state.ForPrivate) {
		result["forPrivate"] = data.ForPrivate.ValueBool()
	}

	return &result
}

func (data *resourceNetworkImplicitDnsRulesModel) getURLObjectNetworkImplicitDnsRules(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
