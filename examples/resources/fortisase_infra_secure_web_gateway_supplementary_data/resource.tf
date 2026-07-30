resource "fortisase_infra_secure_web_gateway_supplementary_data" "example" {
  end_session_after_mins = 10
  primary_key            = "$sase-global"
  session_duration_hours = 1
}
