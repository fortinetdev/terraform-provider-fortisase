// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/setvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceEndpointsQuarantine{}

func newResourceEndpointsQuarantine() resource.Resource {
	return &resourceEndpointsQuarantine{}
}

type resourceEndpointsQuarantine struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointsQuarantineModel describes the resource data model.
type resourceEndpointsQuarantineModel struct {
	ID        types.String `tfsdk:"id"`
	DeviceIds types.Set    `tfsdk:"device_ids"`
}

func (r *resourceEndpointsQuarantine) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_quarantine"
}

func (r *resourceEndpointsQuarantine) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Endpoint management monitor API for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"device_ids": schema.SetAttribute{
				Validators: []validator.Set{
					setvalidatorwarning.SizeAtLeast(1),
				},
				Optional:    true,
				ElementType: types.Int64Type,
			},
		},
	}
}

func (r *resourceEndpointsQuarantine) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoints_quarantine"
}

func (r *resourceEndpointsQuarantine) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointsQuarantineModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointsQuarantine(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateEndpointsQuarantine(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := "EndpointsQuarantine"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointsQuarantine) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointsQuarantineModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointsQuarantineModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointsQuarantine(ctx, state, diags))

	if diags.HasError() {
		return
	}

	output, err := c.CreateEndpointsQuarantine(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointsQuarantine) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointsQuarantine) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourceEndpointsQuarantineModel) getCreateObjectEndpointsQuarantine(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.DeviceIds.IsNull() && !data.DeviceIds.IsUnknown() {
		result["deviceIds"] = expandSetToInt64List(data.DeviceIds)
	}

	return &result
}

func (data *resourceEndpointsQuarantineModel) getUpdateObjectEndpointsQuarantine(ctx context.Context, state resourceEndpointsQuarantineModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.DeviceIds.IsNull() && !data.DeviceIds.IsUnknown() {
		result["deviceIds"] = expandSetToInt64List(data.DeviceIds)
	}

	return &result
}
