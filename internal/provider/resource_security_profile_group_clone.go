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
var _ resource.Resource = &resourceSecurityProfileGroupClone2Edl{}

func newResourceSecurityProfileGroupClone() resource.Resource {
	return &resourceSecurityProfileGroupClone2Edl{}
}

type resourceSecurityProfileGroupClone2Edl struct {
	fortiClient *FortiClient
}

// resourceSecurityProfileGroupClone2EdlModel describes the resource data model.
type resourceSecurityProfileGroupClone2EdlModel struct {
	ID         types.String `tfsdk:"id"`
	PrimaryKey types.String `tfsdk:"primary_key"`
	Direction  types.String `tfsdk:"direction"`
	BasedOn    types.String `tfsdk:"based_on"`
}

func (r *resourceSecurityProfileGroupClone2Edl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_profile_group_clone"
}

func (r *resourceSecurityProfileGroupClone2Edl) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Computed: true,
				Optional: true,
			},
			"direction": schema.StringAttribute{
				Description: "The direction of the target resource.",
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				Computed: true,
				Optional: true,
			},
			"based_on": schema.StringAttribute{
				Description: "The profile group you what to clone.",
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func (r *resourceSecurityProfileGroupClone2Edl) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityProfileGroupClone2Edl) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityProfileGroupClone2EdlModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityProfileGroupClone(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityProfileGroupClone(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	_, err := c.CreateSecurityProfileGroupClone(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := "SecurityProfileGroupClone"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityProfileGroupClone2Edl) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityProfileGroupClone2EdlModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityProfileGroupClone2EdlModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityProfileGroupClone(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityProfileGroupClone(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.CreateSecurityProfileGroupClone(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityProfileGroupClone2Edl) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityProfileGroupClone2Edl) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourceSecurityProfileGroupClone2EdlModel) getCreateObjectSecurityProfileGroupClone(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

func (data *resourceSecurityProfileGroupClone2EdlModel) getUpdateObjectSecurityProfileGroupClone(ctx context.Context, state resourceSecurityProfileGroupClone2EdlModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

func (data *resourceSecurityProfileGroupClone2EdlModel) getURLObjectSecurityProfileGroupClone(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	if !data.BasedOn.IsNull() {
		result["based_on"] = data.BasedOn.ValueString()
	}

	return &result
}
