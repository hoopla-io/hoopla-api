package shop_resource

type ShopsCollections struct {
	ShopID     uint    `json:"shopId"`
	PartnerID  uint    `json:"partnerId"`
	Name       string  `json:"name"`
	PictureURL *string `json:"pictureUrl"`
	Distance   float64 `json:"distance"`

	Location ShopLocationResource   `json:"location"`
	Modules  *[]ShopModulesResource `json:"modules"`
}
