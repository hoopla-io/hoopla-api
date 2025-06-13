package vendor_utils

import "time"

type Poster struct {
	VendorID  string
	VendorKey string
}

func (i Poster) GetAccessToken() (string, time.Time, error) {
	return "", time.Now(), nil
}

func (i Poster) CreateOrder() (string, error) {
	return "pending", nil
}
