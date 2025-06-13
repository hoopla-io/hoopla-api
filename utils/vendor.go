package utils

import (
	vendor_utils "github.com/hoopla/hoopla-api/utils/vendors"
	"time"
)

type VendorInterface interface {
	GetAccessToken() (string, time.Time, error)
	CreateOrder() (string, error)
}

type Vendor struct {
	VendorInterface VendorInterface
}

func (v *Vendor) Init(vendor string, vendorId string, vendorKey string) {
	if vendor == "iiko" {
		v.VendorInterface = vendor_utils.Iiko{
			VendorID:  vendorId,
			VendorKey: vendorKey,
		}
	} else {
		v.VendorInterface = vendor_utils.Poster{
			VendorID:  vendorId,
			VendorKey: vendorKey,
		}
	}
}
