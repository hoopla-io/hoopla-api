package company_social_response

type ListResponse struct {
	ID          int     `json:"id"`
	CompanyID   int     `json:"companyId"`
	Platform    string  `json:"platform"`
	Url         string  `json:"url"`
}