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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceInfraSecureWebGatewaySupplementaryData{}

func newResourceInfraSecureWebGatewaySupplementaryData() resource.Resource {
	return &resourceInfraSecureWebGatewaySupplementaryData{}
}

type resourceInfraSecureWebGatewaySupplementaryData struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceInfraSecureWebGatewaySupplementaryDataModel describes the resource data model.
type resourceInfraSecureWebGatewaySupplementaryDataModel struct {
	ID                   types.String  `tfsdk:"id"`
	PrimaryKey           types.String  `tfsdk:"primary_key"`
	SessionDurationHours types.Float64 `tfsdk:"session_duration_hours"`
	EndSessionAfterMins  types.Float64 `tfsdk:"end_session_after_mins"`
}

func (r *resourceInfraSecureWebGatewaySupplementaryData) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_secure_web_gateway_supplementary_data"
}

func (r *resourceInfraSecureWebGatewaySupplementaryData) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.OneOf("$sase-global"),
				},
				Default:  stringdefault.StaticString("$sase-global"),
				Computed: true,
				Optional: true,
			},
			"session_duration_hours": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"end_session_after_mins": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceInfraSecureWebGatewaySupplementaryData) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_infra_secure_web_gateway_supplementary_data"
}

func (r *resourceInfraSecureWebGatewaySupplementaryData) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("InfraSecureWebGatewaySupplementaryData")
	lock.Lock()
	defer lock.Unlock()
	var data resourceInfraSecureWebGatewaySupplementaryDataModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectInfraSecureWebGatewaySupplementaryData(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateInfraSecureWebGatewaySupplementaryData(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output, err := c.ReadInfraSecureWebGatewaySupplementaryData(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraSecureWebGatewaySupplementaryData(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraSecureWebGatewaySupplementaryData) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("InfraSecureWebGatewaySupplementaryData")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceInfraSecureWebGatewaySupplementaryDataModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceInfraSecureWebGatewaySupplementaryDataModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectInfraSecureWebGatewaySupplementaryData(ctx, state, diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateInfraSecureWebGatewaySupplementaryData(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output, err := c.ReadInfraSecureWebGatewaySupplementaryData(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraSecureWebGatewaySupplementaryData(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraSecureWebGatewaySupplementaryData) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceInfraSecureWebGatewaySupplementaryData) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceInfraSecureWebGatewaySupplementaryDataModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadInfraSecureWebGatewaySupplementaryData(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraSecureWebGatewaySupplementaryData(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraSecureWebGatewaySupplementaryData) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceInfraSecureWebGatewaySupplementaryDataModel) refreshInfraSecureWebGatewaySupplementaryData(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["sessionDurationHours"]; ok {
		m.SessionDurationHours = parseFloat64Value(v)
	}

	if v, ok := o["endSessionAfterMins"]; ok {
		m.EndSessionAfterMins = parseFloat64Value(v)
	}

	return diags
}

func (data *resourceInfraSecureWebGatewaySupplementaryDataModel) getCreateObjectInfraSecureWebGatewaySupplementaryData(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.SessionDurationHours.IsNull() {
		result["sessionDurationHours"] = data.SessionDurationHours.ValueFloat64()
	}

	if !data.EndSessionAfterMins.IsNull() {
		result["endSessionAfterMins"] = data.EndSessionAfterMins.ValueFloat64()
	}

	return &result
}

func (data *resourceInfraSecureWebGatewaySupplementaryDataModel) getUpdateObjectInfraSecureWebGatewaySupplementaryData(ctx context.Context, state resourceInfraSecureWebGatewaySupplementaryDataModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.SessionDurationHours.IsNull() {
		result["sessionDurationHours"] = data.SessionDurationHours.ValueFloat64()
	}

	if !data.EndSessionAfterMins.IsNull() {
		result["endSessionAfterMins"] = data.EndSessionAfterMins.ValueFloat64()
	}

	return &result
}
