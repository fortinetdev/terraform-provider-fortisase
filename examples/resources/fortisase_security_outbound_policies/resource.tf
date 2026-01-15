# GUI: Security > Policies > Secure internet access
resource "fortisase_security_outbound_policies" "example" {
  primary_key = "example_outbound_policy"

  # [Source Scope]
  ### All ###
  scope = "all"
  ### All Agent Devices ###
  # scope = "vpn-user"
  ### All Edge Devices ###
  # scope = "thin-edge"
  ### Specify ###
  # scope = "specify"
  # sources = [
  #   {
  #     primary_key = "existing_host_example"
  #     datasource  = "network/hosts"
  #   }
  # ]


  # [Security Posture Tag]
  # sources = [
  #   {
  #     primary_key = "tag_name"
  #     datasource  = "endpoint/ztna-tags"
  #   }
  # ]


  # [User]
  ### All Users ###
  users                 = []
  captive_portal_exempt = false
  ### Specify ###
  # users = [
  #   {
  #     primary_key = "user@example.com"
  #     datasource = "auth/users"
  #   },
  #   {
  #     primary_key = "user_group_name"
  #     datasource = "auth/user-groups"
  #   }
  # ]
  # captive_portal_exempt = false
  ### Captive Portal Exempty ###
  # Note: Available if scope is "thin-edge" or "specify".
  # users = []
  # captive_portal_exempt = true


  # [Destination]
  ### All Internet Traffic ###
  destinations = [
    {
      primary_key = "all"
      datasource  = "network/hosts"
    }
  ]
  ### Specify ###
  # destinations = [
  #   {
  #     primary_key = "Microsoft-Azure.AD"
  #     datasource  = "network/internet-services"
  #   }
  # ]


  # [Services]
  services = [
    {
      primary_key = "ALL",
      datasource  = "security/services"
    }
  ]


  # [Profile Group]
  ### Internet Access ###
  profile_group = {
    group = {
      primary_key = "outbound"
      datasource  = "security/profile-groups"
    }
    # Basic certificate inspection override
    force_cert_inspection = false
  }
  ### Specify ###
  # profile_group = {
  #   group = {
  #     # If your profile group name is "example", the primary key should be "outbound-example"
  #     primary_key = "outbound-{your profile group name}"
  #     datasource  = "security/profile-groups"
  #   }
  #   # Basic certificate inspection override
  #   force_cert_inspection = false
  # }


  # [Schedule]
  schedule = {
    primary_key = "always"
    datasource  = "security/recurring-schedules"
  }


  # [Action]
  action = "accept" # "accept" or "deny"


  # [Status]
  enabled = true


  # [Logging Options]
  ### All Sessions ###
  # log_traffic = "all"
  ### Disable ###
  log_traffic = "disable"
  ### Security Events ###
  # Note: Available if action is "accept" and scope is "thin-edge" or "specify"
  # log_traffic = "utm"


  # [Comments]
  comments = "Your comments"
}
