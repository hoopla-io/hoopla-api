package repository

import (
	"errors"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type ShopWorkTimeRepository interface {
	Store(data dto.ShopWorktimeDTO) (uint, error)
	GetById(worktimeId uint) (dto.ShopWorktimeDTO, error)
	GetListByShopId(shopId uint) ([]dto.ShopWorktimeDTO, error)
	Edit(data dto.ShopWorktimeDTO) (uint, error)
}

type ShopWorkTimeRepositoryImpl struct {
	db *gorm.DB
}

func NewShopWorkTimeRepository(db *gorm.DB) ShopWorkTimeRepository {
	return &ShopWorkTimeRepositoryImpl{db: db}
}

func (r *ShopWorkTimeRepositoryImpl) Store(data dto.ShopWorktimeDTO) (uint, error) {
	var shopWorktime dto.ShopWorktimeDTO

	query := r.db.Create(&model.ShopWorkTimeModel{
		ShopID:      data.ShopID,
		DayRange:    data.DayRange,
		OpeningTime: data.OpeningTime,
		ClosingTime: data.ClosingTime,
	}).Scan(&shopWorktime)
	if query.Error != nil {
		return 0, query.Error
	}

	return shopWorktime.ID, nil
}

func (r *ShopWorkTimeRepositoryImpl) GetById(worktimeId uint) (dto.ShopWorktimeDTO, error) {
	shopWorktimeModel := model.ShopWorkTimeModel{}

	query := r.db.Where("id = ?", worktimeId).
		First(&shopWorktimeModel)
	if query.Error != nil {
		return dto.ShopWorktimeDTO{}, query.Error
	}

	worktimes := dto.ShopWorktimeDTO{
		ID:          shopWorktimeModel.ID,
		ShopID:      shopWorktimeModel.ShopID,
		DayRange:    shopWorktimeModel.DayRange,
		OpeningTime: shopWorktimeModel.OpeningTime,
		ClosingTime: shopWorktimeModel.ClosingTime,
	}

	return worktimes, nil
}

func (r *ShopWorkTimeRepositoryImpl) GetListByShopId(shopId uint) ([]dto.ShopWorktimeDTO, error) {
	var shopWorktimes []dto.ShopWorktimeDTO
	err := r.db.Model(&model.ShopWorkTimeModel{}).Where("shop_id = ?", shopId).Scan(&shopWorktimes).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return shopWorktimes, nil
}

func (r *ShopWorkTimeRepositoryImpl) Edit(data dto.ShopWorktimeDTO) (uint, error) {
	shopWorktime := model.ShopWorkTimeModel{
		ShopID:      data.ShopID,
		DayRange:    data.DayRange,
		OpeningTime: data.OpeningTime,
		ClosingTime: data.ClosingTime,
	}

	query := r.db.Model(&shopWorktime).Where("id = ?", data.ID).Updates(shopWorktime)
	if query.Error != nil {
		return 0, query.Error
	}

	return shopWorktime.ID, nil
}
