// Package auth 定义签名与验签的最小契约。
// Client 依赖这些接口，具体实现位于 signers 与 verifiers 子包；
// 实现不反向依赖 Client。
package auth

// Signer 生成协议要求的请求签名。实现必须并发安全。
type Signer interface {
	// Algorithm 返回请求公共参数 sign_type 的值，例如 "RSA2"。
	Algorithm() string
	// Sign 对待签名串生成签名，返回按协议编码的签名值（RSA2 为 base64）。
	Sign(data []byte) (string, error)
}

// Verifier 使用可信密钥校验响应签名。实现必须并发安全。
type Verifier interface {
	// Algorithm 返回期望的签名算法 sign_type 值。
	Algorithm() string
	// Verify 校验 data 与签名值；失败时返回错误，成功返回 nil。
	Verify(data []byte, sign string) error
}
