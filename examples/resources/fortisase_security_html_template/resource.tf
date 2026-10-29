# GUI: System -> HTML templates -> Templates
# This resource currently only supports updating existing templates.
# Available primary_key values and corresponding GUI templates:
#   "fsase-appblk-html": Application control block page
#   "fsase-banned-word-html": Banned word block page
#   "auth-disclaimer-page-1": Captive portal disclaimer
#   "auth-login-failed-page": Captive portal login failed page
#   "auth-login-page": Captive portal login page
#   "auth-reject-page": Captive portal reject page
#   "auth-saml-page": Captive portal SAML login page
#   "fsase-dlp-html": Data loss prevention block page
#   "fsase-ftgd-block": FortiGuard block page
#   "fsase-ipsblk-html": Intrusion prevention block page
#   "fsase-url-block": URL block page
#   "fsase-videofilter-block": Video block page
#   "fsase-virus-html": Virus block page
#   "fsase-client-virus-html": Virus upload block page
#   "fsase-videofilter-block-text": YouTube embedded custom block message
#   "Invitation Email": Invitation Email
#   "fsase-ftgd-warning": FortiGuard warning page
resource "fortisase_security_html_template" "example" {
  primary_key = "fsase-videofilter-block-text"
  buffer      = "Video Access Blocked by FortiSASE. %%VIDEOFILTER_DESC%%\n"
}
