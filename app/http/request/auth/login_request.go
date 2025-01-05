package auth_request

type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}
