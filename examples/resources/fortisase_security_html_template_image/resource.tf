# GUI: System -> HTML templates -> Images
resource "fortisase_security_html_template_image" "example" {
  primary_key = "your_image"
  image_type  = "png"
  # 1x1 transparent PNG
  image_base64 = "iVBORw0KGgoAAAANSUhEUgAAAAsdEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=="
}
