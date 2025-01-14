package shop_resource

type ShopResource struct {
	ID         uint    `json:"id"`
	PartnerId  uint    `json:"partnerId"`
	Name       string  `json:"name"`
	PictureUrl *string `json:"pictureUrl"`

	Location         ShopLocationResource        `json:"location"`
	PhoneNumbers     *[]ShopPhoneNumberResource  `json:"phoneNumbers"`
	ShopWorkingHours *[]ShopWorkingHoursResource `json:"workingHours"`
	Pictures         *[]ShopPictureResource      `json:"pictures"`
}
