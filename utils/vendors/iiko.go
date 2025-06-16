package vendor_utils

import (
	"github.com/hoopla/hoopla-api/internal/model"
	"time"
)

type Iiko struct {
	VendorID    string
	VendorKey   string
	AccessToken string
}

func (i *Iiko) SetAccessToken(accessToken string) {
	i.AccessToken = accessToken
}

func (i *Iiko) GetAccessToken() (string, time.Time, error) {
	return "", time.Now(), nil
}

func (i *Iiko) CreateOrder(
	partnerDrink *model.PartnerDrinkModel,
	shop *model.ShopModel,
	partner *model.PartnerModel,
	userOrder *model.UserOrderModel,
	phoneNumber string,
) (string, string, error) {
	return "preparing", "1231231", nil
}
