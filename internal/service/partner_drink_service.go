package service

type PartnerDrinkService interface {
}

type PartnerDrinkServiceImpl struct {
}

func NewPartnerDrinkService() PartnerDrinkService {
	return &PartnerDrinkServiceImpl{}
}
