package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"reflect"
	"strings"
)

// MultipartFile 是 multipart/form-data 请求中的文件字段。Content 必须
// 包含完整文件内容；Filename 是文件名，ContentType 可选。
type MultipartFile struct {
	// Filename 是 multipart 文件部分使用的文件名。
	Filename string
	// ContentType 是文件部分的媒体类型；为空时不单独写入该头。
	ContentType string
	// Content 是文件的完整字节内容。
	Content []byte
}

// MultipartBodyEncoder 返回 multipart/form-data 请求编码器。
// 请求 struct 字段使用 `multipart:"field_name,file"` 写入文件，或
// `multipart:"field_name,json"` 将字段 JSON 编码为文本部分。
func MultipartBodyEncoder[T any]() Encoder[T] {
	return EncoderFunc[T](func(value T) (EncodedParams, error) {
		pathParams, err := extractPathParams(value)
		if err != nil {
			return EncodedParams{}, err
		}
		var body bytes.Buffer
		var signBody []byte
		jsonParts := 0
		writer := multipart.NewWriter(&body)
		rv := reflect.ValueOf(value)
		for rv.Kind() == reflect.Pointer {
			if rv.IsNil() {
				return EncodedParams{}, fmt.Errorf("alipay: multipart encode: nil request")
			}
			rv = rv.Elem()
		}
		if rv.Kind() != reflect.Struct {
			return EncodedParams{}, fmt.Errorf("alipay: multipart encode: expected struct, got %s", rv.Kind())
		}
		rt := rv.Type()
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			tag, ok := field.Tag.Lookup("multipart")
			if !ok || tag == "-" {
				continue
			}
			name, mode, _ := strings.Cut(tag, ",")
			fv := rv.Field(i)
			if fv.Kind() == reflect.Pointer {
				if fv.IsNil() {
					continue
				}
				fv = fv.Elem()
			}
			switch mode {
			case "file":
				file, ok := fv.Interface().(MultipartFile)
				if !ok {
					return EncodedParams{}, fmt.Errorf("alipay: multipart encode field %s: expected MultipartFile", field.Name)
				}
				if file.Filename == "" {
					return EncodedParams{}, fmt.Errorf("alipay: multipart encode field %s: filename is required", field.Name)
				}
				header := make(textproto.MIMEHeader)
				header.Set("Content-Disposition", fmt.Sprintf(`form-data; name=%q; filename=%q`, name, file.Filename))
				if file.ContentType != "" {
					header.Set("Content-Type", file.ContentType)
				}
				part, err := writer.CreatePart(header)
				if err != nil {
					return EncodedParams{}, fmt.Errorf("alipay: multipart encode field %s: %w", field.Name, err)
				}
				if _, err := part.Write(file.Content); err != nil {
					return EncodedParams{}, fmt.Errorf("alipay: multipart encode field %s: %w", field.Name, err)
				}
			case "json":
				jsonParts++
				if jsonParts > 1 {
					return EncodedParams{}, fmt.Errorf("alipay: multipart encode: multiple JSON fields are not supported")
				}
				raw, err := json.Marshal(fv.Interface())
				if err != nil {
					return EncodedParams{}, fmt.Errorf("alipay: multipart encode field %s: %w", field.Name, err)
				}
				signBody = raw
				header := make(textproto.MIMEHeader)
				header.Set("Content-Disposition", fmt.Sprintf(`form-data; name=%q`, name))
				header.Set("Content-Type", "application/json")
				part, err := writer.CreatePart(header)
				if err != nil {
					return EncodedParams{}, fmt.Errorf("alipay: multipart encode field %s: %w", name, err)
				}
				if _, err := part.Write(raw); err != nil {
					return EncodedParams{}, fmt.Errorf("alipay: multipart encode field %s: %w", field.Name, err)
				}
			default:
				return EncodedParams{}, fmt.Errorf("alipay: multipart encode field %s: unsupported mode %q", field.Name, mode)
			}
		}
		if err := writer.Close(); err != nil {
			return EncodedParams{}, fmt.Errorf("alipay: multipart encode: %w", err)
		}
		return EncodedParams{Mode: ParameterModeMultipart, Body: body.Bytes(), SignBody: signBody, ContentType: writer.FormDataContentType(), PathParams: pathParams}, nil
	})
}
