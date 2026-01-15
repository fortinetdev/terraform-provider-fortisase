// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
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
var _ resource.Resource = &resourceEndpointSandboxProfiles{}

func newResourceEndpointSandboxProfiles() resource.Resource {
	return &resourceEndpointSandboxProfiles{}
}

type resourceEndpointSandboxProfiles struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointSandboxProfilesModel describes the resource data model.
type resourceEndpointSandboxProfilesModel struct {
	ID                            types.String                                               `tfsdk:"id"`
	SandboxMode                   types.String                                               `tfsdk:"sandbox_mode"`
	NotificationType              types.Float64                                              `tfsdk:"notification_type"`
	TimeoutAwaitingSandboxResults types.Float64                                              `tfsdk:"timeout_awaiting_sandbox_results"`
	FileSubmissionOptions         *resourceEndpointSandboxProfilesFileSubmissionOptionsModel `tfsdk:"file_submission_options"`
	DetectionVerdictLevel         types.String                                               `tfsdk:"detection_verdict_level"`
	Exceptions                    *resourceEndpointSandboxProfilesExceptionsModel            `tfsdk:"exceptions"`
	RemediationActions            types.String                                               `tfsdk:"remediation_actions"`
	HostName                      types.String                                               `tfsdk:"host_name"`
	Authentication                types.Bool                                                 `tfsdk:"authentication"`
	Username                      types.String                                               `tfsdk:"username"`
	Password                      types.String                                               `tfsdk:"password"`
	PrimaryKey                    types.String                                               `tfsdk:"primary_key"`
}

func (r *resourceEndpointSandboxProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_sandbox_profiles"
}

func (r *resourceEndpointSandboxProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"sandbox_mode": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("Disabled", "FortiSASE", "StandaloneFortiSandbox"),
				},
				Computed: true,
				Optional: true,
			},
			"notification_type": schema.Float64Attribute{
				MarkdownDescription: "Integer representing how notifications should be handled on FortiSandbox file submission. 0 - display notification balloon when malware is detected in a submission. 1 - display a popup for all file submissions.",
				Computed:            true,
				Optional:            true,
			},
			"timeout_awaiting_sandbox_results": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(2147483647),
				},
				Computed: true,
				Optional: true,
			},
			"detection_verdict_level": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("Clean", "Malicious", "High", "Medium", "Low"),
				},
				Computed: true,
				Optional: true,
			},
			"remediation_actions": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("quarantine", "alert"),
				},
				Computed: true,
				Optional: true,
			},
			"host_name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
				Computed: true,
				Optional: true,
			},
			"authentication": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"username": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
				Computed: true,
				Optional: true,
			},
			"password": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"file_submission_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"all_email_downloads": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"all_files_mapped_network_drives": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"all_files_removable_media": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"all_web_downloads": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"exceptions": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"exclude_files_from_trusted_sources": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"files": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"folders": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceEndpointSandboxProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_sandbox_profiles"
}

func (r *resourceEndpointSandboxProfiles) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointSandboxProfiles")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointSandboxProfilesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointSandboxProfiles(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointSandboxProfiles(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateEndpointSandboxProfiles(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectEndpointSandboxProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointSandboxProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSandboxProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSandboxProfiles) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointSandboxProfiles")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointSandboxProfilesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointSandboxProfilesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointSandboxProfiles(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointSandboxProfiles(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointSandboxProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointSandboxProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointSandboxProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSandboxProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSandboxProfiles) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointSandboxProfiles) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointSandboxProfilesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointSandboxProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointSandboxProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSandboxProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSandboxProfiles) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointSandboxProfilesModel) refreshEndpointSandboxProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["sandboxMode"]; ok {
		m.SandboxMode = parseStringValue(v)
	}

	if v, ok := o["notificationType"]; ok {
		m.NotificationType = parseFloat64Value(v)
	}

	if v, ok := o["timeoutAwaitingSandboxResults"]; ok {
		m.TimeoutAwaitingSandboxResults = parseFloat64Value(v)
	}

	if v, ok := o["fileSubmissionOptions"]; ok {
		m.FileSubmissionOptions = m.FileSubmissionOptions.flattenEndpointSandboxProfilesFileSubmissionOptions(ctx, v, &diags)
	}

	if v, ok := o["detectionVerdictLevel"]; ok {
		m.DetectionVerdictLevel = parseStringValue(v)
	}

	if v, ok := o["exceptions"]; ok {
		m.Exceptions = m.Exceptions.flattenEndpointSandboxProfilesExceptions(ctx, v, &diags)
	}

	if v, ok := o["remediationActions"]; ok {
		m.RemediationActions = parseStringValue(v)
	}

	if v, ok := o["hostName"]; ok {
		m.HostName = parseStringValue(v)
	}

	if v, ok := o["authentication"]; ok {
		m.Authentication = parseBoolValue(v)
	}

	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	if v, ok := o["password"]; ok {
		m.Password = parseStringValue(v)
	}

	return diags
}

func (data *resourceEndpointSandboxProfilesModel) getCreateObjectEndpointSandboxProfiles(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SandboxMode.IsNull() {
		result["sandboxMode"] = data.SandboxMode.ValueString()
	}

	if !data.NotificationType.IsNull() {
		result["notificationType"] = data.NotificationType.ValueFloat64()
	}

	if !data.TimeoutAwaitingSandboxResults.IsNull() {
		result["timeoutAwaitingSandboxResults"] = data.TimeoutAwaitingSandboxResults.ValueFloat64()
	}

	if data.FileSubmissionOptions != nil && !isZeroStruct(*data.FileSubmissionOptions) {
		result["fileSubmissionOptions"] = data.FileSubmissionOptions.expandEndpointSandboxProfilesFileSubmissionOptions(ctx, diags)
	}

	if !data.DetectionVerdictLevel.IsNull() {
		result["detectionVerdictLevel"] = data.DetectionVerdictLevel.ValueString()
	}

	if data.Exceptions != nil && !isZeroStruct(*data.Exceptions) {
		result["exceptions"] = data.Exceptions.expandEndpointSandboxProfilesExceptions(ctx, diags)
	}

	if !data.RemediationActions.IsNull() {
		result["remediationActions"] = data.RemediationActions.ValueString()
	}

	if !data.HostName.IsNull() {
		result["hostName"] = data.HostName.ValueString()
	}

	if !data.Authentication.IsNull() {
		result["authentication"] = data.Authentication.ValueBool()
	}

	if !data.Username.IsNull() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	return &result
}

func (data *resourceEndpointSandboxProfilesModel) getUpdateObjectEndpointSandboxProfiles(ctx context.Context, state resourceEndpointSandboxProfilesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SandboxMode.IsNull() {
		result["sandboxMode"] = data.SandboxMode.ValueString()
	}

	if !data.NotificationType.IsNull() {
		result["notificationType"] = data.NotificationType.ValueFloat64()
	}

	if !data.TimeoutAwaitingSandboxResults.IsNull() {
		result["timeoutAwaitingSandboxResults"] = data.TimeoutAwaitingSandboxResults.ValueFloat64()
	}

	if data.FileSubmissionOptions != nil {
		result["fileSubmissionOptions"] = data.FileSubmissionOptions.expandEndpointSandboxProfilesFileSubmissionOptions(ctx, diags)
	}

	if !data.DetectionVerdictLevel.IsNull() {
		result["detectionVerdictLevel"] = data.DetectionVerdictLevel.ValueString()
	}

	if data.Exceptions != nil {
		result["exceptions"] = data.Exceptions.expandEndpointSandboxProfilesExceptions(ctx, diags)
	}

	if !data.RemediationActions.IsNull() {
		result["remediationActions"] = data.RemediationActions.ValueString()
	}

	if !data.HostName.IsNull() {
		result["hostName"] = data.HostName.ValueString()
	}

	if !data.Authentication.IsNull() {
		result["authentication"] = data.Authentication.ValueBool()
	}

	if !data.Username.IsNull() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	return &result
}

func (data *resourceEndpointSandboxProfilesModel) getURLObjectEndpointSandboxProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointSandboxProfilesFileSubmissionOptionsModel struct {
	AllEmailDownloads           types.String `tfsdk:"all_email_downloads"`
	AllFilesMappedNetworkDrives types.String `tfsdk:"all_files_mapped_network_drives"`
	AllFilesRemovableMedia      types.String `tfsdk:"all_files_removable_media"`
	AllWebDownloads             types.String `tfsdk:"all_web_downloads"`
}

type resourceEndpointSandboxProfilesExceptionsModel struct {
	ExcludeFilesFromTrustedSources types.String `tfsdk:"exclude_files_from_trusted_sources"`
	Files                          types.Set    `tfsdk:"files"`
	Folders                        types.Set    `tfsdk:"folders"`
}

func (m *resourceEndpointSandboxProfilesFileSubmissionOptionsModel) flattenEndpointSandboxProfilesFileSubmissionOptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointSandboxProfilesFileSubmissionOptionsModel {
	if input == nil {
		return &resourceEndpointSandboxProfilesFileSubmissionOptionsModel{}
	}
	if m == nil {
		m = &resourceEndpointSandboxProfilesFileSubmissionOptionsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["allEmailDownloads"]; ok {
		m.AllEmailDownloads = parseStringValue(v)
	}

	if v, ok := o["allFilesMappedNetworkDrives"]; ok {
		m.AllFilesMappedNetworkDrives = parseStringValue(v)
	}

	if v, ok := o["allFilesRemovableMedia"]; ok {
		m.AllFilesRemovableMedia = parseStringValue(v)
	}

	if v, ok := o["allWebDownloads"]; ok {
		m.AllWebDownloads = parseStringValue(v)
	}

	return m
}

func (m *resourceEndpointSandboxProfilesExceptionsModel) flattenEndpointSandboxProfilesExceptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointSandboxProfilesExceptionsModel {
	if input == nil {
		return &resourceEndpointSandboxProfilesExceptionsModel{}
	}
	if m == nil {
		m = &resourceEndpointSandboxProfilesExceptionsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["excludeFilesFromTrustedSources"]; ok {
		m.ExcludeFilesFromTrustedSources = parseStringValue(v)
	}

	if v, ok := o["files"]; ok {
		m.Files = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["folders"]; ok {
		m.Folders = parseSetValue(ctx, v, types.StringType)
	}

	return m
}

func (data *resourceEndpointSandboxProfilesFileSubmissionOptionsModel) expandEndpointSandboxProfilesFileSubmissionOptions(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AllEmailDownloads.IsNull() {
		result["allEmailDownloads"] = data.AllEmailDownloads.ValueString()
	}

	if !data.AllFilesMappedNetworkDrives.IsNull() {
		result["allFilesMappedNetworkDrives"] = data.AllFilesMappedNetworkDrives.ValueString()
	}

	if !data.AllFilesRemovableMedia.IsNull() {
		result["allFilesRemovableMedia"] = data.AllFilesRemovableMedia.ValueString()
	}

	if !data.AllWebDownloads.IsNull() {
		result["allWebDownloads"] = data.AllWebDownloads.ValueString()
	}

	return result
}

func (data *resourceEndpointSandboxProfilesExceptionsModel) expandEndpointSandboxProfilesExceptions(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ExcludeFilesFromTrustedSources.IsNull() {
		result["excludeFilesFromTrustedSources"] = data.ExcludeFilesFromTrustedSources.ValueString()
	}

	if !data.Files.IsNull() {
		result["files"] = expandSetToStringList(data.Files)
	}

	if !data.Folders.IsNull() {
		result["folders"] = expandSetToStringList(data.Folders)
	}

	return result
}
