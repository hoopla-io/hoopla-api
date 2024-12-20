package shop_request

type CreateShopWorkTimeRequest struct {
	ShopID      int    `form:"companyId"`
	DayRange    string `form:"dayRange"`
	OpeningTime string `form:"openingTime"`
	ClosingTime string `form:"closingTime"`
}