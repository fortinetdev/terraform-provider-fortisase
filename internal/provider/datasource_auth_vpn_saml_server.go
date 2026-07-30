// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceAuthVpnSamlServer{}

func newDatasourceAuthVpnSamlServer() datasource.DataSource {
	return &datasourceAuthVpnSamlServer{}
}

type datasourceAuthVpnSamlServer struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceAuthVpnSamlServerModel describes the datasource data model.
type datasourceAuthVpnSamlServerModel struct {
	PrimaryKey               types.String                                    `tfsdk:"primary_key"`
	IdpEntityId              types.String                                    `tfsdk:"idp_entity_id"`
	IdpSignOnUrl             types.String                                    `tfsdk:"idp_sign_on_url"`
	IdpLogOutUrl             types.String                                    `tfsdk:"idp_log_out_url"`
	Username                 types.String                                    `tfsdk:"username"`
	GroupName                types.String                                    `tfsdk:"group_name"`
	GroupId                  types.String                                    `tfsdk:"group_id"`
	SpCert                   *datasourceAuthVpnSamlServerSpCertModel         `tfsdk:"sp_cert"`
	IdpCertificate           *datasourceAuthVpnSamlServerIdpCertificateModel `tfsdk:"idp_certificate"`
	DigestMethod             types.String                                    `tfsdk:"digest_method"`
	EntraIdEnabled           types.Bool                                      `tfsdk:"entra_id_enabled"`
	ScimEnabled              types.Bool                                      `tfsdk:"scim_enabled"`
	RequireSignedRespAndAsrt types.Bool                                      `tfsdk:"require_signed_resp_and_asrt"`
	DomainName               types.String                                    `tfsdk:"domain_name"`
	ApplicationId            types.String                                    `tfsdk:"application_id"`
}

func (r *datasourceAuthVpnSamlServer) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_vpn_saml_server"
}

func (r *datasourceAuthVpnSamlServer) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "VPN User SSO Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("$sase-global"),
				},
				Required: true,
			},
			"idp_entity_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
			},
			"idp_sign_on_url": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
			},
			"idp_log_out_url": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
			},
			"username": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
			},
			"group_name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
			},
			"group_id": schema.StringAttribute{
				Computed: true,
			},
			"digest_method": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("sha256", "sha1"),
				},
				Computed: true,
			},
			"entra_id_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"scim_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"require_signed_resp_and_asrt": schema.BoolAttribute{
				Computed: true,
			},
			"domain_name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
			},
			"application_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
			},
			"sp_cert": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("system/certificate/local-certificates"),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"idp_certificate": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("system/certificate/remote-certificates"),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceAuthVpnSamlServer) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_vpn_saml_server"
}

func (r *datasourceAuthVpnSamlServer) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceAuthVpnSamlServerModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadAuthVpnSamlServer(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthVpnSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceAuthVpnSamlServerModel) refreshAuthVpnSamlServer(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
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

type datasourceAuthVpnSamlServerSpCertModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceAuthVpnSamlServerIdpCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceAuthVpnSamlServerSpCertModel) flattenAuthVpnSamlServerSpCert(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthVpnSamlServerSpCertModel {
	if input == nil {
		return &datasourceAuthVpnSamlServerSpCertModel{}
	}
	if m == nil {
		m = &datasourceAuthVpnSamlServerSpCertModel{}
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

func (m *datasourceAuthVpnSamlServerIdpCertificateModel) flattenAuthVpnSamlServerIdpCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthVpnSamlServerIdpCertificateModel {
	if input == nil {
		return &datasourceAuthVpnSamlServerIdpCertificateModel{}
	}
	if m == nil {
		m = &datasourceAuthVpnSamlServerIdpCertificateModel{}
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
