// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityVideoFilterFortiguardCategory{}

func newDatasourceSecurityVideoFilterFortiguardCategory() datasource.DataSource {
	return &datasourceSecurityVideoFilterFortiguardCategory{}
}

type datasourceSecurityVideoFilterFortiguardCategory struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityVideoFilterFortiguardCategoryModel describes the datasource data model.
type datasourceSecurityVideoFilterFortiguardCategoryModel struct {
	PrimaryKey types.String  `tfsdk:"primary_key"`
	Ftntid     types.Float64 `tfsdk:"ftntid"`
}

func (r *datasourceSecurityVideoFilterFortiguardCategory) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_video_filter_fortiguard_category"
}

func (r *datasourceSecurityVideoFilterFortiguardCategory) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Video Filter FortiGuard Category Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"ftntid": schema.Float64Attribute{
				Computed: true,
			},
		},
	}
}

func (r *datasourceSecurityVideoFilterFortiguardCategory) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_video_filter_fortiguard_category"
}

func (r *datasourceSecurityVideoFilterFortiguardCategory) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityVideoFilterFortiguardCategoryModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityVideoFilterFortiguardCategory(ctx, "read", diags))

	read_output, err := c.ReadSecurityVideoFilterFortiguardCategories(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityVideoFilterFortiguardCategory(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityVideoFilterFortiguardCategoryModel) refreshSecurityVideoFilterFortiguardCategory(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["id"]; ok {
		m.Ftntid = parseFloat64Value(v)
	}

	return diags
}

func (data *datasourceSecurityVideoFilterFortiguardCategoryModel) getURLObjectSecurityVideoFilterFortiguardCategory(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
