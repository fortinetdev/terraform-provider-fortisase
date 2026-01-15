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
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

func checkVersionMatch(v string, new_version_map map[string][]string) (bool, error) {
	v1, err := version.NewVersion(v)
	if err != nil {
		return false, err
	}

	for operator, version_list := range new_version_map {
		if operator == "=" {
			for _, cur_version := range version_list {
				if cur_version == v {
					return true, nil
				}
			}
		} else if operator == ">=" {
			min_version, err := version.NewVersion(version_list[0])
			if err != nil {
				continue
			}
			if v1.GreaterThanOrEqual(min_version) {
				return true, nil
			}
		} else if operator == "<=" {
			max_version, err := version.NewVersion(version_list[0])
			if err != nil {
				continue
			}
			if v1.LessThanOrEqual(max_version) {
				return true, nil
			}
		}
	}
	var supported_version_list []string
	for operator, version_list := range new_version_map {
		supported_version_list = append(supported_version_list, operator+strings.Join(version_list, ","))
	}
	err = fmt.Errorf("requires FortiSASE version: %s, your device version is: %s.", strings.Join(supported_version_list, ""), v)
	return false, err
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
