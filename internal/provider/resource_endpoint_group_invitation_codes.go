// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
var _ resource.Resource = &resourceEndpointGroupInvitationCodes{}

func newResourceEndpointGroupInvitationCodes() resource.Resource {
	return &resourceEndpointGroupInvitationCodes{}
}

type resourceEndpointGroupInvitationCodes struct {
	fortiClient *FortiClient
}

// resourceEndpointGroupInvitationCodesModel describes the resource data model.
type resourceEndpointGroupInvitationCodesModel struct {
	ID              types.String                                              `tfsdk:"id"`
	PrimaryKey      types.String                                              `tfsdk:"primary_key"`
	ExpireDate      types.String                                              `tfsdk:"expire_date"`
	GroupAssignment *resourceEndpointGroupInvitationCodesGroupAssignmentModel `tfsdk:"group_assignment"`
}

func (r *resourceEndpointGroupInvitationCodes) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_group_invitation_codes"
}

func (r *resourceEndpointGroupInvitationCodes) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Required: true,
			},
			"expire_date": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"group_assignment": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Computed: true,
						Optional: true,
					},
					"group": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"id": schema.Float64Attribute{
								Validators: []validator.Float64{
									float64validator.AtLeast(1),
								},
								Computed: true,
								Optional: true,
							},
							"path": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceEndpointGroupInvitationCodes) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceEndpointGroupInvitationCodes) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointGroupInvitationCodesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointGroupInvitationCodes(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCodes(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateEndpointGroupInvitationCodes(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCodes(ctx, "read", diags))

	read_output, err := c.ReadEndpointGroupInvitationCodes(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointGroupInvitationCodes(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointGroupInvitationCodes) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointGroupInvitationCodesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointGroupInvitationCodesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointGroupInvitationCodes(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCodes(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateEndpointGroupInvitationCodes(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCodes(ctx, "read", diags))

	read_output, err := c.ReadEndpointGroupInvitationCodes(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointGroupInvitationCodes(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointGroupInvitationCodes) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointGroupInvitationCodesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCodes(ctx, "delete", diags))

	err := c.DeleteEndpointGroupInvitationCodes(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceEndpointGroupInvitationCodes) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointGroupInvitationCodesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCodes(ctx, "read", diags))

	read_output, err := c.ReadEndpointGroupInvitationCodes(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointGroupInvitationCodes(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointGroupInvitationCodes) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointGroupInvitationCodesModel) refreshEndpointGroupInvitationCodes(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["expireDate"]; ok {
		m.ExpireDate = parseStringValue(v)
	}

	if v, ok := o["groupAssignment"]; ok {
		m.GroupAssignment = m.GroupAssignment.flattenEndpointGroupInvitationCodesGroupAssignment(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointGroupInvitationCodesModel) getCreateObjectEndpointGroupInvitationCodes(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ExpireDate.IsNull() {
		result["expireDate"] = data.ExpireDate.ValueString()
	}

	if data.GroupAssignment != nil && !isZeroStruct(*data.GroupAssignment) {
		result["groupAssignment"] = data.GroupAssignment.expandEndpointGroupInvitationCodesGroupAssignment(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointGroupInvitationCodesModel) getUpdateObjectEndpointGroupInvitationCodes(ctx context.Context, state resourceEndpointGroupInvitationCodesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ExpireDate.IsNull() && !data.ExpireDate.Equal(state.ExpireDate) {
		result["expireDate"] = data.ExpireDate.ValueString()
	}

	if data.GroupAssignment != nil && !isSameStruct(data.GroupAssignment, state.GroupAssignment) {
		result["groupAssignment"] = data.GroupAssignment.expandEndpointGroupInvitationCodesGroupAssignment(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointGroupInvitationCodesModel) getURLObjectEndpointGroupInvitationCodes(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointGroupInvitationCodesGroupAssignmentModel struct {
	Enabled types.Bool                                                     `tfsdk:"enabled"`
	Group   *resourceEndpointGroupInvitationCodesGroupAssignmentGroupModel `tfsdk:"group"`
}

type resourceEndpointGroupInvitationCodesGroupAssignmentGroupModel struct {
	Id   types.Float64 `tfsdk:"id"`
	Path types.String  `tfsdk:"path"`
}

func (m *resourceEndpointGroupInvitationCodesGroupAssignmentModel) flattenEndpointGroupInvitationCodesGroupAssignment(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointGroupInvitationCodesGroupAssignmentModel {
	if input == nil {
		return &resourceEndpointGroupInvitationCodesGroupAssignmentModel{}
	}
	if m == nil {
		m = &resourceEndpointGroupInvitationCodesGroupAssignmentModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenEndpointGroupInvitationCodesGroupAssignmentGroup(ctx, v, diags)
	}

	return m
}

func (m *resourceEndpointGroupInvitationCodesGroupAssignmentGroupModel) flattenEndpointGroupInvitationCodesGroupAssignmentGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointGroupInvitationCodesGroupAssignmentGroupModel {
	if input == nil {
		return &resourceEndpointGroupInvitationCodesGroupAssignmentGroupModel{}
	}
	if m == nil {
		m = &resourceEndpointGroupInvitationCodesGroupAssignmentGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	if v, ok := o["path"]; ok {
		m.Path = parseStringValue(v)
	}

	return m
}

func (data *resourceEndpointGroupInvitationCodesGroupAssignmentModel) expandEndpointGroupInvitationCodesGroupAssignment(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandEndpointGroupInvitationCodesGroupAssignmentGroup(ctx, diags)
	}

	return result
}

func (data *resourceEndpointGroupInvitationCodesGroupAssignmentGroupModel) expandEndpointGroupInvitationCodesGroupAssignmentGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() {
		result["id"] = data.Id.ValueFloat64()
	}

	if !data.Path.IsNull() {
		result["path"] = data.Path.ValueString()
	}

	return result
}
