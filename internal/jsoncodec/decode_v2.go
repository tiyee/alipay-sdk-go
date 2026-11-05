//go:build go1.27 && goexperiment.jsonv2

package jsoncodec

import json "encoding/json/v2"

// Unmarshal accepts future response fields and requires exact JSON names.
func Unmarshal(data []byte, dst any) error { return json.Unmarshal(data, dst) }

// UnmarshalStrict also rejects unknown fields in user-supplied request files.
func UnmarshalStrict(data []byte, dst any) error {
	return json.Unmarshal(data, dst, json.RejectUnknownMembers(true))
}
