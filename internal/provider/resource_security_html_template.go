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
var _ resource.Resource = &resourceSecurityHtmlTemplate{}

func newResourceSecurityHtmlTemplate() resource.Resource {
	return &resourceSecurityHtmlTemplate{}
}

type resourceSecurityHtmlTemplate struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityHtmlTemplateModel describes the resource data model.
type resourceSecurityHtmlTemplateModel struct {
	ID            types.String `tfsdk:"id"`
	PrimaryKey    types.String `tfsdk:"primary_key"`
	Buffer        types.String `tfsdk:"buffer"`
	DefaultBuffer types.String `tfsdk:"default_buffer"`
	Subject       types.String `tfsdk:"subject"`
	Description   types.String `tfsdk:"description"`
}

func (r *resourceSecurityHtmlTemplate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_html_template"
}

func (r *resourceSecurityHtmlTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "HTML Templates Resource API V2 for FortiSASE.",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"buffer": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"default_buffer": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"subject": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityHtmlTemplate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_html_template"
}

func (r *resourceSecurityHtmlTemplate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityHtmlTemplate")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityHtmlTemplateModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecurityHtmlTemplate(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityHtmlTemplate(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecurityHtmlTemplate(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := data.PrimaryKey.ValueString()
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityHtmlTemplate(ctx, "read", diags))

	read_output, err := c.ReadSecurityHtmlTemplate(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityHtmlTemplate(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityHtmlTemplate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityHtmlTemplate")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityHtmlTemplateModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityHtmlTemplateModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityHtmlTemplate(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityHtmlTemplate(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityHtmlTemplate(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityHtmlTemplate(ctx, "read", diags))

	read_output, err := c.ReadSecurityHtmlTemplate(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityHtmlTemplate(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityHtmlTemplate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityHtmlTemplate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityHtmlTemplateModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityHtmlTemplate(ctx, "read", diags))

	read_output, err := c.ReadSecurityHtmlTemplate(&input_model)
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

	diags.Append(data.refreshSecurityHtmlTemplate(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityHtmlTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityHtmlTemplateModel) refreshSecurityHtmlTemplate(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["buffer"]; ok {
		m.Buffer = parseStringValue(v)
	}

	if v, ok := o["defaultBuffer"]; ok {
		m.DefaultBuffer = parseStringValue(v)
	}

	if v, ok := o["subject"]; ok {
		m.Subject = parseStringValue(v)
	}

	if v, ok := o["description"]; ok {
		m.Description = parseStringValue(v)
	}

	return diags
}

func (data *resourceSecurityHtmlTemplateModel) getCreateObjectSecurityHtmlTemplate(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Buffer.IsNull() && !data.Buffer.IsUnknown() {
		result["buffer"] = data.Buffer.ValueString()
	}

	if !data.DefaultBuffer.IsNull() && !data.DefaultBuffer.IsUnknown() {
		result["defaultBuffer"] = data.DefaultBuffer.ValueString()
	}

	if !data.Subject.IsNull() && !data.Subject.IsUnknown() {
		result["subject"] = data.Subject.ValueString()
	}

	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		result["description"] = data.Description.ValueString()
	}

	return &result
}

func (data *resourceSecurityHtmlTemplateModel) getUpdateObjectSecurityHtmlTemplate(ctx context.Context, state resourceSecurityHtmlTemplateModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Buffer.IsNull() && !data.Buffer.IsUnknown() {
		result["buffer"] = data.Buffer.ValueString()
	}

	if !data.DefaultBuffer.IsNull() && !data.DefaultBuffer.IsUnknown() {
		result["defaultBuffer"] = data.DefaultBuffer.ValueString()
	}

	if !data.Subject.IsNull() && !data.Subject.IsUnknown() {
		result["subject"] = data.Subject.ValueString()
	}

	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		result["description"] = data.Description.ValueString()
	}

	return &result
}

func (data *resourceSecurityHtmlTemplateModel) getURLObjectSecurityHtmlTemplate(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
