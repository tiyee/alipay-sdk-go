package protocol

import (
	"bytes"
	"encoding/json"
	"encoding/json/jsontext"
	"fmt"
	"io"
)

// RawResponse 保存从网关同步响应中按原样提取的片段。
// Node 与 ErrorResponse 保留原始字节，验签前不重新编码或压缩。
type RawResponse struct {
	// Node 是 responseKey 对应业务节点的原始 JSON 字节。
	Node jsontext.Value
	// ErrorResponse 是 error_response 节点的原始 JSON 字节；
	// 该报文不带签名。
	ErrorResponse jsontext.Value
	// Sign 是顶层 sign 字段的值。
	Sign string
}

// ExtractRawResponse 解析网关响应的顶层 JSON 对象，按名提取
// 业务节点、error_response 与 sign 的原始字节。
// 重复的关键节点名、非对象顶层结构与对象后的尾部内容均为错误。
func ExtractRawResponse(body []byte, responseKey string) (*RawResponse, error) {
	dec := jsontext.NewDecoder(bytes.NewReader(body))
	tok, err := dec.ReadToken()
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
	if tok.Kind() != '{' {
		return nil, fmt.Errorf("response is not a JSON object")
	}
	var raw RawResponse
	var sawNode, sawError, sawSign bool
	for {
		tok, err := dec.ReadToken()
		if err != nil {
			return nil, fmt.Errorf("read response member: %w", err)
		}
		if tok.Kind() == '}' {
			break
		}
		if tok.Kind() != '"' {
			return nil, fmt.Errorf("unexpected response member token %v", tok.Kind())
		}
		name := tok.String()
		switch name {
		case responseKey:
			if sawNode {
				return nil, fmt.Errorf("duplicate response node %q", responseKey)
			}
			sawNode = true
			value, readErr := dec.ReadValue()
			if readErr != nil {
				return nil, fmt.Errorf("read response node: %w", readErr)
			}
			// ReadValue 的结果引用解码器内部缓冲，继续读取会覆写；
			// 立即复制以保留原始字节。
			raw.Node = jsontext.Value(bytes.Clone(value))
		case ErrorKeyName:
			if sawError {
				return nil, fmt.Errorf("duplicate %q node", ErrorKeyName)
			}
			sawError = true
			value, readErr := dec.ReadValue()
			if readErr != nil {
				return nil, fmt.Errorf("read error_response node: %w", readErr)
			}
			raw.ErrorResponse = jsontext.Value(bytes.Clone(value))
		case ParamSign:
			if sawSign {
				return nil, fmt.Errorf("duplicate %q member", ParamSign)
			}
			sawSign = true
			value, readErr := dec.ReadValue()
			if readErr != nil {
				return nil, fmt.Errorf("read sign: %w", readErr)
			}
			if err := json.Unmarshal(value, &raw.Sign); err != nil {
				return nil, fmt.Errorf("decode sign: %w", err)
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return nil, fmt.Errorf("skip response member %q: %w", name, err)
			}
		}
	}
	if !sawNode && !sawError {
		return nil, fmt.Errorf("neither %q nor %q node present", responseKey, ErrorKeyName)
	}
	if _, err := dec.ReadToken(); err != io.EOF {
		return nil, fmt.Errorf("trailing data after response object")
	}
	return &raw, nil
}

// GatewayError 汇总网关报文中的错误字段。
type GatewayError struct {
	Code    string
	Msg     string
	SubCode string
	SubMsg  string
}

// ParseGatewayError 解析业务节点或 error_response 节点中的
// code、msg、sub_code、sub_msg。code 缺失视为格式异常。
func ParseGatewayError(node jsontext.Value) (GatewayError, error) {
	var raw struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	}
	if err := json.Unmarshal(node, &raw); err != nil {
		return GatewayError{}, fmt.Errorf("decode gateway error: %w", err)
	}
	if raw.Code == "" {
		return GatewayError{}, fmt.Errorf("gateway error node has no code")
	}
	return GatewayError{
		Code:    raw.Code,
		Msg:     raw.Msg,
		SubCode: raw.SubCode,
		SubMsg:  raw.SubMsg,
	}, nil
}

// ParseResponseCode 提取业务节点的 code 字段；
// 成功与业务失败报文都必须携带该字段。
func ParseResponseCode(node jsontext.Value) (string, error) {
	var head struct {
		Code string `json:"code"`
	}
	if err := json.Unmarshal(node, &head); err != nil {
		return "", fmt.Errorf("decode response code: %w", err)
	}
	if head.Code == "" {
		return "", fmt.Errorf("response node has no code")
	}
	return head.Code, nil
}
