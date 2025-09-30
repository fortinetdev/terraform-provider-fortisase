resource "fortisase_endpoint_ztna_tags" "ztna_tag" {
  name        = "Compliant"
  primary_key = "Compliant"
}

resource "fortisase_endpoint_ztna_rules" "ztna_rule" {
  primary_key = "Compliant Endpoints"
  status      = "enable"
  comments    = ""
  rules = [
    {
      content = "AV Software is installed and running"
      id      = 1
      negated = false
      os      = "windows"
      type    = "anti-virus"
    },
    {
      content = "AV Signature is up-to-date"
      id      = 2
      negated = false
      os      = "windows"
      type    = "anti-virus"
    },
    {
      content                    = "Windows 11"
      enable_latest_update_check = false
      id                         = 3
      negated                    = false
      os                         = "windows"
      type                       = "os-version"
    },
    {
      content = "High or higher"
      id      = 4
      negated = true
      os      = "windows"
      type    = "vulnerable-devices"
    },
    {
      content = "Windows Firewall is enabled"
      id      = 5
      negated = false
      os      = "windows"
      type    = "windows-security"
    },
  ]
  logic = {
    windows = jsonencode(
      {
        op = "and"
        rules = [
          {
            op = "and"
            rules = [
              {
                id = 1
              },
              {
                id = 2
              },
            ]
          },
          {
            id = 3
          },
          {
            id = 4
          },
          {
            id = 5
          },
        ]
      }
    )
    macos   = ""
    linux   = ""
    ios     = ""
    android = ""
  }
  tag = {
    datasource  = "endpoint/ztna-tags"
    primary_key = fortisase_endpoint_ztna_tags.ztna_tag.primary_key
  }
}