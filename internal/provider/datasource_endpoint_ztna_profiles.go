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
var _ datasource.DataSource = &datasourceEndpointZtnaProfiles{}

func newDatasourceEndpointZtnaProfiles() datasource.DataSource {
	return &datasourceEndpointZtnaProfiles{}
}

type datasourceEndpointZtnaProfiles struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointZtnaProfilesModel describes the datasource data model.
type datasourceEndpointZtnaProfilesModel struct {
	AllowAutomaticSignOn types.String                                         `tfsdk:"allow_automatic_sign_on"`
	ConnectionRules      []datasourceEndpointZtnaProfilesConnectionRulesModel `tfsdk:"connection_rules"`
	EntraId              *datasourceEndpointZtnaProfilesEntraIdModel          `tfsdk:"entra_id"`
	PrimaryKey           types.String                                         `tfsdk:"primary_key"`
}

func (r *datasourceEndpointZtnaProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_profiles"
}

func (r *datasourceEndpointZtnaProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"allow_automatic_sign_on": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
			},
			"connection_rules": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
						"address": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"uid": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"mask": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"port": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"encryption": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"gateways": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"alias": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"private_app_count": schema.Float64Attribute{
										Computed: true,
										Optional: true,
									},
									"vip": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"redirect": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("enable", "disable"),
										},
										Computed: true,
										Optional: true,
									},
								},
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"entra_id": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"application_id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"domain_name": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceEndpointZtnaProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_ztna_profiles"
}

func (r *datasourceEndpointZtnaProfiles) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointZtnaProfilesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointZtnaProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointZtnaProfilesModel) refreshEndpointZtnaProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["allowAutomaticSignOn"]; ok {
		m.AllowAutomaticSignOn = parseStringValue(v)
	}

	if v, ok := o["connectionRules"]; ok {
		m.ConnectionRules = m.flattenEndpointZtnaProfilesConnectionRulesList(ctx, v, &diags)
	}

	if v, ok := o["entraId"]; ok {
		m.EntraId = m.EntraId.flattenEndpointZtnaProfilesEntraId(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointZtnaProfilesModel) getURLObjectEndpointZtnaProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointZtnaProfilesConnectionRulesModel struct {
	Id         types.Float64                                                `tfsdk:"id"`
	Address    types.String                                                 `tfsdk:"address"`
	Uid        types.String                                                 `tfsdk:"uid"`
	Gateways   []datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel `tfsdk:"gateways"`
	Mask       types.String                                                 `tfsdk:"mask"`
	Port       types.String                                                 `tfsdk:"port"`
	Name       types.String                                                 `tfsdk:"name"`
	Encryption types.String                                                 `tfsdk:"encryption"`
}

type datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel struct {
	Alias           types.String  `tfsdk:"alias"`
	PrivateAppCount types.Float64 `tfsdk:"private_app_count"`
	Vip             types.String  `tfsdk:"vip"`
	Redirect        types.String  `tfsdk:"redirect"`
}

type datasourceEndpointZtnaProfilesEntraIdModel struct {
	ApplicationId types.String `tfsdk:"application_id"`
	DomainName    types.String `tfsdk:"domain_name"`
}

func (m *datasourceEndpointZtnaProfilesConnectionRulesModel) flattenEndpointZtnaProfilesConnectionRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaProfilesConnectionRulesModel {
	if input == nil {
		return &datasourceEndpointZtnaProfilesConnectionRulesModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaProfilesConnectionRulesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	if v, ok := o["address"]; ok {
		m.Address = parseStringValue(v)
	}

	if v, ok := o["uid"]; ok {
		m.Uid = parseStringValue(v)
	}

	if v, ok := o["gateways"]; ok {
		m.Gateways = m.flattenEndpointZtnaProfilesConnectionRulesGatewaysList(ctx, v, diags)
	}

	if v, ok := o["mask"]; ok {
		m.Mask = parseStringValue(v)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["encryption"]; ok {
		m.Encryption = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointZtnaProfilesModel) flattenEndpointZtnaProfilesConnectionRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointZtnaProfilesConnectionRulesModel {
	if o == nil {
		return []datasourceEndpointZtnaProfilesConnectionRulesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument connection_rules is not type of []interface{}.", "")
		return []datasourceEndpointZtnaProfilesConnectionRulesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointZtnaProfilesConnectionRulesModel{}
	}

	values := make([]datasourceEndpointZtnaProfilesConnectionRulesModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointZtnaProfilesConnectionRulesModel
		values[i] = *m.flattenEndpointZtnaProfilesConnectionRules(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel) flattenEndpointZtnaProfilesConnectionRulesGateways(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel {
	if input == nil {
		return &datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["alias"]; ok {
		m.Alias = parseStringValue(v)
	}

	if v, ok := o["private_app_count"]; ok {
		m.PrivateAppCount = parseFloat64Value(v)
	}

	if v, ok := o["vip"]; ok {
		m.Vip = parseStringValue(v)
	}

	if v, ok := o["redirect"]; ok {
		m.Redirect = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointZtnaProfilesConnectionRulesModel) flattenEndpointZtnaProfilesConnectionRulesGatewaysList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel {
	if o == nil {
		return []datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument gateways is not type of []interface{}.", "")
		return []datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel{}
	}

	values := make([]datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointZtnaProfilesConnectionRulesGatewaysModel
		values[i] = *m.flattenEndpointZtnaProfilesConnectionRulesGateways(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointZtnaProfilesEntraIdModel) flattenEndpointZtnaProfilesEntraId(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaProfilesEntraIdModel {
	if input == nil {
		return &datasourceEndpointZtnaProfilesEntraIdModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaProfilesEntraIdModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["applicationId"]; ok {
		m.ApplicationId = parseStringValue(v)
	}

	if v, ok := o["domainName"]; ok {
		m.DomainName = parseStringValue(v)
	}

	return m
}
