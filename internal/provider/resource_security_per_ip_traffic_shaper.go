// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
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
var _ resource.Resource = &resourceSecurityPerIpTrafficShaper{}

func newResourceSecurityPerIpTrafficShaper() resource.Resource {
	return &resourceSecurityPerIpTrafficShaper{}
}

type resourceSecurityPerIpTrafficShaper struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityPerIpTrafficShaperModel describes the resource data model.
type resourceSecurityPerIpTrafficShaperModel struct {
	ID                       types.String  `tfsdk:"id"`
	PrimaryKey               types.String  `tfsdk:"primary_key"`
	MaximumBandwidth         types.Float64 `tfsdk:"maximum_bandwidth"`
	BandwidthUnit            types.String  `tfsdk:"bandwidth_unit"`
	MaxConcurrentSessions    types.Float64 `tfsdk:"max_concurrent_sessions"`
	MaxConcurrentTcpSessions types.Float64 `tfsdk:"max_concurrent_tcp_sessions"`
	MaxConcurrentUdpSessions types.Float64 `tfsdk:"max_concurrent_udp_sessions"`
}

func (r *resourceSecurityPerIpTrafficShaper) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_per_ip_traffic_shaper"
}

func (r *resourceSecurityPerIpTrafficShaper) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Per IP Traffic Shaper Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 64),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"maximum_bandwidth": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(80000000),
				},
				Computed: true,
				Optional: true,
			},
			"bandwidth_unit": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("kbps", "mbps", "gbps"),
				},
				Computed: true,
				Optional: true,
			},
			"max_concurrent_sessions": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(2097000),
				},
				Computed: true,
				Optional: true,
			},
			"max_concurrent_tcp_sessions": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(2097000),
				},
				Computed: true,
				Optional: true,
			},
			"max_concurrent_udp_sessions": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(2097000),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityPerIpTrafficShaper) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_per_ip_traffic_shaper"
}

func (r *resourceSecurityPerIpTrafficShaper) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityPerIpTrafficShaper")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityPerIpTrafficShaperModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityPerIpTrafficShaper(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityPerIpTrafficShaper(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityPerIpTrafficShaper(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	if responseMkey, ok := getCreateResponseMkey(output, "primaryKey"); ok {
		mkey = responseMkey
	}
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityPerIpTrafficShaper(ctx, "read", diags))

	read_output, err := c.ReadSecurityPerIpTrafficShaper(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityPerIpTrafficShaper(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityPerIpTrafficShaper) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityPerIpTrafficShaper")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityPerIpTrafficShaperModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityPerIpTrafficShaperModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityPerIpTrafficShaper(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityPerIpTrafficShaper(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityPerIpTrafficShaper(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityPerIpTrafficShaper(ctx, "read", diags))

	read_output, err := c.ReadSecurityPerIpTrafficShaper(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityPerIpTrafficShaper(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityPerIpTrafficShaper) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityPerIpTrafficShaper")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityPerIpTrafficShaperModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityPerIpTrafficShaper(ctx, "delete", diags))

	output, err := c.DeleteSecurityPerIpTrafficShaper(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityPerIpTrafficShaper) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityPerIpTrafficShaperModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityPerIpTrafficShaper(ctx, "read", diags))

	read_output, err := c.ReadSecurityPerIpTrafficShaper(&input_model)
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

	diags.Append(data.refreshSecurityPerIpTrafficShaper(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityPerIpTrafficShaper) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityPerIpTrafficShaperModel) refreshSecurityPerIpTrafficShaper(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["maximumBandwidth"]; ok {
		m.MaximumBandwidth = parseFloat64Value(v)
	}

	if v, ok := o["bandwidthUnit"]; ok {
		m.BandwidthUnit = parseStringValue(v)
	}

	if v, ok := o["maxConcurrentSessions"]; ok {
		m.MaxConcurrentSessions = parseFloat64Value(v)
	}

	if v, ok := o["maxConcurrentTcpSessions"]; ok {
		m.MaxConcurrentTcpSessions = parseFloat64Value(v)
	}

	if v, ok := o["maxConcurrentUdpSessions"]; ok {
		m.MaxConcurrentUdpSessions = parseFloat64Value(v)
	}

	return diags
}

func (data *resourceSecurityPerIpTrafficShaperModel) getCreateObjectSecurityPerIpTrafficShaper(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.MaximumBandwidth.IsNull() && !data.MaximumBandwidth.IsUnknown() {
		result["maximumBandwidth"] = data.MaximumBandwidth.ValueFloat64()
	}

	if !data.BandwidthUnit.IsNull() && !data.BandwidthUnit.IsUnknown() {
		result["bandwidthUnit"] = data.BandwidthUnit.ValueString()
	}

	if !data.MaxConcurrentSessions.IsNull() && !data.MaxConcurrentSessions.IsUnknown() {
		result["maxConcurrentSessions"] = data.MaxConcurrentSessions.ValueFloat64()
	}

	if !data.MaxConcurrentTcpSessions.IsNull() && !data.MaxConcurrentTcpSessions.IsUnknown() {
		result["maxConcurrentTcpSessions"] = data.MaxConcurrentTcpSessions.ValueFloat64()
	}

	if !data.MaxConcurrentUdpSessions.IsNull() && !data.MaxConcurrentUdpSessions.IsUnknown() {
		result["maxConcurrentUdpSessions"] = data.MaxConcurrentUdpSessions.ValueFloat64()
	}

	return &result
}

func (data *resourceSecurityPerIpTrafficShaperModel) getUpdateObjectSecurityPerIpTrafficShaper(ctx context.Context, state resourceSecurityPerIpTrafficShaperModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.MaximumBandwidth.IsNull() && !data.MaximumBandwidth.IsUnknown() {
		result["maximumBandwidth"] = data.MaximumBandwidth.ValueFloat64()
	}

	if !data.BandwidthUnit.IsNull() && !data.BandwidthUnit.IsUnknown() {
		result["bandwidthUnit"] = data.BandwidthUnit.ValueString()
	}

	if !data.MaxConcurrentSessions.IsNull() && !data.MaxConcurrentSessions.IsUnknown() {
		result["maxConcurrentSessions"] = data.MaxConcurrentSessions.ValueFloat64()
	}

	if !data.MaxConcurrentTcpSessions.IsNull() && !data.MaxConcurrentTcpSessions.IsUnknown() {
		result["maxConcurrentTcpSessions"] = data.MaxConcurrentTcpSessions.ValueFloat64()
	}

	if !data.MaxConcurrentUdpSessions.IsNull() && !data.MaxConcurrentUdpSessions.IsUnknown() {
		result["maxConcurrentUdpSessions"] = data.MaxConcurrentUdpSessions.ValueFloat64()
	}

	return &result
}

func (data *resourceSecurityPerIpTrafficShaperModel) getURLObjectSecurityPerIpTrafficShaper(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
