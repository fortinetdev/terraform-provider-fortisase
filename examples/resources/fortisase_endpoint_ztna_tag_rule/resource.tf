# This resource is only available for EMS 7.4
resource "fortisase_endpoint_ztna_tag_rule" "your_rule" {
  primary_key = "your_rule"
  status      = "enable"
  description = "Your description"
  rules = [
    {
      content = "AV Software is installed and running"
      id      = 1
      negated = false
      os      = "windows"
      type    = "anti-virus"
    },
    {
      content = "FortiClient installed and Telemetry connected to EMS"
      id      = 2
      os      = "windows"
      type    = "ems-management"
    },
    {
      content = "AV Software is installed and running"
      id      = 3
      negated = false
      os      = "macos"
      type    = "anti-virus"
    },
    {
      content = "FortiClient installed and Telemetry connected to EMS"
      id      = 4
      os      = "macos"
      type    = "ems-management"
    }
  ]
  logic = {
    windows = jsonencode({
      op = "and"
      rules = [
        {
          id = 1
        },
        {
          id = 2
        }
      ]
    })
    macos = jsonencode({
      op = "and"
      rules = [
        {
          id = 3
        },
        {
          id = 4
        }
      ]
    })
    # linux = null
    # ios = null
    # android = null
  }
}
