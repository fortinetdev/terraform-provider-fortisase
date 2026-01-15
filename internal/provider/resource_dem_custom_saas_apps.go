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
var _ resource.Resource = &resourceDemCustomSaasApps{}

func newResourceDemCustomSaasApps() resource.Resource {
	return &resourceDemCustomSaasApps{}
}

type resourceDemCustomSaasApps struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceDemCustomSaasAppsModel describes the resource data model.
type resourceDemCustomSaasAppsModel struct {
	ID         types.String `tfsdk:"id"`
	PrimaryKey types.String `tfsdk:"primary_key"`
	Alias      types.String `tfsdk:"alias"`
	Fqdn       types.String `tfsdk:"fqdn"`
}

func (r *resourceDemCustomSaasApps) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dem_custom_saas_apps"
}

func (r *resourceDemCustomSaasApps) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthBetween(1, 253),
				},
				MarkdownDescription: "The primary key object of the DEM custom SaaS application. Can not be updated once created.\nLength between 1 and 253.",
				Required:            true,
			},
			"alias": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"fqdn": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 253),
				},
				MarkdownDescription: "The FQDN of the custom SaaS application.\nLength between 1 and 253.",
				Computed:            true,
				Optional:            true,
			},
		},
	}
}

func (r *resourceDemCustomSaasApps) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_dem_custom_saas_apps"
}

func (r *resourceDemCustomSaasApps) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("DemCustomSaasApps")
	lock.Lock()
	defer lock.Unlock()
	var data resourceDemCustomSaasAppsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectDemCustomSaasApps(ctx, diags))
	input_model.URLParams = *(data.getURLObjectDemCustomSaasApps(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateDemCustomSaasApps(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectDemCustomSaasApps(ctx, "read", diags))

	read_output, err := c.ReadDemCustomSaasApps(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshDemCustomSaasApps(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceDemCustomSaasApps) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("DemCustomSaasApps")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceDemCustomSaasAppsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceDemCustomSaasAppsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectDemCustomSaasApps(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectDemCustomSaasApps(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateDemCustomSaasApps(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectDemCustomSaasApps(ctx, "read", diags))

	read_output, err := c.ReadDemCustomSaasApps(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshDemCustomSaasApps(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceDemCustomSaasApps) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("DemCustomSaasApps")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceDemCustomSaasAppsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectDemCustomSaasApps(ctx, "delete", diags))

	output, err := c.DeleteDemCustomSaasApps(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceDemCustomSaasApps) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceDemCustomSaasAppsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectDemCustomSaasApps(ctx, "read", diags))

	read_output, err := c.ReadDemCustomSaasApps(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshDemCustomSaasApps(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceDemCustomSaasApps) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceDemCustomSaasAppsModel) refreshDemCustomSaasApps(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["alias"]; ok {
		m.Alias = parseStringValue(v)
	}

	if v, ok := o["fqdn"]; ok {
		m.Fqdn = parseStringValue(v)
	}

	return diags
}

func (data *resourceDemCustomSaasAppsModel) getCreateObjectDemCustomSaasApps(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Alias.IsNull() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.Fqdn.IsNull() {
		result["fqdn"] = data.Fqdn.ValueString()
	}

	return &result
}

func (data *resourceDemCustomSaasAppsModel) getUpdateObjectDemCustomSaasApps(ctx context.Context, state resourceDemCustomSaasAppsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Alias.IsNull() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.Fqdn.IsNull() {
		result["fqdn"] = data.Fqdn.ValueString()
	}

	return &result
}

func (data *resourceDemCustomSaasAppsModel) getURLObjectDemCustomSaasApps(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
