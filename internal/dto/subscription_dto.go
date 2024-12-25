package dto

type SubscriptionDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	CoffeeLimit int    `json:"coffe_limit"`
	Interval    int    `json:"interval"`
}