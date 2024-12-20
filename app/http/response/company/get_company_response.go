package company_response

type GetCompanyResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"imageUrl"`
	Shops 		[]Shops `json:"shops"`
}

type Shops struct {
	ShopID    int    `json:"shopId"`
	CompanyID int    `json:"companyId"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	ImageUrl  string `json:"imageUrl"`
}