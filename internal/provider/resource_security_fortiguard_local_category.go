// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
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
var _ resource.Resource = &resourceSecurityFortiguardLocalCategory{}
var _ resource.ResourceWithMoveState = &resourceSecurityFortiguardLocalCategory{}

func newResourceSecurityFortiguardLocalCategory() resource.Resource {
	return &resourceSecurityFortiguardLocalCategory{}
}

type resourceSecurityFortiguardLocalCategory struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityFortiguardLocalCategoryModel describes the resource data model.
type resourceSecurityFortiguardLocalCategoryModel struct {
	ID           types.String `tfsdk:"id"`
	PrimaryKey   types.String `tfsdk:"primary_key"`
	ThreatWeight types.String `tfsdk:"threat_weight"`
	Urls         types.Set    `tfsdk:"urls"`
}

func (r *resourceSecurityFortiguardLocalCategory) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_fortiguard_local_category"
}

func (r *resourceSecurityFortiguardLocalCategory) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "FortiGuard Local Category Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 79),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"threat_weight": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("none", "low", "medium", "high", "critical"),
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

func (r *resourceSecurityFortiguardLocalCategory) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_fortiguard_local_category"
}
func (r *resourceSecurityFortiguardLocalCategory) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_fortiguard_local_categories" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityFortiguardLocalCategoryModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityFortiguardLocalCategory) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityFortiguardLocalCategories")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityFortiguardLocalCategoryModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityFortiguardLocalCategory(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategory(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityFortiguardLocalCategories(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategory(ctx, "read", diags))

	read_output, err := c.ReadSecurityFortiguardLocalCategories(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityFortiguardLocalCategory(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityFortiguardLocalCategory) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityFortiguardLocalCategories")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityFortiguardLocalCategoryModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityFortiguardLocalCategoryModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityFortiguardLocalCategory(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategory(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityFortiguardLocalCategories(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategory(ctx, "read", diags))

	read_output, err := c.ReadSecurityFortiguardLocalCategories(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityFortiguardLocalCategory(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityFortiguardLocalCategory) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityFortiguardLocalCategories")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityFortiguardLocalCategoryModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategory(ctx, "delete", diags))

	output, err := c.DeleteSecurityFortiguardLocalCategories(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityFortiguardLocalCategory) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityFortiguardLocalCategoryModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategory(ctx, "read", diags))

	read_output, err := c.ReadSecurityFortiguardLocalCategories(&input_model)
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

	diags.Append(data.refreshSecurityFortiguardLocalCategory(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityFortiguardLocalCategory) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityFortiguardLocalCategoryModel) refreshSecurityFortiguardLocalCategory(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["threatWeight"]; ok {
		m.ThreatWeight = parseStringValue(v)
	}

	if v, ok := o["urls"]; ok {
		m.Urls = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Urls = types.SetNull(types.StringType)
	}

	return diags
}

func (data *resourceSecurityFortiguardLocalCategoryModel) getCreateObjectSecurityFortiguardLocalCategory(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ThreatWeight.IsNull() && !data.ThreatWeight.IsUnknown() {
		result["threatWeight"] = data.ThreatWeight.ValueString()
	}

	if !data.Urls.IsNull() && !data.Urls.IsUnknown() {
		result["urls"] = expandSetToStringList(data.Urls)
	}

	return &result
}

func (data *resourceSecurityFortiguardLocalCategoryModel) getUpdateObjectSecurityFortiguardLocalCategory(ctx context.Context, state resourceSecurityFortiguardLocalCategoryModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ThreatWeight.IsNull() && !data.ThreatWeight.IsUnknown() {
		result["threatWeight"] = data.ThreatWeight.ValueString()
	}

	if !data.Urls.IsNull() && !data.Urls.IsUnknown() {
		result["urls"] = expandSetToStringList(data.Urls)
	}

	return &result
}

func (data *resourceSecurityFortiguardLocalCategoryModel) getURLObjectSecurityFortiguardLocalCategory(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
