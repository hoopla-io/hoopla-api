package subscription_request

type StoreRequest struct {
	Name        string `form:"name"`
	CoffeeLimit int    `form:"coffeeLimit"`
	Interval    int    `form:"interval"`
}
