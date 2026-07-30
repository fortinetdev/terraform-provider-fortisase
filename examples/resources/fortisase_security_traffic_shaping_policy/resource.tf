# GUI: Security -> Bandwidth control -> Bandwidth control policy
resource "fortisase_security_traffic_shaping_policy" "bandwidth_control_policy" {
  primary_key = "bandwidth_control_policy"
  status      = "enable"
  comment     = "Bandwidth Control Policy Example"

  # GUI: If traffic matches
  # [Traffic type]
  ### Option A: Internet access
  traffic_direction = "outbound"


  # [Source scope]
  ### Option A: All
  scope = "all"

  ### Option B: VPN users
  # scope = "vpn-user"


  # [User]
  ### Option A: All Users
  users = []

  ### Option B: Specify users
  # users = [
  #   {
  #     primary_key = "example_user@example.com"
  #     datasource  = "auth/users"
  #   }
  # ]


  # [Destination]
  ### Option A: All destinations
  destination_scope = "all"
  destinations      = []

  ### Option B: Specify
  # destination_scope = "specify"
  # destinations = [
  #   {
  #     primary_key = "fortisandbox_global_1"
  #     datasource  = "network/hosts"
  #   }
  # ]

  # [Schedule]
  ###  Option A: Disable schedule
  schedule = {}

  ###  Option B: Enable schedule
  # schedule = {
  #   primary_key = "always"
  #   datasource  = "security/recurring-schedules"
  # }


  # [Service]
  services = [
    {
      primary_key = "ALL"
      datasource  = "security/services"
    }
  ]


  # [Application]
  ### Option A: No application
  applications = []

  ### Option B: Specify applications
  # applications = [
  #   {
  #     primary_key = "Alipay"
  #     datasource  = "security/applications"
  #   }
  # ]


  # [URL category]
  ### Option A: No URL category
  url_categories = []

  ### Option B: Specify URL category
  # url_categories = [
  #   {
  #     id = 98
  #   }
  # ]

  sources        = []
  app_categories = []


  # GUI: Then apply the following profile
  # [Shared profile]
  ### Option A: No shared profile
  # traffic_shaper = {}

  ### Option B: Specify shared profile
  traffic_shaper = {
    primary_key = "low-priority"
    datasource  = "security/traffic-shapers"
  }


  # [Reverse direction of shared profile]
  ### Option A: No reverse shared profile
  # traffic_shaper_reverse = {}

  ### Option B: Specify reverse shared profile
  traffic_shaper_reverse = {
    primary_key = "low-priority"
    datasource  = "security/traffic-shapers"
  }


  # [Per IP profile]
  ### Option A: No Per IP profile
  per_ip_shaper = {}

  ### Option B: Specify Per IP profile
  # per_ip_shaper = {
  #   primary_key = "your_per_ip_name"
  #   datasource  = "security/per-ip-traffic-shapers"
  # }
}
