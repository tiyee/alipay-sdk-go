package monitor

// HeartbeatSynRequest contains the heartbeat request parameters.
type HeartbeatSynRequest struct {
	// BizContent is passed unchanged as the business content.
	BizContent *string `json:"biz_content,omitempty"`
}
