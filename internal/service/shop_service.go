package service

import (
	shop_request "github.com/qahvazor/qahvazor/app/http/request/shop"
	"github.com/qahvazor/qahvazor/app/http/response"
	shop_response "github.com/qahvazor/qahvazor/app/http/response/shop"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/utils"
)


type ShopServiceImpl struct {
	repository *repository.Repository
}

func NewShopService(repository *repository.Repository) ShopService {
	return &ShopServiceImpl{
		repository: repository,
	}
}

func (s *ShopServiceImpl) CreateShop(data shop_request.CreateShopRequest) (interface{}, error) {
	fileName, filePath, fileExt, err := utils.ConvertAndSaveImage(data.File)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	createImageDTO := dto.ImageDTO{
		FileName: fileName,
		FilePath: filePath,
		FileExt: fileExt[1:],
	}
	imageId, err := s.repository.CreateImage(createImageDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	createShopDTO := dto.ShopDTO{
		CompanyID: data.CompanyID, 
		ImageID: int(imageId),
		Name: data.Name,
		Location: data.Location,
	} 
	shop, err := s.repository.CreateShop(createShopDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
   	}
	
	response := shop_response.CreateShopResponse{
		ID: int(shop.ID),
		CompanyID: shop.CompanyID,
		Name: shop.Name,
		Location: shop.Location,
	}
	return response, nil
}

func (s *ShopServiceImpl) CreateShopWorkTime(data shop_request.CreateShopWorkTimeRequest) (interface{}, error) {
	createShopWorkTimeDTO := dto.ShopWorkTimeDTO{
		ShopID: data.ShopID, 
		DayRange: data.DayRange,
		OpeningTime: data.OpeningTime,
		ClosingTime: data.ClosingTime,
	}
	shopWorkTime, err := s.repository.CreateShopWorkTime(createShopWorkTimeDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
   	}

	response := shop_response.CreateShopWorkTimeResponse{
		ID: int(shopWorkTime.ID),
		ShopID: shopWorkTime.ShopID,
		DayRange: shopWorkTime.DayRange,
		OpeningTime: shopWorkTime.OpeningTime,
		ClosingTime: shopWorkTime.ClosingTime,
	}
	return response, nil
}

func (s *ShopServiceImpl) CreateShopPhone(data shop_request.CreateShopPhoneRequest) (interface{}, error) {
	createShopPhoneDTO := dto.ShopPhoneDTO{
		ShopID: data.ShopID, 
		PhoneNumber: data.PhoneNumber,
	}
	shopPhone, err := s.repository.CreateShopPhone(createShopPhoneDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
   	}

	response := shop_response.CreateShopPhoneResponse{
		ID: int(shopPhone.ID),
		ShopID: shopPhone.ShopID,
		PhoneNumber: shopPhone.PhoneNumber,
	}
	return response, nil
}

func (s *ShopServiceImpl) CreateShopSocial(data shop_request.CreateShopSocialRequest) (interface{}, error) {
	createShopSocialDTO := dto.ShopSocialDTO{
		ShopID: data.ShopID, 
		Platform: data.Platform,
		Url: data.Url,
	}
	shopSocial, err := s.repository.CreateShopSocial(createShopSocialDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
   	}

	response := shop_response.CreateShopSocialResponse{
		ID: int(shopSocial.ID),
		ShopID: shopSocial.ShopID,
		Platform: shopSocial.Platform,
		Url: shopSocial.Url,
	}
	return response, nil
}