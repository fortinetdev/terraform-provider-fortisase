// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceEndpointGroupInvitationCode{}
var _ resource.ResourceWithMoveState = &resourceEndpointGroupInvitationCode{}

func newResourceEndpointGroupInvitationCode() resource.Resource {
	return &resourceEndpointGroupInvitationCode{}
}

type resourceEndpointGroupInvitationCode struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointGroupInvitationCodeModel describes the resource data model.
type resourceEndpointGroupInvitationCodeModel struct {
	ID              types.String                                             `tfsdk:"id"`
	PrimaryKey      types.String                                             `tfsdk:"primary_key"`
	ExpireDate      types.String                                             `tfsdk:"expire_date"`
	GroupAssignment *resourceEndpointGroupInvitationCodeGroupAssignmentModel `tfsdk:"group_assignment"`
}

func (r *resourceEndpointGroupInvitationCode) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_group_invitation_code"
}

func (r *resourceEndpointGroupInvitationCode) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Group-Based Invitation Code Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 128),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"expire_date": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"group_assignment": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.RequiresReplace(),
						},
					},
					"group": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"id": schema.Float64Attribute{
								Validators: []validator.Float64{
									float64validatorwarning.AtLeast(1),
								},
								Computed: true,
								Optional: true,
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.RequiresReplace(),
								},
							},
							"path": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtLeast(1),
								},
								Computed: true,
								Optional: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
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

func (r *resourceEndpointGroupInvitationCode) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_group_invitation_code"
}
func (r *resourceEndpointGroupInvitationCode) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_endpoint_group_invitation_codes" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceEndpointGroupInvitationCodeModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceEndpointGroupInvitationCode) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointGroupInvitationCodes")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointGroupInvitationCodeModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointGroupInvitationCode(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCode(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateEndpointGroupInvitationCodes(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	if responseMkey, ok := getCreateResponseMkey(output, "primaryKey"); ok {
		mkey = responseMkey
	}
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCode(ctx, "read", diags))

	read_output, err := c.ReadEndpointGroupInvitationCodes(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointGroupInvitationCode(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointGroupInvitationCode) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// No update operation for this resource
	resp.Diagnostics.AddError(
		"Update not supported",
		"This resource does not support update. You use terraform taint <resource_type>.<resource_name> to force a replacement.",
	)
}

func (r *resourceEndpointGroupInvitationCode) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointGroupInvitationCodes")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceEndpointGroupInvitationCodeModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCode(ctx, "delete", diags))

	output, err := c.DeleteEndpointGroupInvitationCodes(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceEndpointGroupInvitationCode) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointGroupInvitationCodeModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCode(ctx, "read", diags))

	read_output, err := c.ReadEndpointGroupInvitationCodes(&input_model)
	if err != nil {
		if isNotFoundResponse(read_output) {
			resp.State.RemoveResource(ctx)
			return
		}
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointGroupInvitationCode(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointGroupInvitationCode) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceEndpointGroupInvitationCodeModel) refreshEndpointGroupInvitationCode(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["expireDate"]; ok {
		m.ExpireDate = parseStringValue(v)
	}

	if v, ok := o["groupAssignment"]; ok {
		m.GroupAssignment = m.GroupAssignment.flattenEndpointGroupInvitationCodeGroupAssignment(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointGroupInvitationCodeModel) getCreateObjectEndpointGroupInvitationCode(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ExpireDate.IsNull() && !data.ExpireDate.IsUnknown() {
		result["expireDate"] = data.ExpireDate.ValueString()
	}

	if data.GroupAssignment != nil && !isZeroStruct(*data.GroupAssignment) {
		result["groupAssignment"] = data.GroupAssignment.expandEndpointGroupInvitationCodeGroupAssignment(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointGroupInvitationCodeModel) getUpdateObjectEndpointGroupInvitationCode(ctx context.Context, state resourceEndpointGroupInvitationCodeModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ExpireDate.IsNull() && !data.ExpireDate.IsUnknown() {
		result["expireDate"] = data.ExpireDate.ValueString()
	}

	if data.GroupAssignment != nil {
		result["groupAssignment"] = data.GroupAssignment.expandEndpointGroupInvitationCodeGroupAssignment(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointGroupInvitationCodeModel) getURLObjectEndpointGroupInvitationCode(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointGroupInvitationCodeGroupAssignmentModel struct {
	Enabled types.Bool                                                    `tfsdk:"enabled"`
	Group   *resourceEndpointGroupInvitationCodeGroupAssignmentGroupModel `tfsdk:"group"`
}

type resourceEndpointGroupInvitationCodeGroupAssignmentGroupModel struct {
	Id   types.Float64 `tfsdk:"id"`
	Path types.String  `tfsdk:"path"`
}

func (m *resourceEndpointGroupInvitationCodeGroupAssignmentModel) flattenEndpointGroupInvitationCodeGroupAssignment(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointGroupInvitationCodeGroupAssignmentModel {
	if input == nil {
		return &resourceEndpointGroupInvitationCodeGroupAssignmentModel{}
	}
	if m == nil {
		m = &resourceEndpointGroupInvitationCodeGroupAssignmentModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenEndpointGroupInvitationCodeGroupAssignmentGroup(ctx, v, diags)
	}

	return m
}

func (m *resourceEndpointGroupInvitationCodeGroupAssignmentGroupModel) flattenEndpointGroupInvitationCodeGroupAssignmentGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointGroupInvitationCodeGroupAssignmentGroupModel {
	if input == nil {
		return &resourceEndpointGroupInvitationCodeGroupAssignmentGroupModel{}
	}
	if m == nil {
		m = &resourceEndpointGroupInvitationCodeGroupAssignmentGroupModel{}
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

func (data *resourceEndpointGroupInvitationCodeGroupAssignmentModel) expandEndpointGroupInvitationCodeGroupAssignment(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if data.Group != nil && !isZeroStruct(*data.Group) {
		result["group"] = data.Group.expandEndpointGroupInvitationCodeGroupAssignmentGroup(ctx, diags)
	}

	return result
}

func (data *resourceEndpointGroupInvitationCodeGroupAssignmentGroupModel) expandEndpointGroupInvitationCodeGroupAssignmentGroup(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		result["id"] = data.Id.ValueFloat64()
	}

	if !data.Path.IsNull() && !data.Path.IsUnknown() {
		result["path"] = data.Path.ValueString()
	}

	return result
}
