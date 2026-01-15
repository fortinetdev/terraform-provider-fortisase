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
var _ resource.Resource = &resourceEndpointsAccessProxyAuthorize2Edl{}

func newResourceEndpointsAccessProxyAuthorize() resource.Resource {
	return &resourceEndpointsAccessProxyAuthorize2Edl{}
}

type resourceEndpointsAccessProxyAuthorize2Edl struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointsAccessProxyAuthorize2EdlModel describes the resource data model.
type resourceEndpointsAccessProxyAuthorize2EdlModel struct {
	ID     types.String `tfsdk:"id"`
	SnList types.Set    `tfsdk:"sn_list"`
}

func (r *resourceEndpointsAccessProxyAuthorize2Edl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_access_proxy_authorize"
}

func (r *resourceEndpointsAccessProxyAuthorize2Edl) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *resourceEndpointsAccessProxyAuthorize2Edl) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoints_access_proxy_authorize"
}

func (r *resourceEndpointsAccessProxyAuthorize2Edl) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointsAccessProxyAuthorize2EdlModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointsAccessProxyAuthorize(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateEndpointsAccessProxyAuthorize(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := "EndpointsAccessProxyAuthorize"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointsAccessProxyAuthorize2Edl) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointsAccessProxyAuthorize2EdlModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointsAccessProxyAuthorize2EdlModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointsAccessProxyAuthorize(ctx, state, diags))

	if diags.HasError() {
		return
	}

	output, err := c.CreateEndpointsAccessProxyAuthorize(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointsAccessProxyAuthorize2Edl) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointsAccessProxyAuthorize2Edl) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourceEndpointsAccessProxyAuthorize2EdlModel) getCreateObjectEndpointsAccessProxyAuthorize(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SnList.IsNull() {
		result["snList"] = expandSetToStringList(data.SnList)
	}

	return &result
}

func (data *resourceEndpointsAccessProxyAuthorize2EdlModel) getUpdateObjectEndpointsAccessProxyAuthorize(ctx context.Context, state resourceEndpointsAccessProxyAuthorize2EdlModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SnList.IsNull() {
		result["snList"] = expandSetToStringList(data.SnList)
	}

	return &result
}
