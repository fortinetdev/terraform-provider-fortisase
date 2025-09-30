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
	"strings"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSecurityWebFilterProfile{}

func newResourceSecurityWebFilterProfile() resource.Resource {
	return &resourceSecurityWebFilterProfile{}
}

type resourceSecurityWebFilterProfile struct {
	fortiClient *FortiClient
}

// resourceSecurityWebFilterProfileModel describes the resource data model.
type resourceSecurityWebFilterProfileModel struct {
	ID                             types.String                                                          `tfsdk:"id"`
	PrimaryKey                     types.String                                                          `tfsdk:"primary_key"`
	FortiguardFilters              []resourceSecurityWebFilterProfileFortiguardFiltersModel              `tfsdk:"fortiguard_filters"`
	FortiguardLocalCategoryFilters []resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel `tfsdk:"fortiguard_local_category_filters"`
	FqdnThreatFeedFilters          []resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel          `tfsdk:"fqdn_threat_feed_filters"`
	UseFortiguardFilters           types.String                                                          `tfsdk:"use_fortiguard_filters"`
	BlockInvalidUrl                types.String                                                          `tfsdk:"block_invalid_url"`
	EnforceSafeSearch              types.String                                                          `tfsdk:"enforce_safe_search"`
	TrafficOnRatingError           types.String                                                          `tfsdk:"traffic_on_rating_error"`
	ContentFilters                 []resourceSecurityWebFilterProfileContentFiltersModel                 `tfsdk:"content_filters"`
	UrlFilters                     []resourceSecurityWebFilterProfileUrlFiltersModel                     `tfsdk:"url_filters"`
	HttpHeaders                    []resourceSecurityWebFilterProfileHttpHeadersModel                    `tfsdk:"http_headers"`
	Direction                      types.String                                                          `tfsdk:"direction"`
}

func (r *resourceSecurityWebFilterProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_web_filter_profile"
}

func (r *resourceSecurityWebFilterProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"traffic_on_rating_error": schema.StringAttribute{
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

func (r *resourceSecurityWebFilterProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityWebFilterProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityWebFilterProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecurityWebFilterProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityWebFilterProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecurityWebFilterProfile(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityWebFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityWebFilterProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityWebFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityWebFilterProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityWebFilterProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityWebFilterProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityWebFilterProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityWebFilterProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityWebFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityWebFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityWebFilterProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityWebFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityWebFilterProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityWebFilterProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityWebFilterProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityWebFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityWebFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityWebFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityWebFilterProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected format: direction/primary_key, got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("direction"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), parts[1])...)
}

func (m *resourceSecurityWebFilterProfileModel) refreshSecurityWebFilterProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
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

func (data *resourceSecurityWebFilterProfileModel) getCreateObjectSecurityWebFilterProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	result["fortiguardFilters"] = data.expandSecurityWebFilterProfileFortiguardFiltersList(ctx, data.FortiguardFilters, diags)

	result["fortiguardLocalCategoryFilters"] = data.expandSecurityWebFilterProfileFortiguardLocalCategoryFiltersList(ctx, data.FortiguardLocalCategoryFilters, diags)

	result["fqdnThreatFeedFilters"] = data.expandSecurityWebFilterProfileFqdnThreatFeedFiltersList(ctx, data.FqdnThreatFeedFilters, diags)

	if !data.UseFortiguardFilters.IsNull() {
		result["useFortiguardFilters"] = data.UseFortiguardFilters.ValueString()
	}

	if !data.BlockInvalidUrl.IsNull() {
		result["blockInvalidUrl"] = data.BlockInvalidUrl.ValueString()
	}

	if !data.EnforceSafeSearch.IsNull() {
		result["enforceSafeSearch"] = data.EnforceSafeSearch.ValueString()
	}

	if !data.TrafficOnRatingError.IsNull() {
		result["trafficOnRatingError"] = data.TrafficOnRatingError.ValueString()
	}

	result["contentFilters"] = data.expandSecurityWebFilterProfileContentFiltersList(ctx, data.ContentFilters, diags)

	result["urlFilters"] = data.expandSecurityWebFilterProfileUrlFiltersList(ctx, data.UrlFilters, diags)

	result["httpHeaders"] = data.expandSecurityWebFilterProfileHttpHeadersList(ctx, data.HttpHeaders, diags)

	return &result
}

func (data *resourceSecurityWebFilterProfileModel) getUpdateObjectSecurityWebFilterProfile(ctx context.Context, state resourceSecurityWebFilterProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if len(data.FortiguardFilters) > 0 || !isSameStruct(data.FortiguardFilters, state.FortiguardFilters) {
		result["fortiguardFilters"] = data.expandSecurityWebFilterProfileFortiguardFiltersList(ctx, data.FortiguardFilters, diags)
	}

	if len(data.FortiguardLocalCategoryFilters) > 0 || !isSameStruct(data.FortiguardLocalCategoryFilters, state.FortiguardLocalCategoryFilters) {
		result["fortiguardLocalCategoryFilters"] = data.expandSecurityWebFilterProfileFortiguardLocalCategoryFiltersList(ctx, data.FortiguardLocalCategoryFilters, diags)
	}

	if len(data.FqdnThreatFeedFilters) > 0 || !isSameStruct(data.FqdnThreatFeedFilters, state.FqdnThreatFeedFilters) {
		result["fqdnThreatFeedFilters"] = data.expandSecurityWebFilterProfileFqdnThreatFeedFiltersList(ctx, data.FqdnThreatFeedFilters, diags)
	}

	if !data.UseFortiguardFilters.IsNull() {
		result["useFortiguardFilters"] = data.UseFortiguardFilters.ValueString()
	}

	if !data.BlockInvalidUrl.IsNull() {
		result["blockInvalidUrl"] = data.BlockInvalidUrl.ValueString()
	}

	if !data.EnforceSafeSearch.IsNull() {
		result["enforceSafeSearch"] = data.EnforceSafeSearch.ValueString()
	}

	if !data.TrafficOnRatingError.IsNull() {
		result["trafficOnRatingError"] = data.TrafficOnRatingError.ValueString()
	}

	if len(data.ContentFilters) > 0 || !isSameStruct(data.ContentFilters, state.ContentFilters) {
		result["contentFilters"] = data.expandSecurityWebFilterProfileContentFiltersList(ctx, data.ContentFilters, diags)
	}

	if len(data.UrlFilters) > 0 || !isSameStruct(data.UrlFilters, state.UrlFilters) {
		result["urlFilters"] = data.expandSecurityWebFilterProfileUrlFiltersList(ctx, data.UrlFilters, diags)
	}

	if len(data.HttpHeaders) > 0 || !isSameStruct(data.HttpHeaders, state.HttpHeaders) {
		result["httpHeaders"] = data.expandSecurityWebFilterProfileHttpHeadersList(ctx, data.HttpHeaders, diags)
	}

	return &result
}

func (data *resourceSecurityWebFilterProfileModel) getURLObjectSecurityWebFilterProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityWebFilterProfileFortiguardFiltersModel struct {
	Action          types.String                                                    `tfsdk:"action"`
	Category        *resourceSecurityWebFilterProfileFortiguardFiltersCategoryModel `tfsdk:"category"`
	WarningDuration types.String                                                    `tfsdk:"warning_duration"`
}

type resourceSecurityWebFilterProfileFortiguardFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel struct {
	Action          types.String                                                                 `tfsdk:"action"`
	Category        *resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel `tfsdk:"category"`
	WarningDuration types.String                                                                 `tfsdk:"warning_duration"`
}

type resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel struct {
	Action          types.String                                                        `tfsdk:"action"`
	Category        *resourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel `tfsdk:"category"`
	WarningDuration types.String                                                        `tfsdk:"warning_duration"`
}

type resourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityWebFilterProfileContentFiltersModel struct {
	Status      types.String  `tfsdk:"status"`
	Pattern     types.String  `tfsdk:"pattern"`
	PatternType types.String  `tfsdk:"pattern_type"`
	Lang        types.String  `tfsdk:"lang"`
	Action      types.String  `tfsdk:"action"`
	Score       types.Float64 `tfsdk:"score"`
}

type resourceSecurityWebFilterProfileUrlFiltersModel struct {
	Status types.String `tfsdk:"status"`
	Url    types.String `tfsdk:"url"`
	Type   types.String `tfsdk:"type"`
	Action types.String `tfsdk:"action"`
}

type resourceSecurityWebFilterProfileHttpHeadersModel struct {
	Name         types.String                                                   `tfsdk:"name"`
	Action       types.String                                                   `tfsdk:"action"`
	Content      types.String                                                   `tfsdk:"content"`
	Destinations []resourceSecurityWebFilterProfileHttpHeadersDestinationsModel `tfsdk:"destinations"`
}

type resourceSecurityWebFilterProfileHttpHeadersDestinationsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityWebFilterProfileFortiguardFiltersModel) flattenSecurityWebFilterProfileFortiguardFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityWebFilterProfileFortiguardFiltersModel {
	if input == nil {
		return &resourceSecurityWebFilterProfileFortiguardFiltersModel{}
	}
	if m == nil {
		m = &resourceSecurityWebFilterProfileFortiguardFiltersModel{}
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

func (s *resourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileFortiguardFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityWebFilterProfileFortiguardFiltersModel {
	if o == nil {
		return []resourceSecurityWebFilterProfileFortiguardFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument fortiguard_filters is not type of []interface{}.", "")
		return []resourceSecurityWebFilterProfileFortiguardFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityWebFilterProfileFortiguardFiltersModel{}
	}

	values := make([]resourceSecurityWebFilterProfileFortiguardFiltersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityWebFilterProfileFortiguardFiltersModel
		values[i] = *m.flattenSecurityWebFilterProfileFortiguardFilters(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityWebFilterProfileFortiguardFiltersCategoryModel) flattenSecurityWebFilterProfileFortiguardFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityWebFilterProfileFortiguardFiltersCategoryModel {
	if input == nil {
		return &resourceSecurityWebFilterProfileFortiguardFiltersCategoryModel{}
	}
	if m == nil {
		m = &resourceSecurityWebFilterProfileFortiguardFiltersCategoryModel{}
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

func (m *resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel) flattenSecurityWebFilterProfileFortiguardLocalCategoryFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel {
	if input == nil {
		return &resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel{}
	}
	if m == nil {
		m = &resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel{}
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

func (s *resourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileFortiguardLocalCategoryFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel {
	if o == nil {
		return []resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument fortiguard_local_category_filters is not type of []interface{}.", "")
		return []resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel{}
	}

	values := make([]resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel
		values[i] = *m.flattenSecurityWebFilterProfileFortiguardLocalCategoryFilters(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel) flattenSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel {
	if input == nil {
		return &resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel{}
	}
	if m == nil {
		m = &resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel{}
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

func (m *resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel) flattenSecurityWebFilterProfileFqdnThreatFeedFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel {
	if input == nil {
		return &resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel{}
	}
	if m == nil {
		m = &resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel{}
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

func (s *resourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileFqdnThreatFeedFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel {
	if o == nil {
		return []resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument fqdn_threat_feed_filters is not type of []interface{}.", "")
		return []resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel{}
	}

	values := make([]resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel
		values[i] = *m.flattenSecurityWebFilterProfileFqdnThreatFeedFilters(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel) flattenSecurityWebFilterProfileFqdnThreatFeedFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel {
	if input == nil {
		return &resourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel{}
	}
	if m == nil {
		m = &resourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel{}
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

func (m *resourceSecurityWebFilterProfileContentFiltersModel) flattenSecurityWebFilterProfileContentFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityWebFilterProfileContentFiltersModel {
	if input == nil {
		return &resourceSecurityWebFilterProfileContentFiltersModel{}
	}
	if m == nil {
		m = &resourceSecurityWebFilterProfileContentFiltersModel{}
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

func (s *resourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileContentFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityWebFilterProfileContentFiltersModel {
	if o == nil {
		return []resourceSecurityWebFilterProfileContentFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument content_filters is not type of []interface{}.", "")
		return []resourceSecurityWebFilterProfileContentFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityWebFilterProfileContentFiltersModel{}
	}

	values := make([]resourceSecurityWebFilterProfileContentFiltersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityWebFilterProfileContentFiltersModel
		values[i] = *m.flattenSecurityWebFilterProfileContentFilters(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityWebFilterProfileUrlFiltersModel) flattenSecurityWebFilterProfileUrlFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityWebFilterProfileUrlFiltersModel {
	if input == nil {
		return &resourceSecurityWebFilterProfileUrlFiltersModel{}
	}
	if m == nil {
		m = &resourceSecurityWebFilterProfileUrlFiltersModel{}
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

func (s *resourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileUrlFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityWebFilterProfileUrlFiltersModel {
	if o == nil {
		return []resourceSecurityWebFilterProfileUrlFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument url_filters is not type of []interface{}.", "")
		return []resourceSecurityWebFilterProfileUrlFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityWebFilterProfileUrlFiltersModel{}
	}

	values := make([]resourceSecurityWebFilterProfileUrlFiltersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityWebFilterProfileUrlFiltersModel
		values[i] = *m.flattenSecurityWebFilterProfileUrlFilters(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityWebFilterProfileHttpHeadersModel) flattenSecurityWebFilterProfileHttpHeaders(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityWebFilterProfileHttpHeadersModel {
	if input == nil {
		return &resourceSecurityWebFilterProfileHttpHeadersModel{}
	}
	if m == nil {
		m = &resourceSecurityWebFilterProfileHttpHeadersModel{}
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

func (s *resourceSecurityWebFilterProfileModel) flattenSecurityWebFilterProfileHttpHeadersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityWebFilterProfileHttpHeadersModel {
	if o == nil {
		return []resourceSecurityWebFilterProfileHttpHeadersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument http_headers is not type of []interface{}.", "")
		return []resourceSecurityWebFilterProfileHttpHeadersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityWebFilterProfileHttpHeadersModel{}
	}

	values := make([]resourceSecurityWebFilterProfileHttpHeadersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityWebFilterProfileHttpHeadersModel
		values[i] = *m.flattenSecurityWebFilterProfileHttpHeaders(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityWebFilterProfileHttpHeadersDestinationsModel) flattenSecurityWebFilterProfileHttpHeadersDestinations(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityWebFilterProfileHttpHeadersDestinationsModel {
	if input == nil {
		return &resourceSecurityWebFilterProfileHttpHeadersDestinationsModel{}
	}
	if m == nil {
		m = &resourceSecurityWebFilterProfileHttpHeadersDestinationsModel{}
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

func (s *resourceSecurityWebFilterProfileHttpHeadersModel) flattenSecurityWebFilterProfileHttpHeadersDestinationsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityWebFilterProfileHttpHeadersDestinationsModel {
	if o == nil {
		return []resourceSecurityWebFilterProfileHttpHeadersDestinationsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument destinations is not type of []interface{}.", "")
		return []resourceSecurityWebFilterProfileHttpHeadersDestinationsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityWebFilterProfileHttpHeadersDestinationsModel{}
	}

	values := make([]resourceSecurityWebFilterProfileHttpHeadersDestinationsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityWebFilterProfileHttpHeadersDestinationsModel
		values[i] = *m.flattenSecurityWebFilterProfileHttpHeadersDestinations(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityWebFilterProfileFortiguardFiltersModel) expandSecurityWebFilterProfileFortiguardFilters(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Category != nil && !isZeroStruct(*data.Category) {
		result["category"] = data.Category.expandSecurityWebFilterProfileFortiguardFiltersCategory(ctx, diags)
	}

	if !data.WarningDuration.IsNull() {
		result["warningDuration"] = data.WarningDuration.ValueString()
	}

	return result
}

func (s *resourceSecurityWebFilterProfileModel) expandSecurityWebFilterProfileFortiguardFiltersList(ctx context.Context, l []resourceSecurityWebFilterProfileFortiguardFiltersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityWebFilterProfileFortiguardFilters(ctx, diags)
	}
	return result
}

func (data *resourceSecurityWebFilterProfileFortiguardFiltersCategoryModel) expandSecurityWebFilterProfileFortiguardFiltersCategory(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel) expandSecurityWebFilterProfileFortiguardLocalCategoryFilters(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Category != nil && !isZeroStruct(*data.Category) {
		result["category"] = data.Category.expandSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategory(ctx, diags)
	}

	if !data.WarningDuration.IsNull() {
		result["warningDuration"] = data.WarningDuration.ValueString()
	}

	return result
}

func (s *resourceSecurityWebFilterProfileModel) expandSecurityWebFilterProfileFortiguardLocalCategoryFiltersList(ctx context.Context, l []resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityWebFilterProfileFortiguardLocalCategoryFilters(ctx, diags)
	}
	return result
}

func (data *resourceSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategoryModel) expandSecurityWebFilterProfileFortiguardLocalCategoryFiltersCategory(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel) expandSecurityWebFilterProfileFqdnThreatFeedFilters(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Category != nil && !isZeroStruct(*data.Category) {
		result["category"] = data.Category.expandSecurityWebFilterProfileFqdnThreatFeedFiltersCategory(ctx, diags)
	}

	if !data.WarningDuration.IsNull() {
		result["warningDuration"] = data.WarningDuration.ValueString()
	}

	return result
}

func (s *resourceSecurityWebFilterProfileModel) expandSecurityWebFilterProfileFqdnThreatFeedFiltersList(ctx context.Context, l []resourceSecurityWebFilterProfileFqdnThreatFeedFiltersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityWebFilterProfileFqdnThreatFeedFilters(ctx, diags)
	}
	return result
}

func (data *resourceSecurityWebFilterProfileFqdnThreatFeedFiltersCategoryModel) expandSecurityWebFilterProfileFqdnThreatFeedFiltersCategory(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityWebFilterProfileContentFiltersModel) expandSecurityWebFilterProfileContentFilters(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Pattern.IsNull() {
		result["pattern"] = data.Pattern.ValueString()
	}

	if !data.PatternType.IsNull() {
		result["patternType"] = data.PatternType.ValueString()
	}

	if !data.Lang.IsNull() {
		result["lang"] = data.Lang.ValueString()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Score.IsNull() {
		result["score"] = data.Score.ValueFloat64()
	}

	return result
}

func (s *resourceSecurityWebFilterProfileModel) expandSecurityWebFilterProfileContentFiltersList(ctx context.Context, l []resourceSecurityWebFilterProfileContentFiltersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityWebFilterProfileContentFilters(ctx, diags)
	}
	return result
}

func (data *resourceSecurityWebFilterProfileUrlFiltersModel) expandSecurityWebFilterProfileUrlFilters(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Url.IsNull() {
		result["url"] = data.Url.ValueString()
	}

	if !data.Type.IsNull() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	return result
}

func (s *resourceSecurityWebFilterProfileModel) expandSecurityWebFilterProfileUrlFiltersList(ctx context.Context, l []resourceSecurityWebFilterProfileUrlFiltersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityWebFilterProfileUrlFilters(ctx, diags)
	}
	return result
}

func (data *resourceSecurityWebFilterProfileHttpHeadersModel) expandSecurityWebFilterProfileHttpHeaders(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Name.IsNull() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Content.IsNull() {
		result["content"] = data.Content.ValueString()
	}

	result["destinations"] = data.expandSecurityWebFilterProfileHttpHeadersDestinationsList(ctx, data.Destinations, diags)

	return result
}

func (s *resourceSecurityWebFilterProfileModel) expandSecurityWebFilterProfileHttpHeadersList(ctx context.Context, l []resourceSecurityWebFilterProfileHttpHeadersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityWebFilterProfileHttpHeaders(ctx, diags)
	}
	return result
}

func (data *resourceSecurityWebFilterProfileHttpHeadersDestinationsModel) expandSecurityWebFilterProfileHttpHeadersDestinations(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityWebFilterProfileHttpHeadersModel) expandSecurityWebFilterProfileHttpHeadersDestinationsList(ctx context.Context, l []resourceSecurityWebFilterProfileHttpHeadersDestinationsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityWebFilterProfileHttpHeadersDestinations(ctx, diags)
	}
	return result
}
