package user_request

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}
