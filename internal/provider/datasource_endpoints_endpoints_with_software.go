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
var _ datasource.DataSource = &datasourceEndpointsEndpointsWithSoftware{}

func newDatasourceEndpointsEndpointsWithSoftware() datasource.DataSource {
	return &datasourceEndpointsEndpointsWithSoftware{}
}

type datasourceEndpointsEndpointsWithSoftware struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointsEndpointsWithSoftwareModel describes the datasource data model.
type datasourceEndpointsEndpointsWithSoftwareModel struct {
	Clients    []datasourceEndpointsEndpointsWithSoftwareClientsModel `tfsdk:"clients"`
	Total      types.Float64                                          `tfsdk:"total"`
	SoftwareId types.Float64                                          `tfsdk:"software_id"`
}

func (r *datasourceEndpointsEndpointsWithSoftware) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_endpoints_with_software"
}

func (r *datasourceEndpointsEndpointsWithSoftware) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"total": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"software_id": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
				MarkdownDescription: "The ID property of a specific software.\nValue at least 1.",
				Required:            true,
			},
			"clients": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"client_user_id": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
						"client_id": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
						"app_count": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
						"last_install": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"device_id": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
						"device_ip": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"device_host": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"device_os": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"user_id": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
						"user_name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"user_icon": schema.StringAttribute{
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

func (r *datasourceEndpointsEndpointsWithSoftware) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoints_endpoints_with_software"
}

func (r *datasourceEndpointsEndpointsWithSoftware) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointsEndpointsWithSoftwareModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.SoftwareId.ValueFloat64()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointsEndpointsWithSoftware(ctx, "read", diags))

	read_output, err := c.ReadEndpointsEndpointsWithSoftware(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointsEndpointsWithSoftware(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointsEndpointsWithSoftwareModel) refreshEndpointsEndpointsWithSoftware(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["clients"]; ok {
		m.Clients = m.flattenEndpointsEndpointsWithSoftwareClientsList(ctx, v, &diags)
	}

	if v, ok := o["total"]; ok {
		m.Total = parseFloat64Value(v)
	}

	return diags
}

func (data *datasourceEndpointsEndpointsWithSoftwareModel) getURLObjectEndpointsEndpointsWithSoftware(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SoftwareId.IsNull() {
		result["softwareId"] = data.SoftwareId.ValueFloat64()
	}

	return &result
}

type datasourceEndpointsEndpointsWithSoftwareClientsModel struct {
	ClientUserId types.Float64 `tfsdk:"client_user_id"`
	ClientId     types.Float64 `tfsdk:"client_id"`
	AppCount     types.Float64 `tfsdk:"app_count"`
	LastInstall  types.String  `tfsdk:"last_install"`
	DeviceId     types.Float64 `tfsdk:"device_id"`
	DeviceIp     types.String  `tfsdk:"device_ip"`
	DeviceHost   types.String  `tfsdk:"device_host"`
	DeviceOs     types.String  `tfsdk:"device_os"`
	UserId       types.Float64 `tfsdk:"user_id"`
	UserName     types.String  `tfsdk:"user_name"`
	UserIcon     types.String  `tfsdk:"user_icon"`
}

func (m *datasourceEndpointsEndpointsWithSoftwareClientsModel) flattenEndpointsEndpointsWithSoftwareClients(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsEndpointsWithSoftwareClientsModel {
	if input == nil {
		return &datasourceEndpointsEndpointsWithSoftwareClientsModel{}
	}
	if m == nil {
		m = &datasourceEndpointsEndpointsWithSoftwareClientsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["clientUserId"]; ok {
		m.ClientUserId = parseFloat64Value(v)
	}

	if v, ok := o["clientId"]; ok {
		m.ClientId = parseFloat64Value(v)
	}

	if v, ok := o["appCount"]; ok {
		m.AppCount = parseFloat64Value(v)
	}

	if v, ok := o["lastInstall"]; ok {
		m.LastInstall = parseStringValue(v)
	}

	if v, ok := o["deviceId"]; ok {
		m.DeviceId = parseFloat64Value(v)
	}

	if v, ok := o["deviceIp"]; ok {
		m.DeviceIp = parseStringValue(v)
	}

	if v, ok := o["deviceHost"]; ok {
		m.DeviceHost = parseStringValue(v)
	}

	if v, ok := o["deviceOs"]; ok {
		m.DeviceOs = parseStringValue(v)
	}

	if v, ok := o["userId"]; ok {
		m.UserId = parseFloat64Value(v)
	}

	if v, ok := o["userName"]; ok {
		m.UserName = parseStringValue(v)
	}

	if v, ok := o["userIcon"]; ok {
		m.UserIcon = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointsEndpointsWithSoftwareModel) flattenEndpointsEndpointsWithSoftwareClientsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointsEndpointsWithSoftwareClientsModel {
	if o == nil {
		return []datasourceEndpointsEndpointsWithSoftwareClientsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument clients is not type of []interface{}.", "")
		return []datasourceEndpointsEndpointsWithSoftwareClientsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointsEndpointsWithSoftwareClientsModel{}
	}

	values := make([]datasourceEndpointsEndpointsWithSoftwareClientsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointsEndpointsWithSoftwareClientsModel
		values[i] = *m.flattenEndpointsEndpointsWithSoftwareClients(ctx, ele, diags)
	}

	return values
}
