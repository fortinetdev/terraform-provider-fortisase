// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityTrafficShaper{}

func newDatasourceSecurityTrafficShaper() datasource.DataSource {
	return &datasourceSecurityTrafficShaper{}
}

type datasourceSecurityTrafficShaper struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityTrafficShaperModel describes the datasource data model.
type datasourceSecurityTrafficShaperModel struct {
	PrimaryKey          types.String  `tfsdk:"primary_key"`
	GuaranteedBandwidth types.Float64 `tfsdk:"guaranteed_bandwidth"`
	MaximumBandwidth    types.Float64 `tfsdk:"maximum_bandwidth"`
	BandwidthUnit       types.String  `tfsdk:"bandwidth_unit"`
	Priority            types.String  `tfsdk:"priority"`
	PerPolicy           types.String  `tfsdk:"per_policy"`
}

func (r *datasourceSecurityTrafficShaper) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_traffic_shaper"
}

func (r *datasourceSecurityTrafficShaper) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Shared Traffic Shaper Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 64),
				},
				Required: true,
			},
			"guaranteed_bandwidth": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(80000000),
				},
				Computed: true,
			},
			"maximum_bandwidth": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(80000000),
				},
				Computed: true,
			},
			"bandwidth_unit": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("kbps", "mbps", "gbps"),
				},
				Computed: true,
			},
			"priority": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("low", "medium", "high"),
				},
				Computed: true,
			},
			"per_policy": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("disable", "enable"),
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceSecurityTrafficShaper) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_traffic_shaper"
}

func (r *datasourceSecurityTrafficShaper) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityTrafficShaperModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityTrafficShaper(ctx, "read", diags))

	read_output, err := c.ReadSecurityTrafficShaper(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityTrafficShaper(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityTrafficShaperModel) refreshSecurityTrafficShaper(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["guaranteedBandwidth"]; ok {
		m.GuaranteedBandwidth = parseFloat64Value(v)
	}

	if v, ok := o["maximumBandwidth"]; ok {
		m.MaximumBandwidth = parseFloat64Value(v)
	}

	if v, ok := o["bandwidthUnit"]; ok {
		m.BandwidthUnit = parseStringValue(v)
	}

	if v, ok := o["priority"]; ok {
		m.Priority = parseStringValue(v)
	}

	if v, ok := o["perPolicy"]; ok {
		m.PerPolicy = parseStringValue(v)
	}

	return diags
}

func (data *datasourceSecurityTrafficShaperModel) getURLObjectSecurityTrafficShaper(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
