package subscription_response

type ListResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CoffeeLimit int    `json:"coffeeLimit"`
	Interval    int    `json:"interval"`
	Period      int    `json:"period"`
}