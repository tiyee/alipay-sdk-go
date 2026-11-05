package system

// GrantType 是授权方式，取值依据官方 V3 文档。
type GrantType string

// IsKnown 报告授权方式是否属于文档已知值。
func (g GrantType) IsKnown() bool {
	return g == GrantTypeAuthorizationCode || g == GrantTypeRefreshToken
}
