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
var _ datasource.DataSource = &datasourceSecurityDnsFilterProfile{}

func newDatasourceSecurityDnsFilterProfile() datasource.DataSource {
	return &datasourceSecurityDnsFilterProfile{}
}

type datasourceSecurityDnsFilterProfile struct {
	fortiClient *FortiClient
}

// datasourceSecurityDnsFilterProfileModel describes the datasource data model.
type datasourceSecurityDnsFilterProfileModel struct {
	PrimaryKey                    types.String                                                     `tfsdk:"primary_key"`
	UseForEdgeDevices             types.Bool                                                       `tfsdk:"use_for_edge_devices"`
	UseFortiguardFilters          types.String                                                     `tfsdk:"use_fortiguard_filters"`
	EnableAllLogs                 types.String                                                     `tfsdk:"enable_all_logs"`
	EnableBotnetBlocking          types.String                                                     `tfsdk:"enable_botnet_blocking"`
	EnableSafeSearch              types.String                                                     `tfsdk:"enable_safe_search"`
	AllowDnsRequestsOnRatingError types.String                                                     `tfsdk:"allow_dns_requests_on_rating_error"`
	DnsTranslationEntries         []datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel   `tfsdk:"dns_translation_entries"`
	DomainFilters                 []datasourceSecurityDnsFilterProfileDomainFiltersModel           `tfsdk:"domain_filters"`
	FortiguardFilters             []datasourceSecurityDnsFilterProfileFortiguardFiltersModel       `tfsdk:"fortiguard_filters"`
	DomainThreatFeedFilters       []datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel `tfsdk:"domain_threat_feed_filters"`
	Direction                     types.String                                                     `tfsdk:"direction"`
}

func (r *datasourceSecurityDnsFilterProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dns_filter_profile"
}

func (r *datasourceSecurityDnsFilterProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"use_for_edge_devices": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"use_fortiguard_filters": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"enable_all_logs": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"enable_botnet_blocking": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"enable_safe_search": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"allow_dns_requests_on_rating_error": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"direction": schema.StringAttribute{
				Description: "The direction of the target resource.",
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				Computed: true,
				Optional: true,
			},
			"dns_translation_entries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"src": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"dst": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"netmask": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"domain_filters": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"url": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("simple", "wildcard", "regex"),
							},
							Computed: true,
							Optional: true,
						},
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "block", "exempt", "monitor"),
							},
							Computed: true,
							Optional: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"fortiguard_filters": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "monitor", "block", "warning"),
							},
							Computed: true,
							Optional: true,
						},
						"category": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("security/fortiguard-categories"),
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
			"domain_threat_feed_filters": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "monitor", "block", "warning", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"category": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("security/domain-threat-feeds"),
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

func (r *datasourceSecurityDnsFilterProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityDnsFilterProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityDnsFilterProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDnsFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityDnsFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDnsFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityDnsFilterProfileModel) refreshSecurityDnsFilterProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["useForEdgeDevices"]; ok {
		m.UseForEdgeDevices = parseBoolValue(v)
	}

	if v, ok := o["useFortiguardFilters"]; ok {
		m.UseFortiguardFilters = parseStringValue(v)
	}

	if v, ok := o["enableAllLogs"]; ok {
		m.EnableAllLogs = parseStringValue(v)
	}

	if v, ok := o["enableBotnetBlocking"]; ok {
		m.EnableBotnetBlocking = parseStringValue(v)
	}

	if v, ok := o["enableSafeSearch"]; ok {
		m.EnableSafeSearch = parseStringValue(v)
	}

	if v, ok := o["allowDnsRequestsOnRatingError"]; ok {
		m.AllowDnsRequestsOnRatingError = parseStringValue(v)
	}

	if v, ok := o["dnsTranslationEntries"]; ok {
		m.DnsTranslationEntries = m.flattenSecurityDnsFilterProfileDnsTranslationEntriesList(ctx, v, &diags)
	}

	if v, ok := o["domainFilters"]; ok {
		m.DomainFilters = m.flattenSecurityDnsFilterProfileDomainFiltersList(ctx, v, &diags)
	}

	if v, ok := o["fortiguardFilters"]; ok {
		m.FortiguardFilters = m.flattenSecurityDnsFilterProfileFortiguardFiltersList(ctx, v, &diags)
	}

	if v, ok := o["domainThreatFeedFilters"]; ok {
		m.DomainThreatFeedFilters = m.flattenSecurityDnsFilterProfileDomainThreatFeedFiltersList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityDnsFilterProfileModel) getURLObjectSecurityDnsFilterProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel struct {
	Src     types.String `tfsdk:"src"`
	Dst     types.String `tfsdk:"dst"`
	Netmask types.String `tfsdk:"netmask"`
	Status  types.String `tfsdk:"status"`
}

type datasourceSecurityDnsFilterProfileDomainFiltersModel struct {
	Url     types.String `tfsdk:"url"`
	Type    types.String `tfsdk:"type"`
	Action  types.String `tfsdk:"action"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

type datasourceSecurityDnsFilterProfileFortiguardFiltersModel struct {
	Action   types.String                                                      `tfsdk:"action"`
	Category *datasourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel `tfsdk:"category"`
}

type datasourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel struct {
	Action   types.String                                                            `tfsdk:"action"`
	Category *datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel `tfsdk:"category"`
}

type datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel) flattenSecurityDnsFilterProfileDnsTranslationEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel {
	if input == nil {
		return &datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel{}
	}
	if m == nil {
		m = &datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["src"]; ok {
		m.Src = parseStringValue(v)
	}

	if v, ok := o["dst"]; ok {
		m.Dst = parseStringValue(v)
	}

	if v, ok := o["netmask"]; ok {
		m.Netmask = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	return m
}

func (s *datasourceSecurityDnsFilterProfileModel) flattenSecurityDnsFilterProfileDnsTranslationEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel {
	if o == nil {
		return []datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument dns_translation_entries is not type of []interface{}.", "")
		return []datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel{}
	}

	values := make([]datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDnsFilterProfileDnsTranslationEntriesModel
		values[i] = *m.flattenSecurityDnsFilterProfileDnsTranslationEntries(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityDnsFilterProfileDomainFiltersModel) flattenSecurityDnsFilterProfileDomainFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDnsFilterProfileDomainFiltersModel {
	if input == nil {
		return &datasourceSecurityDnsFilterProfileDomainFiltersModel{}
	}
	if m == nil {
		m = &datasourceSecurityDnsFilterProfileDomainFiltersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["url"]; ok {
		m.Url = parseStringValue(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	return m
}

func (s *datasourceSecurityDnsFilterProfileModel) flattenSecurityDnsFilterProfileDomainFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDnsFilterProfileDomainFiltersModel {
	if o == nil {
		return []datasourceSecurityDnsFilterProfileDomainFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument domain_filters is not type of []interface{}.", "")
		return []datasourceSecurityDnsFilterProfileDomainFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDnsFilterProfileDomainFiltersModel{}
	}

	values := make([]datasourceSecurityDnsFilterProfileDomainFiltersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDnsFilterProfileDomainFiltersModel
		values[i] = *m.flattenSecurityDnsFilterProfileDomainFilters(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityDnsFilterProfileFortiguardFiltersModel) flattenSecurityDnsFilterProfileFortiguardFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDnsFilterProfileFortiguardFiltersModel {
	if input == nil {
		return &datasourceSecurityDnsFilterProfileFortiguardFiltersModel{}
	}
	if m == nil {
		m = &datasourceSecurityDnsFilterProfileFortiguardFiltersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = m.Category.flattenSecurityDnsFilterProfileFortiguardFiltersCategory(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityDnsFilterProfileModel) flattenSecurityDnsFilterProfileFortiguardFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDnsFilterProfileFortiguardFiltersModel {
	if o == nil {
		return []datasourceSecurityDnsFilterProfileFortiguardFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument fortiguard_filters is not type of []interface{}.", "")
		return []datasourceSecurityDnsFilterProfileFortiguardFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDnsFilterProfileFortiguardFiltersModel{}
	}

	values := make([]datasourceSecurityDnsFilterProfileFortiguardFiltersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDnsFilterProfileFortiguardFiltersModel
		values[i] = *m.flattenSecurityDnsFilterProfileFortiguardFilters(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel) flattenSecurityDnsFilterProfileFortiguardFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel {
	if input == nil {
		return &datasourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel{}
	}
	if m == nil {
		m = &datasourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (m *datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel) flattenSecurityDnsFilterProfileDomainThreatFeedFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel {
	if input == nil {
		return &datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel{}
	}
	if m == nil {
		m = &datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = m.Category.flattenSecurityDnsFilterProfileDomainThreatFeedFiltersCategory(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityDnsFilterProfileModel) flattenSecurityDnsFilterProfileDomainThreatFeedFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel {
	if o == nil {
		return []datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument domain_threat_feed_filters is not type of []interface{}.", "")
		return []datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel{}
	}

	values := make([]datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel
		values[i] = *m.flattenSecurityDnsFilterProfileDomainThreatFeedFilters(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel) flattenSecurityDnsFilterProfileDomainThreatFeedFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel {
	if input == nil {
		return &datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel{}
	}
	if m == nil {
		m = &datasourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}
