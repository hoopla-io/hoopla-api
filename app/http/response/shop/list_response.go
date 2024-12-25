package shop_response

type ListResponse struct {
	ID          int     `json:"id"`
	CompanyID   int     `json:"companyId"`
	Name        string  `json:"name"`
	Location    string  `json:"location"`
	ImageUrl    string  `json:"imageUrl"`
}