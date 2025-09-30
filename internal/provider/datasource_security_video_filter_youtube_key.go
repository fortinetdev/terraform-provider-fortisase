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
var _ datasource.DataSource = &datasourceSecurityVideoFilterYoutubeKey{}

func newDatasourceSecurityVideoFilterYoutubeKey() datasource.DataSource {
	return &datasourceSecurityVideoFilterYoutubeKey{}
}

type datasourceSecurityVideoFilterYoutubeKey struct {
	fortiClient *FortiClient
}

// datasourceSecurityVideoFilterYoutubeKeyModel describes the datasource data model.
type datasourceSecurityVideoFilterYoutubeKeyModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	ApiKey     types.String `tfsdk:"api_key"`
}

func (r *datasourceSecurityVideoFilterYoutubeKey) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_video_filter_youtube_key"
}

func (r *datasourceSecurityVideoFilterYoutubeKey) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("$sase-global"),
				},
				Required: true,
			},
			"api_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 47),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityVideoFilterYoutubeKey) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityVideoFilterYoutubeKey) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityVideoFilterYoutubeKeyModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadSecurityVideoFilterYoutubeKey(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityVideoFilterYoutubeKey(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityVideoFilterYoutubeKeyModel) refreshSecurityVideoFilterYoutubeKey(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["apiKey"]; ok {
		m.ApiKey = parseStringValue(v)
	}

	return diags
}
