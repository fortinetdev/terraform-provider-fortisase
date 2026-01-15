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
var _ resource.Resource = &resourceSecurityDnsFilterProfile{}

func newResourceSecurityDnsFilterProfile() resource.Resource {
	return &resourceSecurityDnsFilterProfile{}
}

type resourceSecurityDnsFilterProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityDnsFilterProfileModel describes the resource data model.
type resourceSecurityDnsFilterProfileModel struct {
	ID                            types.String                                                   `tfsdk:"id"`
	PrimaryKey                    types.String                                                   `tfsdk:"primary_key"`
	UseForEdgeDevices             types.Bool                                                     `tfsdk:"use_for_edge_devices"`
	UseFortiguardFilters          types.String                                                   `tfsdk:"use_fortiguard_filters"`
	EnableAllLogs                 types.String                                                   `tfsdk:"enable_all_logs"`
	EnableBotnetBlocking          types.String                                                   `tfsdk:"enable_botnet_blocking"`
	EnableSafeSearch              types.String                                                   `tfsdk:"enable_safe_search"`
	AllowDnsRequestsOnRatingError types.String                                                   `tfsdk:"allow_dns_requests_on_rating_error"`
	DnsTranslationEntries         []resourceSecurityDnsFilterProfileDnsTranslationEntriesModel   `tfsdk:"dns_translation_entries"`
	DomainFilters                 []resourceSecurityDnsFilterProfileDomainFiltersModel           `tfsdk:"domain_filters"`
	FortiguardFilters             []resourceSecurityDnsFilterProfileFortiguardFiltersModel       `tfsdk:"fortiguard_filters"`
	DomainThreatFeedFilters       []resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel `tfsdk:"domain_threat_feed_filters"`
	Direction                     types.String                                                   `tfsdk:"direction"`
}

func (r *resourceSecurityDnsFilterProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dns_filter_profile"
}

func (r *resourceSecurityDnsFilterProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				MarkdownDescription: "The direction of the target resource.\nSupported values: internal-profiles, outbound-profiles.",
				Computed:            true,
				Optional:            true,
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

func (r *resourceSecurityDnsFilterProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dns_filter_profile"
}

func (r *resourceSecurityDnsFilterProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDnsFilterProfile")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityDnsFilterProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecurityDnsFilterProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDnsFilterProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecurityDnsFilterProfile(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityDnsFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityDnsFilterProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDnsFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDnsFilterProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDnsFilterProfile")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDnsFilterProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDnsFilterProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDnsFilterProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDnsFilterProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityDnsFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDnsFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityDnsFilterProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDnsFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDnsFilterProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityDnsFilterProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDnsFilterProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDnsFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityDnsFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDnsFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDnsFilterProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityDnsFilterProfileModel) refreshSecurityDnsFilterProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
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
		convert_v := m.flattenSecurityDnsFilterProfileFortiguardFiltersList(ctx, v, &diags)
		if m.FortiguardFilters == nil || !isSetSuperset(convert_v, m.FortiguardFilters) {
			m.FortiguardFilters = convert_v
		}

	}

	if v, ok := o["domainThreatFeedFilters"]; ok {
		convert_v := m.flattenSecurityDnsFilterProfileDomainThreatFeedFiltersList(ctx, v, &diags)
		if m.DomainThreatFeedFilters == nil || !isSetSuperset(convert_v, m.DomainThreatFeedFilters) {
			m.DomainThreatFeedFilters = convert_v
		}

	}

	return diags
}

func (data *resourceSecurityDnsFilterProfileModel) getCreateObjectSecurityDnsFilterProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.UseForEdgeDevices.IsNull() {
		result["useForEdgeDevices"] = data.UseForEdgeDevices.ValueBool()
	}

	if !data.UseFortiguardFilters.IsNull() {
		result["useFortiguardFilters"] = data.UseFortiguardFilters.ValueString()
	}

	if !data.EnableAllLogs.IsNull() {
		result["enableAllLogs"] = data.EnableAllLogs.ValueString()
	}

	if !data.EnableBotnetBlocking.IsNull() {
		result["enableBotnetBlocking"] = data.EnableBotnetBlocking.ValueString()
	}

	if !data.EnableSafeSearch.IsNull() {
		result["enableSafeSearch"] = data.EnableSafeSearch.ValueString()
	}

	if !data.AllowDnsRequestsOnRatingError.IsNull() {
		result["allowDnsRequestsOnRatingError"] = data.AllowDnsRequestsOnRatingError.ValueString()
	}

	result["dnsTranslationEntries"] = data.expandSecurityDnsFilterProfileDnsTranslationEntriesList(ctx, data.DnsTranslationEntries, diags)

	result["domainFilters"] = data.expandSecurityDnsFilterProfileDomainFiltersList(ctx, data.DomainFilters, diags)

	result["fortiguardFilters"] = data.expandSecurityDnsFilterProfileFortiguardFiltersList(ctx, data.FortiguardFilters, diags)

	result["domainThreatFeedFilters"] = data.expandSecurityDnsFilterProfileDomainThreatFeedFiltersList(ctx, data.DomainThreatFeedFilters, diags)

	return &result
}

func (data *resourceSecurityDnsFilterProfileModel) getUpdateObjectSecurityDnsFilterProfile(ctx context.Context, state resourceSecurityDnsFilterProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.UseForEdgeDevices.IsNull() {
		result["useForEdgeDevices"] = data.UseForEdgeDevices.ValueBool()
	}

	if !data.UseFortiguardFilters.IsNull() {
		result["useFortiguardFilters"] = data.UseFortiguardFilters.ValueString()
	}

	if !data.EnableAllLogs.IsNull() {
		result["enableAllLogs"] = data.EnableAllLogs.ValueString()
	}

	if !data.EnableBotnetBlocking.IsNull() {
		result["enableBotnetBlocking"] = data.EnableBotnetBlocking.ValueString()
	}

	if !data.EnableSafeSearch.IsNull() {
		result["enableSafeSearch"] = data.EnableSafeSearch.ValueString()
	}

	if !data.AllowDnsRequestsOnRatingError.IsNull() {
		result["allowDnsRequestsOnRatingError"] = data.AllowDnsRequestsOnRatingError.ValueString()
	}

	if data.DnsTranslationEntries != nil {
		result["dnsTranslationEntries"] = data.expandSecurityDnsFilterProfileDnsTranslationEntriesList(ctx, data.DnsTranslationEntries, diags)
	}

	if data.DomainFilters != nil {
		result["domainFilters"] = data.expandSecurityDnsFilterProfileDomainFiltersList(ctx, data.DomainFilters, diags)
	}

	if data.FortiguardFilters != nil {
		result["fortiguardFilters"] = data.expandSecurityDnsFilterProfileFortiguardFiltersList(ctx, data.FortiguardFilters, diags)
	}

	if data.DomainThreatFeedFilters != nil {
		result["domainThreatFeedFilters"] = data.expandSecurityDnsFilterProfileDomainThreatFeedFiltersList(ctx, data.DomainThreatFeedFilters, diags)
	}

	return &result
}

func (data *resourceSecurityDnsFilterProfileModel) getURLObjectSecurityDnsFilterProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
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

type resourceSecurityDnsFilterProfileDnsTranslationEntriesModel struct {
	Src     types.String `tfsdk:"src"`
	Dst     types.String `tfsdk:"dst"`
	Netmask types.String `tfsdk:"netmask"`
	Status  types.String `tfsdk:"status"`
}

type resourceSecurityDnsFilterProfileDomainFiltersModel struct {
	Url     types.String `tfsdk:"url"`
	Type    types.String `tfsdk:"type"`
	Action  types.String `tfsdk:"action"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

type resourceSecurityDnsFilterProfileFortiguardFiltersModel struct {
	Action   types.String                                                    `tfsdk:"action"`
	Category *resourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel `tfsdk:"category"`
}

type resourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel struct {
	Action   types.String                                                          `tfsdk:"action"`
	Category *resourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel `tfsdk:"category"`
}

type resourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityDnsFilterProfileDnsTranslationEntriesModel) flattenSecurityDnsFilterProfileDnsTranslationEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDnsFilterProfileDnsTranslationEntriesModel {
	if input == nil {
		return &resourceSecurityDnsFilterProfileDnsTranslationEntriesModel{}
	}
	if m == nil {
		m = &resourceSecurityDnsFilterProfileDnsTranslationEntriesModel{}
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

func (s *resourceSecurityDnsFilterProfileModel) flattenSecurityDnsFilterProfileDnsTranslationEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDnsFilterProfileDnsTranslationEntriesModel {
	if o == nil {
		return []resourceSecurityDnsFilterProfileDnsTranslationEntriesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument dns_translation_entries is not type of []interface{}.", "")
		return []resourceSecurityDnsFilterProfileDnsTranslationEntriesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDnsFilterProfileDnsTranslationEntriesModel{}
	}

	values := make([]resourceSecurityDnsFilterProfileDnsTranslationEntriesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDnsFilterProfileDnsTranslationEntriesModel
		values[i] = *m.flattenSecurityDnsFilterProfileDnsTranslationEntries(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDnsFilterProfileDomainFiltersModel) flattenSecurityDnsFilterProfileDomainFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDnsFilterProfileDomainFiltersModel {
	if input == nil {
		return &resourceSecurityDnsFilterProfileDomainFiltersModel{}
	}
	if m == nil {
		m = &resourceSecurityDnsFilterProfileDomainFiltersModel{}
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

func (s *resourceSecurityDnsFilterProfileModel) flattenSecurityDnsFilterProfileDomainFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDnsFilterProfileDomainFiltersModel {
	if o == nil {
		return []resourceSecurityDnsFilterProfileDomainFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument domain_filters is not type of []interface{}.", "")
		return []resourceSecurityDnsFilterProfileDomainFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDnsFilterProfileDomainFiltersModel{}
	}

	values := make([]resourceSecurityDnsFilterProfileDomainFiltersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDnsFilterProfileDomainFiltersModel
		values[i] = *m.flattenSecurityDnsFilterProfileDomainFilters(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDnsFilterProfileFortiguardFiltersModel) flattenSecurityDnsFilterProfileFortiguardFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDnsFilterProfileFortiguardFiltersModel {
	if input == nil {
		return &resourceSecurityDnsFilterProfileFortiguardFiltersModel{}
	}
	if m == nil {
		m = &resourceSecurityDnsFilterProfileFortiguardFiltersModel{}
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

func (s *resourceSecurityDnsFilterProfileModel) flattenSecurityDnsFilterProfileFortiguardFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDnsFilterProfileFortiguardFiltersModel {
	if o == nil {
		return []resourceSecurityDnsFilterProfileFortiguardFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument fortiguard_filters is not type of []interface{}.", "")
		return []resourceSecurityDnsFilterProfileFortiguardFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDnsFilterProfileFortiguardFiltersModel{}
	}

	values := make([]resourceSecurityDnsFilterProfileFortiguardFiltersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDnsFilterProfileFortiguardFiltersModel
		values[i] = *m.flattenSecurityDnsFilterProfileFortiguardFilters(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel) flattenSecurityDnsFilterProfileFortiguardFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel {
	if input == nil {
		return &resourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel{}
	}
	if m == nil {
		m = &resourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel{}
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

func (m *resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel) flattenSecurityDnsFilterProfileDomainThreatFeedFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel {
	if input == nil {
		return &resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel{}
	}
	if m == nil {
		m = &resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel{}
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

func (s *resourceSecurityDnsFilterProfileModel) flattenSecurityDnsFilterProfileDomainThreatFeedFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel {
	if o == nil {
		return []resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument domain_threat_feed_filters is not type of []interface{}.", "")
		return []resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel{}
	}

	values := make([]resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel
		values[i] = *m.flattenSecurityDnsFilterProfileDomainThreatFeedFilters(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel) flattenSecurityDnsFilterProfileDomainThreatFeedFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel {
	if input == nil {
		return &resourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel{}
	}
	if m == nil {
		m = &resourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel{}
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

func (data *resourceSecurityDnsFilterProfileDnsTranslationEntriesModel) expandSecurityDnsFilterProfileDnsTranslationEntries(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Src.IsNull() {
		result["src"] = data.Src.ValueString()
	}

	if !data.Dst.IsNull() {
		result["dst"] = data.Dst.ValueString()
	}

	if !data.Netmask.IsNull() {
		result["netmask"] = data.Netmask.ValueString()
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	return result
}

func (s *resourceSecurityDnsFilterProfileModel) expandSecurityDnsFilterProfileDnsTranslationEntriesList(ctx context.Context, l []resourceSecurityDnsFilterProfileDnsTranslationEntriesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDnsFilterProfileDnsTranslationEntries(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDnsFilterProfileDomainFiltersModel) expandSecurityDnsFilterProfileDomainFilters(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Url.IsNull() {
		result["url"] = data.Url.ValueString()
	}

	if !data.Type.IsNull() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Enabled.IsNull() {
		result["enabled"] = data.Enabled.ValueBool()
	}

	return result
}

func (s *resourceSecurityDnsFilterProfileModel) expandSecurityDnsFilterProfileDomainFiltersList(ctx context.Context, l []resourceSecurityDnsFilterProfileDomainFiltersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDnsFilterProfileDomainFilters(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDnsFilterProfileFortiguardFiltersModel) expandSecurityDnsFilterProfileFortiguardFilters(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Category != nil && !isZeroStruct(*data.Category) {
		result["category"] = data.Category.expandSecurityDnsFilterProfileFortiguardFiltersCategory(ctx, diags)
	}

	return result
}

func (s *resourceSecurityDnsFilterProfileModel) expandSecurityDnsFilterProfileFortiguardFiltersList(ctx context.Context, l []resourceSecurityDnsFilterProfileFortiguardFiltersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDnsFilterProfileFortiguardFilters(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDnsFilterProfileFortiguardFiltersCategoryModel) expandSecurityDnsFilterProfileFortiguardFiltersCategory(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel) expandSecurityDnsFilterProfileDomainThreatFeedFilters(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if data.Category != nil && !isZeroStruct(*data.Category) {
		result["category"] = data.Category.expandSecurityDnsFilterProfileDomainThreatFeedFiltersCategory(ctx, diags)
	}

	return result
}

func (s *resourceSecurityDnsFilterProfileModel) expandSecurityDnsFilterProfileDomainThreatFeedFiltersList(ctx context.Context, l []resourceSecurityDnsFilterProfileDomainThreatFeedFiltersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDnsFilterProfileDomainThreatFeedFilters(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDnsFilterProfileDomainThreatFeedFiltersCategoryModel) expandSecurityDnsFilterProfileDomainThreatFeedFiltersCategory(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
