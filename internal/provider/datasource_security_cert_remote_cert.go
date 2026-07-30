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
var _ datasource.DataSource = &datasourceSecurityCertRemoteCert{}

func newDatasourceSecurityCertRemoteCert() datasource.DataSource {
	return &datasourceSecurityCertRemoteCert{}
}

type datasourceSecurityCertRemoteCert struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityCertRemoteCertModel describes the datasource data model.
type datasourceSecurityCertRemoteCertModel struct {
	Ftntid       types.Float64                                 `tfsdk:"ftntid"`
	Name         types.String                                  `tfsdk:"name"`
	PrimaryKey   types.String                                  `tfsdk:"primary_key"`
	Type         types.String                                  `tfsdk:"type"`
	Source       types.String                                  `tfsdk:"source"`
	Issuer       *datasourceSecurityCertRemoteCertIssuerModel  `tfsdk:"issuer"`
	ValidFrom    types.String                                  `tfsdk:"valid_from"`
	ValidTo      types.String                                  `tfsdk:"valid_to"`
	SerialNumber types.String                                  `tfsdk:"serial_number"`
	Usages       []datasourceSecurityCertRemoteCertUsagesModel `tfsdk:"usages"`
}

func (r *datasourceSecurityCertRemoteCert) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_remote_cert"
}

func (r *datasourceSecurityCertRemoteCert) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Certificate Resource API for FortiSASE",
		DeprecationMessage:  "Please use fortisase_system_certificate",
		Attributes: map[string]schema.Attribute{
			"ftntid": schema.Float64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"primary_key": schema.StringAttribute{
				Required: true,
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

func (r *datasourceSecurityCertRemoteCert) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_cert_remote_cert"
}

func (r *datasourceSecurityCertRemoteCert) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityCertRemoteCertModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityCertRemoteCert(ctx, "read", diags))

	read_output, err := c.ReadSecurityCertRemoteCerts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityCertRemoteCert(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityCertRemoteCertModel) refreshSecurityCertRemoteCert(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["source"]; ok {
		m.Source = parseStringValue(v)
	}

	if v, ok := o["issuer"]; ok {
		m.Issuer = m.Issuer.flattenSecurityCertRemoteCertIssuer(ctx, v, &diags)
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
		m.Usages = m.flattenSecurityCertRemoteCertUsagesList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityCertRemoteCertModel) getURLObjectSecurityCertRemoteCert(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityCertRemoteCertIssuerModel struct {
	C            types.String `tfsdk:"c"`
	Cn           types.String `tfsdk:"cn"`
	L            types.String `tfsdk:"l"`
	O            types.String `tfsdk:"o"`
	Ou           types.String `tfsdk:"ou"`
	St           types.String `tfsdk:"st"`
	EmailAddress types.String `tfsdk:"email_address"`
}

type datasourceSecurityCertRemoteCertUsagesModel struct {
	Type  types.String  `tfsdk:"type"`
	Count types.Float64 `tfsdk:"count"`
}

func (m *datasourceSecurityCertRemoteCertIssuerModel) flattenSecurityCertRemoteCertIssuer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityCertRemoteCertIssuerModel {
	if input == nil {
		return &datasourceSecurityCertRemoteCertIssuerModel{}
	}
	if m == nil {
		m = &datasourceSecurityCertRemoteCertIssuerModel{}
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

func (m *datasourceSecurityCertRemoteCertUsagesModel) flattenSecurityCertRemoteCertUsages(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityCertRemoteCertUsagesModel {
	if input == nil {
		return &datasourceSecurityCertRemoteCertUsagesModel{}
	}
	if m == nil {
		m = &datasourceSecurityCertRemoteCertUsagesModel{}
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

func (s *datasourceSecurityCertRemoteCertModel) flattenSecurityCertRemoteCertUsagesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityCertRemoteCertUsagesModel {
	if o == nil {
		return []datasourceSecurityCertRemoteCertUsagesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument usages is not type of []interface{}.", "")
		return []datasourceSecurityCertRemoteCertUsagesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityCertRemoteCertUsagesModel{}
	}

	values := make([]datasourceSecurityCertRemoteCertUsagesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityCertRemoteCertUsagesModel
		if i < len(s.Usages) {
			m = s.Usages[i]
		}
		values[i] = *m.flattenSecurityCertRemoteCertUsages(ctx, ele, diags)
	}

	return values
}
