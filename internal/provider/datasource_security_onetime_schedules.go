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
var _ datasource.DataSource = &datasourceSecurityOnetimeSchedules{}

func newDatasourceSecurityOnetimeSchedules() datasource.DataSource {
	return &datasourceSecurityOnetimeSchedules{}
}

type datasourceSecurityOnetimeSchedules struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityOnetimeSchedulesModel describes the datasource data model.
type datasourceSecurityOnetimeSchedulesModel struct {
	PrimaryKey     types.String  `tfsdk:"primary_key"`
	ExpirationDays types.Float64 `tfsdk:"expiration_days"`
	StartUtc       types.Float64 `tfsdk:"start_utc"`
	EndUtc         types.Float64 `tfsdk:"end_utc"`
}

func (r *datasourceSecurityOnetimeSchedules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_onetime_schedules"
}

func (r *datasourceSecurityOnetimeSchedules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 31),
				},
				Required: true,
			},
			"expiration_days": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(1, 100),
				},
				Computed: true,
				Optional: true,
			},
			"start_utc": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"end_utc": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityOnetimeSchedules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_onetime_schedules"
}

func (r *datasourceSecurityOnetimeSchedules) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityOnetimeSchedulesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityOnetimeSchedules(ctx, "read", diags))

	read_output, err := c.ReadSecurityOnetimeSchedules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityOnetimeSchedules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityOnetimeSchedulesModel) refreshSecurityOnetimeSchedules(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["expirationDays"]; ok {
		m.ExpirationDays = parseFloat64Value(v)
	}

	if v, ok := o["startUtc"]; ok {
		m.StartUtc = parseFloat64Value(v)
	}

	if v, ok := o["endUtc"]; ok {
		m.EndUtc = parseFloat64Value(v)
	}

	return diags
}

func (data *datasourceSecurityOnetimeSchedulesModel) getURLObjectSecurityOnetimeSchedules(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
