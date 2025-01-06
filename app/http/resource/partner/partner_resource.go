package partner_resource

type PartnerResource struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	LogoUrl     *string `json:"logoUrl"`

	PartnerPhoneNumbers []PartnerPhoneNumbersCollection `json:"phoneNumbers"`
	PartnerUrls         []PartnerUrlsCollection         `json:"urls"`
}

type PartnerPhoneNumbersCollection struct {
	PhoneNumber string `json:"phoneNumber"`
}

type PartnerUrlsCollection struct {
	UrlType string `json:"urlType"`
	Url     string `json:"url"`
}
