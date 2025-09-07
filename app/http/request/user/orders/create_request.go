package user_orders_request

type CreateRequest struct {
	ShopID        uint   `json:"shopId"`
	DrinkID       uint   `json:"drinkId"`
	VendorAddOnId string `json:"vendorAddOnId" binding:"omitempty"`
}
