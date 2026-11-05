package core

import (
	"fmt"
	"net/url"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

// QueryEncoder 返回 V3 query 参数编码器：按 struct 字段声明顺序，
// 依据 query 标签将请求值写入 url.Values，可用 omitempty 和 required 指定字段选项。
//
// required 标签要求参数存在。必填指针使用 nil 表示缺失，非 nil 保留显式零值。
// 支持的值类型：string、bool、整数及其具名类型、上述类型的指针
// （nil 跳过，非 nil 保留显式零值）与切片（逐元素输出）。不支持嵌套 struct 与其他类型，
// 编码时返回错误；V3 GET 接口的参数都是扁平字符串。
// 输出顺序确定，同一输入产生一致的编码结果，用于签名与传输。
func QueryEncoder[T any]() Encoder[T] {
	return EncoderFunc[T](func(value T) (EncodedParams, error) {
		form, err := encodeQuery(value)
		if err != nil {
			return EncodedParams{}, err
		}
		pathParams, err := extractPathParams(value)
		if err != nil {
			return EncodedParams{}, err
		}
		return EncodedParams{Mode: ParameterModeQuery, Form: form, PathParams: pathParams}, nil
	})
}

// encodeQuery 按 struct 字段顺序展开 query 参数。
func encodeQuery(value any) (url.Values, error) {
	form := url.Values{}
	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return form, nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("alipay: query encode: expected struct, got %s", rv.Kind())
	}
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if _, ok := field.Tag.Lookup("path"); ok {
			continue
		}
		tag, ok := field.Tag.Lookup("query")
		if !ok || tag == "-" {
			continue
		}
		name, opts, _ := strings.Cut(tag, ",")
		options := strings.Split(opts, ",")
		if name == "" {
			return nil, fmt.Errorf("alipay: query encode: field %s has empty query name", field.Name)
		}
		if err := appendQueryField(form, name, slices.Contains(options, "omitempty"), rv.Field(i)); err != nil {
			return nil, fmt.Errorf("alipay: query encode field %s: %w", field.Name, err)
		}
		if slices.Contains(options, "required") && !form.Has(name) {
			return nil, &ValidationError{Field: name, Reason: "required query parameter is missing"}
		}
	}
	return form, nil
}

// appendQueryField 写入单个标量、指针或切片字段。
func appendQueryField(form url.Values, name string, omitEmpty bool, value reflect.Value) error {
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
		omitEmpty = false
	}
	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			if err := appendQueryValue(form, name, omitEmpty, value.Index(i)); err != nil {
				return err
			}
		}
		return nil
	}
	return appendQueryValue(form, name, omitEmpty, value)
}

func appendQueryValue(form url.Values, name string, omitEmpty bool, value reflect.Value) error {
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
		omitEmpty = false
	}
	var text string
	switch value.Kind() {
	case reflect.String:
		text = value.String()
	case reflect.Bool:
		text = strconv.FormatBool(value.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		text = strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		text = strconv.FormatUint(value.Uint(), 10)
	default:
		return fmt.Errorf("unsupported type %s", value.Type())
	}
	if omitEmpty && value.IsZero() {
		return nil
	}
	form.Add(name, text)
	return nil
}
