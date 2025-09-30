package auth

import (
	"os"
)

// Auth describes the authentication information for FortiSASE
type Auth struct {
	AccessToken  string
	RefreshToken string
	Username     string
	Password     string
}

// NewAuth inits Auth object with the given metadata
func NewAuth(username, password, access_token, refresh_token string) *Auth {
	return &Auth{
		Username:     username,
		Password:     password,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}
}

// GetEnvUsername gets username from OS environment
// It returns the username
func (m *Auth) GetEnvUsername() (string, error) {
	h := os.Getenv("FORTISASE_ACCESS_USERNAME")

	m.Username = h

	return h, nil
}

// GetEnvPassword gets password from OS environment
// It returns the hostname
func (m *Auth) GetEnvPassword() (string, error) {
	h := os.Getenv("FORTISASE_IAM_PASSWORD")

	m.Password = h

	return h, nil
}

// GetEnvAccessToken gets AccessToken from OS environment
// It returns the hostname
func (m *Auth) GetEnvAccessToken() (string, error) {
	h := os.Getenv("FORTISASE_ACCESS_TOKEN")

	m.AccessToken = h

	return h, nil
}

// GetEnvRefreshToken gets RefreshToken from OS environment
// It returns the hostname
func (m *Auth) GetEnvRefreshToken() (string, error) {
	h := os.Getenv("FORTISASE_REFRESH_TOKEN")

	m.RefreshToken = h

	return h, nil
}
