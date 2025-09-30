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
var _ resource.Resource = &resourceSecurityEndpointToEndpointPoliciesClone2Edl{}

func newResourceSecurityEndpointToEndpointPoliciesClone() resource.Resource {
	return &resourceSecurityEndpointToEndpointPoliciesClone2Edl{}
}

type resourceSecurityEndpointToEndpointPoliciesClone2Edl struct {
	fortiClient *FortiClient
}

// resourceSecurityEndpointToEndpointPoliciesClone2EdlModel describes the resource data model.
type resourceSecurityEndpointToEndpointPoliciesClone2EdlModel struct {
	ID         types.String `tfsdk:"id"`
	PrimaryKey types.String `tfsdk:"primary_key"`
	BasedOn    types.String `tfsdk:"based_on"`
}

func (r *resourceSecurityEndpointToEndpointPoliciesClone2Edl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_endpoint_to_endpoint_policies_clone"
}

func (r *resourceSecurityEndpointToEndpointPoliciesClone2Edl) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthBetween(1, 35),
				},
				Computed: true,
				Optional: true,
			},
			"based_on": schema.StringAttribute{
				Description: "The policy you what to clone.",
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func (r *resourceSecurityEndpointToEndpointPoliciesClone2Edl) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityEndpointToEndpointPoliciesClone2Edl) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityEndpointToEndpointPoliciesClone2EdlModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityEndpointToEndpointPoliciesClone(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPoliciesClone(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	_, err := c.CreateSecurityEndpointToEndpointPoliciesClone(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := "SecurityEndpointToEndpointPoliciesClone"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityEndpointToEndpointPoliciesClone2Edl) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityEndpointToEndpointPoliciesClone2EdlModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityEndpointToEndpointPoliciesClone2EdlModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityEndpointToEndpointPoliciesClone(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityEndpointToEndpointPoliciesClone(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.CreateSecurityEndpointToEndpointPoliciesClone(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityEndpointToEndpointPoliciesClone2Edl) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityEndpointToEndpointPoliciesClone2Edl) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourceSecurityEndpointToEndpointPoliciesClone2EdlModel) getCreateObjectSecurityEndpointToEndpointPoliciesClone(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

func (data *resourceSecurityEndpointToEndpointPoliciesClone2EdlModel) getUpdateObjectSecurityEndpointToEndpointPoliciesClone(ctx context.Context, state resourceSecurityEndpointToEndpointPoliciesClone2EdlModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

func (data *resourceSecurityEndpointToEndpointPoliciesClone2EdlModel) getURLObjectSecurityEndpointToEndpointPoliciesClone(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.BasedOn.IsNull() {
		result["based_on"] = data.BasedOn.ValueString()
	}

	return &result
}
