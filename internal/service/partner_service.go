package service

import (
	partners_request "github.com/qahvazor/qahvazor/app/http/request/partners"
	partner_resource "github.com/qahvazor/qahvazor/app/http/resource/partner"
	"github.com/qahvazor/qahvazor/internal/repository"
)

type PartnerService interface {
	PartnersList(data partners_request.PartnersRequest) ([]partner_resource.PartnersCollection, int, error)
	PartnerDetail(data partners_request.PartnerRequest)
}

type PartnerServiceImpl struct {
	PartnerRepository repository.PartnerRepository
}

func NewPartnerService(PartnerRepository repository.PartnerRepository) PartnerService {
	return &PartnerServiceImpl{
		PartnerRepository: PartnerRepository,
	}
}

func (s *PartnerServiceImpl) PartnersList(data partners_request.PartnersRequest) ([]partner_resource.PartnersCollection, int, error) {
	list, err := s.PartnerRepository.PartnersList()
	if err != nil {
		return nil, 500, err
	}
	var partners []partner_resource.PartnersCollection
	for _, item := range list {
		partners = append(partners, partner_resource.PartnersCollection{
			ID:      item.ID,
			Name:    item.Name,
			LogoUrl: nil,
		})
	}

	return partners, 200, nil
}

func (s *PartnerServiceImpl) PartnerDetail(data partners_request.PartnerRequest) {

}
