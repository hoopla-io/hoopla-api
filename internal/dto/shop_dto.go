package dto

type ShopDTO struct {
	ID          uint   `json:"id"`
	CompanyID   int    `json:"company_id"`
	ImageID     int    `json:"image_id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
}