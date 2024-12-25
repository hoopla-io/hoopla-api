package shop_worktime_request

type StoreRequest struct {
	ShopID      int    `form:"shopId"`
	DayRange    string `form:"dayRange"`
	OpeningTime string `form:"openingTime"`
	ClosingTime string `form:"closingTime"`
}