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
var _ resource.Resource = &resourceSecurityVideoFilterYoutubeKey{}

func newResourceSecurityVideoFilterYoutubeKey() resource.Resource {
	return &resourceSecurityVideoFilterYoutubeKey{}
}

type resourceSecurityVideoFilterYoutubeKey struct {
	fortiClient *FortiClient
}

// resourceSecurityVideoFilterYoutubeKeyModel describes the resource data model.
type resourceSecurityVideoFilterYoutubeKeyModel struct {
	ID         types.String `tfsdk:"id"`
	PrimaryKey types.String `tfsdk:"primary_key"`
	ApiKey     types.String `tfsdk:"api_key"`
}

func (r *resourceSecurityVideoFilterYoutubeKey) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_video_filter_youtube_key"
}

func (r *resourceSecurityVideoFilterYoutubeKey) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.OneOf("$sase-global"),
				},
				Required: true,
			},
			"api_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 47),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityVideoFilterYoutubeKey) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityVideoFilterYoutubeKey) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityVideoFilterYoutubeKeyModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecurityVideoFilterYoutubeKey(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecurityVideoFilterYoutubeKey(&input_model)
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

	read_output, err := c.ReadSecurityVideoFilterYoutubeKey(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityVideoFilterYoutubeKey(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityVideoFilterYoutubeKey) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityVideoFilterYoutubeKeyModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityVideoFilterYoutubeKeyModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityVideoFilterYoutubeKey(ctx, state, diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityVideoFilterYoutubeKey(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output, err := c.ReadSecurityVideoFilterYoutubeKey(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityVideoFilterYoutubeKey(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityVideoFilterYoutubeKey) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityVideoFilterYoutubeKey) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityVideoFilterYoutubeKeyModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadSecurityVideoFilterYoutubeKey(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityVideoFilterYoutubeKey(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityVideoFilterYoutubeKey) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityVideoFilterYoutubeKeyModel) refreshSecurityVideoFilterYoutubeKey(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["apiKey"]; ok {
		m.ApiKey = parseStringValue(v)
	}

	return diags
}

func (data *resourceSecurityVideoFilterYoutubeKeyModel) getCreateObjectSecurityVideoFilterYoutubeKey(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ApiKey.IsNull() {
		result["apiKey"] = data.ApiKey.ValueString()
	}

	return &result
}

func (data *resourceSecurityVideoFilterYoutubeKeyModel) getUpdateObjectSecurityVideoFilterYoutubeKey(ctx context.Context, state resourceSecurityVideoFilterYoutubeKeyModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ApiKey.IsNull() {
		result["apiKey"] = data.ApiKey.ValueString()
	}

	return &result
}
