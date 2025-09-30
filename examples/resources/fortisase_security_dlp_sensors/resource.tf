resource "fortisase_security_dlp_dictionaries" "dlp_dictionary" {
  primary_key         = "example_dlp_dictionary_name"
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

resource "fortisase_security_dlp_sensors" "dlp_sensor" {
  primary_key                     = "example_dlp_senso_name"
  entry_matches_to_trigger_sensor = "all"
  sensor_dictionaries = [
    {
      dictionary_id = 1
      dictionary = {
        primary_key = fortisase_security_dlp_dictionaries.dlp_dictionary.primary_key
        datasource  = "security/dlp-dictionaries"
      }
      dictionary_matches_to_consider_risk = 255
      status                              = "enable"
    }
  ]
}
