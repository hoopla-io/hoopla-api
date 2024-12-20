package dto

type CompanyDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	ImageID     int    `json:"image_id"`
	Description string `json:"description"`
}