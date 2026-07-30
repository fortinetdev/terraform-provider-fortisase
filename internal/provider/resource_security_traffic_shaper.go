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
var _ resource.Resource = &resourceSecurityTrafficShaper{}

func newResourceSecurityTrafficShaper() resource.Resource {
	return &resourceSecurityTrafficShaper{}
}

type resourceSecurityTrafficShaper struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityTrafficShaperModel describes the resource data model.
type resourceSecurityTrafficShaperModel struct {
	ID                  types.String  `tfsdk:"id"`
	PrimaryKey          types.String  `tfsdk:"primary_key"`
	GuaranteedBandwidth types.Float64 `tfsdk:"guaranteed_bandwidth"`
	MaximumBandwidth    types.Float64 `tfsdk:"maximum_bandwidth"`
	BandwidthUnit       types.String  `tfsdk:"bandwidth_unit"`
	Priority            types.String  `tfsdk:"priority"`
	PerPolicy           types.String  `tfsdk:"per_policy"`
}

func (r *resourceSecurityTrafficShaper) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_traffic_shaper"
}

func (r *resourceSecurityTrafficShaper) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Shared Traffic Shaper Resource API V2 for FortiSASE.",
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
			"guaranteed_bandwidth": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(80000000),
				},
				Computed: true,
				Optional: true,
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
			"priority": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("low", "medium", "high"),
				},
				Computed: true,
				Optional: true,
			},
			"per_policy": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("disable", "enable"),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityTrafficShaper) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_traffic_shaper"
}

func (r *resourceSecurityTrafficShaper) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityTrafficShaper")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityTrafficShaperModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityTrafficShaper(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityTrafficShaper(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityTrafficShaper(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityTrafficShaper(ctx, "read", diags))

	read_output, err := c.ReadSecurityTrafficShaper(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityTrafficShaper(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityTrafficShaper) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityTrafficShaper")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityTrafficShaperModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityTrafficShaperModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityTrafficShaper(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityTrafficShaper(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityTrafficShaper(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityTrafficShaper(ctx, "read", diags))

	read_output, err := c.ReadSecurityTrafficShaper(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityTrafficShaper(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityTrafficShaper) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityTrafficShaper")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityTrafficShaperModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityTrafficShaper(ctx, "delete", diags))

	output, err := c.DeleteSecurityTrafficShaper(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityTrafficShaper) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityTrafficShaperModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityTrafficShaper(ctx, "read", diags))

	read_output, err := c.ReadSecurityTrafficShaper(&input_model)
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

	diags.Append(data.refreshSecurityTrafficShaper(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityTrafficShaper) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityTrafficShaperModel) refreshSecurityTrafficShaper(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["guaranteedBandwidth"]; ok {
		m.GuaranteedBandwidth = parseFloat64Value(v)
	}

	if v, ok := o["maximumBandwidth"]; ok {
		m.MaximumBandwidth = parseFloat64Value(v)
	}

	if v, ok := o["bandwidthUnit"]; ok {
		m.BandwidthUnit = parseStringValue(v)
	}

	if v, ok := o["priority"]; ok {
		m.Priority = parseStringValue(v)
	}

	if v, ok := o["perPolicy"]; ok {
		m.PerPolicy = parseStringValue(v)
	}

	return diags
}

func (data *resourceSecurityTrafficShaperModel) getCreateObjectSecurityTrafficShaper(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.GuaranteedBandwidth.IsNull() && !data.GuaranteedBandwidth.IsUnknown() {
		result["guaranteedBandwidth"] = data.GuaranteedBandwidth.ValueFloat64()
	}

	if !data.MaximumBandwidth.IsNull() && !data.MaximumBandwidth.IsUnknown() {
		result["maximumBandwidth"] = data.MaximumBandwidth.ValueFloat64()
	}

	if !data.BandwidthUnit.IsNull() && !data.BandwidthUnit.IsUnknown() {
		result["bandwidthUnit"] = data.BandwidthUnit.ValueString()
	}

	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		result["priority"] = data.Priority.ValueString()
	}

	if !data.PerPolicy.IsNull() && !data.PerPolicy.IsUnknown() {
		result["perPolicy"] = data.PerPolicy.ValueString()
	}

	return &result
}

func (data *resourceSecurityTrafficShaperModel) getUpdateObjectSecurityTrafficShaper(ctx context.Context, state resourceSecurityTrafficShaperModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.GuaranteedBandwidth.IsNull() && !data.GuaranteedBandwidth.IsUnknown() {
		result["guaranteedBandwidth"] = data.GuaranteedBandwidth.ValueFloat64()
	}

	if !data.MaximumBandwidth.IsNull() && !data.MaximumBandwidth.IsUnknown() {
		result["maximumBandwidth"] = data.MaximumBandwidth.ValueFloat64()
	}

	if !data.BandwidthUnit.IsNull() && !data.BandwidthUnit.IsUnknown() {
		result["bandwidthUnit"] = data.BandwidthUnit.ValueString()
	}

	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		result["priority"] = data.Priority.ValueString()
	}

	if !data.PerPolicy.IsNull() && !data.PerPolicy.IsUnknown() {
		result["perPolicy"] = data.PerPolicy.ValueString()
	}

	return &result
}

func (data *resourceSecurityTrafficShaperModel) getURLObjectSecurityTrafficShaper(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
