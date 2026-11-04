package customapi

import "net/url"

// toURLValues 是 FormEncoder 的内部辅助，保持示例文件聚焦编码逻辑。
func toURLValues(form map[string][]string) url.Values {
	values := url.Values{}
	for key, vs := range form {
		values[key] = vs
	}
	return values
}
