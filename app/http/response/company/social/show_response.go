package company_social_response

type ShowResponse struct {
	ID          int     `json:"id"`
	CompanyID   int     `json:"companyId"`
	Platform    string  `json:"platform"`
	Url         string  `json:"url"`
}