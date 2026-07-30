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
  controls = [
    # You can create multiple rules below to apply different actions to the applications
    # Example 1: Monitor "GAME" and "Social.Media" categories
    {
      action = "monitor" # "monitor", "allow", "block"
      categories = [
        {
          primary_key = "Game"
          datasource  = "security/application-categories"
        },
        {
          primary_key = "Social.Media"
          datasource  = "security/application-categories"
        }
      ]
      applications = []
      risk         = []
      popularity   = [1, 2, 3, 4, 5]
    },
    # Example 2: Block specific applications
    {
      action     = "block" # "monitor", "allow", "block"
      categories = []
      applications = [
        {
          primary_key = "2ch"
          datasource  = "security/applications"
        },
        {
          primary_key = "Facebook"
          datasource  = "security/applications"
        }
      ]
      risk       = []
      popularity = [1, 2, 3, 4, 5]
    },
    # Example 3: Block all risk 4 applications
    {
      action       = "block" # "monitor", "allow", "block"
      categories   = []
      applications = []
      risk         = [4]
      popularity   = [1, 2, 3, 4, 5]
    }
  ]
  unknown_application_action = "allow" # "block", "allow", "monitor"

  # [Network protocol enforcement]
  ## Disable Network protocol enforcement
  # network_protocol_enforcement = "disable"
  ## Enable Network protocol enforcement
  network_protocol_enforcement = "enable"
  network_protocols = [{
    port     = 21
    action   = "monitor" # monitor or block
    services = ["ftp"]   # "dns", "ftp", "http", "https", "imap", "nntp", "pop3", "smtp", "snmp", "ssh", "telnet"
  }]

  # [Block applications detected on non-default ports]
  block_non_default_port_applications = "disable"
}