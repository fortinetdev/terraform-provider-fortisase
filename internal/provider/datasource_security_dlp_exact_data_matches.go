// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityDlpExactDataMatches{}

func newDatasourceSecurityDlpExactDataMatches() datasource.DataSource {
	return &datasourceSecurityDlpExactDataMatches{}
}

type datasourceSecurityDlpExactDataMatches struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityDlpExactDataMatchesModel describes the datasource data model.
type datasourceSecurityDlpExactDataMatchesModel struct {
	PrimaryKey           types.String                                                    `tfsdk:"primary_key"`
	ExternalResourceData *datasourceSecurityDlpExactDataMatchesExternalResourceDataModel `tfsdk:"external_resource_data"`
	Columns              []datasourceSecurityDlpExactDataMatchesColumnsModel             `tfsdk:"columns"`
	OptionalCount        types.Float64                                                   `tfsdk:"optional_count"`
}

func (r *datasourceSecurityDlpExactDataMatches) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_exact_data_matches"
}

func (r *datasourceSecurityDlpExactDataMatches) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"optional_count": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(32),
				},
				Computed: true,
				Optional: true,
			},
			"external_resource_data": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"resource": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"refresh_rate": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.Between(1, 43200),
						},
						Computed: true,
						Optional: true,
					},
					"username": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(64),
						},
						Computed: true,
						Optional: true,
					},
					"password": schema.StringAttribute{
						Sensitive: true,
						Computed:  true,
						Optional:  true,
					},
					"update_method": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("feed", "push"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"columns": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"index": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validator.Between(1, 32),
							},
							Computed: true,
							Optional: true,
						},
						"optional": schema.BoolAttribute{
							Computed: true,
							Optional: true,
						},
						"type": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("security/dlp-data-types"),
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

func (r *datasourceSecurityDlpExactDataMatches) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_exact_data_matches"
}

func (r *datasourceSecurityDlpExactDataMatches) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityDlpExactDataMatchesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatches(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpExactDataMatches(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpExactDataMatches(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityDlpExactDataMatchesModel) refreshSecurityDlpExactDataMatches(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["externalResourceData"]; ok {
		m.ExternalResourceData = m.ExternalResourceData.flattenSecurityDlpExactDataMatchesExternalResourceData(ctx, v, &diags)
	}

	if v, ok := o["columns"]; ok {
		m.Columns = m.flattenSecurityDlpExactDataMatchesColumnsList(ctx, v, &diags)
	}

	if v, ok := o["optionalCount"]; ok {
		m.OptionalCount = parseFloat64Value(v)
	}

	return diags
}

func (data *datasourceSecurityDlpExactDataMatchesModel) getURLObjectSecurityDlpExactDataMatches(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityDlpExactDataMatchesExternalResourceDataModel struct {
	Resource     types.String  `tfsdk:"resource"`
	RefreshRate  types.Float64 `tfsdk:"refresh_rate"`
	Username     types.String  `tfsdk:"username"`
	Password     types.String  `tfsdk:"password"`
	UpdateMethod types.String  `tfsdk:"update_method"`
}

type datasourceSecurityDlpExactDataMatchesColumnsModel struct {
	Index    types.Float64                                          `tfsdk:"index"`
	Type     *datasourceSecurityDlpExactDataMatchesColumnsTypeModel `tfsdk:"type"`
	Optional types.Bool                                             `tfsdk:"optional"`
}

type datasourceSecurityDlpExactDataMatchesColumnsTypeModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityDlpExactDataMatchesExternalResourceDataModel) flattenSecurityDlpExactDataMatchesExternalResourceData(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpExactDataMatchesExternalResourceDataModel {
	if input == nil {
		return &datasourceSecurityDlpExactDataMatchesExternalResourceDataModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpExactDataMatchesExternalResourceDataModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["resource"]; ok {
		m.Resource = parseStringValue(v)
	}

	if v, ok := o["refreshRate"]; ok {
		m.RefreshRate = parseFloat64Value(v)
	}

	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	if v, ok := o["updateMethod"]; ok {
		m.UpdateMethod = parseStringValue(v)
	}

	return m
}

func (m *datasourceSecurityDlpExactDataMatchesColumnsModel) flattenSecurityDlpExactDataMatchesColumns(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpExactDataMatchesColumnsModel {
	if input == nil {
		return &datasourceSecurityDlpExactDataMatchesColumnsModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpExactDataMatchesColumnsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["index"]; ok {
		m.Index = parseFloat64Value(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = m.Type.flattenSecurityDlpExactDataMatchesColumnsType(ctx, v, diags)
	}

	if v, ok := o["optional"]; ok {
		m.Optional = parseBoolValue(v)
	}

	return m
}

func (s *datasourceSecurityDlpExactDataMatchesModel) flattenSecurityDlpExactDataMatchesColumnsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDlpExactDataMatchesColumnsModel {
	if o == nil {
		return []datasourceSecurityDlpExactDataMatchesColumnsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument columns is not type of []interface{}.", "")
		return []datasourceSecurityDlpExactDataMatchesColumnsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDlpExactDataMatchesColumnsModel{}
	}

	values := make([]datasourceSecurityDlpExactDataMatchesColumnsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDlpExactDataMatchesColumnsModel
		values[i] = *m.flattenSecurityDlpExactDataMatchesColumns(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityDlpExactDataMatchesColumnsTypeModel) flattenSecurityDlpExactDataMatchesColumnsType(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpExactDataMatchesColumnsTypeModel {
	if input == nil {
		return &datasourceSecurityDlpExactDataMatchesColumnsTypeModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpExactDataMatchesColumnsTypeModel{}
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
