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
var _ datasource.DataSource = &datasourceNetworkDnsRules{}

func newDatasourceNetworkDnsRules() datasource.DataSource {
	return &datasourceNetworkDnsRules{}
}

type datasourceNetworkDnsRules struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceNetworkDnsRulesModel describes the datasource data model.
type datasourceNetworkDnsRulesModel struct {
	PrimaryKey     types.String                                            `tfsdk:"primary_key"`
	PrimaryDns     types.String                                            `tfsdk:"primary_dns"`
	SecondaryDns   types.String                                            `tfsdk:"secondary_dns"`
	Domains        types.Set                                               `tfsdk:"domains"`
	PopDnsOverride map[string]datasourceNetworkDnsRulesPopDnsOverrideModel `tfsdk:"pop_dns_override"`
	ForPrivate     types.Bool                                              `tfsdk:"for_private"`
}

func (r *datasourceNetworkDnsRules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_dns_rules"
}

func (r *datasourceNetworkDnsRules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(30),
				},
				Required: true,
			},
			"primary_dns": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"secondary_dns": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"domains": schema.SetAttribute{
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
			"pop_dns_override": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pop": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"primary_dns": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"secondary_dns": schema.StringAttribute{
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

func (r *datasourceNetworkDnsRules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_network_dns_rules"
}

func (r *datasourceNetworkDnsRules) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceNetworkDnsRulesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkDnsRules(ctx, "read", diags))

	read_output, err := c.ReadNetworkDnsRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshNetworkDnsRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceNetworkDnsRulesModel) refreshNetworkDnsRules(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryDns"]; ok {
		m.PrimaryDns = parseStringValue(v)
	}

	if v, ok := o["secondaryDns"]; ok {
		m.SecondaryDns = parseStringValue(v)
	}

	if v, ok := o["domains"]; ok {
		m.Domains = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["popDnsOverride"]; ok {
		m.PopDnsOverride = m.flattenNetworkDnsRulesPopDnsOverrideMap(ctx, v.(map[string]interface{}), &diags)
	}

	if v, ok := o["forPrivate"]; ok {
		m.ForPrivate = parseBoolValue(v)
	}

	return diags
}

func (data *datasourceNetworkDnsRulesModel) getURLObjectNetworkDnsRules(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceNetworkDnsRulesPopDnsOverrideModel struct {
	Pop          types.String `tfsdk:"pop"`
	PrimaryDns   types.String `tfsdk:"primary_dns"`
	SecondaryDns types.String `tfsdk:"secondary_dns"`
}

func (m *datasourceNetworkDnsRulesPopDnsOverrideModel) flattenNetworkDnsRulesPopDnsOverride(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceNetworkDnsRulesPopDnsOverrideModel {
	if input == nil {
		return &datasourceNetworkDnsRulesPopDnsOverrideModel{}
	}
	if m == nil {
		m = &datasourceNetworkDnsRulesPopDnsOverrideModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["pop"]; ok {
		m.Pop = parseStringValue(v)
	}

	if v, ok := o["primaryDns"]; ok {
		m.PrimaryDns = parseStringValue(v)
	}

	if v, ok := o["secondaryDns"]; ok {
		m.SecondaryDns = parseStringValue(v)
	}

	return m
}

func (s *datasourceNetworkDnsRulesModel) flattenNetworkDnsRulesPopDnsOverrideMap(ctx context.Context, o map[string]interface{}, diags *diag.Diagnostics) map[string]datasourceNetworkDnsRulesPopDnsOverrideModel {
	result := make(map[string]datasourceNetworkDnsRulesPopDnsOverrideModel)
	for k, v := range o {
		var m datasourceNetworkDnsRulesPopDnsOverrideModel
		m = *m.flattenNetworkDnsRulesPopDnsOverride(ctx, v, diags)
		result[k] = m
	}
	return result
}
