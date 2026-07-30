package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/config"
)

// Request describes the request to FortiSASE service
type Request struct {
	Config       config.Config
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       interface{}
	Data         *bytes.Buffer
}

const (
	defaultFortiSASEAPIURL  = "https://portal.prod.fortisase.com"
	defaultFortiSASEAuthURL = "https://customerapiauth.fortinet.com"
)

// FortiSASEAPIURL returns the FortiSASE API base URL from the environment.
func FortiSASEAPIURL() string {
	return envURL("FORTISASE_API_URL", defaultFortiSASEAPIURL)
}

// FortiSASEAuthURL returns the FortiSASE authentication base URL from the environment.
func FortiSASEAuthURL() string {
	return envURL("FORTISASE_AUTH_URL", defaultFortiSASEAuthURL)
}

func envURL(name, defaultURL string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return defaultURL
	}
	return strings.TrimRight(value, "/")
}

// New creates reqeust object with http method, path, params and data,
// It will save the http request, path, etc. for the next operations
// such as sending data, getting response, etc.
// It returns the created request object to the gobal plugin client.
func New(c config.Config, method string, path string, params interface{}, data *bytes.Buffer) *Request {
	var h *http.Request
	log.Printf("[%v] [REQUEST] [%v] %v, %v", method, path, params, data)

	if data == nil {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
	}

	r := &Request{
		Config:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r
}

// Build Request header

// Build Request Sign/Login Info

// Send request data to FortiSASE.
// If errors are encountered, it returns the error.
func (r *Request) Send() error {
	var err error
	retries := 15
	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	r.HTTPRequest.Header.Set("accept", "application/json")
	access_token := r.Config.Auth.AccessToken
	r.HTTPRequest.Header.Set("Authorization", "Bearer "+access_token)
	u := r.buildURL()

	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", errdo.Error())
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", errdo.Error())
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", errdo.Error(), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

func (r *Request) buildURL() string {
	u := FortiSASEAPIURL()
	u += r.Path

	return u
}

// GenToken logs in to FortiSASE with a username and password and returns tokens.
func GenToken(cfg config.Config) (string, string, error) {
	data := map[string]string{
		"username":   cfg.Auth.Username,
		"password":   cfg.Auth.Password,
		"client_id":  "FortiSASE",
		"grant_type": "password",
	}

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Printf("[ERROR] Encoding body data failed.")
		return "", "", err
	}

	bodyBytes := bytes.NewBuffer(locJSON)
	tokenURL := fmt.Sprintf("%s/api/v1/oauth/token/", FortiSASEAuthURL())
	req, err := http.NewRequest(http.MethodPost, tokenURL, bodyBytes)
	if err != nil {
		return "", "", fmt.Errorf("could not create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rsp, err := cfg.HTTPCon.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("token request failed: %w", err)
	}

	if rsp == nil {
		return "", "", fmt.Errorf("host is unreachable: HTTP response is nil")
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", "", fmt.Errorf("cannot read token response body: %w", err)
	}

	var result struct {
		AccessToken   string `json:"access_token"`
		RefreshToken  string `json:"refresh_token"`
		StatusMessage string `json:"status_message"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", fmt.Errorf("cannot decode token response: %w", err)
	}

	if result.AccessToken == "" || result.RefreshToken == "" {
		if result.StatusMessage == "" {
			result.StatusMessage = "token response does not contain access_token and refresh_token"
		}
		return "", "", fmt.Errorf("login failed: %s", result.StatusMessage)
	}

	return result.AccessToken, result.RefreshToken, nil
}

// Logout current token based authentication.
// If errors are encountered, it returns the error.
func (r *Request) LogoutToken(token string) error {
	// logout the token
	return nil
}

// Get EMS version of FortiSASE.
// If errors are encountered, it returns the error.
func (r *Request) GetEMSVersion() (string, error) {
	var err error
	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	r.HTTPRequest.Header.Set("accept", "application/json")
	access_token := r.Config.Auth.AccessToken
	r.HTTPRequest.Header.Set("Authorization", "Bearer "+access_token)
	r.HTTPRequest.URL, err = url.Parse(r.Path)
	if err != nil {
		return "", err
	}

	rsp, err := r.Config.HTTPCon.Do(r.HTTPRequest)
	if err != nil {
		if strings.Contains(err.Error(), "x509: ") {
			err = fmt.Errorf("HTTP request error: %v", err)
			return "", err
		}
	}

	if rsp == nil {
		err = fmt.Errorf("Host is unreachable. HTTP response is nil.")
		return "", err
	}

	body, err := io.ReadAll(rsp.Body)
	rsp.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("Cannot get response body, %s", err)
		return "", err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if data, ok := result["data"].(map[string]interface{}); ok {
		if emsv, ok := data["emsVersion"].(string); ok {
			return emsv, nil
		} else {
			err = fmt.Errorf("Response data do not contains emsVersion: %v", result["data"])
		}
	} else {
		err = fmt.Errorf("Response do not contains data: %v", result)
	}

	return "", err
}
