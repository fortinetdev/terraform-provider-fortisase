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
var _ resource.Resource = &resourceUserSwgSessionsDeauth2Edl{}

func newResourceUserSwgSessionsDeauth() resource.Resource {
	return &resourceUserSwgSessionsDeauth2Edl{}
}

type resourceUserSwgSessionsDeauth2Edl struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceUserSwgSessionsDeauth2EdlModel describes the resource data model.
type resourceUserSwgSessionsDeauth2EdlModel struct {
	ID         types.String `tfsdk:"id"`
	Usernames  types.Set    `tfsdk:"usernames"`
	SessionIds types.Set    `tfsdk:"session_ids"`
}

func (r *resourceUserSwgSessionsDeauth2Edl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_swg_sessions_deauth"
}

func (r *resourceUserSwgSessionsDeauth2Edl) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"usernames": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"session_ids": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *resourceUserSwgSessionsDeauth2Edl) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_user_swg_sessions_deauth"
}

func (r *resourceUserSwgSessionsDeauth2Edl) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceUserSwgSessionsDeauth2EdlModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectUserSwgSessionsDeauth(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateUserSwgSessionsDeauth(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := "UserSwgSessionsDeauth"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceUserSwgSessionsDeauth2Edl) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceUserSwgSessionsDeauth2EdlModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceUserSwgSessionsDeauth2EdlModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectUserSwgSessionsDeauth(ctx, state, diags))

	if diags.HasError() {
		return
	}

	output, err := c.CreateUserSwgSessionsDeauth(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceUserSwgSessionsDeauth2Edl) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceUserSwgSessionsDeauth2Edl) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourceUserSwgSessionsDeauth2EdlModel) getCreateObjectUserSwgSessionsDeauth(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Usernames.IsNull() {
		result["usernames"] = expandSetToStringList(data.Usernames)
	}

	if !data.SessionIds.IsNull() {
		result["sessionIds"] = expandSetToStringList(data.SessionIds)
	}

	return &result
}

func (data *resourceUserSwgSessionsDeauth2EdlModel) getUpdateObjectUserSwgSessionsDeauth(ctx context.Context, state resourceUserSwgSessionsDeauth2EdlModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Usernames.IsNull() {
		result["usernames"] = expandSetToStringList(data.Usernames)
	}

	if !data.SessionIds.IsNull() {
		result["sessionIds"] = expandSetToStringList(data.SessionIds)
	}

	return &result
}
