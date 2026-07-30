// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
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
var _ resource.Resource = &resourceSecurityDlpFilePattern{}
var _ resource.ResourceWithMoveState = &resourceSecurityDlpFilePattern{}

func newResourceSecurityDlpFilePattern() resource.Resource {
	return &resourceSecurityDlpFilePattern{}
}

type resourceSecurityDlpFilePattern struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityDlpFilePatternModel describes the resource data model.
type resourceSecurityDlpFilePatternModel struct {
	ID         types.String                                 `tfsdk:"id"`
	PrimaryKey types.String                                 `tfsdk:"primary_key"`
	Tag        types.String                                 `tfsdk:"tag"`
	Entries    []resourceSecurityDlpFilePatternEntriesModel `tfsdk:"entries"`
}

func (r *resourceSecurityDlpFilePattern) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_file_pattern"
}

func (r *resourceSecurityDlpFilePattern) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DLP File Pattern Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthBetween(1, 63),
				},
				Computed: true,
				Optional: true,
			},
			"entries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pattern": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"filter_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("type", "pattern"),
							},
							Computed: true,
							Optional: true,
						},
						"file_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("7z", "arj", "cab", "lzh", "rar", "tar", "zip", "bzip", "gzip", "bzip2", "xz", "bat", "uue", "mime", "base64", "binhex", "elf", "exe", "hta", "html", "jad", "class", "cod", "javascript", "msoffice", "msofficex", "fsg", "upx", "petite", "aspack", "sis", "hlp", "activemime", "jpeg", "gif", "tiff", "png", "bmp", "unknown", "mpeg", "mov", "mp3", "wma", "wav", "pdf", "avi", "rm", "torrent", "hibun", "msi", "mach-o", "dmg", ".net", "xar", "chm", "iso", "crx", "flac"),
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

func (r *resourceSecurityDlpFilePattern) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *resourceSecurityDlpFilePattern) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_dlp_file_patterns" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityDlpFilePatternModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityDlpFilePattern) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpFilePatterns")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityDlpFilePatternModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityDlpFilePattern(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpFilePattern(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityDlpFilePatterns(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	if responseMkey, ok := getCreateResponseMkey(output, "primaryKey"); ok {
		mkey = responseMkey
	}
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpFilePattern(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFilePatterns(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFilePattern(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpFilePattern) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpFilePatterns")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDlpFilePatternModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDlpFilePatternModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDlpFilePattern(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpFilePattern(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpFilePattern(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFilePatterns(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFilePattern(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpFilePattern) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDlpFilePatterns")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityDlpFilePatternModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpFilePattern(ctx, "delete", diags))

	output, err := c.DeleteSecurityDlpFilePatterns(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityDlpFilePattern) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpFilePatternModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpFilePattern(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFilePatterns(&input_model)
	if err != nil {
		if isNotFoundResponse(read_output) {
			resp.State.RemoveResource(ctx)
			return
		}
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
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

func (r *resourceSecurityDlpFilePattern) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityDlpFilePatternModel) refreshSecurityDlpFilePattern(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.Entries = m.flattenSecurityDlpFilePatternEntriesList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityDlpFilePatternModel) getCreateObjectSecurityDlpFilePattern(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})

	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		result["tag"] = data.Tag.ValueString()
	}

	result["entries"] = data.expandSecurityDlpFilePatternEntriesList(ctx, data.Entries, diags)

	return &result
}

func (data *resourceSecurityDlpFilePatternModel) getUpdateObjectSecurityDlpFilePattern(ctx context.Context, state resourceSecurityDlpFilePatternModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		result["tag"] = data.Tag.ValueString()
	}

	if data.Entries != nil {
		result["entries"] = data.expandSecurityDlpFilePatternEntriesList(ctx, data.Entries, diags)
	}

	return &result
}

func (data *resourceSecurityDlpFilePatternModel) getURLObjectSecurityDlpFilePattern(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityDlpFilePatternEntriesModel struct {
	Pattern    types.String `tfsdk:"pattern"`
	FilterType types.String `tfsdk:"filter_type"`
	FileType   types.String `tfsdk:"file_type"`
}

func (m *resourceSecurityDlpFilePatternEntriesModel) flattenSecurityDlpFilePatternEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpFilePatternEntriesModel {
	if input == nil {
		return &resourceSecurityDlpFilePatternEntriesModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpFilePatternEntriesModel{}
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

func (s *resourceSecurityDlpFilePatternModel) flattenSecurityDlpFilePatternEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDlpFilePatternEntriesModel {
	if o == nil {
		return []resourceSecurityDlpFilePatternEntriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument entries is not type of []interface{}.", "")
		return []resourceSecurityDlpFilePatternEntriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDlpFilePatternEntriesModel{}
	}

	values := make([]resourceSecurityDlpFilePatternEntriesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDlpFilePatternEntriesModel
		if i < len(s.Entries) {
			m = s.Entries[i]
		}
		values[i] = *m.flattenSecurityDlpFilePatternEntries(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityDlpFilePatternEntriesModel) expandSecurityDlpFilePatternEntries(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Pattern.IsNull() && !data.Pattern.IsUnknown() {
		result["pattern"] = data.Pattern.ValueString()
	}

	if !data.FilterType.IsNull() && !data.FilterType.IsUnknown() {
		result["filterType"] = data.FilterType.ValueString()
	}

	if !data.FileType.IsNull() && !data.FileType.IsUnknown() {
		result["fileType"] = data.FileType.ValueString()
	}

	return result
}

func (s *resourceSecurityDlpFilePatternModel) expandSecurityDlpFilePatternEntriesList(ctx context.Context, l []resourceSecurityDlpFilePatternEntriesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDlpFilePatternEntries(ctx, diags)
	}
	return result
}
