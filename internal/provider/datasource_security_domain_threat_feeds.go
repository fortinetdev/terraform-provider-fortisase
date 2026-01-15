// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityDomainThreatFeeds{}

func newDatasourceSecurityDomainThreatFeeds() datasource.DataSource {
	return &datasourceSecurityDomainThreatFeeds{}
}

type datasourceSecurityDomainThreatFeeds struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityDomainThreatFeedsModel describes the datasource data model.
type datasourceSecurityDomainThreatFeedsModel struct {
	PrimaryKey          types.String  `tfsdk:"primary_key"`
	Comments            types.String  `tfsdk:"comments"`
	Status              types.String  `tfsdk:"status"`
	RefreshRate         types.Float64 `tfsdk:"refresh_rate"`
	Uri                 types.String  `tfsdk:"uri"`
	BasicAuthentication types.String  `tfsdk:"basic_authentication"`
	Username            types.String  `tfsdk:"username"`
}

func (r *datasourceSecurityDomainThreatFeeds) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_domain_threat_feeds"
}

func (r *datasourceSecurityDomainThreatFeeds) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 35),
				},
				Required: true,
			},
			"comments": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
				Computed: true,
				Optional: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"refresh_rate": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(1, 43200),
				},
				Computed: true,
				Optional: true,
			},
			"uri": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"basic_authentication": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"username": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityDomainThreatFeeds) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_domain_threat_feeds"
}

func (r *datasourceSecurityDomainThreatFeeds) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityDomainThreatFeedsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDomainThreatFeeds(ctx, "read", diags))

	read_output, err := c.ReadSecurityDomainThreatFeeds(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDomainThreatFeeds(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityDomainThreatFeedsModel) refreshSecurityDomainThreatFeeds(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["refreshRate"]; ok {
		m.RefreshRate = parseFloat64Value(v)
	}

	if v, ok := o["uri"]; ok {
		m.Uri = parseStringValue(v)
	}

	if v, ok := o["basicAuthentication"]; ok {
		m.BasicAuthentication = parseStringValue(v)
	}

	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	return diags
}

func (data *datasourceSecurityDomainThreatFeedsModel) getURLObjectSecurityDomainThreatFeeds(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
