package provider

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	forticlient "github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func AttributeTypes[T any](ctx context.Context) (map[string]attr.Type, diag.Diagnostics) {
	var diags diag.Diagnostics
	var t T
	val := reflect.ValueOf(t)
	typ := val.Type()

	if typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct {
		val = reflect.New(typ.Elem()).Elem()
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		diags.Append(diag.NewErrorDiagnostic("Invalid type", fmt.Sprintf("%T has unsupported type: %s", t, typ)))
		return nil, diags
	}

	attributeTypes := make(map[string]attr.Type)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			continue // Skip unexported fields.
		}
		tag := field.Tag.Get(`tfsdk`)
		if tag == "-" {
			continue // Skip explicitly excluded fields.
		}
		if tag == "" {
			diags.Append(diag.NewErrorDiagnostic("Invalid type", fmt.Sprintf(`%T needs a struct tag for "tfsdk" on %s`, t, field.Name)))
			return nil, diags
		}

		if v, ok := val.Field(i).Interface().(attr.Value); ok {
			attributeTypes[tag] = v.Type(ctx)
		}
	}

	return attributeTypes, nil
}

func validateConvIPMask2CIDR(oNewIP, oOldIP string) string {
	if oNewIP != oOldIP && strings.Contains(oNewIP, "/") && strings.Contains(oOldIP, " ") {
		line := strings.Split(oOldIP, " ")
		if len(line) >= 2 {
			ip := line[0]
			mask := line[1]
			prefixSize, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
			return ip + "/" + strconv.Itoa(prefixSize)
		}
	}
	return oOldIP
}

func fortiStringValue(t interface{}) string {
	if v, ok := t.(string); ok {
		return v
	} else {
		return ""
	}
}

func fortiIntValue(t interface{}) int {
	if v, ok := t.(float64); ok {
		return int(v)
	} else {
		return 0
	}
}

func escapeFilter(filter string) string {
	var rstSb strings.Builder
	andSlice := strings.Split(filter, "&")

	for i := 0; i < len(andSlice); i++ {
		orSlice := strings.Split(andSlice[i], ",")
		if i > 0 {
			rstSb.WriteString("&")
		}
		rstSb.WriteString("filter=")
		for j := 0; j < len(orSlice); j++ {
			reg := regexp.MustCompile(`([^=*!@><]+)([=*!@><]+)([^=*!@><]+)`)
			match := reg.FindStringSubmatch(orSlice[j])
			if j > 0 {
				rstSb.WriteString(",")
			}
			if match != nil {
				argName := match[1]
				argName = strings.ReplaceAll(argName, "_", "-")
				argName = strings.ReplaceAll(argName, "fssid", "id")
				argName = strings.ReplaceAll(argName, ".", "\\.")
				argName = strings.ReplaceAll(argName, "\\", "\\\\")
				argValue := url.QueryEscape(match[3])
				rstSb.WriteString(argName)
				rstSb.WriteString(match[2])
				rstSb.WriteString(argValue)
			}
		}
	}
	return rstSb.String()
}

func sortStringwithNumber(v string) string {
	i := len(v) - 1
	for ; i >= 0; i-- {
		if '0' > v[i] || v[i] > '9' {
			break
		}
	}
	i++

	b64 := make([]byte, 64/8)
	s64 := v[i:]
	if len(s64) > 0 {
		u64, err := strconv.ParseUint(s64, 10, 64)
		if err == nil {
			binary.BigEndian.PutUint64(b64, u64+1)
		}
	}

	return v[:i] + string(b64)
}

func fortiAPIPatch(t interface{}) bool {
	if t == nil {
		return false
	} else if _, ok := t.(string); ok {
		return true
	} else if _, ok := t.(float64); ok {
		return true
	} else if _, ok := t.([]interface{}); ok {
		return true
	}

	return false
}

func isImportTable() bool {
	itable := os.Getenv("FORTISASE_IMPORT_TABLE")
	if itable == "false" {
		return false
	}
	return true
}

func convintf2i(v interface{}) interface{} {
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return v
		}
		return t[0]
	} else if t, ok := v.(string); ok {
		if t == "" {
			return 0
		} else if iVal, _ := strconv.Atoi(t); ok {
			return iVal
		}
	}
	return v
}

func convintflist2str(v interface{}) interface{} {
	res := ""
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return res
		}

		bFirst := true
		for _, v1 := range t {
			if t1, ok := v1.(float64); ok {
				if bFirst == true {
					res += strconv.Itoa(int(t1))
					bFirst = false
				} else {
					res += " "
					res += strconv.Itoa(int(t1))
				}
			}
		}
	}
	return res
}

func convmap2str(v, tfv interface{}, target_key string) interface{} {
	if vMap, ok := v.([]interface{}); ok {
		if len(vMap) == 0 {
			return ""
		}
		vsList := make([]string, len(vMap))
		for i, r := range vMap {
			if item, ok := r.(map[string]interface{})[target_key]; ok {
				if ts, ok := item.(string); ok {
					vsList[i] = strings.TrimSpace(fmt.Sprintf("%v", ts))
					if strings.Contains(vsList[i], ",") {
						vsList[i] = "'" + vsList[i] + "'"
					}
				}
			}
		}
		if tfv != nil {
			if tfvs := fmt.Sprintf("%v", tfv); tfvs != "" {
				tfvList := flattenStringList(tfv).([]string)
				if len(tfvList) == len(vsList) {
					tfvDict := make(map[string]bool)
					for _, item := range tfvList {
						tfvDict[item] = true
					}
					for _, item := range vsList {
						item = strings.Trim(item, "'\" ")
						if _, ok := tfvDict[item]; !ok {
							return strings.Join(vsList[:], ", ")
						}
					}
					return tfv
				}
			}
		}
		return strings.Join(vsList[:], ", ")

	}
	return v
}

func flattenStringList(v interface{}) interface{} {
	if v == nil {
		return v
	}
	vsList := []string{}
	if cv, ok := v.(string); ok {
		if strings.Contains(cv, "'") || strings.Contains(cv, "\"") {
			re := regexp.MustCompile(`['\"].*?['\"]`)
			comma := re.FindAllString(cv, -1)
			non_comma := re.Split(cv, -1)
			for i := range non_comma {
				cur_list := strings.Split(non_comma[i], ",")
				for _, item := range cur_list {
					item = strings.TrimSpace(item)
					if item != "" {
						vsList = append(vsList, item)
					}
				}
				if i < len(comma) {
					cur_item := strings.Trim(comma[i], "'\" ")
					vsList = append(vsList, cur_item)
				}
			}
		} else {
			vsList = strings.Split(cv, ",")
		}
	} else if vList, ok := v.([]interface{}); ok {
		for _, item := range vList {
			vsList = append(vsList, fmt.Sprintf("%v", item))
		}
	}
	if len(vsList) == 0 {
		return vsList
	}
	for i, item := range vsList {
		vsList[i] = strings.TrimSpace(item)
	}

	return vsList
}

func checkVersionMatch(forticlient *forticlient.FortiSDKClient, supported_version_map map[string][]string) (pass bool, err error) {
	if supported_version_map == nil || len(supported_version_map) == 0 {
		return true, nil
	}

	fssstatus := forticlient.FSSStatus

	if fssstatus == nil || fssstatus.EMSVersion == "" {
		err = forticlient.GetFSSStatus()
		if err != nil {
			return true, err
		}
		fssstatus = forticlient.FSSStatus
	}

	for k, v := range supported_version_map {
		if k == "EMS" {
			fss_ems_version := fssstatus.EMSVersion
			for _, sv := range v {
				if strings.HasPrefix(fss_ems_version, sv) {
					return true, nil
				}
			}
			if err == nil {
				err = fmt.Errorf("Requires FortiSASE EMS version: %s, your device EMS version is: %s.", v, fss_ems_version)
			} else {
				err = fmt.Errorf("Requires FortiSASE EMS version: %s, your device EMS version is: %s. %v", v, fss_ems_version, err)
			}
			return false, err
		}
	}
	return true, nil
}

func toCertFormat(v interface{}) interface{} {
	if t, ok := v.(string); ok {
		if t != "" && !strings.HasPrefix(t, "\"") {
			t = strings.TrimRight(t, "\n")
			return "\"" + t + "\""
		}
	}
	return v
}

func remove_quote(v interface{}) interface{} {
	if t, ok := v.(string); ok {
		t = strings.ReplaceAll(t, "\"", "")
		t = strings.TrimSpace(t)
		return t
	}
	return v
}

func bZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}

func expandSetToStringList(varSet types.Set) []string {
	elements := varSet.Elements()

	result := make([]string, 0, len(elements))
	for _, e := range elements {
		if strVal, ok := e.(types.String); ok {
			result = append(result, strVal.ValueString())
		}
	}
	return result
}

func expandSetToInt64List(varSet types.Set) []int64 {
	elements := varSet.Elements()

	result := make([]int64, 0, len(elements))
	for _, e := range elements {
		if intVal, ok := e.(types.Int64); ok {
			result = append(result, intVal.ValueInt64())
		}
	}
	return result
}

func parseStringValue(v interface{}) basetypes.StringValue {
	if v == nil {
		return types.StringNull()
	}
	switch val := v.(type) {
	case string:
		return types.StringValue(val)
	case bool:
		if val {
			return types.StringValue("enable")
		}
		return types.StringValue("disable")
	default:
		return types.StringValue(v.(string))
	}
}

func parseBoolValue(v interface{}) basetypes.BoolValue {
	if v == nil {
		return types.BoolNull()
	}
	switch val := v.(type) {
	case bool:
		return types.BoolValue(val)
	case string:
		if val == "true" || val == "enable" {
			return types.BoolValue(true)
		} else if val == "false" || val == "disable" {
			return types.BoolValue(false)
		}
	}
	return types.BoolNull()
}

func parseFloat64Value(v interface{}) basetypes.Float64Value {
	if v == nil {
		return types.Float64Null()
	}
	switch val := v.(type) {
	case float64:
		return types.Float64Value(val)
	case int:
		return types.Float64Value(float64(val))
	case string:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return types.Float64Null()
		}
		return types.Float64Value(f)
	}
	return types.Float64Null()
}

func parseMapValue(ctx context.Context, v interface{}, element_type attr.Type) basetypes.MapValue {
	var m basetypes.MapValue
	if v != nil {
		m, _ = types.MapValueFrom(ctx, element_type, v.(map[string]interface{}))
	} else {
		m = types.MapNull(element_type)
	}
	return m
}

func parseSetValue(ctx context.Context, v interface{}, element_type attr.Type) basetypes.SetValue {
	var m basetypes.SetValue
	if v != nil {
		m, _ = types.SetValueFrom(ctx, element_type, v.([]interface{}))
	} else {
		m = types.SetNull(element_type)
	}
	return m
}

func parseListValue(ctx context.Context, v interface{}, element_type attr.Type) basetypes.ListValue {
	var m basetypes.ListValue
	if v != nil {
		m, _ = types.ListValueFrom(ctx, element_type, v.([]interface{}))
	} else {
		m = types.ListNull(element_type)
	}
	return m
}

func isZeroStruct(s interface{}) bool {
	return reflect.ValueOf(s).IsZero()
}

func isSameStruct(s1, s2 interface{}) bool {
	return reflect.DeepEqual(s1, s2)
}

// extractValue extracts the actual value from attr.Value; if not an attr.Value, returns the original value
func extractValue(v any) any {
	// Check if the value implements the attr.Value interface (by checking for a Type() method)
	if valuer, ok := v.(interface {
		Type(context.Context) attr.Type
	}); ok {
		// This is an attr.Value, try to extract its underlying value
		if stringer, ok := v.(interface{ ValueString() string }); ok {
			return stringer.ValueString()
		}
		if inter, ok := v.(interface{ ValueInt64() int64 }); ok {
			return inter.ValueInt64()
		}
		if booler, ok := v.(interface{ ValueBool() bool }); ok {
			return booler.ValueBool()
		}
		if floater, ok := v.(interface{ ValueFloat64() float64 }); ok {
			return floater.ValueFloat64()
		}
		// If value cannot be extracted, return the original value
		_ = valuer // Use valuer to avoid unused variable warning
	}
	return v
}

// isSetSuperset checks if superset contains all elements from subset
// Supports complex nested structures including list of maps, maps with nested maps, etc.
// Handles both native Go types ([]interface{}) and Terraform types ([]attr.Value)
func isSetSuperset(superset any, subset any) bool {
	supersetVal := reflect.ValueOf(superset)
	subsetVal := reflect.ValueOf(subset)

	if supersetVal.Len() < subsetVal.Len() {
		return false
	}

	// For each element in subset, check if it exists in superset
	for i := 0; i < subsetVal.Len(); i++ {
		subsetItem := extractValue(subsetVal.Index(i).Interface())
		found := false

		// Search for this item in superset
		for j := 0; j < supersetVal.Len(); j++ {
			supersetItem := extractValue(supersetVal.Index(j).Interface())
			if isSameStruct(supersetItem, subsetItem) {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func getErrorDetail(input_model *forticlient.InputModel, response map[string]interface{}) string {
	result := ""
	result += fmt.Sprintf("[API Request] %v (%v)\n", input_model.URL, input_model.HTTPMethod)
	request_json_bytes, err := json.MarshalIndent(input_model.BodyParams, "", "  ")
	if err != nil {
		result += fmt.Sprintf("%v\n\n", input_model.HTTPMethod)
	} else {
		result += fmt.Sprintf("%s\n\n", string(request_json_bytes))
	}
	response_json_bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		result += fmt.Sprintf("[API Response]\n%v\n", response)
	} else {
		result += fmt.Sprintf("[API Response]\n%s\n", string(response_json_bytes))
	}
	return result
}

func parseJsonString(apiResp, tfConf interface{}) (result basetypes.StringValue, err error) {
	if apiResp == nil {
		if tfConf != nil {
			if tfVal, ok := tfConf.(basetypes.StringValue); ok && !tfVal.IsNull() && tfVal.ValueString() == "" {
				return tfVal, nil
			}
		}
		return types.StringNull(), nil
	} else if tfConf == nil {
		return types.StringValue(apiResp.(string)), nil
	}
	apiStr := apiResp.(string)
	tfVal := tfConf.(basetypes.StringValue)
	tfStr := tfVal.ValueString()
	var apiObj interface{}
	var tfObj interface{}
	bsame := false
	if err := json.Unmarshal([]byte(apiStr), &apiObj); err != nil {
		err = fmt.Errorf("API response is not JSON format: %v", apiStr)
	}
	if err := json.Unmarshal([]byte(tfStr), &tfObj); err != nil {
		err = fmt.Errorf("Terraform configuration is not JSON format: %v", tfStr)
	}

	if apiObj != nil && tfObj != nil {
		bsame = reflect.DeepEqual(apiObj, tfObj)
	} else if apiObj == nil && tfObj == nil {
		apiStrNoSpace := strings.ReplaceAll(apiStr, " ", "")
		tfStrNoSpace := strings.ReplaceAll(tfStr, " ", "")
		bsame = apiStrNoSpace == tfStrNoSpace
	}

	if bsame {
		return tfVal, nil
	} else {
		return types.StringValue(apiStr), err
	}
}

func getAPICode(output map[string]interface{}) (int, bool) {
	if codeVal, hasCode := output["code"]; hasCode {
		switch v := codeVal.(type) {
		case int:
			return v, true
		case float64:
			return int(v), true
		}
	}
	return 0, false
}

func getErrorCode(output map[string]interface{}) (int, bool) {
	if errorVal, hasError := output["error"]; hasError {
		if errorMap, ok := errorVal.(map[string]interface{}); ok {
			if codeVal, hasCode := errorMap["code"]; hasCode {
				switch v := codeVal.(type) {
				case int:
					return v, true
				case float64:
					return int(v), true
				}
			}
		}
	}
	return 0, false
}

func isNotFoundResponse(response map[string]interface{}) bool {
	if response == nil {
		return false
	}
	if httpStatus, ok := response["http_status"].(float64); ok {
		return httpStatus == 404.0
	}
	if code, ok := response["code"].(float64); ok {
		return code == 404.0
	}
	return false
}

type useStateForUnknownOrNullOnCreateModifier struct{}

func (m useStateForUnknownOrNullOnCreateModifier) Description(_ context.Context) string {
	return "Preserves prior state during updates and keeps omitted create-time API defaults out of the plan."
}

func (m useStateForUnknownOrNullOnCreateModifier) MarkdownDescription(_ context.Context) string {
	return "Preserves prior state during updates and keeps omitted create-time API defaults out of the plan."
}

func UseStateForUnknownOrNullOnCreateString() planmodifier.String {
	return useStateForUnknownOrNullOnCreateModifier{}
}

func UseStateForUnknownOrNullOnCreateFloat64() planmodifier.Float64 {
	return useStateForUnknownOrNullOnCreateModifier{}
}

func UseStateForUnknownOrNullOnCreateBool() planmodifier.Bool {
	return useStateForUnknownOrNullOnCreateModifier{}
}

func UseStateForUnknownOrNullOnCreateObject() planmodifier.Object {
	return useStateForUnknownOrNullOnCreateModifier{}
}

func UseStateForUnknownOrNullOnCreateSet() planmodifier.Set {
	return useStateForUnknownOrNullOnCreateModifier{}
}

func UseStateForUnknownOrNullOnCreateList() planmodifier.List {
	return useStateForUnknownOrNullOnCreateModifier{}
}

func UseStateForUnknownOrNullOnCreateMap() planmodifier.Map {
	return useStateForUnknownOrNullOnCreateModifier{}
}

func (m useStateForUnknownOrNullOnCreateModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.IsNull() {
		resp.PlanValue = req.ConfigValue
		return
	}
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func (m useStateForUnknownOrNullOnCreateModifier) PlanModifyFloat64(ctx context.Context, req planmodifier.Float64Request, resp *planmodifier.Float64Response) {
	if !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.IsNull() {
		resp.PlanValue = req.ConfigValue
		return
	}
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func (m useStateForUnknownOrNullOnCreateModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.IsNull() {
		resp.PlanValue = req.ConfigValue
		return
	}
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func (m useStateForUnknownOrNullOnCreateModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.IsNull() {
		resp.PlanValue = req.ConfigValue
		return
	}
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func (m useStateForUnknownOrNullOnCreateModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	if !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.IsNull() {
		resp.PlanValue = req.ConfigValue
		return
	}
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func (m useStateForUnknownOrNullOnCreateModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.IsNull() {
		resp.PlanValue = req.ConfigValue
		return
	}
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func (m useStateForUnknownOrNullOnCreateModifier) PlanModifyMap(ctx context.Context, req planmodifier.MapRequest, resp *planmodifier.MapResponse) {
	if !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.IsNull() {
		resp.PlanValue = req.ConfigValue
		return
	}
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

func getCreateResponseMkey(response map[string]interface{}, mkeyName string) (string, bool) {
	if response == nil {
		return "", false
	}

	keys := []string{"mkey"}
	if mkeyName != "" && mkeyName != "mkey" {
		keys = append(keys, mkeyName)
	}

	for _, key := range keys {
		switch value := response[key].(type) {
		case string:
			if value != "" && value != "<nil>" {
				return value, true
			}
		case json.Number:
			return value.String(), true
		case float64:
			return strconv.FormatFloat(value, 'f', -1, 64), true
		case float32:
			return strconv.FormatFloat(float64(value), 'f', -1, 32), true
		case int:
			return strconv.Itoa(value), true
		case int64:
			return strconv.FormatInt(value, 10), true
		case uint64:
			return strconv.FormatUint(value, 10), true
		}
	}

	return "", false
}
