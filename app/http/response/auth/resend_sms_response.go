package auth_response

type ResendSmsResponse struct {
	PhoneNumber      string `json:"phoneNumber"`
	SessionID        string `json:"sessionId"`
	SessionExpiresAt int64  `json:"sessionExpiresAt"`
}