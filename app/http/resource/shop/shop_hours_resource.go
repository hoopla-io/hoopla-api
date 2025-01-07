package shop_resource

type ShopWorkingHoursResource struct {
	WeekDay string `json:"weekDay"`
	OpenAt  string `json:"openAt"`
	CloseAt string `json:"closeAt"`
}
