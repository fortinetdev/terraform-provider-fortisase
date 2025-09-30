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
var _ resource.Resource = &resourceEndpointsAccessProxyDisconnect2Edl{}

func newResourceEndpointsAccessProxyDisconnect() resource.Resource {
	return &resourceEndpointsAccessProxyDisconnect2Edl{}
}

type resourceEndpointsAccessProxyDisconnect2Edl struct {
	fortiClient *FortiClient
}

// resourceEndpointsAccessProxyDisconnect2EdlModel describes the resource data model.
type resourceEndpointsAccessProxyDisconnect2EdlModel struct {
	ID     types.String `tfsdk:"id"`
	SnList types.Set    `tfsdk:"sn_list"`
}

func (r *resourceEndpointsAccessProxyDisconnect2Edl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_access_proxy_disconnect"
}

func (r *resourceEndpointsAccessProxyDisconnect2Edl) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"sn_list": schema.SetAttribute{
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

func (r *resourceEndpointsAccessProxyDisconnect2Edl) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceEndpointsAccessProxyDisconnect2Edl) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointsAccessProxyDisconnect2EdlModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointsAccessProxyDisconnect(ctx, diags))

	if diags.HasError() {
		return
	}
	_, err := c.CreateEndpointsAccessProxyDisconnect(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := "EndpointsAccessProxyDisconnect"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointsAccessProxyDisconnect2Edl) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointsAccessProxyDisconnect2EdlModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointsAccessProxyDisconnect2EdlModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointsAccessProxyDisconnect(ctx, state, diags))

	if diags.HasError() {
		return
	}

	_, err := c.CreateEndpointsAccessProxyDisconnect(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointsAccessProxyDisconnect2Edl) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointsAccessProxyDisconnect2Edl) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourceEndpointsAccessProxyDisconnect2EdlModel) getCreateObjectEndpointsAccessProxyDisconnect(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SnList.IsNull() {
		result["snList"] = expandSetToStringList(data.SnList)
	}

	return &result
}

func (data *resourceEndpointsAccessProxyDisconnect2EdlModel) getUpdateObjectEndpointsAccessProxyDisconnect(ctx context.Context, state resourceEndpointsAccessProxyDisconnect2EdlModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SnList.IsNull() {
		result["snList"] = expandSetToStringList(data.SnList)
	}

	return &result
}
