// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	forticlient "github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const systemCertificateAPIBasePath = "/resource-api/v2/system/certificate"

var systemCertificateTypes = map[string]struct{}{
	"ca-certificate":        {},
	"local-certificate":     {},
	"hsm-local-certificate": {},
	"remote-certificate":    {},
	"remote-ca-certificate": {},
}

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSystemCertificate{}
var _ resource.ResourceWithValidateConfig = &resourceSystemCertificate{}
var _ resource.ResourceWithModifyPlan = &resourceSystemCertificate{}
var _ resource.ResourceWithImportState = &resourceSystemCertificate{}

func newResourceSystemCertificate() resource.Resource {
	return &resourceSystemCertificate{}
}

type resourceSystemCertificate struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSystemCertificateModel describes the resource data model.
type resourceSystemCertificateModel struct {
	ID              types.String `tfsdk:"id"`
	CertificateType types.String `tfsdk:"certificate_type"`
	PrimaryKey      types.String `tfsdk:"primary_key"`
	FileContent     types.String `tfsdk:"file_content"`
	KeyFileContent  types.String `tfsdk:"key_file_content"`
	Format          types.String `tfsdk:"format"`
	Password        types.String `tfsdk:"password"`
}

func (r *resourceSystemCertificate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_certificate"
}

func (r *resourceSystemCertificate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	requiresReplace := []planmodifier.String{
		stringplanmodifier.RequiresReplace(),
	}

	resp.Schema = schema.Schema{
		MarkdownDescription: "Imports and manages a FortiSASE system certificate. The resource supports CA, local, HSM local, remote, and remote CA certificates.\nCertificates can be found in the GUI: System -> Certificates",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Certificate type. Valid values are `ca-certificate`, `local-certificate`, `hsm-local-certificate`, `remote-certificate`, and `remote-ca-certificate`.",
				PlanModifiers:       requiresReplace,
			},
			"primary_key": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Primary key of the certificate.",
				PlanModifiers:       requiresReplace,
			},
			"file_content": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: "Base64-encoded certificate file content. Required when creating a certificate and omitted when managing an imported certificate without its original file content.",
				PlanModifiers:       requiresReplace,
			},
			"key_file_content": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: "Base64-encoded private key file content. Required when creating a `ca-certificate` or `local-certificate` with `format` set to `regular`; it may be omitted for an imported certificate.",
				PlanModifiers:       requiresReplace,
			},
			"format": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Certificate format. Required when creating a `ca-certificate` or `local-certificate` and optional for an imported certificate; valid values are `pkcs12` and `regular`.",
				PlanModifiers:       requiresReplace,
			},
			"password": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: "Certificate password. Supported for `ca-certificate` and `local-certificate`.",
				PlanModifiers:       requiresReplace,
			},
		},
	}
}

func (r *resourceSystemCertificate) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data resourceSystemCertificateModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.CertificateType.IsNull() && !data.CertificateType.IsUnknown() {
		certificateType := data.CertificateType.ValueString()
		if _, ok := systemCertificateTypes[certificateType]; !ok {
			resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
				path.Root("certificate_type"),
				"Invalid Certificate Type",
				fmt.Sprintf("Certificate type %q is not supported. Valid values are ca-certificate, local-certificate, hsm-local-certificate, remote-certificate, and remote-ca-certificate.", certificateType),
			))
			return
		}

		requiresFormat := certificateType == "ca-certificate" || certificateType == "local-certificate"
		if requiresFormat {
			if !data.Format.IsNull() && !data.Format.IsUnknown() && data.Format.ValueString() != "regular" && !data.KeyFileContent.IsNull() && !data.KeyFileContent.IsUnknown() {
				resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
					path.Root("key_file_content"),
					"Unsupported Certificate Key File Content",
					"The key_file_content attribute can only be configured when format is regular.",
				))
			}
		} else {
			if !data.Format.IsNull() && !data.Format.IsUnknown() {
				resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
					path.Root("format"),
					"Unsupported Certificate Format",
					fmt.Sprintf("The format attribute cannot be configured when certificate_type is %q.", certificateType),
				))
			}
			if !data.Password.IsNull() && !data.Password.IsUnknown() {
				resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
					path.Root("password"),
					"Unsupported Certificate Password",
					fmt.Sprintf("The password attribute cannot be configured when certificate_type is %q.", certificateType),
				))
			}
			if !data.KeyFileContent.IsNull() && !data.KeyFileContent.IsUnknown() {
				resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
					path.Root("key_file_content"),
					"Unsupported Certificate Key File Content",
					fmt.Sprintf("The key_file_content attribute cannot be configured when certificate_type is %q.", certificateType),
				))
			}
		}
	}

	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() && data.PrimaryKey.ValueString() == "" {
		resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
			path.Root("primary_key"),
			"Invalid Primary Key",
			"The primary_key attribute must not be empty.",
		))
	}

	if !data.FileContent.IsNull() && !data.FileContent.IsUnknown() {
		fileContent := data.FileContent.ValueString()
		if fileContent == "" {
			resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
				path.Root("file_content"),
				"Invalid Certificate File Content",
				"The file_content attribute must not be empty.",
			))
		} else if _, err := base64.StdEncoding.DecodeString(fileContent); err != nil {
			resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
				path.Root("file_content"),
				"Invalid Certificate File Content",
				"The file_content attribute must be a valid base64-encoded string.",
			))
		}
	}

	if !data.KeyFileContent.IsNull() && !data.KeyFileContent.IsUnknown() {
		keyFileContent := data.KeyFileContent.ValueString()
		if keyFileContent == "" {
			resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
				path.Root("key_file_content"),
				"Invalid Certificate Key File Content",
				"The key_file_content attribute must not be empty.",
			))
		} else if _, err := base64.StdEncoding.DecodeString(keyFileContent); err != nil {
			resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
				path.Root("key_file_content"),
				"Invalid Certificate Key File Content",
				"The key_file_content attribute must be a valid base64-encoded string.",
			))
		}
	}

	if !data.Format.IsNull() && !data.Format.IsUnknown() {
		format := data.Format.ValueString()
		if format != "pkcs12" && format != "regular" {
			resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
				path.Root("format"),
				"Invalid Certificate Format",
				fmt.Sprintf("Certificate format %q is not supported. Valid values are pkcs12 and regular.", format),
			))
		}
	}
}

func (r *resourceSystemCertificate) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || !req.State.Raw.IsNull() {
		return
	}

	var data resourceSystemCertificateModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	validateSystemCertificateCreateInputs(data, &resp.Diagnostics, false)
}

func (r *resourceSystemCertificate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_system_certificate"
}

func (r *resourceSystemCertificate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SystemCertificates")
	lock.Lock()
	defer lock.Unlock()

	var data resourceSystemCertificateModel
	diags := &resp.Diagnostics
	diags.Append(req.Plan.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	validateSystemCertificateCreateInputs(data, diags, true)
	if diags.HasError() {
		return
	}

	body := map[string]interface{}{
		"primaryKey":  data.PrimaryKey.ValueString(),
		"fileContent": data.FileContent.ValueString(),
	}
	if !data.Format.IsNull() && !data.Format.IsUnknown() {
		body["format"] = data.Format.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		body["password"] = data.Password.ValueString()
	}
	if !data.KeyFileContent.IsNull() && !data.KeyFileContent.IsUnknown() {
		body["keyFileContent"] = data.KeyFileContent.ValueString()
	}

	certificateType := data.CertificateType.ValueString()
	endpoint := fmt.Sprintf("%s/%ss/import", systemCertificateAPIBasePath, certificateType)
	inputModel := forticlient.InputModel{
		HTTPMethod: http.MethodPost,
		URL:        endpoint,
		BodyParams: body,
	}
	output, err := r.fortiClient.Client.SendRequest(&inputModel)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error creating resource %s: %v", r.resourceName, err),
			fmt.Sprintf("API response: %v", output),
		)
		return
	}

	primaryKey := data.PrimaryKey.ValueString()
	if responsePrimaryKey, ok := getCreateResponseMkey(output, "primaryKey"); ok {
		primaryKey = responsePrimaryKey
		data.PrimaryKey = types.StringValue(primaryKey)
	}
	data.ID = types.StringValue(primaryKey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSystemCertificate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resourceSystemCertificateModel
	diags := &resp.Diagnostics
	diags.Append(req.State.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	primaryKey := data.ID.ValueString()
	if primaryKey == "" {
		primaryKey = data.PrimaryKey.ValueString()
	}
	endpoint := fmt.Sprintf(
		"%s/%ss/%s",
		systemCertificateAPIBasePath,
		data.CertificateType.ValueString(),
		url.PathEscape(primaryKey),
	)

	inputModel := forticlient.InputModel{
		Mkey:       primaryKey,
		MkeyName:   "primaryKey",
		HTTPMethod: http.MethodGet,
		URL:        endpoint,
	}
	output, err := r.fortiClient.Client.ReadRequest(&inputModel)
	if err != nil {
		if isNotFoundResponse(output) {
			resp.State.RemoveResource(ctx)
			return
		}
		diags.AddError(
			fmt.Sprintf("Error reading resource %s: %v", r.resourceName, err),
			fmt.Sprintf("API response: %v", output),
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSystemCertificate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Update Not Supported",
		"This resource does not support update. Changing any configurable attribute requires replacement.",
	)
}

func (r *resourceSystemCertificate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	certificateType, primaryKey, ok := strings.Cut(req.ID, "/")
	certificateType = strings.TrimSpace(certificateType)
	primaryKey = strings.TrimSpace(primaryKey)
	if !ok || certificateType == "" || primaryKey == "" {
		resp.Diagnostics.AddError(
			"Invalid System Certificate Import ID",
			"Use <certificate_type>/<primary_key> as the import ID, for example local-certificate/my-certificate.",
		)
		return
	}
	if _, ok := systemCertificateTypes[certificateType]; !ok {
		resp.Diagnostics.AddError(
			"Invalid System Certificate Import Type",
			fmt.Sprintf("Certificate type %q is not supported. Valid values are ca-certificate, local-certificate, hsm-local-certificate, remote-certificate, and remote-ca-certificate.", certificateType),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), primaryKey)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("certificate_type"), certificateType)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), primaryKey)...)
}

func (r *resourceSystemCertificate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SystemCertificates")
	lock.Lock()
	defer lock.Unlock()

	var data resourceSystemCertificateModel
	diags := &resp.Diagnostics
	diags.Append(req.State.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	primaryKey := data.ID.ValueString()
	if primaryKey == "" {
		primaryKey = data.PrimaryKey.ValueString()
	}
	endpoint := fmt.Sprintf(
		"%s/%ss/%s",
		systemCertificateAPIBasePath,
		data.CertificateType.ValueString(),
		url.PathEscape(primaryKey),
	)

	inputModel := forticlient.InputModel{
		Mkey:       primaryKey,
		MkeyName:   "primaryKey",
		HTTPMethod: http.MethodDelete,
		URL:        endpoint,
	}
	output, err := r.fortiClient.Client.SendRequest(&inputModel)
	if err != nil && !isNotFoundResponse(output) {
		diags.AddError(
			fmt.Sprintf("Error deleting resource %s: %v", r.resourceName, err),
			fmt.Sprintf("API response: %v", output),
		)
	}
}

func validateSystemCertificateCreateInputs(data resourceSystemCertificateModel, diags *diag.Diagnostics, unknownIsError bool) {
	if data.FileContent.IsNull() || (unknownIsError && data.FileContent.IsUnknown()) {
		diags.Append(diag.NewAttributeErrorDiagnostic(
			path.Root("file_content"),
			"Missing Certificate File Content",
			"The file_content attribute is required when creating a system certificate. It may only be omitted for an imported certificate.",
		))
	}

	if data.CertificateType.IsNull() || data.CertificateType.IsUnknown() {
		return
	}
	certificateType := data.CertificateType.ValueString()
	if certificateType != "ca-certificate" && certificateType != "local-certificate" {
		return
	}

	if data.Format.IsNull() || (unknownIsError && data.Format.IsUnknown()) {
		diags.Append(diag.NewAttributeErrorDiagnostic(
			path.Root("format"),
			"Missing Certificate Format",
			fmt.Sprintf("The format attribute is required when creating a certificate with certificate_type %q. It may be omitted for an imported certificate.", certificateType),
		))
		return
	}
	if data.Format.IsUnknown() || data.Format.ValueString() != "regular" {
		return
	}
	if data.KeyFileContent.IsNull() || (unknownIsError && data.KeyFileContent.IsUnknown()) {
		diags.Append(diag.NewAttributeErrorDiagnostic(
			path.Root("key_file_content"),
			"Missing Certificate Key File Content",
			fmt.Sprintf("The key_file_content attribute is required when creating a certificate with certificate_type %q and format regular.", certificateType),
		))
	}
}
