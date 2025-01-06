package user_request

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
