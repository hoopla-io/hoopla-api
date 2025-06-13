package vendor_utils

import (
	"time"
)

type Iiko struct {
	VendorID  string
	VendorKey string
}

func (i Iiko) GetAccessToken() (string, time.Time, error) {
	return "", time.Now(), nil
}

func (i Iiko) CreateOrder() (string, error) {
	return "pending", nil
}
