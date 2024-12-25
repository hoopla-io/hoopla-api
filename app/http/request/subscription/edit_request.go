package subscription_request

type EditRequest struct {
	SubscriptionID uint   `form:"subscriptionId"`
	Name           string `form:"name"`
	CoffeeLimit    int    `form:"coffeeLimit"`
	Interval       int    `form:"interval"`
	Period         int    `form:"period"`
}
