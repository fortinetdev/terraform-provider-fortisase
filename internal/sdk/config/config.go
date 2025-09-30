package config

import (
	"net/http"

	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/auth"
)

// Config provides configuration to a FORTISASE client instance
// It saves authentication information and a http connection
// for FORTISASE Client instance to create New connction to FORTISASE
// and Send data to FORTISASE,  etc. (needs to be extended later.)
type Config struct {
	Auth    *auth.Auth
	HTTPCon *http.Client
}
