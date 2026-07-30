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
var _ resource.Resource = &resourceSecurityAppCustomSignature{}
var _ resource.ResourceWithMoveState = &resourceSecurityAppCustomSignature{}

func newResourceSecurityAppCustomSignature() resource.Resource {
	return &resourceSecurityAppCustomSignature{}
}

type resourceSecurityAppCustomSignature struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityAppCustomSignatureModel describes the resource data model.
type resourceSecurityAppCustomSignatureModel struct {
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

func (r *resourceSecurityAppCustomSignature) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_app_custom_signature"
}

func (r *resourceSecurityAppCustomSignature) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Custom Application Signature Resource API V2 for FortiSASE.",
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
			"signature": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(4095),
				},
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
			"ftntid": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"tag": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(63),
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

func (r *resourceSecurityAppCustomSignature) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_app_custom_signature"
}
func (r *resourceSecurityAppCustomSignature) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_app_custom_signatures" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityAppCustomSignatureModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityAppCustomSignature) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityAppCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityAppCustomSignatureModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityAppCustomSignature(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignature(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityAppCustomSignatures(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignature(ctx, "read", diags))

	read_output, err := c.ReadSecurityAppCustomSignatures(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityAppCustomSignature(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityAppCustomSignature) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityAppCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityAppCustomSignatureModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityAppCustomSignatureModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityAppCustomSignature(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignature(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignature(ctx, "read", diags))

	read_output, err := c.ReadSecurityAppCustomSignatures(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityAppCustomSignature(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityAppCustomSignature) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityAppCustomSignatures")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityAppCustomSignatureModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignature(ctx, "delete", diags))

	output, err := c.DeleteSecurityAppCustomSignatures(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityAppCustomSignature) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityAppCustomSignatureModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignature(ctx, "read", diags))

	read_output, err := c.ReadSecurityAppCustomSignatures(&input_model)
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

	diags.Append(data.refreshSecurityAppCustomSignature(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityAppCustomSignature) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityAppCustomSignatureModel) refreshSecurityAppCustomSignature(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceSecurityAppCustomSignatureModel) getCreateObjectSecurityAppCustomSignature(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Signature.IsNull() && !data.Signature.IsUnknown() {
		result["signature"] = data.Signature.ValueString()
	}

	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		result["comment"] = data.Comment.ValueString()
	}

	if !data.Ftntid.IsNull() && !data.Ftntid.IsUnknown() {
		result["id"] = data.Ftntid.ValueFloat64()
	}

	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Category.IsNull() && !data.Category.IsUnknown() {
		result["category"] = data.Category.ValueFloat64()
	}

	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.Technology.IsNull() && !data.Technology.IsUnknown() {
		result["technology"] = data.Technology.ValueString()
	}

	if !data.Behavior.IsNull() && !data.Behavior.IsUnknown() {
		result["behavior"] = data.Behavior.ValueString()
	}

	if !data.Vendor.IsNull() && !data.Vendor.IsUnknown() {
		result["vendor"] = data.Vendor.ValueString()
	}

	if !data.IconClass.IsNull() && !data.IconClass.IsUnknown() {
		result["iconClass"] = data.IconClass.ValueString()
	}

	return &result
}

func (data *resourceSecurityAppCustomSignatureModel) getUpdateObjectSecurityAppCustomSignature(ctx context.Context, state resourceSecurityAppCustomSignatureModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Signature.IsNull() && !data.Signature.IsUnknown() {
		result["signature"] = data.Signature.ValueString()
	}

	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		result["comment"] = data.Comment.ValueString()
	}

	if !data.Ftntid.IsNull() && !data.Ftntid.IsUnknown() {
		result["id"] = data.Ftntid.ValueFloat64()
	}

	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Category.IsNull() && !data.Category.IsUnknown() {
		result["category"] = data.Category.ValueFloat64()
	}

	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.Technology.IsNull() && !data.Technology.IsUnknown() {
		result["technology"] = data.Technology.ValueString()
	}

	if !data.Behavior.IsNull() && !data.Behavior.IsUnknown() {
		result["behavior"] = data.Behavior.ValueString()
	}

	if !data.Vendor.IsNull() && !data.Vendor.IsUnknown() {
		result["vendor"] = data.Vendor.ValueString()
	}

	if !data.IconClass.IsNull() && !data.IconClass.IsUnknown() {
		result["iconClass"] = data.IconClass.ValueString()
	}

	return &result
}

func (data *resourceSecurityAppCustomSignatureModel) getURLObjectSecurityAppCustomSignature(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
