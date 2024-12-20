package shop_response

type CreateShopResponse struct {
	ID          int     `json:"id"`
	CompanyID   int     `json:"companyId"`
	Name        string  `json:"name"`
	Location    string  `json:"location"`
}