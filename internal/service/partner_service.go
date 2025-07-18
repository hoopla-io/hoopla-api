package service

import (
	"errors"
	partners_request "github.com/hoopla/hoopla-api/app/http/request/partners"
	partner_resource "github.com/hoopla/hoopla-api/app/http/resource/partner"
	"github.com/hoopla/hoopla-api/internal/model"
	"github.com/hoopla/hoopla-api/internal/repository"
	"gorm.io/gorm"
)

type PartnerService interface {
	PartnersList(data partners_request.PartnersRequest) (*[]partner_resource.PartnersCollection, int, error)
	PartnerDetail(data partners_request.PartnerRequest) (*partner_resource.PartnerResource, int, error)
	GetPartnerByVendorId(vendorId string) (*model.PartnerModel, int, error)
	UpdateVendorKey(vendorId string, vendorKey string) (int, error)
}

type PartnerServiceImpl struct {
	PartnerRepository repository.PartnerRepository
}

func NewPartnerService(PartnerRepository repository.PartnerRepository) PartnerService {
	return &PartnerServiceImpl{
		PartnerRepository: PartnerRepository,
	}
}

func (s *PartnerServiceImpl) PartnersList(data partners_request.PartnersRequest) (*[]partner_resource.PartnersCollection, int, error) {
	list, err := s.PartnerRepository.PartnersList()
	if err != nil {
		return nil, 500, err
	}
	var partners []partner_resource.PartnersCollection
	for _, item := range *list {
		var logoUrl *string
		if item.Logo != nil {
			logoUrl = item.Logo.GetUrl()
		}

		partners = append(partners, partner_resource.PartnersCollection{
			ID:      item.ID,
			Name:    item.Name,
			LogoUrl: logoUrl,
		})
	}

	return &partners, 200, nil
}

func (s *PartnerServiceImpl) PartnerDetail(data partners_request.PartnerRequest) (*partner_resource.PartnerResource, int, error) {
	partner, err := s.PartnerRepository.PartnerDetailById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 404, err
		}
		return nil, 500, err
	}

	partnerResource := partner_resource.PartnerResource{
		ID:          partner.ID,
		Name:        partner.Name,
		Description: partner.Description,
		LogoUrl:     nil,
	}

	var logoUrl *string
	if partner.Logo != nil {
		logoUrl = partner.Logo.GetUrl()
	}

	partnerResource.LogoUrl = logoUrl

	var phoneNumbers []partner_resource.PartnerPhoneNumbersCollection
	var urls []partner_resource.PartnerUrlsCollection
	for _, item := range partner.Attributes {
		switch item.AttributeKey {
		case "phone_number":
			phoneNumbers = append(phoneNumbers, partner_resource.PartnerPhoneNumbersCollection{
				PhoneNumber: item.AttributeValue,
			})
			break
		case "web", "instagram":
			urls = append(urls, partner_resource.PartnerUrlsCollection{
				UrlType: item.AttributeKey,
				Url:     item.AttributeValue,
			})
			break
		}
	}
	partnerResource.PartnerPhoneNumbers = &phoneNumbers
	partnerResource.PartnerUrls = &urls

	var partnerDrinks []partner_resource.DrinksCollection
	for _, partnerDrink := range partner.PartnerDrinks {
		drink := partnerDrink.Drink

		var pictureUrl *string
		if drink.Image != nil {
			pictureUrl = drink.Image.GetUrl()
		}
		partnerDrinks = append(partnerDrinks, partner_resource.DrinksCollection{
			ID:         drink.ID,
			Name:       drink.Name,
			PictureUrl: pictureUrl,
		})
	}
	partnerResource.PartnerDrinks = &partnerDrinks

	return &partnerResource, 200, nil
}

func (s *PartnerServiceImpl) GetPartnerByVendorId(vendorId string) (*model.PartnerModel, int, error) {
	partner, err := s.PartnerRepository.GetPartnerByVendorId(vendorId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 404, err
		}
		return nil, 500, err
	}

	return partner, 200, nil
}

func (s *PartnerServiceImpl) UpdateVendorKey(vendorId string, vendorKey string) (int, error) {
	partnerMode, err := s.PartnerRepository.GetPartnerByVendorId(vendorId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}

	err = s.PartnerRepository.UpdateVendorKey(partnerMode, vendorKey)
	if err != nil {
		return 500, err
	}

	return 200, nil
}
