package company_social_request

type StoreRequest struct {
	CompanyID   int     `form:"companyId"`
	Platform    string  `form:"platform"`
	Url         string  `form:"url"`
}
