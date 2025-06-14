package vendor_utils

import (
	"errors"
	"fmt"
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

func (i Poster) GetAccessToken() (string, time.Time, error) {
	formData := url.Values{
		"account":            {i.VendorID},
		"code":               {i.VendorKey},
		"application_id":     {os.Getenv("POSTER_APPLICATION_ID")},
		"application_secret": {os.Getenv("POSTER_APPLICATION_SECRET")},
		"grant_type":         {"authorization_code"},
		"redirect_uri":       {"https://api.hoopla.uz/api/vendors/poster/oauth"},
	}

	req := pkg.Requests{}
	statusCode, data, err := req.PostForm(
		fmt.Sprintf("https://%s.joinposter.com/api/v2/auth/access_token", i.VendorID),
		&formData,
	)
	if err != nil {
		return "", time.Time{}, err
	}

	if statusCode != 200 {
		return "", time.Time{}, errors.New(data["error_message"].(string))
	}

	return data["access_token"].(string), time.Now().AddDate(10, 0, 0), nil
}

func (i Poster) CreateOrder() (string, error) {
	return "pending", nil
}

func (i Poster) GetOrderStatus(orderID int64) (string, error) {
	req := pkg.Requests{}
	statusCode, data, err := req.Get(
		fmt.Sprintf(
			"https://joinposter.com/api/incomingOrders.getIncomingOrder?incoming_order_id=%d&token=%s",
			orderID,
			i.AccessToken,
		),
	)

	if err != nil {
		return "pending", err
	}

	fmt.Println(data, statusCode)

	return "pending", nil
}
