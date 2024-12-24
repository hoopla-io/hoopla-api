package service

import (
	auth_request "github.com/qahvazor/qahvazor/app/http/request/auth"
	coffee_request "github.com/qahvazor/qahvazor/app/http/request/coffee"
	company_request "github.com/qahvazor/qahvazor/app/http/request/company"
	company_social_request "github.com/qahvazor/qahvazor/app/http/request/company/social"
	shop_request "github.com/qahvazor/qahvazor/app/http/request/shop"
	shop_phone_request "github.com/qahvazor/qahvazor/app/http/request/shop/phone"
	shop_worktime_request "github.com/qahvazor/qahvazor/app/http/request/shop/worktime"
	subscription_request "github.com/qahvazor/qahvazor/app/http/request/subscription"
	"github.com/qahvazor/qahvazor/internal/repository"
)

type Service struct {
	AuthService
	CompanyService
	ShopService
	CoffeeService
	SubscriptionService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repo.AuthRepository),
		CompanyService: NewCompanyService(
			repo.CompanyRepository,
			repo.ImageRepository,
			repo.CompanySocialRepository,
		),
		ShopService: NewShopService(
			repo.ShopRepository,
			repo.ImageRepository,
			repo.ShopWorktimeRepository,
			repo.ShopPhoneRepository,
		),
		CoffeeService: NewCoffeeService(
			repo.CoffeeRepository,
			repo.ImageRepository,
		),
		SubscriptionService: NewSubscriptionService(
			repo.SubscriptionRepository,
		),
	}
}

type AuthService interface {
	Login(data auth_request.LoginRequest) (interface{}, error)
    ConfirmSms(data auth_request.ConfirmSmsRequest) (interface{}, error)
	ResendSms(data auth_request.ResendSmsRequest) (interface{}, error)
}

type CompanyService interface {
	Store(data company_request.StoreRequest) (interface{}, error)
	Show(companyId uint) (interface{}, error)
	List() (interface{}, error)
	Edit(data company_request.EditRequest) error 
	StoreCompanySocial(data company_social_request.StoreRequest) (interface{}, error)
	ShowCompanySocial(socialId uint) (interface{}, error)
	ListCompanySocials(companyId uint) (interface{}, error)
	EditCompanySocial(data company_social_request.EditRequest) error
}

type ShopService interface {
	Store(data shop_request.StoreRequest) (interface{}, error)
	Show(shopId uint) (interface{}, error)
	List() (interface{}, error)
	Edit(data shop_request.EditRequest) error
	StoreShopWorktime(data shop_worktime_request.StoreRequest) (interface{}, error)
	ShowWorktime(worktimeId uint) (interface{}, error)
	ListShopWorktimes(shopId uint) (interface{}, error)
	EditShopWorktime(data shop_worktime_request.EditRequest) error
	StoreShopPhone(data shop_phone_request.StoreRequest) (interface{}, error)
	ShowShopPhone(phoneId uint) (interface{}, error)
	ListShopPhones(shopId uint) (interface{}, error)
	EditShopPhone(data shop_phone_request.EditRequest) error
}

type CoffeeService interface {
	Store(data coffee_request.StoreRequest) (interface{}, error)
	Show(coffeeId uint) (interface{}, error)
	List() (interface{}, error)
	Edit(data coffee_request.EditRequest) error 
}

type SubscriptionService interface {
	Store(data subscription_request.StoreRequest) (interface{}, error)
	Show(coffeeId uint) (interface{}, error)
	List() (interface{}, error)
	Edit(data subscription_request.EditRequest) error 
}
