package auth_resource

type LoginResource struct {
	UserID       uint   `json:"userId"`
	PhoneNumber  string `json:"phoneNumber"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireAt     int64  `json:"expireAt"`
	IsNewUser    bool   `json:"isNewUser"`
}
