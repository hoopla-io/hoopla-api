package dto

type UserDTO struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	MobileProvider string `json:"mobile_provider"`
}