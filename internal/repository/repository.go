package repository

import (
	"github.com/qahvazor/qahvazor/internal/dto"
	"gorm.io/gorm"
)

type Repository struct {
	AuthRepository
	ImageRepository
	CompanyRepository
	ShopRepository
	ShopWorkTimeRepository
	ShopPhoneRepository
	ShopSocialRepository
	SubscriptionRepository
	UserSubscriptionRepository
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
	}
}

type AuthRepository interface {
	GetByPhoneNumber(phoneNumber string) (*dto.UserDTO, error)
	CreateUser(data dto.UserDTO) (*dto.UserDTO, error)
}

type ImageRepository interface {
	GetImageById(id uint) (*dto.ImageDTO, error)
	CreateImage(data dto.ImageDTO) (uint, error)
}

type CompanyRepository interface {
	CreateCompany(data dto.CompanyDTO) (*dto.CompanyDTO, error)
	GetCompanyById(id uint) (*dto.CompanyDTO, error)
	GetList() ([]dto.CompanyDTO, error)
}

type ShopRepository interface {
	GetShopsByCompanyId(companyId uint) ([]dto.ShopDTO, error)
	CreateShop(data dto.ShopDTO) (*dto.ShopDTO, error)
}

type ShopWorkTimeRepository interface {
	CreateShopWorkTime(data dto.ShopWorkTimeDTO) (*dto.ShopWorkTimeDTO, error)
	GetShopWorkTimeByShopId(shopId uint) ([]dto.ShopWorkTimeDTO, error)
}

type ShopPhoneRepository interface {
	CreateShopPhone(data dto.ShopPhoneDTO) (*dto.ShopPhoneDTO, error)
	GetShopPhoneByShopId(shopId uint) ([]dto.ShopPhoneDTO, error)
}

type ShopSocialRepository interface {
	CreateShopSocial(data dto.ShopSocialDTO) (*dto.ShopSocialDTO, error)
	GetShopSocialByShopId(shopId uint) ([]dto.ShopSocialDTO, error)
}
