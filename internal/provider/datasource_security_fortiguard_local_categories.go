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
var _ datasource.DataSource = &datasourceSecurityFortiguardLocalCategories{}

func newDatasourceSecurityFortiguardLocalCategories() datasource.DataSource {
	return &datasourceSecurityFortiguardLocalCategories{}
}

type datasourceSecurityFortiguardLocalCategories struct {
	fortiClient *FortiClient
}

// datasourceSecurityFortiguardLocalCategoriesModel describes the datasource data model.
type datasourceSecurityFortiguardLocalCategoriesModel struct {
	PrimaryKey   types.String `tfsdk:"primary_key"`
	ThreatWeight types.String `tfsdk:"threat_weight"`
	Urls         types.Set    `tfsdk:"urls"`
}

func (r *datasourceSecurityFortiguardLocalCategories) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_fortiguard_local_categories"
}

func (r *datasourceSecurityFortiguardLocalCategories) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 79),
				},
				Required: true,
			},
			"threat_weight": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("none", "low", "medium", "high", "critical"),
				},
				Computed: true,
				Optional: true,
			},
			"urls": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *datasourceSecurityFortiguardLocalCategories) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityFortiguardLocalCategories) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityFortiguardLocalCategoriesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityFortiguardLocalCategories(ctx, "read", diags))

	read_output, err := c.ReadSecurityFortiguardLocalCategories(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityFortiguardLocalCategories(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityFortiguardLocalCategoriesModel) refreshSecurityFortiguardLocalCategories(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["threatWeight"]; ok {
		m.ThreatWeight = parseStringValue(v)
	}

	if v, ok := o["urls"]; ok {
		m.Urls = parseSetValue(ctx, v, types.StringType)
	}

	return diags
}

func (data *datasourceSecurityFortiguardLocalCategoriesModel) getURLObjectSecurityFortiguardLocalCategories(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
