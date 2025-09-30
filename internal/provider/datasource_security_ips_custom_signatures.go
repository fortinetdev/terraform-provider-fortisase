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
var _ datasource.DataSource = &datasourceSecurityIpsCustomSignatures{}

func newDatasourceSecurityIpsCustomSignatures() datasource.DataSource {
	return &datasourceSecurityIpsCustomSignatures{}
}

type datasourceSecurityIpsCustomSignatures struct {
	fortiClient *FortiClient
}

// datasourceSecurityIpsCustomSignaturesModel describes the datasource data model.
type datasourceSecurityIpsCustomSignaturesModel struct {
	PrimaryKey  types.String  `tfsdk:"primary_key"`
	Tag         types.String  `tfsdk:"tag"`
	Signature   types.String  `tfsdk:"signature"`
	RuleId      types.Float64 `tfsdk:"rule_id"`
	Status      types.String  `tfsdk:"status"`
	Log         types.String  `tfsdk:"log"`
	LogPacket   types.String  `tfsdk:"log_packet"`
	Action      types.String  `tfsdk:"action"`
	Severity    types.String  `tfsdk:"severity"`
	Location    types.String  `tfsdk:"location"`
	Os          types.String  `tfsdk:"os"`
	Application types.String  `tfsdk:"application"`
	Protocol    types.String  `tfsdk:"protocol"`
	Comment     types.String  `tfsdk:"comment"`
}

func (r *datasourceSecurityIpsCustomSignatures) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ips_custom_signatures"
}

func (r *datasourceSecurityIpsCustomSignatures) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 63),
				},
				Required: true,
			},
			"tag": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"signature": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(4095),
				},
				Computed: true,
				Optional: true,
			},
			"rule_id": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"log": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"log_packet": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"action": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"severity": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"location": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"os": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"application": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"protocol": schema.StringAttribute{
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
		},
	}
}

func (r *datasourceSecurityIpsCustomSignatures) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityIpsCustomSignatures) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityIpsCustomSignaturesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityIpsCustomSignatures(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsCustomSignatures(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityIpsCustomSignatures(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityIpsCustomSignaturesModel) refreshSecurityIpsCustomSignatures(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

	if v, ok := o["signature"]; ok {
		m.Signature = parseStringValue(v)
	}

	if v, ok := o["ruleId"]; ok {
		m.RuleId = parseFloat64Value(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["log"]; ok {
		m.Log = parseStringValue(v)
	}

	if v, ok := o["logPacket"]; ok {
		m.LogPacket = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["severity"]; ok {
		m.Severity = parseStringValue(v)
	}

	if v, ok := o["location"]; ok {
		m.Location = parseStringValue(v)
	}

	if v, ok := o["os"]; ok {
		m.Os = parseStringValue(v)
	}

	if v, ok := o["application"]; ok {
		m.Application = parseStringValue(v)
	}

	if v, ok := o["protocol"]; ok {
		m.Protocol = parseStringValue(v)
	}

	if v, ok := o["comment"]; ok {
		m.Comment = parseStringValue(v)
	}

	return diags
}

func (data *datasourceSecurityIpsCustomSignaturesModel) getURLObjectSecurityIpsCustomSignatures(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
