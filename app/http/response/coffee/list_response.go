package coffee_response

type ListResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"imageUrl"`
}