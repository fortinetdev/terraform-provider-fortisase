// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityApplications{}

func newDatasourceSecurityApplications() datasource.DataSource {
	return &datasourceSecurityApplications{}
}

type datasourceSecurityApplications struct {
	fortiClient *FortiClient
}

// datasourceSecurityApplicationsModel describes the datasource data model.
type datasourceSecurityApplicationsModel struct {
	PrimaryKey                types.String  `tfsdk:"primary_key"`
	Ftntid                    types.Float64 `tfsdk:"ftntid"`
	Category                  types.Float64 `tfsdk:"category"`
	Protocol                  types.String  `tfsdk:"protocol"`
	Popularity                types.Float64 `tfsdk:"popularity"`
	Risk                      types.Float64 `tfsdk:"risk"`
	Behavior                  types.Set     `tfsdk:"behavior"`
	Technology                types.String  `tfsdk:"technology"`
	Vendor                    types.String  `tfsdk:"vendor"`
	IconClass                 types.String  `tfsdk:"icon_class"`
	IsCloudApplication        types.Bool    `tfsdk:"is_cloud_application"`
	RequiresSslDeepInspection types.Bool    `tfsdk:"requires_ssl_deep_inspection"`
	IsDeepInspectionApp       types.Bool    `tfsdk:"is_deep_inspection_app"`
}

func (r *datasourceSecurityApplications) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_applications"
}

func (r *datasourceSecurityApplications) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"ftntid": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"category": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"protocol": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"popularity": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(1, 5),
				},
				Computed: true,
				Optional: true,
			},
			"risk": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(1, 5),
				},
				Computed: true,
				Optional: true,
			},
			"behavior": schema.SetAttribute{
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("", "Cloud", "Excessive-Bandwidth", "Botnet", "Tunneling", "Evasive"),
					),
				},
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"technology": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"vendor": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"icon_class": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"is_cloud_application": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"requires_ssl_deep_inspection": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"is_deep_inspection_app": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityApplications) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityApplications) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityApplicationsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityApplications(ctx, "read", diags))

	read_output, err := c.ReadSecurityApplications(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityApplications(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityApplicationsModel) refreshSecurityApplications(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["id"]; ok {
		m.Ftntid = parseFloat64Value(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = parseFloat64Value(v)
	}

	if v, ok := o["protocol"]; ok {
		m.Protocol = parseStringValue(v)
	}

	if v, ok := o["popularity"]; ok {
		m.Popularity = parseFloat64Value(v)
	}

	if v, ok := o["risk"]; ok {
		m.Risk = parseFloat64Value(v)
	}

	if v, ok := o["behavior"]; ok {
		m.Behavior = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["technology"]; ok {
		m.Technology = parseStringValue(v)
	}

	if v, ok := o["vendor"]; ok {
		m.Vendor = parseStringValue(v)
	}

	if v, ok := o["iconClass"]; ok {
		m.IconClass = parseStringValue(v)
	}

	if v, ok := o["isCloudApplication"]; ok {
		m.IsCloudApplication = parseBoolValue(v)
	}

	if v, ok := o["requiresSslDeepInspection"]; ok {
		m.RequiresSslDeepInspection = parseBoolValue(v)
	}

	if v, ok := o["isDeepInspectionApp"]; ok {
		m.IsDeepInspectionApp = parseBoolValue(v)
	}

	return diags
}

func (data *datasourceSecurityApplicationsModel) getURLObjectSecurityApplications(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
