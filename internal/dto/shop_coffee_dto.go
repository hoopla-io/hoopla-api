package dto

type ShopCoffeeDTO struct {
	ID          uint   `json:"id"`
	Name        *string `json:"name"`
	ShopID      int    `json:"shop_id"`
	CoffeeID    int    `json:"coffee_id"`
	ImageID     *int   `json:"image_id"`
}