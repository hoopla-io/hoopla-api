package shop_request

type CreateShopSocialRequest struct {
	ShopID   int    `form:"companyId"`
	Platform string `form:"platform"`
	Url      string `form:"url"`
}