resource "fortisase_security_profile_group" "example" {
  primary_key = "example_profile_name" # The name of the new profile group

  # Application Control With Inline-CASB
  application_control_profile = {
    status = "enable"
  }
}

# To configure this resource, please disable proxy configuration. "Network" -> "Proxy configuration"
resource "fortisase_security_application_control_profile" "application_control_profile" {
  primary_key = fortisase_security_profile_group.example.primary_key # The name of the existing profile group

  # [Application and Filter overrides]
  # Use filter to monitor/allow/block the applications
  controls = [
    # You can create multiple rules below to apply different actions to the applications
    {
      action = "monitor" # "monitor", "allow", "block"

      # Applications that match the following filters will have the specified action (“monitor”, “allow”, or “block”) applied.

      # [Filter: applications]
      ## Apply to all application
      applications = []
      ## Apply to specific application
      #   applications = [
      #     {
      #       primary_key = "Google.Ads"
      #       datasource  = "security/applications"
      #     }
      #   ]

      # [Filter: categories]
      ## Apply to all category
      categories = []
      ## Apply to specific category
      # categories = [
      #   {
      #     primary_key = "Game"
      #     datasource  = "security/application-categories"
      #   }
      # ]

      # [Filter: risk] Risk level with 0 being lowest and 4 being highest
      ## Apply to all risk
      risk = []
      ## Apply to specific risk
      # risk       = [{ id = 3 }, { id = 4 }]

      # [Filter: protocols]
      ## Apply to all protocol
      protocols = "all"
      ## Apply to specific protocol, different protocols are separated by spaces
      # protocols  = "HTTP HTTPS FTP"

      # [Filter: vendor]
      ## Apply to all vendor
      vendor = "all"
      ## Apply to specific vendor, different vendors are separated by spaces
      # vendor = "Google Meta"

      # [Filter: technology]
      ## Apply to all technology
      technology = "all"
      ## Apply to specific technology
      # technology = "Browser-Based" # 4 possible values: "Browser-Based", "Client-Server", "Network-Protocol", "Peer-to-Peer"

      # [Filter: behavior]
      ## Apply to all behavior
      behavior = "all"
      ## Apply to specific behavior, different behaviors are separated by spaces
      # behavior = "Botnet" # 5 possible values: "Botnet", "Cloud", "Evasive", "Excessive-Bandwidth", "Tunneling"

      # [Filter: popularity]
      ## Apply to all popularity
      popularity = "1 2 3 4 5"
      ## Apply to specific popularity, different popularity are separated by spaces
      # popularity = "1 2 3"
    }
  ]
  unknown_application_action = "monitor" # "block", "allow", "monitor"

  # [Network protocol enforcement]
  ## Disable Network protocol enforcement
  network_protocol_enforcement = "disable"
  ## Enable Network protocol enforcement
  #   network_protocol_enforcement = "enable"
  #   network_protocols = [{
  #     port     = 21
  #     action   = "monitor" # monitor or block
  #     services = ["ftp"]   # "dns", "ftp", "http", "https", "imap", "nntp", "pop3", "smtp", "snmp", "ssh", "telnet"
  #   }]

  # [Block applications detected on non-default ports]
  block_non_default_port_applications = "disable"
}
