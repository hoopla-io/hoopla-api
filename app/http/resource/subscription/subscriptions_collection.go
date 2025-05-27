package subscription_resource

type SubscriptionsCollection struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Days     uint    `json:"days"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
	CupsADay uint   `json:"cupsADay"`

	Features *[]FeaturesCollection `json:"features"`

	WeekDays []string `json:"weekDays"`
}
