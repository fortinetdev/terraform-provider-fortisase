resource "fortisase_security_cert_local_ca_certs" "local_ca_cert" {
  format           = "regular"
  cert_name        = "local_ca_cert"
  password         = "your_password"
  file_content     = base64encode(file("./path/to/ca_cert.crt"))
  key_file_content = base64encode(file("./path/to/private.key"))
}
