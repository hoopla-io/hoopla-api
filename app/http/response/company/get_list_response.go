package company_response

type GetListResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"imageUrl"`
}