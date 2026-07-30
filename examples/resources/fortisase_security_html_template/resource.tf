# GUI: System -> HTML templates -> Templates
resource "fortisase_security_html_template" "example" {
  primary_key = "your_html_template"
  buffer      = "<html><body><h1>FortiSASE HTML Template</h1><p>tf test</p></body></html>"
}
