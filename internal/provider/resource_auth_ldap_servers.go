// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
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
var _ resource.Resource = &resourceAuthLdapServers{}

func newResourceAuthLdapServers() resource.Resource {
	return &resourceAuthLdapServers{}
}

type resourceAuthLdapServers struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceAuthLdapServersModel describes the resource data model.
type resourceAuthLdapServersModel struct {
	ID                           types.String                             `tfsdk:"id"`
	PrimaryKey                   types.String                             `tfsdk:"primary_key"`
	Server                       types.String                             `tfsdk:"server"`
	Port                         types.Float64                            `tfsdk:"port"`
	Cnid                         types.String                             `tfsdk:"cnid"`
	Dn                           types.String                             `tfsdk:"dn"`
	BindType                     types.String                             `tfsdk:"bind_type"`
	SecureConnection             types.Bool                               `tfsdk:"secure_connection"`
	AdvancedGroupMatchingEnabled types.Bool                               `tfsdk:"advanced_group_matching_enabled"`
	GroupMemberCheck             types.String                             `tfsdk:"group_member_check"`
	MemberAttribute              types.String                             `tfsdk:"member_attribute"`
	GroupFilter                  types.String                             `tfsdk:"group_filter"`
	GroupSearchBase              types.String                             `tfsdk:"group_search_base"`
	GroupObjectFilter            types.String                             `tfsdk:"group_object_filter"`
	ServerIdentityCheckEnabled   types.Bool                               `tfsdk:"server_identity_check_enabled"`
	PasswordRenewalEnabled       types.Bool                               `tfsdk:"password_renewal_enabled"`
	Certificate                  *resourceAuthLdapServersCertificateModel `tfsdk:"certificate"`
	ClientCertAuthEnabled        types.Bool                               `tfsdk:"client_cert_auth_enabled"`
	ClientCert                   *resourceAuthLdapServersClientCertModel  `tfsdk:"client_cert"`
	Username                     types.String                             `tfsdk:"username"`
	Password                     types.String                             `tfsdk:"password"`
}

func (r *resourceAuthLdapServers) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_ldap_servers"
}

func (r *resourceAuthLdapServers) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *resourceAuthLdapServers) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceAuthLdapServers) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthLdapServers")
	lock.Lock()
	defer lock.Unlock()
	var data resourceAuthLdapServersModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectAuthLdapServers(ctx, diags))
	input_model.URLParams = *(data.getURLObjectAuthLdapServers(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateAuthLdapServers(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectAuthLdapServers(ctx, "read", diags))

	read_output, err := c.ReadAuthLdapServers(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthLdapServers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthLdapServers) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("AuthLdapServers")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceAuthLdapServersModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceAuthLdapServersModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectAuthLdapServers(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectAuthLdapServers(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateAuthLdapServers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectAuthLdapServers(ctx, "read", diags))

	read_output, err := c.ReadAuthLdapServers(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshAuthLdapServers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceAuthLdapServers) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("AuthLdapServers")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceAuthLdapServersModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthLdapServers(ctx, "delete", diags))

	output, err := c.DeleteAuthLdapServers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceAuthLdapServers) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceAuthLdapServersModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthLdapServers(ctx, "read", diags))

	read_output, err := c.ReadAuthLdapServers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
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

func (r *resourceAuthLdapServers) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceAuthLdapServersModel) refreshAuthLdapServers(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceAuthLdapServersModel) getCreateObjectAuthLdapServers(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Server.IsNull() {
		result["server"] = data.Server.ValueString()
	}

	if !data.Port.IsNull() {
		result["port"] = data.Port.ValueFloat64()
	}

	if !data.Cnid.IsNull() {
		result["cnid"] = data.Cnid.ValueString()
	}

	if !data.Dn.IsNull() {
		result["dn"] = data.Dn.ValueString()
	}

	if !data.BindType.IsNull() {
		result["bindType"] = data.BindType.ValueString()
	}

	if !data.SecureConnection.IsNull() {
		result["secureConnection"] = data.SecureConnection.ValueBool()
	}

	if !data.AdvancedGroupMatchingEnabled.IsNull() {
		result["advancedGroupMatchingEnabled"] = data.AdvancedGroupMatchingEnabled.ValueBool()
	}

	if !data.GroupMemberCheck.IsNull() {
		result["groupMemberCheck"] = data.GroupMemberCheck.ValueString()
	}

	if !data.MemberAttribute.IsNull() {
		result["memberAttribute"] = data.MemberAttribute.ValueString()
	}

	if !data.GroupFilter.IsNull() {
		result["groupFilter"] = data.GroupFilter.ValueString()
	}

	if !data.GroupSearchBase.IsNull() {
		result["groupSearchBase"] = data.GroupSearchBase.ValueString()
	}

	if !data.GroupObjectFilter.IsNull() {
		result["groupObjectFilter"] = data.GroupObjectFilter.ValueString()
	}

	if !data.ServerIdentityCheckEnabled.IsNull() {
		result["serverIdentityCheckEnabled"] = data.ServerIdentityCheckEnabled.ValueBool()
	}

	if !data.PasswordRenewalEnabled.IsNull() {
		result["passwordRenewalEnabled"] = data.PasswordRenewalEnabled.ValueBool()
	}

	if data.Certificate != nil && !isZeroStruct(*data.Certificate) {
		result["certificate"] = data.Certificate.expandAuthLdapServersCertificate(ctx, diags)
	}

	if !data.ClientCertAuthEnabled.IsNull() {
		result["clientCertAuthEnabled"] = data.ClientCertAuthEnabled.ValueBool()
	}

	if data.ClientCert != nil && !isZeroStruct(*data.ClientCert) {
		result["clientCert"] = data.ClientCert.expandAuthLdapServersClientCert(ctx, diags)
	}

	if !data.Username.IsNull() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	return &result
}

func (data *resourceAuthLdapServersModel) getUpdateObjectAuthLdapServers(ctx context.Context, state resourceAuthLdapServersModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Server.IsNull() {
		result["server"] = data.Server.ValueString()
	}

	if !data.Port.IsNull() {
		result["port"] = data.Port.ValueFloat64()
	}

	if !data.Cnid.IsNull() {
		result["cnid"] = data.Cnid.ValueString()
	}

	if !data.Dn.IsNull() {
		result["dn"] = data.Dn.ValueString()
	}

	if !data.BindType.IsNull() {
		result["bindType"] = data.BindType.ValueString()
	}

	if !data.SecureConnection.IsNull() {
		result["secureConnection"] = data.SecureConnection.ValueBool()
	}

	if !data.AdvancedGroupMatchingEnabled.IsNull() {
		result["advancedGroupMatchingEnabled"] = data.AdvancedGroupMatchingEnabled.ValueBool()
	}

	if !data.GroupMemberCheck.IsNull() {
		result["groupMemberCheck"] = data.GroupMemberCheck.ValueString()
	}

	if !data.MemberAttribute.IsNull() {
		result["memberAttribute"] = data.MemberAttribute.ValueString()
	}

	if !data.GroupFilter.IsNull() {
		result["groupFilter"] = data.GroupFilter.ValueString()
	}

	if !data.GroupSearchBase.IsNull() {
		result["groupSearchBase"] = data.GroupSearchBase.ValueString()
	}

	if !data.GroupObjectFilter.IsNull() {
		result["groupObjectFilter"] = data.GroupObjectFilter.ValueString()
	}

	if !data.ServerIdentityCheckEnabled.IsNull() {
		result["serverIdentityCheckEnabled"] = data.ServerIdentityCheckEnabled.ValueBool()
	}

	if !data.PasswordRenewalEnabled.IsNull() {
		result["passwordRenewalEnabled"] = data.PasswordRenewalEnabled.ValueBool()
	}

	if data.Certificate != nil {
		result["certificate"] = data.Certificate.expandAuthLdapServersCertificate(ctx, diags)
	}

	if !data.ClientCertAuthEnabled.IsNull() {
		result["clientCertAuthEnabled"] = data.ClientCertAuthEnabled.ValueBool()
	}

	if data.ClientCert != nil {
		result["clientCert"] = data.ClientCert.expandAuthLdapServersClientCert(ctx, diags)
	}

	if !data.Username.IsNull() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	return &result
}

func (data *resourceAuthLdapServersModel) getURLObjectAuthLdapServers(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceAuthLdapServersCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceAuthLdapServersClientCertModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceAuthLdapServersCertificateModel) flattenAuthLdapServersCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthLdapServersCertificateModel {
	if input == nil {
		return &resourceAuthLdapServersCertificateModel{}
	}
	if m == nil {
		m = &resourceAuthLdapServersCertificateModel{}
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

func (m *resourceAuthLdapServersClientCertModel) flattenAuthLdapServersClientCert(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceAuthLdapServersClientCertModel {
	if input == nil {
		return &resourceAuthLdapServersClientCertModel{}
	}
	if m == nil {
		m = &resourceAuthLdapServersClientCertModel{}
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

func (data *resourceAuthLdapServersCertificateModel) expandAuthLdapServersCertificate(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceAuthLdapServersClientCertModel) expandAuthLdapServersClientCert(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
