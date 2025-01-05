package service

import "github.com/qahvazor/qahvazor/internal/repository"

type PartnerService interface {
	PartnersList()
	PartnerDetail(id uint)
}

type PartnerServiceImpl struct {
	repository repository.PartnerRepository
}

func NewPartnerService(repository repository.PartnerRepository) PartnerService {
	return &PartnerServiceImpl{
		repository: repository,
	}
}

func (s *PartnerServiceImpl) PartnersList() {

}

func (s *PartnerServiceImpl) PartnerDetail(id uint) {

}
