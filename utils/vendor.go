package utils

import (
	"github.com/hoopla/hoopla-api/internal/model"
	vendor_utils "github.com/hoopla/hoopla-api/utils/vendors"
	"time"
)

type VendorInterface interface {
	SetAccessToken(string)
	GetAccessToken() (string, time.Time, error)
	CreateOrder(
		partnerDrink *model.PartnerDrinkModel,
		shop *model.ShopModel,
		partner *model.PartnerModel,
		userOrder *model.UserOrderModel,
		phoneNumber string,
	) (string, string, error)
}

type Vendor struct {
	VendorInterface VendorInterface
}

func (v *Vendor) Init(vendor string, vendorId string, vendorKey string) {
	if vendor == "iiko" {
		v.VendorInterface = &vendor_utils.Iiko{
			VendorID:    vendorId,
			VendorKey:   vendorKey,
			AccessToken: "",
		}
	} else {
		v.VendorInterface = &vendor_utils.Poster{
			VendorID:    vendorId,
			VendorKey:   vendorKey,
			AccessToken: "",
		}
	}
}
