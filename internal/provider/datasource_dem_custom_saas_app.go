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
var _ datasource.DataSource = &datasourceDemCustomSaasApp{}

func newDatasourceDemCustomSaasApp() datasource.DataSource {
	return &datasourceDemCustomSaasApp{}
}

type datasourceDemCustomSaasApp struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceDemCustomSaasAppModel describes the datasource data model.
type datasourceDemCustomSaasAppModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Alias      types.String `tfsdk:"alias"`
	Fqdn       types.String `tfsdk:"fqdn"`
}

func (r *datasourceDemCustomSaasApp) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dem_custom_saas_app"
}

func (r *datasourceDemCustomSaasApp) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DEM Custom SaaS Applications Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 253),
				},
				MarkdownDescription: "The primary key object of the DEM custom SaaS application. Can not be updated once created.\nLength between 1 and 253.",
				Required:            true,
			},
			"alias": schema.StringAttribute{
				Computed: true,
			},
			"fqdn": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 253),
				},
				MarkdownDescription: "The FQDN of the custom SaaS application.\nLength between 1 and 253.",
				Computed:            true,
			},
		},
	}
}

func (r *datasourceDemCustomSaasApp) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_dem_custom_saas_app"
}

func (r *datasourceDemCustomSaasApp) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceDemCustomSaasAppModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectDemCustomSaasApp(ctx, "read", diags))

	read_output, err := c.ReadDemCustomSaasApps(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshDemCustomSaasApp(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceDemCustomSaasAppModel) refreshDemCustomSaasApp(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["alias"]; ok {
		m.Alias = parseStringValue(v)
	}

	if v, ok := o["fqdn"]; ok {
		m.Fqdn = parseStringValue(v)
	}

	return diags
}

func (data *datasourceDemCustomSaasAppModel) getURLObjectDemCustomSaasApp(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
