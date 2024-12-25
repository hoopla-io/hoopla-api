package shop_worktime_response

type ListResponse struct {
	ID          int     `json:"id"`
	ShopID      int     `json:"shopId"`
	DayRange    string  `json:"dayRange"`
	OpeningTime string  `json:"openingTime"`
	ClosingTime string  `json:"closingTime"`
}