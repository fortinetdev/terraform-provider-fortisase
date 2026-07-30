// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
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
var _ resource.Resource = &resourceSecurityService{}
var _ resource.ResourceWithMoveState = &resourceSecurityService{}

func newResourceSecurityService() resource.Resource {
	return &resourceSecurityService{}
}

type resourceSecurityService struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityServiceModel describes the resource data model.
type resourceSecurityServiceModel struct {
	ID             types.String                                `tfsdk:"id"`
	PrimaryKey     types.String                                `tfsdk:"primary_key"`
	Proxy          types.Bool                                  `tfsdk:"proxy"`
	Category       types.String                                `tfsdk:"category"`
	Protocol       types.String                                `tfsdk:"protocol"`
	ProtocolNumber types.Float64                               `tfsdk:"protocol_number"`
	UdpPortrange   []resourceSecurityServiceUdpPortrangeModel  `tfsdk:"udp_portrange"`
	SctpPortrange  []resourceSecurityServiceSctpPortrangeModel `tfsdk:"sctp_portrange"`
	TcpPortrange   []resourceSecurityServiceTcpPortrangeModel  `tfsdk:"tcp_portrange"`
	IcmpType       types.Float64                               `tfsdk:"icmp_type"`
}

func (r *resourceSecurityService) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_service"
}

func (r *resourceSecurityService) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Service Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 79),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"proxy": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"category": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("Authentication", "Email", "File Access", "General", "Network Services", "Remote Access", "Tunneling", "Uncategorized", "VoIP, Messaging & Other Applications", "Web Access", "Web Proxy"),
				},
				Computed: true,
				Optional: true,
			},
			"protocol": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("TCP/UDP/SCTP", "IP"),
				},
				Computed: true,
				Optional: true,
			},
			"protocol_number": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(254),
				},
				Computed: true,
				Optional: true,
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
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
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
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
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
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
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
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
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
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
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
										float64validatorwarning.AtMost(65535),
									},
									Computed: true,
									Optional: true,
								},
								"high": schema.Float64Attribute{
									Validators: []validator.Float64{
										float64validatorwarning.AtMost(65535),
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

func (r *resourceSecurityService) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *resourceSecurityService) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_services" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityServiceModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityService) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityServices")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityServiceModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityService(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityService(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityServices(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	if responseMkey, ok := getCreateResponseMkey(output, "primaryKey"); ok {
		mkey = responseMkey
	}
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityService(ctx, "read", diags))

	read_output, err := c.ReadSecurityServices(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityService(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityService) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityServices")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityServiceModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityServiceModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityService(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityService(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityServices(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityService(ctx, "read", diags))

	read_output, err := c.ReadSecurityServices(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityService(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityService) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityServices")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityServiceModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityService(ctx, "delete", diags))

	output, err := c.DeleteSecurityServices(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityService) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityServiceModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityService(ctx, "read", diags))

	read_output, err := c.ReadSecurityServices(&input_model)
	if err != nil {
		if isNotFoundResponse(read_output) {
			resp.State.RemoveResource(ctx)
			return
		}
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
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

func (r *resourceSecurityService) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityServiceModel) refreshSecurityService(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceSecurityServiceModel) getCreateObjectSecurityService(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Proxy.IsNull() && !data.Proxy.IsUnknown() {
		result["proxy"] = data.Proxy.ValueBool()
	}

	if !data.Category.IsNull() && !data.Category.IsUnknown() {
		result["category"] = data.Category.ValueString()
	}

	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.ProtocolNumber.IsNull() && !data.ProtocolNumber.IsUnknown() {
		result["protocolNumber"] = data.ProtocolNumber.ValueFloat64()
	}

	if data.UdpPortrange != nil {
		result["udpPortrange"] = data.expandSecurityServiceUdpPortrangeList(ctx, data.UdpPortrange, diags)
	}

	if data.SctpPortrange != nil {
		result["sctpPortrange"] = data.expandSecurityServiceSctpPortrangeList(ctx, data.SctpPortrange, diags)
	}

	if data.TcpPortrange != nil {
		result["tcpPortrange"] = data.expandSecurityServiceTcpPortrangeList(ctx, data.TcpPortrange, diags)
	}

	return &result
}

func (data *resourceSecurityServiceModel) getUpdateObjectSecurityService(ctx context.Context, state resourceSecurityServiceModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Proxy.IsNull() && !data.Proxy.IsUnknown() {
		result["proxy"] = data.Proxy.ValueBool()
	}

	if !data.Category.IsNull() && !data.Category.IsUnknown() {
		result["category"] = data.Category.ValueString()
	}

	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.ProtocolNumber.IsNull() && !data.ProtocolNumber.IsUnknown() {
		result["protocolNumber"] = data.ProtocolNumber.ValueFloat64()
	}

	if data.UdpPortrange != nil {
		result["udpPortrange"] = data.expandSecurityServiceUdpPortrangeList(ctx, data.UdpPortrange, diags)
	}

	if data.SctpPortrange != nil {
		result["sctpPortrange"] = data.expandSecurityServiceSctpPortrangeList(ctx, data.SctpPortrange, diags)
	}

	if data.TcpPortrange != nil {
		result["tcpPortrange"] = data.expandSecurityServiceTcpPortrangeList(ctx, data.TcpPortrange, diags)
	}

	return &result
}

func (data *resourceSecurityServiceModel) getURLObjectSecurityService(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityServiceUdpPortrangeModel struct {
	Destination *resourceSecurityServiceUdpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *resourceSecurityServiceUdpPortrangeSourceModel      `tfsdk:"source"`
}

type resourceSecurityServiceUdpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type resourceSecurityServiceUdpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type resourceSecurityServiceSctpPortrangeModel struct {
	Destination *resourceSecurityServiceSctpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *resourceSecurityServiceSctpPortrangeSourceModel      `tfsdk:"source"`
}

type resourceSecurityServiceSctpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type resourceSecurityServiceSctpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type resourceSecurityServiceTcpPortrangeModel struct {
	Destination *resourceSecurityServiceTcpPortrangeDestinationModel `tfsdk:"destination"`
	Source      *resourceSecurityServiceTcpPortrangeSourceModel      `tfsdk:"source"`
}

type resourceSecurityServiceTcpPortrangeDestinationModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

type resourceSecurityServiceTcpPortrangeSourceModel struct {
	Low  types.Float64 `tfsdk:"low"`
	High types.Float64 `tfsdk:"high"`
}

func (m *resourceSecurityServiceUdpPortrangeModel) flattenSecurityServiceUdpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServiceUdpPortrangeModel {
	if input == nil {
		return &resourceSecurityServiceUdpPortrangeModel{}
	}
	if m == nil {
		m = &resourceSecurityServiceUdpPortrangeModel{}
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

func (s *resourceSecurityServiceModel) flattenSecurityServiceUdpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityServiceUdpPortrangeModel {
	if o == nil {
		return []resourceSecurityServiceUdpPortrangeModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument udp_portrange is not type of []interface{}.", "")
		return []resourceSecurityServiceUdpPortrangeModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityServiceUdpPortrangeModel{}
	}

	values := make([]resourceSecurityServiceUdpPortrangeModel, len(l))
	for i, ele := range l {
		var m resourceSecurityServiceUdpPortrangeModel
		if i < len(s.UdpPortrange) {
			m = s.UdpPortrange[i]
		}
		values[i] = *m.flattenSecurityServiceUdpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityServiceUdpPortrangeDestinationModel) flattenSecurityServiceUdpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServiceUdpPortrangeDestinationModel {
	if input == nil {
		return &resourceSecurityServiceUdpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &resourceSecurityServiceUdpPortrangeDestinationModel{}
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

func (m *resourceSecurityServiceUdpPortrangeSourceModel) flattenSecurityServiceUdpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServiceUdpPortrangeSourceModel {
	if input == nil {
		return &resourceSecurityServiceUdpPortrangeSourceModel{}
	}
	if m == nil {
		m = &resourceSecurityServiceUdpPortrangeSourceModel{}
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

func (m *resourceSecurityServiceSctpPortrangeModel) flattenSecurityServiceSctpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServiceSctpPortrangeModel {
	if input == nil {
		return &resourceSecurityServiceSctpPortrangeModel{}
	}
	if m == nil {
		m = &resourceSecurityServiceSctpPortrangeModel{}
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

func (s *resourceSecurityServiceModel) flattenSecurityServiceSctpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityServiceSctpPortrangeModel {
	if o == nil {
		return []resourceSecurityServiceSctpPortrangeModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument sctp_portrange is not type of []interface{}.", "")
		return []resourceSecurityServiceSctpPortrangeModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityServiceSctpPortrangeModel{}
	}

	values := make([]resourceSecurityServiceSctpPortrangeModel, len(l))
	for i, ele := range l {
		var m resourceSecurityServiceSctpPortrangeModel
		if i < len(s.SctpPortrange) {
			m = s.SctpPortrange[i]
		}
		values[i] = *m.flattenSecurityServiceSctpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityServiceSctpPortrangeDestinationModel) flattenSecurityServiceSctpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServiceSctpPortrangeDestinationModel {
	if input == nil {
		return &resourceSecurityServiceSctpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &resourceSecurityServiceSctpPortrangeDestinationModel{}
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

func (m *resourceSecurityServiceSctpPortrangeSourceModel) flattenSecurityServiceSctpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServiceSctpPortrangeSourceModel {
	if input == nil {
		return &resourceSecurityServiceSctpPortrangeSourceModel{}
	}
	if m == nil {
		m = &resourceSecurityServiceSctpPortrangeSourceModel{}
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

func (m *resourceSecurityServiceTcpPortrangeModel) flattenSecurityServiceTcpPortrange(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServiceTcpPortrangeModel {
	if input == nil {
		return &resourceSecurityServiceTcpPortrangeModel{}
	}
	if m == nil {
		m = &resourceSecurityServiceTcpPortrangeModel{}
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

func (s *resourceSecurityServiceModel) flattenSecurityServiceTcpPortrangeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityServiceTcpPortrangeModel {
	if o == nil {
		return []resourceSecurityServiceTcpPortrangeModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument tcp_portrange is not type of []interface{}.", "")
		return []resourceSecurityServiceTcpPortrangeModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityServiceTcpPortrangeModel{}
	}

	values := make([]resourceSecurityServiceTcpPortrangeModel, len(l))
	for i, ele := range l {
		var m resourceSecurityServiceTcpPortrangeModel
		if i < len(s.TcpPortrange) {
			m = s.TcpPortrange[i]
		}
		values[i] = *m.flattenSecurityServiceTcpPortrange(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityServiceTcpPortrangeDestinationModel) flattenSecurityServiceTcpPortrangeDestination(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServiceTcpPortrangeDestinationModel {
	if input == nil {
		return &resourceSecurityServiceTcpPortrangeDestinationModel{}
	}
	if m == nil {
		m = &resourceSecurityServiceTcpPortrangeDestinationModel{}
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

func (m *resourceSecurityServiceTcpPortrangeSourceModel) flattenSecurityServiceTcpPortrangeSource(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServiceTcpPortrangeSourceModel {
	if input == nil {
		return &resourceSecurityServiceTcpPortrangeSourceModel{}
	}
	if m == nil {
		m = &resourceSecurityServiceTcpPortrangeSourceModel{}
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

func (data *resourceSecurityServiceUdpPortrangeModel) expandSecurityServiceUdpPortrange(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Destination != nil && !isZeroStruct(*data.Destination) {
		result["destination"] = data.Destination.expandSecurityServiceUdpPortrangeDestination(ctx, diags)
	}

	if data.Source != nil && !isZeroStruct(*data.Source) {
		result["source"] = data.Source.expandSecurityServiceUdpPortrangeSource(ctx, diags)
	}

	return result
}

func (s *resourceSecurityServiceModel) expandSecurityServiceUdpPortrangeList(ctx context.Context, l []resourceSecurityServiceUdpPortrangeModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityServiceUdpPortrange(ctx, diags)
	}
	return result
}

func (data *resourceSecurityServiceUdpPortrangeDestinationModel) expandSecurityServiceUdpPortrangeDestination(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() && !data.Low.IsUnknown() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() && !data.High.IsUnknown() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityServiceUdpPortrangeSourceModel) expandSecurityServiceUdpPortrangeSource(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() && !data.Low.IsUnknown() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() && !data.High.IsUnknown() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityServiceSctpPortrangeModel) expandSecurityServiceSctpPortrange(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Destination != nil && !isZeroStruct(*data.Destination) {
		result["destination"] = data.Destination.expandSecurityServiceSctpPortrangeDestination(ctx, diags)
	}

	if data.Source != nil && !isZeroStruct(*data.Source) {
		result["source"] = data.Source.expandSecurityServiceSctpPortrangeSource(ctx, diags)
	}

	return result
}

func (s *resourceSecurityServiceModel) expandSecurityServiceSctpPortrangeList(ctx context.Context, l []resourceSecurityServiceSctpPortrangeModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityServiceSctpPortrange(ctx, diags)
	}
	return result
}

func (data *resourceSecurityServiceSctpPortrangeDestinationModel) expandSecurityServiceSctpPortrangeDestination(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() && !data.Low.IsUnknown() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() && !data.High.IsUnknown() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityServiceSctpPortrangeSourceModel) expandSecurityServiceSctpPortrangeSource(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() && !data.Low.IsUnknown() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() && !data.High.IsUnknown() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityServiceTcpPortrangeModel) expandSecurityServiceTcpPortrange(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Destination != nil && !isZeroStruct(*data.Destination) {
		result["destination"] = data.Destination.expandSecurityServiceTcpPortrangeDestination(ctx, diags)
	}

	if data.Source != nil && !isZeroStruct(*data.Source) {
		result["source"] = data.Source.expandSecurityServiceTcpPortrangeSource(ctx, diags)
	}

	return result
}

func (s *resourceSecurityServiceModel) expandSecurityServiceTcpPortrangeList(ctx context.Context, l []resourceSecurityServiceTcpPortrangeModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityServiceTcpPortrange(ctx, diags)
	}
	return result
}

func (data *resourceSecurityServiceTcpPortrangeDestinationModel) expandSecurityServiceTcpPortrangeDestination(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() && !data.Low.IsUnknown() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() && !data.High.IsUnknown() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}

func (data *resourceSecurityServiceTcpPortrangeSourceModel) expandSecurityServiceTcpPortrangeSource(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Low.IsNull() && !data.Low.IsUnknown() {
		result["low"] = data.Low.ValueFloat64()
	}

	if !data.High.IsNull() && !data.High.IsUnknown() {
		result["high"] = data.High.ValueFloat64()
	}

	return result
}
