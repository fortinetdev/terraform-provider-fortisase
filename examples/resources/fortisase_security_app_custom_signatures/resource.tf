resource "fortisase_security_app_custom_signatures" "example" {
  primary_key = "example_name"
  # It is recommended to use ' instead of " in the signature string.
  signature = "F-SBID( --attack_id 6483; --name 'Windows.NT.6.1.Web.Surfing'; --default_action drop_session; --service HTTP; --protocol tcp; --app_cat 25; --flow from_client; --pattern !'FCT'; --pattern 'Windows NT 6.1'; --no_case; --context header; --weight 40; )"
  tag       = "example_name" # should be same as primary_key
}
