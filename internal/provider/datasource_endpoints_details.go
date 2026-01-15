// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointsDetails{}

func newDatasourceEndpointsDetails() datasource.DataSource {
	return &datasourceEndpointsDetails{}
}

type datasourceEndpointsDetails struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointsDetailsModel describes the datasource data model.
type datasourceEndpointsDetailsModel struct {
	DeviceId                types.Float64                                   `tfsdk:"device_id"`
	Host                    types.String                                    `tfsdk:"host"`
	Alias                   types.String                                    `tfsdk:"alias"`
	Name                    types.String                                    `tfsdk:"name"`
	IpAddr                  types.String                                    `tfsdk:"ip_addr"`
	MacAddr                 types.String                                    `tfsdk:"mac_addr"`
	OsVersion               types.String                                    `tfsdk:"os_version"`
	OsServicePack           types.String                                    `tfsdk:"os_service_pack"`
	DistinguishedName       types.String                                    `tfsdk:"distinguished_name"`
	GroupTag                types.String                                    `tfsdk:"group_tag"`
	GroupId                 types.Float64                                   `tfsdk:"group_id"`
	GroupName               types.String                                    `tfsdk:"group_name"`
	OrigGroupName           types.String                                    `tfsdk:"orig_group_name"`
	OrigGroupId             types.Float64                                   `tfsdk:"orig_group_id"`
	DomainId                types.Float64                                   `tfsdk:"domain_id"`
	InstallerName           types.String                                    `tfsdk:"installer_name"`
	DeploymentState         types.Float64                                   `tfsdk:"deployment_state"`
	ScheduledInstallTime    types.String                                    `tfsdk:"scheduled_install_time"`
	InstallationState       types.Float64                                   `tfsdk:"installation_state"`
	DeploymentStateTime     types.String                                    `tfsdk:"deployment_state_time"`
	DeploymentStateData     types.String                                    `tfsdk:"deployment_state_data"`
	ForticlientId           types.Float64                                   `tfsdk:"forticlient_id"`
	Uid                     types.String                                    `tfsdk:"uid"`
	Caps                    types.Float64                                   `tfsdk:"caps"`
	FctSn                   types.String                                    `tfsdk:"fct_sn"`
	FgtSn                   types.String                                    `tfsdk:"fgt_sn"`
	LastSeen                types.Float64                                   `tfsdk:"last_seen"`
	Deregister              types.Float64                                   `tfsdk:"deregister"`
	RunCmd                  types.Float64                                   `tfsdk:"run_cmd"`
	QuarantineMessage       types.String                                    `tfsdk:"quarantine_message"`
	IsInstalled             types.Bool                                      `tfsdk:"is_installed"`
	IsManaged               types.Bool                                      `tfsdk:"is_managed"`
	IsMigrating             types.Bool                                      `tfsdk:"is_migrating"`
	IsEmsRegistered         types.Bool                                      `tfsdk:"is_ems_registered"`
	IsEmsOnline             types.Bool                                      `tfsdk:"is_ems_online"`
	IsEmsOnnet              types.Bool                                      `tfsdk:"is_ems_onnet"`
	IsExcluded              types.Bool                                      `tfsdk:"is_excluded"`
	IsQuarantined           types.Bool                                      `tfsdk:"is_quarantined"`
	QuarantineAccessCode    types.String                                    `tfsdk:"quarantine_access_code"`
	FctVersion              types.String                                    `tfsdk:"fct_version"`
	ComparableFctVersion    types.Float64                                   `tfsdk:"comparable_fct_version"`
	UserDomain              types.String                                    `tfsdk:"user_domain"`
	Service                 types.String                                    `tfsdk:"service"`
	ProfileName             types.String                                    `tfsdk:"profile_name"`
	IpListName              types.String                                    `tfsdk:"ip_list_name"`
	ZtnaEnabled             types.Bool                                      `tfsdk:"ztna_enabled"`
	ZtnaSerial              types.String                                    `tfsdk:"ztna_serial"`
	AvInstalled             types.Bool                                      `tfsdk:"av_installed"`
	AvEnabled               types.Bool                                      `tfsdk:"av_enabled"`
	AvHidden                types.Bool                                      `tfsdk:"av_hidden"`
	FwInstalled             types.Bool                                      `tfsdk:"fw_installed"`
	FwEnabled               types.Bool                                      `tfsdk:"fw_enabled"`
	FwHidden                types.Bool                                      `tfsdk:"fw_hidden"`
	WfInstalled             types.Bool                                      `tfsdk:"wf_installed"`
	WfEnabled               types.Bool                                      `tfsdk:"wf_enabled"`
	WfHidden                types.Bool                                      `tfsdk:"wf_hidden"`
	VpnInstalled            types.Bool                                      `tfsdk:"vpn_installed"`
	VpnEnabled              types.Bool                                      `tfsdk:"vpn_enabled"`
	VpnHidden               types.Bool                                      `tfsdk:"vpn_hidden"`
	VulnInstalled           types.Bool                                      `tfsdk:"vuln_installed"`
	VulnEnabled             types.Bool                                      `tfsdk:"vuln_enabled"`
	VulnHidden              types.Bool                                      `tfsdk:"vuln_hidden"`
	SsomaInstalled          types.Bool                                      `tfsdk:"ssoma_installed"`
	SsomaEnabled            types.Bool                                      `tfsdk:"ssoma_enabled"`
	SsomaHidden             types.Bool                                      `tfsdk:"ssoma_hidden"`
	SbInstalled             types.Bool                                      `tfsdk:"sb_installed"`
	SbEnabled               types.Bool                                      `tfsdk:"sb_enabled"`
	SbHidden                types.Bool                                      `tfsdk:"sb_hidden"`
	RsInstalled             types.Bool                                      `tfsdk:"rs_installed"`
	RsEnabled               types.Bool                                      `tfsdk:"rs_enabled"`
	RsHidden                types.Bool                                      `tfsdk:"rs_hidden"`
	AvLastScanType          types.Float64                                   `tfsdk:"av_last_scan_type"`
	AvLastScanDate          types.String                                    `tfsdk:"av_last_scan_date"`
	AvLastFullScanDate      types.String                                    `tfsdk:"av_last_full_scan_date"`
	AvLastCancelledScanType types.Float64                                   `tfsdk:"av_last_cancelled_scan_type"`
	AvLastCancelledScanDate types.String                                    `tfsdk:"av_last_cancelled_scan_date"`
	AvScanScheduled         types.Bool                                      `tfsdk:"av_scan_scheduled"`
	AvNextSchType           types.Float64                                   `tfsdk:"av_next_sch_type"`
	AvNextScanOn            types.Float64                                   `tfsdk:"av_next_scan_on"`
	AvNextScanHour          types.Float64                                   `tfsdk:"av_next_scan_hour"`
	AvNextScanMin           types.Float64                                   `tfsdk:"av_next_scan_min"`
	AvNextScanType          types.Float64                                   `tfsdk:"av_next_scan_type"`
	IsAvScanning            types.Bool                                      `tfsdk:"is_av_scanning"`
	LastVulnScan            types.Float64                                   `tfsdk:"last_vuln_scan"`
	VulnScanStatus          types.String                                    `tfsdk:"vuln_scan_status"`
	VulnNextScheduled       types.Bool                                      `tfsdk:"vuln_next_scheduled"`
	VulnNextSchType         types.Float64                                   `tfsdk:"vuln_next_sch_type"`
	VulnNextScanOn          types.Float64                                   `tfsdk:"vuln_next_scan_on"`
	VulnNextStartHour       types.Float64                                   `tfsdk:"vuln_next_start_hour"`
	VulnNextStartMin        types.Float64                                   `tfsdk:"vuln_next_start_min"`
	IsVulnScanning          types.Bool                                      `tfsdk:"is_vuln_scanning"`
	AvEventsCount           types.Float64                                   `tfsdk:"av_events_count"`
	SbEventsCount           types.Float64                                   `tfsdk:"sb_events_count"`
	FwEventsCount           types.Float64                                   `tfsdk:"fw_events_count"`
	WfEventsCount           types.Float64                                   `tfsdk:"wf_events_count"`
	VulnEventsCount         types.Float64                                   `tfsdk:"vuln_events_count"`
	SysEventsCount          types.Float64                                   `tfsdk:"sys_events_count"`
	VulnEventsMaxSeverity   types.Float64                                   `tfsdk:"vuln_events_max_severity"`
	ConnDetails             []datasourceEndpointsDetailsConnDetailsModel    `tfsdk:"conn_details"`
	HardwareDetails         *datasourceEndpointsDetailsHardwareDetailsModel `tfsdk:"hardware_details"`
	Forensics               *datasourceEndpointsDetailsForensicsModel       `tfsdk:"forensics"`
	ForensicsEnabled        types.Bool                                      `tfsdk:"forensics_enabled"`
	Tags                    []datasourceEndpointsDetailsTagsModel           `tfsdk:"tags"`
}

func (r *datasourceEndpointsDetails) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoints_details"
}

func (r *datasourceEndpointsDetails) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_id": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
				MarkdownDescription: "The device ID of the endpoint.\nValue at least 1.",
				Required:            true,
			},
			"host": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"alias": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ip_addr": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"mac_addr": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"os_version": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"os_service_pack": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"distinguished_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"group_tag": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"group_id": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"group_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"orig_group_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"orig_group_id": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"domain_id": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"installer_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"deployment_state": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"scheduled_install_time": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"installation_state": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"deployment_state_time": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"deployment_state_data": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"forticlient_id": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"uid": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"caps": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"fct_sn": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"fgt_sn": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"last_seen": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"deregister": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"run_cmd": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"quarantine_message": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"is_installed": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"is_managed": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"is_migrating": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"is_ems_registered": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"is_ems_online": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"is_ems_onnet": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"is_excluded": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"is_quarantined": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"quarantine_access_code": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"fct_version": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"comparable_fct_version": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"user_domain": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"service": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"profile_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ip_list_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ztna_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"ztna_serial": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"av_installed": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"av_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"av_hidden": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"fw_installed": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"fw_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"fw_hidden": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"wf_installed": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"wf_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"wf_hidden": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"vpn_installed": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"vpn_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"vpn_hidden": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"vuln_installed": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"vuln_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"vuln_hidden": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"ssoma_installed": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"ssoma_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"ssoma_hidden": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"sb_installed": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"sb_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"sb_hidden": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"rs_installed": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"rs_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"rs_hidden": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"av_last_scan_type": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"av_last_scan_date": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"av_last_full_scan_date": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"av_last_cancelled_scan_type": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"av_last_cancelled_scan_date": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"av_scan_scheduled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"av_next_sch_type": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"av_next_scan_on": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"av_next_scan_hour": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"av_next_scan_min": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"av_next_scan_type": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"is_av_scanning": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"last_vuln_scan": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"vuln_scan_status": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"vuln_next_scheduled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"vuln_next_sch_type": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"vuln_next_scan_on": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"vuln_next_start_hour": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"vuln_next_start_min": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"is_vuln_scanning": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"av_events_count": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"sb_events_count": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"fw_events_count": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"wf_events_count": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"vuln_events_count": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"sys_events_count": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"vuln_events_max_severity": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"forensics_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"conn_details": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"intf_name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"collapsed": schema.BoolAttribute{
							Computed: true,
							Optional: true,
						},
						"icon": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"connected_icon": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"connected_icon_color": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"active": schema.BoolAttribute{
							Computed: true,
							Optional: true,
						},
						"ssid": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"connections": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"ip_address": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"gateway": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"mac": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"gateway_mac": schema.StringAttribute{
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
			"hardware_details": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"model": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"vendor": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"cpu": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"ram": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"s_n": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"hdd": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"forensics": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"guid": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"status": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"verdict": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"report_url": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"completion_time": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"update_time": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"fsr_task_id": schema.Float64Attribute{
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"tags": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceEndpointsDetails) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoints_details"
}

func (r *datasourceEndpointsDetails) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointsDetailsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.DeviceId.ValueFloat64()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointsDetails(ctx, "read", diags))

	read_output, err := c.ReadEndpointsDetails(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointsDetails(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointsDetailsModel) refreshEndpointsDetails(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["host"]; ok {
		m.Host = parseStringValue(v)
	}

	if v, ok := o["alias"]; ok {
		m.Alias = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["ipAddr"]; ok {
		m.IpAddr = parseStringValue(v)
	}

	if v, ok := o["macAddr"]; ok {
		m.MacAddr = parseStringValue(v)
	}

	if v, ok := o["osVersion"]; ok {
		m.OsVersion = parseStringValue(v)
	}

	if v, ok := o["osServicePack"]; ok {
		m.OsServicePack = parseStringValue(v)
	}

	if v, ok := o["distinguishedName"]; ok {
		m.DistinguishedName = parseStringValue(v)
	}

	if v, ok := o["groupTag"]; ok {
		m.GroupTag = parseStringValue(v)
	}

	if v, ok := o["groupId"]; ok {
		m.GroupId = parseFloat64Value(v)
	}

	if v, ok := o["groupName"]; ok {
		m.GroupName = parseStringValue(v)
	}

	if v, ok := o["origGroupName"]; ok {
		m.OrigGroupName = parseStringValue(v)
	}

	if v, ok := o["origGroupId"]; ok {
		m.OrigGroupId = parseFloat64Value(v)
	}

	if v, ok := o["domainId"]; ok {
		m.DomainId = parseFloat64Value(v)
	}

	if v, ok := o["installerName"]; ok {
		m.InstallerName = parseStringValue(v)
	}

	if v, ok := o["deploymentState"]; ok {
		m.DeploymentState = parseFloat64Value(v)
	}

	if v, ok := o["scheduledInstallTime"]; ok {
		m.ScheduledInstallTime = parseStringValue(v)
	}

	if v, ok := o["installationState"]; ok {
		m.InstallationState = parseFloat64Value(v)
	}

	if v, ok := o["deploymentStateTime"]; ok {
		m.DeploymentStateTime = parseStringValue(v)
	}

	if v, ok := o["deploymentStateData"]; ok {
		m.DeploymentStateData = parseStringValue(v)
	}

	if v, ok := o["forticlientId"]; ok {
		m.ForticlientId = parseFloat64Value(v)
	}

	if v, ok := o["uid"]; ok {
		m.Uid = parseStringValue(v)
	}

	if v, ok := o["caps"]; ok {
		m.Caps = parseFloat64Value(v)
	}

	if v, ok := o["fctSn"]; ok {
		m.FctSn = parseStringValue(v)
	}

	if v, ok := o["fgtSn"]; ok {
		m.FgtSn = parseStringValue(v)
	}

	if v, ok := o["lastSeen"]; ok {
		m.LastSeen = parseFloat64Value(v)
	}

	if v, ok := o["deregister"]; ok {
		m.Deregister = parseFloat64Value(v)
	}

	if v, ok := o["runCmd"]; ok {
		m.RunCmd = parseFloat64Value(v)
	}

	if v, ok := o["quarantineMessage"]; ok {
		m.QuarantineMessage = parseStringValue(v)
	}

	if v, ok := o["isInstalled"]; ok {
		m.IsInstalled = parseBoolValue(v)
	}

	if v, ok := o["isManaged"]; ok {
		m.IsManaged = parseBoolValue(v)
	}

	if v, ok := o["isMigrating"]; ok {
		m.IsMigrating = parseBoolValue(v)
	}

	if v, ok := o["isEmsRegistered"]; ok {
		m.IsEmsRegistered = parseBoolValue(v)
	}

	if v, ok := o["isEmsOnline"]; ok {
		m.IsEmsOnline = parseBoolValue(v)
	}

	if v, ok := o["isEmsOnnet"]; ok {
		m.IsEmsOnnet = parseBoolValue(v)
	}

	if v, ok := o["isExcluded"]; ok {
		m.IsExcluded = parseBoolValue(v)
	}

	if v, ok := o["isQuarantined"]; ok {
		m.IsQuarantined = parseBoolValue(v)
	}

	if v, ok := o["quarantineAccessCode"]; ok {
		m.QuarantineAccessCode = parseStringValue(v)
	}

	if v, ok := o["fctVersion"]; ok {
		m.FctVersion = parseStringValue(v)
	}

	if v, ok := o["comparableFctVersion"]; ok {
		m.ComparableFctVersion = parseFloat64Value(v)
	}

	if v, ok := o["userDomain"]; ok {
		m.UserDomain = parseStringValue(v)
	}

	if v, ok := o["service"]; ok {
		m.Service = parseStringValue(v)
	}

	if v, ok := o["profileName"]; ok {
		m.ProfileName = parseStringValue(v)
	}

	if v, ok := o["ipListName"]; ok {
		m.IpListName = parseStringValue(v)
	}

	if v, ok := o["ztnaEnabled"]; ok {
		m.ZtnaEnabled = parseBoolValue(v)
	}

	if v, ok := o["ztnaSerial"]; ok {
		m.ZtnaSerial = parseStringValue(v)
	}

	if v, ok := o["avInstalled"]; ok {
		m.AvInstalled = parseBoolValue(v)
	}

	if v, ok := o["avEnabled"]; ok {
		m.AvEnabled = parseBoolValue(v)
	}

	if v, ok := o["avHidden"]; ok {
		m.AvHidden = parseBoolValue(v)
	}

	if v, ok := o["fwInstalled"]; ok {
		m.FwInstalled = parseBoolValue(v)
	}

	if v, ok := o["fwEnabled"]; ok {
		m.FwEnabled = parseBoolValue(v)
	}

	if v, ok := o["fwHidden"]; ok {
		m.FwHidden = parseBoolValue(v)
	}

	if v, ok := o["wfInstalled"]; ok {
		m.WfInstalled = parseBoolValue(v)
	}

	if v, ok := o["wfEnabled"]; ok {
		m.WfEnabled = parseBoolValue(v)
	}

	if v, ok := o["wfHidden"]; ok {
		m.WfHidden = parseBoolValue(v)
	}

	if v, ok := o["vpnInstalled"]; ok {
		m.VpnInstalled = parseBoolValue(v)
	}

	if v, ok := o["vpnEnabled"]; ok {
		m.VpnEnabled = parseBoolValue(v)
	}

	if v, ok := o["vpnHidden"]; ok {
		m.VpnHidden = parseBoolValue(v)
	}

	if v, ok := o["vulnInstalled"]; ok {
		m.VulnInstalled = parseBoolValue(v)
	}

	if v, ok := o["vulnEnabled"]; ok {
		m.VulnEnabled = parseBoolValue(v)
	}

	if v, ok := o["vulnHidden"]; ok {
		m.VulnHidden = parseBoolValue(v)
	}

	if v, ok := o["ssomaInstalled"]; ok {
		m.SsomaInstalled = parseBoolValue(v)
	}

	if v, ok := o["ssomaEnabled"]; ok {
		m.SsomaEnabled = parseBoolValue(v)
	}

	if v, ok := o["ssomaHidden"]; ok {
		m.SsomaHidden = parseBoolValue(v)
	}

	if v, ok := o["sbInstalled"]; ok {
		m.SbInstalled = parseBoolValue(v)
	}

	if v, ok := o["sbEnabled"]; ok {
		m.SbEnabled = parseBoolValue(v)
	}

	if v, ok := o["sbHidden"]; ok {
		m.SbHidden = parseBoolValue(v)
	}

	if v, ok := o["rsInstalled"]; ok {
		m.RsInstalled = parseBoolValue(v)
	}

	if v, ok := o["rsEnabled"]; ok {
		m.RsEnabled = parseBoolValue(v)
	}

	if v, ok := o["rsHidden"]; ok {
		m.RsHidden = parseBoolValue(v)
	}

	if v, ok := o["avLastScanType"]; ok {
		m.AvLastScanType = parseFloat64Value(v)
	}

	if v, ok := o["avLastScanDate"]; ok {
		m.AvLastScanDate = parseStringValue(v)
	}

	if v, ok := o["avLastFullScanDate"]; ok {
		m.AvLastFullScanDate = parseStringValue(v)
	}

	if v, ok := o["avLastCancelledScanType"]; ok {
		m.AvLastCancelledScanType = parseFloat64Value(v)
	}

	if v, ok := o["avLastCancelledScanDate"]; ok {
		m.AvLastCancelledScanDate = parseStringValue(v)
	}

	if v, ok := o["avScanScheduled"]; ok {
		m.AvScanScheduled = parseBoolValue(v)
	}

	if v, ok := o["avNextSchType"]; ok {
		m.AvNextSchType = parseFloat64Value(v)
	}

	if v, ok := o["avNextScanOn"]; ok {
		m.AvNextScanOn = parseFloat64Value(v)
	}

	if v, ok := o["avNextScanHour"]; ok {
		m.AvNextScanHour = parseFloat64Value(v)
	}

	if v, ok := o["avNextScanMin"]; ok {
		m.AvNextScanMin = parseFloat64Value(v)
	}

	if v, ok := o["avNextScanType"]; ok {
		m.AvNextScanType = parseFloat64Value(v)
	}

	if v, ok := o["isAvScanning"]; ok {
		m.IsAvScanning = parseBoolValue(v)
	}

	if v, ok := o["lastVulnScan"]; ok {
		m.LastVulnScan = parseFloat64Value(v)
	}

	if v, ok := o["vulnScanStatus"]; ok {
		m.VulnScanStatus = parseStringValue(v)
	}

	if v, ok := o["vulnNextScheduled"]; ok {
		m.VulnNextScheduled = parseBoolValue(v)
	}

	if v, ok := o["vulnNextSchType"]; ok {
		m.VulnNextSchType = parseFloat64Value(v)
	}

	if v, ok := o["vulnNextScanOn"]; ok {
		m.VulnNextScanOn = parseFloat64Value(v)
	}

	if v, ok := o["vulnNextStartHour"]; ok {
		m.VulnNextStartHour = parseFloat64Value(v)
	}

	if v, ok := o["vulnNextStartMin"]; ok {
		m.VulnNextStartMin = parseFloat64Value(v)
	}

	if v, ok := o["isVulnScanning"]; ok {
		m.IsVulnScanning = parseBoolValue(v)
	}

	if v, ok := o["avEventsCount"]; ok {
		m.AvEventsCount = parseFloat64Value(v)
	}

	if v, ok := o["sbEventsCount"]; ok {
		m.SbEventsCount = parseFloat64Value(v)
	}

	if v, ok := o["fwEventsCount"]; ok {
		m.FwEventsCount = parseFloat64Value(v)
	}

	if v, ok := o["wfEventsCount"]; ok {
		m.WfEventsCount = parseFloat64Value(v)
	}

	if v, ok := o["vulnEventsCount"]; ok {
		m.VulnEventsCount = parseFloat64Value(v)
	}

	if v, ok := o["sysEventsCount"]; ok {
		m.SysEventsCount = parseFloat64Value(v)
	}

	if v, ok := o["vulnEventsMaxSeverity"]; ok {
		m.VulnEventsMaxSeverity = parseFloat64Value(v)
	}

	if v, ok := o["connDetails"]; ok {
		m.ConnDetails = m.flattenEndpointsDetailsConnDetailsList(ctx, v, &diags)
	}

	if v, ok := o["hardwareDetails"]; ok {
		m.HardwareDetails = m.HardwareDetails.flattenEndpointsDetailsHardwareDetails(ctx, v, &diags)
	}

	if v, ok := o["forensics"]; ok {
		m.Forensics = m.Forensics.flattenEndpointsDetailsForensics(ctx, v, &diags)
	}

	if v, ok := o["forensicsEnabled"]; ok {
		m.ForensicsEnabled = parseBoolValue(v)
	}

	if v, ok := o["tags"]; ok {
		m.Tags = m.flattenEndpointsDetailsTagsList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointsDetailsModel) getURLObjectEndpointsDetails(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.DeviceId.IsNull() {
		result["deviceId"] = data.DeviceId.ValueFloat64()
	}

	return &result
}

type datasourceEndpointsDetailsConnDetailsModel struct {
	IntfName           types.String                                           `tfsdk:"intf_name"`
	Collapsed          types.Bool                                             `tfsdk:"collapsed"`
	Icon               types.String                                           `tfsdk:"icon"`
	ConnectedIcon      types.String                                           `tfsdk:"connected_icon"`
	ConnectedIconColor types.String                                           `tfsdk:"connected_icon_color"`
	Active             types.Bool                                             `tfsdk:"active"`
	Ssid               types.String                                           `tfsdk:"ssid"`
	Connections        *datasourceEndpointsDetailsConnDetailsConnectionsModel `tfsdk:"connections"`
}

type datasourceEndpointsDetailsConnDetailsConnectionsModel struct {
	IpAddress  types.String `tfsdk:"ip_address"`
	Gateway    types.String `tfsdk:"gateway"`
	Mac        types.String `tfsdk:"mac"`
	GatewayMac types.String `tfsdk:"gateway_mac"`
}

type datasourceEndpointsDetailsHardwareDetailsModel struct {
	Model  types.String `tfsdk:"model"`
	Vendor types.String `tfsdk:"vendor"`
	Cpu    types.String `tfsdk:"cpu"`
	Ram    types.String `tfsdk:"ram"`
	SN     types.String `tfsdk:"s_n"`
	Hdd    types.String `tfsdk:"hdd"`
}

type datasourceEndpointsDetailsForensicsModel struct {
	Guid           types.String  `tfsdk:"guid"`
	Status         types.String  `tfsdk:"status"`
	Verdict        types.String  `tfsdk:"verdict"`
	ReportUrl      types.String  `tfsdk:"report_url"`
	CompletionTime types.String  `tfsdk:"completion_time"`
	UpdateTime     types.String  `tfsdk:"update_time"`
	FsrTaskId      types.Float64 `tfsdk:"fsr_task_id"`
}

type datasourceEndpointsDetailsTagsModel struct {
	Id   types.Float64 `tfsdk:"id"`
	Name types.String  `tfsdk:"name"`
}

func (m *datasourceEndpointsDetailsConnDetailsModel) flattenEndpointsDetailsConnDetails(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsDetailsConnDetailsModel {
	if input == nil {
		return &datasourceEndpointsDetailsConnDetailsModel{}
	}
	if m == nil {
		m = &datasourceEndpointsDetailsConnDetailsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["intfName"]; ok {
		m.IntfName = parseStringValue(v)
	}

	if v, ok := o["collapsed"]; ok {
		m.Collapsed = parseBoolValue(v)
	}

	if v, ok := o["icon"]; ok {
		m.Icon = parseStringValue(v)
	}

	if v, ok := o["connectedIcon"]; ok {
		m.ConnectedIcon = parseStringValue(v)
	}

	if v, ok := o["connectedIconColor"]; ok {
		m.ConnectedIconColor = parseStringValue(v)
	}

	if v, ok := o["active"]; ok {
		m.Active = parseBoolValue(v)
	}

	if v, ok := o["ssid"]; ok {
		m.Ssid = parseStringValue(v)
	}

	if v, ok := o["connections"]; ok {
		m.Connections = m.Connections.flattenEndpointsDetailsConnDetailsConnections(ctx, v, diags)
	}

	return m
}

func (s *datasourceEndpointsDetailsModel) flattenEndpointsDetailsConnDetailsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointsDetailsConnDetailsModel {
	if o == nil {
		return []datasourceEndpointsDetailsConnDetailsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument conn_details is not type of []interface{}.", "")
		return []datasourceEndpointsDetailsConnDetailsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointsDetailsConnDetailsModel{}
	}

	values := make([]datasourceEndpointsDetailsConnDetailsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointsDetailsConnDetailsModel
		values[i] = *m.flattenEndpointsDetailsConnDetails(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointsDetailsConnDetailsConnectionsModel) flattenEndpointsDetailsConnDetailsConnections(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsDetailsConnDetailsConnectionsModel {
	if input == nil {
		return &datasourceEndpointsDetailsConnDetailsConnectionsModel{}
	}
	if m == nil {
		m = &datasourceEndpointsDetailsConnDetailsConnectionsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["ip address"]; ok {
		m.IpAddress = parseStringValue(v)
	}

	if v, ok := o["gateway"]; ok {
		m.Gateway = parseStringValue(v)
	}

	if v, ok := o["mac"]; ok {
		m.Mac = parseStringValue(v)
	}

	if v, ok := o["gateway mac"]; ok {
		m.GatewayMac = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointsDetailsHardwareDetailsModel) flattenEndpointsDetailsHardwareDetails(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsDetailsHardwareDetailsModel {
	if input == nil {
		return &datasourceEndpointsDetailsHardwareDetailsModel{}
	}
	if m == nil {
		m = &datasourceEndpointsDetailsHardwareDetailsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["model"]; ok {
		m.Model = parseStringValue(v)
	}

	if v, ok := o["vendor"]; ok {
		m.Vendor = parseStringValue(v)
	}

	if v, ok := o["cpu"]; ok {
		m.Cpu = parseStringValue(v)
	}

	if v, ok := o["ram"]; ok {
		m.Ram = parseStringValue(v)
	}

	if v, ok := o["s/n"]; ok {
		m.SN = parseStringValue(v)
	}

	if v, ok := o["hdd"]; ok {
		m.Hdd = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointsDetailsForensicsModel) flattenEndpointsDetailsForensics(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsDetailsForensicsModel {
	if input == nil {
		return &datasourceEndpointsDetailsForensicsModel{}
	}
	if m == nil {
		m = &datasourceEndpointsDetailsForensicsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["guid"]; ok {
		m.Guid = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["verdict"]; ok {
		m.Verdict = parseStringValue(v)
	}

	if v, ok := o["reportUrl"]; ok {
		m.ReportUrl = parseStringValue(v)
	}

	if v, ok := o["completionTime"]; ok {
		m.CompletionTime = parseStringValue(v)
	}

	if v, ok := o["updateTime"]; ok {
		m.UpdateTime = parseStringValue(v)
	}

	if v, ok := o["fsrTaskId"]; ok {
		m.FsrTaskId = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceEndpointsDetailsTagsModel) flattenEndpointsDetailsTags(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointsDetailsTagsModel {
	if input == nil {
		return &datasourceEndpointsDetailsTagsModel{}
	}
	if m == nil {
		m = &datasourceEndpointsDetailsTagsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointsDetailsModel) flattenEndpointsDetailsTagsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointsDetailsTagsModel {
	if o == nil {
		return []datasourceEndpointsDetailsTagsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument tags is not type of []interface{}.", "")
		return []datasourceEndpointsDetailsTagsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointsDetailsTagsModel{}
	}

	values := make([]datasourceEndpointsDetailsTagsModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointsDetailsTagsModel
		values[i] = *m.flattenEndpointsDetailsTags(ctx, ele, diags)
	}

	return values
}
