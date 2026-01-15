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
var _ resource.Resource = &resourceEndpointFssoProfiles{}

func newResourceEndpointFssoProfiles() resource.Resource {
	return &resourceEndpointFssoProfiles{}
}

type resourceEndpointFssoProfiles struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointFssoProfilesModel describes the resource data model.
type resourceEndpointFssoProfilesModel struct {
	ID            types.String  `tfsdk:"id"`
	Enabled       types.Bool    `tfsdk:"enabled"`
	PreferEntraId types.String  `tfsdk:"prefer_entra_id"`
	Host          types.String  `tfsdk:"host"`
	Port          types.Float64 `tfsdk:"port"`
	PreSharedKey  types.String  `tfsdk:"pre_shared_key"`
	PrimaryKey    types.String  `tfsdk:"primary_key"`
}

func (r *resourceEndpointFssoProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_fsso_profiles"
}

func (r *resourceEndpointFssoProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"prefer_entra_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"host": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"port": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(65535),
				},
				Computed: true,
				Optional: true,
			},
			"pre_shared_key": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *resourceEndpointFssoProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_fsso_profiles"
}

func (r *resourceEndpointFssoProfiles) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointFssoProfiles")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointFssoProfilesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointFssoProfiles(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointFssoProfiles(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateEndpointFssoProfiles(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectEndpointFssoProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointFssoProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointFssoProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointFssoProfiles) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointFssoProfiles")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointFssoProfilesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointFssoProfilesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointFssoProfiles(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointFssoProfiles(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointFssoProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointFssoProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointFssoProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointFssoProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointFssoProfiles) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointFssoProfiles) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointFssoProfilesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointFssoProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointFssoProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointFssoProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointFssoProfiles) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointFssoProfilesModel) refreshEndpointFssoProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["preferEntraId"]; ok {
		m.PreferEntraId = parseStringValue(v)
	}

	if v, ok := o["host"]; ok {
		m.Host = parseStringValue(v)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	if v, ok := o["preSharedKey"]; ok {
		m.PreSharedKey = parseStringValue(v)
	}

	return diags
}

func (data *resourceEndpointFssoProfilesModel) getCreateObjectEndpointFssoProfiles(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.PreferEntraId.IsNull() {
		result["preferEntraId"] = data.PreferEntraId.ValueString()
	}

	if !data.Host.IsNull() {
		result["host"] = data.Host.ValueString()
	}

	if !data.Port.IsNull() {
		result["port"] = data.Port.ValueFloat64()
	}

	if !data.PreSharedKey.IsNull() {
		result["preSharedKey"] = data.PreSharedKey.ValueString()
	}

	return &result
}

func (data *resourceEndpointFssoProfilesModel) getUpdateObjectEndpointFssoProfiles(ctx context.Context, state resourceEndpointFssoProfilesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.PreferEntraId.IsNull() {
		result["preferEntraId"] = data.PreferEntraId.ValueString()
	}

	if !data.Host.IsNull() {
		result["host"] = data.Host.ValueString()
	}

	if !data.Port.IsNull() {
		result["port"] = data.Port.ValueFloat64()
	}

	if !data.PreSharedKey.IsNull() {
		result["preSharedKey"] = data.PreSharedKey.ValueString()
	}

	return &result
}

func (data *resourceEndpointFssoProfilesModel) getURLObjectEndpointFssoProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
