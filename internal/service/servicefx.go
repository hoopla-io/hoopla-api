package service

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewAuthService),
	fx.Provide(NewCoffeeService),
	fx.Provide(NewCompanyService),
	fx.Provide(NewShopService),
	fx.Provide(NewSubscriptionService),
	fx.Provide(NewUserSubscriptionService),
)

//func NewService(repo *repository.Repository) *Service {
//	return &Service{
//		UserSubscriptionService: NewUserSubscriptionService(repo),
//		AuthService:             NewAuthService(repo.AuthRepository),
//		CompanyService: NewCompanyService(
//			repo.CompanyRepository,
//			repo.ImageRepository,
//			repo.CompanySocialRepository,
//			repo.ShopRepository,
//		),
//		ShopService: NewShopService(
//			repo.ShopRepository,
//			repo.ImageRepository,
//			repo.ShopWorkTimeRepository,
//			repo.ShopPhoneRepository,
//			repo.ShopCoffeeRepository,
//			repo.CompanySocialRepository,
//		),
//		CoffeeService: NewCoffeeService(
//			repo.CoffeeRepository,
//			repo.ImageRepository,
//		),
//		SubscriptionService: NewSubscriptionService(
//			repo.SubscriptionRepository,
//		),
//	}
//}
