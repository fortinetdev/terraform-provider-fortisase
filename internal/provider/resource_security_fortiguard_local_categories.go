// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
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
var _ resource.Resource = &resourceSecurityFortiguardLocalCategories{}

func newResourceSecurityFortiguardLocalCategories() resource.Resource {
	return &resourceSecurityFortiguardLocalCategories{}
}

type resourceSecurityFortiguardLocalCategories struct {
	fortiClient *FortiClient
}

// resourceSecurityFortiguardLocalCategoriesModel describes the resource data model.
type resourceSecurityFortiguardLocalCategoriesModel struct {
	ID           types.String `tfsdk:"id"`
	PrimaryKey   types.String `tfsdk:"primary_key"`
	ThreatWeight types.String `tfsdk:"threat_weight"`
	Urls         types.Set    `tfsdk:"urls"`
}

func (r *resourceSecurityFortiguardLocalCategories) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_fortiguard_local_categories"
}

func (r *resourceSecurityFortiguardLocalCategories) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthBetween(1, 79),
				},
				Required: true,
			},
			"threat_weight": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("none", "low", "medium", "high", "critical"),
				},
				Computed: true,
				Optional: true,
			},
			"urls": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *resourceSecurityFortiguardLocalCategories) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityFortiguardLocalCategories) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityFortiguardLocalCategoriesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityFortiguardLocalCategories(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategories(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityFortiguardLocalCategories(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategories(ctx, "read", diags))

	read_output, err := c.ReadSecurityFortiguardLocalCategories(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityFortiguardLocalCategories(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityFortiguardLocalCategories) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityFortiguardLocalCategoriesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityFortiguardLocalCategoriesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityFortiguardLocalCategories(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategories(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityFortiguardLocalCategories(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategories(ctx, "read", diags))

	read_output, err := c.ReadSecurityFortiguardLocalCategories(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityFortiguardLocalCategories(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityFortiguardLocalCategories) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityFortiguardLocalCategoriesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategories(ctx, "delete", diags))

	err := c.DeleteSecurityFortiguardLocalCategories(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceSecurityFortiguardLocalCategories) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityFortiguardLocalCategoriesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategories(ctx, "read", diags))

	read_output, err := c.ReadSecurityFortiguardLocalCategories(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityFortiguardLocalCategories(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityFortiguardLocalCategories) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityFortiguardLocalCategoriesModel) refreshSecurityFortiguardLocalCategories(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["threatWeight"]; ok {
		m.ThreatWeight = parseStringValue(v)
	}

	if v, ok := o["urls"]; ok {
		m.Urls = parseSetValue(ctx, v, types.StringType)
	}

	return diags
}

func (data *resourceSecurityFortiguardLocalCategoriesModel) getCreateObjectSecurityFortiguardLocalCategories(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ThreatWeight.IsNull() {
		result["threatWeight"] = data.ThreatWeight.ValueString()
	}

	if !data.Urls.IsNull() {
		result["urls"] = expandSetToStringList(data.Urls)
	}

	return &result
}

func (data *resourceSecurityFortiguardLocalCategoriesModel) getUpdateObjectSecurityFortiguardLocalCategories(ctx context.Context, state resourceSecurityFortiguardLocalCategoriesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ThreatWeight.IsNull() {
		result["threatWeight"] = data.ThreatWeight.ValueString()
	}

	if !data.Urls.IsNull() {
		result["urls"] = expandSetToStringList(data.Urls)
	}

	return &result
}

func (data *resourceSecurityFortiguardLocalCategoriesModel) getURLObjectSecurityFortiguardLocalCategories(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
