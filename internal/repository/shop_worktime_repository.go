package repository

import (
	"errors"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type ShopWorkTimeRepositoryImpl struct {
	db *gorm.DB
}

func NewShopWorktimeRepository(db *gorm.DB) ShopWorkTimeRepository {
	return &ShopWorkTimeRepositoryImpl{db: db}
}

func (r *ShopWorkTimeRepositoryImpl) CreateShopWorkTime(data dto.ShopWorkTimeDTO) (*dto.ShopWorkTimeDTO, error) {
	var createShopWorkTime dto.ShopWorkTimeDTO

	query := r.db.Create(&model.ShopWorkTimeModel{
		ShopID: data.ShopID,
		DayRange: data.DayRange,
		OpeningTime: data.OpeningTime,
		ClosingTime: data.ClosingTime,
	}).Scan(&createShopWorkTime)
	if query.Error != nil {
		return nil, query.Error
	}

	return &createShopWorkTime, nil
}

func (r *ShopWorkTimeRepositoryImpl) GetShopWorkTimeByShopId(shopId uint) ([]dto.ShopWorkTimeDTO, error) {
	var shopWorkTimes []dto.ShopWorkTimeDTO
	err := r.db.Model(&model.ShopWorkTimeModel{}).Where("shop_id = ?", shopId).Scan(&shopWorkTimes).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return shopWorkTimes, nil
} 
