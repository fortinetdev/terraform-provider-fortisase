// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityPkiUsers{}

func newDatasourceSecurityPkiUsers() datasource.DataSource {
	return &datasourceSecurityPkiUsers{}
}

type datasourceSecurityPkiUsers struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityPkiUsersModel describes the datasource data model.
type datasourceSecurityPkiUsersModel struct {
	PrimaryKey     types.String                       `tfsdk:"primary_key"`
	Subject        types.String                       `tfsdk:"subject"`
	Ca             *datasourceSecurityPkiUsersCaModel `tfsdk:"ca"`
	IsStaticObject types.Bool                         `tfsdk:"is_static_object"`
	References     types.Float64                      `tfsdk:"references"`
	IsGlobalEntry  types.Bool                         `tfsdk:"is_global_entry"`
}

func (r *datasourceSecurityPkiUsers) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_pki_users"
}

func (r *datasourceSecurityPkiUsers) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "Primary Key of PKI User.",
				Required:            true,
			},
			"subject": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"is_static_object": schema.BoolAttribute{
				Computed: true,
			},
			"references": schema.Float64Attribute{
				Computed: true,
			},
			"is_global_entry": schema.BoolAttribute{
				Computed: true,
			},
			"ca": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: "CA Cert Name",
						Computed:            true,
						Optional:            true,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityPkiUsers) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_pki_users"
}

func (r *datasourceSecurityPkiUsers) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityPkiUsersModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityPkiUsers(ctx, "read", diags))

	read_output, err := c.ReadSecurityPkiUsers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityPkiUsers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityPkiUsersModel) refreshSecurityPkiUsers(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["subject"]; ok {
		m.Subject = parseStringValue(v)
	}

	if v, ok := o["ca"]; ok {
		m.Ca = m.Ca.flattenSecurityPkiUsersCa(ctx, v, &diags)
	}

	if v, ok := o["isStaticObject"]; ok {
		m.IsStaticObject = parseBoolValue(v)
	}

	if v, ok := o["references"]; ok {
		m.References = parseFloat64Value(v)
	}

	if v, ok := o["isGlobalEntry"]; ok {
		m.IsGlobalEntry = parseBoolValue(v)
	}

	return diags
}

func (data *datasourceSecurityPkiUsersModel) getURLObjectSecurityPkiUsers(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityPkiUsersCaModel struct {
	Name types.String `tfsdk:"name"`
}

func (m *datasourceSecurityPkiUsersCaModel) flattenSecurityPkiUsersCa(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityPkiUsersCaModel {
	if input == nil {
		return &datasourceSecurityPkiUsersCaModel{}
	}
	if m == nil {
		m = &datasourceSecurityPkiUsersCaModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	return m
}
