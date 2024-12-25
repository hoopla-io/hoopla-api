package dto

type CompanySocialDTO struct {
	ID        uint   `json:"id"`
	CompanyID int    `json:"company_id"`
	Platform  string `json:"platform"`
	Url       string `json:"url"`
}