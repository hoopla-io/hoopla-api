package user_orders_request

type ValidateOrderRequest struct {
	ShopID  uint `json:"shopId"`
	DrinkID uint `json:"drinkId"`
}
