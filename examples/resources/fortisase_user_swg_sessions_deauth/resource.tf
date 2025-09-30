resource "fortisase_user_swg_sessions_deauth" "user_swg_sessions_deauth" {
  usernames   = ["employee@company.com", "user@company.com"]
  session_ids = ["Ottawa - Canada-demo@demo.com-0", "San Jose - USA-demo@demo.com-1"]
}
