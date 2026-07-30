// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
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
var _ resource.Resource = &resourceEndpointOnNetRule{}
var _ resource.ResourceWithMoveState = &resourceEndpointOnNetRule{}

func newResourceEndpointOnNetRule() resource.Resource {
	return &resourceEndpointOnNetRule{}
}

type resourceEndpointOnNetRule struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointOnNetRuleModel describes the resource data model.
type resourceEndpointOnNetRuleModel struct {
	ID              types.String                                    `tfsdk:"id"`
	PrimaryKey      types.String                                    `tfsdk:"primary_key"`
	PublicIp        types.String                                    `tfsdk:"public_ip"`
	DhcpServerIp    types.String                                    `tfsdk:"dhcp_server_ip"`
	DhcpServerMac   types.String                                    `tfsdk:"dhcp_server_mac"`
	DhcpServerCode  types.String                                    `tfsdk:"dhcp_server_code"`
	DnsServerIp     types.String                                    `tfsdk:"dns_server_ip"`
	PingServer      types.String                                    `tfsdk:"ping_server"`
	LocalIp         types.String                                    `tfsdk:"local_ip"`
	GatewayMac      types.String                                    `tfsdk:"gateway_mac"`
	WebRequestHttp  types.String                                    `tfsdk:"web_request_http"`
	WebRequestHttps []resourceEndpointOnNetRuleWebRequestHttpsModel `tfsdk:"web_request_https"`
	DnsRequest      []resourceEndpointOnNetRuleDnsRequestModel      `tfsdk:"dns_request"`
}

func (r *resourceEndpointOnNetRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_on_net_rule"
}

func (r *resourceEndpointOnNetRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Endpoint on net rule",
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
					stringvalidatorwarning.LengthBetween(1, 128),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"public_ip": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"dhcp_server_ip": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"dhcp_server_mac": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"dhcp_server_code": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"dns_server_ip": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ping_server": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"local_ip": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"gateway_mac": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"web_request_http": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"web_request_https": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"hostname": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"dns_request": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"hostname": schema.StringAttribute{
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

func (r *resourceEndpointOnNetRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_on_net_rule"
}
func (r *resourceEndpointOnNetRule) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_endpoint_on_net_rules" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceEndpointOnNetRuleModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceEndpointOnNetRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointOnNetRuleModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointOnNetRule(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointOnNetRule(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateEndpointOnNetRules(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectEndpointOnNetRule(ctx, "read", diags))

	read_output, err := c.ReadEndpointOnNetRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointOnNetRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointOnNetRule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointOnNetRuleModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointOnNetRuleModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointOnNetRule(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointOnNetRule(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointOnNetRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointOnNetRule(ctx, "read", diags))

	read_output, err := c.ReadEndpointOnNetRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointOnNetRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointOnNetRule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointOnNetRuleModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointOnNetRule(ctx, "delete", diags))

	output, err := c.DeleteEndpointOnNetRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceEndpointOnNetRule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointOnNetRuleModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointOnNetRule(ctx, "read", diags))

	read_output, err := c.ReadEndpointOnNetRules(&input_model)
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

	diags.Append(data.refreshEndpointOnNetRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointOnNetRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceEndpointOnNetRuleModel) refreshEndpointOnNetRule(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["publicIp"]; ok {
		m.PublicIp = parseStringValue(v)
	}

	if v, ok := o["dhcpServerIp"]; ok {
		m.DhcpServerIp = parseStringValue(v)
	}

	if v, ok := o["dhcpServerMac"]; ok {
		m.DhcpServerMac = parseStringValue(v)
	}

	if v, ok := o["dhcpServerCode"]; ok {
		m.DhcpServerCode = parseStringValue(v)
	}

	if v, ok := o["dnsServerIp"]; ok {
		m.DnsServerIp = parseStringValue(v)
	}

	if v, ok := o["pingServer"]; ok {
		m.PingServer = parseStringValue(v)
	}

	if v, ok := o["localIp"]; ok {
		m.LocalIp = parseStringValue(v)
	}

	if v, ok := o["gatewayMac"]; ok {
		m.GatewayMac = parseStringValue(v)
	}

	if v, ok := o["webRequestHttp"]; ok {
		m.WebRequestHttp = parseStringValue(v)
	}

	if v, ok := o["webRequestHttps"]; ok {
		m.WebRequestHttps = m.flattenEndpointOnNetRuleWebRequestHttpsList(ctx, v, &diags)
	}

	if v, ok := o["dnsRequest"]; ok {
		m.DnsRequest = m.flattenEndpointOnNetRuleDnsRequestList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointOnNetRuleModel) getCreateObjectEndpointOnNetRule(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.PublicIp.IsNull() && !data.PublicIp.IsUnknown() {
		result["publicIp"] = data.PublicIp.ValueString()
	}

	if !data.DhcpServerIp.IsNull() && !data.DhcpServerIp.IsUnknown() {
		result["dhcpServerIp"] = data.DhcpServerIp.ValueString()
	}

	if !data.DhcpServerMac.IsNull() && !data.DhcpServerMac.IsUnknown() {
		result["dhcpServerMac"] = data.DhcpServerMac.ValueString()
	}

	if !data.DhcpServerCode.IsNull() && !data.DhcpServerCode.IsUnknown() {
		result["dhcpServerCode"] = data.DhcpServerCode.ValueString()
	}

	if !data.DnsServerIp.IsNull() && !data.DnsServerIp.IsUnknown() {
		result["dnsServerIp"] = data.DnsServerIp.ValueString()
	}

	if !data.PingServer.IsNull() && !data.PingServer.IsUnknown() {
		result["pingServer"] = data.PingServer.ValueString()
	}

	if !data.LocalIp.IsNull() && !data.LocalIp.IsUnknown() {
		result["localIp"] = data.LocalIp.ValueString()
	}

	if !data.GatewayMac.IsNull() && !data.GatewayMac.IsUnknown() {
		result["gatewayMac"] = data.GatewayMac.ValueString()
	}

	if !data.WebRequestHttp.IsNull() && !data.WebRequestHttp.IsUnknown() {
		result["webRequestHttp"] = data.WebRequestHttp.ValueString()
	}

	if data.WebRequestHttps != nil {
		result["webRequestHttps"] = data.expandEndpointOnNetRuleWebRequestHttpsList(ctx, data.WebRequestHttps, diags)
	}

	if data.DnsRequest != nil {
		result["dnsRequest"] = data.expandEndpointOnNetRuleDnsRequestList(ctx, data.DnsRequest, diags)
	}

	return &result
}

func (data *resourceEndpointOnNetRuleModel) getUpdateObjectEndpointOnNetRule(ctx context.Context, state resourceEndpointOnNetRuleModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.PublicIp.IsNull() && !data.PublicIp.IsUnknown() {
		result["publicIp"] = data.PublicIp.ValueString()
	}

	if !data.DhcpServerIp.IsNull() && !data.DhcpServerIp.IsUnknown() {
		result["dhcpServerIp"] = data.DhcpServerIp.ValueString()
	}

	if !data.DhcpServerMac.IsNull() && !data.DhcpServerMac.IsUnknown() {
		result["dhcpServerMac"] = data.DhcpServerMac.ValueString()
	}

	if !data.DhcpServerCode.IsNull() && !data.DhcpServerCode.IsUnknown() {
		result["dhcpServerCode"] = data.DhcpServerCode.ValueString()
	}

	if !data.DnsServerIp.IsNull() && !data.DnsServerIp.IsUnknown() {
		result["dnsServerIp"] = data.DnsServerIp.ValueString()
	}

	if !data.PingServer.IsNull() && !data.PingServer.IsUnknown() {
		result["pingServer"] = data.PingServer.ValueString()
	}

	if !data.LocalIp.IsNull() && !data.LocalIp.IsUnknown() {
		result["localIp"] = data.LocalIp.ValueString()
	}

	if !data.GatewayMac.IsNull() && !data.GatewayMac.IsUnknown() {
		result["gatewayMac"] = data.GatewayMac.ValueString()
	}

	if !data.WebRequestHttp.IsNull() && !data.WebRequestHttp.IsUnknown() {
		result["webRequestHttp"] = data.WebRequestHttp.ValueString()
	}

	if data.WebRequestHttps != nil {
		result["webRequestHttps"] = data.expandEndpointOnNetRuleWebRequestHttpsList(ctx, data.WebRequestHttps, diags)
	}

	if data.DnsRequest != nil {
		result["dnsRequest"] = data.expandEndpointOnNetRuleDnsRequestList(ctx, data.DnsRequest, diags)
	}

	return &result
}

func (data *resourceEndpointOnNetRuleModel) getURLObjectEndpointOnNetRule(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointOnNetRuleWebRequestHttpsModel struct {
	Ip       types.String `tfsdk:"ip"`
	Hostname types.String `tfsdk:"hostname"`
}

type resourceEndpointOnNetRuleDnsRequestModel struct {
	Ip       types.String `tfsdk:"ip"`
	Hostname types.String `tfsdk:"hostname"`
}

func (m *resourceEndpointOnNetRuleWebRequestHttpsModel) flattenEndpointOnNetRuleWebRequestHttps(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointOnNetRuleWebRequestHttpsModel {
	if input == nil {
		return &resourceEndpointOnNetRuleWebRequestHttpsModel{}
	}
	if m == nil {
		m = &resourceEndpointOnNetRuleWebRequestHttpsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["ip"]; ok {
		m.Ip = parseStringValue(v)
	}

	if v, ok := o["hostname"]; ok {
		m.Hostname = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointOnNetRuleModel) flattenEndpointOnNetRuleWebRequestHttpsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointOnNetRuleWebRequestHttpsModel {
	if o == nil {
		return []resourceEndpointOnNetRuleWebRequestHttpsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument web_request_https is not type of []interface{}.", "")
		return []resourceEndpointOnNetRuleWebRequestHttpsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointOnNetRuleWebRequestHttpsModel{}
	}

	values := make([]resourceEndpointOnNetRuleWebRequestHttpsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointOnNetRuleWebRequestHttpsModel
		if i < len(s.WebRequestHttps) {
			m = s.WebRequestHttps[i]
		}
		values[i] = *m.flattenEndpointOnNetRuleWebRequestHttps(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointOnNetRuleDnsRequestModel) flattenEndpointOnNetRuleDnsRequest(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointOnNetRuleDnsRequestModel {
	if input == nil {
		return &resourceEndpointOnNetRuleDnsRequestModel{}
	}
	if m == nil {
		m = &resourceEndpointOnNetRuleDnsRequestModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["ip"]; ok {
		m.Ip = parseStringValue(v)
	}

	if v, ok := o["hostname"]; ok {
		m.Hostname = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointOnNetRuleModel) flattenEndpointOnNetRuleDnsRequestList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointOnNetRuleDnsRequestModel {
	if o == nil {
		return []resourceEndpointOnNetRuleDnsRequestModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument dns_request is not type of []interface{}.", "")
		return []resourceEndpointOnNetRuleDnsRequestModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointOnNetRuleDnsRequestModel{}
	}

	values := make([]resourceEndpointOnNetRuleDnsRequestModel, len(l))
	for i, ele := range l {
		var m resourceEndpointOnNetRuleDnsRequestModel
		if i < len(s.DnsRequest) {
			m = s.DnsRequest[i]
		}
		values[i] = *m.flattenEndpointOnNetRuleDnsRequest(ctx, ele, diags)
	}

	return values
}

func (data *resourceEndpointOnNetRuleWebRequestHttpsModel) expandEndpointOnNetRuleWebRequestHttps(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		result["ip"] = data.Ip.ValueString()
	}

	if !data.Hostname.IsNull() && !data.Hostname.IsUnknown() {
		result["hostname"] = data.Hostname.ValueString()
	}

	return result
}

func (s *resourceEndpointOnNetRuleModel) expandEndpointOnNetRuleWebRequestHttpsList(ctx context.Context, l []resourceEndpointOnNetRuleWebRequestHttpsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointOnNetRuleWebRequestHttps(ctx, diags)
	}
	return result
}

func (data *resourceEndpointOnNetRuleDnsRequestModel) expandEndpointOnNetRuleDnsRequest(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		result["ip"] = data.Ip.ValueString()
	}

	if !data.Hostname.IsNull() && !data.Hostname.IsUnknown() {
		result["hostname"] = data.Hostname.ValueString()
	}

	return result
}

func (s *resourceEndpointOnNetRuleModel) expandEndpointOnNetRuleDnsRequestList(ctx context.Context, l []resourceEndpointOnNetRuleDnsRequestModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointOnNetRuleDnsRequest(ctx, diags)
	}
	return result
}
