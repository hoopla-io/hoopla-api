package auth_request

type ConfirmSmsRequest struct {
	SessionID string `json:"sessionId" binding:"required"`
	Code      int    `json:"code" binding:"required"`
}
