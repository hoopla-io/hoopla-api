package shop_response

type CreateShopSocialResponse struct {
	ID       int    `json:"id"`
	ShopID   int    `json:"shopId"`
	Platform string `json:"platform"`
	Url      string `json:"url"`
}