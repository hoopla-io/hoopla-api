package partner_resource

type PartnersCollection struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	LogoUrl *string `json:"logoUrl"`
}
