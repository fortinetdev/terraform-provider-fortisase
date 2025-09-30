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
var _ datasource.DataSource = &datasourceSecurityAppCustomSignatures{}

func newDatasourceSecurityAppCustomSignatures() datasource.DataSource {
	return &datasourceSecurityAppCustomSignatures{}
}

type datasourceSecurityAppCustomSignatures struct {
	fortiClient *FortiClient
}

// datasourceSecurityAppCustomSignaturesModel describes the datasource data model.
type datasourceSecurityAppCustomSignaturesModel struct {
	PrimaryKey types.String  `tfsdk:"primary_key"`
	Signature  types.String  `tfsdk:"signature"`
	Comment    types.String  `tfsdk:"comment"`
	Ftntid     types.Float64 `tfsdk:"ftntid"`
	Tag        types.String  `tfsdk:"tag"`
	Name       types.String  `tfsdk:"name"`
	Category   types.Float64 `tfsdk:"category"`
	Protocol   types.String  `tfsdk:"protocol"`
	Technology types.String  `tfsdk:"technology"`
	Behavior   types.String  `tfsdk:"behavior"`
	Vendor     types.String  `tfsdk:"vendor"`
	IconClass  types.String  `tfsdk:"icon_class"`
}

func (r *datasourceSecurityAppCustomSignatures) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_app_custom_signatures"
}

func (r *datasourceSecurityAppCustomSignatures) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 63),
				},
				Required: true,
			},
			"signature": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(4095),
				},
				Computed: true,
				Optional: true,
			},
			"comment": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"ftntid": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"tag": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"category": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"protocol": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"technology": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"behavior": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"vendor": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"icon_class": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityAppCustomSignatures) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityAppCustomSignatures) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityAppCustomSignaturesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityAppCustomSignatures(ctx, "read", diags))

	read_output, err := c.ReadSecurityAppCustomSignatures(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityAppCustomSignatures(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityAppCustomSignaturesModel) refreshSecurityAppCustomSignatures(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["signature"]; ok {
		m.Signature = parseStringValue(v)
	}

	if v, ok := o["comment"]; ok {
		m.Comment = parseStringValue(v)
	}

	if v, ok := o["id"]; ok {
		m.Ftntid = parseFloat64Value(v)
	}

	if v, ok := o["tag"]; ok {
		m.Tag = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["category"]; ok {
		m.Category = parseFloat64Value(v)
	}

	if v, ok := o["protocol"]; ok {
		m.Protocol = parseStringValue(v)
	}

	if v, ok := o["technology"]; ok {
		m.Technology = parseStringValue(v)
	}

	if v, ok := o["behavior"]; ok {
		m.Behavior = parseStringValue(v)
	}

	if v, ok := o["vendor"]; ok {
		m.Vendor = parseStringValue(v)
	}

	if v, ok := o["iconClass"]; ok {
		m.IconClass = parseStringValue(v)
	}

	return diags
}

func (data *datasourceSecurityAppCustomSignaturesModel) getURLObjectSecurityAppCustomSignatures(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
