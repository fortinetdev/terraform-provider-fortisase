// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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
var _ resource.Resource = &resourceSecurityAntivirusProfile{}

func newResourceSecurityAntivirusProfile() resource.Resource {
	return &resourceSecurityAntivirusProfile{}
}

type resourceSecurityAntivirusProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityAntivirusProfileModel describes the resource data model.
type resourceSecurityAntivirusProfileModel struct {
	ID         types.String                              `tfsdk:"id"`
	PrimaryKey types.String                              `tfsdk:"primary_key"`
	Http       types.String                              `tfsdk:"http"`
	Smtp       types.String                              `tfsdk:"smtp"`
	Pop3       types.String                              `tfsdk:"pop3"`
	Imap       types.String                              `tfsdk:"imap"`
	Ftp        types.String                              `tfsdk:"ftp"`
	Cifs       types.String                              `tfsdk:"cifs"`
	Cdr        *resourceSecurityAntivirusProfileCdrModel `tfsdk:"cdr"`
	Direction  types.String                              `tfsdk:"direction"`
}

func (r *resourceSecurityAntivirusProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_antivirus_profile"
}

func (r *resourceSecurityAntivirusProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Required: true,
			},
			"http": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"smtp": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"pop3": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"imap": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"ftp": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"cifs": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				MarkdownDescription: "The direction of the target resource.\nSupported values: internal-profiles, outbound-profiles.",
				Computed:            true,
				Optional:            true,
			},
			"cdr": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"enable": schema.BoolAttribute{
						Computed: true,
						Optional: true,
					},
					"file_types": schema.SetAttribute{
						Validators: []validator.Set{
							setvalidator.ValueStringsAre(
								stringvalidator.OneOf("pdf", "office"),
							),
						},
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"allow_error_transmission": schema.BoolAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityAntivirusProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_antivirus_profile"
}

func (r *resourceSecurityAntivirusProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityAntivirusProfile")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityAntivirusProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecurityAntivirusProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityAntivirusProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecurityAntivirusProfile(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityAntivirusProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityAntivirusProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityAntivirusProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityAntivirusProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityAntivirusProfile")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityAntivirusProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityAntivirusProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityAntivirusProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityAntivirusProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityAntivirusProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityAntivirusProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityAntivirusProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityAntivirusProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityAntivirusProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityAntivirusProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityAntivirusProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityAntivirusProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityAntivirusProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityAntivirusProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityAntivirusProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityAntivirusProfileModel) refreshSecurityAntivirusProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["http"]; ok {
		m.Http = parseStringValue(v)
	}

	if v, ok := o["smtp"]; ok {
		m.Smtp = parseStringValue(v)
	}

	if v, ok := o["pop3"]; ok {
		m.Pop3 = parseStringValue(v)
	}

	if v, ok := o["imap"]; ok {
		m.Imap = parseStringValue(v)
	}

	if v, ok := o["ftp"]; ok {
		m.Ftp = parseStringValue(v)
	}

	if v, ok := o["cifs"]; ok {
		m.Cifs = parseStringValue(v)
	}

	if v, ok := o["cdr"]; ok {
		m.Cdr = m.Cdr.flattenSecurityAntivirusProfileCdr(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityAntivirusProfileModel) getCreateObjectSecurityAntivirusProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Http.IsNull() {
		result["http"] = data.Http.ValueString()
	}

	if !data.Smtp.IsNull() {
		result["smtp"] = data.Smtp.ValueString()
	}

	if !data.Pop3.IsNull() {
		result["pop3"] = data.Pop3.ValueString()
	}

	if !data.Imap.IsNull() {
		result["imap"] = data.Imap.ValueString()
	}

	if !data.Ftp.IsNull() {
		result["ftp"] = data.Ftp.ValueString()
	}

	if !data.Cifs.IsNull() {
		result["cifs"] = data.Cifs.ValueString()
	}

	if data.Cdr != nil && !isZeroStruct(*data.Cdr) {
		result["cdr"] = data.Cdr.expandSecurityAntivirusProfileCdr(ctx, diags)
	}

	return &result
}

func (data *resourceSecurityAntivirusProfileModel) getUpdateObjectSecurityAntivirusProfile(ctx context.Context, state resourceSecurityAntivirusProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Http.IsNull() {
		result["http"] = data.Http.ValueString()
	}

	if !data.Smtp.IsNull() {
		result["smtp"] = data.Smtp.ValueString()
	}

	if !data.Pop3.IsNull() {
		result["pop3"] = data.Pop3.ValueString()
	}

	if !data.Imap.IsNull() {
		result["imap"] = data.Imap.ValueString()
	}

	if !data.Ftp.IsNull() {
		result["ftp"] = data.Ftp.ValueString()
	}

	if !data.Cifs.IsNull() {
		result["cifs"] = data.Cifs.ValueString()
	}

	if data.Cdr != nil {
		result["cdr"] = data.Cdr.expandSecurityAntivirusProfileCdr(ctx, diags)
	}

	return &result
}

func (data *resourceSecurityAntivirusProfileModel) getURLObjectSecurityAntivirusProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		diags.AddWarning("\"direction\" is deprecated and may be removed in future.",
			"It is recommended to recreate the resource without \"direction\" to avoid unexpected behavior in future.",
		)
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityAntivirusProfileCdrModel struct {
	Enable                 types.Bool `tfsdk:"enable"`
	FileTypes              types.Set  `tfsdk:"file_types"`
	AllowErrorTransmission types.Bool `tfsdk:"allow_error_transmission"`
}

func (m *resourceSecurityAntivirusProfileCdrModel) flattenSecurityAntivirusProfileCdr(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityAntivirusProfileCdrModel {
	if input == nil {
		return &resourceSecurityAntivirusProfileCdrModel{}
	}
	if m == nil {
		m = &resourceSecurityAntivirusProfileCdrModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["enable"]; ok {
		m.Enable = parseBoolValue(v)
	}

	if v, ok := o["fileTypes"]; ok {
		m.FileTypes = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["allowErrorTransmission"]; ok {
		m.AllowErrorTransmission = parseBoolValue(v)
	}

	return m
}

func (data *resourceSecurityAntivirusProfileCdrModel) expandSecurityAntivirusProfileCdr(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Enable.IsNull() {
		result["enable"] = data.Enable.ValueBool()
	}

	if !data.FileTypes.IsNull() {
		result["fileTypes"] = expandSetToStringList(data.FileTypes)
	}

	if !data.AllowErrorTransmission.IsNull() {
		result["allowErrorTransmission"] = data.AllowErrorTransmission.ValueBool()
	}

	return result
}
