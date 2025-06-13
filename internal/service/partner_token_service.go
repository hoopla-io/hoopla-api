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
	GetAccessToken(shop model.ShopModel) (string, error)
}

type PartnerTokenServiceImpl struct {
	repository repository.PartnerTokenRepository
}

func NewPartnerTokenService(repository repository.PartnerTokenRepository) PartnerTokenService {
	return &PartnerTokenServiceImpl{
		repository: repository,
	}
}

func (s *PartnerTokenServiceImpl) GetAccessToken(shop model.ShopModel) (string, error) {
	partnerToken, err := s.repository.GetTokenByShopID(shop.ID)
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
		vendor.Init(shop.Partner.Vendor, shop.Partner.VendorID, shop.Partner.VendorKey)
		token, expiresAt, err := vendor.VendorInterface.GetAccessToken()
		if err != nil {
			return "", err
		}

		if partnerToken == nil {
			partnerToken = &model.PartnerTokenModel{
				ShopID:      shop.ID,
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
