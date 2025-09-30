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
var _ datasource.DataSource = &datasourceSecurityVideoFilterProfile{}

func newDatasourceSecurityVideoFilterProfile() datasource.DataSource {
	return &datasourceSecurityVideoFilterProfile{}
}

type datasourceSecurityVideoFilterProfile struct {
	fortiClient *FortiClient
}

// datasourceSecurityVideoFilterProfileModel describes the datasource data model.
type datasourceSecurityVideoFilterProfileModel struct {
	PrimaryKey        types.String                                                 `tfsdk:"primary_key"`
	FortiguardFilters []datasourceSecurityVideoFilterProfileFortiguardFiltersModel `tfsdk:"fortiguard_filters"`
	DefaultAction     types.String                                                 `tfsdk:"default_action"`
	Channels          []datasourceSecurityVideoFilterProfileChannelsModel          `tfsdk:"channels"`
	Direction         types.String                                                 `tfsdk:"direction"`
}

func (r *datasourceSecurityVideoFilterProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_video_filter_profile"
}

func (r *datasourceSecurityVideoFilterProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"default_action": schema.StringAttribute{
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
								stringvalidator.OneOf("allow", "monitor", "block", "warning", "default"),
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
										stringvalidator.OneOf("security/video-filter-fortiguard-categories"),
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
			"channels": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "monitor", "block"),
							},
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
							Computed: true,
							Optional: true,
						},
						"channel_id": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 64),
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

func (r *datasourceSecurityVideoFilterProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityVideoFilterProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityVideoFilterProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityVideoFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityVideoFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityVideoFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityVideoFilterProfileModel) refreshSecurityVideoFilterProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["fortiguardFilters"]; ok {
		m.FortiguardFilters = m.flattenSecurityVideoFilterProfileFortiguardFiltersList(ctx, v, &diags)
	}

	if v, ok := o["defaultAction"]; ok {
		m.DefaultAction = parseStringValue(v)
	}

	if v, ok := o["channels"]; ok {
		m.Channels = m.flattenSecurityVideoFilterProfileChannelsList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityVideoFilterProfileModel) getURLObjectSecurityVideoFilterProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityVideoFilterProfileFortiguardFiltersModel struct {
	Action   types.String                                                        `tfsdk:"action"`
	Category *datasourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel `tfsdk:"category"`
}

type datasourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityVideoFilterProfileChannelsModel struct {
	Action    types.String `tfsdk:"action"`
	Name      types.String `tfsdk:"name"`
	ChannelId types.String `tfsdk:"channel_id"`
}

func (m *datasourceSecurityVideoFilterProfileFortiguardFiltersModel) flattenSecurityVideoFilterProfileFortiguardFilters(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityVideoFilterProfileFortiguardFiltersModel {
	if input == nil {
		return &datasourceSecurityVideoFilterProfileFortiguardFiltersModel{}
	}
	if m == nil {
		m = &datasourceSecurityVideoFilterProfileFortiguardFiltersModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = m.Category.flattenSecurityVideoFilterProfileFortiguardFiltersCategory(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityVideoFilterProfileModel) flattenSecurityVideoFilterProfileFortiguardFiltersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityVideoFilterProfileFortiguardFiltersModel {
	if o == nil {
		return []datasourceSecurityVideoFilterProfileFortiguardFiltersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument fortiguard_filters is not type of []interface{}.", "")
		return []datasourceSecurityVideoFilterProfileFortiguardFiltersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityVideoFilterProfileFortiguardFiltersModel{}
	}

	values := make([]datasourceSecurityVideoFilterProfileFortiguardFiltersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityVideoFilterProfileFortiguardFiltersModel
		values[i] = *m.flattenSecurityVideoFilterProfileFortiguardFilters(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel) flattenSecurityVideoFilterProfileFortiguardFiltersCategory(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel {
	if input == nil {
		return &datasourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel{}
	}
	if m == nil {
		m = &datasourceSecurityVideoFilterProfileFortiguardFiltersCategoryModel{}
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

func (m *datasourceSecurityVideoFilterProfileChannelsModel) flattenSecurityVideoFilterProfileChannels(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityVideoFilterProfileChannelsModel {
	if input == nil {
		return &datasourceSecurityVideoFilterProfileChannelsModel{}
	}
	if m == nil {
		m = &datasourceSecurityVideoFilterProfileChannelsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["channelId"]; ok {
		m.ChannelId = parseStringValue(v)
	}

	return m
}

func (s *datasourceSecurityVideoFilterProfileModel) flattenSecurityVideoFilterProfileChannelsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityVideoFilterProfileChannelsModel {
	if o == nil {
		return []datasourceSecurityVideoFilterProfileChannelsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument channels is not type of []interface{}.", "")
		return []datasourceSecurityVideoFilterProfileChannelsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityVideoFilterProfileChannelsModel{}
	}

	values := make([]datasourceSecurityVideoFilterProfileChannelsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityVideoFilterProfileChannelsModel
		values[i] = *m.flattenSecurityVideoFilterProfileChannels(ctx, ele, diags)
	}

	return values
}
