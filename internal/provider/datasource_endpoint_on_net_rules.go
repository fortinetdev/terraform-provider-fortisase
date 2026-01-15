// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointOnNetRules{}

func newDatasourceEndpointOnNetRules() datasource.DataSource {
	return &datasourceEndpointOnNetRules{}
}

type datasourceEndpointOnNetRules struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointOnNetRulesModel describes the datasource data model.
type datasourceEndpointOnNetRulesModel struct {
	PrimaryKey      types.String                                       `tfsdk:"primary_key"`
	PublicIp        types.String                                       `tfsdk:"public_ip"`
	DhcpServerIp    types.String                                       `tfsdk:"dhcp_server_ip"`
	DhcpServerMac   types.String                                       `tfsdk:"dhcp_server_mac"`
	DhcpServerCode  types.String                                       `tfsdk:"dhcp_server_code"`
	DnsServerIp     types.String                                       `tfsdk:"dns_server_ip"`
	PingServer      types.String                                       `tfsdk:"ping_server"`
	LocalIp         types.String                                       `tfsdk:"local_ip"`
	GatewayMac      types.String                                       `tfsdk:"gateway_mac"`
	WebRequestHttp  types.String                                       `tfsdk:"web_request_http"`
	WebRequestHttps []datasourceEndpointOnNetRulesWebRequestHttpsModel `tfsdk:"web_request_https"`
	DnsRequest      []datasourceEndpointOnNetRulesDnsRequestModel      `tfsdk:"dns_request"`
}

func (r *datasourceEndpointOnNetRules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_on_net_rules"
}

func (r *datasourceEndpointOnNetRules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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

func (r *datasourceEndpointOnNetRules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceEndpointOnNetRules) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointOnNetRulesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointOnNetRules(ctx, "read", diags))

	read_output, err := c.ReadEndpointOnNetRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
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

func (m *datasourceEndpointOnNetRulesModel) refreshEndpointOnNetRules(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *datasourceEndpointOnNetRulesModel) getURLObjectEndpointOnNetRules(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointOnNetRulesWebRequestHttpsModel struct {
	Ip       types.String `tfsdk:"ip"`
	Hostname types.String `tfsdk:"hostname"`
}

type datasourceEndpointOnNetRulesDnsRequestModel struct {
	Ip       types.String `tfsdk:"ip"`
	Hostname types.String `tfsdk:"hostname"`
}

func (m *datasourceEndpointOnNetRulesWebRequestHttpsModel) flattenEndpointOnNetRulesWebRequestHttps(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointOnNetRulesWebRequestHttpsModel {
	if input == nil {
		return &datasourceEndpointOnNetRulesWebRequestHttpsModel{}
	}
	if m == nil {
		m = &datasourceEndpointOnNetRulesWebRequestHttpsModel{}
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

func (s *datasourceEndpointOnNetRulesModel) flattenEndpointOnNetRulesWebRequestHttpsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointOnNetRulesWebRequestHttpsModel {
	if o == nil {
		return []datasourceEndpointOnNetRulesWebRequestHttpsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument web_request_https is not type of []interface{}.", "")
		return []datasourceEndpointOnNetRulesWebRequestHttpsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointOnNetRulesWebRequestHttpsModel{}
	}

	values := make([]datasourceEndpointOnNetRulesWebRequestHttpsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointOnNetRulesWebRequestHttpsModel
		values[i] = *m.flattenEndpointOnNetRulesWebRequestHttps(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointOnNetRulesDnsRequestModel) flattenEndpointOnNetRulesDnsRequest(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointOnNetRulesDnsRequestModel {
	if input == nil {
		return &datasourceEndpointOnNetRulesDnsRequestModel{}
	}
	if m == nil {
		m = &datasourceEndpointOnNetRulesDnsRequestModel{}
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

func (s *datasourceEndpointOnNetRulesModel) flattenEndpointOnNetRulesDnsRequestList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointOnNetRulesDnsRequestModel {
	if o == nil {
		return []datasourceEndpointOnNetRulesDnsRequestModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument dns_request is not type of []interface{}.", "")
		return []datasourceEndpointOnNetRulesDnsRequestModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointOnNetRulesDnsRequestModel{}
	}

	values := make([]datasourceEndpointOnNetRulesDnsRequestModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointOnNetRulesDnsRequestModel
		values[i] = *m.flattenEndpointOnNetRulesDnsRequest(ctx, ele, diags)
	}

	return values
}
