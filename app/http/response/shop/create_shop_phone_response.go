package shop_response

type CreateShopPhoneResponse struct {
	ID          int    `json:"id"`
	ShopID      int    `json:"shopId"`
	PhoneNumber string `json:"phoneNumber"`
}