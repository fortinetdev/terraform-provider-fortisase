// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceEndpointsDisableManagement{}

func newResourceEndpointsDisableManagement() resource.Resource {
	return &resourceEndpointsDisableManagement{}
}

type resourceEndpointsDisableManagement struct {
	fortiClient *FortiClient
}

// resourceEndpointsDisableManagementModel describes the resource data model.
type resourceEndpointsDisableManagementModel struct {
	ID        types.String                                       `tfsdk:"id"`
	Endpoints []resourceEndpointsDisableManagementEndpointsModel `tfsdk:"endpoints"`
}

func (r *resourceEndpointsDisableManagement) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_disable_management"
}

func (r *resourceEndpointsDisableManagement) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"endpoints": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"device_id": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"hostname": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceEndpointsDisableManagement) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceEndpointsDisableManagement) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointsDisableManagementModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointsDisableManagement(ctx, diags))

	if diags.HasError() {
		return
	}
	_, err := c.CreateEndpointsDisableManagement(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := "EndpointsDisableManagement"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointsDisableManagement) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointsDisableManagementModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointsDisableManagementModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointsDisableManagement(ctx, state, diags))

	if diags.HasError() {
		return
	}

	_, err := c.CreateEndpointsDisableManagement(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointsDisableManagement) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointsDisableManagement) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourceEndpointsDisableManagementModel) getCreateObjectEndpointsDisableManagement(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})

	result["endpoints"] = data.expandEndpointsDisableManagementEndpointsList(ctx, data.Endpoints, diags)

	return &result
}

func (data *resourceEndpointsDisableManagementModel) getUpdateObjectEndpointsDisableManagement(ctx context.Context, state resourceEndpointsDisableManagementModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if len(data.Endpoints) > 0 || !isSameStruct(data.Endpoints, state.Endpoints) {
		result["endpoints"] = data.expandEndpointsDisableManagementEndpointsList(ctx, data.Endpoints, diags)
	}

	return &result
}

type resourceEndpointsDisableManagementEndpointsModel struct {
	DeviceId types.String `tfsdk:"device_id"`
	Hostname types.String `tfsdk:"hostname"`
}

func (data *resourceEndpointsDisableManagementEndpointsModel) expandEndpointsDisableManagementEndpoints(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.DeviceId.IsNull() {
		result["deviceId"] = data.DeviceId.ValueString()
	}

	if !data.Hostname.IsNull() {
		result["hostname"] = data.Hostname.ValueString()
	}

	return result
}

func (s *resourceEndpointsDisableManagementModel) expandEndpointsDisableManagementEndpointsList(ctx context.Context, l []resourceEndpointsDisableManagementEndpointsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointsDisableManagementEndpoints(ctx, diags)
	}
	return result
}
