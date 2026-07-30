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
var _ datasource.DataSource = &datasourceSecurityService{}

func newDatasourceSecurityService() datasource.DataSource {
	return &datasourceSecurityService{}
}

type datasourceSecurityService struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityServiceModel describes the datasource data model.
type datasourceSecurityServiceModel struct {
	PrimaryKey     types.String                                  `tfsdk:"primary_key"`
	Proxy          types.Bool                                    `tfsdk:"proxy"`
	Category       types.String                                  `tfsdk:"category"`
	Protocol       types.String                                  `tfsdk:"protocol"`
	ProtocolNumber types.Float64                                 `tfsdk:"protocol_number"`
	UdpPortrange   []datasourceSecurityServiceUdpPortrangeModel  `tfsdk:"udp_portrange"`
	SctpPortrange  []datasourceSecurityServiceSctpPortrangeModel `tfsdk:"sctp_portrange"`
	TcpPortrange   []datasourceSecurityServiceTcpPortrangeModel  `tfsdk:"tcp_portrange"`
	IcmpType       types.Float64                                 `tfsdk:"icmp_type"`
}

func (r *datasourceSecurityService) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_service"
}

func (r *datasourceSecurityService) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Service Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 79),
				},
				Required: true,
			},
			"proxy": schema.BoolAttribute{
				Computed: true,
			},
			"category": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("Authentication", "Email", "File Access", "General", "Network Services", "Remote Access", "Tunneling", "Uncategorized", "VoIP, Messaging & Other Applications", "Web Access", "Web Proxy"),
				},
				Computed: true,
			},
			"protocol": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("TCP/UDP/SCTP", "IP"),
				},
				Computed: true,
			},
			"protocol_number": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(254),
				},
				Computed: true,
			},
			"icmp_type": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(4294967295),
				},
				Computed: true,
			},
			"udp_portrange": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"destination": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
							},
							Computed: true,
						},
						"source": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"sctp_portrange": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"destination": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
							},
							Computed: true,
						},
						"source": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"tcp_portrange": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"destination": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
							},
							Computed: true,
						},
						"source": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"low": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceSecurityService) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_service"
}

func (r *datasourceSecurityService) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityServiceModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityService(ctx, "read", diags))

	read_output, err := c.ReadSecurityServices(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityService(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityServiceModel) refreshSecurityService(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
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

	if v, ok := o["udpPortrange"]; ok {
		m.UdpPortrange = m.flattenSecurityServiceUdpPortrangeList(ctx, v, &diags)
	}

	if v, ok := o["sctpPortrange"]; ok {
		m.SctpPortrange = m.flattenSecurityServiceSctpPortrangeList(ctx, v, &diags)
	}

	if v, ok := o["tcpPortrange"]; ok {
		m.TcpPortrange = m.flattenSecurityServiceTcpPortrangeList(ctx, v, &diags)
	}

	if v, ok := o["icmpType"]; ok {
		m.IcmpType = parseFloat64Value(v)
	}

	return diags
}

func (data *datasourceSecurityServiceModel) getURLObjectSecurityService(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityServiceUdpPortrangeModel struct {
	Destination *datasourceSecurityServiceUdpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *datasourceSecurityServiceUdpPortrangeSourceModel      `tfsdk:"source"`
}

type datasourceSecurityServiceUdpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type datasourceSecurityServiceUdpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type datasourceSecurityServiceSctpPortrangeModel struct {
	Destination *datasourceSecurityServiceSctpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *datasourceSecurityServiceSctpPortrangeSourceModel      `tfsdk:"source"`
}

type datasourceSecurityServiceSctpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type datasourceSecurityServiceSctpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type datasourceSecurityServiceTcpPortrangeModel struct {
	Destination *datasourceSecurityServiceTcpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *datasourceSecurityServiceTcpPortrangeSourceModel      `tfsdk:"source"`
}

type datasourceSecurityServiceTcpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type datasourceSecurityServiceTcpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

func (m *datasourceSecurityServiceUdpPortrangeModel) flattenSecurityServiceUdpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceUdpPortrangeModel {
	if input == nil {
		return &datasourceSecurityServiceUdpPortrangeModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceUdpPortrangeModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["destination"]; ok {
		m.Destination = m.Destination.flattenSecurityServiceUdpPortrangeDestination(ctx, v, diags)
	}

	if v, ok := o["source"]; ok {
		m.Source = m.Source.flattenSecurityServiceUdpPortrangeSource(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityServiceModel) flattenSecurityServiceUdpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityServiceUdpPortrangeModel {
	if o == nil {
		return []datasourceSecurityServiceUdpPortrangeModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument udp_portrange is not type of []interface{}.", "")
		return []datasourceSecurityServiceUdpPortrangeModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityServiceUdpPortrangeModel{}
	}

	values := make([]datasourceSecurityServiceUdpPortrangeModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityServiceUdpPortrangeModel
		if i < len(s.UdpPortrange) {
			m = s.UdpPortrange[i]
		}
		values[i] = *m.flattenSecurityServiceUdpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityServiceUdpPortrangeDestinationModel) flattenSecurityServiceUdpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceUdpPortrangeDestinationModel {
	if input == nil {
		return &datasourceSecurityServiceUdpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceUdpPortrangeDestinationModel{}
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

func (m *datasourceSecurityServiceUdpPortrangeSourceModel) flattenSecurityServiceUdpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceUdpPortrangeSourceModel {
	if input == nil {
		return &datasourceSecurityServiceUdpPortrangeSourceModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceUdpPortrangeSourceModel{}
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

func (m *datasourceSecurityServiceSctpPortrangeModel) flattenSecurityServiceSctpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceSctpPortrangeModel {
	if input == nil {
		return &datasourceSecurityServiceSctpPortrangeModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceSctpPortrangeModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["destination"]; ok {
		m.Destination = m.Destination.flattenSecurityServiceSctpPortrangeDestination(ctx, v, diags)
	}

	if v, ok := o["source"]; ok {
		m.Source = m.Source.flattenSecurityServiceSctpPortrangeSource(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityServiceModel) flattenSecurityServiceSctpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityServiceSctpPortrangeModel {
	if o == nil {
		return []datasourceSecurityServiceSctpPortrangeModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sctp_portrange is not type of []interface{}.", "")
		return []datasourceSecurityServiceSctpPortrangeModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityServiceSctpPortrangeModel{}
	}

	values := make([]datasourceSecurityServiceSctpPortrangeModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityServiceSctpPortrangeModel
		if i < len(s.SctpPortrange) {
			m = s.SctpPortrange[i]
		}
		values[i] = *m.flattenSecurityServiceSctpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityServiceSctpPortrangeDestinationModel) flattenSecurityServiceSctpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceSctpPortrangeDestinationModel {
	if input == nil {
		return &datasourceSecurityServiceSctpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceSctpPortrangeDestinationModel{}
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

func (m *datasourceSecurityServiceSctpPortrangeSourceModel) flattenSecurityServiceSctpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceSctpPortrangeSourceModel {
	if input == nil {
		return &datasourceSecurityServiceSctpPortrangeSourceModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceSctpPortrangeSourceModel{}
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

func (m *datasourceSecurityServiceTcpPortrangeModel) flattenSecurityServiceTcpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceTcpPortrangeModel {
	if input == nil {
		return &datasourceSecurityServiceTcpPortrangeModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceTcpPortrangeModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["destination"]; ok {
		m.Destination = m.Destination.flattenSecurityServiceTcpPortrangeDestination(ctx, v, diags)
	}

	if v, ok := o["source"]; ok {
		m.Source = m.Source.flattenSecurityServiceTcpPortrangeSource(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityServiceModel) flattenSecurityServiceTcpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityServiceTcpPortrangeModel {
	if o == nil {
		return []datasourceSecurityServiceTcpPortrangeModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument tcp_portrange is not type of []interface{}.", "")
		return []datasourceSecurityServiceTcpPortrangeModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityServiceTcpPortrangeModel{}
	}

	values := make([]datasourceSecurityServiceTcpPortrangeModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityServiceTcpPortrangeModel
		if i < len(s.TcpPortrange) {
			m = s.TcpPortrange[i]
		}
		values[i] = *m.flattenSecurityServiceTcpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityServiceTcpPortrangeDestinationModel) flattenSecurityServiceTcpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceTcpPortrangeDestinationModel {
	if input == nil {
		return &datasourceSecurityServiceTcpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceTcpPortrangeDestinationModel{}
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

func (m *datasourceSecurityServiceTcpPortrangeSourceModel) flattenSecurityServiceTcpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceTcpPortrangeSourceModel {
	if input == nil {
		return &datasourceSecurityServiceTcpPortrangeSourceModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceTcpPortrangeSourceModel{}
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
