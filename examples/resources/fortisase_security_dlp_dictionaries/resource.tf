resource "fortisase_security_dlp_dictionaries" "dlp_dictionary" {
  primary_key         = "dlp_dictionary"
  dictionary_type     = "sensor"
  entries_to_evaluate = "all"
  entries = [
    {
      dlp_data_type = {
        primary_key = "regex"
        datasource  = "security/dlp-data-types"
      }
      pattern = "string1"
      status  = "enable"
      repeat  = "enable"
    }
  ]
}
