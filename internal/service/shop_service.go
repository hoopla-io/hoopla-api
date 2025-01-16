package service

import (
	"errors"
	"fmt"

	shops_request "github.com/hoopla/hoopla-api/app/http/request/shops"
	partner_resource "github.com/hoopla/hoopla-api/app/http/resource/partner"
	shop_resource "github.com/hoopla/hoopla-api/app/http/resource/shop"
	"github.com/hoopla/hoopla-api/internal/repository"
	"gorm.io/gorm"
)

type ShopService interface {
	NearShopsList(data shops_request.NearShopsRequest) (*[]partner_resource.ShopsCollections, int, error)
	PartnerShopsList(data shops_request.PartnerShopsRequest) (*[]partner_resource.ShopsCollections, int, error)
	ShopDetail(data shops_request.ShopRequest) (*shop_resource.ShopResource, int, error)
}

type ShopServiceImpl struct {
	ShopRepository repository.ShopRepository
}

func NewShopService(shopRepository repository.ShopRepository) ShopService {
	return &ShopServiceImpl{
		ShopRepository: shopRepository,
	}
}

func (s *ShopServiceImpl) NearShopsList(data shops_request.NearShopsRequest) (*[]partner_resource.ShopsCollections, int, error) {
	shops, err := s.ShopRepository.GetShopsByDistance(data.Lat, data.Long)
	if err != nil {
		return nil, 500, err
	}

	var shopsCollection []partner_resource.ShopsCollections
	for _, shop := range *shops {
		fmt.Println(shop)
		shopResource := partner_resource.ShopsCollections{
			ShopID: shop.ID,
			Name:   shop.Name,
			Location: partner_resource.ShopLocationResource{
				Lat: shop.LocationLat,
				Lng: shop.LocationLong,
			},
			Distance: shop.Distance,
		}

		var pictureUrl *string
		picture := shop.Image
		if picture != nil {
			pictureUrl = picture.GetUrl()
			shopResource.PictureURL = pictureUrl
		}

		shopsCollection = append(shopsCollection, shopResource)
	}

	return &shopsCollection, 200, nil
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
			Location: partner_resource.ShopLocationResource{
				Lat: shop.LocationLat,
				Lng: shop.LocationLong,
			},
		}

		var phoneNumbers []partner_resource.ShopPhoneNumbersCollection
		for _, attribute := range *shop.Attributes {
			switch attribute.AttributeKey {
			case "phone_number":
				phoneNumbers = append(phoneNumbers, partner_resource.ShopPhoneNumbersCollection{
					PhoneNumber: attribute.AttributeValue,
				})
			}
		}
		shopResource.PhoneNumbers = &phoneNumbers

		var pictureUrl *string
		picture := shop.Image
		if picture != nil {
			pictureUrl = picture.GetUrl()
			shopResource.PictureURL = pictureUrl
		}

		shopsCollection = append(shopsCollection, shopResource)
	}

	return &shopsCollection, 200, nil
}

func (s *ShopServiceImpl) ShopDetail(data shops_request.ShopRequest) (*shop_resource.ShopResource, int, error) {
	shop, err := s.ShopRepository.ShopDetailById(data.ShopId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 404, err
		}
		return nil, 500, err
	}

	shopResource := shop_resource.ShopResource{
		ID:        shop.ID,
		PartnerId: shop.PartnerID,
		Name:      shop.Name,
		Location: shop_resource.ShopLocationResource{
			Lat: shop.LocationLat,
			Lng: shop.LocationLong,
		},
	}

	var phoneNumbers []shop_resource.ShopPhoneNumberResource
	for _, attribute := range *shop.Attributes {
		switch attribute.AttributeKey {
		case "phone_number":
			phoneNumbers = append(phoneNumbers, shop_resource.ShopPhoneNumberResource{
				PhoneNumber: attribute.AttributeValue,
			})
		}
	}
	shopResource.PhoneNumbers = &phoneNumbers

	var workingHours []shop_resource.ShopWorkingHoursResource
	for _, workingHour := range *shop.WorkingHours {
		workingHours = append(workingHours, shop_resource.ShopWorkingHoursResource{
			WeekDay: workingHour.WeekDay,
			OpenAt:  workingHour.OpenAt,
			CloseAt: workingHour.CloseAt,
		})
	}
	shopResource.ShopWorkingHours = &workingHours

	var pictures []shop_resource.ShopPictureResource
	for _, picture := range *shop.Pictures {
		var pictureUrl *string
		if picture.Image != nil {
			pictureUrl = picture.Image.GetUrl()
			pictures = append(pictures, shop_resource.ShopPictureResource{
				PictureUrl: pictureUrl,
			})
		}
	}
	shopResource.Pictures = &pictures

	var pictureUrl *string
	picture := shop.Image
	if picture != nil {
		pictureUrl = picture.GetUrl()
		shopResource.PictureUrl = pictureUrl
	}

	return &shopResource, 200, nil
}
