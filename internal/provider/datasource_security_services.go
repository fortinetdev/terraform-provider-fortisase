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
var _ datasource.DataSource = &datasourceSecurityServices{}

func newDatasourceSecurityServices() datasource.DataSource {
	return &datasourceSecurityServices{}
}

type datasourceSecurityServices struct {
	fortiClient *FortiClient
}

// datasourceSecurityServicesModel describes the datasource data model.
type datasourceSecurityServicesModel struct {
	PrimaryKey     types.String                                   `tfsdk:"primary_key"`
	Proxy          types.Bool                                     `tfsdk:"proxy"`
	Category       types.String                                   `tfsdk:"category"`
	Protocol       types.String                                   `tfsdk:"protocol"`
	ProtocolNumber types.Float64                                  `tfsdk:"protocol_number"`
	IcmpType       types.Float64                                  `tfsdk:"icmp_type"`
	UdpPortrange   []datasourceSecurityServicesUdpPortrangeModel  `tfsdk:"udp_portrange"`
	SctpPortrange  []datasourceSecurityServicesSctpPortrangeModel `tfsdk:"sctp_portrange"`
	TcpPortrange   []datasourceSecurityServicesTcpPortrangeModel  `tfsdk:"tcp_portrange"`
}

func (r *datasourceSecurityServices) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_services"
}

func (r *datasourceSecurityServices) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 79),
				},
				Required: true,
			},
			"proxy": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"category": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("Authentication", "Email", "File Access", "General", "Network Services", "Remote Access", "Tunneling", "Uncategorized", "VoIP, Messaging & Other Applications", "Web Access", "Web Proxy"),
				},
				Computed: true,
				Optional: true,
			},
			"protocol": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("TCP/UDP/SCTP"),
				},
				Computed: true,
				Optional: true,
			},
			"protocol_number": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(254),
				},
				Computed: true,
				Optional: true,
			},
			"icmp_type": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(4294967295),
				},
				Computed: true,
				Optional: true,
			},
			"udp_portrange": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"destination": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
							},
							Computed: true,
							Optional: true,
						},
						"source": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"sctp_portrange": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"destination": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
							},
							Computed: true,
							Optional: true,
						},
						"source": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"tcp_portrange": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"destination": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
							},
							Computed: true,
							Optional: true,
						},
						"source": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validator.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityServices) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityServices) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityServicesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityServices(ctx, "read", diags))

	read_output, err := c.ReadSecurityServices(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityServices(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityServicesModel) refreshSecurityServices(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["proxy"]; ok {
		m.Proxy = parseBoolValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = parseStringValue(v)
	}

	if v, ok := o["protocol"]; ok {
		m.Protocol = parseStringValue(v)
	}

	if v, ok := o["protocolNumber"]; ok {
		m.ProtocolNumber = parseFloat64Value(v)
	}

	if v, ok := o["icmpType"]; ok {
		m.IcmpType = parseFloat64Value(v)
	}

	if v, ok := o["udpPortrange"]; ok {
		m.UdpPortrange = m.flattenSecurityServicesUdpPortrangeList(ctx, v, &diags)
	}

	if v, ok := o["sctpPortrange"]; ok {
		m.SctpPortrange = m.flattenSecurityServicesSctpPortrangeList(ctx, v, &diags)
	}

	if v, ok := o["tcpPortrange"]; ok {
		m.TcpPortrange = m.flattenSecurityServicesTcpPortrangeList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityServicesModel) getURLObjectSecurityServices(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityServicesUdpPortrangeModel struct {
	Destination *datasourceSecurityServicesUdpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *datasourceSecurityServicesUdpPortrangeSourceModel      `tfsdk:"source"`
}

type datasourceSecurityServicesUdpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type datasourceSecurityServicesUdpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type datasourceSecurityServicesSctpPortrangeModel struct {
	Destination *datasourceSecurityServicesSctpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *datasourceSecurityServicesSctpPortrangeSourceModel      `tfsdk:"source"`
}

type datasourceSecurityServicesSctpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type datasourceSecurityServicesSctpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type datasourceSecurityServicesTcpPortrangeModel struct {
	Destination *datasourceSecurityServicesTcpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *datasourceSecurityServicesTcpPortrangeSourceModel      `tfsdk:"source"`
}

type datasourceSecurityServicesTcpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type datasourceSecurityServicesTcpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

func (m *datasourceSecurityServicesUdpPortrangeModel) flattenSecurityServicesUdpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServicesUdpPortrangeModel {
	if input == nil {
		return &datasourceSecurityServicesUdpPortrangeModel{}
	}
	if m == nil {
		m = &datasourceSecurityServicesUdpPortrangeModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["destination"]; ok {
		m.Destination = m.Destination.flattenSecurityServicesUdpPortrangeDestination(ctx, v, diags)
	}

	if v, ok := o["source"]; ok {
		m.Source = m.Source.flattenSecurityServicesUdpPortrangeSource(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityServicesModel) flattenSecurityServicesUdpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityServicesUdpPortrangeModel {
	if o == nil {
		return []datasourceSecurityServicesUdpPortrangeModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument udp_portrange is not type of []interface{}.", "")
		return []datasourceSecurityServicesUdpPortrangeModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityServicesUdpPortrangeModel{}
	}

	values := make([]datasourceSecurityServicesUdpPortrangeModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityServicesUdpPortrangeModel
		values[i] = *m.flattenSecurityServicesUdpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityServicesUdpPortrangeDestinationModel) flattenSecurityServicesUdpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServicesUdpPortrangeDestinationModel {
	if input == nil {
		return &datasourceSecurityServicesUdpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &datasourceSecurityServicesUdpPortrangeDestinationModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["low"]; ok {
		m.Low = parseFloat64Value(v)
	}

	if v, ok := o["high"]; ok {
		m.High = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceSecurityServicesUdpPortrangeSourceModel) flattenSecurityServicesUdpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServicesUdpPortrangeSourceModel {
	if input == nil {
		return &datasourceSecurityServicesUdpPortrangeSourceModel{}
	}
	if m == nil {
		m = &datasourceSecurityServicesUdpPortrangeSourceModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["low"]; ok {
		m.Low = parseFloat64Value(v)
	}

	if v, ok := o["high"]; ok {
		m.High = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceSecurityServicesSctpPortrangeModel) flattenSecurityServicesSctpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServicesSctpPortrangeModel {
	if input == nil {
		return &datasourceSecurityServicesSctpPortrangeModel{}
	}
	if m == nil {
		m = &datasourceSecurityServicesSctpPortrangeModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["destination"]; ok {
		m.Destination = m.Destination.flattenSecurityServicesSctpPortrangeDestination(ctx, v, diags)
	}

	if v, ok := o["source"]; ok {
		m.Source = m.Source.flattenSecurityServicesSctpPortrangeSource(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityServicesModel) flattenSecurityServicesSctpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityServicesSctpPortrangeModel {
	if o == nil {
		return []datasourceSecurityServicesSctpPortrangeModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sctp_portrange is not type of []interface{}.", "")
		return []datasourceSecurityServicesSctpPortrangeModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityServicesSctpPortrangeModel{}
	}

	values := make([]datasourceSecurityServicesSctpPortrangeModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityServicesSctpPortrangeModel
		values[i] = *m.flattenSecurityServicesSctpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityServicesSctpPortrangeDestinationModel) flattenSecurityServicesSctpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServicesSctpPortrangeDestinationModel {
	if input == nil {
		return &datasourceSecurityServicesSctpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &datasourceSecurityServicesSctpPortrangeDestinationModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["low"]; ok {
		m.Low = parseFloat64Value(v)
	}

	if v, ok := o["high"]; ok {
		m.High = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceSecurityServicesSctpPortrangeSourceModel) flattenSecurityServicesSctpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServicesSctpPortrangeSourceModel {
	if input == nil {
		return &datasourceSecurityServicesSctpPortrangeSourceModel{}
	}
	if m == nil {
		m = &datasourceSecurityServicesSctpPortrangeSourceModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["low"]; ok {
		m.Low = parseFloat64Value(v)
	}

	if v, ok := o["high"]; ok {
		m.High = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceSecurityServicesTcpPortrangeModel) flattenSecurityServicesTcpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServicesTcpPortrangeModel {
	if input == nil {
		return &datasourceSecurityServicesTcpPortrangeModel{}
	}
	if m == nil {
		m = &datasourceSecurityServicesTcpPortrangeModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["destination"]; ok {
		m.Destination = m.Destination.flattenSecurityServicesTcpPortrangeDestination(ctx, v, diags)
	}

	if v, ok := o["source"]; ok {
		m.Source = m.Source.flattenSecurityServicesTcpPortrangeSource(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityServicesModel) flattenSecurityServicesTcpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityServicesTcpPortrangeModel {
	if o == nil {
		return []datasourceSecurityServicesTcpPortrangeModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument tcp_portrange is not type of []interface{}.", "")
		return []datasourceSecurityServicesTcpPortrangeModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityServicesTcpPortrangeModel{}
	}

	values := make([]datasourceSecurityServicesTcpPortrangeModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityServicesTcpPortrangeModel
		values[i] = *m.flattenSecurityServicesTcpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityServicesTcpPortrangeDestinationModel) flattenSecurityServicesTcpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServicesTcpPortrangeDestinationModel {
	if input == nil {
		return &datasourceSecurityServicesTcpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &datasourceSecurityServicesTcpPortrangeDestinationModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["low"]; ok {
		m.Low = parseFloat64Value(v)
	}

	if v, ok := o["high"]; ok {
		m.High = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceSecurityServicesTcpPortrangeSourceModel) flattenSecurityServicesTcpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServicesTcpPortrangeSourceModel {
	if input == nil {
		return &datasourceSecurityServicesTcpPortrangeSourceModel{}
	}
	if m == nil {
		m = &datasourceSecurityServicesTcpPortrangeSourceModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["low"]; ok {
		m.Low = parseFloat64Value(v)
	}

	if v, ok := o["high"]; ok {
		m.High = parseFloat64Value(v)
	}

	return m
}
