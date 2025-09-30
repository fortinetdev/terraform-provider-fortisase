// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceNetworkImplicitDnsRules{}

func newDatasourceNetworkImplicitDnsRules() datasource.DataSource {
	return &datasourceNetworkImplicitDnsRules{}
}

type datasourceNetworkImplicitDnsRules struct {
	fortiClient *FortiClient
}

// datasourceNetworkImplicitDnsRulesModel describes the datasource data model.
type datasourceNetworkImplicitDnsRulesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	DnsServer  types.String `tfsdk:"dns_server"`
	DnsServer1 types.String `tfsdk:"dns_server1"`
	DnsServer2 types.String `tfsdk:"dns_server2"`
	Protocols  types.Set    `tfsdk:"protocols"`
	ForPrivate types.Bool   `tfsdk:"for_private"`
}

func (r *datasourceNetworkImplicitDnsRules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_implicit_dns_rules"
}

func (r *datasourceNetworkImplicitDnsRules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("vpn", "other", "implicit_all"),
				},
				Required: true,
			},
			"dns_server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("fortiguard", "google", "quad9", "cloudflare", "endpoint", "custom"),
				},
				Computed: true,
				Optional: true,
			},
			"dns_server1": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"dns_server2": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"protocols": schema.SetAttribute{
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
				},
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"for_private": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceNetworkImplicitDnsRules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceNetworkImplicitDnsRules) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceNetworkImplicitDnsRulesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRules(ctx, "read", diags))

	read_output, err := c.ReadNetworkImplicitDnsRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkImplicitDnsRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceNetworkImplicitDnsRulesModel) refreshNetworkImplicitDnsRules(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["dnsServer"]; ok {
		m.DnsServer = parseStringValue(v)
	}

	if v, ok := o["dnsServer1"]; ok {
		m.DnsServer1 = parseStringValue(v)
	}

	if v, ok := o["dnsServer2"]; ok {
		m.DnsServer2 = parseStringValue(v)
	}

	if v, ok := o["protocols"]; ok {
		m.Protocols = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["forPrivate"]; ok {
		m.ForPrivate = parseBoolValue(v)
	}

	return diags
}

func (data *datasourceNetworkImplicitDnsRulesModel) getURLObjectNetworkImplicitDnsRules(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
