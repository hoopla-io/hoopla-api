package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	user_pay_request "github.com/hoopla/hoopla-api/app/http/request/user/pay"
	user_pay_resource "github.com/hoopla/hoopla-api/app/http/resource/user/pay"
	"github.com/hoopla/hoopla-api/utils"
	"net/http"
)

type PayService interface {
	Services(data user_pay_request.ServicesRequest) (*[]user_pay_resource.ServicesCollection, int, error)
	TopUp(data user_pay_request.TopUpRequest, user *utils.UserHelper) (*user_pay_resource.TopUpResource, int, error)
}

type PayServiceImpl struct {
}

func NewPayService() PayService {
	return &PayServiceImpl{}
}

func (s *PayServiceImpl) Services(data user_pay_request.ServicesRequest) (*[]user_pay_resource.ServicesCollection, int, error) {
	var payServices []user_pay_resource.ServicesCollection

	payServices = append(payServices, user_pay_resource.ServicesCollection{
		ID:      1,
		Name:    "Payme",
		LogoURL: "https://files.itv.uz/uploads/payments/types/2019/08/16//2a221daf4bf444b7b3c464d820ecb965.jpg",
	})

	return &payServices, http.StatusOK, nil
}

func (s *PayServiceImpl) TopUp(data user_pay_request.TopUpRequest, user *utils.UserHelper) (*user_pay_resource.TopUpResource, int, error) {
	var topUpResource user_pay_resource.TopUpResource
	var checkoutUrl string
	if data.ID == 1 {
		checkoutUrl = fmt.Sprintf("m=67c6d714e4b4003392f2e809;ac.phoneNumber=%s;a=%f;c=%s",
			user.PhoneNumber,
			data.Amount*100,
			"https://hoopla.uz",
		)
		checkoutUrl = fmt.Sprintf("https://checkout.paycom.uz/%s", base64.StdEncoding.EncodeToString([]byte(checkoutUrl)))
	} else {
		return nil, 500, errors.New("invalid pay-service id")
	}

	topUpResource = user_pay_resource.TopUpResource{
		CheckoutUrl: checkoutUrl,
	}
	return &topUpResource, http.StatusOK, nil
}
