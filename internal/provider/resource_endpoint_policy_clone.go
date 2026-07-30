// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceEndpointPolicyClone2Edl{}
var _ resource.ResourceWithMoveState = &resourceEndpointPolicyClone2Edl{}

func newResourceEndpointPolicyClone() resource.Resource {
	return &resourceEndpointPolicyClone2Edl{}
}

type resourceEndpointPolicyClone2Edl struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointPolicyClone2EdlModel describes the resource data model.
type resourceEndpointPolicyClone2EdlModel struct {
	ID                              types.String `tfsdk:"id"`
	PrimaryKey                      types.String `tfsdk:"primary_key"`
	Enabled                         types.Bool   `tfsdk:"enabled"`
	SkipOffNetProfileCreationOnEdit types.Bool   `tfsdk:"skip_off_net_profile_creation_on_edit"`
	BasedOn                         types.String `tfsdk:"based_on"`
}

func (r *resourceEndpointPolicyClone2Edl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_policy_clone"
}

func (r *resourceEndpointPolicyClone2Edl) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Endpoint Policy Resource API V2 for FortiSASE.",
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
				Optional: true,
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
			},
			"skip_off_net_profile_creation_on_edit": schema.BoolAttribute{
				Optional: true,
			},
			"based_on": schema.StringAttribute{
				MarkdownDescription: "The endpoint profile you what to clone.",
				Optional:            true,
			},
		},
	}
}

func (r *resourceEndpointPolicyClone2Edl) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_policy_clone"
}
func (r *resourceEndpointPolicyClone2Edl) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_endpoint_policies_clone" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceEndpointPolicyClone2EdlModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceEndpointPolicyClone2Edl) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointPolicyClone2EdlModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointPolicyClone(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointPolicyClone(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateEndpointPoliciesClone(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := "EndpointPoliciesClone"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointPolicyClone2Edl) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointPolicyClone2EdlModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointPolicyClone2EdlModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointPolicyClone(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointPolicyClone(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.CreateEndpointPoliciesClone(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointPolicyClone2Edl) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointPolicyClone2Edl) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourceEndpointPolicyClone2EdlModel) getCreateObjectEndpointPolicyClone(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.SkipOffNetProfileCreationOnEdit.IsNull() && !data.SkipOffNetProfileCreationOnEdit.IsUnknown() {
		result["skipOffNetProfileCreationOnEdit"] = data.SkipOffNetProfileCreationOnEdit.ValueBool()
	}

	return &result
}

func (data *resourceEndpointPolicyClone2EdlModel) getUpdateObjectEndpointPolicyClone(ctx context.Context, state resourceEndpointPolicyClone2EdlModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.SkipOffNetProfileCreationOnEdit.IsNull() && !data.SkipOffNetProfileCreationOnEdit.IsUnknown() {
		result["skipOffNetProfileCreationOnEdit"] = data.SkipOffNetProfileCreationOnEdit.ValueBool()
	}

	return &result
}

func (data *resourceEndpointPolicyClone2EdlModel) getURLObjectEndpointPolicyClone(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.BasedOn.IsNull() && !data.BasedOn.IsUnknown() {
		result["based_on"] = data.BasedOn.ValueString()
	}

	return &result
}
