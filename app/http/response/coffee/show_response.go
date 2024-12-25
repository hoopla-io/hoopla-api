package coffee_response

type ShowResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"imageUrl"`
}