package shop_response

type ShowResponse struct {
	ID          int     `json:"id"`
	CompanyID   int     `json:"companyId"`
	Name        string  `json:"name"`
	Location    string  `json:"location"`
	ImageUrl    string  `json:"imageUrl"`
}