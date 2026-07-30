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
var _ resource.Resource = &resourceAuthFssoAgent{}
var _ resource.ResourceWithMoveState = &resourceAuthFssoAgent{}

func newResourceAuthFssoAgent() resource.Resource {
	return &resourceAuthFssoAgent{}
}

type resourceAuthFssoAgent struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceAuthFssoAgentModel describes the resource data model.
type resourceAuthFssoAgentModel struct {
	ID             types.String `tfsdk:"id"`
	PrimaryKey     types.String `tfsdk:"primary_key"`
	ActiveServer   types.String `tfsdk:"active_server"`
	Status         types.String `tfsdk:"status"`
	Name           types.String `tfsdk:"name"`
	Server         types.String `tfsdk:"server"`
	Port           types.String `tfsdk:"port"`
	Password       types.String `tfsdk:"password"`
	Server2        types.String `tfsdk:"server2"`
	Port2          types.String `tfsdk:"port2"`
	Password2      types.String `tfsdk:"password2"`
	Server3        types.String `tfsdk:"server3"`
	Port3          types.String `tfsdk:"port3"`
	Password3      types.String `tfsdk:"password3"`
	Server4        types.String `tfsdk:"server4"`
	Port4          types.String `tfsdk:"port4"`
	Password4      types.String `tfsdk:"password4"`
	Server5        types.String `tfsdk:"server5"`
	Port5          types.String `tfsdk:"port5"`
	Password5      types.String `tfsdk:"password5"`
	SslTrustedCert types.String `tfsdk:"ssl_trusted_cert"`
}

func (r *resourceAuthFssoAgent) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_fsso_agent"
}

func (r *resourceAuthFssoAgent) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "FSSO Agent Resource API V2 for FortiSASE.",
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
			"active_server": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("connected", "disconnected"),
				},
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 35),
				},
				Computed: true,
				Optional: true,
			},
			"server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"port": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(5),
				},
				Computed: true,
				Optional: true,
			},
			"password": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"server2": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"port2": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(5),
				},
				Computed: true,
				Optional: true,
			},
			"password2": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"server3": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"port3": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(5),
				},
				Computed: true,
				Optional: true,
			},
			"password3": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"server4": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"port4": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(5),
				},
				Computed: true,
				Optional: true,
			},
			"password4": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"server5": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"port5": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(5),
				},
				Computed: true,
				Optional: true,
			},
			"password5": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"ssl_trusted_cert": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(79),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceAuthFssoAgent) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_fsso_agent"
}
func (r *resourceAuthFssoAgent) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_auth_fsso_agents" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceAuthFssoAgentModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceAuthFssoAgent) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthFssoAgents")
	lock.Lock()
	defer lock.Unlock()
	var data resourceAuthFssoAgentModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectAuthFssoAgent(ctx, diags))
	input_model.URLParams = *(data.getURLObjectAuthFssoAgent(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateAuthFssoAgents(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectAuthFssoAgent(ctx, "read", diags))

	read_output, err := c.ReadAuthFssoAgents(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthFssoAgent(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthFssoAgent) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthFssoAgents")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceAuthFssoAgentModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceAuthFssoAgentModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectAuthFssoAgent(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectAuthFssoAgent(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectAuthFssoAgent(ctx, "read", diags))

	read_output, err := c.ReadAuthFssoAgents(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthFssoAgent(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthFssoAgent) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("AuthFssoAgents")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceAuthFssoAgentModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthFssoAgent(ctx, "delete", diags))

	output, err := c.DeleteAuthFssoAgents(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceAuthFssoAgent) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthFssoAgentModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthFssoAgent(ctx, "read", diags))

	read_output, err := c.ReadAuthFssoAgents(&input_model)
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

	diags.Append(data.refreshAuthFssoAgent(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthFssoAgent) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceAuthFssoAgentModel) refreshAuthFssoAgent(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

	if v, ok := o["port"]; ok {
		m.Port = parseStringValue(v)
	}

	if v, ok := o["server2"]; ok {
		m.Server2 = parseStringValue(v)
	}

	if v, ok := o["port2"]; ok {
		m.Port2 = parseStringValue(v)
	}

	if v, ok := o["server3"]; ok {
		m.Server3 = parseStringValue(v)
	}

	if v, ok := o["port3"]; ok {
		m.Port3 = parseStringValue(v)
	}

	if v, ok := o["server4"]; ok {
		m.Server4 = parseStringValue(v)
	}

	if v, ok := o["port4"]; ok {
		m.Port4 = parseStringValue(v)
	}

	if v, ok := o["server5"]; ok {
		m.Server5 = parseStringValue(v)
	}

	if v, ok := o["port5"]; ok {
		m.Port5 = parseStringValue(v)
	}

	if v, ok := o["sslTrustedCert"]; ok {
		m.SslTrustedCert = parseStringValue(v)
	}

	return diags
}

func (data *resourceAuthFssoAgentModel) getCreateObjectAuthFssoAgent(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ActiveServer.IsNull() && !data.ActiveServer.IsUnknown() {
		result["activeServer"] = data.ActiveServer.ValueString()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Server.IsNull() && !data.Server.IsUnknown() {
		result["server"] = data.Server.ValueString()
	}

	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		result["port"] = data.Port.ValueString()
	}

	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		result["password"] = data.Password.ValueString()
	}

	if !data.Server2.IsNull() && !data.Server2.IsUnknown() {
		result["server2"] = data.Server2.ValueString()
	}

	if !data.Port2.IsNull() && !data.Port2.IsUnknown() {
		result["port2"] = data.Port2.ValueString()
	}

	if !data.Password2.IsNull() && !data.Password2.IsUnknown() {
		result["password2"] = data.Password2.ValueString()
	}

	if !data.Server3.IsNull() && !data.Server3.IsUnknown() {
		result["server3"] = data.Server3.ValueString()
	}

	if !data.Port3.IsNull() && !data.Port3.IsUnknown() {
		result["port3"] = data.Port3.ValueString()
	}

	if !data.Password3.IsNull() && !data.Password3.IsUnknown() {
		result["password3"] = data.Password3.ValueString()
	}

	if !data.Server4.IsNull() && !data.Server4.IsUnknown() {
		result["server4"] = data.Server4.ValueString()
	}

	if !data.Port4.IsNull() && !data.Port4.IsUnknown() {
		result["port4"] = data.Port4.ValueString()
	}

	if !data.Password4.IsNull() && !data.Password4.IsUnknown() {
		result["password4"] = data.Password4.ValueString()
	}

	if !data.Server5.IsNull() && !data.Server5.IsUnknown() {
		result["server5"] = data.Server5.ValueString()
	}

	if !data.Port5.IsNull() && !data.Port5.IsUnknown() {
		result["port5"] = data.Port5.ValueString()
	}

	if !data.Password5.IsNull() && !data.Password5.IsUnknown() {
		result["password5"] = data.Password5.ValueString()
	}

	if !data.SslTrustedCert.IsNull() && !data.SslTrustedCert.IsUnknown() {
		result["sslTrustedCert"] = data.SslTrustedCert.ValueString()
	}

	return &result
}

func (data *resourceAuthFssoAgentModel) getUpdateObjectAuthFssoAgent(ctx context.Context, state resourceAuthFssoAgentModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ActiveServer.IsNull() && !data.ActiveServer.IsUnknown() {
		result["activeServer"] = data.ActiveServer.ValueString()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Server.IsNull() && !data.Server.IsUnknown() {
		result["server"] = data.Server.ValueString()
	}

	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		result["port"] = data.Port.ValueString()
	}

	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		result["password"] = data.Password.ValueString()
	}

	if !data.Server2.IsNull() && !data.Server2.IsUnknown() {
		result["server2"] = data.Server2.ValueString()
	}

	if !data.Port2.IsNull() && !data.Port2.IsUnknown() {
		result["port2"] = data.Port2.ValueString()
	}

	if !data.Password2.IsNull() && !data.Password2.IsUnknown() {
		result["password2"] = data.Password2.ValueString()
	}

	if !data.Server3.IsNull() && !data.Server3.IsUnknown() {
		result["server3"] = data.Server3.ValueString()
	}

	if !data.Port3.IsNull() && !data.Port3.IsUnknown() {
		result["port3"] = data.Port3.ValueString()
	}

	if !data.Password3.IsNull() && !data.Password3.IsUnknown() {
		result["password3"] = data.Password3.ValueString()
	}

	if !data.Server4.IsNull() && !data.Server4.IsUnknown() {
		result["server4"] = data.Server4.ValueString()
	}

	if !data.Port4.IsNull() && !data.Port4.IsUnknown() {
		result["port4"] = data.Port4.ValueString()
	}

	if !data.Password4.IsNull() && !data.Password4.IsUnknown() {
		result["password4"] = data.Password4.ValueString()
	}

	if !data.Server5.IsNull() && !data.Server5.IsUnknown() {
		result["server5"] = data.Server5.ValueString()
	}

	if !data.Port5.IsNull() && !data.Port5.IsUnknown() {
		result["port5"] = data.Port5.ValueString()
	}

	if !data.Password5.IsNull() && !data.Password5.IsUnknown() {
		result["password5"] = data.Password5.ValueString()
	}

	if !data.SslTrustedCert.IsNull() && !data.SslTrustedCert.IsUnknown() {
		result["sslTrustedCert"] = data.SslTrustedCert.ValueString()
	}

	return &result
}

func (data *resourceAuthFssoAgentModel) getURLObjectAuthFssoAgent(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
