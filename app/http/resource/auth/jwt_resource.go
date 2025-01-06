package auth_resource

type JwtResource struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireAt     int64  `json:"expireAt"`
	ExpireAtMs   int64  `json:"expireAtMs"`
}
