resource "fortisase_security_cert_local_certs" "local_cert" {
  format           = "regular"
  cert_name        = "local_cert_name"
  password         = "your_password"
  file_content     = base64encode(file("./path/to/cert.pem"))
  key_file_content = base64encode(file("./path/to/key.pem"))
}
