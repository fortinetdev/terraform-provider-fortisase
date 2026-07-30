# GUI: Endpoint management -> Configuration -> Profiles
resource "fortisase_endpoint_profile" "endpoint_profile" {
  primary_key = "example_endpoint_profile"
  enabled     = true
}

resource "fortisase_endpoint_connection_profile" "connection_profile" {
  primary_key = fortisase_endpoint_profile.endpoint_profile.primary_key

  # [Endpoint connects to FortiSASE Cloud Security]
  connect_to_fortisase = "automatically" # "automatically" or "manually"

  # [Show option to disconnect from security PoP]
  show_disconnect_btn = "enable"

  secure_internet_access = {
    # [Authenticate with SSO]
    authenticate_with_sso       = "enable"
    external_browser_saml_login = "disable" # "enable" or "disable". Required if authenticate_with_sso is "enable"
    allow_fido_auth             = "disable"

    # [Failover sequence]
    # failover_sequence = ["newdomain.com"]

    # [Run posture check before initiating FortiSASE Cloud Security tunnel]
    ## Optiona A: Enable posture check
    # posture_check = {
    #   action               = "prohibit"
    #   tag                  = "your_posture_check_name"
    #   check_failed_message = "Your Comment"
    # }
    ## Optiona B: Remove posture check
    # posture_check = {
    #   action               = "allow" # must be "allow"
    #   tag                  = "" # must be ""
    #   check_failed_message = "" # must be ""
    # }

    # [Allow local LAN access]
    enable_local_lan = "enable"
  }

  # [On/off-net detection]
  # on_fabric_rule_set = {
  #   datasource  = "endpoint/on-net-rules"
  #   primary_key = "example_name"
  # }

  # [IPsec tunnel MTU size]
  # mtu_size = 1280
}
