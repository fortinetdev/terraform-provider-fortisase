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
var _ resource.Resource = &resourceAuthSwgSamlServer{}

func newResourceAuthSwgSamlServer() resource.Resource {
	return &resourceAuthSwgSamlServer{}
}

type resourceAuthSwgSamlServer struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceAuthSwgSamlServerModel describes the resource data model.
type resourceAuthSwgSamlServerModel struct {
	ID                       types.String                                  `tfsdk:"id"`
	PrimaryKey               types.String                                  `tfsdk:"primary_key"`
	IdpEntityId              types.String                                  `tfsdk:"idp_entity_id"`
	IdpSignOnUrl             types.String                                  `tfsdk:"idp_sign_on_url"`
	IdpLogOutUrl             types.String                                  `tfsdk:"idp_log_out_url"`
	Username                 types.String                                  `tfsdk:"username"`
	GroupName                types.String                                  `tfsdk:"group_name"`
	GroupMatch               types.String                                  `tfsdk:"group_match"`
	SpCert                   *resourceAuthSwgSamlServerSpCertModel         `tfsdk:"sp_cert"`
	IdpCertificate           *resourceAuthSwgSamlServerIdpCertificateModel `tfsdk:"idp_certificate"`
	DigestMethod             types.String                                  `tfsdk:"digest_method"`
	ScimEnabled              types.Bool                                    `tfsdk:"scim_enabled"`
	RequireSignedRespAndAsrt types.Bool                                    `tfsdk:"require_signed_resp_and_asrt"`
	Scim                     *resourceAuthSwgSamlServerScimModel           `tfsdk:"scim"`
}

func (r *resourceAuthSwgSamlServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_swg_saml_server"
}

func (r *resourceAuthSwgSamlServer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SWG User SSO Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.OneOf("$sase-global"),
				},
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"idp_entity_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"idp_sign_on_url": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"idp_log_out_url": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"username": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"group_name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"group_match": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"digest_method": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("sha256", "sha1"),
				},
				Computed: true,
				Optional: true,
			},
			"scim_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"require_signed_resp_and_asrt": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"sp_cert": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("system/certificate/local-certificates"),
						},
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"idp_certificate": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("system/certificate/remote-certificates"),
						},
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"scim": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"scim_url": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"auth_method": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("token"),
						},
						Computed: true,
						Optional: true,
					},
					"token": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.LengthBetween(1, 128),
						},
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

func (r *resourceAuthSwgSamlServer) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_swg_saml_server"
}

func (r *resourceAuthSwgSamlServer) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceAuthSwgSamlServerModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectAuthSwgSamlServer(ctx, diags))
	input_model.BodyParams["enabled"] = true

	if diags.HasError() {
		return
	}
	output, err := c.UpdateAuthSwgSamlServer(&input_model)
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

	read_output, err := c.ReadAuthSwgSamlServer(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthSwgSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthSwgSamlServer) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthSwgSamlServer")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceAuthSwgSamlServerModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceAuthSwgSamlServerModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectAuthSwgSamlServer(ctx, state, diags))
	input_model.BodyParams["enabled"] = true

	if diags.HasError() {
		return
	}

	output, err := c.UpdateAuthSwgSamlServer(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output, err := c.ReadAuthSwgSamlServer(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthSwgSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthSwgSamlServer) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("AuthSwgSamlServer")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceAuthSwgSamlServerModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	result_model := make(map[string]interface{})
	result_model["primaryKey"] = mkey
	result_model["enabled"] = false
	input_model.BodyParams = result_model

	output, err := c.UpdateAuthSwgSamlServer(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceAuthSwgSamlServer) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthSwgSamlServerModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadAuthSwgSamlServer(&input_model)
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

	diags.Append(data.refreshAuthSwgSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthSwgSamlServer) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceAuthSwgSamlServerModel) refreshAuthSwgSamlServer(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["idpEntityId"]; ok {
		m.IdpEntityId = parseStringValue(v)
	}

	if v, ok := o["idpSignOnUrl"]; ok {
		m.IdpSignOnUrl = parseStringValue(v)
	}

	if v, ok := o["idpLogOutUrl"]; ok {
		m.IdpLogOutUrl = parseStringValue(v)
	}

	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	if v, ok := o["groupName"]; ok {
		m.GroupName = parseStringValue(v)
	}

	if v, ok := o["groupMatch"]; ok {
		m.GroupMatch = parseStringValue(v)
	}

	if v, ok := o["spCert"]; ok {
		m.SpCert = m.SpCert.flattenAuthSwgSamlServerSpCert(ctx, v, &diags)
	}

	if v, ok := o["idpCertificate"]; ok {
		m.IdpCertificate = m.IdpCertificate.flattenAuthSwgSamlServerIdpCertificate(ctx, v, &diags)
	}

	if v, ok := o["digestMethod"]; ok {
		m.DigestMethod = parseStringValue(v)
	}

	if v, ok := o["scimEnabled"]; ok {
		m.ScimEnabled = parseBoolValue(v)
	}

	if v, ok := o["requireSignedRespAndAsrt"]; ok {
		m.RequireSignedRespAndAsrt = parseBoolValue(v)
	}

	if v, ok := o["scim"]; ok {
		m.Scim = m.Scim.flattenAuthSwgSamlServerScim(ctx, v, &diags)
	}

	return diags
}

func (data *resourceAuthSwgSamlServerModel) getCreateObjectAuthSwgSamlServer(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.IdpEntityId.IsNull() && !data.IdpEntityId.IsUnknown() {
		result["idpEntityId"] = data.IdpEntityId.ValueString()
	}

	if !data.IdpSignOnUrl.IsNull() && !data.IdpSignOnUrl.IsUnknown() {
		result["idpSignOnUrl"] = data.IdpSignOnUrl.ValueString()
	}

	if !data.IdpLogOutUrl.IsNull() && !data.IdpLogOutUrl.IsUnknown() {
		result["idpLogOutUrl"] = data.IdpLogOutUrl.ValueString()
	}

	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		result["username"] = data.Username.ValueString()
	}

	if !data.GroupName.IsNull() && !data.GroupName.IsUnknown() {
		result["groupName"] = data.GroupName.ValueString()
	}

	if !data.GroupMatch.IsNull() && !data.GroupMatch.IsUnknown() {
		result["groupMatch"] = data.GroupMatch.ValueString()
	}

	result["spCert"] = nil
	if data.SpCert != nil && !isZeroStruct(*data.SpCert) {
		result["spCert"] = data.SpCert.expandAuthSwgSamlServerSpCert(ctx, diags)
	}

	result["idpCertificate"] = nil
	if data.IdpCertificate != nil && !isZeroStruct(*data.IdpCertificate) {
		result["idpCertificate"] = data.IdpCertificate.expandAuthSwgSamlServerIdpCertificate(ctx, diags)
	}

	if !data.DigestMethod.IsNull() && !data.DigestMethod.IsUnknown() {
		result["digestMethod"] = data.DigestMethod.ValueString()
	}

	if !data.ScimEnabled.IsNull() && !data.ScimEnabled.IsUnknown() {
		result["scimEnabled"] = data.ScimEnabled.ValueBool()
	}

	if !data.RequireSignedRespAndAsrt.IsNull() && !data.RequireSignedRespAndAsrt.IsUnknown() {
		result["requireSignedRespAndAsrt"] = data.RequireSignedRespAndAsrt.ValueBool()
	}

	if data.Scim != nil && !isZeroStruct(*data.Scim) {
		result["scim"] = data.Scim.expandAuthSwgSamlServerScim(ctx, diags)
	}

	return &result
}

func (data *resourceAuthSwgSamlServerModel) getUpdateObjectAuthSwgSamlServer(ctx context.Context, state resourceAuthSwgSamlServerModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.IdpEntityId.IsNull() && !data.IdpEntityId.IsUnknown() {
		result["idpEntityId"] = data.IdpEntityId.ValueString()
	}

	if !data.IdpSignOnUrl.IsNull() && !data.IdpSignOnUrl.IsUnknown() {
		result["idpSignOnUrl"] = data.IdpSignOnUrl.ValueString()
	}

	if !data.IdpLogOutUrl.IsNull() && !data.IdpLogOutUrl.IsUnknown() {
		result["idpLogOutUrl"] = data.IdpLogOutUrl.ValueString()
	}

	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		result["username"] = data.Username.ValueString()
	}

	if !data.GroupName.IsNull() && !data.GroupName.IsUnknown() {
		result["groupName"] = data.GroupName.ValueString()
	}

	if !data.GroupMatch.IsNull() && !data.GroupMatch.IsUnknown() {
		result["groupMatch"] = data.GroupMatch.ValueString()
	}

	result["spCert"] = nil
	if data.SpCert != nil && !isZeroStruct(*data.SpCert) {
		result["spCert"] = data.SpCert.expandAuthSwgSamlServerSpCert(ctx, diags)
	}

	result["idpCertificate"] = nil
	if data.IdpCertificate != nil && !isZeroStruct(*data.IdpCertificate) {
		result["idpCertificate"] = data.IdpCertificate.expandAuthSwgSamlServerIdpCertificate(ctx, diags)
	}

	if !data.DigestMethod.IsNull() && !data.DigestMethod.IsUnknown() {
		result["digestMethod"] = data.DigestMethod.ValueString()
	}

	if !data.ScimEnabled.IsNull() && !data.ScimEnabled.IsUnknown() {
		result["scimEnabled"] = data.ScimEnabled.ValueBool()
	}

	if !data.RequireSignedRespAndAsrt.IsNull() && !data.RequireSignedRespAndAsrt.IsUnknown() {
		result["requireSignedRespAndAsrt"] = data.RequireSignedRespAndAsrt.ValueBool()
	}

	if data.Scim != nil {
		result["scim"] = data.Scim.expandAuthSwgSamlServerScim(ctx, diags)
	}

	return &result
}

type resourceAuthSwgSamlServerSpCertModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceAuthSwgSamlServerIdpCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceAuthSwgSamlServerScimModel struct {
	ScimUrl    types.String `tfsdk:"scim_url"`
	AuthMethod types.String `tfsdk:"auth_method"`
	Token      types.String `tfsdk:"token"`
}

func (m *resourceAuthSwgSamlServerSpCertModel) flattenAuthSwgSamlServerSpCert(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthSwgSamlServerSpCertModel {
	if input == nil {
		return &resourceAuthSwgSamlServerSpCertModel{}
	}
	if m == nil {
		m = &resourceAuthSwgSamlServerSpCertModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (m *resourceAuthSwgSamlServerIdpCertificateModel) flattenAuthSwgSamlServerIdpCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthSwgSamlServerIdpCertificateModel {
	if input == nil {
		return &resourceAuthSwgSamlServerIdpCertificateModel{}
	}
	if m == nil {
		m = &resourceAuthSwgSamlServerIdpCertificateModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (m *resourceAuthSwgSamlServerScimModel) flattenAuthSwgSamlServerScim(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthSwgSamlServerScimModel {
	if input == nil {
		return &resourceAuthSwgSamlServerScimModel{}
	}
	if m == nil {
		m = &resourceAuthSwgSamlServerScimModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["scimUrl"]; ok {
		m.ScimUrl = parseStringValue(v)
	}

	if v, ok := o["authMethod"]; ok {
		m.AuthMethod = parseStringValue(v)
	}

	if v, ok := o["token"]; ok {
		m.Token = parseStringValue(v)
	}

	return m
}

func (data *resourceAuthSwgSamlServerSpCertModel) expandAuthSwgSamlServerSpCert(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceAuthSwgSamlServerIdpCertificateModel) expandAuthSwgSamlServerIdpCertificate(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceAuthSwgSamlServerScimModel) expandAuthSwgSamlServerScim(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ScimUrl.IsNull() && !data.ScimUrl.IsUnknown() {
		result["scimUrl"] = data.ScimUrl.ValueString()
	}

	if !data.AuthMethod.IsNull() && !data.AuthMethod.IsUnknown() {
		result["authMethod"] = data.AuthMethod.ValueString()
	}

	if !data.Token.IsNull() && !data.Token.IsUnknown() {
		result["token"] = data.Token.ValueString()
	}

	return result
}
