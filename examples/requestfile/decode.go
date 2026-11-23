// Package requestfile 读取 SDK 示例和沙箱测试使用的请求 JSON 文件。
package requestfile

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/tiyee/alipay-sdk-go/internal/jsoncodec"
)

// Decode 将单个 JSON 对象严格解码为具体请求结构体，拒绝未知或重复字段。
// json:"-" 的混合请求字段使用其 query 或 path 协议名，其余字段按 JSON 标签或 Go 字段名读取。
func Decode[T any](body []byte) (T, error) {
	var zero, req T
	trimmed := bytes.Trim(body, " \t\r\n")
	if len(trimmed) == 0 || trimmed[0] != '{' {
		return zero, errors.New("request must be a JSON object")
	}
	// 混合请求的 query/path 字段标记为 json:"-"，避免它们被写进网关请求体。
	// 请求文件可以按协议名提供这些字段，其余内容继续严格解码。
	value := reflect.ValueOf(&req).Elem()
	var fields map[string]json.RawMessage
	if value.Kind() != reflect.Struct {
		return zero, errors.New("request type must be a struct")
	}
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		if !field.IsExported() || field.Tag.Get("json") != "-" {
			continue
		}
		name := field.Tag.Get("path")
		if name == "" {
			name, _, _ = strings.Cut(field.Tag.Get("query"), ",")
		}
		if name == "" || name == "-" {
			continue
		}
		if fields == nil {
			if err := jsoncodec.Unmarshal(trimmed, &fields); err != nil {
				return zero, err
			}
		}
		if raw, ok := fields[name]; ok {
			if err := jsoncodec.UnmarshalStrict(raw, value.Field(i).Addr().Interface()); err != nil {
				return zero, err
			}
			delete(fields, name)
		}
	}
	if fields != nil {
		rest, err := json.Marshal(fields)
		if err != nil {
			return zero, err
		}
		trimmed = rest
	}
	if err := jsoncodec.UnmarshalStrict(trimmed, &req); err != nil {
		return zero, err
	}
	return req, nil
}
