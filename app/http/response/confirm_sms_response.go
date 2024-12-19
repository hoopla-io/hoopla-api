package response

type ConfirmSmsResponse struct {
	AccessToken      string `json:"accessToken"`
	RefreshToken     string `json:"refreshToken"`
	ExpireAt         int64  `json:"expireAt"`
	PhoneNumber      string `json:"phone_number"`
}