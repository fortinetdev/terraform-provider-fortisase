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
var _ resource.Resource = &resourceAuthRadiusServer{}
var _ resource.ResourceWithMoveState = &resourceAuthRadiusServer{}

func newResourceAuthRadiusServer() resource.Resource {
	return &resourceAuthRadiusServer{}
}

type resourceAuthRadiusServer struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceAuthRadiusServerModel describes the resource data model.
type resourceAuthRadiusServerModel struct {
	ID                         types.String `tfsdk:"id"`
	PrimaryKey                 types.String `tfsdk:"primary_key"`
	AuthType                   types.String `tfsdk:"auth_type"`
	PrimaryServer              types.String `tfsdk:"primary_server"`
	PrimarySecret              types.String `tfsdk:"primary_secret"`
	IncludedInDefaultUserGroup types.Bool   `tfsdk:"included_in_default_user_group"`
	SecondaryServer            types.String `tfsdk:"secondary_server"`
	SecondarySecret            types.String `tfsdk:"secondary_secret"`
}

func (r *resourceAuthRadiusServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_radius_server"
}

func (r *resourceAuthRadiusServer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "RADIUS Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 35),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"auth_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("auto", "pap", "chap", "ms_chap", "ms_chap_v2"),
				},
				Computed: true,
				Optional: true,
			},
			"primary_server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"primary_secret": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
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
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"secondary_secret": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
		},
	}
}

func (r *resourceAuthRadiusServer) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_radius_server"
}
func (r *resourceAuthRadiusServer) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_auth_radius_servers" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceAuthRadiusServerModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceAuthRadiusServer) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthRadiusServers")
	lock.Lock()
	defer lock.Unlock()
	var data resourceAuthRadiusServerModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectAuthRadiusServer(ctx, diags))
	input_model.URLParams = *(data.getURLObjectAuthRadiusServer(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateAuthRadiusServers(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectAuthRadiusServer(ctx, "read", diags))

	read_output, err := c.ReadAuthRadiusServers(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthRadiusServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthRadiusServer) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthRadiusServers")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceAuthRadiusServerModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceAuthRadiusServerModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectAuthRadiusServer(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectAuthRadiusServer(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectAuthRadiusServer(ctx, "read", diags))

	read_output, err := c.ReadAuthRadiusServers(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthRadiusServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthRadiusServer) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("AuthRadiusServers")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceAuthRadiusServerModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthRadiusServer(ctx, "delete", diags))

	output, err := c.DeleteAuthRadiusServers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceAuthRadiusServer) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthRadiusServerModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthRadiusServer(ctx, "read", diags))

	read_output, err := c.ReadAuthRadiusServers(&input_model)
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

	diags.Append(data.refreshAuthRadiusServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthRadiusServer) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceAuthRadiusServerModel) refreshAuthRadiusServer(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceAuthRadiusServerModel) getCreateObjectAuthRadiusServer(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.AuthType.IsNull() && !data.AuthType.IsUnknown() {
		result["authType"] = data.AuthType.ValueString()
	}

	if !data.PrimaryServer.IsNull() && !data.PrimaryServer.IsUnknown() {
		result["primaryServer"] = data.PrimaryServer.ValueString()
	}

	if !data.PrimarySecret.IsNull() && !data.PrimarySecret.IsUnknown() {
		result["primarySecret"] = data.PrimarySecret.ValueString()
	}

	if !data.IncludedInDefaultUserGroup.IsNull() && !data.IncludedInDefaultUserGroup.IsUnknown() {
		result["includedInDefaultUserGroup"] = data.IncludedInDefaultUserGroup.ValueBool()
	}

	if !data.SecondaryServer.IsNull() && !data.SecondaryServer.IsUnknown() {
		result["secondaryServer"] = data.SecondaryServer.ValueString()
	}

	if !data.SecondarySecret.IsNull() && !data.SecondarySecret.IsUnknown() {
		result["secondarySecret"] = data.SecondarySecret.ValueString()
	}

	return &result
}

func (data *resourceAuthRadiusServerModel) getUpdateObjectAuthRadiusServer(ctx context.Context, state resourceAuthRadiusServerModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.AuthType.IsNull() && !data.AuthType.IsUnknown() {
		result["authType"] = data.AuthType.ValueString()
	}

	if !data.PrimaryServer.IsNull() && !data.PrimaryServer.IsUnknown() {
		result["primaryServer"] = data.PrimaryServer.ValueString()
	}

	if !data.PrimarySecret.IsNull() && !data.PrimarySecret.IsUnknown() {
		result["primarySecret"] = data.PrimarySecret.ValueString()
	}

	if !data.IncludedInDefaultUserGroup.IsNull() && !data.IncludedInDefaultUserGroup.IsUnknown() {
		result["includedInDefaultUserGroup"] = data.IncludedInDefaultUserGroup.ValueBool()
	}

	if !data.SecondaryServer.IsNull() && !data.SecondaryServer.IsUnknown() {
		result["secondaryServer"] = data.SecondaryServer.ValueString()
	}

	if !data.SecondarySecret.IsNull() && !data.SecondarySecret.IsUnknown() {
		result["secondarySecret"] = data.SecondarySecret.ValueString()
	}

	return &result
}

func (data *resourceAuthRadiusServerModel) getURLObjectAuthRadiusServer(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
