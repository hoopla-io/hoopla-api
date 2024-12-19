package response

type LoginResponse struct {
	PhoneNumber      string `json:"phone_number"`
	SessionID        string `json:"session_id"`
	SessionExpiresAt int64    `json:"session_expires_at"`
}