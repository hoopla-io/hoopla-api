package dto

type ShopSocialDTO struct {
	ID       uint   `json:"id"`
	ShopID   int    `json:"shop_id"`
	Platform string `json:"platform"`
	Url      string `json:"url"`
}