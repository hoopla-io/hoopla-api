package repository

import (
	"errors"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type ShopSocialRepositoryImpl struct {
	db *gorm.DB
}

func NewShopSocialRepository(db *gorm.DB) ShopSocialRepository {
	return &ShopSocialRepositoryImpl{db: db}
}

func (r *ShopSocialRepositoryImpl) CreateShopSocial(data dto.ShopSocialDTO) (*dto.ShopSocialDTO, error) {
	var createShopSocial dto.ShopSocialDTO

	query := r.db.Create(&model.ShopSocialModel{
		ShopID: data.ShopID,
		Platform: data.Platform,
		Url: data.Url,
	}).Scan(&createShopSocial)
	if query.Error != nil {
		return nil, query.Error
	}

	return &createShopSocial, nil
}

func (r *ShopSocialRepositoryImpl) GetShopSocialByShopId(shopId uint) ([]dto.ShopSocialDTO, error) {
	var shopSocials []dto.ShopSocialDTO
	err := r.db.Model(&model.ShopSocialModel{}).Where("shop_id = ?", shopId).Scan(&shopSocials).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return shopSocials, nil
} 
