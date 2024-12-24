package shop_phone_request

type EditRequest struct {
	PhoneID     uint   `form:"phoneId"`
	ShopID      uint   `form:"shopId"`
	PhoneNumber string `form:"phoneNumber"`
}
