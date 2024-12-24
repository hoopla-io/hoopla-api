package repository

import (
	"errors"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type ShopWorktimeRepositoryImpl struct {
	db *gorm.DB
}

func NewShopWorktimeRepository(db *gorm.DB) ShopWorktimeRepository {
	return &ShopWorktimeRepositoryImpl{db: db}
}

func (r *ShopWorktimeRepositoryImpl) Store(data dto.ShopWorktimeDTO) (uint, error) {
	var shopWorktime dto.ShopWorktimeDTO

	query := r.db.Create(&model.ShopWorkTimeModel{
		ShopID: data.ShopID,
		DayRange: data.DayRange,
		OpeningTime: data.OpeningTime,
		ClosingTime: data.ClosingTime,
	}).Scan(&shopWorktime)
	if query.Error != nil {
		return 0, query.Error
	}

	return shopWorktime.ID, nil
}

func (r *ShopWorktimeRepositoryImpl) GetById(worktimeId uint) (dto.ShopWorktimeDTO, error) {
	shopWorktimeModel := model.ShopWorkTimeModel{}

	query := r.db.Where("id = ?", worktimeId).
		First(&shopWorktimeModel)
	if query.Error != nil {
		return dto.ShopWorktimeDTO{}, query.Error
	}

	worktimes := dto.ShopWorktimeDTO{
		ID:        shopWorktimeModel.ID,
		ShopID:    shopWorktimeModel.ShopID,
		DayRange:  shopWorktimeModel.DayRange,
		OpeningTime: shopWorktimeModel.OpeningTime,
		ClosingTime: shopWorktimeModel.ClosingTime,
	}

	return worktimes, nil
}

func (r *ShopWorktimeRepositoryImpl) GetListByShopId(shopId uint) ([]dto.ShopWorktimeDTO, error) {
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

func (r *ShopWorktimeRepositoryImpl) Edit(data dto.ShopWorktimeDTO) (uint, error){
	shopWorktime := model.ShopWorkTimeModel{
		Model:    gorm.Model{ID: data.ID},
		ShopID:   data.ShopID,
		DayRange: data.DayRange,
		OpeningTime: data.OpeningTime,
		ClosingTime: data.ClosingTime,
	}

	query := r.db.Updates(&shopWorktime)
	if query.Error != nil {
		return 0, query.Error
	}

	return shopWorktime.ID, nil
}