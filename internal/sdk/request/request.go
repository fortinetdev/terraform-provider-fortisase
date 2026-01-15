package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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
	u := "https://portal.prod.fortisase.com"
	u += r.Path

	return u
}

// Login FortiSASE using username and password in token mode, and return Cookies.
// If errors are encountered, it returns the error.
func (r *Request) GenToken() (string, string, error) {
	// todo
	// generate access token and refresh token

	var err error
	var access_token string
	var refresh_token string

	data := make(map[string]interface{})
	data["username"] = r.Config.Auth.Username
	data["password"] = r.Config.Auth.Password
	data["client_id"] = "FortiSASE"
	data["grant_type"] = "password"

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Printf("[ERROR] Encoding body data failed.")
		return access_token, refresh_token, err
	}

	bodyBytes := bytes.NewBuffer(locJSON)

	req, _ := http.NewRequest("POST", "", bodyBytes)
	req.Header.Set("Content-Type", "application/json")
	req.URL, err = url.Parse("https://customerapiauth.fortinet.com/api/v1/oauth/token/")
	if err != nil {
		err = fmt.Errorf("Could not parse URL: %s", err)
		return access_token, refresh_token, err
	}

	rsp, err := r.Config.HTTPCon.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "x509: ") {
			err = fmt.Errorf("HTTP request error: %v", err)
			return access_token, refresh_token, err
		}
	}

	if rsp == nil {
		err = fmt.Errorf("Host is unreachable. HTTP response is nil.")
		return access_token, refresh_token, err
	}

	body, err := io.ReadAll(rsp.Body)
	rsp.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body, %s", err)
		return access_token, refresh_token, err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if at, ok := result["access_token"]; ok {
		access_token = at.(string)
		refresh_token = result["refresh_token"].(string)
	} else {
		err = fmt.Errorf("Login failed: %S.", result["status_message"])
	}

	return access_token, refresh_token, err
}

// Logout current token based authentication.
// If errors are encountered, it returns the error.
func (r *Request) LogoutToken(token string) error {
	// logout the token
	return nil
}
