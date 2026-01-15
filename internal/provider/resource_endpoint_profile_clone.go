// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceEndpointProfileClone2Edl{}

func newResourceEndpointProfileClone() resource.Resource {
	return &resourceEndpointProfileClone2Edl{}
}

type resourceEndpointProfileClone2Edl struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointProfileClone2EdlModel describes the resource data model.
type resourceEndpointProfileClone2EdlModel struct {
	ID                              types.String `tfsdk:"id"`
	PrimaryKey                      types.String `tfsdk:"primary_key"`
	Enabled                         types.Bool   `tfsdk:"enabled"`
	SkipOffNetProfileCreationOnEdit types.Bool   `tfsdk:"skip_off_net_profile_creation_on_edit"`
	BasedOn                         types.String `tfsdk:"based_on"`
}

func (r *resourceEndpointProfileClone2Edl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_profile_clone"
}

func (r *resourceEndpointProfileClone2Edl) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthBetween(1, 128),
				},
				Computed: true,
				Optional: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"skip_off_net_profile_creation_on_edit": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"based_on": schema.StringAttribute{
				MarkdownDescription: "The endpoint profile you what to clone.",
				Computed:            true,
				Optional:            true,
			},
		},
	}
}

func (r *resourceEndpointProfileClone2Edl) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_profile_clone"
}

func (r *resourceEndpointProfileClone2Edl) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointProfileClone2EdlModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointProfileClone(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointProfileClone(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateEndpointProfileClone(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := "EndpointProfileClone"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointProfileClone2Edl) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointProfileClone2EdlModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointProfileClone2EdlModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointProfileClone(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointProfileClone(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.CreateEndpointProfileClone(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointProfileClone2Edl) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointProfileClone2Edl) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourceEndpointProfileClone2EdlModel) getCreateObjectEndpointProfileClone(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.SkipOffNetProfileCreationOnEdit.IsNull() {
		result["skipOffNetProfileCreationOnEdit"] = data.SkipOffNetProfileCreationOnEdit.ValueBool()
	}

	return &result
}

func (data *resourceEndpointProfileClone2EdlModel) getUpdateObjectEndpointProfileClone(ctx context.Context, state resourceEndpointProfileClone2EdlModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.SkipOffNetProfileCreationOnEdit.IsNull() {
		result["skipOffNetProfileCreationOnEdit"] = data.SkipOffNetProfileCreationOnEdit.ValueBool()
	}

	return &result
}

func (data *resourceEndpointProfileClone2EdlModel) getURLObjectEndpointProfileClone(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.BasedOn.IsNull() {
		result["based_on"] = data.BasedOn.ValueString()
	}

	return &result
}
