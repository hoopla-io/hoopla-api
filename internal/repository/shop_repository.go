package repository

import (
	"github.com/hoopla/hoopla-api/internal/model"
	"gorm.io/gorm"
)

type ShopRepository interface {
	GetPartnerShops(partnerId uint) (*[]model.ShopModel, error)
	ShopDetailById(shopId uint) (*model.ShopModel, error)
	GetShopsByDistance(userLat float64, userLong float64) (*[]model.ShopModel, error)
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
	err := r.db.Where("partner_id = ?", partnerId).
		Preload("Attributes").
		Preload("Image").
		Find(&shops).Error
	if err != nil {
		return nil, err
	}

	return &shops, nil
}

func (r *ShopRepositoryImpl) ShopDetailById(shopId uint) (*model.ShopModel, error) {
	var shop model.ShopModel
	err := r.db.Where("id = ?", shopId).
		Preload("Attributes").
		Preload("WorkingHours").
		Preload("Pictures.Image").
		Preload("Image").
		First(&shop).Error
	if err != nil {
		return nil, err
	}

	return &shop, nil
}

func (r *ShopRepositoryImpl) GetShopsByDistance(userLat float64, userLong float64) (*[]model.ShopModel, error) {
	var shops []model.ShopModel

	err := r.db.Model(&model.ShopModel{}).
		Preload("Image").
		Select(`id, image_id, partner_id, name, location_lat, location_long,
			(6371 * acos(cos(radians(?)) * cos(radians(location_lat)) * cos(radians(location_long) - radians(?)) + sin(radians(?)) * sin(radians(location_lat)))) as distance,
			created_at, updated_at, deleted_at`,
			userLat, userLong, userLat).
		Order("distance ASC").
		Find(&shops).Error

	if err != nil {
		return nil, err
	}

	return &shops, nil
}
