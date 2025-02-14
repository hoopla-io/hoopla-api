package user_resource

type UserBaseResource struct {
	UserID       uint                  `json:"userId"`
	PhoneNumber  string                `json:"phoneNumber"`
	Name         string                `json:"name"`
	Balance      float64               `json:"balance"`
	Currency     string                `json:"currency"`
	Subscription *SubscriptionResource `json:"subscription"`
}

type SubscriptionResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	EndDate     string `json:"endDate"`
	EndDateUnix int64  `json:"endDateUnix"`
}
