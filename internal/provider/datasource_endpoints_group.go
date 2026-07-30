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
var _ datasource.DataSource = &datasourceEndpointsGroup{}

func newDatasourceEndpointsGroup() datasource.DataSource {
	return &datasourceEndpointsGroup{}
}

type datasourceEndpointsGroup struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointsGroupModel describes the datasource data model.
type datasourceEndpointsGroupModel struct {
	AdGroups    *datasourceEndpointsGroupAdGroupsModel    `tfsdk:"ad_groups"`
	NonAdGroups *datasourceEndpointsGroupNonAdGroupsModel `tfsdk:"non_ad_groups"`
	Guid        types.String                              `tfsdk:"guid"`
	Offset      types.Float64                             `tfsdk:"offset"`
	PrimaryKey  types.String                              `tfsdk:"primary_key"`
}

func (r *datasourceEndpointsGroup) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_group"
}

func (r *datasourceEndpointsGroup) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Endpoint Domain monitor API for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"guid": schema.StringAttribute{
				MarkdownDescription: "UID of the group to expand to find child groups.",
				Optional:            true,
			},
			"offset": schema.Float64Attribute{
				MarkdownDescription: "Specifies the starting position of AD groups. Based on this the results will be seperated in AD groups and non AD groups, with AD groups containing a \"total\" count.",
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
					},
					"data": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.Float64Attribute{
									MarkdownDescription: "Id of the group.",
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: "Name of the group.",
									Computed:            true,
								},
								"parent_id": schema.Float64Attribute{
									MarkdownDescription: "Parent id of the group.",
									Computed:            true,
								},
								"guid": schema.StringAttribute{
									MarkdownDescription: "UID of the group.",
									Computed:            true,
								},
								"path": schema.StringAttribute{
									MarkdownDescription: "Path of the group.",
									Computed:            true,
								},
								"has_child": schema.BoolAttribute{
									MarkdownDescription: "Indicate if the group has child or not.",
									Computed:            true,
								},
								"is_custom_group": schema.BoolAttribute{
									MarkdownDescription: "Indicate if the group is custom group or not.",
									Computed:            true,
								},
								"domain_type": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("azure", "adfs"),
									},
									MarkdownDescription: "Type of the endpint/domains entry the group belongs to.\nSupported values: azure, adfs.",
									Computed:            true,
								},
								"domain": schema.SingleNestedAttribute{
									MarkdownDescription: "Reference of the endpoint/domains entry the group belongs to.",
									Attributes: map[string]schema.Attribute{
										"primary_key": schema.StringAttribute{
											Computed: true,
										},
										"datasource": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidatorwarning.OneOf("endpoint/domains"),
											},
											Computed: true,
										},
									},
									Computed: true,
								},
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"non_ad_groups": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"data": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.Float64Attribute{
									MarkdownDescription: "Id of the group.",
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: "Name of the group.",
									Computed:            true,
								},
								"parent_id": schema.Float64Attribute{
									MarkdownDescription: "Parent id of the group.",
									Computed:            true,
								},
								"guid": schema.StringAttribute{
									MarkdownDescription: "UID of the group.",
									Computed:            true,
								},
								"path": schema.StringAttribute{
									MarkdownDescription: "Path of the group.",
									Computed:            true,
								},
								"has_child": schema.BoolAttribute{
									MarkdownDescription: "Indicate if the group has child or not.",
									Computed:            true,
								},
								"is_custom_group": schema.BoolAttribute{
									MarkdownDescription: "Indicate if the group is custom group or not.",
									Computed:            true,
								},
								"domain_type": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("azure", "adfs"),
									},
									MarkdownDescription: "Type of the endpint/domains entry the group belongs to.\nSupported values: azure, adfs.",
									Computed:            true,
								},
								"domain": schema.SingleNestedAttribute{
									MarkdownDescription: "Reference of the endpoint/domains entry the group belongs to.",
									Attributes: map[string]schema.Attribute{
										"primary_key": schema.StringAttribute{
											Computed: true,
										},
										"datasource": schema.StringAttribute{
											Validators: []validator.String{
												stringvalidatorwarning.OneOf("endpoint/domains"),
											},
											Computed: true,
										},
									},
									Computed: true,
								},
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceEndpointsGroup) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoints_group"
}

func (r *datasourceEndpointsGroup) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointsGroupModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointsGroup(ctx, "read", diags))

	read_output, err := c.ReadEndpointsGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointsGroup(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointsGroupModel) refreshEndpointsGroup(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["adGroups"]; ok {
		m.AdGroups = m.AdGroups.flattenEndpointsGroupAdGroups(ctx, v, &diags)
	}

	if v, ok := o["nonAdGroups"]; ok {
		m.NonAdGroups = m.NonAdGroups.flattenEndpointsGroupNonAdGroups(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointsGroupModel) getURLObjectEndpointsGroup(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Guid.IsNull() && !data.Guid.IsUnknown() {
		result["guid"] = data.Guid.ValueString()
	}

	if !data.Offset.IsNull() && !data.Offset.IsUnknown() {
		result["offset"] = data.Offset.ValueFloat64()
	}

	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointsGroupAdGroupsModel struct {
	Data  []datasourceEndpointsGroupAdGroupsDataModel `tfsdk:"data"`
	Total types.Float64                               `tfsdk:"total"`
}

type datasourceEndpointsGroupAdGroupsDataModel struct {
	Id            types.Float64                                    `tfsdk:"id"`
	Name          types.String                                     `tfsdk:"name"`
	ParentId      types.Float64                                    `tfsdk:"parent_id"`
	Guid          types.String                                     `tfsdk:"guid"`
	Path          types.String                                     `tfsdk:"path"`
	HasChild      types.Bool                                       `tfsdk:"has_child"`
	IsCustomGroup types.Bool                                       `tfsdk:"is_custom_group"`
	DomainType    types.String                                     `tfsdk:"domain_type"`
	Domain        *datasourceEndpointsGroupAdGroupsDataDomainModel `tfsdk:"domain"`
}

type datasourceEndpointsGroupAdGroupsDataDomainModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointsGroupNonAdGroupsModel struct {
	Data []datasourceEndpointsGroupNonAdGroupsDataModel `tfsdk:"data"`
}

type datasourceEndpointsGroupNonAdGroupsDataModel struct {
	Id            types.Float64                                       `tfsdk:"id"`
	Name          types.String                                        `tfsdk:"name"`
	ParentId      types.Float64                                       `tfsdk:"parent_id"`
	Guid          types.String                                        `tfsdk:"guid"`
	Path          types.String                                        `tfsdk:"path"`
	HasChild      types.Bool                                          `tfsdk:"has_child"`
	IsCustomGroup types.Bool                                          `tfsdk:"is_custom_group"`
	DomainType    types.String                                        `tfsdk:"domain_type"`
	Domain        *datasourceEndpointsGroupNonAdGroupsDataDomainModel `tfsdk:"domain"`
}

type datasourceEndpointsGroupNonAdGroupsDataDomainModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceEndpointsGroupAdGroupsModel) flattenEndpointsGroupAdGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupAdGroupsModel {
	if input == nil {
		return &datasourceEndpointsGroupAdGroupsModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupAdGroupsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["data"]; ok {
		m.Data = m.flattenEndpointsGroupAdGroupsDataList(ctx, v, diags)
	}

	if v, ok := o["total"]; ok {
		m.Total = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceEndpointsGroupAdGroupsDataModel) flattenEndpointsGroupAdGroupsData(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupAdGroupsDataModel {
	if input == nil {
		return &datasourceEndpointsGroupAdGroupsDataModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupAdGroupsDataModel{}
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
		m.Domain = m.Domain.flattenEndpointsGroupAdGroupsDataDomain(ctx, v, diags)
	}

	return m
}

func (s *datasourceEndpointsGroupAdGroupsModel) flattenEndpointsGroupAdGroupsDataList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointsGroupAdGroupsDataModel {
	if o == nil {
		return []datasourceEndpointsGroupAdGroupsDataModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument data is not type of []interface{}.", "")
		return []datasourceEndpointsGroupAdGroupsDataModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointsGroupAdGroupsDataModel{}
	}

	values := make([]datasourceEndpointsGroupAdGroupsDataModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointsGroupAdGroupsDataModel
		if i < len(s.Data) {
			m = s.Data[i]
		}
		values[i] = *m.flattenEndpointsGroupAdGroupsData(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointsGroupAdGroupsDataDomainModel) flattenEndpointsGroupAdGroupsDataDomain(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupAdGroupsDataDomainModel {
	if input == nil {
		return &datasourceEndpointsGroupAdGroupsDataDomainModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupAdGroupsDataDomainModel{}
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

func (m *datasourceEndpointsGroupNonAdGroupsModel) flattenEndpointsGroupNonAdGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupNonAdGroupsModel {
	if input == nil {
		return &datasourceEndpointsGroupNonAdGroupsModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupNonAdGroupsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["data"]; ok {
		m.Data = m.flattenEndpointsGroupNonAdGroupsDataList(ctx, v, diags)
	}

	return m
}

func (m *datasourceEndpointsGroupNonAdGroupsDataModel) flattenEndpointsGroupNonAdGroupsData(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupNonAdGroupsDataModel {
	if input == nil {
		return &datasourceEndpointsGroupNonAdGroupsDataModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupNonAdGroupsDataModel{}
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
		m.Domain = m.Domain.flattenEndpointsGroupNonAdGroupsDataDomain(ctx, v, diags)
	}

	return m
}

func (s *datasourceEndpointsGroupNonAdGroupsModel) flattenEndpointsGroupNonAdGroupsDataList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointsGroupNonAdGroupsDataModel {
	if o == nil {
		return []datasourceEndpointsGroupNonAdGroupsDataModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument data is not type of []interface{}.", "")
		return []datasourceEndpointsGroupNonAdGroupsDataModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointsGroupNonAdGroupsDataModel{}
	}

	values := make([]datasourceEndpointsGroupNonAdGroupsDataModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointsGroupNonAdGroupsDataModel
		if i < len(s.Data) {
			m = s.Data[i]
		}
		values[i] = *m.flattenEndpointsGroupNonAdGroupsData(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointsGroupNonAdGroupsDataDomainModel) flattenEndpointsGroupNonAdGroupsDataDomain(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsGroupNonAdGroupsDataDomainModel {
	if input == nil {
		return &datasourceEndpointsGroupNonAdGroupsDataDomainModel{}
	}
	if m == nil {
		m = &datasourceEndpointsGroupNonAdGroupsDataDomainModel{}
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
