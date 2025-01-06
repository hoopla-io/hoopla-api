package repository

import (
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type PartnerRepository interface {
	PartnersList() ([]model.PartnerModel, error)
	PartnerDetailById(id uint) (*model.PartnerModel, error)
}

type PartnerRepositoryImpl struct {
	db *gorm.DB
}

func NewPartnerRepository(db *gorm.DB) PartnerRepository {
	return &PartnerRepositoryImpl{
		db: db,
	}
}

func (p *PartnerRepositoryImpl) PartnersList() ([]model.PartnerModel, error) {
	var partners []model.PartnerModel
	if err := p.db.Model(&model.PartnerModel{}).Find(&partners).Error; err != nil {
		return nil, err
	}

	return partners, nil
}

func (p *PartnerRepositoryImpl) PartnerDetailById(id uint) (*model.PartnerModel, error) {
	var partner model.PartnerModel
	if err := p.db.Where("id = ?", id).First(&partner).Error; err != nil {
		return nil, err
	}
	return &partner, nil
}
