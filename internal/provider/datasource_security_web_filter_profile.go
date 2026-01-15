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
var _ datasource.DataSource = &datasourceSecurityWebFilterProfile{}

func newDatasourceSecurityWebFilterProfile() datasource.DataSource {
	return &datasourceSecurityWebFilterProfile{}
}

type datasourceSecurityWebFilterProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityWebFilterProfileModel describes the datasource data model.
type datasourceSecurityWebFilterProfileModel struct {
	PrimaryKey                     types.String                                                            `tfsdk:"primary_key"`
	FortiguardFilters              []datasourceSecurityWebFilterProfileFortiguardFiltersModel              `tfsdk:"fortiguard_filters"`
	FortiguardLocalCategoryFilters []datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel `tfsdk:"fortiguard_local_category_filters"`
	FqdnThreatFeedFilters          []datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel          `tfsdk:"fqdn_threat_feed_filters"`
	UseFortiguardFilters           types.String                                                            `tfsdk:"use_fortiguard_filters"`
	BlockInvalidUrl                types.String                                                            `tfsdk:"block_invalid_url"`
	EnforceSafeSearch              types.String                                                            `tfsdk:"enforce_safe_search"`
	LogSearchedKeywords            types.String                                                            `tfsdk:"log_searched_keywords"`
	TrafficOnRatingError           types.String                                                            `tfsdk:"traffic_on_rating_error"`
	ContentFilters                 []datasourceSecurityWebFilterProfileContentFiltersModel                 `tfsdk:"content_filters"`
	UrlFilters                     []datasourceSecurityWebFilterProfileUrlFiltersModel                     `tfsdk:"url_filters"`
	HttpHeaders                    []datasourceSecurityWebFilterProfileHttpHeadersModel                    `tfsdk:"http_headers"`
	Direction                      types.String                                                            `tfsdk:"direction"`
}

func (r *datasourceSecurityWebFilterProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_web_filter_profile"
}

func (r *datasourceSecurityWebFilterProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"use_fortiguard_filters": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"block_invalid_url": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"enforce_safe_search": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"log_searched_keywords": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"traffic_on_rating_error": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				MarkdownDescription: "The direction of the target resource.\nSupported values: internal-profiles, outbound-profiles.",
				Computed:            true,
				Optional:            true,
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
						"warning_duration": schema.StringAttribute{
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
			"fortiguard_local_category_filters": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "monitor", "block", "warning", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"warning_duration": schema.StringAttribute{
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
										stringvalidator.OneOf("security/fortiguard-local-categories"),
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
			"fqdn_threat_feed_filters": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "monitor", "block", "warning", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"warning_duration": schema.StringAttribute{
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
										stringvalidator.OneOf("security/url-threat-feeds"),
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
			"content_filters": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"pattern": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"pattern_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("wildcard", "regexp"),
							},
							Computed: true,
							Optional: true,
						},
						"lang": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("western", "simch", "trach", "japanese", "korean", "french", "thai", "spanish", "cyrillic"),
							},
							Computed: true,
							Optional: true,
						},
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("exempt", "block"),
							},
							Computed: true,
							Optional: true,
						},
						"score": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"url_filters": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
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
					},
				},
				Computed: true,
				Optional: true,
			},
			"http_headers": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("add-to-request", "add-to-response", "remove-from-request", "remove-from-response"),
							},
							Computed: true,
							Optional: true,
						},
						"content": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"destinations": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("network/hosts", "network/host-groups"),
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
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityWebFilterProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_web_filter_profile"
}

func (r *datasourceSecurityWebFilterProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityWebFilterProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityWebFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityWebFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityWebFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityWebFilterProfileModel) refreshSecurityWebFilterProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["fortiguardFilters"]; ok {
		m.FortiguardFilters = m.flattenSecurityWebFilterProfileFortiguardFiltersList(ctx, v, &diags)
	}

	if v, ok := o["fortiguardLocalCategoryFilters"]; ok {
		m.FortiguardLocalCategoryFilters = m.flattenSecurityWebFilterProfileFortiguardLocalCategoryFiltersList(ctx, v, &diags)
	}

	if v, ok := o["fqdnThreatFeedFilters"]; ok {
		m.FqdnThreatFeedFilters = m.flattenSecurityWebFilterProfileFqdnThreatFeedFiltersList(ctx, v, &diags)
	}

	if v, ok := o["useFortiguardFilters"]; ok {
		m.UseFortiguardFilters = parseStringValue(v)
	}

	if v, ok := o["blockInvalidUrl"]; ok {
		m.BlockInvalidUrl = parseStringValue(v)
	}

	if v, ok := o["enforceSafeSearch"]; ok {
		m.EnforceSafeSearch = parseStringValue(v)
	}

	if v, ok := o["logSearchedKeywords"]; ok {
		m.LogSearchedKeywords = parseStringValue(v)
	}

	if v, ok := o["trafficOnRatingError"]; ok {
		m.TrafficOnRatingError = parseStringValue(v)
	}

	if v, ok := o["contentFilters"]; ok {
		m.ContentFilters = m.flattenSecurityWebFilterProfileContentFiltersList(ctx, v, &diags)
	}

	if v, ok := o["urlFilters"]; ok {
		m.UrlFilters = m.flattenSecurityWebFilterProfileUrlFiltersList(ctx, v, &diags)
	}

	if v, ok := o["httpHeaders"]; ok {
		m.HttpHeaders = m.flattenSecurityWebFilterProfileHttpHeadersList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityWebFilterProfileModel) getURLObjectSecurityWebFilterProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		diags.AddWarning("\"direction\" is deprecated and may be removed in future.",
			"It is recommended to recreate the resource without \"direction\" to avoid unexpected behavior in future.",
		)
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityWebFilterProfileFortiguardFiltersModel struct {
	Action          types.String                                                      `tfsdk:"action"`
	Category        *datasourceSecurityWebFilterProfileFortiguardFiltersCategoryModel `tfsdk:"category"`
	WarningDuration types.String                                                      `tfsdk:"warning_duration"`
}

type datasourceSecurityWebFilterProfileFortiguardFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel struct {
	Action          types.String                                                                   `tfsdk:"action"`
	Category        *datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel `tfsdk:"category"`
	WarningDuration types.String                                                                   `tfsdk:"warning_duration"`
}

type datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel struct {
	Action          types.String                                                          `tfsdk:"action"`
	Category        *datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel `tfsdk:"category"`
	WarningDuration types.String                                                          `tfsdk:"warning_duration"`
}

type datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityWebFilterProfileContentFiltersModel struct {
	Status      types.String  `tfsdk:"status"`
	Pattern     types.String  `tfsdk:"pattern"`
	PatternType types.String  `tfsdk:"pattern_type"`
	Lang        types.String  `tfsdk:"lang"`
	Action      types.String  `tfsdk:"action"`
	Score       types.Float64 `tfsdk:"score"`
}

type datasourceSecurityWebFilterProfileUrlFiltersModel struct {
	Status types.String `tfsdk:"status"`
	Url    types.String `tfsdk:"url"`
	Type   types.String `tfsdk:"type"`
	Action types.String `tfsdk:"action"`
}

type datasourceSecurityWebFilterProfileHttpHeadersModel struct {
	Name         types.String                                                     `tfsdk:"name"`
	Action       types.String                                                     `tfsdk:"action"`
	Content      types.String                                                     `tfsdk:"content"`
	Destinations []datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel `tfsdk:"destinations"`
}

type datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityWebFilterProfileFortiguardFiltersModel) flattenSecurityWebFilterProfileFortiguardFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityWebFilterProfileFortiguardFiltersModel {
	if input == nil {
		return &datasourceSecurityWebFilterProfileFortiguardFiltersModel{}
	}
	if m == nil {
		m = &datasourceSecurityWebFilterProfileFortiguardFiltersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = m.Category.flattenSecurityWebFilterProfileFortiguardFiltersCategory(ctx, v, diags)
	}

	if v, ok := o["warningDuration"]; ok {
		m.WarningDuration = parseStringValue(v)
	}

	return m
}

func (s *datasourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileFortiguardFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityWebFilterProfileFortiguardFiltersModel {
	if o == nil {
		return []datasourceSecurityWebFilterProfileFortiguardFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument fortiguard_filters is not type of []interface{}.", "")
		return []datasourceSecurityWebFilterProfileFortiguardFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityWebFilterProfileFortiguardFiltersModel{}
	}

	values := make([]datasourceSecurityWebFilterProfileFortiguardFiltersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityWebFilterProfileFortiguardFiltersModel
		values[i] = *m.flattenSecurityWebFilterProfileFortiguardFilters(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityWebFilterProfileFortiguardFiltersCategoryModel) flattenSecurityWebFilterProfileFortiguardFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityWebFilterProfileFortiguardFiltersCategoryModel {
	if input == nil {
		return &datasourceSecurityWebFilterProfileFortiguardFiltersCategoryModel{}
	}
	if m == nil {
		m = &datasourceSecurityWebFilterProfileFortiguardFiltersCategoryModel{}
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

func (m *datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel) flattenSecurityWebFilterProfileFortiguardLocalCategoryFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel {
	if input == nil {
		return &datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel{}
	}
	if m == nil {
		m = &datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = m.Category.flattenSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategory(ctx, v, diags)
	}

	if v, ok := o["warningDuration"]; ok {
		m.WarningDuration = parseStringValue(v)
	}

	return m
}

func (s *datasourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileFortiguardLocalCategoryFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel {
	if o == nil {
		return []datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument fortiguard_local_category_filters is not type of []interface{}.", "")
		return []datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel{}
	}

	values := make([]datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel
		values[i] = *m.flattenSecurityWebFilterProfileFortiguardLocalCategoryFilters(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel) flattenSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel {
	if input == nil {
		return &datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel{}
	}
	if m == nil {
		m = &datasourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel{}
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

func (m *datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel) flattenSecurityWebFilterProfileFqdnThreatFeedFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel {
	if input == nil {
		return &datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel{}
	}
	if m == nil {
		m = &datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = m.Category.flattenSecurityWebFilterProfileFqdnThreatFeedFiltersCategory(ctx, v, diags)
	}

	if v, ok := o["warningDuration"]; ok {
		m.WarningDuration = parseStringValue(v)
	}

	return m
}

func (s *datasourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileFqdnThreatFeedFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel {
	if o == nil {
		return []datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument fqdn_threat_feed_filters is not type of []interface{}.", "")
		return []datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel{}
	}

	values := make([]datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel
		values[i] = *m.flattenSecurityWebFilterProfileFqdnThreatFeedFilters(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel) flattenSecurityWebFilterProfileFqdnThreatFeedFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel {
	if input == nil {
		return &datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel{}
	}
	if m == nil {
		m = &datasourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel{}
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

func (m *datasourceSecurityWebFilterProfileContentFiltersModel) flattenSecurityWebFilterProfileContentFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityWebFilterProfileContentFiltersModel {
	if input == nil {
		return &datasourceSecurityWebFilterProfileContentFiltersModel{}
	}
	if m == nil {
		m = &datasourceSecurityWebFilterProfileContentFiltersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["pattern"]; ok {
		m.Pattern = parseStringValue(v)
	}

	if v, ok := o["patternType"]; ok {
		m.PatternType = parseStringValue(v)
	}

	if v, ok := o["lang"]; ok {
		m.Lang = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["score"]; ok {
		m.Score = parseFloat64Value(v)
	}

	return m
}

func (s *datasourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileContentFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityWebFilterProfileContentFiltersModel {
	if o == nil {
		return []datasourceSecurityWebFilterProfileContentFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument content_filters is not type of []interface{}.", "")
		return []datasourceSecurityWebFilterProfileContentFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityWebFilterProfileContentFiltersModel{}
	}

	values := make([]datasourceSecurityWebFilterProfileContentFiltersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityWebFilterProfileContentFiltersModel
		values[i] = *m.flattenSecurityWebFilterProfileContentFilters(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityWebFilterProfileUrlFiltersModel) flattenSecurityWebFilterProfileUrlFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityWebFilterProfileUrlFiltersModel {
	if input == nil {
		return &datasourceSecurityWebFilterProfileUrlFiltersModel{}
	}
	if m == nil {
		m = &datasourceSecurityWebFilterProfileUrlFiltersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["url"]; ok {
		m.Url = parseStringValue(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	return m
}

func (s *datasourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileUrlFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityWebFilterProfileUrlFiltersModel {
	if o == nil {
		return []datasourceSecurityWebFilterProfileUrlFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument url_filters is not type of []interface{}.", "")
		return []datasourceSecurityWebFilterProfileUrlFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityWebFilterProfileUrlFiltersModel{}
	}

	values := make([]datasourceSecurityWebFilterProfileUrlFiltersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityWebFilterProfileUrlFiltersModel
		values[i] = *m.flattenSecurityWebFilterProfileUrlFilters(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityWebFilterProfileHttpHeadersModel) flattenSecurityWebFilterProfileHttpHeaders(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityWebFilterProfileHttpHeadersModel {
	if input == nil {
		return &datasourceSecurityWebFilterProfileHttpHeadersModel{}
	}
	if m == nil {
		m = &datasourceSecurityWebFilterProfileHttpHeadersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["content"]; ok {
		m.Content = parseStringValue(v)
	}

	if v, ok := o["destinations"]; ok {
		m.Destinations = m.flattenSecurityWebFilterProfileHttpHeadersDestinationsList(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileHttpHeadersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityWebFilterProfileHttpHeadersModel {
	if o == nil {
		return []datasourceSecurityWebFilterProfileHttpHeadersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument http_headers is not type of []interface{}.", "")
		return []datasourceSecurityWebFilterProfileHttpHeadersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityWebFilterProfileHttpHeadersModel{}
	}

	values := make([]datasourceSecurityWebFilterProfileHttpHeadersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityWebFilterProfileHttpHeadersModel
		values[i] = *m.flattenSecurityWebFilterProfileHttpHeaders(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel) flattenSecurityWebFilterProfileHttpHeadersDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel {
	if input == nil {
		return &datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel{}
	}
	if m == nil {
		m = &datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel{}
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

func (s *datasourceSecurityWebFilterProfileHttpHeadersModel) flattenSecurityWebFilterProfileHttpHeadersDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel {
	if o == nil {
		return []datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel{}
	}

	values := make([]datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityWebFilterProfileHttpHeadersDestinationsModel
		values[i] = *m.flattenSecurityWebFilterProfileHttpHeadersDestinations(ctx, ele, diags)
	}

	return values
}
