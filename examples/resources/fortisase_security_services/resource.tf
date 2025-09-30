resource "fortisase_security_services" "service" {
  primary_key = "service_name"
  proxy       = false
  category    = "Email"
  protocol    = "TCP/UDP/SCTP"
  tcp_portrange = [
    {
      destination = {
        low = 25
      }
    }
  ]
}
