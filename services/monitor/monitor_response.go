package monitor

// HeartbeatSynResponse contains the merchant identifier returned by the API.
type HeartbeatSynResponse struct {
	// PID is the merchant identifier.
	PID *string `json:"pid,omitempty"`
}
