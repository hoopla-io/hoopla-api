package repository

import (
	"github.com/hoopla/hoopla-api/internal/model"
	"gorm.io/gorm"
)

type PartnerRepository interface {
	PartnersList() (*[]model.PartnerModel, error)
	PartnerDetailById(uint) (*model.PartnerModel, error)
	GetPartnerByVendorId(string) (*model.PartnerModel, error)
	UpdateVendorKey(*model.PartnerModel, string) error
}

type PartnerRepositoryImpl struct {
	db *gorm.DB
}

func NewPartnerRepository(db *gorm.DB) PartnerRepository {
	return &PartnerRepositoryImpl{
		db: db,
	}
}

// PartnersList fetches all partners from the database
func (p *PartnerRepositoryImpl) PartnersList() (*[]model.PartnerModel, error) {
	var partners *[]model.PartnerModel
	err := p.db.Model(&model.PartnerModel{}).
		Order("id desc").
		Preload("Logo").
		Find(&partners).Error
	if err != nil {
		return nil, err
	}

	return partners, nil
}

func (p *PartnerRepositoryImpl) PartnerDetailById(id uint) (*model.PartnerModel, error) {
	var partner model.PartnerModel
	if err := p.db.Where("id = ?", id).
		Preload("Logo").
		Preload("Attributes").
		Preload("PartnerDrinks.Drink.Image").
		First(&partner).Error; err != nil {
		return nil, err
	}
	return &partner, nil
}

func (p *PartnerRepositoryImpl) GetPartnerByVendorId(vendorId string) (*model.PartnerModel, error) {
	var partner model.PartnerModel
	if err := p.db.Where("vendor_id = ?", vendorId).First(&partner).Error; err != nil {
		return nil, err
	}

	return &partner, nil
}

func (p *PartnerRepositoryImpl) UpdateVendorKey(partner *model.PartnerModel, vendorKey string) error {
	err := p.db.Model(partner).Update("vendor_key", vendorKey).Error
	if err != nil {
		return err
	}

	return nil
}
