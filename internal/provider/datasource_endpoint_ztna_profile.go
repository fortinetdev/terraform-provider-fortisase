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
var _ datasource.DataSource = &datasourceEndpointZtnaProfile{}

func newDatasourceEndpointZtnaProfile() datasource.DataSource {
	return &datasourceEndpointZtnaProfile{}
}

type datasourceEndpointZtnaProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointZtnaProfileModel describes the datasource data model.
type datasourceEndpointZtnaProfileModel struct {
	AllowAutomaticSignOn types.String                                        `tfsdk:"allow_automatic_sign_on"`
	Status               types.String                                        `tfsdk:"status"`
	ConnectionRules      []datasourceEndpointZtnaProfileConnectionRulesModel `tfsdk:"connection_rules"`
	EntraId              *datasourceEndpointZtnaProfileEntraIdModel          `tfsdk:"entra_id"`
	PrimaryKey           types.String                                        `tfsdk:"primary_key"`
}

func (r *datasourceEndpointZtnaProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_profile"
}

func (r *datasourceEndpointZtnaProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ZTNA Profile Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"allow_automatic_sign_on": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
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
						},
						"address": schema.StringAttribute{
							Computed: true,
						},
						"uid": schema.StringAttribute{
							Computed: true,
						},
						"mask": schema.StringAttribute{
							Computed: true,
						},
						"port": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"encryption": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"gateways": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.Float64Attribute{
										Computed: true,
									},
									"alias": schema.StringAttribute{
										Computed: true,
									},
									"private_app_count": schema.Float64Attribute{
										Computed: true,
									},
									"vip": schema.StringAttribute{
										Computed: true,
									},
									"redirect": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidatorwarning.OneOf("enable", "disable"),
										},
										Computed: true,
									},
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"entra_id": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"application_id": schema.StringAttribute{
						Computed: true,
					},
					"domain_name": schema.StringAttribute{
						Computed: true,
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceEndpointZtnaProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_ztna_profile"
}

func (r *datasourceEndpointZtnaProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointZtnaProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointZtnaProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointZtnaProfileModel) refreshEndpointZtnaProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["allowAutomaticSignOn"]; ok {
		m.AllowAutomaticSignOn = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["connectionRules"]; ok {
		m.ConnectionRules = m.flattenEndpointZtnaProfileConnectionRulesList(ctx, v, &diags)
	}

	if v, ok := o["entraId"]; ok {
		m.EntraId = m.EntraId.flattenEndpointZtnaProfileEntraId(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointZtnaProfileModel) getURLObjectEndpointZtnaProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointZtnaProfileConnectionRulesModel struct {
	Id         types.Float64                                               `tfsdk:"id"`
	Address    types.String                                                `tfsdk:"address"`
	Uid        types.String                                                `tfsdk:"uid"`
	Gateways   []datasourceEndpointZtnaProfileConnectionRulesGatewaysModel `tfsdk:"gateways"`
	Mask       types.String                                                `tfsdk:"mask"`
	Port       types.String                                                `tfsdk:"port"`
	Name       types.String                                                `tfsdk:"name"`
	Encryption types.String                                                `tfsdk:"encryption"`
}

type datasourceEndpointZtnaProfileConnectionRulesGatewaysModel struct {
	Id              types.Float64 `tfsdk:"id"`
	Alias           types.String  `tfsdk:"alias"`
	PrivateAppCount types.Float64 `tfsdk:"private_app_count"`
	Vip             types.String  `tfsdk:"vip"`
	Redirect        types.String  `tfsdk:"redirect"`
}

type datasourceEndpointZtnaProfileEntraIdModel struct {
	ApplicationId types.String `tfsdk:"application_id"`
	DomainName    types.String `tfsdk:"domain_name"`
}

func (m *datasourceEndpointZtnaProfileConnectionRulesModel) flattenEndpointZtnaProfileConnectionRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaProfileConnectionRulesModel {
	if input == nil {
		return &datasourceEndpointZtnaProfileConnectionRulesModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaProfileConnectionRulesModel{}
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
		m.Gateways = m.flattenEndpointZtnaProfileConnectionRulesGatewaysList(ctx, v, diags)
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

func (s *datasourceEndpointZtnaProfileModel) flattenEndpointZtnaProfileConnectionRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointZtnaProfileConnectionRulesModel {
	if o == nil {
		return []datasourceEndpointZtnaProfileConnectionRulesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument connection_rules is not type of []interface{}.", "")
		return []datasourceEndpointZtnaProfileConnectionRulesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointZtnaProfileConnectionRulesModel{}
	}

	values := make([]datasourceEndpointZtnaProfileConnectionRulesModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointZtnaProfileConnectionRulesModel
		if i < len(s.ConnectionRules) {
			m = s.ConnectionRules[i]
		}
		values[i] = *m.flattenEndpointZtnaProfileConnectionRules(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointZtnaProfileConnectionRulesGatewaysModel) flattenEndpointZtnaProfileConnectionRulesGateways(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaProfileConnectionRulesGatewaysModel {
	if input == nil {
		return &datasourceEndpointZtnaProfileConnectionRulesGatewaysModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaProfileConnectionRulesGatewaysModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

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

func (s *datasourceEndpointZtnaProfileConnectionRulesModel) flattenEndpointZtnaProfileConnectionRulesGatewaysList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointZtnaProfileConnectionRulesGatewaysModel {
	if o == nil {
		return []datasourceEndpointZtnaProfileConnectionRulesGatewaysModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument gateways is not type of []interface{}.", "")
		return []datasourceEndpointZtnaProfileConnectionRulesGatewaysModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointZtnaProfileConnectionRulesGatewaysModel{}
	}

	values := make([]datasourceEndpointZtnaProfileConnectionRulesGatewaysModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointZtnaProfileConnectionRulesGatewaysModel
		if i < len(s.Gateways) {
			m = s.Gateways[i]
		}
		values[i] = *m.flattenEndpointZtnaProfileConnectionRulesGateways(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointZtnaProfileEntraIdModel) flattenEndpointZtnaProfileEntraId(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaProfileEntraIdModel {
	if input == nil {
		return &datasourceEndpointZtnaProfileEntraIdModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaProfileEntraIdModel{}
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
