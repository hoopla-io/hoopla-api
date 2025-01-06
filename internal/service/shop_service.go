package service

import (
	"fmt"
	shops_request "github.com/qahvazor/qahvazor/app/http/request/shops"
	partner_resource "github.com/qahvazor/qahvazor/app/http/resource/partner"
	"github.com/qahvazor/qahvazor/internal/repository"
	"strconv"
)

type ShopService interface {
	PartnerShopsList(data shops_request.PartnerShopsRequest) (*[]partner_resource.ShopsCollections, int, error)
}

type ShopServiceImpl struct {
	ShopRepository repository.ShopRepository
}

func NewShopService(shopRepository repository.ShopRepository) ShopService {
	return &ShopServiceImpl{
		ShopRepository: shopRepository,
	}
}

func (s *ShopServiceImpl) PartnerShopsList(data shops_request.PartnerShopsRequest) (*[]partner_resource.ShopsCollections, int, error) {
	shops, err := s.ShopRepository.GetPartnerShops(data.PartnerID)
	if err != nil {
		return nil, 500, err
	}

	var shopsCollection []partner_resource.ShopsCollections
	for _, shop := range *shops {
		shopResource := partner_resource.ShopsCollections{
			ShopID: shop.ID,
			Name:   shop.Name,
		}

		var shopLocation partner_resource.ShopLocationResource
		var phoneNumbers []partner_resource.ShopPhoneNumbersCollection
		for _, attribute := range *shop.Attributes {
			switch attribute.AttributeKey {
			case "lat":
				floatValue, err := strconv.ParseFloat(attribute.AttributeValue, 64)
				if err != nil {
					fmt.Println("Error:", err)
					break
				}
				shopLocation.Lat = floatValue
				break
			case "lng":
				floatValue, err := strconv.ParseFloat(attribute.AttributeValue, 64)
				if err != nil {
					fmt.Println("Error:", err)
					break
				}
				shopLocation.Lng = floatValue
				break
			case "phone_number":
				phoneNumbers = append(phoneNumbers, partner_resource.ShopPhoneNumbersCollection{
					PhoneNumber: attribute.AttributeValue,
				})
			}
		}
		shopResource.Location = &shopLocation
		shopResource.PhoneNumbers = &phoneNumbers

		shopsCollection = append(shopsCollection, shopResource)
	}

	return &shopsCollection, 200, nil
}
