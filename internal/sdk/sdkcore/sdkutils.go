package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func sendRequest(c *FortiSDKClient, input_model *InputModel) (output map[string]interface{}, code float64, err error) {
	method := input_model.HTTPMethod
	path := input_model.URL
	body_params := input_model.BodyParams
	head_params := input_model.HeadParams
	err = nil

	var body_bytes *bytes.Buffer
	if body_params != nil {
		locJSON, err := json.Marshal(body_params)
		if err != nil {
			log.Fatal(err)
			return nil, -101, err
		}
		body_bytes = bytes.NewBuffer(locJSON)
	} else {
		body_bytes = nil
	}

	req := c.NewRequest(method, path, head_params, body_bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("Cannot send request: %v", err)
		return nil, -102, err
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("Cannot get response body: %v", err)
		return nil, -103, err
	}
	log.Printf("[%v] [RESPONSE] [%v] %s", method, path, string(body))
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	code, err = fortiAPIErrorFormat(result, string(body))
	return result, code, err
}

func createUpdate(c *FortiSDKClient, input_model *InputModel) (map[string]interface{}, error) {
	var result map[string]interface{}
	var code float64
	var err error
	for i := 0; i < 100; {
		result, code, err = sendRequest(c, input_model)
		if err == nil {
			// return empty map if data is nil
			if result["data"] == nil {
				return result, nil
			} else if convered_rst, ok := result["data"].(map[string]interface{}); ok {
				return convered_rst, nil
			} else {
				return result, nil
			}
		} else if code == 429.0 {
			log.Printf("[%v] [RETRY] [%v] retry again due to 429", input_model.HTTPMethod, input_model.URL)
			time.Sleep(time.Second * 2)
			i = i + 10
			continue
		} else if code == 500.0 {
			log.Printf("[%v] [RETRY] [%v] retry again due to 500", input_model.HTTPMethod, input_model.URL)
			time.Sleep(time.Second * 2)
			i = i + 30
			continue
		} else {
			return nil, err
		}
	}
	return nil, err
}

func delete(c *FortiSDKClient, input_model *InputModel) error {
	var code float64
	var err error
	for i := 0; i < 100; {
		_, code, err = sendRequest(c, input_model)
		if err == nil {
			return err
		} else if code == 429.0 {
			log.Printf("[%v] [RETRY] [%v] retry again due to 429", input_model.HTTPMethod, input_model.URL)
			time.Sleep(time.Second * 2)
			i = i + 10
			continue
		} else if code == 500.0 {
			log.Printf("[%v] [RETRY] [%v] retry again due to 500", input_model.HTTPMethod, input_model.URL)
			time.Sleep(time.Second * 2)
			i = i + 30
			continue
		} else {
			return err
		}
	}
	return err
}

func read(c *FortiSDKClient, input_model *InputModel) (map[string]interface{}, error) {
	var result map[string]interface{}
	var code float64
	var err error
	for i := 0; i < 100; {
		result, code, err = sendRequest(c, input_model)
		if err == nil {
			if result["data"] == nil {
				return result, nil
			} else if convered_rst, ok := result["data"].(map[string]interface{}); ok {
				return convered_rst, nil
			} else if convered_rst, ok := result["data"].([]interface{}); ok {
				if len(convered_rst) == 0 {
					return nil, nil
				}
				return convered_rst[0].(map[string]interface{}), err
			} else {
				err = fmt.Errorf("Cannot convert respound type: %T", result["data"])
				return nil, err
			}
		} else if code == 404.0 {
			return nil, err
		} else if code == 429.0 {
			log.Printf("[%v] [ERROR] [%v] retry again due to 429", input_model.HTTPMethod, input_model.URL)
			time.Sleep(time.Second * 2)
			i = i + 10
			continue
		} else if code == 500.0 {
			log.Printf("[%v] [RETRY] [%v] retry again due to 500", input_model.HTTPMethod, input_model.URL)
			time.Sleep(time.Second * 2)
			i = i + 30
			continue
		} else {
			return nil, err
		}
	}
	return nil, err
}
