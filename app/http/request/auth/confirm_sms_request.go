package auth_request

type ConfirmSmsRequest struct{
	SessionID string `json:"sessionId"`
	Code 	  int    `json:"code"`
}