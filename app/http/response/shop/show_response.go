package shop_response

type ShowResponse struct {
	ID          int      `json:"id"`
	CompanyID   int      `json:"companyId"`
	Name        string   `json:"name"`
	Location    string   `json:"location"`
	ImageUrl    string   `json:"imageUrl"`
	Coffees     []Coffee `json:"coffees"`
}

type Coffee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}
