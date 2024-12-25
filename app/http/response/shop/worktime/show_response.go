package shop_worktime_response

type ShowResponse struct {
	ID          int     `json:"id"`
	ShopID      int     `json:"shopId"`
	DayRange    string  `json:"dayRange"`
	OpeningTime string  `json:"openingTime"`
	ClosingTime string  `json:"closingTime"`
}