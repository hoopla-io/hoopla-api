package shop_worktime_request

type EditRequest struct {
	WorktimeID  uint   `form:"worktimeId"`
	ShopID      uint   `form:"shopId"`
	DayRange    string `form:"dayRange"`
	OpeningTime string `form:"openingTime"`
	ClosingTime string `form:"closingTime"`
}
