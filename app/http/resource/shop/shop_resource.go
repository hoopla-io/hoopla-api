package shop_resource

import partner_resource "github.com/hoopla/hoopla-api/app/http/resource/partner"

type ShopResource struct {
	ID              uint    `json:"id"`
	PartnerId       uint    `json:"partnerId"`
	Name            string  `json:"name"`
	PictureUrl      *string `json:"pictureUrl"`
	CanAcceptOrders bool    `json:"canAcceptOrders"`

	Location         ShopLocationResource        `json:"location"`
	PhoneNumbers     *[]ShopPhoneNumberResource  `json:"phoneNumbers"`
	ShopWorkingHours *[]ShopWorkingHoursResource `json:"workingHours"`
	Pictures         *[]ShopPictureResource      `json:"pictures"`

	PartnerUrls   *[]partner_resource.PartnerUrlsCollection `json:"urls"`
	PartnerDrinks *[]partner_resource.DrinksCollection      `json:"drinks"`
}
