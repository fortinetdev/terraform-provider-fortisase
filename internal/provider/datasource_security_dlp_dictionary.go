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
var _ datasource.DataSource = &datasourceSecurityDlpDictionary{}

func newDatasourceSecurityDlpDictionary() datasource.DataSource {
	return &datasourceSecurityDlpDictionary{}
}

type datasourceSecurityDlpDictionary struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityDlpDictionaryModel describes the datasource data model.
type datasourceSecurityDlpDictionaryModel struct {
	PrimaryKey        types.String                                  `tfsdk:"primary_key"`
	EntriesToEvaluate types.String                                  `tfsdk:"entries_to_evaluate"`
	DictionaryType    types.String                                  `tfsdk:"dictionary_type"`
	Entries           []datasourceSecurityDlpDictionaryEntriesModel `tfsdk:"entries"`
}

func (r *datasourceSecurityDlpDictionary) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_dictionary"
}

func (r *datasourceSecurityDlpDictionary) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DLP Dictionary Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 64),
				},
				Required: true,
			},
			"entries_to_evaluate": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("all", "any"),
				},
				Computed: true,
			},
			"dictionary_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("mip-label", "sensor"),
				},
				MarkdownDescription: "This property is used to classify DLP Dictionaries. It is server-generated and cannot be modified.\nSupported values: mip-label, sensor.",
				Computed:            true,
			},
			"entries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"repeat": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"pattern": schema.StringAttribute{
							Computed: true,
						},
						"case_sensitive": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"dlp_data_type": schema.SingleNestedAttribute{
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

func (r *datasourceSecurityDlpDictionary) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_dictionary"
}

func (r *datasourceSecurityDlpDictionary) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityDlpDictionaryModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpDictionary(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpDictionaries(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpDictionary(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityDlpDictionaryModel) refreshSecurityDlpDictionary(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["entriesToEvaluate"]; ok {
		m.EntriesToEvaluate = parseStringValue(v)
	}

	if v, ok := o["dictionaryType"]; ok {
		m.DictionaryType = parseStringValue(v)
	}

	if v, ok := o["entries"]; ok {
		m.Entries = m.flattenSecurityDlpDictionaryEntriesList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityDlpDictionaryModel) getURLObjectSecurityDlpDictionary(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityDlpDictionaryEntriesModel struct {
	DlpDataType   *datasourceSecurityDlpDictionaryEntriesDlpDataTypeModel `tfsdk:"dlp_data_type"`
	Status        types.String                                            `tfsdk:"status"`
	Repeat        types.String                                            `tfsdk:"repeat"`
	Pattern       types.String                                            `tfsdk:"pattern"`
	CaseSensitive types.String                                            `tfsdk:"case_sensitive"`
}

type datasourceSecurityDlpDictionaryEntriesDlpDataTypeModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityDlpDictionaryEntriesModel) flattenSecurityDlpDictionaryEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpDictionaryEntriesModel {
	if input == nil {
		return &datasourceSecurityDlpDictionaryEntriesModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpDictionaryEntriesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["dlpDataType"]; ok {
		m.DlpDataType = m.DlpDataType.flattenSecurityDlpDictionaryEntriesDlpDataType(ctx, v, diags)
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

func (s *datasourceSecurityDlpDictionaryModel) flattenSecurityDlpDictionaryEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDlpDictionaryEntriesModel {
	if o == nil {
		return []datasourceSecurityDlpDictionaryEntriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument entries is not type of []interface{}.", "")
		return []datasourceSecurityDlpDictionaryEntriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDlpDictionaryEntriesModel{}
	}

	values := make([]datasourceSecurityDlpDictionaryEntriesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDlpDictionaryEntriesModel
		if i < len(s.Entries) {
			m = s.Entries[i]
		}
		values[i] = *m.flattenSecurityDlpDictionaryEntries(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityDlpDictionaryEntriesDlpDataTypeModel) flattenSecurityDlpDictionaryEntriesDlpDataType(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpDictionaryEntriesDlpDataTypeModel {
	if input == nil {
		return &datasourceSecurityDlpDictionaryEntriesDlpDataTypeModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpDictionaryEntriesDlpDataTypeModel{}
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
