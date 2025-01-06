package partner_resource

type ShopsCollections struct {
	ShopID uint   `json:"shopId"`
	Name   string `json:"name"`

	Location     *ShopLocationResource         `json:"location"`
	PhoneNumbers *[]ShopPhoneNumbersCollection `json:"phoneNumbers"`
}

type ShopLocationResource struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ShopPhoneNumbersCollection struct {
	PhoneNumber string `json:"phoneNumber"`
}
