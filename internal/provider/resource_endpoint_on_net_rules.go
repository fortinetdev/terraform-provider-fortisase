// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
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
var _ resource.Resource = &resourceEndpointOnNetRules{}

func newResourceEndpointOnNetRules() resource.Resource {
	return &resourceEndpointOnNetRules{}
}

type resourceEndpointOnNetRules struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointOnNetRulesModel describes the resource data model.
type resourceEndpointOnNetRulesModel struct {
	ID              types.String                                     `tfsdk:"id"`
	PrimaryKey      types.String                                     `tfsdk:"primary_key"`
	PublicIp        types.String                                     `tfsdk:"public_ip"`
	DhcpServerIp    types.String                                     `tfsdk:"dhcp_server_ip"`
	DhcpServerMac   types.String                                     `tfsdk:"dhcp_server_mac"`
	DhcpServerCode  types.String                                     `tfsdk:"dhcp_server_code"`
	DnsServerIp     types.String                                     `tfsdk:"dns_server_ip"`
	PingServer      types.String                                     `tfsdk:"ping_server"`
	LocalIp         types.String                                     `tfsdk:"local_ip"`
	GatewayMac      types.String                                     `tfsdk:"gateway_mac"`
	WebRequestHttp  types.String                                     `tfsdk:"web_request_http"`
	WebRequestHttps []resourceEndpointOnNetRulesWebRequestHttpsModel `tfsdk:"web_request_https"`
	DnsRequest      []resourceEndpointOnNetRulesDnsRequestModel      `tfsdk:"dns_request"`
}

func (r *resourceEndpointOnNetRules) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_on_net_rules"
}

func (r *resourceEndpointOnNetRules) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthBetween(1, 128),
				},
				Required: true,
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

func (r *resourceEndpointOnNetRules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_on_net_rules"
}

func (r *resourceEndpointOnNetRules) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointOnNetRulesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointOnNetRules(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointOnNetRules(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateEndpointOnNetRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointOnNetRules(ctx, "read", diags))

	read_output, err := c.ReadEndpointOnNetRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointOnNetRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointOnNetRules) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointOnNetRulesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointOnNetRulesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointOnNetRules(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointOnNetRules(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectEndpointOnNetRules(ctx, "read", diags))

	read_output, err := c.ReadEndpointOnNetRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointOnNetRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointOnNetRules) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointOnNetRulesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointOnNetRules(ctx, "delete", diags))

	output, err := c.DeleteEndpointOnNetRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceEndpointOnNetRules) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointOnNetRulesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointOnNetRules(ctx, "read", diags))

	read_output, err := c.ReadEndpointOnNetRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointOnNetRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointOnNetRules) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointOnNetRulesModel) refreshEndpointOnNetRules(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.WebRequestHttps = m.flattenEndpointOnNetRulesWebRequestHttpsList(ctx, v, &diags)
	}

	if v, ok := o["dnsRequest"]; ok {
		m.DnsRequest = m.flattenEndpointOnNetRulesDnsRequestList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointOnNetRulesModel) getCreateObjectEndpointOnNetRules(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.PublicIp.IsNull() {
		result["publicIp"] = data.PublicIp.ValueString()
	}

	if !data.DhcpServerIp.IsNull() {
		result["dhcpServerIp"] = data.DhcpServerIp.ValueString()
	}

	if !data.DhcpServerMac.IsNull() {
		result["dhcpServerMac"] = data.DhcpServerMac.ValueString()
	}

	if !data.DhcpServerCode.IsNull() {
		result["dhcpServerCode"] = data.DhcpServerCode.ValueString()
	}

	if !data.DnsServerIp.IsNull() {
		result["dnsServerIp"] = data.DnsServerIp.ValueString()
	}

	if !data.PingServer.IsNull() {
		result["pingServer"] = data.PingServer.ValueString()
	}

	if !data.LocalIp.IsNull() {
		result["localIp"] = data.LocalIp.ValueString()
	}

	if !data.GatewayMac.IsNull() {
		result["gatewayMac"] = data.GatewayMac.ValueString()
	}

	if !data.WebRequestHttp.IsNull() {
		result["webRequestHttp"] = data.WebRequestHttp.ValueString()
	}

	if data.WebRequestHttps != nil {
		result["webRequestHttps"] = data.expandEndpointOnNetRulesWebRequestHttpsList(ctx, data.WebRequestHttps, diags)
	}

	if data.DnsRequest != nil {
		result["dnsRequest"] = data.expandEndpointOnNetRulesDnsRequestList(ctx, data.DnsRequest, diags)
	}

	return &result
}

func (data *resourceEndpointOnNetRulesModel) getUpdateObjectEndpointOnNetRules(ctx context.Context, state resourceEndpointOnNetRulesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.PublicIp.IsNull() {
		result["publicIp"] = data.PublicIp.ValueString()
	}

	if !data.DhcpServerIp.IsNull() {
		result["dhcpServerIp"] = data.DhcpServerIp.ValueString()
	}

	if !data.DhcpServerMac.IsNull() {
		result["dhcpServerMac"] = data.DhcpServerMac.ValueString()
	}

	if !data.DhcpServerCode.IsNull() {
		result["dhcpServerCode"] = data.DhcpServerCode.ValueString()
	}

	if !data.DnsServerIp.IsNull() {
		result["dnsServerIp"] = data.DnsServerIp.ValueString()
	}

	if !data.PingServer.IsNull() {
		result["pingServer"] = data.PingServer.ValueString()
	}

	if !data.LocalIp.IsNull() {
		result["localIp"] = data.LocalIp.ValueString()
	}

	if !data.GatewayMac.IsNull() {
		result["gatewayMac"] = data.GatewayMac.ValueString()
	}

	if !data.WebRequestHttp.IsNull() {
		result["webRequestHttp"] = data.WebRequestHttp.ValueString()
	}

	if data.WebRequestHttps != nil {
		result["webRequestHttps"] = data.expandEndpointOnNetRulesWebRequestHttpsList(ctx, data.WebRequestHttps, diags)
	}

	if data.DnsRequest != nil {
		result["dnsRequest"] = data.expandEndpointOnNetRulesDnsRequestList(ctx, data.DnsRequest, diags)
	}

	return &result
}

func (data *resourceEndpointOnNetRulesModel) getURLObjectEndpointOnNetRules(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointOnNetRulesWebRequestHttpsModel struct {
	Ip       types.String `tfsdk:"ip"`
	Hostname types.String `tfsdk:"hostname"`
}

type resourceEndpointOnNetRulesDnsRequestModel struct {
	Ip       types.String `tfsdk:"ip"`
	Hostname types.String `tfsdk:"hostname"`
}

func (m *resourceEndpointOnNetRulesWebRequestHttpsModel) flattenEndpointOnNetRulesWebRequestHttps(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointOnNetRulesWebRequestHttpsModel {
	if input == nil {
		return &resourceEndpointOnNetRulesWebRequestHttpsModel{}
	}
	if m == nil {
		m = &resourceEndpointOnNetRulesWebRequestHttpsModel{}
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

func (s *resourceEndpointOnNetRulesModel) flattenEndpointOnNetRulesWebRequestHttpsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointOnNetRulesWebRequestHttpsModel {
	if o == nil {
		return []resourceEndpointOnNetRulesWebRequestHttpsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument web_request_https is not type of []interface{}.", "")
		return []resourceEndpointOnNetRulesWebRequestHttpsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointOnNetRulesWebRequestHttpsModel{}
	}

	values := make([]resourceEndpointOnNetRulesWebRequestHttpsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointOnNetRulesWebRequestHttpsModel
		values[i] = *m.flattenEndpointOnNetRulesWebRequestHttps(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointOnNetRulesDnsRequestModel) flattenEndpointOnNetRulesDnsRequest(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointOnNetRulesDnsRequestModel {
	if input == nil {
		return &resourceEndpointOnNetRulesDnsRequestModel{}
	}
	if m == nil {
		m = &resourceEndpointOnNetRulesDnsRequestModel{}
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

func (s *resourceEndpointOnNetRulesModel) flattenEndpointOnNetRulesDnsRequestList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointOnNetRulesDnsRequestModel {
	if o == nil {
		return []resourceEndpointOnNetRulesDnsRequestModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument dns_request is not type of []interface{}.", "")
		return []resourceEndpointOnNetRulesDnsRequestModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointOnNetRulesDnsRequestModel{}
	}

	values := make([]resourceEndpointOnNetRulesDnsRequestModel, len(l))
	for i, ele := range l {
		var m resourceEndpointOnNetRulesDnsRequestModel
		values[i] = *m.flattenEndpointOnNetRulesDnsRequest(ctx, ele, diags)
	}

	return values
}

func (data *resourceEndpointOnNetRulesWebRequestHttpsModel) expandEndpointOnNetRulesWebRequestHttps(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Ip.IsNull() {
		result["ip"] = data.Ip.ValueString()
	}

	if !data.Hostname.IsNull() {
		result["hostname"] = data.Hostname.ValueString()
	}

	return result
}

func (s *resourceEndpointOnNetRulesModel) expandEndpointOnNetRulesWebRequestHttpsList(ctx context.Context, l []resourceEndpointOnNetRulesWebRequestHttpsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointOnNetRulesWebRequestHttps(ctx, diags)
	}
	return result
}

func (data *resourceEndpointOnNetRulesDnsRequestModel) expandEndpointOnNetRulesDnsRequest(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Ip.IsNull() {
		result["ip"] = data.Ip.ValueString()
	}

	if !data.Hostname.IsNull() {
		result["hostname"] = data.Hostname.ValueString()
	}

	return result
}

func (s *resourceEndpointOnNetRulesModel) expandEndpointOnNetRulesDnsRequestList(ctx context.Context, l []resourceEndpointOnNetRulesDnsRequestModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointOnNetRulesDnsRequest(ctx, diags)
	}
	return result
}
