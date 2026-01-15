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
var _ datasource.DataSource = &datasourceEndpointsGroups{}

func newDatasourceEndpointsGroups() datasource.DataSource {
	return &datasourceEndpointsGroups{}
}

type datasourceEndpointsGroups struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointsGroupsModel describes the datasource data model.
type datasourceEndpointsGroupsModel struct {
	AdGroups    *datasourceEndpointsGroupsAdGroupsModel    `tfsdk:"ad_groups"`
	NonAdGroups *datasourceEndpointsGroupsNonAdGroupsModel `tfsdk:"non_ad_groups"`
	Guid        types.String                               `tfsdk:"guid"`
	Offset      types.Float64                              `tfsdk:"offset"`
	PrimaryKey  types.String                               `tfsdk:"primary_key"`
}

func (r *datasourceEndpointsGroups) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_groups"
}

func (r *datasourceEndpointsGroups) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"guid": schema.StringAttribute{
				MarkdownDescription: "UID of the group to expand to find child groups.",
				Computed:            true,
				Optional:            true,
			},
			"offset": schema.Float64Attribute{
				MarkdownDescription: "Specifies the starting position of AD groups. Based on this the results will be seperated in AD groups and non AD groups, with AD groups containing a \"total\" count.",
				Computed:            true,
				Optional:            true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "Primary key of the endpoint/domains entry.",
				Required:            true,
			},
			"ad_groups": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"total": schema.Float64Attribute{
						Computed: true,
						Optional: true,
					},
					"data": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.Float64Attribute{
									MarkdownDescription: "Id of the group.",
									Computed:            true,
									Optional:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: "Name of the group.",
									Computed:            true,
									Optional:            true,
								},
								"parent_id": schema.Float64Attribute{
									MarkdownDescription: "Parent id of the group.",
									Computed:            true,
									Optional:            true,
								},
								"guid": schema.StringAttribute{
									MarkdownDescription: "UID of the group.",
									Computed:            true,
									Optional:            true,
								},
								"path": schema.StringAttribute{
									MarkdownDescription: "Path of the group.",
									Computed:            true,
									Optional:            true,
								},
								"has_child": schema.BoolAttribute{
									MarkdownDescription: "Indicate if the group has child or not.",
									Computed:            true,
									Optional:            true,
								},
								"is_custom_group": schema.BoolAttribute{
									MarkdownDescription: "Indicate if the group is custom group or not.",
									Computed:            true,
									Optional:            true,
								},
								"domain_type": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("azure", "adfs"),
									},
									MarkdownDescription: "Type of the endpint/domains entry the group belongs to.\nSupported values: azure, adfs.",
									Computed:            true,
									Optional:            true,
								},
								"domain": schema.SingleNestedAttribute{
									MarkdownDescription: "Reference of the endpoint/domains entry the group belongs to.",
									Attributes: map[string]schema.Attribute{
										"primary_key": schema.StringAttribute{
											Computed: true,
											Optional: true,
										},
										"datasource": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("endpoint/domains"),
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
				Computed: true,
				Optional: true,
			},
			"non_ad_groups": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"data": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.Float64Attribute{
									MarkdownDescription: "Id of the group.",
									Computed:            true,
									Optional:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: "Name of the group.",
									Computed:            true,
									Optional:            true,
								},
								"parent_id": schema.Float64Attribute{
									MarkdownDescription: "Parent id of the group.",
									Computed:            true,
									Optional:            true,
								},
								"guid": schema.StringAttribute{
									MarkdownDescription: "UID of the group.",
									Computed:            true,
									Optional:            true,
								},
								"path": schema.StringAttribute{
									MarkdownDescription: "Path of the group.",
									Computed:            true,
									Optional:            true,
								},
								"has_child": schema.BoolAttribute{
									MarkdownDescription: "Indicate if the group has child or not.",
									Computed:            true,
									Optional:            true,
								},
								"is_custom_group": schema.BoolAttribute{
									MarkdownDescription: "Indicate if the group is custom group or not.",
									Computed:            true,
									Optional:            true,
								},
								"domain_type": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("azure", "adfs"),
									},
									MarkdownDescription: "Type of the endpint/domains entry the group belongs to.\nSupported values: azure, adfs.",
									Computed:            true,
									Optional:            true,
								},
								"domain": schema.SingleNestedAttribute{
									MarkdownDescription: "Reference of the endpoint/domains entry the group belongs to.",
									Attributes: map[string]schema.Attribute{
										"primary_key": schema.StringAttribute{
											Computed: true,
											Optional: true,
										},
										"datasource": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidator.OneOf("endpoint/domains"),
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
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceEndpointsGroups) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoints_groups"
}

func (r *datasourceEndpointsGroups) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointsGroupsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointsGroups(ctx, "read", diags))

	read_output, err := c.ReadEndpointsGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointsGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointsGroupsModel) refreshEndpointsGroups(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["adGroups"]; ok {
		m.AdGroups = m.AdGroups.flattenEndpointsGroupsAdGroups(ctx, v, &diags)
	}

	if v, ok := o["nonAdGroups"]; ok {
		m.NonAdGroups = m.NonAdGroups.flattenEndpointsGroupsNonAdGroups(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointsGroupsModel) getURLObjectEndpointsGroups(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Guid.IsNull() {
		result["guid"] = data.Guid.ValueString()
	}

	if !data.Offset.IsNull() {
		result["offset"] = data.Offset.ValueFloat64()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointsGroupsAdGroupsModel struct {
	Data  []datasourceEndpointsGroupsAdGroupsDataModel `tfsdk:"data"`
	Total types.Float64                                `tfsdk:"total"`
}

type datasourceEndpointsGroupsAdGroupsDataModel struct {
	Id            types.Float64                                     `tfsdk:"id"`
	Name          types.String                                      `tfsdk:"name"`
	ParentId      types.Float64                                     `tfsdk:"parent_id"`
	Guid          types.String                                      `tfsdk:"guid"`
	Path          types.String                                      `tfsdk:"path"`
	HasChild      types.Bool                                        `tfsdk:"has_child"`
	IsCustomGroup types.Bool                                        `tfsdk:"is_custom_group"`
	DomainType    types.String                                      `tfsdk:"domain_type"`
	Domain        *datasourceEndpointsGroupsAdGroupsDataDomainModel `tfsdk:"domain"`
}

type datasourceEndpointsGroupsAdGroupsDataDomainModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointsGroupsNonAdGroupsModel struct {
	Data []datasourceEndpointsGroupsNonAdGroupsDataModel `tfsdk:"data"`
}

type datasourceEndpointsGroupsNonAdGroupsDataModel struct {
	Id            types.Float64                                        `tfsdk:"id"`
	Name          types.String                                         `tfsdk:"name"`
	ParentId      types.Float64                                        `tfsdk:"parent_id"`
	Guid          types.String                                         `tfsdk:"guid"`
	Path          types.String                                         `tfsdk:"path"`
	HasChild      types.Bool                                           `tfsdk:"has_child"`
	IsCustomGroup types.Bool                                           `tfsdk:"is_custom_group"`
	DomainType    types.String                                         `tfsdk:"domain_type"`
	Domain        *datasourceEndpointsGroupsNonAdGroupsDataDomainModel `tfsdk:"domain"`
}

type datasourceEndpointsGroupsNonAdGroupsDataDomainModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceEndpointsGroupsAdGroupsModel) flattenEndpointsGroupsAdGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupsAdGroupsModel {
	if input == nil {
		return &datasourceEndpointsGroupsAdGroupsModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupsAdGroupsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["data"]; ok {
		m.Data = m.flattenEndpointsGroupsAdGroupsDataList(ctx, v, diags)
	}

	if v, ok := o["total"]; ok {
		m.Total = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceEndpointsGroupsAdGroupsDataModel) flattenEndpointsGroupsAdGroupsData(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupsAdGroupsDataModel {
	if input == nil {
		return &datasourceEndpointsGroupsAdGroupsDataModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupsAdGroupsDataModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["parentId"]; ok {
		m.ParentId = parseFloat64Value(v)
	}

	if v, ok := o["guid"]; ok {
		m.Guid = parseStringValue(v)
	}

	if v, ok := o["path"]; ok {
		m.Path = parseStringValue(v)
	}

	if v, ok := o["hasChild"]; ok {
		m.HasChild = parseBoolValue(v)
	}

	if v, ok := o["isCustomGroup"]; ok {
		m.IsCustomGroup = parseBoolValue(v)
	}

	if v, ok := o["domainType"]; ok {
		m.DomainType = parseStringValue(v)
	}

	if v, ok := o["domain"]; ok {
		m.Domain = m.Domain.flattenEndpointsGroupsAdGroupsDataDomain(ctx, v, diags)
	}

	return m
}

func (s *datasourceEndpointsGroupsAdGroupsModel) flattenEndpointsGroupsAdGroupsDataList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointsGroupsAdGroupsDataModel {
	if o == nil {
		return []datasourceEndpointsGroupsAdGroupsDataModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument data is not type of []interface{}.", "")
		return []datasourceEndpointsGroupsAdGroupsDataModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointsGroupsAdGroupsDataModel{}
	}

	values := make([]datasourceEndpointsGroupsAdGroupsDataModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointsGroupsAdGroupsDataModel
		values[i] = *m.flattenEndpointsGroupsAdGroupsData(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointsGroupsAdGroupsDataDomainModel) flattenEndpointsGroupsAdGroupsDataDomain(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupsAdGroupsDataDomainModel {
	if input == nil {
		return &datasourceEndpointsGroupsAdGroupsDataDomainModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupsAdGroupsDataDomainModel{}
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

func (m *datasourceEndpointsGroupsNonAdGroupsModel) flattenEndpointsGroupsNonAdGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupsNonAdGroupsModel {
	if input == nil {
		return &datasourceEndpointsGroupsNonAdGroupsModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupsNonAdGroupsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["data"]; ok {
		m.Data = m.flattenEndpointsGroupsNonAdGroupsDataList(ctx, v, diags)
	}

	return m
}

func (m *datasourceEndpointsGroupsNonAdGroupsDataModel) flattenEndpointsGroupsNonAdGroupsData(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupsNonAdGroupsDataModel {
	if input == nil {
		return &datasourceEndpointsGroupsNonAdGroupsDataModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupsNonAdGroupsDataModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["parentId"]; ok {
		m.ParentId = parseFloat64Value(v)
	}

	if v, ok := o["guid"]; ok {
		m.Guid = parseStringValue(v)
	}

	if v, ok := o["path"]; ok {
		m.Path = parseStringValue(v)
	}

	if v, ok := o["hasChild"]; ok {
		m.HasChild = parseBoolValue(v)
	}

	if v, ok := o["isCustomGroup"]; ok {
		m.IsCustomGroup = parseBoolValue(v)
	}

	if v, ok := o["domainType"]; ok {
		m.DomainType = parseStringValue(v)
	}

	if v, ok := o["domain"]; ok {
		m.Domain = m.Domain.flattenEndpointsGroupsNonAdGroupsDataDomain(ctx, v, diags)
	}

	return m
}

func (s *datasourceEndpointsGroupsNonAdGroupsModel) flattenEndpointsGroupsNonAdGroupsDataList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointsGroupsNonAdGroupsDataModel {
	if o == nil {
		return []datasourceEndpointsGroupsNonAdGroupsDataModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument data is not type of []interface{}.", "")
		return []datasourceEndpointsGroupsNonAdGroupsDataModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointsGroupsNonAdGroupsDataModel{}
	}

	values := make([]datasourceEndpointsGroupsNonAdGroupsDataModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointsGroupsNonAdGroupsDataModel
		values[i] = *m.flattenEndpointsGroupsNonAdGroupsData(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointsGroupsNonAdGroupsDataDomainModel) flattenEndpointsGroupsNonAdGroupsDataDomain(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupsNonAdGroupsDataDomainModel {
	if input == nil {
		return &datasourceEndpointsGroupsNonAdGroupsDataDomainModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupsNonAdGroupsDataDomainModel{}
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
