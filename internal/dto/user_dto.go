package dto

import "time"

type UserDTO struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	PhoneNumber    string    `json:"phone_number"`
	MobileProvider string    `json:"mobile_provider"`
	CreateAt       time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
