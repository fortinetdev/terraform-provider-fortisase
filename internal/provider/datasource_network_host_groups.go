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
var _ datasource.DataSource = &datasourceNetworkHostGroups{}

func newDatasourceNetworkHostGroups() datasource.DataSource {
	return &datasourceNetworkHostGroups{}
}

type datasourceNetworkHostGroups struct {
	fortiClient *FortiClient
}

// datasourceNetworkHostGroupsModel describes the datasource data model.
type datasourceNetworkHostGroupsModel struct {
	PrimaryKey types.String                              `tfsdk:"primary_key"`
	Members    []datasourceNetworkHostGroupsMembersModel `tfsdk:"members"`
}

func (r *datasourceNetworkHostGroups) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_host_groups"
}

func (r *datasourceNetworkHostGroups) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 79),
				},
				Required: true,
			},
			"members": schema.ListNestedAttribute{
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
	}
}

func (r *datasourceNetworkHostGroups) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceNetworkHostGroups) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceNetworkHostGroupsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkHostGroups(ctx, "read", diags))

	read_output, err := c.ReadNetworkHostGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkHostGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceNetworkHostGroupsModel) refreshNetworkHostGroups(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["members"]; ok {
		m.Members = m.flattenNetworkHostGroupsMembersList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceNetworkHostGroupsModel) getURLObjectNetworkHostGroups(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceNetworkHostGroupsMembersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceNetworkHostGroupsMembersModel) flattenNetworkHostGroupsMembers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceNetworkHostGroupsMembersModel {
	if input == nil {
		return &datasourceNetworkHostGroupsMembersModel{}
	}
	if m == nil {
		m = &datasourceNetworkHostGroupsMembersModel{}
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

func (s *datasourceNetworkHostGroupsModel) flattenNetworkHostGroupsMembersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceNetworkHostGroupsMembersModel {
	if o == nil {
		return []datasourceNetworkHostGroupsMembersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument members is not type of []interface{}.", "")
		return []datasourceNetworkHostGroupsMembersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceNetworkHostGroupsMembersModel{}
	}

	values := make([]datasourceNetworkHostGroupsMembersModel, len(l))
	for i, ele := range l {
		var m datasourceNetworkHostGroupsMembersModel
		values[i] = *m.flattenNetworkHostGroupsMembers(ctx, ele, diags)
	}

	return values
}
