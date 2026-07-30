// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceEndpointConnectionProfile{}
var _ resource.ResourceWithMoveState = &resourceEndpointConnectionProfile{}

func newResourceEndpointConnectionProfile() resource.Resource {
	return &resourceEndpointConnectionProfile{}
}

type resourceEndpointConnectionProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointConnectionProfileModel describes the resource data model.
type resourceEndpointConnectionProfileModel struct {
	ID                             types.String                                                `tfsdk:"id"`
	ConnectToFortisase             types.String                                                `tfsdk:"connect_to_fortisase"`
	ConnectToFortiSase             types.String                                                `tfsdk:"connect_to_forti_sase"`
	AvailableVpns                  []resourceEndpointConnectionProfileAvailableVpnsModel       `tfsdk:"available_vpns"`
	AvailableVpNs                  []resourceEndpointConnectionProfileAvailableVpnsModel       `tfsdk:"available_vp_ns"`
	Lockdown                       *resourceEndpointConnectionProfileLockdownModel             `tfsdk:"lockdown"`
	OnFabricRuleSet                *resourceEndpointConnectionProfileOnFabricRuleSetModel      `tfsdk:"on_fabric_rule_set"`
	OffNetSplitTunnel              *resourceEndpointConnectionProfileOffNetSplitTunnelModel    `tfsdk:"off_net_split_tunnel"`
	SplitTunnel                    *resourceEndpointConnectionProfileSplitTunnelModel          `tfsdk:"split_tunnel"`
	AllowInvalidServerCertificate  types.String                                                `tfsdk:"allow_invalid_server_certificate"`
	EndpointOnNetBypass            types.Bool                                                  `tfsdk:"endpoint_on_net_bypass"`
	AuthBeforeUserLogon            types.Bool                                                  `tfsdk:"auth_before_user_logon"`
	SecureInternetAccess           *resourceEndpointConnectionProfileSecureInternetAccessModel `tfsdk:"secure_internet_access"`
	PreferredDtlsTunnel            types.String                                                `tfsdk:"preferred_dtls_tunnel"`
	UseGuiSamlAuth                 types.String                                                `tfsdk:"use_gui_saml_auth"`
	UseWebview2SamlAuth            types.String                                                `tfsdk:"use_webview2_saml_auth"`
	BeforeLogonSamlAuth            types.String                                                `tfsdk:"before_logon_saml_auth"`
	AfterLogonSamlAuth             types.String                                                `tfsdk:"after_logon_saml_auth"`
	AllowPersonalVpns              types.Bool                                                  `tfsdk:"allow_personal_vpns"`
	MtuSize                        types.Float64                                               `tfsdk:"mtu_size"`
	VpnType                        types.String                                                `tfsdk:"vpn_type"`
	DisableInternetCheck           types.String                                                `tfsdk:"disable_internet_check"`
	ShowDisconnectBtn              types.String                                                `tfsdk:"show_disconnect_btn"`
	EnableInvalidServerCertWarning types.String                                                `tfsdk:"enable_invalid_server_cert_warning"`
	PreLogon                       *resourceEndpointConnectionProfilePreLogonModel             `tfsdk:"pre_logon"`
	PrimaryKey                     types.String                                                `tfsdk:"primary_key"`
}

func (r *resourceEndpointConnectionProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_connection_profile"
}

func (r *resourceEndpointConnectionProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Connection Profile Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"connect_to_fortisase": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("automatically", "manually"),
				},
				Computed: true,
				Optional: true,
			},
			"connect_to_forti_sase": schema.StringAttribute{
				DeprecationMessage: "\"connect_to_forti_sase\" is deprecated; use \"connect_to_fortisase\" instead.",
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("automatically", "manually"),
				},
				Computed: true,
				Optional: true,
			},
			"allow_invalid_server_certificate": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
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
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"use_gui_saml_auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"use_webview2_saml_auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"before_logon_saml_auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("webBrowser", "electron"),
				},
				MarkdownDescription: "Specifies the browser framework used for Pre-logon VPN SAML authentication.\nSupported values: webBrowser, electron.",
				Computed:            true,
				Optional:            true,
			},
			"after_logon_saml_auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("webBrowser", "electron", "webView2"),
				},
				MarkdownDescription: "Specifies the browser framework used for normal VPN SAML authentication.\nSupported values: webBrowser, electron, webView2.",
				Computed:            true,
				Optional:            true,
			},
			"allow_personal_vpns": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"mtu_size": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.Between(576, 1500),
				},
				Computed: true,
				Optional: true,
			},
			"vpn_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("sslVPN", "ipSecVPN"),
				},
				Computed: true,
				Optional: true,
			},
			"disable_internet_check": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"show_disconnect_btn": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"enable_invalid_server_cert_warning": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"available_vpns": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("sslVPN", "ipSecVPN"),
							},
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"remote_gateway": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"username_prompt": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"save_username": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_always_up": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_auto_connect": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_remember_password": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"authenticate_with_sso": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"allow_fido_auth": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"enable_local_lan": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"encapsulation_mode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("Auto", "TCP", "UDP"),
							},
							Computed: true,
							Optional: true,
						},
						"udp_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(500, 65535),
							},
							Computed: true,
							Optional: true,
						},
						"tcp_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(1, 65535),
							},
							Computed: true,
							Optional: true,
						},
						"port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtMost(65535),
							},
							Computed: true,
							Optional: true,
						},
						"require_certificate": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"external_browser_saml_login": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"auth_method": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("preSharedKey", "smartCardCert", "systemStoreCert"),
							},
							Computed: true,
							Optional: true,
						},
						"dns_suffixes": schema.SetAttribute{
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
						"show_passcode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"eap_enabled": schema.BoolAttribute{
							MarkdownDescription: "Per-tunnel EAP for this manual IPsec VPN entry in availableVPNs.",
							Computed:            true,
							Optional:            true,
						},
						"saml_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtMost(65535),
							},
							Computed: true,
							Optional: true,
						},
						"pre_shared_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"connect_disconnect_scripts": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"on_connect_windows": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
									Optional: true,
								},
								"on_connect_mac": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
									Optional: true,
								},
								"on_disconnect_windows": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
									Optional: true,
								},
								"on_disconnect_mac": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
									Optional: true,
								},
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
										stringvalidatorwarning.OneOf("allow", "prohibit"),
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
			"available_vp_ns": schema.ListNestedAttribute{
				DeprecationMessage: "\"available_vp_ns\" is deprecated; use \"available_vpns\" instead.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("sslVPN", "ipSecVPN"),
							},
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"remote_gateway": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"username_prompt": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"save_username": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_always_up": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_auto_connect": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_remember_password": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"authenticate_with_sso": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"allow_fido_auth": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"enable_local_lan": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"encapsulation_mode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("Auto", "TCP", "UDP"),
							},
							Computed: true,
							Optional: true,
						},
						"udp_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(500, 65535),
							},
							Computed: true,
							Optional: true,
						},
						"tcp_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(1, 65535),
							},
							Computed: true,
							Optional: true,
						},
						"port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtMost(65535),
							},
							Computed: true,
							Optional: true,
						},
						"require_certificate": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"external_browser_saml_login": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"auth_method": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("preSharedKey", "smartCardCert", "systemStoreCert"),
							},
							Computed: true,
							Optional: true,
						},
						"dns_suffixes": schema.SetAttribute{
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
						"show_passcode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"eap_enabled": schema.BoolAttribute{
							MarkdownDescription: "Per-tunnel EAP for this manual IPsec VPN entry in availableVPNs.",
							Computed:            true,
							Optional:            true,
						},
						"saml_port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtMost(65535),
							},
							Computed: true,
							Optional: true,
						},
						"pre_shared_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"connect_disconnect_scripts": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"on_connect_windows": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
									Optional: true,
								},
								"on_connect_mac": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
									Optional: true,
								},
								"on_disconnect_windows": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
									Optional: true,
								},
								"on_disconnect_mac": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthAtMost(1023),
									},
									Computed: true,
									Optional: true,
								},
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
										stringvalidatorwarning.OneOf("allow", "prohibit"),
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
			"lockdown": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
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
							float64validatorwarning.AtLeast(1),
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
										stringvalidatorwarning.OneOf("tcp", "udp", "icmp", ""),
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
									stringvalidatorwarning.OneOf("enable", "disable"),
								},
								Computed: true,
								Optional: true,
							},
							"disable_windows_captive_portal": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.OneOf("enable", "disable"),
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
						Required: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("endpoint/on-net-rules"),
						},
						Required: true,
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
					"split_tunnel_mode": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("exclude", "include"),
						},
						Computed: true,
						Optional: true,
					},
					"isdbs": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("network/basic-internet-services"),
									},
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
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("network/hosts", "network/host-groups"),
									},
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
					"split_tunnel_mode": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("exclude", "include"),
						},
						Computed: true,
						Optional: true,
					},
					"isdbs": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("network/basic-internet-services"),
									},
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
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("network/hosts", "network/host-groups"),
									},
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
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"allow_fido_auth": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"dns_suffixes": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"enable_local_lan": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"failover_sequence": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"eap_enabled": schema.BoolAttribute{
						MarkdownDescription: "When vpnType is ipSecVPN, sets EAP (eap_method) on the Secure Internet Access tunnel(s) only (SIA-named connections), for both on-net and off-net EMS profiles. Custom/manual IPsec tunnels use availableVPNs[].eapEnabled.",
						Computed:            true,
						Optional:            true,
					},
					"encapsulation_mode": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("Auto", "TCP", "UDP"),
						},
						Computed: true,
						Optional: true,
					},
					"external_browser_saml_login": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"connect_disconnect_scripts": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"on_connect_windows": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtMost(1023),
								},
								Computed: true,
								Optional: true,
							},
							"on_connect_mac": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtMost(1023),
								},
								Computed: true,
								Optional: true,
							},
							"on_disconnect_windows": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtMost(1023),
								},
								Computed: true,
								Optional: true,
							},
							"on_disconnect_mac": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtMost(1023),
								},
								Computed: true,
								Optional: true,
							},
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
									stringvalidatorwarning.OneOf("allow", "prohibit"),
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
			"pre_logon": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"vpn_type": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("sslVPN", "ipSecVPN"),
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
							float64validatorwarning.AtMost(65535),
						},
						Computed: true,
						Optional: true,
					},
					"common_name": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"match_type": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.OneOf("wildcard", "regex"),
								},
								Computed: true,
								Optional: true,
							},
							"pattern": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtLeast(1),
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
									stringvalidatorwarning.OneOf("wildcard", "regex"),
								},
								Computed: true,
								Optional: true,
							},
							"pattern": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtLeast(1),
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

func (r *resourceEndpointConnectionProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *resourceEndpointConnectionProfile) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_endpoint_connection_profiles" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceEndpointConnectionProfileModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceEndpointConnectionProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointConnectionProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	for i := 0; i < 4; i++ {
		c := r.fortiClient.Client
		var input_model forticlient.InputModel
		input_model.Mkey = data.PrimaryKey.ValueString()
		input_model.BodyParams = *(data.getCreateObjectEndpointConnectionProfile(ctx, diags))
		input_model.URLParams = *(data.getURLObjectEndpointConnectionProfile(ctx, "create", diags))

		if diags.HasError() {
			return
		}
		output, err := c.UpdateEndpointConnectionProfiles(&input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
				getErrorDetail(&input_model, output),
			)
			return
		}

		mkey := fmt.Sprintf("%v", output["primaryKey"])
		data.ID = types.StringValue(mkey)
		var read_input_model forticlient.InputModel
		read_input_model.Mkey = mkey
		read_input_model.URLParams = *(data.getURLObjectEndpointConnectionProfile(ctx, "read", diags))

		read_output, err := c.ReadEndpointConnectionProfiles(&read_input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
				getErrorDetail(&read_input_model, read_output),
			)
			return
		}

		if read_output["connectToFortiSASE"] != data.ConnectToFortiSase.ValueString() {
			// diags.AddWarning(
			// 	"Detected that connectToFortiSASE was not accepted by the server. Resending the request...",
			// 	"This issue is handled automatically by Terraform. No user action is required.",
			// )
			continue
		}

		if data.SecureInternetAccess != nil && data.SecureInternetAccess.PostureCheck != nil {
			if secureInternetAccess, ok := read_output["secureInternetAccess"].(map[string]interface{}); ok {
				if postureCheck, ok := secureInternetAccess["postureCheck"].(map[string]interface{}); ok {
					if fmt.Sprintf("%v", postureCheck["action"]) != data.SecureInternetAccess.PostureCheck.Action.ValueString() ||
						fmt.Sprintf("%v", postureCheck["tag"]) != data.SecureInternetAccess.PostureCheck.Tag.ValueString() ||
						fmt.Sprintf("%v", postureCheck["checkFailedMessage"]) != data.SecureInternetAccess.PostureCheck.CheckFailedMessage.ValueString() {
						// diags.AddWarning(
						// 	"Detected that secureInternetAccess was not accepted by the server. Resending the request...",
						// 	"This issue is handled automatically by Terraform. No user action is required.",
						// )
						continue
					}
				}
			}
		}

		diags.Append(data.refreshEndpointConnectionProfile(ctx, read_output)...)
		if diags.HasError() {
			return
		}
		break
	}
	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointConnectionProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointConnectionProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointConnectionProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointConnectionProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointConnectionProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointConnectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointConnectionProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointConnectionProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointConnectionProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointConnectionProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointConnectionProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}
	result_model := make(map[string]interface{})
	result_model["onFabricRuleSet"] = nil
	secureInternetAccess := make(map[string]interface{})
	postureTag := make(map[string]interface{})
	postureTag["tag"] = ""
	postureTag["action"] = "allow"
	postureTag["checkFailedMessage"] = ""
	secureInternetAccess["postureCheck"] = postureTag
	result_model["secureInternetAccess"] = secureInternetAccess

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = result_model
	input_model.URLParams = *(state.getURLObjectEndpointConnectionProfile(ctx, "delete", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointConnectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceEndpointConnectionProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointConnectionProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointConnectionProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointConnectionProfiles(&input_model)
	if err != nil {
		if isNotFoundResponse(read_output) {
			resp.State.RemoveResource(ctx)
			return
		}
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
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

func (r *resourceEndpointConnectionProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceEndpointConnectionProfileModel) refreshEndpointConnectionProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceEndpointConnectionProfileModel) getCreateObjectEndpointConnectionProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	connectToFortisaseValue := data.ConnectToFortisase
	connectToFortisaseValueSource := "connect_to_fortisase"
	if connectToFortisaseValue.IsNull() || connectToFortisaseValue.IsUnknown() {
		if !data.ConnectToFortiSase.IsNull() && !data.ConnectToFortiSase.IsUnknown() {
			connectToFortisaseValue = data.ConnectToFortiSase
			connectToFortisaseValueSource = "connect_to_forti_sase"
		}
	}
	if !data.ConnectToFortiSase.IsNull() && !data.ConnectToFortiSase.IsUnknown() && connectToFortisaseValueSource != "connect_to_forti_sase" && !connectToFortisaseValue.Equal(data.ConnectToFortiSase) {
		diags.AddError("Conflicting Terraform arguments",
			fmt.Sprintf("Arguments %q and %q cannot both be configured with different values.", connectToFortisaseValueSource, "connect_to_forti_sase"),
		)
	}
	if !connectToFortisaseValue.IsNull() && !connectToFortisaseValue.IsUnknown() {
		data.ConnectToFortisase = connectToFortisaseValue
		data.ConnectToFortiSase = connectToFortisaseValue
	}
	if !connectToFortisaseValue.IsNull() && !connectToFortisaseValue.IsUnknown() {
		result["connectToFortiSASE"] = connectToFortisaseValue.ValueString()
	}

	availableVpnsValue := data.AvailableVpns
	availableVpnsValueSource := "available_vpns"
	if availableVpnsValue == nil && data.AvailableVpNs != nil {
		availableVpnsValue = data.AvailableVpNs
		availableVpnsValueSource = "available_vp_ns"
	}
	if data.AvailableVpNs != nil && availableVpnsValueSource != "available_vp_ns" && !reflect.DeepEqual(availableVpnsValue, data.AvailableVpNs) {
		diags.AddError("Conflicting Terraform arguments",
			fmt.Sprintf("Arguments %q and %q cannot both be configured with different values.", availableVpnsValueSource, "available_vp_ns"),
		)
	}
	if availableVpnsValue != nil {
		data.AvailableVpns = availableVpnsValue
		data.AvailableVpNs = availableVpnsValue
	}
	if availableVpnsValue != nil {
		result["availableVPNs"] = data.expandEndpointConnectionProfileAvailableVpnsList(ctx, availableVpnsValue, diags)
	}

	if data.Lockdown != nil && !isZeroStruct(*data.Lockdown) {
		result["lockdown"] = data.Lockdown.expandEndpointConnectionProfileLockdown(ctx, diags)
	}

	result["onFabricRuleSet"] = nil
	if data.OnFabricRuleSet != nil && !isZeroStruct(*data.OnFabricRuleSet) {
		result["onFabricRuleSet"] = data.OnFabricRuleSet.expandEndpointConnectionProfileOnFabricRuleSet(ctx, diags)
	}

	if data.OffNetSplitTunnel != nil && !isZeroStruct(*data.OffNetSplitTunnel) {
		result["offNetSplitTunnel"] = data.OffNetSplitTunnel.expandEndpointConnectionProfileOffNetSplitTunnel(ctx, diags)
	}

	if data.SplitTunnel != nil && !isZeroStruct(*data.SplitTunnel) {
		result["splitTunnel"] = data.SplitTunnel.expandEndpointConnectionProfileSplitTunnel(ctx, diags)
	}

	if !data.AllowInvalidServerCertificate.IsNull() && !data.AllowInvalidServerCertificate.IsUnknown() {
		result["allowInvalidServerCertificate"] = data.AllowInvalidServerCertificate.ValueString()
	}

	if !data.EndpointOnNetBypass.IsNull() && !data.EndpointOnNetBypass.IsUnknown() {
		result["endpointOnNetBypass"] = data.EndpointOnNetBypass.ValueBool()
	}

	if !data.AuthBeforeUserLogon.IsNull() && !data.AuthBeforeUserLogon.IsUnknown() {
		result["authBeforeUserLogon"] = data.AuthBeforeUserLogon.ValueBool()
	}

	if data.SecureInternetAccess != nil && !isZeroStruct(*data.SecureInternetAccess) {
		result["secureInternetAccess"] = data.SecureInternetAccess.expandEndpointConnectionProfileSecureInternetAccess(ctx, diags)
	}

	if !data.PreferredDtlsTunnel.IsNull() && !data.PreferredDtlsTunnel.IsUnknown() {
		result["preferredDTLSTunnel"] = data.PreferredDtlsTunnel.ValueString()
	}

	if !data.UseGuiSamlAuth.IsNull() && !data.UseGuiSamlAuth.IsUnknown() {
		result["useGuiSamlAuth"] = data.UseGuiSamlAuth.ValueString()
	}

	if !data.UseWebview2SamlAuth.IsNull() && !data.UseWebview2SamlAuth.IsUnknown() {
		result["useWebview2SamlAuth"] = data.UseWebview2SamlAuth.ValueString()
	}

	if !data.BeforeLogonSamlAuth.IsNull() && !data.BeforeLogonSamlAuth.IsUnknown() {
		result["beforeLogonSamlAuth"] = data.BeforeLogonSamlAuth.ValueString()
	}

	if !data.AfterLogonSamlAuth.IsNull() && !data.AfterLogonSamlAuth.IsUnknown() {
		result["afterLogonSamlAuth"] = data.AfterLogonSamlAuth.ValueString()
	}

	if !data.AllowPersonalVpns.IsNull() && !data.AllowPersonalVpns.IsUnknown() {
		result["allowPersonalVpns"] = data.AllowPersonalVpns.ValueBool()
	}

	if !data.MtuSize.IsNull() && !data.MtuSize.IsUnknown() {
		result["mtuSize"] = data.MtuSize.ValueFloat64()
	}

	if !data.VpnType.IsNull() && !data.VpnType.IsUnknown() {
		result["vpnType"] = data.VpnType.ValueString()
	}

	if !data.DisableInternetCheck.IsNull() && !data.DisableInternetCheck.IsUnknown() {
		result["disableInternetCheck"] = data.DisableInternetCheck.ValueString()
	}

	if !data.ShowDisconnectBtn.IsNull() && !data.ShowDisconnectBtn.IsUnknown() {
		result["showDisconnectBtn"] = data.ShowDisconnectBtn.ValueString()
	}

	if !data.EnableInvalidServerCertWarning.IsNull() && !data.EnableInvalidServerCertWarning.IsUnknown() {
		result["enableInvalidServerCertWarning"] = data.EnableInvalidServerCertWarning.ValueString()
	}

	if data.PreLogon != nil && !isZeroStruct(*data.PreLogon) {
		result["preLogon"] = data.PreLogon.expandEndpointConnectionProfilePreLogon(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointConnectionProfileModel) getUpdateObjectEndpointConnectionProfile(ctx context.Context, state resourceEndpointConnectionProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	connectToFortisaseValue := data.ConnectToFortisase
	connectToFortisaseValueSource := "connect_to_fortisase"
	if connectToFortisaseValue.IsNull() || connectToFortisaseValue.IsUnknown() {
		if !data.ConnectToFortiSase.IsNull() && !data.ConnectToFortiSase.IsUnknown() {
			connectToFortisaseValue = data.ConnectToFortiSase
			connectToFortisaseValueSource = "connect_to_forti_sase"
		}
	}
	if !data.ConnectToFortiSase.IsNull() && !data.ConnectToFortiSase.IsUnknown() && connectToFortisaseValueSource != "connect_to_forti_sase" && !connectToFortisaseValue.Equal(data.ConnectToFortiSase) {
		diags.AddError("Conflicting Terraform arguments",
			fmt.Sprintf("Arguments %q and %q cannot both be configured with different values.", connectToFortisaseValueSource, "connect_to_forti_sase"),
		)
	}
	if !connectToFortisaseValue.IsNull() && !connectToFortisaseValue.IsUnknown() {
		data.ConnectToFortisase = connectToFortisaseValue
		data.ConnectToFortiSase = connectToFortisaseValue
	}
	if !connectToFortisaseValue.IsNull() && !connectToFortisaseValue.IsUnknown() {
		result["connectToFortiSASE"] = connectToFortisaseValue.ValueString()
	}

	availableVpnsValue := data.AvailableVpns
	availableVpnsValueSource := "available_vpns"
	if availableVpnsValue == nil && data.AvailableVpNs != nil {
		availableVpnsValue = data.AvailableVpNs
		availableVpnsValueSource = "available_vp_ns"
	}
	if data.AvailableVpNs != nil && availableVpnsValueSource != "available_vp_ns" && !reflect.DeepEqual(availableVpnsValue, data.AvailableVpNs) {
		diags.AddError("Conflicting Terraform arguments",
			fmt.Sprintf("Arguments %q and %q cannot both be configured with different values.", availableVpnsValueSource, "available_vp_ns"),
		)
	}
	if availableVpnsValue != nil {
		data.AvailableVpns = availableVpnsValue
		data.AvailableVpNs = availableVpnsValue
	}
	if availableVpnsValue != nil {
		result["availableVPNs"] = data.expandEndpointConnectionProfileAvailableVpnsList(ctx, availableVpnsValue, diags)
	}

	if data.Lockdown != nil {
		result["lockdown"] = data.Lockdown.expandEndpointConnectionProfileLockdown(ctx, diags)
	}

	result["onFabricRuleSet"] = nil
	if data.OnFabricRuleSet != nil && !isZeroStruct(*data.OnFabricRuleSet) {
		result["onFabricRuleSet"] = data.OnFabricRuleSet.expandEndpointConnectionProfileOnFabricRuleSet(ctx, diags)
	}

	if data.OffNetSplitTunnel != nil {
		result["offNetSplitTunnel"] = data.OffNetSplitTunnel.expandEndpointConnectionProfileOffNetSplitTunnel(ctx, diags)
	}

	if data.SplitTunnel != nil {
		result["splitTunnel"] = data.SplitTunnel.expandEndpointConnectionProfileSplitTunnel(ctx, diags)
	}

	if !data.AllowInvalidServerCertificate.IsNull() && !data.AllowInvalidServerCertificate.IsUnknown() {
		result["allowInvalidServerCertificate"] = data.AllowInvalidServerCertificate.ValueString()
	}

	if !data.EndpointOnNetBypass.IsNull() && !data.EndpointOnNetBypass.IsUnknown() {
		result["endpointOnNetBypass"] = data.EndpointOnNetBypass.ValueBool()
	}

	if !data.AuthBeforeUserLogon.IsNull() && !data.AuthBeforeUserLogon.IsUnknown() {
		result["authBeforeUserLogon"] = data.AuthBeforeUserLogon.ValueBool()
	}

	if data.SecureInternetAccess != nil {
		result["secureInternetAccess"] = data.SecureInternetAccess.expandEndpointConnectionProfileSecureInternetAccess(ctx, diags)
	}

	if !data.PreferredDtlsTunnel.IsNull() && !data.PreferredDtlsTunnel.IsUnknown() {
		result["preferredDTLSTunnel"] = data.PreferredDtlsTunnel.ValueString()
	}

	if !data.UseGuiSamlAuth.IsNull() && !data.UseGuiSamlAuth.IsUnknown() {
		result["useGuiSamlAuth"] = data.UseGuiSamlAuth.ValueString()
	}

	if !data.UseWebview2SamlAuth.IsNull() && !data.UseWebview2SamlAuth.IsUnknown() {
		result["useWebview2SamlAuth"] = data.UseWebview2SamlAuth.ValueString()
	}

	if !data.BeforeLogonSamlAuth.IsNull() && !data.BeforeLogonSamlAuth.IsUnknown() {
		result["beforeLogonSamlAuth"] = data.BeforeLogonSamlAuth.ValueString()
	}

	if !data.AfterLogonSamlAuth.IsNull() && !data.AfterLogonSamlAuth.IsUnknown() {
		result["afterLogonSamlAuth"] = data.AfterLogonSamlAuth.ValueString()
	}

	if !data.AllowPersonalVpns.IsNull() && !data.AllowPersonalVpns.IsUnknown() {
		result["allowPersonalVpns"] = data.AllowPersonalVpns.ValueBool()
	}

	if !data.MtuSize.IsNull() && !data.MtuSize.IsUnknown() {
		result["mtuSize"] = data.MtuSize.ValueFloat64()
	}

	if !data.VpnType.IsNull() && !data.VpnType.IsUnknown() {
		result["vpnType"] = data.VpnType.ValueString()
	}

	if !data.DisableInternetCheck.IsNull() && !data.DisableInternetCheck.IsUnknown() {
		result["disableInternetCheck"] = data.DisableInternetCheck.ValueString()
	}

	if !data.ShowDisconnectBtn.IsNull() && !data.ShowDisconnectBtn.IsUnknown() {
		result["showDisconnectBtn"] = data.ShowDisconnectBtn.ValueString()
	}

	if !data.EnableInvalidServerCertWarning.IsNull() && !data.EnableInvalidServerCertWarning.IsUnknown() {
		result["enableInvalidServerCertWarning"] = data.EnableInvalidServerCertWarning.ValueString()
	}

	if data.PreLogon != nil {
		result["preLogon"] = data.PreLogon.expandEndpointConnectionProfilePreLogon(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointConnectionProfileModel) getURLObjectEndpointConnectionProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointConnectionProfileAvailableVpnsModel struct {
	Type                     types.String                                                                 `tfsdk:"type"`
	Name                     types.String                                                                 `tfsdk:"name"`
	RemoteGateway            types.String                                                                 `tfsdk:"remote_gateway"`
	UsernamePrompt           types.String                                                                 `tfsdk:"username_prompt"`
	SaveUsername             types.String                                                                 `tfsdk:"save_username"`
	ShowAlwaysUp             types.String                                                                 `tfsdk:"show_always_up"`
	ShowAutoConnect          types.String                                                                 `tfsdk:"show_auto_connect"`
	ShowRememberPassword     types.String                                                                 `tfsdk:"show_remember_password"`
	AuthenticateWithSso      types.String                                                                 `tfsdk:"authenticate_with_sso"`
	AllowFidoAuth            types.String                                                                 `tfsdk:"allow_fido_auth"`
	EnableLocalLan           types.String                                                                 `tfsdk:"enable_local_lan"`
	EncapsulationMode        types.String                                                                 `tfsdk:"encapsulation_mode"`
	UdpPort                  types.Float64                                                                `tfsdk:"udp_port"`
	TcpPort                  types.Float64                                                                `tfsdk:"tcp_port"`
	ConnectDisconnectScripts *resourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel `tfsdk:"connect_disconnect_scripts"`
	Port                     types.Float64                                                                `tfsdk:"port"`
	RequireCertificate       types.String                                                                 `tfsdk:"require_certificate"`
	ExternalBrowserSamlLogin types.String                                                                 `tfsdk:"external_browser_saml_login"`
	AuthMethod               types.String                                                                 `tfsdk:"auth_method"`
	DnsSuffixes              types.Set                                                                    `tfsdk:"dns_suffixes"`
	ShowPasscode             types.String                                                                 `tfsdk:"show_passcode"`
	PostureCheck             *resourceEndpointConnectionProfileAvailableVpnsPostureCheckModel             `tfsdk:"posture_check"`
	EapEnabled               types.Bool                                                                   `tfsdk:"eap_enabled"`
	SamlPort                 types.Float64                                                                `tfsdk:"saml_port"`
	PreSharedKey             types.String                                                                 `tfsdk:"pre_shared_key"`
}

type resourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel struct {
	OnConnectWindows    types.String `tfsdk:"on_connect_windows"`
	OnConnectMac        types.String `tfsdk:"on_connect_mac"`
	OnDisconnectWindows types.String `tfsdk:"on_disconnect_windows"`
	OnDisconnectMac     types.String `tfsdk:"on_disconnect_mac"`
}

type resourceEndpointConnectionProfileAvailableVpnsPostureCheckModel struct {
	Tag                types.String `tfsdk:"tag"`
	Action             types.String `tfsdk:"action"`
	CheckFailedMessage types.String `tfsdk:"check_failed_message"`
}

type resourceEndpointConnectionProfileLockdownModel struct {
	Status              types.String                                                       `tfsdk:"status"`
	GracePeriod         types.Float64                                                      `tfsdk:"grace_period"`
	MaxAttempts         types.Float64                                                      `tfsdk:"max_attempts"`
	Ips                 []resourceEndpointConnectionProfileLockdownIpsModel                `tfsdk:"ips"`
	Domains             []resourceEndpointConnectionProfileLockdownDomainsModel            `tfsdk:"domains"`
	DetectCaptivePortal *resourceEndpointConnectionProfileLockdownDetectCaptivePortalModel `tfsdk:"detect_captive_portal"`
}

type resourceEndpointConnectionProfileLockdownIpsModel struct {
	Ip       types.String `tfsdk:"ip"`
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
}

type resourceEndpointConnectionProfileLockdownDomainsModel struct {
	Address types.String `tfsdk:"address"`
}

type resourceEndpointConnectionProfileLockdownDetectCaptivePortalModel struct {
	Status                      types.String `tfsdk:"status"`
	DisableWindowsCaptivePortal types.String `tfsdk:"disable_windows_captive_portal"`
}

type resourceEndpointConnectionProfileOnFabricRuleSetModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointConnectionProfileOffNetSplitTunnelModel struct {
	LocalApps       types.Set                                                        `tfsdk:"local_apps"`
	Isdbs           []resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel   `tfsdk:"isdbs"`
	Fqdns           types.Set                                                        `tfsdk:"fqdns"`
	Subnets         []resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel `tfsdk:"subnets"`
	SubnetsIpsec    types.Set                                                        `tfsdk:"subnets_ipsec"`
	SplitTunnelMode types.String                                                     `tfsdk:"split_tunnel_mode"`
}

type resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointConnectionProfileSplitTunnelModel struct {
	LocalApps       types.Set                                                  `tfsdk:"local_apps"`
	Isdbs           []resourceEndpointConnectionProfileSplitTunnelIsdbsModel   `tfsdk:"isdbs"`
	Fqdns           types.Set                                                  `tfsdk:"fqdns"`
	Subnets         []resourceEndpointConnectionProfileSplitTunnelSubnetsModel `tfsdk:"subnets"`
	SubnetsIpsec    types.Set                                                  `tfsdk:"subnets_ipsec"`
	SplitTunnelMode types.String                                               `tfsdk:"split_tunnel_mode"`
}

type resourceEndpointConnectionProfileSplitTunnelIsdbsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointConnectionProfileSplitTunnelSubnetsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointConnectionProfileSecureInternetAccessModel struct {
	AuthenticateWithSso      types.String                                                                        `tfsdk:"authenticate_with_sso"`
	AllowFidoAuth            types.String                                                                        `tfsdk:"allow_fido_auth"`
	ConnectDisconnectScripts *resourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel `tfsdk:"connect_disconnect_scripts"`
	DnsSuffixes              types.Set                                                                           `tfsdk:"dns_suffixes"`
	EnableLocalLan           types.String                                                                        `tfsdk:"enable_local_lan"`
	FailoverSequence         types.Set                                                                           `tfsdk:"failover_sequence"`
	PostureCheck             *resourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel             `tfsdk:"posture_check"`
	EapEnabled               types.Bool                                                                          `tfsdk:"eap_enabled"`
	EncapsulationMode        types.String                                                                        `tfsdk:"encapsulation_mode"`
	ExternalBrowserSamlLogin types.String                                                                        `tfsdk:"external_browser_saml_login"`
}

type resourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel struct {
	OnConnectWindows    types.String `tfsdk:"on_connect_windows"`
	OnConnectMac        types.String `tfsdk:"on_connect_mac"`
	OnDisconnectWindows types.String `tfsdk:"on_disconnect_windows"`
	OnDisconnectMac     types.String `tfsdk:"on_disconnect_mac"`
}

type resourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel struct {
	Tag                types.String `tfsdk:"tag"`
	Action             types.String `tfsdk:"action"`
	CheckFailedMessage types.String `tfsdk:"check_failed_message"`
}

type resourceEndpointConnectionProfilePreLogonModel struct {
	VpnType       types.String                                              `tfsdk:"vpn_type"`
	RemoteGateway types.String                                              `tfsdk:"remote_gateway"`
	CommonName    *resourceEndpointConnectionProfilePreLogonCommonNameModel `tfsdk:"common_name"`
	Issuer        *resourceEndpointConnectionProfilePreLogonIssuerModel     `tfsdk:"issuer"`
	Port          types.Float64                                             `tfsdk:"port"`
}

type resourceEndpointConnectionProfilePreLogonCommonNameModel struct {
	MatchType types.String `tfsdk:"match_type"`
	Pattern   types.String `tfsdk:"pattern"`
}

type resourceEndpointConnectionProfilePreLogonIssuerModel struct {
	MatchType types.String `tfsdk:"match_type"`
	Pattern   types.String `tfsdk:"pattern"`
}

func (m *resourceEndpointConnectionProfileAvailableVpnsModel) flattenEndpointConnectionProfileAvailableVpns(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileAvailableVpnsModel {
	if input == nil {
		return &resourceEndpointConnectionProfileAvailableVpnsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileAvailableVpnsModel{}
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

func (s *resourceEndpointConnectionProfileModel) flattenEndpointConnectionProfileAvailableVpnsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfileAvailableVpnsModel {
	if o == nil {
		return []resourceEndpointConnectionProfileAvailableVpnsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument available_vpns is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfileAvailableVpnsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfileAvailableVpnsModel{}
	}

	values := make([]resourceEndpointConnectionProfileAvailableVpnsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfileAvailableVpnsModel
		if i < len(s.AvailableVpns) {
			m = s.AvailableVpns[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileAvailableVpns(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel) flattenEndpointConnectionProfileAvailableVpnsConnectDisconnectScripts(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel {
	if input == nil {
		return &resourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel{}
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

func (m *resourceEndpointConnectionProfileAvailableVpnsPostureCheckModel) flattenEndpointConnectionProfileAvailableVpnsPostureCheck(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileAvailableVpnsPostureCheckModel {
	if input == nil {
		return &resourceEndpointConnectionProfileAvailableVpnsPostureCheckModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileAvailableVpnsPostureCheckModel{}
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

func (m *resourceEndpointConnectionProfileLockdownModel) flattenEndpointConnectionProfileLockdown(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileLockdownModel {
	if input == nil {
		return &resourceEndpointConnectionProfileLockdownModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileLockdownModel{}
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

func (m *resourceEndpointConnectionProfileLockdownIpsModel) flattenEndpointConnectionProfileLockdownIps(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileLockdownIpsModel {
	if input == nil {
		return &resourceEndpointConnectionProfileLockdownIpsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileLockdownIpsModel{}
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

func (s *resourceEndpointConnectionProfileLockdownModel) flattenEndpointConnectionProfileLockdownIpsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfileLockdownIpsModel {
	if o == nil {
		return []resourceEndpointConnectionProfileLockdownIpsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument ips is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfileLockdownIpsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfileLockdownIpsModel{}
	}

	values := make([]resourceEndpointConnectionProfileLockdownIpsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfileLockdownIpsModel
		if i < len(s.Ips) {
			m = s.Ips[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileLockdownIps(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfileLockdownDomainsModel) flattenEndpointConnectionProfileLockdownDomains(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileLockdownDomainsModel {
	if input == nil {
		return &resourceEndpointConnectionProfileLockdownDomainsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileLockdownDomainsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["address"]; ok {
		m.Address = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointConnectionProfileLockdownModel) flattenEndpointConnectionProfileLockdownDomainsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfileLockdownDomainsModel {
	if o == nil {
		return []resourceEndpointConnectionProfileLockdownDomainsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument domains is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfileLockdownDomainsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfileLockdownDomainsModel{}
	}

	values := make([]resourceEndpointConnectionProfileLockdownDomainsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfileLockdownDomainsModel
		if i < len(s.Domains) {
			m = s.Domains[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileLockdownDomains(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfileLockdownDetectCaptivePortalModel) flattenEndpointConnectionProfileLockdownDetectCaptivePortal(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileLockdownDetectCaptivePortalModel {
	if input == nil {
		return &resourceEndpointConnectionProfileLockdownDetectCaptivePortalModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileLockdownDetectCaptivePortalModel{}
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

func (m *resourceEndpointConnectionProfileOnFabricRuleSetModel) flattenEndpointConnectionProfileOnFabricRuleSet(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileOnFabricRuleSetModel {
	if input == nil {
		return &resourceEndpointConnectionProfileOnFabricRuleSetModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileOnFabricRuleSetModel{}
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

func (m *resourceEndpointConnectionProfileOffNetSplitTunnelModel) flattenEndpointConnectionProfileOffNetSplitTunnel(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileOffNetSplitTunnelModel {
	if input == nil {
		return &resourceEndpointConnectionProfileOffNetSplitTunnelModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileOffNetSplitTunnelModel{}
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

func (m *resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel) flattenEndpointConnectionProfileOffNetSplitTunnelIsdbs(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel {
	if input == nil {
		return &resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel{}
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

func (s *resourceEndpointConnectionProfileOffNetSplitTunnelModel) flattenEndpointConnectionProfileOffNetSplitTunnelIsdbsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel {
	if o == nil {
		return []resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument isdbs is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel{}
	}

	values := make([]resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel
		if i < len(s.Isdbs) {
			m = s.Isdbs[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileOffNetSplitTunnelIsdbs(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel) flattenEndpointConnectionProfileOffNetSplitTunnelSubnets(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel {
	if input == nil {
		return &resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel{}
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

func (s *resourceEndpointConnectionProfileOffNetSplitTunnelModel) flattenEndpointConnectionProfileOffNetSplitTunnelSubnetsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel {
	if o == nil {
		return []resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument subnets is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel{}
	}

	values := make([]resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel
		if i < len(s.Subnets) {
			m = s.Subnets[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileOffNetSplitTunnelSubnets(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfileSplitTunnelModel) flattenEndpointConnectionProfileSplitTunnel(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileSplitTunnelModel {
	if input == nil {
		return &resourceEndpointConnectionProfileSplitTunnelModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileSplitTunnelModel{}
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

func (m *resourceEndpointConnectionProfileSplitTunnelIsdbsModel) flattenEndpointConnectionProfileSplitTunnelIsdbs(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileSplitTunnelIsdbsModel {
	if input == nil {
		return &resourceEndpointConnectionProfileSplitTunnelIsdbsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileSplitTunnelIsdbsModel{}
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

func (s *resourceEndpointConnectionProfileSplitTunnelModel) flattenEndpointConnectionProfileSplitTunnelIsdbsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfileSplitTunnelIsdbsModel {
	if o == nil {
		return []resourceEndpointConnectionProfileSplitTunnelIsdbsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument isdbs is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfileSplitTunnelIsdbsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfileSplitTunnelIsdbsModel{}
	}

	values := make([]resourceEndpointConnectionProfileSplitTunnelIsdbsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfileSplitTunnelIsdbsModel
		if i < len(s.Isdbs) {
			m = s.Isdbs[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileSplitTunnelIsdbs(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfileSplitTunnelSubnetsModel) flattenEndpointConnectionProfileSplitTunnelSubnets(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileSplitTunnelSubnetsModel {
	if input == nil {
		return &resourceEndpointConnectionProfileSplitTunnelSubnetsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileSplitTunnelSubnetsModel{}
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

func (s *resourceEndpointConnectionProfileSplitTunnelModel) flattenEndpointConnectionProfileSplitTunnelSubnetsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfileSplitTunnelSubnetsModel {
	if o == nil {
		return []resourceEndpointConnectionProfileSplitTunnelSubnetsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument subnets is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfileSplitTunnelSubnetsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfileSplitTunnelSubnetsModel{}
	}

	values := make([]resourceEndpointConnectionProfileSplitTunnelSubnetsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfileSplitTunnelSubnetsModel
		if i < len(s.Subnets) {
			m = s.Subnets[i]
		}
		values[i] = *m.flattenEndpointConnectionProfileSplitTunnelSubnets(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfileSecureInternetAccessModel) flattenEndpointConnectionProfileSecureInternetAccess(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileSecureInternetAccessModel {
	if input == nil {
		return &resourceEndpointConnectionProfileSecureInternetAccessModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileSecureInternetAccessModel{}
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
		if m.FailoverSequence.IsNull() || !isSetSuperset(v, m.FailoverSequence.Elements()) {
			m.FailoverSequence = parseSetValue(ctx, v, types.StringType)
		}
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

func (m *resourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel) flattenEndpointConnectionProfileSecureInternetAccessConnectDisconnectScripts(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel {
	if input == nil {
		return &resourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel{}
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

func (m *resourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel) flattenEndpointConnectionProfileSecureInternetAccessPostureCheck(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel {
	if input == nil {
		return &resourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel{}
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

func (m *resourceEndpointConnectionProfilePreLogonModel) flattenEndpointConnectionProfilePreLogon(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilePreLogonModel {
	if input == nil {
		return &resourceEndpointConnectionProfilePreLogonModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilePreLogonModel{}
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

func (m *resourceEndpointConnectionProfilePreLogonCommonNameModel) flattenEndpointConnectionProfilePreLogonCommonName(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilePreLogonCommonNameModel {
	if input == nil {
		return &resourceEndpointConnectionProfilePreLogonCommonNameModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilePreLogonCommonNameModel{}
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

func (m *resourceEndpointConnectionProfilePreLogonIssuerModel) flattenEndpointConnectionProfilePreLogonIssuer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilePreLogonIssuerModel {
	if input == nil {
		return &resourceEndpointConnectionProfilePreLogonIssuerModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilePreLogonIssuerModel{}
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

func (data *resourceEndpointConnectionProfileAvailableVpnsModel) expandEndpointConnectionProfileAvailableVpns(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		result["name"] = data.Name.ValueString()
	}

	if !data.RemoteGateway.IsNull() && !data.RemoteGateway.IsUnknown() {
		result["remoteGateway"] = data.RemoteGateway.ValueString()
	}

	if !data.UsernamePrompt.IsNull() && !data.UsernamePrompt.IsUnknown() {
		result["usernamePrompt"] = data.UsernamePrompt.ValueString()
	}

	if !data.SaveUsername.IsNull() && !data.SaveUsername.IsUnknown() {
		result["saveUsername"] = data.SaveUsername.ValueString()
	}

	if !data.ShowAlwaysUp.IsNull() && !data.ShowAlwaysUp.IsUnknown() {
		result["showAlwaysUp"] = data.ShowAlwaysUp.ValueString()
	}

	if !data.ShowAutoConnect.IsNull() && !data.ShowAutoConnect.IsUnknown() {
		result["showAutoConnect"] = data.ShowAutoConnect.ValueString()
	}

	if !data.ShowRememberPassword.IsNull() && !data.ShowRememberPassword.IsUnknown() {
		result["showRememberPassword"] = data.ShowRememberPassword.ValueString()
	}

	if !data.AuthenticateWithSso.IsNull() && !data.AuthenticateWithSso.IsUnknown() {
		result["authenticateWithSSO"] = data.AuthenticateWithSso.ValueString()
	}

	if !data.AllowFidoAuth.IsNull() && !data.AllowFidoAuth.IsUnknown() {
		result["allowFidoAuth"] = data.AllowFidoAuth.ValueString()
	}

	if !data.EnableLocalLan.IsNull() && !data.EnableLocalLan.IsUnknown() {
		result["enableLocalLan"] = data.EnableLocalLan.ValueString()
	}

	if !data.EncapsulationMode.IsNull() && !data.EncapsulationMode.IsUnknown() {
		result["encapsulationMode"] = data.EncapsulationMode.ValueString()
	}

	if !data.UdpPort.IsNull() && !data.UdpPort.IsUnknown() {
		result["udpPort"] = data.UdpPort.ValueFloat64()
	}

	if !data.TcpPort.IsNull() && !data.TcpPort.IsUnknown() {
		result["tcpPort"] = data.TcpPort.ValueFloat64()
	}

	if data.ConnectDisconnectScripts != nil && !isZeroStruct(*data.ConnectDisconnectScripts) {
		result["connectDisconnectScripts"] = data.ConnectDisconnectScripts.expandEndpointConnectionProfileAvailableVpnsConnectDisconnectScripts(ctx, diags)
	}

	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		result["port"] = data.Port.ValueFloat64()
	}

	if !data.RequireCertificate.IsNull() && !data.RequireCertificate.IsUnknown() {
		result["requireCertificate"] = data.RequireCertificate.ValueString()
	}

	if !data.ExternalBrowserSamlLogin.IsNull() && !data.ExternalBrowserSamlLogin.IsUnknown() {
		result["externalBrowserSamlLogin"] = data.ExternalBrowserSamlLogin.ValueString()
	}

	if !data.AuthMethod.IsNull() && !data.AuthMethod.IsUnknown() {
		result["authMethod"] = data.AuthMethod.ValueString()
	}

	if !data.DnsSuffixes.IsNull() && !data.DnsSuffixes.IsUnknown() {
		result["dnsSuffixes"] = expandSetToStringList(data.DnsSuffixes)
	}

	if !data.ShowPasscode.IsNull() && !data.ShowPasscode.IsUnknown() {
		result["showPasscode"] = data.ShowPasscode.ValueString()
	}

	if data.PostureCheck != nil && !isZeroStruct(*data.PostureCheck) {
		result["postureCheck"] = data.PostureCheck.expandEndpointConnectionProfileAvailableVpnsPostureCheck(ctx, diags)
	}

	if !data.EapEnabled.IsNull() && !data.EapEnabled.IsUnknown() {
		result["eapEnabled"] = data.EapEnabled.ValueBool()
	}

	if !data.SamlPort.IsNull() && !data.SamlPort.IsUnknown() {
		result["samlPort"] = data.SamlPort.ValueFloat64()
	}

	if !data.PreSharedKey.IsNull() && !data.PreSharedKey.IsUnknown() {
		result["preSharedKey"] = data.PreSharedKey.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfileModel) expandEndpointConnectionProfileAvailableVpnsList(ctx context.Context, l []resourceEndpointConnectionProfileAvailableVpnsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfileAvailableVpns(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfileAvailableVpnsConnectDisconnectScriptsModel) expandEndpointConnectionProfileAvailableVpnsConnectDisconnectScripts(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.OnConnectWindows.IsNull() && !data.OnConnectWindows.IsUnknown() {
		result["onConnectWindows"] = data.OnConnectWindows.ValueString()
	}

	if !data.OnConnectMac.IsNull() && !data.OnConnectMac.IsUnknown() {
		result["onConnectMac"] = data.OnConnectMac.ValueString()
	}

	if !data.OnDisconnectWindows.IsNull() && !data.OnDisconnectWindows.IsUnknown() {
		result["onDisconnectWindows"] = data.OnDisconnectWindows.ValueString()
	}

	if !data.OnDisconnectMac.IsNull() && !data.OnDisconnectMac.IsUnknown() {
		result["onDisconnectMac"] = data.OnDisconnectMac.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfileAvailableVpnsPostureCheckModel) expandEndpointConnectionProfileAvailableVpnsPostureCheck(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	if !data.CheckFailedMessage.IsNull() && !data.CheckFailedMessage.IsUnknown() {
		result["checkFailedMessage"] = data.CheckFailedMessage.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfileLockdownModel) expandEndpointConnectionProfileLockdown(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if !data.GracePeriod.IsNull() && !data.GracePeriod.IsUnknown() {
		result["gracePeriod"] = data.GracePeriod.ValueFloat64()
	}

	if !data.MaxAttempts.IsNull() && !data.MaxAttempts.IsUnknown() {
		result["maxAttempts"] = data.MaxAttempts.ValueFloat64()
	}

	if data.Ips != nil {
		result["ips"] = data.expandEndpointConnectionProfileLockdownIpsList(ctx, data.Ips, diags)
	}

	if data.Domains != nil {
		result["domains"] = data.expandEndpointConnectionProfileLockdownDomainsList(ctx, data.Domains, diags)
	}

	if data.DetectCaptivePortal != nil && !isZeroStruct(*data.DetectCaptivePortal) {
		result["detectCaptivePortal"] = data.DetectCaptivePortal.expandEndpointConnectionProfileLockdownDetectCaptivePortal(ctx, diags)
	}

	return result
}

func (data *resourceEndpointConnectionProfileLockdownIpsModel) expandEndpointConnectionProfileLockdownIps(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		result["ip"] = data.Ip.ValueString()
	}

	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		result["port"] = data.Port.ValueString()
	}

	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		result["protocol"] = data.Protocol.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfileLockdownModel) expandEndpointConnectionProfileLockdownIpsList(ctx context.Context, l []resourceEndpointConnectionProfileLockdownIpsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfileLockdownIps(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfileLockdownDomainsModel) expandEndpointConnectionProfileLockdownDomains(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Address.IsNull() && !data.Address.IsUnknown() {
		result["address"] = data.Address.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfileLockdownModel) expandEndpointConnectionProfileLockdownDomainsList(ctx context.Context, l []resourceEndpointConnectionProfileLockdownDomainsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfileLockdownDomains(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfileLockdownDetectCaptivePortalModel) expandEndpointConnectionProfileLockdownDetectCaptivePortal(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if !data.DisableWindowsCaptivePortal.IsNull() && !data.DisableWindowsCaptivePortal.IsUnknown() {
		result["disableWindowsCaptivePortal"] = data.DisableWindowsCaptivePortal.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfileOnFabricRuleSetModel) expandEndpointConnectionProfileOnFabricRuleSet(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfileOffNetSplitTunnelModel) expandEndpointConnectionProfileOffNetSplitTunnel(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.LocalApps.IsNull() && !data.LocalApps.IsUnknown() {
		result["localApps"] = expandSetToStringList(data.LocalApps)
	}

	result["isdbs"] = data.expandEndpointConnectionProfileOffNetSplitTunnelIsdbsList(ctx, data.Isdbs, diags)

	if !data.Fqdns.IsNull() && !data.Fqdns.IsUnknown() {
		result["fqdns"] = expandSetToStringList(data.Fqdns)
	}

	if data.Subnets != nil {
		result["subnets"] = data.expandEndpointConnectionProfileOffNetSplitTunnelSubnetsList(ctx, data.Subnets, diags)
	}

	if !data.SubnetsIpsec.IsNull() && !data.SubnetsIpsec.IsUnknown() {
		result["subnetsIpsec"] = expandSetToStringList(data.SubnetsIpsec)
	}

	if !data.SplitTunnelMode.IsNull() && !data.SplitTunnelMode.IsUnknown() {
		result["splitTunnelMode"] = data.SplitTunnelMode.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel) expandEndpointConnectionProfileOffNetSplitTunnelIsdbs(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfileOffNetSplitTunnelModel) expandEndpointConnectionProfileOffNetSplitTunnelIsdbsList(ctx context.Context, l []resourceEndpointConnectionProfileOffNetSplitTunnelIsdbsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfileOffNetSplitTunnelIsdbs(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel) expandEndpointConnectionProfileOffNetSplitTunnelSubnets(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfileOffNetSplitTunnelModel) expandEndpointConnectionProfileOffNetSplitTunnelSubnetsList(ctx context.Context, l []resourceEndpointConnectionProfileOffNetSplitTunnelSubnetsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfileOffNetSplitTunnelSubnets(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfileSplitTunnelModel) expandEndpointConnectionProfileSplitTunnel(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.LocalApps.IsNull() && !data.LocalApps.IsUnknown() {
		result["localApps"] = expandSetToStringList(data.LocalApps)
	}

	result["isdbs"] = data.expandEndpointConnectionProfileSplitTunnelIsdbsList(ctx, data.Isdbs, diags)

	if !data.Fqdns.IsNull() && !data.Fqdns.IsUnknown() {
		result["fqdns"] = expandSetToStringList(data.Fqdns)
	}

	if data.Subnets != nil {
		result["subnets"] = data.expandEndpointConnectionProfileSplitTunnelSubnetsList(ctx, data.Subnets, diags)
	}

	if !data.SubnetsIpsec.IsNull() && !data.SubnetsIpsec.IsUnknown() {
		result["subnetsIpsec"] = expandSetToStringList(data.SubnetsIpsec)
	}

	if !data.SplitTunnelMode.IsNull() && !data.SplitTunnelMode.IsUnknown() {
		result["splitTunnelMode"] = data.SplitTunnelMode.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfileSplitTunnelIsdbsModel) expandEndpointConnectionProfileSplitTunnelIsdbs(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfileSplitTunnelModel) expandEndpointConnectionProfileSplitTunnelIsdbsList(ctx context.Context, l []resourceEndpointConnectionProfileSplitTunnelIsdbsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfileSplitTunnelIsdbs(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfileSplitTunnelSubnetsModel) expandEndpointConnectionProfileSplitTunnelSubnets(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfileSplitTunnelModel) expandEndpointConnectionProfileSplitTunnelSubnetsList(ctx context.Context, l []resourceEndpointConnectionProfileSplitTunnelSubnetsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfileSplitTunnelSubnets(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfileSecureInternetAccessModel) expandEndpointConnectionProfileSecureInternetAccess(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AuthenticateWithSso.IsNull() && !data.AuthenticateWithSso.IsUnknown() {
		result["authenticateWithSSO"] = data.AuthenticateWithSso.ValueString()
	}

	if !data.AllowFidoAuth.IsNull() && !data.AllowFidoAuth.IsUnknown() {
		result["allowFidoAuth"] = data.AllowFidoAuth.ValueString()
	}

	if data.ConnectDisconnectScripts != nil && !isZeroStruct(*data.ConnectDisconnectScripts) {
		result["connectDisconnectScripts"] = data.ConnectDisconnectScripts.expandEndpointConnectionProfileSecureInternetAccessConnectDisconnectScripts(ctx, diags)
	}

	if !data.DnsSuffixes.IsNull() && !data.DnsSuffixes.IsUnknown() {
		result["dnsSuffixes"] = expandSetToStringList(data.DnsSuffixes)
	}

	if !data.EnableLocalLan.IsNull() && !data.EnableLocalLan.IsUnknown() {
		result["enableLocalLan"] = data.EnableLocalLan.ValueString()
	}

	if !data.FailoverSequence.IsNull() && !data.FailoverSequence.IsUnknown() {
		result["failoverSequence"] = expandSetToStringList(data.FailoverSequence)
	}

	if data.PostureCheck != nil && !isZeroStruct(*data.PostureCheck) {
		result["postureCheck"] = data.PostureCheck.expandEndpointConnectionProfileSecureInternetAccessPostureCheck(ctx, diags)
	}

	if !data.EapEnabled.IsNull() && !data.EapEnabled.IsUnknown() {
		result["eapEnabled"] = data.EapEnabled.ValueBool()
	}

	if !data.EncapsulationMode.IsNull() && !data.EncapsulationMode.IsUnknown() {
		result["encapsulationMode"] = data.EncapsulationMode.ValueString()
	}

	if !data.ExternalBrowserSamlLogin.IsNull() && !data.ExternalBrowserSamlLogin.IsUnknown() {
		result["externalBrowserSamlLogin"] = data.ExternalBrowserSamlLogin.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfileSecureInternetAccessConnectDisconnectScriptsModel) expandEndpointConnectionProfileSecureInternetAccessConnectDisconnectScripts(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.OnConnectWindows.IsNull() && !data.OnConnectWindows.IsUnknown() {
		result["onConnectWindows"] = data.OnConnectWindows.ValueString()
	}

	if !data.OnConnectMac.IsNull() && !data.OnConnectMac.IsUnknown() {
		result["onConnectMac"] = data.OnConnectMac.ValueString()
	}

	if !data.OnDisconnectWindows.IsNull() && !data.OnDisconnectWindows.IsUnknown() {
		result["onDisconnectWindows"] = data.OnDisconnectWindows.ValueString()
	}

	if !data.OnDisconnectMac.IsNull() && !data.OnDisconnectMac.IsUnknown() {
		result["onDisconnectMac"] = data.OnDisconnectMac.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfileSecureInternetAccessPostureCheckModel) expandEndpointConnectionProfileSecureInternetAccessPostureCheck(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	if !data.CheckFailedMessage.IsNull() && !data.CheckFailedMessage.IsUnknown() {
		result["checkFailedMessage"] = data.CheckFailedMessage.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfilePreLogonModel) expandEndpointConnectionProfilePreLogon(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.VpnType.IsNull() && !data.VpnType.IsUnknown() {
		result["vpnType"] = data.VpnType.ValueString()
	}

	if !data.RemoteGateway.IsNull() && !data.RemoteGateway.IsUnknown() {
		result["remoteGateway"] = data.RemoteGateway.ValueString()
	}

	if data.CommonName != nil && !isZeroStruct(*data.CommonName) {
		result["commonName"] = data.CommonName.expandEndpointConnectionProfilePreLogonCommonName(ctx, diags)
	}

	if data.Issuer != nil && !isZeroStruct(*data.Issuer) {
		result["issuer"] = data.Issuer.expandEndpointConnectionProfilePreLogonIssuer(ctx, diags)
	}

	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		result["port"] = data.Port.ValueFloat64()
	}

	return result
}

func (data *resourceEndpointConnectionProfilePreLogonCommonNameModel) expandEndpointConnectionProfilePreLogonCommonName(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.MatchType.IsNull() && !data.MatchType.IsUnknown() {
		result["matchType"] = data.MatchType.ValueString()
	}

	if !data.Pattern.IsNull() && !data.Pattern.IsUnknown() {
		result["pattern"] = data.Pattern.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfilePreLogonIssuerModel) expandEndpointConnectionProfilePreLogonIssuer(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.MatchType.IsNull() && !data.MatchType.IsUnknown() {
		result["matchType"] = data.MatchType.ValueString()
	}

	if !data.Pattern.IsNull() && !data.Pattern.IsUnknown() {
		result["pattern"] = data.Pattern.ValueString()
	}

	return result
}
