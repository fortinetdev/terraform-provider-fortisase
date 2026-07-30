package forticlient

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/auth"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/config"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/request"
)

// MultValue describes the nested structure in the results
type MultValue struct {
	Name string `json:"name"`
}

// MultValues describes the nested structure in the results
type MultValues []MultValue

// FortiSDKClient describes the global FortiSASE plugin client instance
type FortiSDKClient struct {
	Config    config.Config
	FSSStatus *FSSStatus
	Retries   int
}

// FortiSASE system information
type FSSStatus struct {
	EMSVersion string
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
	err := c.GenToken()
	if err != nil {
		err = fmt.Errorf("Generate token failed: %v", err)
		return c, err
	}

	err = c.GetFSSStatus()
	if err != nil {
		log.Printf("Could not get EMS version: %v", err)
	}

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
		access_token, refresh_token, err := request.GenToken(c.Config)
		if err == nil {
			c.Config.Auth.AccessToken = access_token
			c.Config.Auth.RefreshToken = refresh_token
		}
		return err
	}
	return nil
}

// Get the FortiSASE status
// If errors are encountered, it returns the error.
func (c *FortiSDKClient) GetFSSStatus() error {
	fssstatus := &FSSStatus{}

	// Get EMS version of FortiSASE
	base_url := request.FortiSASEAPIURL()
	req := c.NewRequest("GET", fmt.Sprintf("%s/resource-api/v1/admin/config?include_ems_version=true", base_url), nil, nil)
	ems_version, err := req.GetEMSVersion()
	if err != nil {
		err = fmt.Errorf("Could not get EMS version: %v", err)
		log.Printf("%v", err)
	} else if ems_version == "" {
		err = fmt.Errorf("EMS version is empty.")
		log.Printf("%v", err)
	} else {
		fssstatus.EMSVersion = ems_version
		c.FSSStatus = fssstatus
	}

	return err
}

// CheckUP checks whether username and password is valid
// If errors are encountered, it returns the error.
func (c *FortiSDKClient) CheckUP() error {
	return nil
}

func fortiAPIErrorFormat(result map[string]interface{}, body string) (code float64, err error) {
	code = -100
	if result != nil {
		if code, ok := result["code"].(float64); ok {
			// 200	OK: Request returns successful
			if code == 200.0 {
				return code, nil
			} else if code == 400.0 {
				err = fmt.Errorf("Bad Request - Request cannot be processed by the API (%.0f)", code)
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
