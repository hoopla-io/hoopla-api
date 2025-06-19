package utils

import (
	"github.com/hoopla/hoopla-api/internal/model"
	vendor_utils "github.com/hoopla/hoopla-api/utils/vendors"
	"time"
)

// created - when order is created
// pending - when order is sent to vendor partner
// preparing - when order is accepted
// completed - when order is completed
// cancelled - when order is cancelled
type VendorInterface interface {
	SetAccessToken(string)
	GetAccessToken() (string, time.Time, error)
	CreateOrder(
		partnerDrink *model.PartnerDrinkModel,
		shop *model.ShopModel,
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
