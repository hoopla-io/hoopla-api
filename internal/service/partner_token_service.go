package service

import (
	"errors"
	"github.com/hoopla/hoopla-api/internal/model"
	"github.com/hoopla/hoopla-api/internal/repository"
	"github.com/hoopla/hoopla-api/utils"
	"gorm.io/gorm"
	"time"
)

type PartnerTokenService interface {
	GetAccessToken(partner *model.PartnerModel) (string, error)
}

type PartnerTokenServiceImpl struct {
	repository repository.PartnerTokenRepository
}

func NewPartnerTokenService(repository repository.PartnerTokenRepository) PartnerTokenService {
	return &PartnerTokenServiceImpl{
		repository: repository,
	}
}

func (s *PartnerTokenServiceImpl) GetAccessToken(partner *model.PartnerModel) (string, error) {
	partnerToken, err := s.repository.GetTokenByPartnerID(partner.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	isAccessTokenNeeded := false
	if errors.Is(err, gorm.ErrRecordNotFound) {
		isAccessTokenNeeded = true
	} else if partnerToken.ExpiresAt.Unix() <= time.Now().Add(time.Minute*10).Unix() { // update token
		isAccessTokenNeeded = true
	}

	if isAccessTokenNeeded {
		vendor := utils.Vendor{}
		vendor.Init(partner.Vendor, partner.VendorID, partner.VendorKey)
		token, expiresAt, err := vendor.VendorInterface.GetAccessToken()
		if err != nil {
			return "", err
		}

		if token == "" {
			return "", err
		}

		if partnerToken == nil {
			partnerToken = &model.PartnerTokenModel{
				PartnerID:   partner.ID,
				AccessToken: token,
				ExpiresAt:   expiresAt,
			}

			err := s.repository.CreatePartnerToken(partnerToken)
			if err != nil {
				return "", err
			}
		} else {
			err := s.repository.UpdatePartnerToken(partnerToken, token, expiresAt)
			if err != nil {
				return "", err
			}
		}
	}

	return partnerToken.AccessToken, nil
}
