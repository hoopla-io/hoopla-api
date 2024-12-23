package shop_request

type CreateShopPhoneRequest struct {
	ShopID      int    `form:"companyId"`
	PhoneNumber string `form:"phoneNumber"`
}