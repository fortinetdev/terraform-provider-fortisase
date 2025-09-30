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
var _ datasource.DataSource = &datasourceSecurityServiceGroups{}

func newDatasourceSecurityServiceGroups() datasource.DataSource {
	return &datasourceSecurityServiceGroups{}
}

type datasourceSecurityServiceGroups struct {
	fortiClient *FortiClient
}

// datasourceSecurityServiceGroupsModel describes the datasource data model.
type datasourceSecurityServiceGroupsModel struct {
	PrimaryKey types.String                                  `tfsdk:"primary_key"`
	Proxy      types.Bool                                    `tfsdk:"proxy"`
	Members    []datasourceSecurityServiceGroupsMembersModel `tfsdk:"members"`
}

func (r *datasourceSecurityServiceGroups) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_service_groups"
}

func (r *datasourceSecurityServiceGroups) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(79),
				},
				Required: true,
			},
			"proxy": schema.BoolAttribute{
				Computed: true,
				Optional: true,
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
								stringvalidator.OneOf("security/services", "security/service-groups"),
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

func (r *datasourceSecurityServiceGroups) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityServiceGroups) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityServiceGroupsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityServiceGroups(ctx, "read", diags))

	read_output, err := c.ReadSecurityServiceGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityServiceGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityServiceGroupsModel) refreshSecurityServiceGroups(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["proxy"]; ok {
		m.Proxy = parseBoolValue(v)
	}

	if v, ok := o["members"]; ok {
		m.Members = m.flattenSecurityServiceGroupsMembersList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityServiceGroupsModel) getURLObjectSecurityServiceGroups(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityServiceGroupsMembersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityServiceGroupsMembersModel) flattenSecurityServiceGroupsMembers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceGroupsMembersModel {
	if input == nil {
		return &datasourceSecurityServiceGroupsMembersModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceGroupsMembersModel{}
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

func (s *datasourceSecurityServiceGroupsModel) flattenSecurityServiceGroupsMembersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityServiceGroupsMembersModel {
	if o == nil {
		return []datasourceSecurityServiceGroupsMembersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument members is not type of []interface{}.", "")
		return []datasourceSecurityServiceGroupsMembersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityServiceGroupsMembersModel{}
	}

	values := make([]datasourceSecurityServiceGroupsMembersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityServiceGroupsMembersModel
		values[i] = *m.flattenSecurityServiceGroupsMembers(ctx, ele, diags)
	}

	return values
}
