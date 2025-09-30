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
var _ datasource.DataSource = &datasourceEndpointsDonut{}

func newDatasourceEndpointsDonut() datasource.DataSource {
	return &datasourceEndpointsDonut{}
}

type datasourceEndpointsDonut struct {
	fortiClient *FortiClient
}

// datasourceEndpointsDonutModel describes the datasource data model.
type datasourceEndpointsDonutModel struct {
	Token     types.String  `tfsdk:"token"`
	Value     types.Float64 `tfsdk:"value"`
	Name      types.String  `tfsdk:"name"`
	DonutType types.String  `tfsdk:"donut_type"`
}

func (r *datasourceEndpointsDonut) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_donut"
}

func (r *datasourceEndpointsDonut) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"token": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"value": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"donut_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("connection", "management"),
				},
				Required: true,
			},
		},
	}
}

func (r *datasourceEndpointsDonut) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceEndpointsDonut) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointsDonutModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.DonutType.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointsDonut(ctx, "read", diags))

	read_output, err := c.ReadEndpointsDonut(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointsDonut(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointsDonutModel) refreshEndpointsDonut(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["token"]; ok {
		m.Token = parseStringValue(v)
	}

	if v, ok := o["value"]; ok {
		m.Value = parseFloat64Value(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	return diags
}

func (data *datasourceEndpointsDonutModel) getURLObjectEndpointsDonut(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.DonutType.IsNull() {
		result["donutType"] = data.DonutType.ValueString()
	}

	return &result
}
