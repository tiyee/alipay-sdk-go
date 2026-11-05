package core

import (
	"fmt"
	"reflect"
)

// extractPathParams reads path-tagged fields from an API request. It accepts
// strings and pointers to strings so callers can distinguish an absent value.
func extractPathParams(value any) (map[string]string, error) {
	rv := reflect.ValueOf(value)
	for rv.IsValid() && rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil, &ValidationError{Field: "path_params", Reason: "request must not be nil"}
		}
		rv = rv.Elem()
	}
	if !rv.IsValid() || rv.Kind() != reflect.Struct {
		return nil, nil
	}

	params := make(map[string]string)
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		name, ok := field.Tag.Lookup("path")
		if !ok || name == "" || name == "-" {
			continue
		}
		fv := rv.Field(i)
		if fv.Kind() == reflect.Pointer {
			if fv.IsNil() {
				continue
			}
			fv = fv.Elem()
		}
		if fv.Kind() != reflect.String {
			return nil, fmt.Errorf("alipay: encode path parameter %s: expected string, got %s", name, fv.Kind())
		}
		if _, exists := params[name]; exists {
			return nil, fmt.Errorf("alipay: encode path parameter %s: duplicate field", name)
		}
		params[name] = fv.String()
	}
	if len(params) == 0 {
		return nil, nil
	}
	return params, nil
}
