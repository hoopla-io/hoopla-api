package service

import (
	auth_request "github.com/qahvazor/qahvazor/app/http/request/auth"
	company_request "github.com/qahvazor/qahvazor/app/http/request/company"
	shop_request "github.com/qahvazor/qahvazor/app/http/request/shop"
	"github.com/qahvazor/qahvazor/internal/repository"
)

type Service struct {
	AuthService
	CompanyService
	ShopService
	SubscriptionService
	UserSubscriptionService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService:             NewAuthService(repo),
		CompanyService:          NewCompanyService(repo),
		ShopService:             NewShopService(repo),
		SubscriptionService:     NewSubscriptionService(repo),
		UserSubscriptionService: NewUserSubscriptionService(repo),
	}
}

type AuthService interface {
	Login(data auth_request.LoginRequest) (interface{}, error)
	ConfirmSms(data auth_request.ConfirmSmsRequest) (interface{}, error)
	ResendSms(data auth_request.ResendSmsRequest) (interface{}, error)
}

type CompanyService interface {
	CreateCompany(data company_request.CreateCompanyRequest) (interface{}, error)
	GetCompany(data company_request.GetCompanyRequest) (interface{}, error)
	GetList() (interface{}, error)
}

type ShopService interface {
	CreateShop(data shop_request.CreateShopRequest) (interface{}, error)
	CreateShopWorkTime(data shop_request.CreateShopWorkTimeRequest) (interface{}, error)
	CreateShopPhone(data shop_request.CreateShopPhoneRequest) (interface{}, error)
	CreateShopSocial(data shop_request.CreateShopSocialRequest) (interface{}, error)
}
