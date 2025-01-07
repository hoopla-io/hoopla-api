package repository

import (
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type ShopRepository interface {
	GetPartnerShops(partnerId uint) (*[]model.ShopModel, error)
	ShopDetailById(shopId uint) (*model.ShopModel, error)
}

type ShopRepositoryImpl struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &ShopRepositoryImpl{
		db: db,
	}
}

func (r *ShopRepositoryImpl) GetPartnerShops(partnerId uint) (*[]model.ShopModel, error) {
	var shops []model.ShopModel
	err := r.db.Where("partner_id = ?", partnerId).Preload("Attributes").Find(&shops).Error
	if err != nil {
		return nil, err
	}

	return &shops, nil
}

func (r *ShopRepositoryImpl) ShopDetailById(shopId uint) (*model.ShopModel, error) {
	var shop model.ShopModel
	err := r.db.Where("id = ?", shopId).Preload("Attributes").First(&shop).Error
	if err != nil {
		return nil, err
	}

	return &shop, nil
}
