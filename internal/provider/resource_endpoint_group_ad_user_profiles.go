// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceEndpointGroupAdUserProfiles{}

func newResourceEndpointGroupAdUserProfiles() resource.Resource {
	return &resourceEndpointGroupAdUserProfiles{}
}

type resourceEndpointGroupAdUserProfiles struct {
	fortiClient *FortiClient
}

// resourceEndpointGroupAdUserProfilesModel describes the resource data model.
type resourceEndpointGroupAdUserProfilesModel struct {
	ID         types.String `tfsdk:"id"`
	AdUserIds  types.Set    `tfsdk:"ad_user_ids"`
	GroupIds   types.Set    `tfsdk:"group_ids"`
	PrimaryKey types.String `tfsdk:"primary_key"`
}

func (r *resourceEndpointGroupAdUserProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_group_ad_user_profiles"
}

func (r *resourceEndpointGroupAdUserProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ad_user_ids": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.Int64Type,
			},
			"group_ids": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.Int64Type,
			},
			"primary_key": schema.StringAttribute{
				Description: "The primary key of the object. Can be found in the response from the get request.",
				Required:    true,
			},
		},
	}
}

func (r *resourceEndpointGroupAdUserProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceEndpointGroupAdUserProfiles) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointGroupAdUserProfilesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointGroupAdUserProfiles(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointGroupAdUserProfiles(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateEndpointGroupAdUserProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointGroupAdUserProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointGroupAdUserProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointGroupAdUserProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointGroupAdUserProfiles) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointGroupAdUserProfilesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointGroupAdUserProfilesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointGroupAdUserProfiles(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointGroupAdUserProfiles(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateEndpointGroupAdUserProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointGroupAdUserProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointGroupAdUserProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointGroupAdUserProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointGroupAdUserProfiles) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointGroupAdUserProfiles) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointGroupAdUserProfilesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointGroupAdUserProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointGroupAdUserProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointGroupAdUserProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointGroupAdUserProfiles) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointGroupAdUserProfilesModel) refreshEndpointGroupAdUserProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["adUserIds"]; ok {
		m.AdUserIds = parseSetValue(ctx, v, types.Int64Type)
	}

	if v, ok := o["groupIds"]; ok {
		m.GroupIds = parseSetValue(ctx, v, types.Int64Type)
	}

	return diags
}

func (data *resourceEndpointGroupAdUserProfilesModel) getCreateObjectEndpointGroupAdUserProfiles(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AdUserIds.IsNull() {
		result["adUserIds"] = expandSetToStringList(data.AdUserIds)
	}

	if !data.GroupIds.IsNull() {
		result["groupIds"] = expandSetToStringList(data.GroupIds)
	}

	return &result
}

func (data *resourceEndpointGroupAdUserProfilesModel) getUpdateObjectEndpointGroupAdUserProfiles(ctx context.Context, state resourceEndpointGroupAdUserProfilesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AdUserIds.IsNull() {
		result["adUserIds"] = expandSetToStringList(data.AdUserIds)
	}

	if !data.GroupIds.IsNull() {
		result["groupIds"] = expandSetToStringList(data.GroupIds)
	}

	return &result
}

func (data *resourceEndpointGroupAdUserProfilesModel) getURLObjectEndpointGroupAdUserProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
