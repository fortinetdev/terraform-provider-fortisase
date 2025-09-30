// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceAuthSwgSamlServer{}

func newDatasourceAuthSwgSamlServer() datasource.DataSource {
	return &datasourceAuthSwgSamlServer{}
}

type datasourceAuthSwgSamlServer struct {
	fortiClient *FortiClient
}

// datasourceAuthSwgSamlServerModel describes the datasource data model.
type datasourceAuthSwgSamlServerModel struct {
	PrimaryKey     types.String                                    `tfsdk:"primary_key"`
	Enabled        types.Bool                                      `tfsdk:"enabled"`
	IdpEntityId    types.String                                    `tfsdk:"idp_entity_id"`
	IdpSignOnUrl   types.String                                    `tfsdk:"idp_sign_on_url"`
	IdpLogOutUrl   types.String                                    `tfsdk:"idp_log_out_url"`
	Username       types.String                                    `tfsdk:"username"`
	GroupName      types.String                                    `tfsdk:"group_name"`
	GroupMatch     types.String                                    `tfsdk:"group_match"`
	SpCert         *datasourceAuthSwgSamlServerSpCertModel         `tfsdk:"sp_cert"`
	IdpCertificate *datasourceAuthSwgSamlServerIdpCertificateModel `tfsdk:"idp_certificate"`
	DigestMethod   types.String                                    `tfsdk:"digest_method"`
	ScimEnabled    types.Bool                                      `tfsdk:"scim_enabled"`
	Scim           *datasourceAuthSwgSamlServerScimModel           `tfsdk:"scim"`
}

func (r *datasourceAuthSwgSamlServer) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_swg_saml_server"
}

func (r *datasourceAuthSwgSamlServer) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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
			"group_match": schema.StringAttribute{
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
			"scim_enabled": schema.BoolAttribute{
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
			"scim": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"scim_url": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"auth_method": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("token"),
						},
						Computed: true,
						Optional: true,
					},
					"token": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthBetween(1, 128),
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

func (r *datasourceAuthSwgSamlServer) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceAuthSwgSamlServer) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceAuthSwgSamlServerModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadAuthSwgSamlServer(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshAuthSwgSamlServer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceAuthSwgSamlServerModel) refreshAuthSwgSamlServer(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

	if v, ok := o["scim"]; ok {
		m.Scim = m.Scim.flattenAuthSwgSamlServerScim(ctx, v, &diags)
	}

	return diags
}

type datasourceAuthSwgSamlServerSpCertModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceAuthSwgSamlServerIdpCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceAuthSwgSamlServerScimModel struct {
	ScimUrl    types.String `tfsdk:"scim_url"`
	AuthMethod types.String `tfsdk:"auth_method"`
	Token      types.String `tfsdk:"token"`
}

func (m *datasourceAuthSwgSamlServerSpCertModel) flattenAuthSwgSamlServerSpCert(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthSwgSamlServerSpCertModel {
	if input == nil {
		return &datasourceAuthSwgSamlServerSpCertModel{}
	}
	if m == nil {
		m = &datasourceAuthSwgSamlServerSpCertModel{}
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

func (m *datasourceAuthSwgSamlServerIdpCertificateModel) flattenAuthSwgSamlServerIdpCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthSwgSamlServerIdpCertificateModel {
	if input == nil {
		return &datasourceAuthSwgSamlServerIdpCertificateModel{}
	}
	if m == nil {
		m = &datasourceAuthSwgSamlServerIdpCertificateModel{}
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

func (m *datasourceAuthSwgSamlServerScimModel) flattenAuthSwgSamlServerScim(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthSwgSamlServerScimModel {
	if input == nil {
		return &datasourceAuthSwgSamlServerScimModel{}
	}
	if m == nil {
		m = &datasourceAuthSwgSamlServerScimModel{}
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
