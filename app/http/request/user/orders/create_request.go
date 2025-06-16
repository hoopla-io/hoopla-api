package user_orders_request

type CreateRequest struct {
	ShopID  uint `json:"shop_id"`
	DrinkID uint `json:"drink_id"`
}
