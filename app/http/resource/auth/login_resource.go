package auth_resource

type LoginResource struct {
	UserID      uint        `json:"userId"`
	PhoneNumber string      `json:"phoneNumber"`
	IsNewUser   bool        `json:"isNewUser"`
	Jwt         JwtResource `json:"jwt"`
}
