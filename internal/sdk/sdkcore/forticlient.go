package forticlient

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/auth"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/config"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/request"
	// "strconv"
)

// MultValue describes the nested structure in the results
type MultValue struct {
	Name string `json:"name"`
}

// MultValues describes the nested structure in the results
type MultValues []MultValue

// FortiSDKClient describes the global FortiSASE plugin client instance
type FortiSDKClient struct {
	Config  config.Config
	Retries int
}

// ExtractString extracts strings from result and put them into a string array,
// and return the string array
func ExtractString(members []MultValue) []string {
	vs := make([]string, 0, len(members))
	for _, v := range members {
		c := v.Name
		vs = append(vs, c)
	}
	return vs
}

// EscapeURLString escapes the string so it can be safely placed inside a URL query
func EscapeURLString(v string) string { // doesn't support "<>()"'#"
	return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

func escapeURLString(v string) string { // doesn't support "<>()"'#"
	return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

// NewClient initializes a new global plugin client
// It returns the created client object
func NewClient(auth *auth.Auth, client *http.Client) (*FortiSDKClient, error) {
	c := &FortiSDKClient{}

	c.Config.Auth = auth
	c.Config.HTTPCon = client
	c.GenToken()

	return c, nil
}

// NewRequest creates the request to FortiSASE for the client
// and return it to the client
func (c *FortiSDKClient) NewRequest(method string, path string, params interface{}, data *bytes.Buffer) *request.Request {
	return request.New(c.Config, method, path, params, data)
}

// GenToken generate access tokan and refresh token
// If errors are encountered, it returns the error.
func (c *FortiSDKClient) GenToken() error {
	// var err error
	if c.Config.Auth.AccessToken != "" {
		// todo: need check the validation of the access token
	} else if c.Config.Auth.RefreshToken != "" {
		// todo: generate access token by refresh token
	} else {
		req := c.NewRequest("POST", "https://customerapiauth.fortinet.com/api/v1/oauth/token/", nil, nil) // todo: may could move the url into the request.GenToken()
		access_token, refresh_token, err := req.GenToken()
		if err == nil {
			c.Config.Auth.AccessToken = access_token
			c.Config.Auth.RefreshToken = refresh_token
		}
		return err
	}
	return nil
}

// CheckUP checks whether username and password is valid
// If errors are encountered, it returns the error.
func (c *FortiSDKClient) CheckUP() error {
	// todo
	// var err error

	// req := c.NewRequest("GET", "/api/v2/monitor/system/status", nil, nil)
	// err = req.CheckValid()
	// if err != nil {
	// 	if c.Config.Auth.Token == "" {
	// 		err = fmt.Errorf("Error using Username/Password to login: %v", err)
	// 	} else {
	// 		err = fmt.Errorf("Error using Token to login: %v", err)
	// 	}
	// }
	return nil
}

func fortiAPIHttpStatus404Checking(result map[string]interface{}) (b404 bool) {
	b404 = false

	if result != nil {
		if result["code"] != nil && result["code"] == 404 {
			b404 = true
			return
		}
	}

	return
}

func fortiAPIErrorFormat(result map[string]interface{}, body string) (code float64, err error) {
	code = -100
	if result != nil {
		if code, ok := result["code"].(float64); ok {
			// 200	OK: Request returns successful
			if code == 200.0 {
				return code, nil
			} else if code == 400.0 {
				err = fmt.Errorf("Bad Request - Request cannot be processed by the API (%v)", result)
			} else if code == 401.0 {
				err = fmt.Errorf("Not Authorized - Request without successful login session (%.0f)", code)
			} else if code == 403.0 {
				err = fmt.Errorf("Forbidden - Request is missing CSRF token or administrator is missing access profile permissions (%.0f)", code)
			} else if code == 404.0 {
				err = fmt.Errorf("Resource Not Found - Unable to find the specified resource (%.0f)", code)
			} else if code == 405.0 {
				err = fmt.Errorf("Method Not Allowed - Specified HTTP method is not allowed for this resource (%.0f)", code)
			} else if code == 413.0 {
				err = fmt.Errorf("Request Entity Too Large - Request cannot be processed due to large entity (%.0f)", code)
			} else if code == 424.0 {
				err = fmt.Errorf("Failed Dependency - Fail dependency can be duplicate resource, missing required parameter, missing required attribute, invalid attribute value (%.0f)", code)
			} else if code == 429.0 {
				err = fmt.Errorf("Access temporarily blocked - Maximum failed authentications reached. The offended source is temporarily blocked for certain amount of time (%.0f)", code)
			} else if code == 500.0 {
				err = fmt.Errorf("Internal Server Error - Internal error when processing the request (%.0f)", code)
			} else {
				err = fmt.Errorf("Unknow Error (%.0f)", code)
			}
			return code, err
		}
		err = fmt.Errorf("\n%v", body)
		return code, err
	}

	// Authorization Required, etc. | Attention: scalable here
	err = fmt.Errorf("\n%v", body)
	return code, err
}

//Build input data by sdk
