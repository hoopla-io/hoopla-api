package company_social_request

type EditRequest struct {
	SocialID    int     `form:"socialId"`
	CompanyID   int     `form:"companyId"`
	Platform    string  `form:"platform"`
	Url         string  `form:"url"`
}
