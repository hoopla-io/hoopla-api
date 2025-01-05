package auth_resource

type SessionResource struct {
	PhoneNumber      string `json:"phoneNumber"`
	SessionID        string `json:"sessionId"`
	SessionExpiresAt int64  `json:"sessionExpiresAt"`
}
