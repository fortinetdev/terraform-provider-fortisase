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
var _ resource.Resource = &resourceEndpointMdmIntegration{}

func newResourceEndpointMdmIntegration() resource.Resource {
	return &resourceEndpointMdmIntegration{}
}

type resourceEndpointMdmIntegration struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointMdmIntegrationModel describes the resource data model.
type resourceEndpointMdmIntegrationModel struct {
	ID                types.String                                          `tfsdk:"id"`
	PrimaryKey        types.String                                          `tfsdk:"primary_key"`
	Enabled           types.Bool                                            `tfsdk:"enabled"`
	Vendor            types.String                                          `tfsdk:"vendor"`
	TenantId          types.String                                          `tfsdk:"tenant_id"`
	ClientId          types.String                                          `tfsdk:"client_id"`
	ClientSecret      types.String                                          `tfsdk:"client_secret"`
	Url               types.String                                          `tfsdk:"url"`
	SmartGroup        types.String                                          `tfsdk:"smart_group"`
	ApiKey            types.String                                          `tfsdk:"api_key"`
	Password          types.String                                          `tfsdk:"password"`
	Username          types.String                                          `tfsdk:"username"`
	SiteName          types.String                                          `tfsdk:"site_name"`
	AuthType          types.String                                          `tfsdk:"auth_type"`
	DeploymentType    types.String                                          `tfsdk:"deployment_type"`
	Region            types.String                                          `tfsdk:"region"`
	ClientCertificate *resourceEndpointMdmIntegrationClientCertificateModel `tfsdk:"client_certificate"`
}

func (r *resourceEndpointMdmIntegration) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_mdm_integration"
}

func (r *resourceEndpointMdmIntegration) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MDM Integration Resource API V2 for FortiSASE. This resource is restricted to EMS version: 7.4.",
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
			"vendor": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("intune", "airwatch", "jamf", "manageengine"),
				},
				Computed: true,
				Optional: true,
			},
			"tenant_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"client_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"client_secret": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"smart_group": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"api_key": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"password": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"username": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"site_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"auth_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("client_secret", "client_certificate", "basic", "oauth2"),
				},
				Computed: true,
				Optional: true,
			},
			"deployment_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("onprem", "cloud"),
				},
				Computed: true,
				Optional: true,
			},
			"region": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("northAmerica", "europe", "asiaPacific", "eu", "inec", "sg", "in", "uk", "us", "au", "jp", "ca", "sa"),
				},
				Computed: true,
				Optional: true,
			},
			"client_certificate": schema.SingleNestedAttribute{
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

func (r *resourceEndpointMdmIntegration) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	support_versions := map[string][]string{
		"EMS": {"7.4"},
	}
	ok, err := checkVersionMatch(client.Client, support_versions)
	if !ok {
		resp.Diagnostics.AddError(
			"FortiSASE EMS version do not support this resource.",
			fmt.Sprintf("%v", err),
		)

		return
	}

	r.fortiClient = client
	r.resourceName = "fortisase_endpoint_mdm_integration"
}

func (r *resourceEndpointMdmIntegration) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointMdmIntegration")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointMdmIntegrationModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointMdmIntegration(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateEndpointMdmIntegration(&input_model)
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

	read_output, err := c.ReadEndpointMdmIntegration(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointMdmIntegration(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointMdmIntegration) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointMdmIntegration")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointMdmIntegrationModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointMdmIntegrationModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointMdmIntegration(ctx, state, diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointMdmIntegration(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output, err := c.ReadEndpointMdmIntegration(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointMdmIntegration(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointMdmIntegration) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointMdmIntegration) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointMdmIntegrationModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadEndpointMdmIntegration(&input_model)
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

	diags.Append(data.refreshEndpointMdmIntegration(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointMdmIntegration) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointMdmIntegrationModel) refreshEndpointMdmIntegration(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

	if v, ok := o["vendor"]; ok {
		m.Vendor = parseStringValue(v)
	}

	if v, ok := o["tenantId"]; ok {
		m.TenantId = parseStringValue(v)
	}

	if v, ok := o["clientId"]; ok {
		m.ClientId = parseStringValue(v)
	}

	if v, ok := o["clientSecret"]; ok {
		m.ClientSecret = parseStringValue(v)
	}

	if v, ok := o["url"]; ok {
		m.Url = parseStringValue(v)
	}

	if v, ok := o["smartGroup"]; ok {
		m.SmartGroup = parseStringValue(v)
	}

	if v, ok := o["apiKey"]; ok {
		m.ApiKey = parseStringValue(v)
	}

	if v, ok := o["password"]; ok {
		m.Password = parseStringValue(v)
	}

	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	if v, ok := o["siteName"]; ok {
		m.SiteName = parseStringValue(v)
	}

	if v, ok := o["authType"]; ok {
		m.AuthType = parseStringValue(v)
	}

	if v, ok := o["deploymentType"]; ok {
		m.DeploymentType = parseStringValue(v)
	}

	if v, ok := o["region"]; ok {
		m.Region = parseStringValue(v)
	}

	if v, ok := o["clientCertificate"]; ok {
		m.ClientCertificate = m.ClientCertificate.flattenEndpointMdmIntegrationClientCertificate(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointMdmIntegrationModel) getCreateObjectEndpointMdmIntegration(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.Vendor.IsNull() && !data.Vendor.IsUnknown() {
		result["vendor"] = data.Vendor.ValueString()
	}

	if !data.TenantId.IsNull() && !data.TenantId.IsUnknown() {
		result["tenantId"] = data.TenantId.ValueString()
	}

	if !data.ClientId.IsNull() && !data.ClientId.IsUnknown() {
		result["clientId"] = data.ClientId.ValueString()
	}

	if !data.ClientSecret.IsNull() && !data.ClientSecret.IsUnknown() {
		result["clientSecret"] = data.ClientSecret.ValueString()
	}

	if !data.Url.IsNull() && !data.Url.IsUnknown() {
		result["url"] = data.Url.ValueString()
	}

	if !data.SmartGroup.IsNull() && !data.SmartGroup.IsUnknown() {
		result["smartGroup"] = data.SmartGroup.ValueString()
	}

	if !data.ApiKey.IsNull() && !data.ApiKey.IsUnknown() {
		result["apiKey"] = data.ApiKey.ValueString()
	}

	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		result["password"] = data.Password.ValueString()
	}

	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		result["username"] = data.Username.ValueString()
	}

	if !data.SiteName.IsNull() && !data.SiteName.IsUnknown() {
		result["siteName"] = data.SiteName.ValueString()
	}

	if !data.AuthType.IsNull() && !data.AuthType.IsUnknown() {
		result["authType"] = data.AuthType.ValueString()
	}

	if !data.DeploymentType.IsNull() && !data.DeploymentType.IsUnknown() {
		result["deploymentType"] = data.DeploymentType.ValueString()
	}

	if !data.Region.IsNull() && !data.Region.IsUnknown() {
		result["region"] = data.Region.ValueString()
	}

	result["clientCertificate"] = nil
	if data.ClientCertificate != nil && !isZeroStruct(*data.ClientCertificate) {
		result["clientCertificate"] = data.ClientCertificate.expandEndpointMdmIntegrationClientCertificate(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointMdmIntegrationModel) getUpdateObjectEndpointMdmIntegration(ctx context.Context, state resourceEndpointMdmIntegrationModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Enabled.IsNull() && !data.Enabled.IsUnknown() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	if !data.Vendor.IsNull() && !data.Vendor.IsUnknown() {
		result["vendor"] = data.Vendor.ValueString()
	}

	if !data.TenantId.IsNull() && !data.TenantId.IsUnknown() {
		result["tenantId"] = data.TenantId.ValueString()
	}

	if !data.ClientId.IsNull() && !data.ClientId.IsUnknown() {
		result["clientId"] = data.ClientId.ValueString()
	}

	if !data.ClientSecret.IsNull() && !data.ClientSecret.IsUnknown() {
		result["clientSecret"] = data.ClientSecret.ValueString()
	}

	if !data.Url.IsNull() && !data.Url.IsUnknown() {
		result["url"] = data.Url.ValueString()
	}

	if !data.SmartGroup.IsNull() && !data.SmartGroup.IsUnknown() {
		result["smartGroup"] = data.SmartGroup.ValueString()
	}

	if !data.ApiKey.IsNull() && !data.ApiKey.IsUnknown() {
		result["apiKey"] = data.ApiKey.ValueString()
	}

	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		result["password"] = data.Password.ValueString()
	}

	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		result["username"] = data.Username.ValueString()
	}

	if !data.SiteName.IsNull() && !data.SiteName.IsUnknown() {
		result["siteName"] = data.SiteName.ValueString()
	}

	if !data.AuthType.IsNull() && !data.AuthType.IsUnknown() {
		result["authType"] = data.AuthType.ValueString()
	}

	if !data.DeploymentType.IsNull() && !data.DeploymentType.IsUnknown() {
		result["deploymentType"] = data.DeploymentType.ValueString()
	}

	if !data.Region.IsNull() && !data.Region.IsUnknown() {
		result["region"] = data.Region.ValueString()
	}

	result["clientCertificate"] = nil
	if data.ClientCertificate != nil && !isZeroStruct(*data.ClientCertificate) {
		result["clientCertificate"] = data.ClientCertificate.expandEndpointMdmIntegrationClientCertificate(ctx, diags)
	}

	return &result
}

type resourceEndpointMdmIntegrationClientCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceEndpointMdmIntegrationClientCertificateModel) flattenEndpointMdmIntegrationClientCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointMdmIntegrationClientCertificateModel {
	if input == nil {
		return &resourceEndpointMdmIntegrationClientCertificateModel{}
	}
	if m == nil {
		m = &resourceEndpointMdmIntegrationClientCertificateModel{}
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

func (data *resourceEndpointMdmIntegrationClientCertificateModel) expandEndpointMdmIntegrationClientCertificate(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
