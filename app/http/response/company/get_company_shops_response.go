package company_response

type GetCompanyShopsResponse struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageUrl    string   `json:"imageUrl"`
	Shops      []Shop    `json:"shops"`
}

type Shop struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	ImageUrl string `json:"imageUrl"`
}
