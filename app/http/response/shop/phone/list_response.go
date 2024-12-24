package shop_phone_response

type ListResponse struct {
	ID          int     `json:"id"`
	ShopID      int     `json:"shopId"`
	PhoneNumber string  `json:"phoneNumber"`
}