package forticlient

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

// Helper function to check if a value has a specific method
func hasMethod(v reflect.Value, methodName string) bool {
	method := v.MethodByName(methodName)
	return method.IsValid()
}

// MultValue describes the nested structure in the results
type InputModel struct {
	Mkey       interface{}            `json:"mkey"` // mkey may be string or int
	URL        string                 `json:"url"`
	HTTPMethod string                 `json:"http_method"`
	HeadParams map[string]interface{} `json:"head_params"`
	BodyParams map[string]interface{} `json:"body_params"`
	URLParams  map[string]interface{} `json:"url_params"`
}

func (input_model *InputModel) update() {
	if !strings.ContainsAny(input_model.URL, "[ | {") {
		return
	}
	// FortiSASE Terraform 1.1.0, direction has been deprecated, so we need to remove it from the URL.
	if strings.Contains(input_model.URL, "/{direction}") && input_model.URLParams["direction"] == nil {
		input_model.URL = strings.ReplaceAll(input_model.URL, "/{direction}", "s")
	}
	// replace mkey
	// Find all placeholder patterns {.*?}
	re := regexp.MustCompile(`{.*?}`)
	placeholders := re.FindAllString(input_model.URL, -1)
	placeholderCount := len(placeholders)

	if placeholderCount == 1 && input_model.Mkey != nil && input_model.Mkey != "<nil>" {
		// Only one placeholder, replace with Mkey
		input_model.URL = re.ReplaceAllString(input_model.URL, fmt.Sprintf("%v", input_model.Mkey))
	} else if placeholderCount > 0 {
		// Multiple placeholders found, handle each one individually
		updatedURL := input_model.URL
		for _, placeholder := range placeholders {
			// Extract content between {} - remove the braces
			paramKey := placeholder[1 : len(placeholder)-1]

			// Try to find the parameter in URLParams
			if input_model.URLParams != nil {
				if value, exists := input_model.URLParams[paramKey]; exists {
					// Replace this specific placeholder with the found value
					updatedURL = strings.ReplaceAll(updatedURL, placeholder, fmt.Sprintf("%v", value))
				} else {
					// Parameter not found, log and keep original placeholder
					log.Printf("[WARNING] URL parameter '%s' (from placeholder '%s') not found in URLParams", paramKey, placeholder)
				}
			} else {
				log.Printf("[WARNING] URLParams is nil, cannot replace placeholder '%s'", placeholder)
			}
		}

		input_model.URL = updatedURL
	}
}
