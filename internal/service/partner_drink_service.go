package service

import "github.com/qahvazor/qahvazor/internal/repository"

type PartnerDrinkService interface {
	GetPartnerDrinks()
}

type PartnerDrinkServiceImpl struct {
	partnerDrinkRepository repository.PartnerDrinkRepository
}

func NewPartnerDrinkService(partnerDrinkRepository repository.PartnerDrinkRepository) PartnerDrinkService {
	return &PartnerDrinkServiceImpl{
		partnerDrinkRepository: partnerDrinkRepository,
	}
}

func (s *PartnerDrinkServiceImpl) GetPartnerDrinks() {

}
