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
var _ resource.Resource = &resourceAuthRadiusServers{}

func newResourceAuthRadiusServers() resource.Resource {
	return &resourceAuthRadiusServers{}
}

type resourceAuthRadiusServers struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceAuthRadiusServersModel describes the resource data model.
type resourceAuthRadiusServersModel struct {
	ID                         types.String `tfsdk:"id"`
	PrimaryKey                 types.String `tfsdk:"primary_key"`
	AuthType                   types.String `tfsdk:"auth_type"`
	PrimaryServer              types.String `tfsdk:"primary_server"`
	PrimarySecret              types.String `tfsdk:"primary_secret"`
	IncludedInDefaultUserGroup types.Bool   `tfsdk:"included_in_default_user_group"`
	SecondaryServer            types.String `tfsdk:"secondary_server"`
	SecondarySecret            types.String `tfsdk:"secondary_secret"`
}

func (r *resourceAuthRadiusServers) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_radius_servers"
}

func (r *resourceAuthRadiusServers) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Required: true,
			},
			"auth_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("auto", "pap", "chap", "ms_chap", "ms_chap_v2"),
				},
				Computed: true,
				Optional: true,
			},
			"primary_server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"primary_secret": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"included_in_default_user_group": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"secondary_server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"secondary_secret": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
		},
	}
}

func (r *resourceAuthRadiusServers) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_radius_servers"
}

func (r *resourceAuthRadiusServers) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthRadiusServers")
	lock.Lock()
	defer lock.Unlock()
	var data resourceAuthRadiusServersModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectAuthRadiusServers(ctx, diags))
	input_model.URLParams = *(data.getURLObjectAuthRadiusServers(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateAuthRadiusServers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectAuthRadiusServers(ctx, "read", diags))

	read_output, err := c.ReadAuthRadiusServers(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthRadiusServers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthRadiusServers) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthRadiusServers")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceAuthRadiusServersModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceAuthRadiusServersModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectAuthRadiusServers(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectAuthRadiusServers(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateAuthRadiusServers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectAuthRadiusServers(ctx, "read", diags))

	read_output, err := c.ReadAuthRadiusServers(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthRadiusServers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthRadiusServers) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("AuthRadiusServers")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceAuthRadiusServersModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthRadiusServers(ctx, "delete", diags))

	output, err := c.DeleteAuthRadiusServers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceAuthRadiusServers) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthRadiusServersModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthRadiusServers(ctx, "read", diags))

	read_output, err := c.ReadAuthRadiusServers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthRadiusServers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthRadiusServers) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceAuthRadiusServersModel) refreshAuthRadiusServers(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["authType"]; ok {
		m.AuthType = parseStringValue(v)
	}

	if v, ok := o["primaryServer"]; ok {
		m.PrimaryServer = parseStringValue(v)
	}

	if v, ok := o["includedInDefaultUserGroup"]; ok {
		m.IncludedInDefaultUserGroup = parseBoolValue(v)
	}

	if v, ok := o["secondaryServer"]; ok {
		m.SecondaryServer = parseStringValue(v)
	}

	return diags
}

func (data *resourceAuthRadiusServersModel) getCreateObjectAuthRadiusServers(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.AuthType.IsNull() {
		result["authType"] = data.AuthType.ValueString()
	}

	if !data.PrimaryServer.IsNull() {
		result["primaryServer"] = data.PrimaryServer.ValueString()
	}

	if !data.PrimarySecret.IsNull() {
		result["primarySecret"] = data.PrimarySecret.ValueString()
	}

	if !data.IncludedInDefaultUserGroup.IsNull() {
		result["includedInDefaultUserGroup"] = data.IncludedInDefaultUserGroup.ValueBool()
	}

	if !data.SecondaryServer.IsNull() {
		result["secondaryServer"] = data.SecondaryServer.ValueString()
	}

	if !data.SecondarySecret.IsNull() {
		result["secondarySecret"] = data.SecondarySecret.ValueString()
	}

	return &result
}

func (data *resourceAuthRadiusServersModel) getUpdateObjectAuthRadiusServers(ctx context.Context, state resourceAuthRadiusServersModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.AuthType.IsNull() {
		result["authType"] = data.AuthType.ValueString()
	}

	if !data.PrimaryServer.IsNull() {
		result["primaryServer"] = data.PrimaryServer.ValueString()
	}

	if !data.PrimarySecret.IsNull() {
		result["primarySecret"] = data.PrimarySecret.ValueString()
	}

	if !data.IncludedInDefaultUserGroup.IsNull() {
		result["includedInDefaultUserGroup"] = data.IncludedInDefaultUserGroup.ValueBool()
	}

	if !data.SecondaryServer.IsNull() {
		result["secondaryServer"] = data.SecondaryServer.ValueString()
	}

	if !data.SecondarySecret.IsNull() {
		result["secondarySecret"] = data.SecondarySecret.ValueString()
	}

	return &result
}

func (data *resourceAuthRadiusServersModel) getURLObjectAuthRadiusServers(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
