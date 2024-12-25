package repository

import (
	"context"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	AuthRepository
	ImageRepository
	CompanyRepository
	CompanySocialRepository
	ShopRepository
	ShopWorkTimeRepository
	ShopPhoneRepository
	ShopSocialRepository
	SubscriptionRepository
	UserSubscriptionRepository
	CoffeeRepository
	ShopCoffeeRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		AuthRepository:             NewAuthRepository(db),
		ImageRepository:            NewImageRepository(db),
		CompanyRepository:          NewCompanyRepository(db),
		ShopRepository:             NewShopRepository(db),
		ShopWorkTimeRepository:     NewShopWorktimeRepository(db),
		ShopPhoneRepository:        NewShopPhoneRepository(db),
		ShopSocialRepository:       NewShopSocialRepository(db),
		SubscriptionRepository:     NewSubscriptionRepository(db),
		UserSubscriptionRepository: NewUserSubscriptionRepository(db),
		CompanySocialRepository:    NewCompanySocialRepository(db),
		CoffeeRepository:           NewCoffeeRepository(db),
		ShopCoffeeRepository:       NewShopCoffeeRepository(db),
	}
}

type AuthRepository interface {
	GetByPhoneNumber(phoneNumber string) (*dto.UserDTO, error)
	CreateUser(data dto.UserDTO) (*dto.UserDTO, error)
}

type ImageRepository interface {
	GetImageById(id uint) (*dto.ImageDTO, error)
	CreateImage(data dto.ImageDTO) (int, error)
}

type CompanyRepository interface {
	Store(data dto.CompanyDTO) (uint, error)
	GetById(categoryId uint) (dto.CompanyDTO, error)
	List() ([]dto.CompanyDTO, error)
	Edit(data dto.CompanyDTO) (uint, error)
}

type CompanySocialRepository interface {
	Store(data dto.CompanySocialDTO) (uint, error)
	GetById(socialId uint) (dto.CompanySocialDTO, error)
	GetListByCompanyId(companyId uint) ([]dto.CompanySocialDTO, error)
	Edit(data dto.CompanySocialDTO) (uint, error)
}

type ShopRepository interface {
	GetByCompanyId(companyId uint) ([]dto.ShopDTO, error)
	Store(data dto.ShopDTO) (uint, error)
	GetById(shopId uint) (dto.ShopDTO, error)
	List() ([]dto.ShopDTO, error)
	Edit(data dto.ShopDTO) (uint, error)
}

type ShopWorkTimeRepository interface {
	Store(data dto.ShopWorktimeDTO) (uint, error)
	GetById(worktimeId uint) (dto.ShopWorktimeDTO, error)
	GetListByShopId(shopId uint) ([]dto.ShopWorktimeDTO, error)
	Edit(data dto.ShopWorktimeDTO) (uint, error)
}

type ShopPhoneRepository interface {
	Store(data dto.ShopPhoneDTO) (uint, error)
	GetById(phoneId uint) (dto.ShopPhoneDTO, error)
	GetListByShopId(shopId uint) ([]dto.ShopPhoneDTO, error)
	Edit(data dto.ShopPhoneDTO) (uint, error)
}

type ShopSocialRepository interface {
	CreateShopSocial(data dto.ShopSocialDTO) (*dto.ShopSocialDTO, error)
	GetShopSocialByShopId(shopId uint) ([]dto.ShopSocialDTO, error)
}

type CoffeeRepository interface {
	Store(data dto.CoffeeDTO) (uint, error)
	GetById(coffeeId uint) (dto.CoffeeDTO, error)
	List() ([]dto.CoffeeDTO, error)
	Edit(data dto.CoffeeDTO) (uint, error)
}

type ShopCoffeeRepository interface {
	Store(shopId uint, coffeeId uint) error
	GetListByShopId(shopId uint) ([]dto.ShopCoffeeDTO, error)
	DeleteByShopId(shopId uint) error
}

type SubscriptionRepository interface {
	GetAllSubscriptions(ctx context.Context) ([]model.SubscriptionModel, error) // Получить все подписки
	GetSubscriptionByID(ctx context.Context, id uint) (*model.SubscriptionModel, error)
	Store(data dto.SubscriptionDTO) (uint, error)
	GetById(coffeeId uint) (dto.SubscriptionDTO, error)
	List() ([]dto.SubscriptionDTO, error)
	Edit(data dto.SubscriptionDTO) (uint, error)
}
