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
var _ datasource.DataSource = &datasourceSecurityServiceGroup{}

func newDatasourceSecurityServiceGroup() datasource.DataSource {
	return &datasourceSecurityServiceGroup{}
}

type datasourceSecurityServiceGroup struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityServiceGroupModel describes the datasource data model.
type datasourceSecurityServiceGroupModel struct {
	PrimaryKey types.String                                 `tfsdk:"primary_key"`
	Proxy      types.Bool                                   `tfsdk:"proxy"`
	Members    []datasourceSecurityServiceGroupMembersModel `tfsdk:"members"`
}

func (r *datasourceSecurityServiceGroup) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_service_group"
}

func (r *datasourceSecurityServiceGroup) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Service Group Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(79),
				},
				Required: true,
			},
			"proxy": schema.BoolAttribute{
				Computed: true,
			},
			"members": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("security/services", "security/service-groups"),
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceSecurityServiceGroup) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_service_group"
}

func (r *datasourceSecurityServiceGroup) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityServiceGroupModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityServiceGroup(ctx, "read", diags))

	read_output, err := c.ReadSecurityServiceGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityServiceGroup(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityServiceGroupModel) refreshSecurityServiceGroup(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["proxy"]; ok {
		m.Proxy = parseBoolValue(v)
	}

	if v, ok := o["members"]; ok {
		m.Members = m.flattenSecurityServiceGroupMembersList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityServiceGroupModel) getURLObjectSecurityServiceGroup(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityServiceGroupMembersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityServiceGroupMembersModel) flattenSecurityServiceGroupMembers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityServiceGroupMembersModel {
	if input == nil {
		return &datasourceSecurityServiceGroupMembersModel{}
	}
	if m == nil {
		m = &datasourceSecurityServiceGroupMembersModel{}
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

func (s *datasourceSecurityServiceGroupModel) flattenSecurityServiceGroupMembersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityServiceGroupMembersModel {
	if o == nil {
		return []datasourceSecurityServiceGroupMembersModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument members is not type of []interface{}.", "")
		return []datasourceSecurityServiceGroupMembersModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityServiceGroupMembersModel{}
	}

	values := make([]datasourceSecurityServiceGroupMembersModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityServiceGroupMembersModel
		if i < len(s.Members) {
			m = s.Members[i]
		}
		values[i] = *m.flattenSecurityServiceGroupMembers(ctx, ele, diags)
	}

	return values
}
