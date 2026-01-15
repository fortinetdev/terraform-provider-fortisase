resource "fortisase_network_dns_rules" "example" {
  # [Primary key]
  primary_key = "UniqueID8324" # Unique String to identify the rule

  # [Primary DNS Server]
  primary_dns = "2.2.2.2" # IPv4 or domain name

  # [Secondary DNS Server]
  secondary_dns = "3.3.3.3" # IPv4 or domain name

  # [Domains]
  domains = ["example1.com", "example2.com"] # List of IPv4 or domain name. At least one domain is required.

  # [Access type]
  ## Public
  for_private = false
  ## Private, If this DNS rule is for SPA traffic, private should be selected.
  # for_private = true

  # [PoP DNS override]
  ## Disable PoP DNS override
  pop_dns_override = {}
  ## Enable PoP DNS override
  # pop_dns_override = {
  #   "region13" = {
  #     pop           = "region13"
  #     primary_dns   = "5.5.5.6"
  #     secondary_dns = "6.6.6.6"
  #   }
  # }
}
