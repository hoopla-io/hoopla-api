package auth_request

type ResendSmsRequest struct {
	SessionID string `json:"sessionId" binding:"required"`
}
