// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceEndpointsEnableManagement{}

func newResourceEndpointsEnableManagement() resource.Resource {
	return &resourceEndpointsEnableManagement{}
}

type resourceEndpointsEnableManagement struct {
	fortiClient *FortiClient
}

// resourceEndpointsEnableManagementModel describes the resource data model.
type resourceEndpointsEnableManagementModel struct {
	ID        types.String `tfsdk:"id"`
	DeviceIds types.Set    `tfsdk:"device_ids"`
}

func (r *resourceEndpointsEnableManagement) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_enable_management"
}

func (r *resourceEndpointsEnableManagement) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
					setvalidator.SizeAtLeast(1),
				},
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *resourceEndpointsEnableManagement) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceEndpointsEnableManagement) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointsEnableManagementModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointsEnableManagement(ctx, diags))

	if diags.HasError() {
		return
	}
	_, err := c.CreateEndpointsEnableManagement(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := "EndpointsEnableManagement"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointsEnableManagement) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointsEnableManagementModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointsEnableManagementModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointsEnableManagement(ctx, state, diags))

	if diags.HasError() {
		return
	}

	_, err := c.CreateEndpointsEnableManagement(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointsEnableManagement) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointsEnableManagement) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourceEndpointsEnableManagementModel) getCreateObjectEndpointsEnableManagement(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.DeviceIds.IsNull() {
		result["deviceIds"] = expandSetToStringList(data.DeviceIds)
	}

	return &result
}

func (data *resourceEndpointsEnableManagementModel) getUpdateObjectEndpointsEnableManagement(ctx context.Context, state resourceEndpointsEnableManagementModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.DeviceIds.IsNull() {
		result["deviceIds"] = expandSetToStringList(data.DeviceIds)
	}

	return &result
}
