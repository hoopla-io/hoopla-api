package dto

type ShopPhoneDTO struct {
	ID          uint   `json:"id"`
	ShopID      int    `json:"shop_id"`
	PhoneNumber string `json:"phone_number"`
}