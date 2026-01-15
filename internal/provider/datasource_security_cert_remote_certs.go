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
var _ datasource.DataSource = &datasourceSecurityCertRemoteCerts{}

func newDatasourceSecurityCertRemoteCerts() datasource.DataSource {
	return &datasourceSecurityCertRemoteCerts{}
}

type datasourceSecurityCertRemoteCerts struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityCertRemoteCertsModel describes the datasource data model.
type datasourceSecurityCertRemoteCertsModel struct {
	Ftntid       types.Float64                                  `tfsdk:"ftntid"`
	Name         types.String                                   `tfsdk:"name"`
	PrimaryKey   types.String                                   `tfsdk:"primary_key"`
	Type         types.String                                   `tfsdk:"type"`
	Source       types.String                                   `tfsdk:"source"`
	Issuer       *datasourceSecurityCertRemoteCertsIssuerModel  `tfsdk:"issuer"`
	ValidFrom    types.String                                   `tfsdk:"valid_from"`
	ValidTo      types.String                                   `tfsdk:"valid_to"`
	SerialNumber types.String                                   `tfsdk:"serial_number"`
	Usages       []datasourceSecurityCertRemoteCertsUsagesModel `tfsdk:"usages"`
	CertName     types.String                                   `tfsdk:"cert_name"`
	FileContent  types.String                                   `tfsdk:"file_content"`
}

func (r *datasourceSecurityCertRemoteCerts) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_cert_remote_certs"
}

func (r *datasourceSecurityCertRemoteCerts) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
			"cert_name": schema.StringAttribute{
				Optional: true,
			},
			"file_content": schema.StringAttribute{
				Optional: true,
			},
			"issuer": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"c": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"cn": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"l": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"o": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"ou": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"st": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"email_address": schema.StringAttribute{
						Computed: true,
						Optional: true,
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
						},
						"count": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceSecurityCertRemoteCerts) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_cert_remote_certs"
}

func (r *datasourceSecurityCertRemoteCerts) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityCertRemoteCertsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityCertRemoteCerts(ctx, "read", diags))

	read_output, err := c.ReadSecurityCertRemoteCerts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityCertRemoteCerts(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityCertRemoteCertsModel) refreshSecurityCertRemoteCerts(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.Issuer = m.Issuer.flattenSecurityCertRemoteCertsIssuer(ctx, v, &diags)
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
		m.Usages = m.flattenSecurityCertRemoteCertsUsagesList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityCertRemoteCertsModel) getURLObjectSecurityCertRemoteCerts(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityCertRemoteCertsIssuerModel struct {
	C            types.String `tfsdk:"c"`
	Cn           types.String `tfsdk:"cn"`
	L            types.String `tfsdk:"l"`
	O            types.String `tfsdk:"o"`
	Ou           types.String `tfsdk:"ou"`
	St           types.String `tfsdk:"st"`
	EmailAddress types.String `tfsdk:"email_address"`
}

type datasourceSecurityCertRemoteCertsUsagesModel struct {
	Type  types.String  `tfsdk:"type"`
	Count types.Float64 `tfsdk:"count"`
}

func (m *datasourceSecurityCertRemoteCertsIssuerModel) flattenSecurityCertRemoteCertsIssuer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityCertRemoteCertsIssuerModel {
	if input == nil {
		return &datasourceSecurityCertRemoteCertsIssuerModel{}
	}
	if m == nil {
		m = &datasourceSecurityCertRemoteCertsIssuerModel{}
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

func (m *datasourceSecurityCertRemoteCertsUsagesModel) flattenSecurityCertRemoteCertsUsages(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityCertRemoteCertsUsagesModel {
	if input == nil {
		return &datasourceSecurityCertRemoteCertsUsagesModel{}
	}
	if m == nil {
		m = &datasourceSecurityCertRemoteCertsUsagesModel{}
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

func (s *datasourceSecurityCertRemoteCertsModel) flattenSecurityCertRemoteCertsUsagesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityCertRemoteCertsUsagesModel {
	if o == nil {
		return []datasourceSecurityCertRemoteCertsUsagesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument usages is not type of []interface{}.", "")
		return []datasourceSecurityCertRemoteCertsUsagesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityCertRemoteCertsUsagesModel{}
	}

	values := make([]datasourceSecurityCertRemoteCertsUsagesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityCertRemoteCertsUsagesModel
		values[i] = *m.flattenSecurityCertRemoteCertsUsages(ctx, ele, diags)
	}

	return values
}
