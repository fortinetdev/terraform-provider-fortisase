// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSecurityServices{}

func newResourceSecurityServices() resource.Resource {
	return &resourceSecurityServices{}
}

type resourceSecurityServices struct {
	fortiClient *FortiClient
}

// resourceSecurityServicesModel describes the resource data model.
type resourceSecurityServicesModel struct {
	ID             types.String                                 `tfsdk:"id"`
	PrimaryKey     types.String                                 `tfsdk:"primary_key"`
	Proxy          types.Bool                                   `tfsdk:"proxy"`
	Category       types.String                                 `tfsdk:"category"`
	Protocol       types.String                                 `tfsdk:"protocol"`
	ProtocolNumber types.Float64                                `tfsdk:"protocol_number"`
	IcmpType       types.Float64                                `tfsdk:"icmp_type"`
	UdpPortrange   []resourceSecurityServicesUdpPortrangeModel  `tfsdk:"udp_portrange"`
	SctpPortrange  []resourceSecurityServicesSctpPortrangeModel `tfsdk:"sctp_portrange"`
	TcpPortrange   []resourceSecurityServicesTcpPortrangeModel  `tfsdk:"tcp_portrange"`
}

func (r *resourceSecurityServices) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_services"
}

func (r *resourceSecurityServices) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
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

func (r *resourceSecurityServices) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityServices) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityServicesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityServices(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityServices(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityServices(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityServices(ctx, "read", diags))

	read_output, err := c.ReadSecurityServices(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
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

func (r *resourceSecurityServices) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityServicesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityServicesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityServices(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityServices(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityServices(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityServices(ctx, "read", diags))

	read_output, err := c.ReadSecurityServices(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
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

func (r *resourceSecurityServices) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityServicesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityServices(ctx, "delete", diags))

	err := c.DeleteSecurityServices(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceSecurityServices) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityServicesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityServices(ctx, "read", diags))

	read_output, err := c.ReadSecurityServices(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
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

func (r *resourceSecurityServices) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityServicesModel) refreshSecurityServices(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceSecurityServicesModel) getCreateObjectSecurityServices(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Proxy.IsNull() {
		result["proxy"] = data.Proxy.ValueBool()
	}

	if !data.Category.IsNull() {
		result["category"] = data.Category.ValueString()
	}

	if !data.Protocol.IsNull() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.ProtocolNumber.IsNull() {
		result["protocolNumber"] = data.ProtocolNumber.ValueFloat64()
	}

	if !data.IcmpType.IsNull() {
		result["icmpType"] = data.IcmpType.ValueFloat64()
	}

	if len(data.UdpPortrange) > 0 {
		result["udpPortrange"] = data.expandSecurityServicesUdpPortrangeList(ctx, data.UdpPortrange, diags)
	}

	if len(data.SctpPortrange) > 0 {
		result["sctpPortrange"] = data.expandSecurityServicesSctpPortrangeList(ctx, data.SctpPortrange, diags)
	}

	if len(data.TcpPortrange) > 0 {
		result["tcpPortrange"] = data.expandSecurityServicesTcpPortrangeList(ctx, data.TcpPortrange, diags)
	}

	return &result
}

func (data *resourceSecurityServicesModel) getUpdateObjectSecurityServices(ctx context.Context, state resourceSecurityServicesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Proxy.IsNull() {
		result["proxy"] = data.Proxy.ValueBool()
	}

	if !data.Category.IsNull() {
		result["category"] = data.Category.ValueString()
	}

	if !data.Protocol.IsNull() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.ProtocolNumber.IsNull() && !data.ProtocolNumber.Equal(state.ProtocolNumber) {
		result["protocolNumber"] = data.ProtocolNumber.ValueFloat64()
	}

	if !data.IcmpType.IsNull() && !data.IcmpType.Equal(state.IcmpType) {
		result["icmpType"] = data.IcmpType.ValueFloat64()
	}

	if len(data.UdpPortrange) > 0 || !isSameStruct(data.UdpPortrange, state.UdpPortrange) {
		result["udpPortrange"] = data.expandSecurityServicesUdpPortrangeList(ctx, data.UdpPortrange, diags)
	}

	if len(data.SctpPortrange) > 0 || !isSameStruct(data.SctpPortrange, state.SctpPortrange) {
		result["sctpPortrange"] = data.expandSecurityServicesSctpPortrangeList(ctx, data.SctpPortrange, diags)
	}

	if len(data.TcpPortrange) > 0 || !isSameStruct(data.TcpPortrange, state.TcpPortrange) {
		result["tcpPortrange"] = data.expandSecurityServicesTcpPortrangeList(ctx, data.TcpPortrange, diags)
	}

	return &result
}

func (data *resourceSecurityServicesModel) getURLObjectSecurityServices(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityServicesUdpPortrangeModel struct {
	Destination *resourceSecurityServicesUdpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *resourceSecurityServicesUdpPortrangeSourceModel      `tfsdk:"source"`
}

type resourceSecurityServicesUdpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type resourceSecurityServicesUdpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type resourceSecurityServicesSctpPortrangeModel struct {
	Destination *resourceSecurityServicesSctpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *resourceSecurityServicesSctpPortrangeSourceModel      `tfsdk:"source"`
}

type resourceSecurityServicesSctpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type resourceSecurityServicesSctpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type resourceSecurityServicesTcpPortrangeModel struct {
	Destination *resourceSecurityServicesTcpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *resourceSecurityServicesTcpPortrangeSourceModel      `tfsdk:"source"`
}

type resourceSecurityServicesTcpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type resourceSecurityServicesTcpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

func (m *resourceSecurityServicesUdpPortrangeModel) flattenSecurityServicesUdpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServicesUdpPortrangeModel {
	if input == nil {
		return &resourceSecurityServicesUdpPortrangeModel{}
	}
	if m == nil {
		m = &resourceSecurityServicesUdpPortrangeModel{}
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

func (s *resourceSecurityServicesModel) flattenSecurityServicesUdpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityServicesUdpPortrangeModel {
	if o == nil {
		return []resourceSecurityServicesUdpPortrangeModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument udp_portrange is not type of []interface{}.", "")
		return []resourceSecurityServicesUdpPortrangeModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityServicesUdpPortrangeModel{}
	}

	values := make([]resourceSecurityServicesUdpPortrangeModel, len(l))
	for i, ele := range l {
		var m resourceSecurityServicesUdpPortrangeModel
		values[i] = *m.flattenSecurityServicesUdpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityServicesUdpPortrangeDestinationModel) flattenSecurityServicesUdpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServicesUdpPortrangeDestinationModel {
	if input == nil {
		return &resourceSecurityServicesUdpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &resourceSecurityServicesUdpPortrangeDestinationModel{}
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

func (m *resourceSecurityServicesUdpPortrangeSourceModel) flattenSecurityServicesUdpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServicesUdpPortrangeSourceModel {
	if input == nil {
		return &resourceSecurityServicesUdpPortrangeSourceModel{}
	}
	if m == nil {
		m = &resourceSecurityServicesUdpPortrangeSourceModel{}
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

func (m *resourceSecurityServicesSctpPortrangeModel) flattenSecurityServicesSctpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServicesSctpPortrangeModel {
	if input == nil {
		return &resourceSecurityServicesSctpPortrangeModel{}
	}
	if m == nil {
		m = &resourceSecurityServicesSctpPortrangeModel{}
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

func (s *resourceSecurityServicesModel) flattenSecurityServicesSctpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityServicesSctpPortrangeModel {
	if o == nil {
		return []resourceSecurityServicesSctpPortrangeModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument sctp_portrange is not type of []interface{}.", "")
		return []resourceSecurityServicesSctpPortrangeModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityServicesSctpPortrangeModel{}
	}

	values := make([]resourceSecurityServicesSctpPortrangeModel, len(l))
	for i, ele := range l {
		var m resourceSecurityServicesSctpPortrangeModel
		values[i] = *m.flattenSecurityServicesSctpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityServicesSctpPortrangeDestinationModel) flattenSecurityServicesSctpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServicesSctpPortrangeDestinationModel {
	if input == nil {
		return &resourceSecurityServicesSctpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &resourceSecurityServicesSctpPortrangeDestinationModel{}
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

func (m *resourceSecurityServicesSctpPortrangeSourceModel) flattenSecurityServicesSctpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServicesSctpPortrangeSourceModel {
	if input == nil {
		return &resourceSecurityServicesSctpPortrangeSourceModel{}
	}
	if m == nil {
		m = &resourceSecurityServicesSctpPortrangeSourceModel{}
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

func (m *resourceSecurityServicesTcpPortrangeModel) flattenSecurityServicesTcpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServicesTcpPortrangeModel {
	if input == nil {
		return &resourceSecurityServicesTcpPortrangeModel{}
	}
	if m == nil {
		m = &resourceSecurityServicesTcpPortrangeModel{}
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

func (s *resourceSecurityServicesModel) flattenSecurityServicesTcpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityServicesTcpPortrangeModel {
	if o == nil {
		return []resourceSecurityServicesTcpPortrangeModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument tcp_portrange is not type of []interface{}.", "")
		return []resourceSecurityServicesTcpPortrangeModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityServicesTcpPortrangeModel{}
	}

	values := make([]resourceSecurityServicesTcpPortrangeModel, len(l))
	for i, ele := range l {
		var m resourceSecurityServicesTcpPortrangeModel
		values[i] = *m.flattenSecurityServicesTcpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityServicesTcpPortrangeDestinationModel) flattenSecurityServicesTcpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServicesTcpPortrangeDestinationModel {
	if input == nil {
		return &resourceSecurityServicesTcpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &resourceSecurityServicesTcpPortrangeDestinationModel{}
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

func (m *resourceSecurityServicesTcpPortrangeSourceModel) flattenSecurityServicesTcpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServicesTcpPortrangeSourceModel {
	if input == nil {
		return &resourceSecurityServicesTcpPortrangeSourceModel{}
	}
	if m == nil {
		m = &resourceSecurityServicesTcpPortrangeSourceModel{}
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

func (data *resourceSecurityServicesUdpPortrangeModel) expandSecurityServicesUdpPortrange(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Destination != nil && !isZeroStruct(*data.Destination) {
		result["destination"] = data.Destination.expandSecurityServicesUdpPortrangeDestination(ctx, diags)
	}

	if data.Source != nil && !isZeroStruct(*data.Source) {
		result["source"] = data.Source.expandSecurityServicesUdpPortrangeSource(ctx, diags)
	}

	return result
}

func (s *resourceSecurityServicesModel) expandSecurityServicesUdpPortrangeList(ctx context.Context, l []resourceSecurityServicesUdpPortrangeModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityServicesUdpPortrange(ctx, diags)
	}
	return result
}

func (data *resourceSecurityServicesUdpPortrangeDestinationModel) expandSecurityServicesUdpPortrangeDestination(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityServicesUdpPortrangeSourceModel) expandSecurityServicesUdpPortrangeSource(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityServicesSctpPortrangeModel) expandSecurityServicesSctpPortrange(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Destination != nil && !isZeroStruct(*data.Destination) {
		result["destination"] = data.Destination.expandSecurityServicesSctpPortrangeDestination(ctx, diags)
	}

	if data.Source != nil && !isZeroStruct(*data.Source) {
		result["source"] = data.Source.expandSecurityServicesSctpPortrangeSource(ctx, diags)
	}

	return result
}

func (s *resourceSecurityServicesModel) expandSecurityServicesSctpPortrangeList(ctx context.Context, l []resourceSecurityServicesSctpPortrangeModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityServicesSctpPortrange(ctx, diags)
	}
	return result
}

func (data *resourceSecurityServicesSctpPortrangeDestinationModel) expandSecurityServicesSctpPortrangeDestination(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityServicesSctpPortrangeSourceModel) expandSecurityServicesSctpPortrangeSource(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityServicesTcpPortrangeModel) expandSecurityServicesTcpPortrange(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Destination != nil && !isZeroStruct(*data.Destination) {
		result["destination"] = data.Destination.expandSecurityServicesTcpPortrangeDestination(ctx, diags)
	}

	if data.Source != nil && !isZeroStruct(*data.Source) {
		result["source"] = data.Source.expandSecurityServicesTcpPortrangeSource(ctx, diags)
	}

	return result
}

func (s *resourceSecurityServicesModel) expandSecurityServicesTcpPortrangeList(ctx context.Context, l []resourceSecurityServicesTcpPortrangeModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityServicesTcpPortrange(ctx, diags)
	}
	return result
}

func (data *resourceSecurityServicesTcpPortrangeDestinationModel) expandSecurityServicesTcpPortrangeDestination(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityServicesTcpPortrangeSourceModel) expandSecurityServicesTcpPortrangeSource(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}
