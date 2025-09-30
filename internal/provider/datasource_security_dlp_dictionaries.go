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
var _ datasource.DataSource = &datasourceSecurityDlpDictionaries{}

func newDatasourceSecurityDlpDictionaries() datasource.DataSource {
	return &datasourceSecurityDlpDictionaries{}
}

type datasourceSecurityDlpDictionaries struct {
	fortiClient *FortiClient
}

// datasourceSecurityDlpDictionariesModel describes the datasource data model.
type datasourceSecurityDlpDictionariesModel struct {
	PrimaryKey           types.String                                    `tfsdk:"primary_key"`
	DictionaryType       types.String                                    `tfsdk:"dictionary_type"`
	SensitivityLabelGuid types.String                                    `tfsdk:"sensitivity_label_guid"`
	EntriesToEvaluate    types.String                                    `tfsdk:"entries_to_evaluate"`
	Entries              []datasourceSecurityDlpDictionariesEntriesModel `tfsdk:"entries"`
}

func (r *datasourceSecurityDlpDictionaries) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_dictionaries"
}

func (r *datasourceSecurityDlpDictionaries) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"dictionary_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("sensor", "mip-label"),
				},
				Computed: true,
				Optional: true,
			},
			"sensitivity_label_guid": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"entries_to_evaluate": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("all", "any"),
				},
				Computed: true,
				Optional: true,
			},
			"entries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"repeat": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"pattern": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"case_sensitive": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"dlp_data_type": schema.SingleNestedAttribute{
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

func (r *datasourceSecurityDlpDictionaries) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityDlpDictionaries) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityDlpDictionariesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpDictionaries(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpDictionaries(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpDictionaries(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityDlpDictionariesModel) refreshSecurityDlpDictionaries(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["dictionaryType"]; ok {
		m.DictionaryType = parseStringValue(v)
	}

	if v, ok := o["sensitivityLabelGuid"]; ok {
		m.SensitivityLabelGuid = parseStringValue(v)
	}

	if v, ok := o["entriesToEvaluate"]; ok {
		m.EntriesToEvaluate = parseStringValue(v)
	}

	if v, ok := o["entries"]; ok {
		m.Entries = m.flattenSecurityDlpDictionariesEntriesList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityDlpDictionariesModel) getURLObjectSecurityDlpDictionaries(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityDlpDictionariesEntriesModel struct {
	DlpDataType   *datasourceSecurityDlpDictionariesEntriesDlpDataTypeModel `tfsdk:"dlp_data_type"`
	Status        types.String                                              `tfsdk:"status"`
	Repeat        types.String                                              `tfsdk:"repeat"`
	Pattern       types.String                                              `tfsdk:"pattern"`
	CaseSensitive types.String                                              `tfsdk:"case_sensitive"`
}

type datasourceSecurityDlpDictionariesEntriesDlpDataTypeModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityDlpDictionariesEntriesModel) flattenSecurityDlpDictionariesEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpDictionariesEntriesModel {
	if input == nil {
		return &datasourceSecurityDlpDictionariesEntriesModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpDictionariesEntriesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["dlpDataType"]; ok {
		m.DlpDataType = m.DlpDataType.flattenSecurityDlpDictionariesEntriesDlpDataType(ctx, v, diags)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["repeat"]; ok {
		m.Repeat = parseStringValue(v)
	}

	if v, ok := o["pattern"]; ok {
		m.Pattern = parseStringValue(v)
	}

	if v, ok := o["caseSensitive"]; ok {
		m.CaseSensitive = parseStringValue(v)
	}

	return m
}

func (s *datasourceSecurityDlpDictionariesModel) flattenSecurityDlpDictionariesEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDlpDictionariesEntriesModel {
	if o == nil {
		return []datasourceSecurityDlpDictionariesEntriesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument entries is not type of []interface{}.", "")
		return []datasourceSecurityDlpDictionariesEntriesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDlpDictionariesEntriesModel{}
	}

	values := make([]datasourceSecurityDlpDictionariesEntriesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDlpDictionariesEntriesModel
		values[i] = *m.flattenSecurityDlpDictionariesEntries(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityDlpDictionariesEntriesDlpDataTypeModel) flattenSecurityDlpDictionariesEntriesDlpDataType(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpDictionariesEntriesDlpDataTypeModel {
	if input == nil {
		return &datasourceSecurityDlpDictionariesEntriesDlpDataTypeModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpDictionariesEntriesDlpDataTypeModel{}
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
