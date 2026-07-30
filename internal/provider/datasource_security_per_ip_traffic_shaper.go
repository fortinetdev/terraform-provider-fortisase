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
var _ datasource.DataSource = &datasourceSecurityPerIpTrafficShaper{}

func newDatasourceSecurityPerIpTrafficShaper() datasource.DataSource {
	return &datasourceSecurityPerIpTrafficShaper{}
}

type datasourceSecurityPerIpTrafficShaper struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityPerIpTrafficShaperModel describes the datasource data model.
type datasourceSecurityPerIpTrafficShaperModel struct {
	PrimaryKey               types.String  `tfsdk:"primary_key"`
	MaximumBandwidth         types.Float64 `tfsdk:"maximum_bandwidth"`
	BandwidthUnit            types.String  `tfsdk:"bandwidth_unit"`
	MaxConcurrentSessions    types.Float64 `tfsdk:"max_concurrent_sessions"`
	MaxConcurrentTcpSessions types.Float64 `tfsdk:"max_concurrent_tcp_sessions"`
	MaxConcurrentUdpSessions types.Float64 `tfsdk:"max_concurrent_udp_sessions"`
}

func (r *datasourceSecurityPerIpTrafficShaper) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_per_ip_traffic_shaper"
}

func (r *datasourceSecurityPerIpTrafficShaper) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Per IP Traffic Shaper Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 64),
				},
				Required: true,
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
			"max_concurrent_sessions": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(2097000),
				},
				Computed: true,
			},
			"max_concurrent_tcp_sessions": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(2097000),
				},
				Computed: true,
			},
			"max_concurrent_udp_sessions": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(2097000),
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceSecurityPerIpTrafficShaper) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_per_ip_traffic_shaper"
}

func (r *datasourceSecurityPerIpTrafficShaper) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityPerIpTrafficShaperModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityPerIpTrafficShaper(ctx, "read", diags))

	read_output, err := c.ReadSecurityPerIpTrafficShaper(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityPerIpTrafficShaper(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityPerIpTrafficShaperModel) refreshSecurityPerIpTrafficShaper(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["maximumBandwidth"]; ok {
		m.MaximumBandwidth = parseFloat64Value(v)
	}

	if v, ok := o["bandwidthUnit"]; ok {
		m.BandwidthUnit = parseStringValue(v)
	}

	if v, ok := o["maxConcurrentSessions"]; ok {
		m.MaxConcurrentSessions = parseFloat64Value(v)
	}

	if v, ok := o["maxConcurrentTcpSessions"]; ok {
		m.MaxConcurrentTcpSessions = parseFloat64Value(v)
	}

	if v, ok := o["maxConcurrentUdpSessions"]; ok {
		m.MaxConcurrentUdpSessions = parseFloat64Value(v)
	}

	return diags
}

func (data *datasourceSecurityPerIpTrafficShaperModel) getURLObjectSecurityPerIpTrafficShaper(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
