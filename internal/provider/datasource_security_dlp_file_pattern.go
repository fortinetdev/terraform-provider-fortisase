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
var _ datasource.DataSource = &datasourceSecurityDlpFilePattern{}

func newDatasourceSecurityDlpFilePattern() datasource.DataSource {
	return &datasourceSecurityDlpFilePattern{}
}

type datasourceSecurityDlpFilePattern struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityDlpFilePatternModel describes the datasource data model.
type datasourceSecurityDlpFilePatternModel struct {
	PrimaryKey types.String                                   `tfsdk:"primary_key"`
	Tag        types.String                                   `tfsdk:"tag"`
	Entries    []datasourceSecurityDlpFilePatternEntriesModel `tfsdk:"entries"`
}

func (r *datasourceSecurityDlpFilePattern) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_file_pattern"
}

func (r *datasourceSecurityDlpFilePattern) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DLP File Pattern Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"tag": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 63),
				},
				Computed: true,
			},
			"entries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pattern": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
						},
						"filter_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("type", "pattern"),
							},
							Computed: true,
						},
						"file_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("7z", "arj", "cab", "lzh", "rar", "tar", "zip", "bzip", "gzip", "bzip2", "xz", "bat", "uue", "mime", "base64", "binhex", "elf", "exe", "hta", "html", "jad", "class", "cod", "javascript", "msoffice", "msofficex", "fsg", "upx", "petite", "aspack", "sis", "hlp", "activemime", "jpeg", "gif", "tiff", "png", "bmp", "unknown", "mpeg", "mov", "mp3", "wma", "wav", "pdf", "avi", "rm", "torrent", "hibun", "msi", "mach-o", "dmg", ".net", "xar", "chm", "iso", "crx", "flac"),
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

func (r *datasourceSecurityDlpFilePattern) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_file_pattern"
}

func (r *datasourceSecurityDlpFilePattern) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityDlpFilePatternModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpFilePattern(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFilePatterns(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFilePattern(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityDlpFilePatternModel) refreshSecurityDlpFilePattern(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["tag"]; ok {
		m.Tag = parseStringValue(v)
	}

	if v, ok := o["entries"]; ok {
		m.Entries = m.flattenSecurityDlpFilePatternEntriesList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityDlpFilePatternModel) getURLObjectSecurityDlpFilePattern(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityDlpFilePatternEntriesModel struct {
	Pattern    types.String `tfsdk:"pattern"`
	FilterType types.String `tfsdk:"filter_type"`
	FileType   types.String `tfsdk:"file_type"`
}

func (m *datasourceSecurityDlpFilePatternEntriesModel) flattenSecurityDlpFilePatternEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpFilePatternEntriesModel {
	if input == nil {
		return &datasourceSecurityDlpFilePatternEntriesModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpFilePatternEntriesModel{}
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

func (s *datasourceSecurityDlpFilePatternModel) flattenSecurityDlpFilePatternEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDlpFilePatternEntriesModel {
	if o == nil {
		return []datasourceSecurityDlpFilePatternEntriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument entries is not type of []interface{}.", "")
		return []datasourceSecurityDlpFilePatternEntriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDlpFilePatternEntriesModel{}
	}

	values := make([]datasourceSecurityDlpFilePatternEntriesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDlpFilePatternEntriesModel
		if i < len(s.Entries) {
			m = s.Entries[i]
		}
		values[i] = *m.flattenSecurityDlpFilePatternEntries(ctx, ele, diags)
	}

	return values
}
