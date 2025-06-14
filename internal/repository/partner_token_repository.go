package repository

import (
	"github.com/hoopla/hoopla-api/internal/model"
	"gorm.io/gorm"
	"time"
)

type PartnerTokenRepository interface {
	GetTokenByPartnerID(partnerID uint) (*model.PartnerTokenModel, error)
	UpdatePartnerToken(*model.PartnerTokenModel, string, time.Time) error
	CreatePartnerToken(*model.PartnerTokenModel) error
}

type PartnerTokenRepositoryImpl struct {
	db *gorm.DB
}

func NewPartnerTokenRepository(db *gorm.DB) PartnerTokenRepository {
	return &PartnerRepositoryImpl{
		db: db,
	}
}

func (r *PartnerRepositoryImpl) GetTokenByPartnerID(partnerID uint) (*model.PartnerTokenModel, error) {
	var token model.PartnerTokenModel

	err := r.db.First(&token, "partner_id = ?", partnerID).Error
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *PartnerRepositoryImpl) UpdatePartnerToken(partnerToken *model.PartnerTokenModel, accessToken string, expiresAt time.Time) error {
	err := r.db.Model(partnerToken).Update("access_token", accessToken).Update("expires_at", expiresAt).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *PartnerRepositoryImpl) CreatePartnerToken(partnerToken *model.PartnerTokenModel) error {
	err := r.db.Create(partnerToken).Error
	if err != nil {
		return err
	}

	return nil
}
