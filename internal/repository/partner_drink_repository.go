package repository

import (
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type PartnerDrinkRepository interface {
	DrinksByPartnerId(partnerId uint) (*[]model.PartnerDrinkModel, error)
}

type PartnerDrinkRepositoryImpl struct {
	db *gorm.DB
}

func NewPartnerDrinkRepository(db *gorm.DB) PartnerDrinkRepository {
	return &PartnerDrinkRepositoryImpl{
		db: db,
	}
}

func (r PartnerDrinkRepositoryImpl) DrinksByPartnerId(partnerId uint) (*[]model.PartnerDrinkModel, error) {
	var drinks []model.PartnerDrinkModel
	if err := r.db.Where("partner_id = ?", partnerId).Find(&drinks).Error; err != nil {
		return nil, err
	}
	return &drinks, nil
}
