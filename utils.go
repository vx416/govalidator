package govalidator

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
)

// containsRequiredField check rules contain any required field
func isContainRequiredField(rules []string) bool {
	for _, rule := range rules {
		if rule == "required" {
			return true
		}
	}
	return false
}

// isRuleExist check if the provided rule name is exist or not
func isRuleExist(rule string) bool {
	if strings.Contains(rule, ":") {
		rule = strings.Split(rule, ":")[0]
	}
	extendedRules := []string{"size", "mime", "ext"}
	for _, r := range extendedRules {
		if r == rule {
			return true
		}
	}
	if _, ok := rulesFuncMap[rule]; ok {
		return true
	}
	return false
}

// toString force data to be string
func toString(v interface{}) string {
	str, ok := v.(string)
	if !ok {
		str = fmt.Sprintf("%v", v)
	}
	return str
}

// isEmpty check a type is Zero
func isEmpty(x interface{}) bool {
	rt := reflect.TypeOf(x)
	if rt == nil {
		return true
	}
	rv := reflect.ValueOf(x)
	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice:
		return rv.Len() == 0
	}
	return reflect.DeepEqual(x, reflect.Zero(rt).Interface())
}

// isValuer check the filed implement driver.Valuer or not
func isValuer(field reflect.Value) (driver.Valuer, bool) {
	var fieldRaw interface{}
	fieldRaw = field.Interface()
	if scanner, ok := fieldRaw.(driver.Valuer); ok {
		return scanner, ok
	}
	if field.CanAddr() {
		fieldRaw = field.Addr().Interface()
	}
	if scanner, ok := fieldRaw.(driver.Valuer); ok {
		return scanner, ok
	}
	return nil, false
}
