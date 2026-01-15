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
var _ resource.Resource = &resourceSecurityAppCustomSignatures{}

func newResourceSecurityAppCustomSignatures() resource.Resource {
	return &resourceSecurityAppCustomSignatures{}
}

type resourceSecurityAppCustomSignatures struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityAppCustomSignaturesModel describes the resource data model.
type resourceSecurityAppCustomSignaturesModel struct {
	ID         types.String  `tfsdk:"id"`
	PrimaryKey types.String  `tfsdk:"primary_key"`
	Signature  types.String  `tfsdk:"signature"`
	Comment    types.String  `tfsdk:"comment"`
	Ftntid     types.Float64 `tfsdk:"ftntid"`
	Tag        types.String  `tfsdk:"tag"`
	Name       types.String  `tfsdk:"name"`
	Category   types.Float64 `tfsdk:"category"`
	Protocol   types.String  `tfsdk:"protocol"`
	Technology types.String  `tfsdk:"technology"`
	Behavior   types.String  `tfsdk:"behavior"`
	Vendor     types.String  `tfsdk:"vendor"`
	IconClass  types.String  `tfsdk:"icon_class"`
}

func (r *resourceSecurityAppCustomSignatures) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_app_custom_signatures"
}

func (r *resourceSecurityAppCustomSignatures) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"signature": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(4095),
				},
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
			"ftntid": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"tag": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"category": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"protocol": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"technology": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"behavior": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"vendor": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"icon_class": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityAppCustomSignatures) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_app_custom_signatures"
}

func (r *resourceSecurityAppCustomSignatures) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityAppCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityAppCustomSignaturesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityAppCustomSignatures(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignatures(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityAppCustomSignatures(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignatures(ctx, "read", diags))

	read_output, err := c.ReadSecurityAppCustomSignatures(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityAppCustomSignatures(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityAppCustomSignatures) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityAppCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityAppCustomSignaturesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityAppCustomSignaturesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityAppCustomSignatures(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignatures(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityAppCustomSignatures(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignatures(ctx, "read", diags))

	read_output, err := c.ReadSecurityAppCustomSignatures(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityAppCustomSignatures(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityAppCustomSignatures) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityAppCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityAppCustomSignaturesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignatures(ctx, "delete", diags))

	output, err := c.DeleteSecurityAppCustomSignatures(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityAppCustomSignatures) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityAppCustomSignaturesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignatures(ctx, "read", diags))

	read_output, err := c.ReadSecurityAppCustomSignatures(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityAppCustomSignatures(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityAppCustomSignatures) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityAppCustomSignaturesModel) refreshSecurityAppCustomSignatures(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["comment"]; ok {
		m.Comment = parseStringValue(v)
	}

	if v, ok := o["id"]; ok {
		m.Ftntid = parseFloat64Value(v)
	}

	if v, ok := o["tag"]; ok {
		m.Tag = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = parseFloat64Value(v)
	}

	if v, ok := o["protocol"]; ok {
		m.Protocol = parseStringValue(v)
	}

	if v, ok := o["technology"]; ok {
		m.Technology = parseStringValue(v)
	}

	if v, ok := o["behavior"]; ok {
		m.Behavior = parseStringValue(v)
	}

	if v, ok := o["vendor"]; ok {
		m.Vendor = parseStringValue(v)
	}

	if v, ok := o["iconClass"]; ok {
		m.IconClass = parseStringValue(v)
	}

	return diags
}

func (data *resourceSecurityAppCustomSignaturesModel) getCreateObjectSecurityAppCustomSignatures(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Signature.IsNull() {
		result["signature"] = data.Signature.ValueString()
	}

	if !data.Comment.IsNull() {
		result["comment"] = data.Comment.ValueString()
	}

	if !data.Ftntid.IsNull() {
		result["id"] = data.Ftntid.ValueFloat64()
	}

	if !data.Tag.IsNull() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Name.IsNull() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Category.IsNull() {
		result["category"] = data.Category.ValueFloat64()
	}

	if !data.Protocol.IsNull() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.Technology.IsNull() {
		result["technology"] = data.Technology.ValueString()
	}

	if !data.Behavior.IsNull() {
		result["behavior"] = data.Behavior.ValueString()
	}

	if !data.Vendor.IsNull() {
		result["vendor"] = data.Vendor.ValueString()
	}

	if !data.IconClass.IsNull() {
		result["iconClass"] = data.IconClass.ValueString()
	}

	return &result
}

func (data *resourceSecurityAppCustomSignaturesModel) getUpdateObjectSecurityAppCustomSignatures(ctx context.Context, state resourceSecurityAppCustomSignaturesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Signature.IsNull() {
		result["signature"] = data.Signature.ValueString()
	}

	if !data.Comment.IsNull() {
		result["comment"] = data.Comment.ValueString()
	}

	if !data.Ftntid.IsNull() {
		result["id"] = data.Ftntid.ValueFloat64()
	}

	if !data.Tag.IsNull() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Name.IsNull() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Category.IsNull() {
		result["category"] = data.Category.ValueFloat64()
	}

	if !data.Protocol.IsNull() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.Technology.IsNull() {
		result["technology"] = data.Technology.ValueString()
	}

	if !data.Behavior.IsNull() {
		result["behavior"] = data.Behavior.ValueString()
	}

	if !data.Vendor.IsNull() {
		result["vendor"] = data.Vendor.ValueString()
	}

	if !data.IconClass.IsNull() {
		result["iconClass"] = data.IconClass.ValueString()
	}

	return &result
}

func (data *resourceSecurityAppCustomSignaturesModel) getURLObjectSecurityAppCustomSignatures(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
