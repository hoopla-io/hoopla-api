package repository

import (
	"errors"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type ShopRepositoryImpl struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &ShopRepositoryImpl{db: db}
}

func (r *ShopRepositoryImpl) GetShopsByCompanyId(companyId uint) ([]dto.ShopDTO, error) {
	var shops []dto.ShopDTO
	err := r.db.Model(&model.ShopModel{}).Where("company_id = ?", companyId).Scan(&shops).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return shops, nil
} 

func (r *ShopRepositoryImpl) CreateShop(data dto.ShopDTO) (*dto.ShopDTO, error) {
	var createShop dto.ShopDTO

	query := r.db.Create(&model.ShopModel{
		ImageID: data.ImageID,
		CompanyID: data.CompanyID,
		Name: data.Name,
		Location: data.Location,
	}).Scan(&createShop)
	if query.Error != nil {
		return nil, query.Error
	}

	return &createShop, nil
}
