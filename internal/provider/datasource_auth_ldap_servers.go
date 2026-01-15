// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceAuthLdapServers{}

func newDatasourceAuthLdapServers() datasource.DataSource {
	return &datasourceAuthLdapServers{}
}

type datasourceAuthLdapServers struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceAuthLdapServersModel describes the datasource data model.
type datasourceAuthLdapServersModel struct {
	PrimaryKey                   types.String                               `tfsdk:"primary_key"`
	Server                       types.String                               `tfsdk:"server"`
	Port                         types.Float64                              `tfsdk:"port"`
	Cnid                         types.String                               `tfsdk:"cnid"`
	Dn                           types.String                               `tfsdk:"dn"`
	BindType                     types.String                               `tfsdk:"bind_type"`
	SecureConnection             types.Bool                                 `tfsdk:"secure_connection"`
	AdvancedGroupMatchingEnabled types.Bool                                 `tfsdk:"advanced_group_matching_enabled"`
	GroupMemberCheck             types.String                               `tfsdk:"group_member_check"`
	MemberAttribute              types.String                               `tfsdk:"member_attribute"`
	GroupFilter                  types.String                               `tfsdk:"group_filter"`
	GroupSearchBase              types.String                               `tfsdk:"group_search_base"`
	GroupObjectFilter            types.String                               `tfsdk:"group_object_filter"`
	ServerIdentityCheckEnabled   types.Bool                                 `tfsdk:"server_identity_check_enabled"`
	PasswordRenewalEnabled       types.Bool                                 `tfsdk:"password_renewal_enabled"`
	Certificate                  *datasourceAuthLdapServersCertificateModel `tfsdk:"certificate"`
	ClientCertAuthEnabled        types.Bool                                 `tfsdk:"client_cert_auth_enabled"`
	ClientCert                   *datasourceAuthLdapServersClientCertModel  `tfsdk:"client_cert"`
	Username                     types.String                               `tfsdk:"username"`
	Password                     types.String                               `tfsdk:"password"`
}

func (r *datasourceAuthLdapServers) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_ldap_servers"
}

func (r *datasourceAuthLdapServers) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 35),
				},
				Required: true,
			},
			"server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 63),
				},
				Computed: true,
				Optional: true,
			},
			"port": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(1, 65535),
				},
				Computed: true,
				Optional: true,
			},
			"cnid": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 20),
				},
				Computed: true,
				Optional: true,
			},
			"dn": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 511),
				},
				Computed: true,
				Optional: true,
			},
			"bind_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("simple", "anonymous", "regular"),
				},
				Computed: true,
				Optional: true,
			},
			"secure_connection": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"advanced_group_matching_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"group_member_check": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("user-attr", "group-object", "posix-group-object"),
				},
				Computed: true,
				Optional: true,
			},
			"member_attribute": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"group_filter": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(2047),
				},
				Computed: true,
				Optional: true,
			},
			"group_search_base": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(511),
				},
				Computed: true,
				Optional: true,
			},
			"group_object_filter": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(2047),
				},
				Computed: true,
				Optional: true,
			},
			"server_identity_check_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"password_renewal_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"client_cert_auth_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"username": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
				Computed: true,
				Optional: true,
			},
			"password": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
				Computed: true,
				Optional: true,
			},
			"certificate": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("system/certificate/ca-certificates"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"client_cert": schema.SingleNestedAttribute{
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
		},
	}
}

func (r *datasourceAuthLdapServers) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_auth_ldap_servers"
}

func (r *datasourceAuthLdapServers) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceAuthLdapServersModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthLdapServers(ctx, "read", diags))

	read_output, err := c.ReadAuthLdapServers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthLdapServers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceAuthLdapServersModel) refreshAuthLdapServers(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["server"]; ok {
		m.Server = parseStringValue(v)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	if v, ok := o["cnid"]; ok {
		m.Cnid = parseStringValue(v)
	}

	if v, ok := o["dn"]; ok {
		m.Dn = parseStringValue(v)
	}

	if v, ok := o["bindType"]; ok {
		m.BindType = parseStringValue(v)
	}

	if v, ok := o["secureConnection"]; ok {
		m.SecureConnection = parseBoolValue(v)
	}

	if v, ok := o["advancedGroupMatchingEnabled"]; ok {
		m.AdvancedGroupMatchingEnabled = parseBoolValue(v)
	}

	if v, ok := o["groupMemberCheck"]; ok {
		m.GroupMemberCheck = parseStringValue(v)
	}

	if v, ok := o["memberAttribute"]; ok {
		m.MemberAttribute = parseStringValue(v)
	}

	if v, ok := o["groupFilter"]; ok {
		m.GroupFilter = parseStringValue(v)
	}

	if v, ok := o["groupSearchBase"]; ok {
		m.GroupSearchBase = parseStringValue(v)
	}

	if v, ok := o["groupObjectFilter"]; ok {
		m.GroupObjectFilter = parseStringValue(v)
	}

	if v, ok := o["serverIdentityCheckEnabled"]; ok {
		m.ServerIdentityCheckEnabled = parseBoolValue(v)
	}

	if v, ok := o["passwordRenewalEnabled"]; ok {
		m.PasswordRenewalEnabled = parseBoolValue(v)
	}

	if v, ok := o["certificate"]; ok {
		m.Certificate = m.Certificate.flattenAuthLdapServersCertificate(ctx, v, &diags)
	}

	if v, ok := o["clientCertAuthEnabled"]; ok {
		m.ClientCertAuthEnabled = parseBoolValue(v)
	}

	if v, ok := o["clientCert"]; ok {
		m.ClientCert = m.ClientCert.flattenAuthLdapServersClientCert(ctx, v, &diags)
	}

	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	if v, ok := o["password"]; ok {
		m.Password = parseStringValue(v)
	}

	return diags
}

func (data *datasourceAuthLdapServersModel) getURLObjectAuthLdapServers(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceAuthLdapServersCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceAuthLdapServersClientCertModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceAuthLdapServersCertificateModel) flattenAuthLdapServersCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthLdapServersCertificateModel {
	if input == nil {
		return &datasourceAuthLdapServersCertificateModel{}
	}
	if m == nil {
		m = &datasourceAuthLdapServersCertificateModel{}
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

func (m *datasourceAuthLdapServersClientCertModel) flattenAuthLdapServersClientCert(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceAuthLdapServersClientCertModel {
	if input == nil {
		return &datasourceAuthLdapServersClientCertModel{}
	}
	if m == nil {
		m = &datasourceAuthLdapServersClientCertModel{}
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
