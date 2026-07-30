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
var _ resource.Resource = &resourceAuthSslvpnSamlServer{}

func newResourceAuthSslvpnSamlServer() resource.Resource {
	return &resourceAuthSslvpnSamlServer{}
}

type resourceAuthSslvpnSamlServer struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceAuthSslvpnSamlServerModel describes the resource data model.
type resourceAuthSslvpnSamlServerModel struct {
	ID                       types.String                                     `tfsdk:"id"`
	PrimaryKey               types.String                                     `tfsdk:"primary_key"`
	Enabled                  types.Bool                                       `tfsdk:"enabled"`
	IdpEntityId              types.String                                     `tfsdk:"idp_entity_id"`
	IdpSignOnUrl             types.String                                     `tfsdk:"idp_sign_on_url"`
	IdpLogOutUrl             types.String                                     `tfsdk:"idp_log_out_url"`
	Username                 types.String                                     `tfsdk:"username"`
	GroupName                types.String                                     `tfsdk:"group_name"`
	GroupId                  types.String                                     `tfsdk:"group_id"`
	SpCert                   *resourceAuthSslvpnSamlServerSpCertModel         `tfsdk:"sp_cert"`
	IdpCertificate           *resourceAuthSslvpnSamlServerIdpCertificateModel `tfsdk:"idp_certificate"`
	DigestMethod             types.String                                     `tfsdk:"digest_method"`
	EntraIdEnabled           types.Bool                                       `tfsdk:"entra_id_enabled"`
	ScimEnabled              types.Bool                                       `tfsdk:"scim_enabled"`
	RequireSignedRespAndAsrt types.Bool                                       `tfsdk:"require_signed_resp_and_asrt"`
	DomainName               types.String                                     `tfsdk:"domain_name"`
	ApplicationId            types.String                                     `tfsdk:"application_id"`
}

func (r *resourceAuthSslvpnSamlServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_sslvpn_saml_server"
}

func (r *resourceAuthSslvpnSamlServer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SSLVPN User SSO Resource API V2 for FortiSASE.",
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
			"enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
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
			"group_id": schema.StringAttribute{
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
			"entra_id_enabled": schema.BoolAttribute{
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
			"domain_name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"application_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
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
		},
	}
}

func (r *resourceAuthSslvpnSamlServer) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_sslvpn_saml_server"
}

func (r *resourceAuthSslvpnSamlServer) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthSslvpnSamlServer")
	lock.Lock()
	defer lock.Unlock()
	var data resourceAuthSslvpnSamlServerModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectAuthSslvpnSamlServer(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateAuthSslvpnSamlServer(&input_model)
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

	read_output, err := c.ReadAuthSslvpnSamlServer(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthSslvpnSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthSslvpnSamlServer) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthSslvpnSamlServer")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceAuthSslvpnSamlServerModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceAuthSslvpnSamlServerModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectAuthSslvpnSamlServer(ctx, state, diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateAuthSslvpnSamlServer(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output, err := c.ReadAuthSslvpnSamlServer(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthSslvpnSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthSslvpnSamlServer) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceAuthSslvpnSamlServer) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthSslvpnSamlServerModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadAuthSslvpnSamlServer(&input_model)
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

	diags.Append(data.refreshAuthSslvpnSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthSslvpnSamlServer) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceAuthSslvpnSamlServerModel) refreshAuthSslvpnSamlServer(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
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

	if v, ok := o["groupId"]; ok {
		m.GroupId = parseStringValue(v)
	}

	if v, ok := o["spCert"]; ok {
		m.SpCert = m.SpCert.flattenAuthSslvpnSamlServerSpCert(ctx, v, &diags)
	}

	if v, ok := o["idpCertificate"]; ok {
		m.IdpCertificate = m.IdpCertificate.flattenAuthSslvpnSamlServerIdpCertificate(ctx, v, &diags)
	}

	if v, ok := o["digestMethod"]; ok {
		m.DigestMethod = parseStringValue(v)
	}

	if v, ok := o["entraIdEnabled"]; ok {
		m.EntraIdEnabled = parseBoolValue(v)
	}

	if v, ok := o["scimEnabled"]; ok {
		m.ScimEnabled = parseBoolValue(v)
	}

	if v, ok := o["requireSignedRespAndAsrt"]; ok {
		m.RequireSignedRespAndAsrt = parseBoolValue(v)
	}

	if v, ok := o["domainName"]; ok {
		m.DomainName = parseStringValue(v)
	}

	if v, ok := o["applicationId"]; ok {
		m.ApplicationId = parseStringValue(v)
	}

	return diags
}

func (data *resourceAuthSslvpnSamlServerModel) getCreateObjectAuthSslvpnSamlServer(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
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

	if !data.GroupId.IsNull() && !data.GroupId.IsUnknown() {
		result["groupId"] = data.GroupId.ValueString()
	}

	result["spCert"] = nil
	if data.SpCert != nil && !isZeroStruct(*data.SpCert) {
		result["spCert"] = data.SpCert.expandAuthSslvpnSamlServerSpCert(ctx, diags)
	}

	result["idpCertificate"] = nil
	if data.IdpCertificate != nil && !isZeroStruct(*data.IdpCertificate) {
		result["idpCertificate"] = data.IdpCertificate.expandAuthSslvpnSamlServerIdpCertificate(ctx, diags)
	}

	if !data.DigestMethod.IsNull() && !data.DigestMethod.IsUnknown() {
		result["digestMethod"] = data.DigestMethod.ValueString()
	}

	if !data.EntraIdEnabled.IsNull() && !data.EntraIdEnabled.IsUnknown() {
		result["entraIdEnabled"] = data.EntraIdEnabled.ValueBool()
	}

	if !data.ScimEnabled.IsNull() && !data.ScimEnabled.IsUnknown() {
		result["scimEnabled"] = data.ScimEnabled.ValueBool()
	}

	if !data.RequireSignedRespAndAsrt.IsNull() && !data.RequireSignedRespAndAsrt.IsUnknown() {
		result["requireSignedRespAndAsrt"] = data.RequireSignedRespAndAsrt.ValueBool()
	}

	if !data.DomainName.IsNull() && !data.DomainName.IsUnknown() {
		result["domainName"] = data.DomainName.ValueString()
	}

	if !data.ApplicationId.IsNull() && !data.ApplicationId.IsUnknown() {
		result["applicationId"] = data.ApplicationId.ValueString()
	}

	return &result
}

func (data *resourceAuthSslvpnSamlServerModel) getUpdateObjectAuthSslvpnSamlServer(ctx context.Context, state resourceAuthSslvpnSamlServerModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
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

	if !data.GroupId.IsNull() && !data.GroupId.IsUnknown() {
		result["groupId"] = data.GroupId.ValueString()
	}

	result["spCert"] = nil
	if data.SpCert != nil && !isZeroStruct(*data.SpCert) {
		result["spCert"] = data.SpCert.expandAuthSslvpnSamlServerSpCert(ctx, diags)
	}

	result["idpCertificate"] = nil
	if data.IdpCertificate != nil && !isZeroStruct(*data.IdpCertificate) {
		result["idpCertificate"] = data.IdpCertificate.expandAuthSslvpnSamlServerIdpCertificate(ctx, diags)
	}

	if !data.DigestMethod.IsNull() && !data.DigestMethod.IsUnknown() {
		result["digestMethod"] = data.DigestMethod.ValueString()
	}

	if !data.EntraIdEnabled.IsNull() && !data.EntraIdEnabled.IsUnknown() {
		result["entraIdEnabled"] = data.EntraIdEnabled.ValueBool()
	}

	if !data.ScimEnabled.IsNull() && !data.ScimEnabled.IsUnknown() {
		result["scimEnabled"] = data.ScimEnabled.ValueBool()
	}

	if !data.RequireSignedRespAndAsrt.IsNull() && !data.RequireSignedRespAndAsrt.IsUnknown() {
		result["requireSignedRespAndAsrt"] = data.RequireSignedRespAndAsrt.ValueBool()
	}

	if !data.DomainName.IsNull() && !data.DomainName.IsUnknown() {
		result["domainName"] = data.DomainName.ValueString()
	}

	if !data.ApplicationId.IsNull() && !data.ApplicationId.IsUnknown() {
		result["applicationId"] = data.ApplicationId.ValueString()
	}

	return &result
}

type resourceAuthSslvpnSamlServerSpCertModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceAuthSslvpnSamlServerIdpCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceAuthSslvpnSamlServerSpCertModel) flattenAuthSslvpnSamlServerSpCert(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthSslvpnSamlServerSpCertModel {
	if input == nil {
		return &resourceAuthSslvpnSamlServerSpCertModel{}
	}
	if m == nil {
		m = &resourceAuthSslvpnSamlServerSpCertModel{}
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

func (m *resourceAuthSslvpnSamlServerIdpCertificateModel) flattenAuthSslvpnSamlServerIdpCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthSslvpnSamlServerIdpCertificateModel {
	if input == nil {
		return &resourceAuthSslvpnSamlServerIdpCertificateModel{}
	}
	if m == nil {
		m = &resourceAuthSslvpnSamlServerIdpCertificateModel{}
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

func (data *resourceAuthSslvpnSamlServerSpCertModel) expandAuthSslvpnSamlServerSpCert(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceAuthSslvpnSamlServerIdpCertificateModel) expandAuthSslvpnSamlServerIdpCertificate(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
