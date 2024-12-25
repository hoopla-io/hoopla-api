package shop_phone_request

type StoreRequest struct {
	ShopID      int    `form:"shopId"`
	PhoneNumber string `form:"phoneNumber"`
}