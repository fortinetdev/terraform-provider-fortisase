// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSecurityDlpFilePatterns{}

func newResourceSecurityDlpFilePatterns() resource.Resource {
	return &resourceSecurityDlpFilePatterns{}
}

type resourceSecurityDlpFilePatterns struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityDlpFilePatternsModel describes the resource data model.
type resourceSecurityDlpFilePatternsModel struct {
	ID         types.String                                  `tfsdk:"id"`
	PrimaryKey types.String                                  `tfsdk:"primary_key"`
	Tag        types.String                                  `tfsdk:"tag"`
	Entries    []resourceSecurityDlpFilePatternsEntriesModel `tfsdk:"entries"`
}

func (r *resourceSecurityDlpFilePatterns) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_file_patterns"
}

func (r *resourceSecurityDlpFilePatterns) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"primary_key": schema.StringAttribute{
				Computed: true,
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
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
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

func (r *resourceSecurityDlpFilePatterns) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_file_patterns"
}

func (r *resourceSecurityDlpFilePatterns) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpFilePatterns")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityDlpFilePatternsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityDlpFilePatterns(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpFilePatterns(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityDlpFilePatterns(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpFilePatterns(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFilePatterns(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFilePatterns(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpFilePatterns) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpFilePatterns")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDlpFilePatternsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDlpFilePatternsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDlpFilePatterns(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpFilePatterns(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityDlpFilePatterns(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpFilePatterns(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFilePatterns(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFilePatterns(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpFilePatterns) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpFilePatterns")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityDlpFilePatternsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpFilePatterns(ctx, "delete", diags))

	output, err := c.DeleteSecurityDlpFilePatterns(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityDlpFilePatterns) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpFilePatternsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpFilePatterns(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFilePatterns(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFilePatterns(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpFilePatterns) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityDlpFilePatternsModel) refreshSecurityDlpFilePatterns(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceSecurityDlpFilePatternsModel) getCreateObjectSecurityDlpFilePatterns(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})

	if !data.Tag.IsNull() {
		result["tag"] = data.Tag.ValueString()
	}

	result["entries"] = data.expandSecurityDlpFilePatternsEntriesList(ctx, data.Entries, diags)

	return &result
}

func (data *resourceSecurityDlpFilePatternsModel) getUpdateObjectSecurityDlpFilePatterns(ctx context.Context, state resourceSecurityDlpFilePatternsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Tag.IsNull() {
		result["tag"] = data.Tag.ValueString()
	}

	if data.Entries != nil {
		result["entries"] = data.expandSecurityDlpFilePatternsEntriesList(ctx, data.Entries, diags)
	}

	return &result
}

func (data *resourceSecurityDlpFilePatternsModel) getURLObjectSecurityDlpFilePatterns(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityDlpFilePatternsEntriesModel struct {
	Pattern    types.String `tfsdk:"pattern"`
	FilterType types.String `tfsdk:"filter_type"`
	FileType   types.String `tfsdk:"file_type"`
}

func (m *resourceSecurityDlpFilePatternsEntriesModel) flattenSecurityDlpFilePatternsEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpFilePatternsEntriesModel {
	if input == nil {
		return &resourceSecurityDlpFilePatternsEntriesModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpFilePatternsEntriesModel{}
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

func (s *resourceSecurityDlpFilePatternsModel) flattenSecurityDlpFilePatternsEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDlpFilePatternsEntriesModel {
	if o == nil {
		return []resourceSecurityDlpFilePatternsEntriesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument entries is not type of []interface{}.", "")
		return []resourceSecurityDlpFilePatternsEntriesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDlpFilePatternsEntriesModel{}
	}

	values := make([]resourceSecurityDlpFilePatternsEntriesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDlpFilePatternsEntriesModel
		values[i] = *m.flattenSecurityDlpFilePatternsEntries(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityDlpFilePatternsEntriesModel) expandSecurityDlpFilePatternsEntries(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Pattern.IsNull() {
		result["pattern"] = data.Pattern.ValueString()
	}

	if !data.FilterType.IsNull() {
		result["filterType"] = data.FilterType.ValueString()
	}

	if !data.FileType.IsNull() {
		result["fileType"] = data.FileType.ValueString()
	}

	return result
}

func (s *resourceSecurityDlpFilePatternsModel) expandSecurityDlpFilePatternsEntriesList(ctx context.Context, l []resourceSecurityDlpFilePatternsEntriesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDlpFilePatternsEntries(ctx, diags)
	}
	return result
}
