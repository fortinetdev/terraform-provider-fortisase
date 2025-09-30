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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceNetworkHosts{}

func newResourceNetworkHosts() resource.Resource {
	return &resourceNetworkHosts{}
}

type resourceNetworkHosts struct {
	fortiClient *FortiClient
}

// resourceNetworkHostsModel describes the resource data model.
type resourceNetworkHostsModel struct {
	ID         types.String `tfsdk:"id"`
	PrimaryKey types.String `tfsdk:"primary_key"`
	Type       types.String `tfsdk:"type"`
	Location   types.String `tfsdk:"location"`
	Subnet     types.String `tfsdk:"subnet"`
	StartIp    types.String `tfsdk:"start_ip"`
	EndIp      types.String `tfsdk:"end_ip"`
	Fqdn       types.String `tfsdk:"fqdn"`
	CountryId  types.String `tfsdk:"country_id"`
}

func (r *resourceNetworkHosts) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hosts"
}

func (r *resourceNetworkHosts) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthBetween(1, 79),
				},
				Required: true,
			},
			"type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("ipmask", "iprange", "fqdn", "geography"),
				},
				Computed: true,
				Optional: true,
			},
			"location": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("internal", "external", "private-access", "unspecified"),
				},
				Computed: true,
				Optional: true,
			},
			"subnet": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"start_ip": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"end_ip": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"fqdn": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
				Computed: true,
				Optional: true,
			},
			"country_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(2, 2),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceNetworkHosts) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceNetworkHosts) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceNetworkHostsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectNetworkHosts(ctx, diags))
	input_model.URLParams = *(data.getURLObjectNetworkHosts(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateNetworkHosts(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectNetworkHosts(ctx, "read", diags))

	read_output, err := c.ReadNetworkHosts(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkHosts(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkHosts) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceNetworkHostsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceNetworkHostsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectNetworkHosts(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectNetworkHosts(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateNetworkHosts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectNetworkHosts(ctx, "read", diags))

	read_output, err := c.ReadNetworkHosts(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkHosts(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkHosts) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceNetworkHostsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkHosts(ctx, "delete", diags))

	err := c.DeleteNetworkHosts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceNetworkHosts) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceNetworkHostsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkHosts(ctx, "read", diags))

	read_output, err := c.ReadNetworkHosts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkHosts(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkHosts) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceNetworkHostsModel) refreshNetworkHosts(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["location"]; ok {
		m.Location = parseStringValue(v)
	}

	if v, ok := o["subnet"]; ok {
		m.Subnet = parseStringValue(v)
	}

	if v, ok := o["startIp"]; ok {
		m.StartIp = parseStringValue(v)
	}

	if v, ok := o["endIp"]; ok {
		m.EndIp = parseStringValue(v)
	}

	if v, ok := o["fqdn"]; ok {
		m.Fqdn = parseStringValue(v)
	}

	if v, ok := o["countryId"]; ok {
		m.CountryId = parseStringValue(v)
	}

	return diags
}

func (data *resourceNetworkHostsModel) getCreateObjectNetworkHosts(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Type.IsNull() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Location.IsNull() {
		result["location"] = data.Location.ValueString()
	}

	if !data.Subnet.IsNull() {
		result["subnet"] = data.Subnet.ValueString()
	}

	if !data.StartIp.IsNull() {
		result["startIp"] = data.StartIp.ValueString()
	}

	if !data.EndIp.IsNull() {
		result["endIp"] = data.EndIp.ValueString()
	}

	if !data.Fqdn.IsNull() {
		result["fqdn"] = data.Fqdn.ValueString()
	}

	if !data.CountryId.IsNull() {
		result["countryId"] = data.CountryId.ValueString()
	}

	return &result
}

func (data *resourceNetworkHostsModel) getUpdateObjectNetworkHosts(ctx context.Context, state resourceNetworkHostsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Type.IsNull() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Location.IsNull() {
		result["location"] = data.Location.ValueString()
	}

	if !data.Subnet.IsNull() && !data.Subnet.Equal(state.Subnet) {
		result["subnet"] = data.Subnet.ValueString()
	}

	if !data.StartIp.IsNull() && !data.StartIp.Equal(state.StartIp) {
		result["startIp"] = data.StartIp.ValueString()
	}

	if !data.EndIp.IsNull() && !data.EndIp.Equal(state.EndIp) {
		result["endIp"] = data.EndIp.ValueString()
	}

	if !data.Fqdn.IsNull() && !data.Fqdn.Equal(state.Fqdn) {
		result["fqdn"] = data.Fqdn.ValueString()
	}

	if !data.CountryId.IsNull() && !data.CountryId.Equal(state.CountryId) {
		result["countryId"] = data.CountryId.ValueString()
	}

	return &result
}

func (data *resourceNetworkHostsModel) getURLObjectNetworkHosts(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
