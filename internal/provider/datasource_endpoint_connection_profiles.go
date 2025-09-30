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
var _ datasource.DataSource = &datasourceEndpointConnectionProfiles{}

func newDatasourceEndpointConnectionProfiles() datasource.DataSource {
	return &datasourceEndpointConnectionProfiles{}
}

type datasourceEndpointConnectionProfiles struct {
	fortiClient *FortiClient
}

// datasourceEndpointConnectionProfilesModel describes the datasource data model.
type datasourceEndpointConnectionProfilesModel struct {
	ConnectToFortiSase             types.String                                                   `tfsdk:"connect_to_forti_sase"`
	Lockdown                       *datasourceEndpointConnectionProfilesLockdownModel             `tfsdk:"lockdown"`
	OnFabricRuleSet                *datasourceEndpointConnectionProfilesOnFabricRuleSetModel      `tfsdk:"on_fabric_rule_set"`
	OffNetSplitTunnel              *datasourceEndpointConnectionProfilesOffNetSplitTunnelModel    `tfsdk:"off_net_split_tunnel"`
	SplitTunnel                    *datasourceEndpointConnectionProfilesSplitTunnelModel          `tfsdk:"split_tunnel"`
	AllowInvalidServerCertificate  types.String                                                   `tfsdk:"allow_invalid_server_certificate"`
	EndpointOnNetBypass            types.Bool                                                     `tfsdk:"endpoint_on_net_bypass"`
	AuthBeforeUserLogon            types.Bool                                                     `tfsdk:"auth_before_user_logon"`
	SecureInternetAccess           *datasourceEndpointConnectionProfilesSecureInternetAccessModel `tfsdk:"secure_internet_access"`
	PreferredDtlsTunnel            types.String                                                   `tfsdk:"preferred_dtls_tunnel"`
	UseGuiSamlAuth                 types.String                                                   `tfsdk:"use_gui_saml_auth"`
	AllowPersonalVpns              types.Bool                                                     `tfsdk:"allow_personal_vpns"`
	MtuSize                        types.Float64                                                  `tfsdk:"mtu_size"`
	AvailableVpNs                  []datasourceEndpointConnectionProfilesAvailableVpNsModel       `tfsdk:"available_vp_ns"`
	ShowDisconnectBtn              types.String                                                   `tfsdk:"show_disconnect_btn"`
	EnableInvalidServerCertWarning types.String                                                   `tfsdk:"enable_invalid_server_cert_warning"`
	PreLogon                       *datasourceEndpointConnectionProfilesPreLogonModel             `tfsdk:"pre_logon"`
	PrimaryKey                     types.String                                                   `tfsdk:"primary_key"`
}

func (r *datasourceEndpointConnectionProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_connection_profiles"
}

func (r *datasourceEndpointConnectionProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"connect_to_forti_sase": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("automatically", "manually"),
				},
				Computed: true,
				Optional: true,
			},
			"allow_invalid_server_certificate": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"endpoint_on_net_bypass": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"auth_before_user_logon": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"preferred_dtls_tunnel": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"use_gui_saml_auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"allow_personal_vpns": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"mtu_size": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(576, 1500),
				},
				Computed: true,
				Optional: true,
			},
			"show_disconnect_btn": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"enable_invalid_server_cert_warning": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				Description: "The primary key of the object. Can be found in the response from the get request.",
				Required:    true,
			},
			"lockdown": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"grace_period": schema.Float64Attribute{
						Computed: true,
						Optional: true,
					},
					"max_attempts": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.AtLeast(1),
						},
						Computed: true,
						Optional: true,
					},
					"ips": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"port": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"protocol": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("tcp", "udp", "icmp", ""),
									},
									Computed: true,
									Optional: true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
					"domains": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
					"detect_captive_portal": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("enable", "disable"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"on_fabric_rule_set": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("endpoint/on-net-rules"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"off_net_split_tunnel": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"local_apps": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"fqdns": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"subnets_ipsec": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"isdbs": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("network/basic-internet-services"),
									},
									Computed: true,
									Optional: true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
					"subnets": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("network/hosts", "network/host-groups"),
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
				Computed: true,
				Optional: true,
			},
			"split_tunnel": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"local_apps": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"fqdns": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"subnets_ipsec": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"isdbs": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("network/basic-internet-services"),
									},
									Computed: true,
									Optional: true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
					"subnets": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("network/hosts", "network/host-groups"),
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
				Computed: true,
				Optional: true,
			},
			"secure_internet_access": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"authenticate_with_sso": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"enable_local_lan": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"failover_sequence": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"external_browser_saml_login": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"posture_check": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"tag": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"action": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("allow", "prohibit"),
								},
								Computed: true,
								Optional: true,
							},
							"check_failed_message": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"available_vp_ns": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("sslVPN", "ipSecVPN"),
							},
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"remote_gateway": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"username_prompt": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"save_username": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_always_up": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_auto_connect": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_remember_password": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"authenticate_with_sso": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"enable_local_lan": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"external_browser_saml_login": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validator.AtMost(65535),
							},
							Computed: true,
							Optional: true,
						},
						"require_certificate": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"auth_method": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("preSharedKey", "smartCardCert", "systemStoreCert"),
							},
							Computed: true,
							Optional: true,
						},
						"show_passcode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"pre_shared_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"posture_check": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"tag": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"action": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("allow", "prohibit"),
									},
									Computed: true,
									Optional: true,
								},
								"check_failed_message": schema.StringAttribute{
									Computed: true,
									Optional: true,
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
			"pre_logon": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"vpn_type": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("sslVPN", "ipSecVPN"),
						},
						Computed: true,
						Optional: true,
					},
					"remote_gateway": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"port": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.AtMost(65535),
						},
						Computed: true,
						Optional: true,
					},
					"common_name": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"match_type": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("wildcard", "regex"),
								},
								Computed: true,
								Optional: true,
							},
							"pattern": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
					"issuer": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"match_type": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("wildcard", "regex"),
								},
								Computed: true,
								Optional: true,
							},
							"pattern": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
								Computed: true,
								Optional: true,
							},
						},
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

func (r *datasourceEndpointConnectionProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceEndpointConnectionProfiles) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointConnectionProfilesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointConnectionProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointConnectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointConnectionProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointConnectionProfilesModel) refreshEndpointConnectionProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["connectToFortiSASE"]; ok {
		m.ConnectToFortiSase = parseStringValue(v)
	}

	if v, ok := o["lockdown"]; ok {
		m.Lockdown = m.Lockdown.flattenEndpointConnectionProfilesLockdown(ctx, v, &diags)
	}

	if v, ok := o["onFabricRuleSet"]; ok {
		m.OnFabricRuleSet = m.OnFabricRuleSet.flattenEndpointConnectionProfilesOnFabricRuleSet(ctx, v, &diags)
	}

	if v, ok := o["offNetSplitTunnel"]; ok {
		m.OffNetSplitTunnel = m.OffNetSplitTunnel.flattenEndpointConnectionProfilesOffNetSplitTunnel(ctx, v, &diags)
	}

	if v, ok := o["splitTunnel"]; ok {
		m.SplitTunnel = m.SplitTunnel.flattenEndpointConnectionProfilesSplitTunnel(ctx, v, &diags)
	}

	if v, ok := o["allowInvalidServerCertificate"]; ok {
		m.AllowInvalidServerCertificate = parseStringValue(v)
	}

	if v, ok := o["endpointOnNetBypass"]; ok {
		m.EndpointOnNetBypass = parseBoolValue(v)
	}

	if v, ok := o["authBeforeUserLogon"]; ok {
		m.AuthBeforeUserLogon = parseBoolValue(v)
	}

	if v, ok := o["secureInternetAccess"]; ok {
		m.SecureInternetAccess = m.SecureInternetAccess.flattenEndpointConnectionProfilesSecureInternetAccess(ctx, v, &diags)
	}

	if v, ok := o["preferredDTLSTunnel"]; ok {
		m.PreferredDtlsTunnel = parseStringValue(v)
	}

	if v, ok := o["useGuiSamlAuth"]; ok {
		m.UseGuiSamlAuth = parseStringValue(v)
	}

	if v, ok := o["allowPersonalVpns"]; ok {
		m.AllowPersonalVpns = parseBoolValue(v)
	}

	if v, ok := o["mtuSize"]; ok {
		m.MtuSize = parseFloat64Value(v)
	}

	if v, ok := o["availableVPNs"]; ok {
		m.AvailableVpNs = m.flattenEndpointConnectionProfilesAvailableVpNsList(ctx, v, &diags)
	}

	if v, ok := o["showDisconnectBtn"]; ok {
		m.ShowDisconnectBtn = parseStringValue(v)
	}

	if v, ok := o["enableInvalidServerCertWarning"]; ok {
		m.EnableInvalidServerCertWarning = parseStringValue(v)
	}

	if v, ok := o["preLogon"]; ok {
		m.PreLogon = m.PreLogon.flattenEndpointConnectionProfilesPreLogon(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointConnectionProfilesModel) getURLObjectEndpointConnectionProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointConnectionProfilesLockdownModel struct {
	Status              types.String                                                          `tfsdk:"status"`
	GracePeriod         types.Float64                                                         `tfsdk:"grace_period"`
	MaxAttempts         types.Float64                                                         `tfsdk:"max_attempts"`
	Ips                 []datasourceEndpointConnectionProfilesLockdownIpsModel                `tfsdk:"ips"`
	Domains             []datasourceEndpointConnectionProfilesLockdownDomainsModel            `tfsdk:"domains"`
	DetectCaptivePortal *datasourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel `tfsdk:"detect_captive_portal"`
}

type datasourceEndpointConnectionProfilesLockdownIpsModel struct {
	Ip       types.String `tfsdk:"ip"`
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
}

type datasourceEndpointConnectionProfilesLockdownDomainsModel struct {
	Address types.String `tfsdk:"address"`
}

type datasourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel struct {
	Status types.String `tfsdk:"status"`
}

type datasourceEndpointConnectionProfilesOnFabricRuleSetModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointConnectionProfilesOffNetSplitTunnelModel struct {
	LocalApps    types.Set                                                           `tfsdk:"local_apps"`
	Isdbs        []datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel   `tfsdk:"isdbs"`
	Fqdns        types.Set                                                           `tfsdk:"fqdns"`
	Subnets      []datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel `tfsdk:"subnets"`
	SubnetsIpsec types.Set                                                           `tfsdk:"subnets_ipsec"`
}

type datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointConnectionProfilesSplitTunnelModel struct {
	LocalApps    types.Set                                                     `tfsdk:"local_apps"`
	Isdbs        []datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel   `tfsdk:"isdbs"`
	Fqdns        types.Set                                                     `tfsdk:"fqdns"`
	Subnets      []datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel `tfsdk:"subnets"`
	SubnetsIpsec types.Set                                                     `tfsdk:"subnets_ipsec"`
}

type datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointConnectionProfilesSecureInternetAccessModel struct {
	AuthenticateWithSso      types.String                                                               `tfsdk:"authenticate_with_sso"`
	EnableLocalLan           types.String                                                               `tfsdk:"enable_local_lan"`
	FailoverSequence         types.Set                                                                  `tfsdk:"failover_sequence"`
	PostureCheck             *datasourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel `tfsdk:"posture_check"`
	ExternalBrowserSamlLogin types.String                                                               `tfsdk:"external_browser_saml_login"`
}

type datasourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel struct {
	Tag                types.String `tfsdk:"tag"`
	Action             types.String `tfsdk:"action"`
	CheckFailedMessage types.String `tfsdk:"check_failed_message"`
}

type datasourceEndpointConnectionProfilesAvailableVpNsModel struct {
	Type                     types.String                                                        `tfsdk:"type"`
	Name                     types.String                                                        `tfsdk:"name"`
	RemoteGateway            types.String                                                        `tfsdk:"remote_gateway"`
	UsernamePrompt           types.String                                                        `tfsdk:"username_prompt"`
	SaveUsername             types.String                                                        `tfsdk:"save_username"`
	ShowAlwaysUp             types.String                                                        `tfsdk:"show_always_up"`
	ShowAutoConnect          types.String                                                        `tfsdk:"show_auto_connect"`
	ShowRememberPassword     types.String                                                        `tfsdk:"show_remember_password"`
	AuthenticateWithSso      types.String                                                        `tfsdk:"authenticate_with_sso"`
	EnableLocalLan           types.String                                                        `tfsdk:"enable_local_lan"`
	ExternalBrowserSamlLogin types.String                                                        `tfsdk:"external_browser_saml_login"`
	Port                     types.Float64                                                       `tfsdk:"port"`
	RequireCertificate       types.String                                                        `tfsdk:"require_certificate"`
	AuthMethod               types.String                                                        `tfsdk:"auth_method"`
	ShowPasscode             types.String                                                        `tfsdk:"show_passcode"`
	PostureCheck             *datasourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel `tfsdk:"posture_check"`
	PreSharedKey             types.String                                                        `tfsdk:"pre_shared_key"`
}

type datasourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel struct {
	Tag                types.String `tfsdk:"tag"`
	Action             types.String `tfsdk:"action"`
	CheckFailedMessage types.String `tfsdk:"check_failed_message"`
}

type datasourceEndpointConnectionProfilesPreLogonModel struct {
	VpnType       types.String                                                 `tfsdk:"vpn_type"`
	RemoteGateway types.String                                                 `tfsdk:"remote_gateway"`
	CommonName    *datasourceEndpointConnectionProfilesPreLogonCommonNameModel `tfsdk:"common_name"`
	Issuer        *datasourceEndpointConnectionProfilesPreLogonIssuerModel     `tfsdk:"issuer"`
	Port          types.Float64                                                `tfsdk:"port"`
}

type datasourceEndpointConnectionProfilesPreLogonCommonNameModel struct {
	MatchType types.String `tfsdk:"match_type"`
	Pattern   types.String `tfsdk:"pattern"`
}

type datasourceEndpointConnectionProfilesPreLogonIssuerModel struct {
	MatchType types.String `tfsdk:"match_type"`
	Pattern   types.String `tfsdk:"pattern"`
}

func (m *datasourceEndpointConnectionProfilesLockdownModel) flattenEndpointConnectionProfilesLockdown(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesLockdownModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesLockdownModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesLockdownModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["gracePeriod"]; ok {
		m.GracePeriod = parseFloat64Value(v)
	}

	if v, ok := o["maxAttempts"]; ok {
		m.MaxAttempts = parseFloat64Value(v)
	}

	if v, ok := o["ips"]; ok {
		m.Ips = m.flattenEndpointConnectionProfilesLockdownIpsList(ctx, v, diags)
	}

	if v, ok := o["domains"]; ok {
		m.Domains = m.flattenEndpointConnectionProfilesLockdownDomainsList(ctx, v, diags)
	}

	if v, ok := o["detectCaptivePortal"]; ok {
		m.DetectCaptivePortal = m.DetectCaptivePortal.flattenEndpointConnectionProfilesLockdownDetectCaptivePortal(ctx, v, diags)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilesLockdownIpsModel) flattenEndpointConnectionProfilesLockdownIps(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesLockdownIpsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesLockdownIpsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesLockdownIpsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["ip"]; ok {
		m.Ip = parseStringValue(v)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseStringValue(v)
	}

	if v, ok := o["protocol"]; ok {
		m.Protocol = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointConnectionProfilesLockdownModel) flattenEndpointConnectionProfilesLockdownIpsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfilesLockdownIpsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfilesLockdownIpsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument ips is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfilesLockdownIpsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfilesLockdownIpsModel{}
	}

	values := make([]datasourceEndpointConnectionProfilesLockdownIpsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfilesLockdownIpsModel
		values[i] = *m.flattenEndpointConnectionProfilesLockdownIps(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfilesLockdownDomainsModel) flattenEndpointConnectionProfilesLockdownDomains(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesLockdownDomainsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesLockdownDomainsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesLockdownDomainsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["address"]; ok {
		m.Address = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointConnectionProfilesLockdownModel) flattenEndpointConnectionProfilesLockdownDomainsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfilesLockdownDomainsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfilesLockdownDomainsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument domains is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfilesLockdownDomainsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfilesLockdownDomainsModel{}
	}

	values := make([]datasourceEndpointConnectionProfilesLockdownDomainsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfilesLockdownDomainsModel
		values[i] = *m.flattenEndpointConnectionProfilesLockdownDomains(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel) flattenEndpointConnectionProfilesLockdownDetectCaptivePortal(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilesOnFabricRuleSetModel) flattenEndpointConnectionProfilesOnFabricRuleSet(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesOnFabricRuleSetModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesOnFabricRuleSetModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesOnFabricRuleSetModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilesOffNetSplitTunnelModel) flattenEndpointConnectionProfilesOffNetSplitTunnel(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesOffNetSplitTunnelModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesOffNetSplitTunnelModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesOffNetSplitTunnelModel{}
	}
	o := input.(map[string]interface{})
	m.LocalApps = types.SetNull(types.StringType)
	m.Fqdns = types.SetNull(types.StringType)
	m.SubnetsIpsec = types.SetNull(types.StringType)

	if v, ok := o["localApps"]; ok {
		m.LocalApps = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["isdbs"]; ok {
		m.Isdbs = m.flattenEndpointConnectionProfilesOffNetSplitTunnelIsdbsList(ctx, v, diags)
	}

	if v, ok := o["fqdns"]; ok {
		m.Fqdns = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["subnets"]; ok {
		m.Subnets = m.flattenEndpointConnectionProfilesOffNetSplitTunnelSubnetsList(ctx, v, diags)
	}

	if v, ok := o["subnetsIpsec"]; ok {
		m.SubnetsIpsec = parseSetValue(ctx, v, types.StringType)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel) flattenEndpointConnectionProfilesOffNetSplitTunnelIsdbs(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointConnectionProfilesOffNetSplitTunnelModel) flattenEndpointConnectionProfilesOffNetSplitTunnelIsdbsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument isdbs is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel{}
	}

	values := make([]datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel
		values[i] = *m.flattenEndpointConnectionProfilesOffNetSplitTunnelIsdbs(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel) flattenEndpointConnectionProfilesOffNetSplitTunnelSubnets(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointConnectionProfilesOffNetSplitTunnelModel) flattenEndpointConnectionProfilesOffNetSplitTunnelSubnetsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument subnets is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel{}
	}

	values := make([]datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel
		values[i] = *m.flattenEndpointConnectionProfilesOffNetSplitTunnelSubnets(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfilesSplitTunnelModel) flattenEndpointConnectionProfilesSplitTunnel(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesSplitTunnelModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesSplitTunnelModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesSplitTunnelModel{}
	}
	o := input.(map[string]interface{})
	m.LocalApps = types.SetNull(types.StringType)
	m.Fqdns = types.SetNull(types.StringType)
	m.SubnetsIpsec = types.SetNull(types.StringType)

	if v, ok := o["localApps"]; ok {
		m.LocalApps = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["isdbs"]; ok {
		m.Isdbs = m.flattenEndpointConnectionProfilesSplitTunnelIsdbsList(ctx, v, diags)
	}

	if v, ok := o["fqdns"]; ok {
		m.Fqdns = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["subnets"]; ok {
		m.Subnets = m.flattenEndpointConnectionProfilesSplitTunnelSubnetsList(ctx, v, diags)
	}

	if v, ok := o["subnetsIpsec"]; ok {
		m.SubnetsIpsec = parseSetValue(ctx, v, types.StringType)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel) flattenEndpointConnectionProfilesSplitTunnelIsdbs(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointConnectionProfilesSplitTunnelModel) flattenEndpointConnectionProfilesSplitTunnelIsdbsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument isdbs is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel{}
	}

	values := make([]datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfilesSplitTunnelIsdbsModel
		values[i] = *m.flattenEndpointConnectionProfilesSplitTunnelIsdbs(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel) flattenEndpointConnectionProfilesSplitTunnelSubnets(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointConnectionProfilesSplitTunnelModel) flattenEndpointConnectionProfilesSplitTunnelSubnetsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument subnets is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel{}
	}

	values := make([]datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfilesSplitTunnelSubnetsModel
		values[i] = *m.flattenEndpointConnectionProfilesSplitTunnelSubnets(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfilesSecureInternetAccessModel) flattenEndpointConnectionProfilesSecureInternetAccess(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesSecureInternetAccessModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesSecureInternetAccessModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesSecureInternetAccessModel{}
	}
	o := input.(map[string]interface{})
	m.FailoverSequence = types.SetNull(types.StringType)

	if v, ok := o["authenticateWithSSO"]; ok {
		m.AuthenticateWithSso = parseStringValue(v)
	}

	if v, ok := o["enableLocalLan"]; ok {
		m.EnableLocalLan = parseStringValue(v)
	}

	if v, ok := o["failoverSequence"]; ok {
		m.FailoverSequence = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["postureCheck"]; ok {
		m.PostureCheck = m.PostureCheck.flattenEndpointConnectionProfilesSecureInternetAccessPostureCheck(ctx, v, diags)
	}

	if v, ok := o["externalBrowserSamlLogin"]; ok {
		m.ExternalBrowserSamlLogin = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel) flattenEndpointConnectionProfilesSecureInternetAccessPostureCheck(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["tag"]; ok {
		m.Tag = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["checkFailedMessage"]; ok {
		m.CheckFailedMessage = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilesAvailableVpNsModel) flattenEndpointConnectionProfilesAvailableVpNs(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesAvailableVpNsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesAvailableVpNsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesAvailableVpNsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["remoteGateway"]; ok {
		m.RemoteGateway = parseStringValue(v)
	}

	if v, ok := o["usernamePrompt"]; ok {
		m.UsernamePrompt = parseStringValue(v)
	}

	if v, ok := o["saveUsername"]; ok {
		m.SaveUsername = parseStringValue(v)
	}

	if v, ok := o["showAlwaysUp"]; ok {
		m.ShowAlwaysUp = parseStringValue(v)
	}

	if v, ok := o["showAutoConnect"]; ok {
		m.ShowAutoConnect = parseStringValue(v)
	}

	if v, ok := o["showRememberPassword"]; ok {
		m.ShowRememberPassword = parseStringValue(v)
	}

	if v, ok := o["authenticateWithSSO"]; ok {
		m.AuthenticateWithSso = parseStringValue(v)
	}

	if v, ok := o["enableLocalLan"]; ok {
		m.EnableLocalLan = parseStringValue(v)
	}

	if v, ok := o["externalBrowserSamlLogin"]; ok {
		m.ExternalBrowserSamlLogin = parseStringValue(v)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	if v, ok := o["requireCertificate"]; ok {
		m.RequireCertificate = parseStringValue(v)
	}

	if v, ok := o["authMethod"]; ok {
		m.AuthMethod = parseStringValue(v)
	}

	if v, ok := o["showPasscode"]; ok {
		m.ShowPasscode = parseStringValue(v)
	}

	if v, ok := o["postureCheck"]; ok {
		m.PostureCheck = m.PostureCheck.flattenEndpointConnectionProfilesAvailableVpNsPostureCheck(ctx, v, diags)
	}

	if v, ok := o["preSharedKey"]; ok {
		m.PreSharedKey = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointConnectionProfilesModel) flattenEndpointConnectionProfilesAvailableVpNsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfilesAvailableVpNsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfilesAvailableVpNsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument available_vp_ns is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfilesAvailableVpNsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfilesAvailableVpNsModel{}
	}

	values := make([]datasourceEndpointConnectionProfilesAvailableVpNsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfilesAvailableVpNsModel
		values[i] = *m.flattenEndpointConnectionProfilesAvailableVpNs(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel) flattenEndpointConnectionProfilesAvailableVpNsPostureCheck(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["tag"]; ok {
		m.Tag = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["checkFailedMessage"]; ok {
		m.CheckFailedMessage = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilesPreLogonModel) flattenEndpointConnectionProfilesPreLogon(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesPreLogonModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesPreLogonModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesPreLogonModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["vpnType"]; ok {
		m.VpnType = parseStringValue(v)
	}

	if v, ok := o["remoteGateway"]; ok {
		m.RemoteGateway = parseStringValue(v)
	}

	if v, ok := o["commonName"]; ok {
		m.CommonName = m.CommonName.flattenEndpointConnectionProfilesPreLogonCommonName(ctx, v, diags)
	}

	if v, ok := o["issuer"]; ok {
		m.Issuer = m.Issuer.flattenEndpointConnectionProfilesPreLogonIssuer(ctx, v, diags)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilesPreLogonCommonNameModel) flattenEndpointConnectionProfilesPreLogonCommonName(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesPreLogonCommonNameModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesPreLogonCommonNameModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesPreLogonCommonNameModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["matchType"]; ok {
		m.MatchType = parseStringValue(v)
	}

	if v, ok := o["pattern"]; ok {
		m.Pattern = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilesPreLogonIssuerModel) flattenEndpointConnectionProfilesPreLogonIssuer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilesPreLogonIssuerModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilesPreLogonIssuerModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilesPreLogonIssuerModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["matchType"]; ok {
		m.MatchType = parseStringValue(v)
	}

	if v, ok := o["pattern"]; ok {
		m.Pattern = parseStringValue(v)
	}

	return m
}
