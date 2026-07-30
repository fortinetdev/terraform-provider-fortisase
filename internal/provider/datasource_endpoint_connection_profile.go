// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointConnectionProfile{}

func newDatasourceEndpointConnectionProfile() datasource.DataSource {
	return &datasourceEndpointConnectionProfile{}
}

type datasourceEndpointConnectionProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointConnectionProfileModel describes the datasource data model.
type datasourceEndpointConnectionProfileModel struct {
	ConnectToFortisase             types.String                                                  `tfsdk:"connect_to_fortisase"`
	ConnectToFortiSase             types.String                                                  `tfsdk:"connect_to_forti_sase"`
	AvailableVpns                  []datasourceEndpointConnectionProfileAvailableVpnsModel       `tfsdk:"available_vpns"`
	AvailableVpNs                  []datasourceEndpointConnectionProfileAvailableVpnsModel       `tfsdk:"available_vp_ns"`
	Lockdown                       *datasourceEndpointConnectionProfileLockdownModel             `tfsdk:"lockdown"`
	OnFabricRuleSet                *datasourceEndpointConnectionProfileOnFabricRuleSetModel      `tfsdk:"on_fabric_rule_set"`
	OffNetSplitTunnel              *datasourceEndpointConnectionProfileOffNetSplitTunnelModel    `tfsdk:"off_net_split_tunnel"`
	SplitTunnel                    *datasourceEndpointConnectionProfileSplitTunnelModel          `tfsdk:"split_tunnel"`
	AllowInvalidServerCertificate  types.String                                                  `tfsdk:"allow_invalid_server_certificate"`
	EndpointOnNetBypass            types.Bool                                                    `tfsdk:"endpoint_on_net_bypass"`
	AuthBeforeUserLogon            types.Bool                                                    `tfsdk:"auth_before_user_logon"`
	SecureInternetAccess           *datasourceEndpointConnectionProfileSecureInternetAccessModel `tfsdk:"secure_internet_access"`
	PreferredDtlsTunnel            types.String                                                  `tfsdk:"preferred_dtls_tunnel"`
	UseGuiSamlAuth                 types.String                                                  `tfsdk:"use_gui_saml_auth"`
	UseWebview2SamlAuth            types.String                                                  `tfsdk:"use_webview2_saml_auth"`
	BeforeLogonSamlAuth            types.String                                                  `tfsdk:"before_logon_saml_auth"`
	AfterLogonSamlAuth             types.String                                                  `tfsdk:"after_logon_saml_auth"`
	AllowPersonalVpns              types.Bool                                                    `tfsdk:"allow_personal_vpns"`
	MtuSize                        types.Float64                                                 `tfsdk:"mtu_size"`
	VpnType                        types.String                                                  `tfsdk:"vpn_type"`
	DisableInternetCheck           types.String                                                  `tfsdk:"disable_internet_check"`
	ShowDisconnectBtn              types.String                                                  `tfsdk:"show_disconnect_btn"`
	EnableInvalidServerCertWarning types.String                                                  `tfsdk:"enable_invalid_server_cert_warning"`
	PreLogon                       *datasourceEndpointConnectionProfilePreLogonModel             `tfsdk:"pre_logon"`
	PrimaryKey                     types.String                                                  `tfsdk:"primary_key"`
}

func (r *datasourceEndpointConnectionProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_connection_profile"
}

func (r *datasourceEndpointConnectionProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Connection Profile Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"connect_to_fortisase": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("automatically", "manually"),
				},
				Computed: true,
			},
			"connect_to_forti_sase": schema.StringAttribute{
				DeprecationMessage: "\"connect_to_forti_sase\" is deprecated; use \"connect_to_fortisase\" instead.",
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("automatically", "manually"),
				},
				Computed: true,
			},
			"allow_invalid_server_certificate": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"endpoint_on_net_bypass": schema.BoolAttribute{
				Computed: true,
			},
			"auth_before_user_logon": schema.BoolAttribute{
				Computed: true,
			},
			"preferred_dtls_tunnel": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"use_gui_saml_auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"use_webview2_saml_auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"before_logon_saml_auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("webBrowser", "electron"),
				},
				MarkdownDescription: "Specifies the browser framework used for Pre-logon VPN SAML authentication.\nSupported values: webBrowser, electron.",
				Computed:            true,
			},
			"after_logon_saml_auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("webBrowser", "electron", "webView2"),
				},
				MarkdownDescription: "Specifies the browser framework used for normal VPN SAML authentication.\nSupported values: webBrowser, electron, webView2.",
				Computed:            true,
			},
			"allow_personal_vpns": schema.BoolAttribute{
				Computed: true,
			},
			"mtu_size": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.Between(576, 1500),
				},
				Computed: true,
			},
			"vpn_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("sslVPN", "ipSecVPN"),
				},
				Computed: true,
			},
			"disable_internet_check": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"show_disconnect_btn": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"enable_invalid_server_cert_warning": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
			},
			"available_vpns": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("sslVPN", "ipSecVPN"),
							},
							Computed: true,
						},
						"name": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
						},
						"remote_gateway": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
						},
						"username_prompt": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"save_username": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"show_always_up": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"show_auto_connect": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"show_remember_password": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"authenticate_with_sso": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"allow_fido_auth": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"enable_local_lan": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"encapsulation_mode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("Auto", "TCP", "UDP"),
							},
							Computed: true,
						},
						"udp_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(500, 65535),
							},
							Computed: true,
						},
						"tcp_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(1, 65535),
							},
							Computed: true,
						},
						"port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtMost(65535),
							},
							Computed: true,
						},
						"require_certificate": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"external_browser_saml_login": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"auth_method": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("preSharedKey", "smartCardCert", "systemStoreCert"),
							},
							Computed: true,
						},
						"dns_suffixes": schema.SetAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"show_passcode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"eap_enabled": schema.BoolAttribute{
							MarkdownDescription: "Per-tunnel EAP for this manual IPsec VPN entry in availableVPNs.",
							Computed:            true,
						},
						"saml_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtMost(65535),
							},
							Computed: true,
						},
						"pre_shared_key": schema.StringAttribute{
							Computed: true,
						},
						"connect_disconnect_scripts": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"on_connect_windows": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
								},
								"on_connect_mac": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
								},
								"on_disconnect_windows": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
								},
								"on_disconnect_mac": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
								},
							},
							Computed: true,
						},
						"posture_check": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"tag": schema.StringAttribute{
									Computed: true,
								},
								"action": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("allow", "prohibit"),
									},
									Computed: true,
								},
								"check_failed_message": schema.StringAttribute{
									Computed: true,
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"available_vp_ns": schema.ListNestedAttribute{
				DeprecationMessage: "\"available_vp_ns\" is deprecated; use \"available_vpns\" instead.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("sslVPN", "ipSecVPN"),
							},
							Computed: true,
						},
						"name": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
						},
						"remote_gateway": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
						},
						"username_prompt": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"save_username": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"show_always_up": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"show_auto_connect": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"show_remember_password": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"authenticate_with_sso": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"allow_fido_auth": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"enable_local_lan": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"encapsulation_mode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("Auto", "TCP", "UDP"),
							},
							Computed: true,
						},
						"udp_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(500, 65535),
							},
							Computed: true,
						},
						"tcp_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(1, 65535),
							},
							Computed: true,
						},
						"port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtMost(65535),
							},
							Computed: true,
						},
						"require_certificate": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"external_browser_saml_login": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"auth_method": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("preSharedKey", "smartCardCert", "systemStoreCert"),
							},
							Computed: true,
						},
						"dns_suffixes": schema.SetAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"show_passcode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"eap_enabled": schema.BoolAttribute{
							MarkdownDescription: "Per-tunnel EAP for this manual IPsec VPN entry in availableVPNs.",
							Computed:            true,
						},
						"saml_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtMost(65535),
							},
							Computed: true,
						},
						"pre_shared_key": schema.StringAttribute{
							Computed: true,
						},
						"connect_disconnect_scripts": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"on_connect_windows": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
								},
								"on_connect_mac": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
								},
								"on_disconnect_windows": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
								},
								"on_disconnect_mac": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
								},
							},
							Computed: true,
						},
						"posture_check": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"tag": schema.StringAttribute{
									Computed: true,
								},
								"action": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("allow", "prohibit"),
									},
									Computed: true,
								},
								"check_failed_message": schema.StringAttribute{
									Computed: true,
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"lockdown": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
					},
					"grace_period": schema.Float64Attribute{
						Computed: true,
					},
					"max_attempts": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validatorwarning.AtLeast(1),
						},
						Computed: true,
					},
					"ips": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Computed: true,
								},
								"port": schema.StringAttribute{
									Computed: true,
								},
								"protocol": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("tcp", "udp", "icmp", ""),
									},
									Computed: true,
								},
							},
						},
						Computed: true,
					},
					"domains": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						Computed: true,
					},
					"detect_captive_portal": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.OneOf("enable", "disable"),
								},
								Computed: true,
							},
							"disable_windows_captive_portal": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.OneOf("enable", "disable"),
								},
								Computed: true,
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"on_fabric_rule_set": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("endpoint/on-net-rules"),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"off_net_split_tunnel": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"local_apps": schema.SetAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"fqdns": schema.SetAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"subnets_ipsec": schema.SetAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"split_tunnel_mode": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("exclude", "include"),
						},
						Computed: true,
					},
					"isdbs": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("network/basic-internet-services"),
									},
									Computed: true,
								},
							},
						},
						Computed: true,
					},
					"subnets": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("network/hosts", "network/host-groups"),
									},
									Computed: true,
								},
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"split_tunnel": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"local_apps": schema.SetAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"fqdns": schema.SetAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"subnets_ipsec": schema.SetAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"split_tunnel_mode": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("exclude", "include"),
						},
						Computed: true,
					},
					"isdbs": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("network/basic-internet-services"),
									},
									Computed: true,
								},
							},
						},
						Computed: true,
					},
					"subnets": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("network/hosts", "network/host-groups"),
									},
									Computed: true,
								},
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"secure_internet_access": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"authenticate_with_sso": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
					},
					"allow_fido_auth": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
					},
					"dns_suffixes": schema.SetAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"enable_local_lan": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
					},
					"failover_sequence": schema.SetAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"eap_enabled": schema.BoolAttribute{
						MarkdownDescription: "When vpnType is ipSecVPN, sets EAP (eap_method) on the Secure Internet Access tunnel(s) only (SIA-named connections), for both on-net and off-net EMS profiles. Custom/manual IPsec tunnels use availableVPNs[].eapEnabled.",
						Computed:            true,
					},
					"encapsulation_mode": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("Auto", "TCP", "UDP"),
						},
						Computed: true,
					},
					"external_browser_saml_login": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
					},
					"connect_disconnect_scripts": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"on_connect_windows": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtMost(1023),
								},
								Computed: true,
							},
							"on_connect_mac": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtMost(1023),
								},
								Computed: true,
							},
							"on_disconnect_windows": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtMost(1023),
								},
								Computed: true,
							},
							"on_disconnect_mac": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtMost(1023),
								},
								Computed: true,
							},
						},
						Computed: true,
					},
					"posture_check": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"tag": schema.StringAttribute{
								Computed: true,
							},
							"action": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.OneOf("allow", "prohibit"),
								},
								Computed: true,
							},
							"check_failed_message": schema.StringAttribute{
								Computed: true,
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"pre_logon": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"vpn_type": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("sslVPN", "ipSecVPN"),
						},
						Computed: true,
					},
					"remote_gateway": schema.StringAttribute{
						Computed: true,
					},
					"port": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validatorwarning.AtMost(65535),
						},
						Computed: true,
					},
					"common_name": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"match_type": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.OneOf("wildcard", "regex"),
								},
								Computed: true,
							},
							"pattern": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtLeast(1),
								},
								Computed: true,
							},
						},
						Computed: true,
					},
					"issuer": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"match_type": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.OneOf("wildcard", "regex"),
								},
								Computed: true,
							},
							"pattern": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtLeast(1),
								},
								Computed: true,
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceEndpointConnectionProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_connection_profile"
}

func (r *datasourceEndpointConnectionProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointConnectionProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointConnectionProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointConnectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointConnectionProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointConnectionProfileModel) refreshEndpointConnectionProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["connectToFortiSASE"]; ok {
		connectToFortisaseValue := parseStringValue(v)
		m.ConnectToFortisase = connectToFortisaseValue
		m.ConnectToFortiSase = connectToFortisaseValue
	}

	if v, ok := o["availableVPNs"]; ok {
		if m.AvailableVpns == nil {
			if m.AvailableVpNs != nil {
				m.AvailableVpns = m.AvailableVpNs
			}
		}
		AvailableVpnsValue := m.flattenEndpointConnectionProfileAvailableVpnsList(ctx, v, &diags)
		m.AvailableVpns = AvailableVpnsValue
		m.AvailableVpNs = AvailableVpnsValue
	}

	if v, ok := o["lockdown"]; ok {
		m.Lockdown = m.Lockdown.flattenEndpointConnectionProfileLockdown(ctx, v, &diags)
	}

	if v, ok := o["onFabricRuleSet"]; ok {
		m.OnFabricRuleSet = m.OnFabricRuleSet.flattenEndpointConnectionProfileOnFabricRuleSet(ctx, v, &diags)
	}

	if v, ok := o["offNetSplitTunnel"]; ok {
		m.OffNetSplitTunnel = m.OffNetSplitTunnel.flattenEndpointConnectionProfileOffNetSplitTunnel(ctx, v, &diags)
	}

	if v, ok := o["splitTunnel"]; ok {
		m.SplitTunnel = m.SplitTunnel.flattenEndpointConnectionProfileSplitTunnel(ctx, v, &diags)
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
		m.SecureInternetAccess = m.SecureInternetAccess.flattenEndpointConnectionProfileSecureInternetAccess(ctx, v, &diags)
	}

	if v, ok := o["preferredDTLSTunnel"]; ok {
		m.PreferredDtlsTunnel = parseStringValue(v)
	}

	if v, ok := o["useGuiSamlAuth"]; ok {
		m.UseGuiSamlAuth = parseStringValue(v)
	}

	if v, ok := o["useWebview2SamlAuth"]; ok {
		m.UseWebview2SamlAuth = parseStringValue(v)
	}

	if v, ok := o["beforeLogonSamlAuth"]; ok {
		m.BeforeLogonSamlAuth = parseStringValue(v)
	}

	if v, ok := o["afterLogonSamlAuth"]; ok {
		m.AfterLogonSamlAuth = parseStringValue(v)
	}

	if v, ok := o["allowPersonalVpns"]; ok {
		m.AllowPersonalVpns = parseBoolValue(v)
	}

	if v, ok := o["mtuSize"]; ok {
		m.MtuSize = parseFloat64Value(v)
	}

	if v, ok := o["vpnType"]; ok {
		m.VpnType = parseStringValue(v)
	}

	if v, ok := o["disableInternetCheck"]; ok {
		m.DisableInternetCheck = parseStringValue(v)
	}

	if v, ok := o["showDisconnectBtn"]; ok {
		m.ShowDisconnectBtn = parseStringValue(v)
	}

	if v, ok := o["enableInvalidServerCertWarning"]; ok {
		m.EnableInvalidServerCertWarning = parseStringValue(v)
	}

	if v, ok := o["preLogon"]; ok {
		m.PreLogon = m.PreLogon.flattenEndpointConnectionProfilePreLogon(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointConnectionProfileModel) getURLObjectEndpointConnectionProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointConnectionProfileAvailableVpnsModel struct {
	Type                     types.String                                                                   `tfsdk:"type"`
	Name                     types.String                                                                   `tfsdk:"name"`
	RemoteGateway            types.String                                                                   `tfsdk:"remote_gateway"`
	UsernamePrompt           types.String                                                                   `tfsdk:"username_prompt"`
	SaveUsername             types.String                                                                   `tfsdk:"save_username"`
	ShowAlwaysUp             types.String                                                                   `tfsdk:"show_always_up"`
	ShowAutoConnect          types.String                                                                   `tfsdk:"show_auto_connect"`
	ShowRememberPassword     types.String                                                                   `tfsdk:"show_remember_password"`
	AuthenticateWithSso      types.String                                                                   `tfsdk:"authenticate_with_sso"`
	AllowFidoAuth            types.String                                                                   `tfsdk:"allow_fido_auth"`
	EnableLocalLan           types.String                                                                   `tfsdk:"enable_local_lan"`
	EncapsulationMode        types.String                                                                   `tfsdk:"encapsulation_mode"`
	UdpPort                  types.Float64                                                                  `tfsdk:"udp_port"`
	TcpPort                  types.Float64                                                                  `tfsdk:"tcp_port"`
	ConnectDisconnectScripts *datasourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel `tfsdk:"connect_disconnect_scripts"`
	Port                     types.Float64                                                                  `tfsdk:"port"`
	RequireCertificate       types.String                                                                   `tfsdk:"require_certificate"`
	ExternalBrowserSamlLogin types.String                                                                   `tfsdk:"external_browser_saml_login"`
	AuthMethod               types.String                                                                   `tfsdk:"auth_method"`
	DnsSuffixes              types.Set                                                                      `tfsdk:"dns_suffixes"`
	ShowPasscode             types.String                                                                   `tfsdk:"show_passcode"`
	PostureCheck             *datasourceEndpointConnectionProfileAvailableVpnsPostureCheckModel             `tfsdk:"posture_check"`
	EapEnabled               types.Bool                                                                     `tfsdk:"eap_enabled"`
	SamlPort                 types.Float64                                                                  `tfsdk:"saml_port"`
	PreSharedKey             types.String                                                                   `tfsdk:"pre_shared_key"`
}

type datasourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel struct {
	OnConnectWindows    types.String `tfsdk:"on_connect_windows"`
	OnConnectMac        types.String `tfsdk:"on_connect_mac"`
	OnDisconnectWindows types.String `tfsdk:"on_disconnect_windows"`
	OnDisconnectMac     types.String `tfsdk:"on_disconnect_mac"`
}

type datasourceEndpointConnectionProfileAvailableVpnsPostureCheckModel struct {
	Tag                types.String `tfsdk:"tag"`
	Action             types.String `tfsdk:"action"`
	CheckFailedMessage types.String `tfsdk:"check_failed_message"`
}

type datasourceEndpointConnectionProfileLockdownModel struct {
	Status              types.String                                                         `tfsdk:"status"`
	GracePeriod         types.Float64                                                        `tfsdk:"grace_period"`
	MaxAttempts         types.Float64                                                        `tfsdk:"max_attempts"`
	Ips                 []datasourceEndpointConnectionProfileLockdownIpsModel                `tfsdk:"ips"`
	Domains             []datasourceEndpointConnectionProfileLockdownDomainsModel            `tfsdk:"domains"`
	DetectCaptivePortal *datasourceEndpointConnectionProfileLockdownDetectCaptivePortalModel `tfsdk:"detect_captive_portal"`
}

type datasourceEndpointConnectionProfileLockdownIpsModel struct {
	Ip       types.String `tfsdk:"ip"`
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
}

type datasourceEndpointConnectionProfileLockdownDomainsModel struct {
	Address types.String `tfsdk:"address"`
}

type datasourceEndpointConnectionProfileLockdownDetectCaptivePortalModel struct {
	Status                      types.String `tfsdk:"status"`
	DisableWindowsCaptivePortal types.String `tfsdk:"disable_windows_captive_portal"`
}

type datasourceEndpointConnectionProfileOnFabricRuleSetModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointConnectionProfileOffNetSplitTunnelModel struct {
	LocalApps       types.Set                                                          `tfsdk:"local_apps"`
	Isdbs           []datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel   `tfsdk:"isdbs"`
	Fqdns           types.Set                                                          `tfsdk:"fqdns"`
	Subnets         []datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel `tfsdk:"subnets"`
	SubnetsIpsec    types.Set                                                          `tfsdk:"subnets_ipsec"`
	SplitTunnelMode types.String                                                       `tfsdk:"split_tunnel_mode"`
}

type datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointConnectionProfileSplitTunnelModel struct {
	LocalApps       types.Set                                                    `tfsdk:"local_apps"`
	Isdbs           []datasourceEndpointConnectionProfileSplitTunnelIsdbsModel   `tfsdk:"isdbs"`
	Fqdns           types.Set                                                    `tfsdk:"fqdns"`
	Subnets         []datasourceEndpointConnectionProfileSplitTunnelSubnetsModel `tfsdk:"subnets"`
	SubnetsIpsec    types.Set                                                    `tfsdk:"subnets_ipsec"`
	SplitTunnelMode types.String                                                 `tfsdk:"split_tunnel_mode"`
}

type datasourceEndpointConnectionProfileSplitTunnelIsdbsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointConnectionProfileSplitTunnelSubnetsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointConnectionProfileSecureInternetAccessModel struct {
	AuthenticateWithSso      types.String                                                                          `tfsdk:"authenticate_with_sso"`
	AllowFidoAuth            types.String                                                                          `tfsdk:"allow_fido_auth"`
	ConnectDisconnectScripts *datasourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel `tfsdk:"connect_disconnect_scripts"`
	DnsSuffixes              types.Set                                                                             `tfsdk:"dns_suffixes"`
	EnableLocalLan           types.String                                                                          `tfsdk:"enable_local_lan"`
	FailoverSequence         types.Set                                                                             `tfsdk:"failover_sequence"`
	PostureCheck             *datasourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel             `tfsdk:"posture_check"`
	EapEnabled               types.Bool                                                                            `tfsdk:"eap_enabled"`
	EncapsulationMode        types.String                                                                          `tfsdk:"encapsulation_mode"`
	ExternalBrowserSamlLogin types.String                                                                          `tfsdk:"external_browser_saml_login"`
}

type datasourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel struct {
	OnConnectWindows    types.String `tfsdk:"on_connect_windows"`
	OnConnectMac        types.String `tfsdk:"on_connect_mac"`
	OnDisconnectWindows types.String `tfsdk:"on_disconnect_windows"`
	OnDisconnectMac     types.String `tfsdk:"on_disconnect_mac"`
}

type datasourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel struct {
	Tag                types.String `tfsdk:"tag"`
	Action             types.String `tfsdk:"action"`
	CheckFailedMessage types.String `tfsdk:"check_failed_message"`
}

type datasourceEndpointConnectionProfilePreLogonModel struct {
	VpnType       types.String                                                `tfsdk:"vpn_type"`
	RemoteGateway types.String                                                `tfsdk:"remote_gateway"`
	CommonName    *datasourceEndpointConnectionProfilePreLogonCommonNameModel `tfsdk:"common_name"`
	Issuer        *datasourceEndpointConnectionProfilePreLogonIssuerModel     `tfsdk:"issuer"`
	Port          types.Float64                                               `tfsdk:"port"`
}

type datasourceEndpointConnectionProfilePreLogonCommonNameModel struct {
	MatchType types.String `tfsdk:"match_type"`
	Pattern   types.String `tfsdk:"pattern"`
}

type datasourceEndpointConnectionProfilePreLogonIssuerModel struct {
	MatchType types.String `tfsdk:"match_type"`
	Pattern   types.String `tfsdk:"pattern"`
}

func (m *datasourceEndpointConnectionProfileAvailableVpnsModel) flattenEndpointConnectionProfileAvailableVpns(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileAvailableVpnsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileAvailableVpnsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileAvailableVpnsModel{}
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

	if v, ok := o["allowFidoAuth"]; ok {
		m.AllowFidoAuth = parseStringValue(v)
	}

	if v, ok := o["enableLocalLan"]; ok {
		m.EnableLocalLan = parseStringValue(v)
	}

	if v, ok := o["encapsulationMode"]; ok {
		m.EncapsulationMode = parseStringValue(v)
	}

	if v, ok := o["udpPort"]; ok {
		m.UdpPort = parseFloat64Value(v)
	}

	if v, ok := o["tcpPort"]; ok {
		m.TcpPort = parseFloat64Value(v)
	}

	if v, ok := o["connectDisconnectScripts"]; ok {
		m.ConnectDisconnectScripts = m.ConnectDisconnectScripts.flattenEndpointConnectionProfileAvailableVpnsConnectDisconnectScripts(ctx, v, diags)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	if v, ok := o["requireCertificate"]; ok {
		m.RequireCertificate = parseStringValue(v)
	}

	if v, ok := o["externalBrowserSamlLogin"]; ok {
		m.ExternalBrowserSamlLogin = parseStringValue(v)
	}

	if v, ok := o["authMethod"]; ok {
		m.AuthMethod = parseStringValue(v)
	}

	if v, ok := o["dnsSuffixes"]; ok {
		m.DnsSuffixes = parseSetValue(ctx, v, types.StringType)
	} else {
		m.DnsSuffixes = types.SetNull(types.StringType)
	}

	if v, ok := o["showPasscode"]; ok {
		m.ShowPasscode = parseStringValue(v)
	}

	if v, ok := o["postureCheck"]; ok {
		m.PostureCheck = m.PostureCheck.flattenEndpointConnectionProfileAvailableVpnsPostureCheck(ctx, v, diags)
	}

	if v, ok := o["eapEnabled"]; ok {
		m.EapEnabled = parseBoolValue(v)
	}

	if v, ok := o["samlPort"]; ok {
		m.SamlPort = parseFloat64Value(v)
	}

	if v, ok := o["preSharedKey"]; ok {
		m.PreSharedKey = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointConnectionProfileModel) flattenEndpointConnectionProfileAvailableVpnsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfileAvailableVpnsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfileAvailableVpnsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument available_vpns is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfileAvailableVpnsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfileAvailableVpnsModel{}
	}

	values := make([]datasourceEndpointConnectionProfileAvailableVpnsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfileAvailableVpnsModel
		if i < len(s.AvailableVpns) {
			m = s.AvailableVpns[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileAvailableVpns(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel) flattenEndpointConnectionProfileAvailableVpnsConnectDisconnectScripts(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["onConnectWindows"]; ok {
		m.OnConnectWindows = parseStringValue(v)
	}

	if v, ok := o["onConnectMac"]; ok {
		m.OnConnectMac = parseStringValue(v)
	}

	if v, ok := o["onDisconnectWindows"]; ok {
		m.OnDisconnectWindows = parseStringValue(v)
	}

	if v, ok := o["onDisconnectMac"]; ok {
		m.OnDisconnectMac = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfileAvailableVpnsPostureCheckModel) flattenEndpointConnectionProfileAvailableVpnsPostureCheck(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileAvailableVpnsPostureCheckModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileAvailableVpnsPostureCheckModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileAvailableVpnsPostureCheckModel{}
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

func (m *datasourceEndpointConnectionProfileLockdownModel) flattenEndpointConnectionProfileLockdown(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileLockdownModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileLockdownModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileLockdownModel{}
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
		m.Ips = m.flattenEndpointConnectionProfileLockdownIpsList(ctx, v, diags)
	}

	if v, ok := o["domains"]; ok {
		m.Domains = m.flattenEndpointConnectionProfileLockdownDomainsList(ctx, v, diags)
	}

	if v, ok := o["detectCaptivePortal"]; ok {
		m.DetectCaptivePortal = m.DetectCaptivePortal.flattenEndpointConnectionProfileLockdownDetectCaptivePortal(ctx, v, diags)
	}

	return m
}

func (m *datasourceEndpointConnectionProfileLockdownIpsModel) flattenEndpointConnectionProfileLockdownIps(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileLockdownIpsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileLockdownIpsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileLockdownIpsModel{}
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

func (s *datasourceEndpointConnectionProfileLockdownModel) flattenEndpointConnectionProfileLockdownIpsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfileLockdownIpsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfileLockdownIpsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument ips is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfileLockdownIpsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfileLockdownIpsModel{}
	}

	values := make([]datasourceEndpointConnectionProfileLockdownIpsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfileLockdownIpsModel
		if i < len(s.Ips) {
			m = s.Ips[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileLockdownIps(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfileLockdownDomainsModel) flattenEndpointConnectionProfileLockdownDomains(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileLockdownDomainsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileLockdownDomainsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileLockdownDomainsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["address"]; ok {
		m.Address = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointConnectionProfileLockdownModel) flattenEndpointConnectionProfileLockdownDomainsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfileLockdownDomainsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfileLockdownDomainsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument domains is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfileLockdownDomainsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfileLockdownDomainsModel{}
	}

	values := make([]datasourceEndpointConnectionProfileLockdownDomainsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfileLockdownDomainsModel
		if i < len(s.Domains) {
			m = s.Domains[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileLockdownDomains(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfileLockdownDetectCaptivePortalModel) flattenEndpointConnectionProfileLockdownDetectCaptivePortal(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileLockdownDetectCaptivePortalModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileLockdownDetectCaptivePortalModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileLockdownDetectCaptivePortalModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["disableWindowsCaptivePortal"]; ok {
		m.DisableWindowsCaptivePortal = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfileOnFabricRuleSetModel) flattenEndpointConnectionProfileOnFabricRuleSet(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileOnFabricRuleSetModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileOnFabricRuleSetModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileOnFabricRuleSetModel{}
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

func (m *datasourceEndpointConnectionProfileOffNetSplitTunnelModel) flattenEndpointConnectionProfileOffNetSplitTunnel(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileOffNetSplitTunnelModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileOffNetSplitTunnelModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileOffNetSplitTunnelModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["localApps"]; ok {
		m.LocalApps = parseSetValue(ctx, v, types.StringType)
	} else {
		m.LocalApps = types.SetNull(types.StringType)
	}

	if v, ok := o["isdbs"]; ok {
		m.Isdbs = m.flattenEndpointConnectionProfileOffNetSplitTunnelIsdbsList(ctx, v, diags)
	}

	if v, ok := o["fqdns"]; ok {
		m.Fqdns = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Fqdns = types.SetNull(types.StringType)
	}

	if v, ok := o["subnets"]; ok {
		m.Subnets = m.flattenEndpointConnectionProfileOffNetSplitTunnelSubnetsList(ctx, v, diags)
	}

	if v, ok := o["subnetsIpsec"]; ok {
		m.SubnetsIpsec = parseSetValue(ctx, v, types.StringType)
	} else {
		m.SubnetsIpsec = types.SetNull(types.StringType)
	}

	if v, ok := o["splitTunnelMode"]; ok {
		m.SplitTunnelMode = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel) flattenEndpointConnectionProfileOffNetSplitTunnelIsdbs(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel{}
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

func (s *datasourceEndpointConnectionProfileOffNetSplitTunnelModel) flattenEndpointConnectionProfileOffNetSplitTunnelIsdbsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument isdbs is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel{}
	}

	values := make([]datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel
		if i < len(s.Isdbs) {
			m = s.Isdbs[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileOffNetSplitTunnelIsdbs(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel) flattenEndpointConnectionProfileOffNetSplitTunnelSubnets(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel{}
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

func (s *datasourceEndpointConnectionProfileOffNetSplitTunnelModel) flattenEndpointConnectionProfileOffNetSplitTunnelSubnetsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument subnets is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel{}
	}

	values := make([]datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel
		if i < len(s.Subnets) {
			m = s.Subnets[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileOffNetSplitTunnelSubnets(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfileSplitTunnelModel) flattenEndpointConnectionProfileSplitTunnel(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileSplitTunnelModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileSplitTunnelModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileSplitTunnelModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["localApps"]; ok {
		m.LocalApps = parseSetValue(ctx, v, types.StringType)
	} else {
		m.LocalApps = types.SetNull(types.StringType)
	}

	if v, ok := o["isdbs"]; ok {
		m.Isdbs = m.flattenEndpointConnectionProfileSplitTunnelIsdbsList(ctx, v, diags)
	}

	if v, ok := o["fqdns"]; ok {
		m.Fqdns = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Fqdns = types.SetNull(types.StringType)
	}

	if v, ok := o["subnets"]; ok {
		m.Subnets = m.flattenEndpointConnectionProfileSplitTunnelSubnetsList(ctx, v, diags)
	}

	if v, ok := o["subnetsIpsec"]; ok {
		m.SubnetsIpsec = parseSetValue(ctx, v, types.StringType)
	} else {
		m.SubnetsIpsec = types.SetNull(types.StringType)
	}

	if v, ok := o["splitTunnelMode"]; ok {
		m.SplitTunnelMode = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfileSplitTunnelIsdbsModel) flattenEndpointConnectionProfileSplitTunnelIsdbs(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileSplitTunnelIsdbsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileSplitTunnelIsdbsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileSplitTunnelIsdbsModel{}
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

func (s *datasourceEndpointConnectionProfileSplitTunnelModel) flattenEndpointConnectionProfileSplitTunnelIsdbsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfileSplitTunnelIsdbsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfileSplitTunnelIsdbsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument isdbs is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfileSplitTunnelIsdbsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfileSplitTunnelIsdbsModel{}
	}

	values := make([]datasourceEndpointConnectionProfileSplitTunnelIsdbsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfileSplitTunnelIsdbsModel
		if i < len(s.Isdbs) {
			m = s.Isdbs[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileSplitTunnelIsdbs(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfileSplitTunnelSubnetsModel) flattenEndpointConnectionProfileSplitTunnelSubnets(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileSplitTunnelSubnetsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileSplitTunnelSubnetsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileSplitTunnelSubnetsModel{}
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

func (s *datasourceEndpointConnectionProfileSplitTunnelModel) flattenEndpointConnectionProfileSplitTunnelSubnetsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointConnectionProfileSplitTunnelSubnetsModel {
	if o == nil {
		return []datasourceEndpointConnectionProfileSplitTunnelSubnetsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument subnets is not type of []interface{}.", "")
		return []datasourceEndpointConnectionProfileSplitTunnelSubnetsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointConnectionProfileSplitTunnelSubnetsModel{}
	}

	values := make([]datasourceEndpointConnectionProfileSplitTunnelSubnetsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointConnectionProfileSplitTunnelSubnetsModel
		if i < len(s.Subnets) {
			m = s.Subnets[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileSplitTunnelSubnets(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointConnectionProfileSecureInternetAccessModel) flattenEndpointConnectionProfileSecureInternetAccess(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileSecureInternetAccessModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileSecureInternetAccessModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileSecureInternetAccessModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["authenticateWithSSO"]; ok {
		m.AuthenticateWithSso = parseStringValue(v)
	}

	if v, ok := o["allowFidoAuth"]; ok {
		m.AllowFidoAuth = parseStringValue(v)
	}

	if v, ok := o["connectDisconnectScripts"]; ok {
		m.ConnectDisconnectScripts = m.ConnectDisconnectScripts.flattenEndpointConnectionProfileSecureInternetAccessConnectDisconnectScripts(ctx, v, diags)
	}

	if v, ok := o["dnsSuffixes"]; ok {
		m.DnsSuffixes = parseSetValue(ctx, v, types.StringType)
	} else {
		m.DnsSuffixes = types.SetNull(types.StringType)
	}

	if v, ok := o["enableLocalLan"]; ok {
		m.EnableLocalLan = parseStringValue(v)
	}

	if v, ok := o["failoverSequence"]; ok {
		m.FailoverSequence = parseSetValue(ctx, v, types.StringType)
	} else {
		m.FailoverSequence = types.SetNull(types.StringType)
	}

	if v, ok := o["postureCheck"]; ok {
		m.PostureCheck = m.PostureCheck.flattenEndpointConnectionProfileSecureInternetAccessPostureCheck(ctx, v, diags)
	}

	if v, ok := o["eapEnabled"]; ok {
		m.EapEnabled = parseBoolValue(v)
	}

	if v, ok := o["encapsulationMode"]; ok {
		m.EncapsulationMode = parseStringValue(v)
	}

	if v, ok := o["externalBrowserSamlLogin"]; ok {
		m.ExternalBrowserSamlLogin = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel) flattenEndpointConnectionProfileSecureInternetAccessConnectDisconnectScripts(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["onConnectWindows"]; ok {
		m.OnConnectWindows = parseStringValue(v)
	}

	if v, ok := o["onConnectMac"]; ok {
		m.OnConnectMac = parseStringValue(v)
	}

	if v, ok := o["onDisconnectWindows"]; ok {
		m.OnDisconnectWindows = parseStringValue(v)
	}

	if v, ok := o["onDisconnectMac"]; ok {
		m.OnDisconnectMac = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel) flattenEndpointConnectionProfileSecureInternetAccessPostureCheck(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel {
	if input == nil {
		return &datasourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel{}
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

func (m *datasourceEndpointConnectionProfilePreLogonModel) flattenEndpointConnectionProfilePreLogon(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilePreLogonModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilePreLogonModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilePreLogonModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["vpnType"]; ok {
		m.VpnType = parseStringValue(v)
	}

	if v, ok := o["remoteGateway"]; ok {
		m.RemoteGateway = parseStringValue(v)
	}

	if v, ok := o["commonName"]; ok {
		m.CommonName = m.CommonName.flattenEndpointConnectionProfilePreLogonCommonName(ctx, v, diags)
	}

	if v, ok := o["issuer"]; ok {
		m.Issuer = m.Issuer.flattenEndpointConnectionProfilePreLogonIssuer(ctx, v, diags)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceEndpointConnectionProfilePreLogonCommonNameModel) flattenEndpointConnectionProfilePreLogonCommonName(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilePreLogonCommonNameModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilePreLogonCommonNameModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilePreLogonCommonNameModel{}
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

func (m *datasourceEndpointConnectionProfilePreLogonIssuerModel) flattenEndpointConnectionProfilePreLogonIssuer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointConnectionProfilePreLogonIssuerModel {
	if input == nil {
		return &datasourceEndpointConnectionProfilePreLogonIssuerModel{}
	}
	if m == nil {
		m = &datasourceEndpointConnectionProfilePreLogonIssuerModel{}
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
