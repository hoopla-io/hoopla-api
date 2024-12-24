package subscription_response

type ShowResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CoffeeLimit int    `json:"coffeeLimit"`
	Interval    int    `json:"interval"`
}