package subscription_resource

type SubscriptionsCollection struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Days     uint    `json:"days"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`

	Features *[]FeaturesCollection `json:"features"`
}
