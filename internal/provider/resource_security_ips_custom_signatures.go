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
var _ resource.Resource = &resourceSecurityIpsCustomSignatures{}

func newResourceSecurityIpsCustomSignatures() resource.Resource {
	return &resourceSecurityIpsCustomSignatures{}
}

type resourceSecurityIpsCustomSignatures struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityIpsCustomSignaturesModel describes the resource data model.
type resourceSecurityIpsCustomSignaturesModel struct {
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

func (r *resourceSecurityIpsCustomSignatures) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ips_custom_signatures"
}

func (r *resourceSecurityIpsCustomSignatures) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthBetween(1, 63),
				},
				Required: true,
			},
			"tag": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"signature": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(4095),
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
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityIpsCustomSignatures) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_ips_custom_signatures"
}

func (r *resourceSecurityIpsCustomSignatures) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityIpsCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityIpsCustomSignaturesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityIpsCustomSignatures(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignatures(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityIpsCustomSignatures(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignatures(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsCustomSignatures(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityIpsCustomSignatures(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityIpsCustomSignatures) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityIpsCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityIpsCustomSignaturesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityIpsCustomSignaturesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityIpsCustomSignatures(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignatures(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignatures(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsCustomSignatures(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityIpsCustomSignatures(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityIpsCustomSignatures) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityIpsCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityIpsCustomSignaturesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignatures(ctx, "delete", diags))

	output, err := c.DeleteSecurityIpsCustomSignatures(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityIpsCustomSignatures) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityIpsCustomSignaturesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignatures(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsCustomSignatures(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityIpsCustomSignatures(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityIpsCustomSignatures) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityIpsCustomSignaturesModel) refreshSecurityIpsCustomSignatures(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceSecurityIpsCustomSignaturesModel) getCreateObjectSecurityIpsCustomSignatures(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Tag.IsNull() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Signature.IsNull() {
		result["signature"] = data.Signature.ValueString()
	}

	if !data.RuleId.IsNull() {
		result["ruleId"] = data.RuleId.ValueFloat64()
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Log.IsNull() {
		result["log"] = data.Log.ValueString()
	}

	if !data.LogPacket.IsNull() {
		result["logPacket"] = data.LogPacket.ValueString()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Severity.IsNull() {
		result["severity"] = data.Severity.ValueString()
	}

	if !data.Location.IsNull() {
		result["location"] = data.Location.ValueString()
	}

	if !data.Os.IsNull() {
		result["os"] = data.Os.ValueString()
	}

	if !data.Application.IsNull() {
		result["application"] = data.Application.ValueString()
	}

	if !data.Protocol.IsNull() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.Comment.IsNull() {
		result["comment"] = data.Comment.ValueString()
	}

	return &result
}

func (data *resourceSecurityIpsCustomSignaturesModel) getUpdateObjectSecurityIpsCustomSignatures(ctx context.Context, state resourceSecurityIpsCustomSignaturesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Tag.IsNull() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Signature.IsNull() {
		result["signature"] = data.Signature.ValueString()
	}

	if !data.RuleId.IsNull() {
		result["ruleId"] = data.RuleId.ValueFloat64()
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Log.IsNull() {
		result["log"] = data.Log.ValueString()
	}

	if !data.LogPacket.IsNull() {
		result["logPacket"] = data.LogPacket.ValueString()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Severity.IsNull() {
		result["severity"] = data.Severity.ValueString()
	}

	if !data.Location.IsNull() {
		result["location"] = data.Location.ValueString()
	}

	if !data.Os.IsNull() {
		result["os"] = data.Os.ValueString()
	}

	if !data.Application.IsNull() {
		result["application"] = data.Application.ValueString()
	}

	if !data.Protocol.IsNull() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.Comment.IsNull() {
		result["comment"] = data.Comment.ValueString()
	}

	return &result
}

func (data *resourceSecurityIpsCustomSignaturesModel) getURLObjectSecurityIpsCustomSignatures(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
