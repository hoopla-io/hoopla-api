package vendor_utils

import (
	"errors"
	"fmt"
	"github.com/hoopla/hoopla-api/internal/model"
	"github.com/hoopla/hoopla-api/pkg"
	"net/url"
	"os"
	"time"
)

type Poster struct {
	VendorID    string
	VendorKey   string
	AccessToken string
}

func (p *Poster) SetAccessToken(accessToken string) {
	p.AccessToken = accessToken
}

func (p *Poster) GetAccessToken() (string, time.Time, error) {
	formData := url.Values{
		"account":            {p.VendorID},
		"code":               {p.VendorKey},
		"application_id":     {os.Getenv("POSTER_APPLICATION_ID")},
		"application_secret": {os.Getenv("POSTER_APPLICATION_SECRET")},
		"grant_type":         {"authorization_code"},
		"redirect_uri":       {"https://api.hoopla.uz/api/vendors/poster/oauth"},
	}

	req := pkg.Requests{}
	statusCode, data, err := req.PostForm(
		fmt.Sprintf("https://%s.joinposter.com/api/v2/auth/access_token", p.VendorID),
		&formData,
	)
	if err != nil {
		return "", time.Now(), err
	}

	if statusCode != 200 {
		return "", time.Now(), errors.New(data["error_message"].(string))
	}

	return data["access_token"].(string), time.Now().AddDate(10, 0, 0), nil
}

func (p *Poster) CreateOrder(
	partnerDrink *model.PartnerDrinkModel,
	shop *model.ShopModel,
	partner *model.PartnerModel,
	userOrder *model.UserOrderModel,
	phoneNumber string,
) (string, string, error) {
	formData := url.Values{
		"spot_id":                 {shop.VendorTerminalID},
		"phone":                   {fmt.Sprintf("+%s", phoneNumber)},
		"products[0][product_id]": {partnerDrink.VendorProductID},
		"products[0][count]":      {"1"},
		"products[0][price]":      {fmt.Sprintf("%f", partnerDrink.ProductPrice)},
		"service_mode":            {"2"},
		"comment":                 {"hoopla"},
		"payment[type]":           {"1"},
		"payment[sum]":            {fmt.Sprintf("%f", partnerDrink.ProductPrice)},
	}

	req := pkg.Requests{}
	statusCode, data, err := req.PostForm(
		fmt.Sprintf("https://joinposter.com/api/incomingOrders.createIncomingOrder?token=%s", p.AccessToken),
		&formData,
	)
	if err != nil {
		return "error", "", err
	}

	if statusCode != 200 {
		return "error", "", errors.New(data["error_message"].(string))
	}

	response := data["response"].(map[string]interface{})

	return "pending", fmt.Sprintf("%.0f", response["incoming_order_id"]), nil
}

func (p *Poster) GetOrderStatus(orderID int64) (string, error) {
	req := pkg.Requests{}
	statusCode, data, err := req.Get(
		fmt.Sprintf(
			"https://joinposter.com/api/incomingOrders.getIncomingOrder?incoming_order_id=%d&token=%s",
			orderID,
			p.AccessToken,
		),
	)

	if err != nil {
		return "pending", err
	}

	if data["error"] != nil || statusCode != 200 {
		return "pending", errors.New(data["message"].(string))
	}

	response, ok := data["response"].(map[string]interface{})
	if !ok {
		return "pending", errors.New("could not parse response")
	}

	if response["status"].(float64) == 1 {
		return "completed", nil
	}

	if response["status"].(float64) == 0 {
		return "preparing", nil
	}

	return "canceled", nil
}
