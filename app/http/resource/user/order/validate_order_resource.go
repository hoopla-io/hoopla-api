package user_order_resource

import "time"

type UserValidateOrderResource struct {
	Partner         UserValidateOrderPartnerStruct `json:"partner"`
	Shop            UserValidateOrderShopStruct    `json:"shop"`
	Drink           UserValidateOrderDrinkStruct   `json:"drink"`
	ValidatedAt     time.Time                      `json:"validatedAt"`
	ValidatedAtUnix int64                          `json:"validatedAtUnix"`
	AddOns          *[]OrderAddOnsCollection       `json:"addOns"`
}

type OrderAddOnsCollection struct {
	VendorAddOnId string `json:"vendorAddOnId"`
	AddOn         string `json:"addOn"`
}

type UserValidateOrderPartnerStruct struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserValidateOrderShopStruct struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserValidateOrderDrinkStruct struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
