package user_order_resource

type DrinksStatCollection struct {
	Available uint  `json:"available"`
	Used      int64 `json:"used"`
}
