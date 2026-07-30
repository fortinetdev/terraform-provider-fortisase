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
var _ datasource.DataSource = &datasourceEndpointMdmIntegration{}

func newDatasourceEndpointMdmIntegration() datasource.DataSource {
	return &datasourceEndpointMdmIntegration{}
}

type datasourceEndpointMdmIntegration struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointMdmIntegrationModel describes the datasource data model.
type datasourceEndpointMdmIntegrationModel struct {
	PrimaryKey        types.String                                            `tfsdk:"primary_key"`
	Enabled           types.Bool                                              `tfsdk:"enabled"`
	Vendor            types.String                                            `tfsdk:"vendor"`
	TenantId          types.String                                            `tfsdk:"tenant_id"`
	ClientId          types.String                                            `tfsdk:"client_id"`
	ClientSecret      types.String                                            `tfsdk:"client_secret"`
	Url               types.String                                            `tfsdk:"url"`
	SmartGroup        types.String                                            `tfsdk:"smart_group"`
	ApiKey            types.String                                            `tfsdk:"api_key"`
	Password          types.String                                            `tfsdk:"password"`
	Username          types.String                                            `tfsdk:"username"`
	SiteName          types.String                                            `tfsdk:"site_name"`
	AuthType          types.String                                            `tfsdk:"auth_type"`
	DeploymentType    types.String                                            `tfsdk:"deployment_type"`
	Region            types.String                                            `tfsdk:"region"`
	ClientCertificate *datasourceEndpointMdmIntegrationClientCertificateModel `tfsdk:"client_certificate"`
}

func (r *datasourceEndpointMdmIntegration) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_mdm_integration"
}

func (r *datasourceEndpointMdmIntegration) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MDM Integration Resource API V2 for FortiSASE. This resource is restricted to EMS version: 7.4.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("$sase-global"),
				},
				Required: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"vendor": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("intune", "airwatch", "jamf", "manageengine"),
				},
				Computed: true,
			},
			"tenant_id": schema.StringAttribute{
				Computed: true,
			},
			"client_id": schema.StringAttribute{
				Computed: true,
			},
			"client_secret": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"smart_group": schema.StringAttribute{
				Computed: true,
			},
			"api_key": schema.StringAttribute{
				Computed: true,
			},
			"password": schema.StringAttribute{
				Computed: true,
			},
			"username": schema.StringAttribute{
				Computed: true,
			},
			"site_name": schema.StringAttribute{
				Computed: true,
			},
			"auth_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("client_secret", "client_certificate", "basic", "oauth2"),
				},
				Computed: true,
			},
			"deployment_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("onprem", "cloud"),
				},
				Computed: true,
			},
			"region": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("northAmerica", "europe", "asiaPacific", "eu", "inec", "sg", "in", "uk", "us", "au", "jp", "ca", "sa"),
				},
				Computed: true,
			},
			"client_certificate": schema.SingleNestedAttribute{
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

func (r *datasourceEndpointMdmIntegration) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceEndpointMdmIntegration) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointMdmIntegrationModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadEndpointMdmIntegration(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
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

func (m *datasourceEndpointMdmIntegrationModel) refreshEndpointMdmIntegration(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
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

type datasourceEndpointMdmIntegrationClientCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceEndpointMdmIntegrationClientCertificateModel) flattenEndpointMdmIntegrationClientCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointMdmIntegrationClientCertificateModel {
	if input == nil {
		return &datasourceEndpointMdmIntegrationClientCertificateModel{}
	}
	if m == nil {
		m = &datasourceEndpointMdmIntegrationClientCertificateModel{}
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
