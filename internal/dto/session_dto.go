package dto

type SessionDTO struct {
	PhoneNumber    string `json:"phone_number"`
	MobileProvider string `json:"mobile_provder"`
	Session        Session `json:"session"`
}

type Session struct {
	Attempt   int `json:"attempt"`
	Code      int `json:"code"`
	ExpiresAt int `json:"expires_at"`
}