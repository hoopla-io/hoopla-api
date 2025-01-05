package user_resource

type UserBaseResource struct {
	UserID      uint   `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}
