// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointsSoftwareOnClientUser{}

func newDatasourceEndpointsSoftwareOnClientUser() datasource.DataSource {
	return &datasourceEndpointsSoftwareOnClientUser{}
}

type datasourceEndpointsSoftwareOnClientUser struct {
	fortiClient *FortiClient
}

// datasourceEndpointsSoftwareOnClientUserModel describes the datasource data model.
type datasourceEndpointsSoftwareOnClientUserModel struct {
	Software     []datasourceEndpointsSoftwareOnClientUserSoftwareModel `tfsdk:"software"`
	ClientUserId types.Float64                                          `tfsdk:"client_user_id"`
}

func (r *datasourceEndpointsSoftwareOnClientUser) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_software_on_client_user"
}

func (r *datasourceEndpointsSoftwareOnClientUser) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_user_id": schema.Float64Attribute{
				Description: "The client user ID of the endpoint client",
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
				Required: true,
			},
			"software": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"vendor": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"version": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"icon": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"install_date": schema.StringAttribute{
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

func (r *datasourceEndpointsSoftwareOnClientUser) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceEndpointsSoftwareOnClientUser) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointsSoftwareOnClientUserModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ClientUserId.ValueFloat64()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointsSoftwareOnClientUser(ctx, "read", diags))

	read_output, err := c.ReadEndpointsSoftwareOnClientUser(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointsSoftwareOnClientUser(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointsSoftwareOnClientUserModel) refreshEndpointsSoftwareOnClientUser(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["software"]; ok {
		m.Software = m.flattenEndpointsSoftwareOnClientUserSoftwareList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointsSoftwareOnClientUserModel) getURLObjectEndpointsSoftwareOnClientUser(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ClientUserId.IsNull() {
		result["clientUserId"] = data.ClientUserId.ValueFloat64()
	}

	return &result
}

type datasourceEndpointsSoftwareOnClientUserSoftwareModel struct {
	Id          types.Float64 `tfsdk:"id"`
	Name        types.String  `tfsdk:"name"`
	Vendor      types.String  `tfsdk:"vendor"`
	Version     types.String  `tfsdk:"version"`
	Icon        types.String  `tfsdk:"icon"`
	InstallDate types.String  `tfsdk:"install_date"`
}

func (m *datasourceEndpointsSoftwareOnClientUserSoftwareModel) flattenEndpointsSoftwareOnClientUserSoftware(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsSoftwareOnClientUserSoftwareModel {
	if input == nil {
		return &datasourceEndpointsSoftwareOnClientUserSoftwareModel{}
	}
	if m == nil {
		m = &datasourceEndpointsSoftwareOnClientUserSoftwareModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["vendor"]; ok {
		m.Vendor = parseStringValue(v)
	}

	if v, ok := o["version"]; ok {
		m.Version = parseStringValue(v)
	}

	if v, ok := o["icon"]; ok {
		m.Icon = parseStringValue(v)
	}

	if v, ok := o["installDate"]; ok {
		m.InstallDate = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointsSoftwareOnClientUserModel) flattenEndpointsSoftwareOnClientUserSoftwareList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointsSoftwareOnClientUserSoftwareModel {
	if o == nil {
		return []datasourceEndpointsSoftwareOnClientUserSoftwareModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument software is not type of []interface{}.", "")
		return []datasourceEndpointsSoftwareOnClientUserSoftwareModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointsSoftwareOnClientUserSoftwareModel{}
	}

	values := make([]datasourceEndpointsSoftwareOnClientUserSoftwareModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointsSoftwareOnClientUserSoftwareModel
		values[i] = *m.flattenEndpointsSoftwareOnClientUserSoftware(ctx, ele, diags)
	}

	return values
}
