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
var _ datasource.DataSource = &datasourceInfraExtenders{}

func newDatasourceInfraExtenders() datasource.DataSource {
	return &datasourceInfraExtenders{}
}

type datasourceInfraExtenders struct {
	fortiClient *FortiClient
}

// datasourceInfraExtendersModel describes the datasource data model.
type datasourceInfraExtendersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Status     types.String `tfsdk:"status"`
}

func (r *datasourceInfraExtenders) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_extenders"
}

func (r *datasourceInfraExtenders) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("authorizing", "authorized", "deauthorizing", "deauthorized", "unknown", "failed"),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceInfraExtenders) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceInfraExtenders) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceInfraExtendersModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectInfraExtenders(ctx, "read", diags))

	read_output, err := c.ReadInfraExtenders(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshInfraExtenders(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceInfraExtendersModel) refreshInfraExtenders(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	return diags
}

func (data *datasourceInfraExtendersModel) getURLObjectInfraExtenders(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
