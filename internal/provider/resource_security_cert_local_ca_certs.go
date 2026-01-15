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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSecurityCertLocalCaCerts{}

func newResourceSecurityCertLocalCaCerts() resource.Resource {
	return &resourceSecurityCertLocalCaCerts{}
}

type resourceSecurityCertLocalCaCerts struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityCertLocalCaCertsModel describes the resource data model.
type resourceSecurityCertLocalCaCertsModel struct {
	ID             types.String                                  `tfsdk:"id"`
	Ftntid         types.Float64                                 `tfsdk:"ftntid"`
	Name           types.String                                  `tfsdk:"name"`
	PrimaryKey     types.String                                  `tfsdk:"primary_key"`
	Type           types.String                                  `tfsdk:"type"`
	Source         types.String                                  `tfsdk:"source"`
	Issuer         *resourceSecurityCertLocalCaCertsIssuerModel  `tfsdk:"issuer"`
	ValidFrom      types.String                                  `tfsdk:"valid_from"`
	ValidTo        types.String                                  `tfsdk:"valid_to"`
	SerialNumber   types.String                                  `tfsdk:"serial_number"`
	Usages         []resourceSecurityCertLocalCaCertsUsagesModel `tfsdk:"usages"`
	Format         types.String                                  `tfsdk:"format"`
	CertName       types.String                                  `tfsdk:"cert_name"`
	Password       types.String                                  `tfsdk:"password"`
	FileContent    types.String                                  `tfsdk:"file_content"`
	KeyFileContent types.String                                  `tfsdk:"key_file_content"`
}

func (r *resourceSecurityCertLocalCaCerts) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_local_ca_certs"
}

func (r *resourceSecurityCertLocalCaCerts) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
					stringvalidator.OneOf("local-cer", "emote-ca"),
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
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"cn": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"l": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"o": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"ou": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"st": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"email_address": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
				Computed: true,
			},
			"usages": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"count": schema.Float64Attribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Float64{
								float64planmodifier.RequiresReplace(),
							},
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *resourceSecurityCertLocalCaCerts) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_cert_local_ca_certs"
}

func (r *resourceSecurityCertLocalCaCerts) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityCertLocalCaCertsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityCertLocalCaCerts(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityCertLocalCaCerts(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityCertLocalCaCerts(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityCertLocalCaCerts(ctx, "read", diags))

	read_output, err := c.ReadSecurityCertLocalCaCerts(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityCertLocalCaCerts(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityCertLocalCaCerts) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// No update operation for this resource
	resp.Diagnostics.AddError(
		"Update not supported",
		"This resource does not support update. You use terraform taint <resource_type>.<resource_name> to force a replacement.",
	)
}

func (r *resourceSecurityCertLocalCaCerts) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityCertLocalCaCertsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityCertLocalCaCerts(ctx, "delete", diags))

	output, err := c.DeleteSecurityCertLocalCaCerts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityCertLocalCaCerts) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityCertLocalCaCertsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityCertLocalCaCerts(ctx, "read", diags))

	read_output, err := c.ReadSecurityCertLocalCaCerts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityCertLocalCaCerts(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityCertLocalCaCerts) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityCertLocalCaCertsModel) refreshSecurityCertLocalCaCerts(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.Issuer = m.Issuer.flattenSecurityCertLocalCaCertsIssuer(ctx, v, &diags)
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
		m.Usages = m.flattenSecurityCertLocalCaCertsUsagesList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityCertLocalCaCertsModel) getCreateObjectSecurityCertLocalCaCerts(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})

	if !data.Format.IsNull() {
		result["format"] = data.Format.ValueString()
	}

	if !data.CertName.IsNull() {
		result["certName"] = data.CertName.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	if !data.FileContent.IsNull() {
		result["fileContent"] = data.FileContent.ValueString()
	}

	if !data.KeyFileContent.IsNull() {
		result["keyFileContent"] = data.KeyFileContent.ValueString()
	}

	return &result
}

func (data *resourceSecurityCertLocalCaCertsModel) getUpdateObjectSecurityCertLocalCaCerts(ctx context.Context, state resourceSecurityCertLocalCaCertsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})

	return &result
}

func (data *resourceSecurityCertLocalCaCertsModel) getURLObjectSecurityCertLocalCaCerts(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityCertLocalCaCertsIssuerModel struct {
	C            types.String `tfsdk:"c"`
	Cn           types.String `tfsdk:"cn"`
	L            types.String `tfsdk:"l"`
	O            types.String `tfsdk:"o"`
	Ou           types.String `tfsdk:"ou"`
	St           types.String `tfsdk:"st"`
	EmailAddress types.String `tfsdk:"email_address"`
}

type resourceSecurityCertLocalCaCertsUsagesModel struct {
	Type  types.String  `tfsdk:"type"`
	Count types.Float64 `tfsdk:"count"`
}

func (m *resourceSecurityCertLocalCaCertsIssuerModel) flattenSecurityCertLocalCaCertsIssuer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityCertLocalCaCertsIssuerModel {
	if input == nil {
		return &resourceSecurityCertLocalCaCertsIssuerModel{}
	}
	if m == nil {
		m = &resourceSecurityCertLocalCaCertsIssuerModel{}
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

func (m *resourceSecurityCertLocalCaCertsUsagesModel) flattenSecurityCertLocalCaCertsUsages(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityCertLocalCaCertsUsagesModel {
	if input == nil {
		return &resourceSecurityCertLocalCaCertsUsagesModel{}
	}
	if m == nil {
		m = &resourceSecurityCertLocalCaCertsUsagesModel{}
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

func (s *resourceSecurityCertLocalCaCertsModel) flattenSecurityCertLocalCaCertsUsagesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityCertLocalCaCertsUsagesModel {
	if o == nil {
		return []resourceSecurityCertLocalCaCertsUsagesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument usages is not type of []interface{}.", "")
		return []resourceSecurityCertLocalCaCertsUsagesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityCertLocalCaCertsUsagesModel{}
	}

	values := make([]resourceSecurityCertLocalCaCertsUsagesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityCertLocalCaCertsUsagesModel
		values[i] = *m.flattenSecurityCertLocalCaCertsUsages(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityCertLocalCaCertsIssuerModel) expandSecurityCertLocalCaCertsIssuer(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.C.IsNull() {
		result["C"] = data.C.ValueString()
	}

	if !data.Cn.IsNull() {
		result["CN"] = data.Cn.ValueString()
	}

	if !data.L.IsNull() {
		result["L"] = data.L.ValueString()
	}

	if !data.O.IsNull() {
		result["O"] = data.O.ValueString()
	}

	if !data.Ou.IsNull() {
		result["OU"] = data.Ou.ValueString()
	}

	if !data.St.IsNull() {
		result["ST"] = data.St.ValueString()
	}

	if !data.EmailAddress.IsNull() {
		result["emailAddress"] = data.EmailAddress.ValueString()
	}

	return result
}

func (data *resourceSecurityCertLocalCaCertsUsagesModel) expandSecurityCertLocalCaCertsUsages(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Type.IsNull() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Count.IsNull() {
		result["count"] = data.Count.ValueFloat64()
	}

	return result
}

func (s *resourceSecurityCertLocalCaCertsModel) expandSecurityCertLocalCaCertsUsagesList(ctx context.Context, l []resourceSecurityCertLocalCaCertsUsagesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityCertLocalCaCertsUsages(ctx, diags)
	}
	return result
}
