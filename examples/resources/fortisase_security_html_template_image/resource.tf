# GUI: System -> HTML templates -> Images
resource "fortisase_security_html_template_image" "example" {
  primary_key  = "your_image"
  image_type   = "png"
  image_base64 = filebase64("./path/to/fortinet.png")
}
