package request

type ResendSmsRequest struct{
	SessionID string `json:"sessionId"`
}