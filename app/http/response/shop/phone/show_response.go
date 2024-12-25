package shop_phone_response

type ShowResponse struct {
	ID          int     `json:"id"`
	ShopID      int     `json:"shopId"`
	PhoneNumber string  `json:"phoneNumber"`
}