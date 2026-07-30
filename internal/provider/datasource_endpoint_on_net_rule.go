// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointOnNetRule{}

func newDatasourceEndpointOnNetRule() datasource.DataSource {
	return &datasourceEndpointOnNetRule{}
}

type datasourceEndpointOnNetRule struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointOnNetRuleModel describes the datasource data model.
type datasourceEndpointOnNetRuleModel struct {
	PrimaryKey      types.String                                      `tfsdk:"primary_key"`
	PublicIp        types.String                                      `tfsdk:"public_ip"`
	DhcpServerIp    types.String                                      `tfsdk:"dhcp_server_ip"`
	DhcpServerMac   types.String                                      `tfsdk:"dhcp_server_mac"`
	DhcpServerCode  types.String                                      `tfsdk:"dhcp_server_code"`
	DnsServerIp     types.String                                      `tfsdk:"dns_server_ip"`
	PingServer      types.String                                      `tfsdk:"ping_server"`
	LocalIp         types.String                                      `tfsdk:"local_ip"`
	GatewayMac      types.String                                      `tfsdk:"gateway_mac"`
	WebRequestHttp  types.String                                      `tfsdk:"web_request_http"`
	WebRequestHttps []datasourceEndpointOnNetRuleWebRequestHttpsModel `tfsdk:"web_request_https"`
	DnsRequest      []datasourceEndpointOnNetRuleDnsRequestModel      `tfsdk:"dns_request"`
}

func (r *datasourceEndpointOnNetRule) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_on_net_rule"
}

func (r *datasourceEndpointOnNetRule) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Endpoint on net rule",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 128),
				},
				Required: true,
			},
			"public_ip": schema.StringAttribute{
				Computed: true,
			},
			"dhcp_server_ip": schema.StringAttribute{
				Computed: true,
			},
			"dhcp_server_mac": schema.StringAttribute{
				Computed: true,
			},
			"dhcp_server_code": schema.StringAttribute{
				Computed: true,
			},
			"dns_server_ip": schema.StringAttribute{
				Computed: true,
			},
			"ping_server": schema.StringAttribute{
				Computed: true,
			},
			"local_ip": schema.StringAttribute{
				Computed: true,
			},
			"gateway_mac": schema.StringAttribute{
				Computed: true,
			},
			"web_request_http": schema.StringAttribute{
				Computed: true,
			},
			"web_request_https": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							Computed: true,
						},
						"hostname": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"dns_request": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							Computed: true,
						},
						"hostname": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceEndpointOnNetRule) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceEndpointOnNetRule) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointOnNetRuleModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointOnNetRule(ctx, "read", diags))

	read_output, err := c.ReadEndpointOnNetRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
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

func (m *datasourceEndpointOnNetRuleModel) refreshEndpointOnNetRule(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *datasourceEndpointOnNetRuleModel) getURLObjectEndpointOnNetRule(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointOnNetRuleWebRequestHttpsModel struct {
	Ip       types.String `tfsdk:"ip"`
	Hostname types.String `tfsdk:"hostname"`
}

type datasourceEndpointOnNetRuleDnsRequestModel struct {
	Ip       types.String `tfsdk:"ip"`
	Hostname types.String `tfsdk:"hostname"`
}

func (m *datasourceEndpointOnNetRuleWebRequestHttpsModel) flattenEndpointOnNetRuleWebRequestHttps(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointOnNetRuleWebRequestHttpsModel {
	if input == nil {
		return &datasourceEndpointOnNetRuleWebRequestHttpsModel{}
	}
	if m == nil {
		m = &datasourceEndpointOnNetRuleWebRequestHttpsModel{}
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

func (s *datasourceEndpointOnNetRuleModel) flattenEndpointOnNetRuleWebRequestHttpsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointOnNetRuleWebRequestHttpsModel {
	if o == nil {
		return []datasourceEndpointOnNetRuleWebRequestHttpsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument web_request_https is not type of []interface{}.", "")
		return []datasourceEndpointOnNetRuleWebRequestHttpsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointOnNetRuleWebRequestHttpsModel{}
	}

	values := make([]datasourceEndpointOnNetRuleWebRequestHttpsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointOnNetRuleWebRequestHttpsModel
		if i < len(s.WebRequestHttps) {
			m = s.WebRequestHttps[i]
		}
		values[i] = *m.flattenEndpointOnNetRuleWebRequestHttps(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointOnNetRuleDnsRequestModel) flattenEndpointOnNetRuleDnsRequest(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointOnNetRuleDnsRequestModel {
	if input == nil {
		return &datasourceEndpointOnNetRuleDnsRequestModel{}
	}
	if m == nil {
		m = &datasourceEndpointOnNetRuleDnsRequestModel{}
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

func (s *datasourceEndpointOnNetRuleModel) flattenEndpointOnNetRuleDnsRequestList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointOnNetRuleDnsRequestModel {
	if o == nil {
		return []datasourceEndpointOnNetRuleDnsRequestModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument dns_request is not type of []interface{}.", "")
		return []datasourceEndpointOnNetRuleDnsRequestModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointOnNetRuleDnsRequestModel{}
	}

	values := make([]datasourceEndpointOnNetRuleDnsRequestModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointOnNetRuleDnsRequestModel
		if i < len(s.DnsRequest) {
			m = s.DnsRequest[i]
		}
		values[i] = *m.flattenEndpointOnNetRuleDnsRequest(ctx, ele, diags)
	}

	return values
}
