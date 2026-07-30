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
var _ resource.Resource = &resourceSecurityIpsCustomSignature{}
var _ resource.ResourceWithMoveState = &resourceSecurityIpsCustomSignature{}

func newResourceSecurityIpsCustomSignature() resource.Resource {
	return &resourceSecurityIpsCustomSignature{}
}

type resourceSecurityIpsCustomSignature struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityIpsCustomSignatureModel describes the resource data model.
type resourceSecurityIpsCustomSignatureModel struct {
	ID          types.String  `tfsdk:"id"`
	PrimaryKey  types.String  `tfsdk:"primary_key"`
	Tag         types.String  `tfsdk:"tag"`
	Signature   types.String  `tfsdk:"signature"`
	RuleId      types.Float64 `tfsdk:"rule_id"`
	Status      types.String  `tfsdk:"status"`
	Log         types.String  `tfsdk:"log"`
	LogPacket   types.String  `tfsdk:"log_packet"`
	Action      types.String  `tfsdk:"action"`
	Severity    types.String  `tfsdk:"severity"`
	Location    types.String  `tfsdk:"location"`
	Os          types.String  `tfsdk:"os"`
	Application types.String  `tfsdk:"application"`
	Protocol    types.String  `tfsdk:"protocol"`
	Comment     types.String  `tfsdk:"comment"`
}

func (r *resourceSecurityIpsCustomSignature) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ips_custom_signature"
}

func (r *resourceSecurityIpsCustomSignature) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IPS Custom Signature Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 63),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"tag": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"signature": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(4095),
				},
				Computed: true,
				Optional: true,
			},
			"rule_id": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"log": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"log_packet": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"action": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"severity": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"location": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"os": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"application": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"protocol": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"comment": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityIpsCustomSignature) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_ips_custom_signature"
}
func (r *resourceSecurityIpsCustomSignature) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_ips_custom_signatures" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityIpsCustomSignatureModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityIpsCustomSignature) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityIpsCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityIpsCustomSignatureModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityIpsCustomSignature(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignature(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityIpsCustomSignatures(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignature(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsCustomSignatures(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityIpsCustomSignature(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityIpsCustomSignature) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityIpsCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityIpsCustomSignatureModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityIpsCustomSignatureModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityIpsCustomSignature(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignature(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityIpsCustomSignatures(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignature(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsCustomSignatures(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityIpsCustomSignature(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityIpsCustomSignature) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityIpsCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityIpsCustomSignatureModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignature(ctx, "delete", diags))

	output, err := c.DeleteSecurityIpsCustomSignatures(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityIpsCustomSignature) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityIpsCustomSignatureModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignature(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsCustomSignatures(&input_model)
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

	diags.Append(data.refreshSecurityIpsCustomSignature(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityIpsCustomSignature) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityIpsCustomSignatureModel) refreshSecurityIpsCustomSignature(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["tag"]; ok {
		m.Tag = parseStringValue(v)
	}

	if v, ok := o["ruleId"]; ok {
		m.RuleId = parseFloat64Value(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["log"]; ok {
		m.Log = parseStringValue(v)
	}

	if v, ok := o["logPacket"]; ok {
		m.LogPacket = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["severity"]; ok {
		m.Severity = parseStringValue(v)
	}

	if v, ok := o["location"]; ok {
		m.Location = parseStringValue(v)
	}

	if v, ok := o["os"]; ok {
		m.Os = parseStringValue(v)
	}

	if v, ok := o["application"]; ok {
		m.Application = parseStringValue(v)
	}

	if v, ok := o["protocol"]; ok {
		m.Protocol = parseStringValue(v)
	}

	if v, ok := o["comment"]; ok {
		m.Comment = parseStringValue(v)
	}

	return diags
}

func (data *resourceSecurityIpsCustomSignatureModel) getCreateObjectSecurityIpsCustomSignature(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Signature.IsNull() && !data.Signature.IsUnknown() {
		result["signature"] = data.Signature.ValueString()
	}

	if !data.RuleId.IsNull() && !data.RuleId.IsUnknown() {
		result["ruleId"] = data.RuleId.ValueFloat64()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Log.IsNull() && !data.Log.IsUnknown() {
		result["log"] = data.Log.ValueString()
	}

	if !data.LogPacket.IsNull() && !data.LogPacket.IsUnknown() {
		result["logPacket"] = data.LogPacket.ValueString()
	}

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Severity.IsNull() && !data.Severity.IsUnknown() {
		result["severity"] = data.Severity.ValueString()
	}

	if !data.Location.IsNull() && !data.Location.IsUnknown() {
		result["location"] = data.Location.ValueString()
	}

	if !data.Os.IsNull() && !data.Os.IsUnknown() {
		result["os"] = data.Os.ValueString()
	}

	if !data.Application.IsNull() && !data.Application.IsUnknown() {
		result["application"] = data.Application.ValueString()
	}

	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		result["comment"] = data.Comment.ValueString()
	}

	return &result
}

func (data *resourceSecurityIpsCustomSignatureModel) getUpdateObjectSecurityIpsCustomSignature(ctx context.Context, state resourceSecurityIpsCustomSignatureModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Signature.IsNull() && !data.Signature.IsUnknown() {
		result["signature"] = data.Signature.ValueString()
	}

	if !data.RuleId.IsNull() && !data.RuleId.IsUnknown() {
		result["ruleId"] = data.RuleId.ValueFloat64()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Log.IsNull() && !data.Log.IsUnknown() {
		result["log"] = data.Log.ValueString()
	}

	if !data.LogPacket.IsNull() && !data.LogPacket.IsUnknown() {
		result["logPacket"] = data.LogPacket.ValueString()
	}

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Severity.IsNull() && !data.Severity.IsUnknown() {
		result["severity"] = data.Severity.ValueString()
	}

	if !data.Location.IsNull() && !data.Location.IsUnknown() {
		result["location"] = data.Location.ValueString()
	}

	if !data.Os.IsNull() && !data.Os.IsUnknown() {
		result["os"] = data.Os.ValueString()
	}

	if !data.Application.IsNull() && !data.Application.IsUnknown() {
		result["application"] = data.Application.ValueString()
	}

	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		result["comment"] = data.Comment.ValueString()
	}

	return &result
}

func (data *resourceSecurityIpsCustomSignatureModel) getURLObjectSecurityIpsCustomSignature(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
