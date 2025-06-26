package user_order_resource

import "time"

type UserOrderResource struct {
	ID              uint      `json:"id"`
	PartnerName     string    `json:"partnerName"`
	ShopName        string    `json:"shopName"`
	PurchasedAt     time.Time `json:"purchasedAt"`
	PurchasedAtUnix int64     `json:"purchasedAtUnix"`
	DrinkName       string    `json:"drinkName"`
	OrderStatus     string    `json:"orderStatus"`
}
