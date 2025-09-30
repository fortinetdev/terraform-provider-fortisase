resource "fortisase_security_cert_remote_ca_certs" "remote_ca_cert" {
  cert_name    = "remote_ca_cert_name"
  file_content = base64encode(file("./path/to/certificate.crt"))
}
