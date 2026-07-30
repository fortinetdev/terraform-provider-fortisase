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
var _ resource.Resource = &resourceSecurityCertLocalCaCert{}
var _ resource.ResourceWithMoveState = &resourceSecurityCertLocalCaCert{}

func newResourceSecurityCertLocalCaCert() resource.Resource {
	return &resourceSecurityCertLocalCaCert{}
}

type resourceSecurityCertLocalCaCert struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityCertLocalCaCertModel describes the resource data model.
type resourceSecurityCertLocalCaCertModel struct {
	ID             types.String                                 `tfsdk:"id"`
	Ftntid         types.Float64                                `tfsdk:"ftntid"`
	Name           types.String                                 `tfsdk:"name"`
	PrimaryKey     types.String                                 `tfsdk:"primary_key"`
	Type           types.String                                 `tfsdk:"type"`
	Source         types.String                                 `tfsdk:"source"`
	Issuer         *resourceSecurityCertLocalCaCertIssuerModel  `tfsdk:"issuer"`
	ValidFrom      types.String                                 `tfsdk:"valid_from"`
	ValidTo        types.String                                 `tfsdk:"valid_to"`
	SerialNumber   types.String                                 `tfsdk:"serial_number"`
	Usages         []resourceSecurityCertLocalCaCertUsagesModel `tfsdk:"usages"`
	Format         types.String                                 `tfsdk:"format"`
	CertName       types.String                                 `tfsdk:"cert_name"`
	Password       types.String                                 `tfsdk:"password"`
	FileContent    types.String                                 `tfsdk:"file_content"`
	KeyFileContent types.String                                 `tfsdk:"key_file_content"`
}

func (r *resourceSecurityCertLocalCaCert) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_local_ca_cert"
}

func (r *resourceSecurityCertLocalCaCert) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Certificate Resource API for FortiSASE",
		DeprecationMessage:  "Please use fortisase_system_certificate",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ftntid": schema.Float64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"primary_key": schema.StringAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("local-cer", "emote-ca"),
				},
				Computed: true,
			},
			"source": schema.StringAttribute{
				Computed: true,
			},
			"valid_from": schema.StringAttribute{
				Computed: true,
			},
			"valid_to": schema.StringAttribute{
				Computed: true,
			},
			"serial_number": schema.StringAttribute{
				Computed: true,
			},
			"format": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"cert_name": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"password": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"file_content": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"key_file_content": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"issuer": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"c": schema.StringAttribute{
						Computed: true,
					},
					"cn": schema.StringAttribute{
						Computed: true,
					},
					"l": schema.StringAttribute{
						Computed: true,
					},
					"o": schema.StringAttribute{
						Computed: true,
					},
					"ou": schema.StringAttribute{
						Computed: true,
					},
					"st": schema.StringAttribute{
						Computed: true,
					},
					"email_address": schema.StringAttribute{
						Computed: true,
					},
				},
				Computed: true,
			},
			"usages": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Computed: true,
						},
						"count": schema.Float64Attribute{
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *resourceSecurityCertLocalCaCert) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_cert_local_ca_cert"
}
func (r *resourceSecurityCertLocalCaCert) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_security_cert_local_ca_certs" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceSecurityCertLocalCaCertModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceSecurityCertLocalCaCert) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SystemCertificates")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityCertLocalCaCertModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityCertLocalCaCert(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityCertLocalCaCert(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateSecurityCertLocalCaCerts(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityCertLocalCaCert(ctx, "read", diags))

	read_output, err := c.ReadSecurityCertLocalCaCerts(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityCertLocalCaCert(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityCertLocalCaCert) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// No update operation for this resource
	resp.Diagnostics.AddError(
		"Update not supported",
		"This resource does not support update. You use terraform taint <resource_type>.<resource_name> to force a replacement.",
	)
}

func (r *resourceSecurityCertLocalCaCert) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SystemCertificates")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityCertLocalCaCertModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityCertLocalCaCert(ctx, "delete", diags))

	output, err := c.DeleteSecurityCertLocalCaCerts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityCertLocalCaCert) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityCertLocalCaCertModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityCertLocalCaCert(ctx, "read", diags))

	read_output, err := c.ReadSecurityCertLocalCaCerts(&input_model)
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

	diags.Append(data.refreshSecurityCertLocalCaCert(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityCertLocalCaCert) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceSecurityCertLocalCaCertModel) refreshSecurityCertLocalCaCert(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["id"]; ok {
		m.Ftntid = parseFloat64Value(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["source"]; ok {
		m.Source = parseStringValue(v)
	}

	if v, ok := o["issuer"]; ok {
		m.Issuer = m.Issuer.flattenSecurityCertLocalCaCertIssuer(ctx, v, &diags)
	}

	if v, ok := o["validFrom"]; ok {
		m.ValidFrom = parseStringValue(v)
	}

	if v, ok := o["validTo"]; ok {
		m.ValidTo = parseStringValue(v)
	}

	if v, ok := o["serialNumber"]; ok {
		m.SerialNumber = parseStringValue(v)
	}

	if v, ok := o["usages"]; ok {
		m.Usages = m.flattenSecurityCertLocalCaCertUsagesList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityCertLocalCaCertModel) getCreateObjectSecurityCertLocalCaCert(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})

	if !data.Format.IsNull() && !data.Format.IsUnknown() {
		result["format"] = data.Format.ValueString()
	}

	if !data.CertName.IsNull() && !data.CertName.IsUnknown() {
		result["certName"] = data.CertName.ValueString()
	}

	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		result["password"] = data.Password.ValueString()
	}

	if !data.FileContent.IsNull() && !data.FileContent.IsUnknown() {
		result["fileContent"] = data.FileContent.ValueString()
	}

	if !data.KeyFileContent.IsNull() && !data.KeyFileContent.IsUnknown() {
		result["keyFileContent"] = data.KeyFileContent.ValueString()
	}

	return &result
}

func (data *resourceSecurityCertLocalCaCertModel) getUpdateObjectSecurityCertLocalCaCert(ctx context.Context, state resourceSecurityCertLocalCaCertModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Ftntid.IsNull() && !data.Ftntid.IsUnknown() {
		result["id"] = data.Ftntid.ValueFloat64()
	}

	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		result["name"] = data.Name.ValueString()
	}

	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Source.IsNull() && !data.Source.IsUnknown() {
		result["source"] = data.Source.ValueString()
	}

	if !data.ValidFrom.IsNull() && !data.ValidFrom.IsUnknown() {
		result["validFrom"] = data.ValidFrom.ValueString()
	}

	if !data.ValidTo.IsNull() && !data.ValidTo.IsUnknown() {
		result["validTo"] = data.ValidTo.ValueString()
	}

	if !data.SerialNumber.IsNull() && !data.SerialNumber.IsUnknown() {
		result["serialNumber"] = data.SerialNumber.ValueString()
	}

	if !data.Format.IsNull() && !data.Format.IsUnknown() {
		result["format"] = data.Format.ValueString()
	}

	if !data.CertName.IsNull() && !data.CertName.IsUnknown() {
		result["certName"] = data.CertName.ValueString()
	}

	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		result["password"] = data.Password.ValueString()
	}

	if !data.FileContent.IsNull() && !data.FileContent.IsUnknown() {
		result["fileContent"] = data.FileContent.ValueString()
	}

	if !data.KeyFileContent.IsNull() && !data.KeyFileContent.IsUnknown() {
		result["keyFileContent"] = data.KeyFileContent.ValueString()
	}

	return &result
}

func (data *resourceSecurityCertLocalCaCertModel) getURLObjectSecurityCertLocalCaCert(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityCertLocalCaCertIssuerModel struct {
	C            types.String `tfsdk:"c"`
	Cn           types.String `tfsdk:"cn"`
	L            types.String `tfsdk:"l"`
	O            types.String `tfsdk:"o"`
	Ou           types.String `tfsdk:"ou"`
	St           types.String `tfsdk:"st"`
	EmailAddress types.String `tfsdk:"email_address"`
}

type resourceSecurityCertLocalCaCertUsagesModel struct {
	Type  types.String  `tfsdk:"type"`
	Count types.Float64 `tfsdk:"count"`
}

func (m *resourceSecurityCertLocalCaCertIssuerModel) flattenSecurityCertLocalCaCertIssuer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityCertLocalCaCertIssuerModel {
	if input == nil {
		return &resourceSecurityCertLocalCaCertIssuerModel{}
	}
	if m == nil {
		m = &resourceSecurityCertLocalCaCertIssuerModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["C"]; ok {
		m.C = parseStringValue(v)
	}

	if v, ok := o["CN"]; ok {
		m.Cn = parseStringValue(v)
	}

	if v, ok := o["L"]; ok {
		m.L = parseStringValue(v)
	}

	if v, ok := o["O"]; ok {
		m.O = parseStringValue(v)
	}

	if v, ok := o["OU"]; ok {
		m.Ou = parseStringValue(v)
	}

	if v, ok := o["ST"]; ok {
		m.St = parseStringValue(v)
	}

	if v, ok := o["emailAddress"]; ok {
		m.EmailAddress = parseStringValue(v)
	}

	return m
}

func (m *resourceSecurityCertLocalCaCertUsagesModel) flattenSecurityCertLocalCaCertUsages(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityCertLocalCaCertUsagesModel {
	if input == nil {
		return &resourceSecurityCertLocalCaCertUsagesModel{}
	}
	if m == nil {
		m = &resourceSecurityCertLocalCaCertUsagesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["count"]; ok {
		m.Count = parseFloat64Value(v)
	}

	return m
}

func (s *resourceSecurityCertLocalCaCertModel) flattenSecurityCertLocalCaCertUsagesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityCertLocalCaCertUsagesModel {
	if o == nil {
		return []resourceSecurityCertLocalCaCertUsagesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument usages is not type of []interface{}.", "")
		return []resourceSecurityCertLocalCaCertUsagesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityCertLocalCaCertUsagesModel{}
	}

	values := make([]resourceSecurityCertLocalCaCertUsagesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityCertLocalCaCertUsagesModel
		if i < len(s.Usages) {
			m = s.Usages[i]
		}
		values[i] = *m.flattenSecurityCertLocalCaCertUsages(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityCertLocalCaCertIssuerModel) expandSecurityCertLocalCaCertIssuer(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})

	return result
}

func (data *resourceSecurityCertLocalCaCertUsagesModel) expandSecurityCertLocalCaCertUsages(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})

	return result
}

func (s *resourceSecurityCertLocalCaCertModel) expandSecurityCertLocalCaCertUsagesList(ctx context.Context, l []resourceSecurityCertLocalCaCertUsagesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityCertLocalCaCertUsages(ctx, diags)
	}
	return result
}
