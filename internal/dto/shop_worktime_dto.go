package dto

type ShopWorkTimeDTO struct {
	ID          uint   `json:"id"`
	ShopID      int    `json:"shop_id"`
	DayRange    string `json:"day_range"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}