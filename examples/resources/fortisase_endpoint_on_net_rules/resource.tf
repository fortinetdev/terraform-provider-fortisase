# Note: For string arguments, use ; to separate multiple addresses.
resource "fortisase_endpoint_on_net_rules" "example" {
  primary_key = "example_name"

  # Specify one or more of the following rules, and delete the ones that are not applicable:

  # (1) Receives a successful HTTP(S) 200 OK response from a known server
  # HTTPS server IP addresses
  web_request_https = [{
    hostname = "example.com"
    ip       = "1.1.1.1"
  }]
  # HTTP server IP addresses
  web_request_http = "1.1.1.1;2.2.2.2"

  # (2) Connects with a known public IP
  public_ip = "1.1.1.1;2.2.2.2" # Known public (WAN) IP addresses, 0.0.0.0 or 0.0.0.0/32

  # (3) Is connected to a known DNS server
  dns_server_ip = "1.1.1.1;2.2.2.2" # Known server IP addresses

  # (4) Makes a successful query to a known DNS server
  # DNS query
  dns_request = [
    {
      hostname = "example1.com"
      ip       = "1.1.1.1"
    },
    {
      hostname = "example2.com"
      ip       = "2.2.2.2"
    }
  ]

  # (5) Is connected to a known DHCP server
  # Identify servers by IP/MAC addresses
  dhcp_server_ip  = "1.1.1.1"           # Known server IP addresses
  dhcp_server_mac = "AA-BB-CC-DD-EE-FF" # Known MAC addresses
  # Identify servers by DHCP option 224
  dhcp_server_code = "example_string" # Strings assigned to option 224

  # (6) Connects from a known local subnet
  local_ip    = "1.1.1.1"           # Known subnets, 0.0.0.0 or 0.0.0.0/0 or 0.0.0.0-1.1.1.1
  gateway_mac = "AA-BB-CC-DD-EE-FF" # Known gateway MAC addresses

  # (7) Can ping a known server
  ping_server = "1.1.1.1" # Known server IP addresses
}
