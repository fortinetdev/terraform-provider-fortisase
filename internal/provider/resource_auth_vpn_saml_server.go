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
	"time"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceAuthVpnSamlServer{}

func newResourceAuthVpnSamlServer() resource.Resource {
	return &resourceAuthVpnSamlServer{}
}

type resourceAuthVpnSamlServer struct {
	fortiClient *FortiClient
}

// resourceAuthVpnSamlServerModel describes the resource data model.
type resourceAuthVpnSamlServerModel struct {
	ID             types.String                                  `tfsdk:"id"`
	PrimaryKey     types.String                                  `tfsdk:"primary_key"`
	Enabled        types.Bool                                    `tfsdk:"enabled"`
	IdpEntityId    types.String                                  `tfsdk:"idp_entity_id"`
	IdpSignOnUrl   types.String                                  `tfsdk:"idp_sign_on_url"`
	IdpLogOutUrl   types.String                                  `tfsdk:"idp_log_out_url"`
	Username       types.String                                  `tfsdk:"username"`
	GroupName      types.String                                  `tfsdk:"group_name"`
	GroupId        types.String                                  `tfsdk:"group_id"`
	SpCert         *resourceAuthVpnSamlServerSpCertModel         `tfsdk:"sp_cert"`
	IdpCertificate *resourceAuthVpnSamlServerIdpCertificateModel `tfsdk:"idp_certificate"`
	DigestMethod   types.String                                  `tfsdk:"digest_method"`
	EntraIdEnabled types.Bool                                    `tfsdk:"entra_id_enabled"`
	ScimEnabled    types.Bool                                    `tfsdk:"scim_enabled"`
	DomainName     types.String                                  `tfsdk:"domain_name"`
	ApplicationId  types.String                                  `tfsdk:"application_id"`
}

func (r *resourceAuthVpnSamlServer) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_vpn_saml_server"
}

func (r *resourceAuthVpnSamlServer) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.OneOf("$sase-global"),
				},
				Required: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"idp_entity_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"idp_sign_on_url": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"idp_log_out_url": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"username": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"group_name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
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
					stringvalidator.OneOf("sha256", "sha1"),
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
			"domain_name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"application_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"sp_cert": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("system/certificate/local-certificates"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"idp_certificate": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("system/certificate/remote-certificates"),
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

func (r *resourceAuthVpnSamlServer) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
}

func (r *resourceAuthVpnSamlServer) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceAuthVpnSamlServerModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectAuthVpnSamlServer(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateAuthVpnSamlServer(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output := make(map[string]interface{})
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadAuthVpnSamlServer(&read_input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to read resource: %v", err),
				"",
			)
			return
		}
		if v, ok := read_output["$meta"].(map[string]interface{})["state"]; ok {
			if v != "done" {
				continue
			}
		}
		break
	}

	diags.Append(data.refreshAuthVpnSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthVpnSamlServer) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceAuthVpnSamlServerModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceAuthVpnSamlServerModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectAuthVpnSamlServer(ctx, state, diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateAuthVpnSamlServer(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output := make(map[string]interface{})
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadAuthVpnSamlServer(&read_input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to read resource: %v", err),
				"",
			)
			return
		}
		if v, ok := read_output["$meta"].(map[string]interface{})["state"]; ok {
			if v != "done" {
				continue
			}
		}
		break
	}

	diags.Append(data.refreshAuthVpnSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthVpnSamlServer) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthVpnSamlServerModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	if !data.Enabled.ValueBool() {
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

	_, err := c.UpdateAuthVpnSamlServer(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_output := make(map[string]interface{})
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadAuthVpnSamlServer(&read_input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to read resource: %v", err),
				"",
			)
			return
		}
		if v, ok := read_output["$meta"]; ok {
			// if "state" is not in the map, return success
			if _, ok := v.(map[string]interface{})["state"]; !ok {
				return
			}
		}
	}
	diags.AddError(
		fmt.Sprintf("Error to delete resource: %v", err),
		fmt.Sprintf("The resource still exists: %v", read_output),
	)
}

func (r *resourceAuthVpnSamlServer) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthVpnSamlServerModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadAuthVpnSamlServer(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshAuthVpnSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthVpnSamlServer) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceAuthVpnSamlServerModel) refreshAuthVpnSamlServer(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.SpCert = m.SpCert.flattenAuthVpnSamlServerSpCert(ctx, v, &diags)
	}

	if v, ok := o["idpCertificate"]; ok {
		m.IdpCertificate = m.IdpCertificate.flattenAuthVpnSamlServerIdpCertificate(ctx, v, &diags)
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

	if v, ok := o["domainName"]; ok {
		m.DomainName = parseStringValue(v)
	}

	if v, ok := o["applicationId"]; ok {
		m.ApplicationId = parseStringValue(v)
	}

	return diags
}

func (data *resourceAuthVpnSamlServerModel) getCreateObjectAuthVpnSamlServer(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.IdpEntityId.IsNull() {
		result["idpEntityId"] = data.IdpEntityId.ValueString()
	}

	if !data.IdpSignOnUrl.IsNull() {
		result["idpSignOnUrl"] = data.IdpSignOnUrl.ValueString()
	}

	if !data.IdpLogOutUrl.IsNull() {
		result["idpLogOutUrl"] = data.IdpLogOutUrl.ValueString()
	}

	if !data.Username.IsNull() {
		result["username"] = data.Username.ValueString()
	}

	if !data.GroupName.IsNull() {
		result["groupName"] = data.GroupName.ValueString()
	}

	if !data.GroupId.IsNull() {
		result["groupId"] = data.GroupId.ValueString()
	}

	if data.SpCert != nil && !isZeroStruct(*data.SpCert) {
		result["spCert"] = data.SpCert.expandAuthVpnSamlServerSpCert(ctx, diags)
	}

	if data.IdpCertificate != nil && !isZeroStruct(*data.IdpCertificate) {
		result["idpCertificate"] = data.IdpCertificate.expandAuthVpnSamlServerIdpCertificate(ctx, diags)
	}

	if !data.DigestMethod.IsNull() {
		result["digestMethod"] = data.DigestMethod.ValueString()
	}

	if !data.EntraIdEnabled.IsNull() {
		result["entraIdEnabled"] = data.EntraIdEnabled.ValueBool()
	}

	if !data.ScimEnabled.IsNull() {
		result["scimEnabled"] = data.ScimEnabled.ValueBool()
	}

	if !data.DomainName.IsNull() {
		result["domainName"] = data.DomainName.ValueString()
	}

	if !data.ApplicationId.IsNull() {
		result["applicationId"] = data.ApplicationId.ValueString()
	}

	return &result
}

func (data *resourceAuthVpnSamlServerModel) getUpdateObjectAuthVpnSamlServer(ctx context.Context, state resourceAuthVpnSamlServerModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.IdpEntityId.IsNull() && !data.IdpEntityId.Equal(state.IdpEntityId) {
		result["idpEntityId"] = data.IdpEntityId.ValueString()
	}

	if !data.IdpSignOnUrl.IsNull() && !data.IdpSignOnUrl.Equal(state.IdpSignOnUrl) {
		result["idpSignOnUrl"] = data.IdpSignOnUrl.ValueString()
	}

	if !data.IdpLogOutUrl.IsNull() && !data.IdpLogOutUrl.Equal(state.IdpLogOutUrl) {
		result["idpLogOutUrl"] = data.IdpLogOutUrl.ValueString()
	}

	if !data.Username.IsNull() && !data.Username.Equal(state.Username) {
		result["username"] = data.Username.ValueString()
	}

	if !data.GroupName.IsNull() && !data.GroupName.Equal(state.GroupName) {
		result["groupName"] = data.GroupName.ValueString()
	}

	if !data.GroupId.IsNull() && !data.GroupId.Equal(state.GroupId) {
		result["groupId"] = data.GroupId.ValueString()
	}

	if data.SpCert != nil && !isSameStruct(data.SpCert, state.SpCert) {
		result["spCert"] = data.SpCert.expandAuthVpnSamlServerSpCert(ctx, diags)
	}

	if data.IdpCertificate != nil && !isSameStruct(data.IdpCertificate, state.IdpCertificate) {
		result["idpCertificate"] = data.IdpCertificate.expandAuthVpnSamlServerIdpCertificate(ctx, diags)
	}

	if !data.DigestMethod.IsNull() && !data.DigestMethod.Equal(state.DigestMethod) {
		result["digestMethod"] = data.DigestMethod.ValueString()
	}

	if !data.EntraIdEnabled.IsNull() && !data.EntraIdEnabled.Equal(state.EntraIdEnabled) {
		result["entraIdEnabled"] = data.EntraIdEnabled.ValueBool()
	}

	if !data.ScimEnabled.IsNull() && !data.ScimEnabled.Equal(state.ScimEnabled) {
		result["scimEnabled"] = data.ScimEnabled.ValueBool()
	}

	if !data.DomainName.IsNull() && !data.DomainName.Equal(state.DomainName) {
		result["domainName"] = data.DomainName.ValueString()
	}

	if !data.ApplicationId.IsNull() && !data.ApplicationId.Equal(state.ApplicationId) {
		result["applicationId"] = data.ApplicationId.ValueString()
	}

	return &result
}

type resourceAuthVpnSamlServerSpCertModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceAuthVpnSamlServerIdpCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceAuthVpnSamlServerSpCertModel) flattenAuthVpnSamlServerSpCert(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthVpnSamlServerSpCertModel {
	if input == nil {
		return &resourceAuthVpnSamlServerSpCertModel{}
	}
	if m == nil {
		m = &resourceAuthVpnSamlServerSpCertModel{}
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

func (m *resourceAuthVpnSamlServerIdpCertificateModel) flattenAuthVpnSamlServerIdpCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthVpnSamlServerIdpCertificateModel {
	if input == nil {
		return &resourceAuthVpnSamlServerIdpCertificateModel{}
	}
	if m == nil {
		m = &resourceAuthVpnSamlServerIdpCertificateModel{}
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

func (data *resourceAuthVpnSamlServerSpCertModel) expandAuthVpnSamlServerSpCert(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceAuthVpnSamlServerIdpCertificateModel) expandAuthVpnSamlServerIdpCertificate(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
