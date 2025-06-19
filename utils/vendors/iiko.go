package vendor_utils

import (
	"errors"
	"fmt"
	"github.com/hoopla/hoopla-api/internal/model"
	"github.com/hoopla/hoopla-api/pkg"
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
	data := map[string]interface{}{
		"apiLogin": i.VendorKey,
	}

	req := pkg.Requests{}
	statusCode, res, err := req.Post("https://api-ru.iiko.services/api/1/access_token", data)
	if err != nil {
		return "", time.Now(), err
	}

	if statusCode != 200 {
		return "", time.Now(), errors.New(res["errorDescription"].(string))
	}

	accessToken := res["token"].(string)
	expiresAt := time.Now().Add(time.Hour)

	return accessToken, expiresAt, nil
}

func (i *Iiko) CreateOrder(
	partnerDrink *model.PartnerDrinkModel,
	shop *model.ShopModel,
	userOrder *model.UserOrderModel,
	phoneNumber string,
) (string, string, error) {

	data := map[string]interface{}{
		"organizationId":  i.VendorID,
		"terminalGroupId": shop.VendorTerminalID,
		"order": map[string]interface{}{
			"externalNumber":   fmt.Sprintf("hoopla-%d", userOrder.ID),
			"phone":            fmt.Sprintf("+%s", phoneNumber),
			"orderServiceType": "DeliveryByClient",
			"sourceKey":        "hoopla",
			"items": []map[string]interface{}{
				{
					"productId": partnerDrink.VendorProductID,
					"type":      "Product",
					"amount":    1,
					"price":     partnerDrink.ProductPrice,
				},
			},
		},
	}

	req := pkg.Requests{}
	req.Headers = map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", i.AccessToken),
	}

	statusCode, res, err := req.Post("https://api-ru.iiko.services/api/1/deliveries/create", data)
	if err != nil {
		return "error", "", errors.New("error creating order")
	}

	if statusCode != 200 {
		return "error", "", errors.New(res["errorDescription"].(string))
	}

	orderInfo := res["orderInfo"].(map[string]interface{})

	return "pending", orderInfo["id"].(string), nil
}
