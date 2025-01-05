package dto

type SessionDTO struct {
	PhoneNumber    string  `json:"phone_number"`
	MobileProvider string  `json:"mobile_provder"`
	Session        Session `json:"session"`
}

type Session struct {
	Code      int   `json:"code"`
	ExpiresAt int64 `json:"expires_at"`
}
