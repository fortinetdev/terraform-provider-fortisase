resource "fortisase_security_cert_remote_certs" "remote_cert" {
  cert_name    = "remote_cert_name"
  file_content = base64encode(file("./path/to/cert.pem"))
}
