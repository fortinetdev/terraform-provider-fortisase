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
var _ datasource.DataSource = &datasourceSecurityDlpFilePatterns{}

func newDatasourceSecurityDlpFilePatterns() datasource.DataSource {
	return &datasourceSecurityDlpFilePatterns{}
}

type datasourceSecurityDlpFilePatterns struct {
	fortiClient *FortiClient
}

// datasourceSecurityDlpFilePatternsModel describes the datasource data model.
type datasourceSecurityDlpFilePatternsModel struct {
	PrimaryKey types.String                                    `tfsdk:"primary_key"`
	Tag        types.String                                    `tfsdk:"tag"`
	Entries    []datasourceSecurityDlpFilePatternsEntriesModel `tfsdk:"entries"`
}

func (r *datasourceSecurityDlpFilePatterns) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_file_patterns"
}

func (r *datasourceSecurityDlpFilePatterns) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"tag": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 63),
				},
				Computed: true,
				Optional: true,
			},
			"entries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pattern": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"filter_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("type", "pattern"),
							},
							Computed: true,
							Optional: true,
						},
						"file_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("7z", "arj", "cab", "lzh", "rar", "tar", "zip", "bzip", "gzip", "bzip2", "xz", "bat", "uue", "mime", "base64", "binhex", "elf", "exe", "hta", "html", "jad", "class", "cod", "javascript", "msoffice", "msofficex", "fsg", "upx", "petite", "aspack", "sis", "hlp", "activemime", "jpeg", "gif", "tiff", "png", "bmp", "unknown", "mpeg", "mov", "mp3", "wma", "wav", "pdf", "avi", "rm", "torrent", "hibun", "msi", "mach-o", "dmg", ".net", "xar", "chm", "iso", "crx", "flac"),
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

func (r *datasourceSecurityDlpFilePatterns) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityDlpFilePatterns) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityDlpFilePatternsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpFilePatterns(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFilePatterns(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFilePatterns(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityDlpFilePatternsModel) refreshSecurityDlpFilePatterns(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["tag"]; ok {
		m.Tag = parseStringValue(v)
	}

	if v, ok := o["entries"]; ok {
		m.Entries = m.flattenSecurityDlpFilePatternsEntriesList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityDlpFilePatternsModel) getURLObjectSecurityDlpFilePatterns(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityDlpFilePatternsEntriesModel struct {
	Pattern    types.String `tfsdk:"pattern"`
	FilterType types.String `tfsdk:"filter_type"`
	FileType   types.String `tfsdk:"file_type"`
}

func (m *datasourceSecurityDlpFilePatternsEntriesModel) flattenSecurityDlpFilePatternsEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpFilePatternsEntriesModel {
	if input == nil {
		return &datasourceSecurityDlpFilePatternsEntriesModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpFilePatternsEntriesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["pattern"]; ok {
		m.Pattern = parseStringValue(v)
	}

	if v, ok := o["filterType"]; ok {
		m.FilterType = parseStringValue(v)
	}

	if v, ok := o["fileType"]; ok {
		m.FileType = parseStringValue(v)
	}

	return m
}

func (s *datasourceSecurityDlpFilePatternsModel) flattenSecurityDlpFilePatternsEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDlpFilePatternsEntriesModel {
	if o == nil {
		return []datasourceSecurityDlpFilePatternsEntriesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument entries is not type of []interface{}.", "")
		return []datasourceSecurityDlpFilePatternsEntriesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDlpFilePatternsEntriesModel{}
	}

	values := make([]datasourceSecurityDlpFilePatternsEntriesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDlpFilePatternsEntriesModel
		values[i] = *m.flattenSecurityDlpFilePatternsEntries(ctx, ele, diags)
	}

	return values
}
