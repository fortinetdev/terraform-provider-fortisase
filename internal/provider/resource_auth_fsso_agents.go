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
var _ resource.Resource = &resourceAuthFssoAgents{}

func newResourceAuthFssoAgents() resource.Resource {
	return &resourceAuthFssoAgents{}
}

type resourceAuthFssoAgents struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceAuthFssoAgentsModel describes the resource data model.
type resourceAuthFssoAgentsModel struct {
	ID             types.String `tfsdk:"id"`
	PrimaryKey     types.String `tfsdk:"primary_key"`
	ActiveServer   types.String `tfsdk:"active_server"`
	Status         types.String `tfsdk:"status"`
	Name           types.String `tfsdk:"name"`
	Server         types.String `tfsdk:"server"`
	Password       types.String `tfsdk:"password"`
	Server2        types.String `tfsdk:"server2"`
	Password2      types.String `tfsdk:"password2"`
	Server3        types.String `tfsdk:"server3"`
	Password3      types.String `tfsdk:"password3"`
	Server4        types.String `tfsdk:"server4"`
	Password4      types.String `tfsdk:"password4"`
	Server5        types.String `tfsdk:"server5"`
	Password5      types.String `tfsdk:"password5"`
	SslTrustedCert types.String `tfsdk:"ssl_trusted_cert"`
}

func (r *resourceAuthFssoAgents) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_fsso_agents"
}

func (r *resourceAuthFssoAgents) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"active_server": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("connected", "disconnected"),
				},
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 35),
				},
				Computed: true,
				Optional: true,
			},
			"server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"password": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"server2": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"password2": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"server3": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"password3": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"server4": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"password4": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"server5": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"password5": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"ssl_trusted_cert": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(79),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceAuthFssoAgents) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_fsso_agents"
}

func (r *resourceAuthFssoAgents) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthFssoAgents")
	lock.Lock()
	defer lock.Unlock()
	var data resourceAuthFssoAgentsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectAuthFssoAgents(ctx, diags))
	input_model.URLParams = *(data.getURLObjectAuthFssoAgents(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateAuthFssoAgents(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectAuthFssoAgents(ctx, "read", diags))

	read_output, err := c.ReadAuthFssoAgents(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthFssoAgents(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthFssoAgents) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthFssoAgents")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceAuthFssoAgentsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceAuthFssoAgentsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectAuthFssoAgents(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectAuthFssoAgents(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateAuthFssoAgents(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectAuthFssoAgents(ctx, "read", diags))

	read_output, err := c.ReadAuthFssoAgents(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthFssoAgents(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthFssoAgents) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("AuthFssoAgents")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceAuthFssoAgentsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthFssoAgents(ctx, "delete", diags))

	output, err := c.DeleteAuthFssoAgents(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceAuthFssoAgents) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthFssoAgentsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthFssoAgents(ctx, "read", diags))

	read_output, err := c.ReadAuthFssoAgents(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthFssoAgents(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthFssoAgents) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceAuthFssoAgentsModel) refreshAuthFssoAgents(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["activeServer"]; ok {
		m.ActiveServer = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["server"]; ok {
		m.Server = parseStringValue(v)
	}

	if v, ok := o["server2"]; ok {
		m.Server2 = parseStringValue(v)
	}

	if v, ok := o["server3"]; ok {
		m.Server3 = parseStringValue(v)
	}

	if v, ok := o["server4"]; ok {
		m.Server4 = parseStringValue(v)
	}

	if v, ok := o["server5"]; ok {
		m.Server5 = parseStringValue(v)
	}

	if v, ok := o["sslTrustedCert"]; ok {
		m.SslTrustedCert = parseStringValue(v)
	}

	return diags
}

func (data *resourceAuthFssoAgentsModel) getCreateObjectAuthFssoAgents(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ActiveServer.IsNull() {
		result["activeServer"] = data.ActiveServer.ValueString()
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Name.IsNull() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Server.IsNull() {
		result["server"] = data.Server.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	if !data.Server2.IsNull() {
		result["server2"] = data.Server2.ValueString()
	}

	if !data.Password2.IsNull() {
		result["password2"] = data.Password2.ValueString()
	}

	if !data.Server3.IsNull() {
		result["server3"] = data.Server3.ValueString()
	}

	if !data.Password3.IsNull() {
		result["password3"] = data.Password3.ValueString()
	}

	if !data.Server4.IsNull() {
		result["server4"] = data.Server4.ValueString()
	}

	if !data.Password4.IsNull() {
		result["password4"] = data.Password4.ValueString()
	}

	if !data.Server5.IsNull() {
		result["server5"] = data.Server5.ValueString()
	}

	if !data.Password5.IsNull() {
		result["password5"] = data.Password5.ValueString()
	}

	if !data.SslTrustedCert.IsNull() {
		result["sslTrustedCert"] = data.SslTrustedCert.ValueString()
	}

	return &result
}

func (data *resourceAuthFssoAgentsModel) getUpdateObjectAuthFssoAgents(ctx context.Context, state resourceAuthFssoAgentsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ActiveServer.IsNull() {
		result["activeServer"] = data.ActiveServer.ValueString()
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Name.IsNull() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Server.IsNull() {
		result["server"] = data.Server.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	if !data.Server2.IsNull() {
		result["server2"] = data.Server2.ValueString()
	}

	if !data.Password2.IsNull() {
		result["password2"] = data.Password2.ValueString()
	}

	if !data.Server3.IsNull() {
		result["server3"] = data.Server3.ValueString()
	}

	if !data.Password3.IsNull() {
		result["password3"] = data.Password3.ValueString()
	}

	if !data.Server4.IsNull() {
		result["server4"] = data.Server4.ValueString()
	}

	if !data.Password4.IsNull() {
		result["password4"] = data.Password4.ValueString()
	}

	if !data.Server5.IsNull() {
		result["server5"] = data.Server5.ValueString()
	}

	if !data.Password5.IsNull() {
		result["password5"] = data.Password5.ValueString()
	}

	if !data.SslTrustedCert.IsNull() {
		result["sslTrustedCert"] = data.SslTrustedCert.ValueString()
	}

	return &result
}

func (data *resourceAuthFssoAgentsModel) getURLObjectAuthFssoAgents(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
