package dto

type CoffeeDTO struct {
	ID          uint   `json:"id"`
	ImageID     int    `json:"image_id"`
	Name        string `json:"name"`
}