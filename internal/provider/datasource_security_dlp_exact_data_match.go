// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityDlpExactDataMatch{}

func newDatasourceSecurityDlpExactDataMatch() datasource.DataSource {
	return &datasourceSecurityDlpExactDataMatch{}
}

type datasourceSecurityDlpExactDataMatch struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityDlpExactDataMatchModel describes the datasource data model.
type datasourceSecurityDlpExactDataMatchModel struct {
	PrimaryKey           types.String                                                  `tfsdk:"primary_key"`
	ExternalResourceData *datasourceSecurityDlpExactDataMatchExternalResourceDataModel `tfsdk:"external_resource_data"`
	Columns              []datasourceSecurityDlpExactDataMatchColumnsModel             `tfsdk:"columns"`
	OptionalCount        types.Float64                                                 `tfsdk:"optional_count"`
}

func (r *datasourceSecurityDlpExactDataMatch) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_exact_data_match"
}

func (r *datasourceSecurityDlpExactDataMatch) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DLP Exact Data Match Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"optional_count": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(32),
				},
				Computed: true,
			},
			"external_resource_data": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"resource": schema.StringAttribute{
						Computed: true,
					},
					"refresh_rate": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validatorwarning.Between(1, 43200),
						},
						Computed: true,
					},
					"username": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.LengthAtMost(64),
						},
						Computed: true,
					},
					"password": schema.StringAttribute{
						Sensitive: true,
						Computed:  true,
					},
					"update_method": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("feed", "push"),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"columns": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"index": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(1, 32),
							},
							Computed: true,
						},
						"optional": schema.BoolAttribute{
							Computed: true,
						},
						"type": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("security/dlp-data-types"),
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
	}
}

func (r *datasourceSecurityDlpExactDataMatch) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_exact_data_match"
}

func (r *datasourceSecurityDlpExactDataMatch) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityDlpExactDataMatchModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpExactDataMatch(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpExactDataMatches(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpExactDataMatch(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityDlpExactDataMatchModel) refreshSecurityDlpExactDataMatch(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["externalResourceData"]; ok {
		m.ExternalResourceData = m.ExternalResourceData.flattenSecurityDlpExactDataMatchExternalResourceData(ctx, v, &diags)
	}

	if v, ok := o["columns"]; ok {
		m.Columns = m.flattenSecurityDlpExactDataMatchColumnsList(ctx, v, &diags)
	}

	if v, ok := o["optionalCount"]; ok {
		m.OptionalCount = parseFloat64Value(v)
	}

	return diags
}

func (data *datasourceSecurityDlpExactDataMatchModel) getURLObjectSecurityDlpExactDataMatch(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityDlpExactDataMatchExternalResourceDataModel struct {
	Resource     types.String  `tfsdk:"resource"`
	RefreshRate  types.Float64 `tfsdk:"refresh_rate"`
	Username     types.String  `tfsdk:"username"`
	Password     types.String  `tfsdk:"password"`
	UpdateMethod types.String  `tfsdk:"update_method"`
}

type datasourceSecurityDlpExactDataMatchColumnsModel struct {
	Index    types.Float64                                        `tfsdk:"index"`
	Type     *datasourceSecurityDlpExactDataMatchColumnsTypeModel `tfsdk:"type"`
	Optional types.Bool                                           `tfsdk:"optional"`
}

type datasourceSecurityDlpExactDataMatchColumnsTypeModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityDlpExactDataMatchExternalResourceDataModel) flattenSecurityDlpExactDataMatchExternalResourceData(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpExactDataMatchExternalResourceDataModel {
	if input == nil {
		return &datasourceSecurityDlpExactDataMatchExternalResourceDataModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpExactDataMatchExternalResourceDataModel{}
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

func (m *datasourceSecurityDlpExactDataMatchColumnsModel) flattenSecurityDlpExactDataMatchColumns(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpExactDataMatchColumnsModel {
	if input == nil {
		return &datasourceSecurityDlpExactDataMatchColumnsModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpExactDataMatchColumnsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["index"]; ok {
		m.Index = parseFloat64Value(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = m.Type.flattenSecurityDlpExactDataMatchColumnsType(ctx, v, diags)
	}

	if v, ok := o["optional"]; ok {
		m.Optional = parseBoolValue(v)
	}

	return m
}

func (s *datasourceSecurityDlpExactDataMatchModel) flattenSecurityDlpExactDataMatchColumnsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDlpExactDataMatchColumnsModel {
	if o == nil {
		return []datasourceSecurityDlpExactDataMatchColumnsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument columns is not type of []interface{}.", "")
		return []datasourceSecurityDlpExactDataMatchColumnsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDlpExactDataMatchColumnsModel{}
	}

	values := make([]datasourceSecurityDlpExactDataMatchColumnsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDlpExactDataMatchColumnsModel
		if i < len(s.Columns) {
			m = s.Columns[i]
		}
		values[i] = *m.flattenSecurityDlpExactDataMatchColumns(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityDlpExactDataMatchColumnsTypeModel) flattenSecurityDlpExactDataMatchColumnsType(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpExactDataMatchColumnsTypeModel {
	if input == nil {
		return &datasourceSecurityDlpExactDataMatchColumnsTypeModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpExactDataMatchColumnsTypeModel{}
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
