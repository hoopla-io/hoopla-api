package service

import (
	"fmt"

	shop_request "github.com/qahvazor/qahvazor/app/http/request/shop"
	shop_phone_request "github.com/qahvazor/qahvazor/app/http/request/shop/phone"
	shop_worktime_request "github.com/qahvazor/qahvazor/app/http/request/shop/worktime"
	"github.com/qahvazor/qahvazor/app/http/response"
	shop_response "github.com/qahvazor/qahvazor/app/http/response/shop"
	shop_phone_response "github.com/qahvazor/qahvazor/app/http/response/shop/phone"
	shop_worktime_response "github.com/qahvazor/qahvazor/app/http/response/shop/worktime"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/utils"
)


type ShopServiceImpl struct {
	ShopRepository repository.ShopRepository
	ImageRepository   repository.ImageRepository
	ShopWorktimeRepository repository.ShopWorktimeRepository
	ShopPhoneRepository repository.ShopPhoneRepository
	ShopCoffeeRepository repository.ShopCoffeeRepository
}

func NewShopService(
	ShopRepository repository.ShopRepository, 
	ImageRepository repository.ImageRepository,
	ShopWorktimeRepository repository.ShopWorktimeRepository,
	ShopPhoneRepository repository.ShopPhoneRepository,
	ShopCoffeeRepository repository.ShopCoffeeRepository,
	) ShopService {
	return &ShopServiceImpl{
		ShopRepository: ShopRepository,
		ImageRepository: ImageRepository,
		ShopWorktimeRepository: ShopWorktimeRepository,
		ShopPhoneRepository: ShopPhoneRepository,
		ShopCoffeeRepository: ShopCoffeeRepository,
	}
}

func (s *ShopServiceImpl) Store(data shop_request.StoreRequest) (interface{}, error) {
	fileName, filePath, fileExt, err := utils.ConvertAndSaveImage(data.File)
	if err != nil {
		return response.NewErrorResponse(500, "Invalid file upload. Please check the file and try again."), nil
	}

	createImageDTO := dto.ImageDTO{
		FileName: fileName,
		FilePath: filePath,
		FileExt: fileExt[1:],
	}
	imageId, err := s.ImageRepository.CreateImage(createImageDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Unable to save image information. Please try again later."), nil
	}

	createShopDTO := dto.ShopDTO{
		Name: data.Name,
		Location: data.Location,
		CompanyID: data.CompanyId,
		ImageID: int(imageId),
	}
	shopId, err := s.ShopRepository.Store(createShopDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Failed to create the shop."), nil
   	}

	for _, coffeId := range *data.CoffeeIds {
		if err := s.ShopCoffeeRepository.Store(shopId, coffeId); err != nil {
			return response.NewErrorResponse(500, "Failed to associate coffee items with the shop."), nil
		}
	}
	response := shop_response.StoreResponse{
		ShopID: int(shopId),
	}
	return response, nil
}

func (s *ShopServiceImpl) Show(shopId uint) (interface{}, error) {
	shop, err := s.ShopRepository.GetById(shopId)
	if err != nil {
		return response.NewErrorResponse(404, "Shop not found. Please check the shop ID and try again."), nil
	}

	coffees, err := s.ShopCoffeeRepository.GetListByShopId(shopId)
	if err != nil {
		return nil, err
	}

	image, _ := s.ImageRepository.GetImageById(uint(shop.ImageID))
	shopImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)

	var coffeeResponses []shop_response.Coffee
	for _, coffee := range coffees {
		coffeeImage, _ := s.ImageRepository.GetImageById(uint(*coffee.ImageID))
		coffeeImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", coffeeImage.FilePath, coffeeImage.FileName, coffeeImage.FileExt)
		coffeeResponses = append(coffeeResponses, shop_response.Coffee{
			ID:       coffee.CoffeeID,
			Name:     *coffee.Name,
			ImageUrl: coffeeImageUrl,
		})
	}

	showResponse := shop_response.ShowResponse{
		ID:        int(shop.ID),
		CompanyID: shop.CompanyID,
		Name:      shop.Name,
		Location:  shop.Location,
		ImageUrl:  shopImageUrl,
		Coffees:   coffeeResponses,
	}

	return showResponse, nil
}

func (s *ShopServiceImpl) List() (interface{}, error) {
	data, err := s.ShopRepository.List()
	if err != nil {
		return nil, err
	}

	var response []shop_response.ListResponse
	for _, item := range data {
		image, _ := s.ImageRepository.GetImageById(uint(item.ImageID))
		companyImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)
		response = append(response, shop_response.ListResponse{
			ID:          int(item.ID),
			CompanyID:   item.CompanyID,
			Name:        item.Name,
			Location:    item.Location,
			ImageUrl:    companyImageUrl,
		})
	}

	return response, nil
}

func (s *ShopServiceImpl) Edit(data shop_request.EditRequest) error {
	fileName, filePath, fileExt, err := utils.ConvertAndSaveImage(data.File)
	if err != nil {
		return err
	}

	createImageDTO := dto.ImageDTO{
		FileName: fileName,
		FilePath: filePath,
		FileExt: fileExt[1:],
	}
	imageId, err := s.ImageRepository.CreateImage(createImageDTO)
	if err != nil {
		return err
	}

	editDTO := dto.ShopDTO{
		ID: uint(data.CompanyID),
		CompanyID: int(data.CompanyID),
		Name: data.Name,
		Location: data.Location,
		ImageID: int(imageId),
	}
	if _, err := s.ShopRepository.Edit(editDTO); err != nil {
		return err
	}

	if data.CoffeeIds != nil {
		if err := s.ShopCoffeeRepository.DeleteByShopId(data.ShopID); err != nil {
			return err
		}

		for _, coffeeId := range *data.CoffeeIds {
			if err := s.ShopCoffeeRepository.Store(data.ShopID, coffeeId); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *ShopServiceImpl) StoreShopWorktime(data shop_worktime_request.StoreRequest) (interface{}, error) {
	storeShopWorkTimeDTO := dto.ShopWorktimeDTO{
		ShopID: data.ShopID, 
		DayRange: data.DayRange,
		OpeningTime: data.OpeningTime,
		ClosingTime: data.ClosingTime,
	}
	shopWorkTimeId, err := s.ShopWorktimeRepository.Store(storeShopWorkTimeDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
   	}

	response := shop_worktime_response.StoreResponse{
		ID: int(shopWorkTimeId),
	}
	return response, nil
}

func (s *ShopServiceImpl) ShowWorktime(worktimeId uint) (interface{}, error) {
	shopWorktime, err := s.ShopWorktimeRepository.GetById(worktimeId)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	showWorktimeResponse := shop_worktime_response.ShowResponse{
		ID:     int(shopWorktime.ID),
		ShopID: shopWorktime.ShopID,
		DayRange: shopWorktime.DayRange,
		OpeningTime: shopWorktime.OpeningTime,
		ClosingTime: shopWorktime.ClosingTime,
	}

	return showWorktimeResponse, nil
}

func (s *ShopServiceImpl) ListShopWorktimes(shopId uint) (interface{}, error){
	data, err := s.ShopWorktimeRepository.GetListByShopId(shopId)
	if err != nil {
		return nil, err
	}

	var response []shop_worktime_response.ListResponse
	for _, item := range data {
		response = append(response, shop_worktime_response.ListResponse{
			ID:          int(item.ID),
			ShopID: item.ShopID,
			DayRange: item.DayRange,
			OpeningTime: item.OpeningTime,
			ClosingTime: item.ClosingTime,
		})
	}

	return response, nil
}

func (s *ShopServiceImpl) EditShopWorktime(data shop_worktime_request.EditRequest) error {
	editDTO := dto.ShopWorktimeDTO{
		ID: uint(data.WorktimeID),
		ShopID: int(data.ShopID),
		DayRange: data.DayRange,
		OpeningTime: data.OpeningTime,
		ClosingTime: data.ClosingTime,
	}
	if _, err := s.ShopWorktimeRepository.Edit(editDTO); err != nil {
		return err
	}
	
	return nil
}

func (s *ShopServiceImpl) StoreShopPhone(data shop_phone_request.StoreRequest) (interface{}, error) {
	storeShopPhoneDTO := dto.ShopPhoneDTO{
		ShopID: data.ShopID, 
		PhoneNumber: data.PhoneNumber,
	}
	shopPhoneId, err := s.ShopPhoneRepository.Store(storeShopPhoneDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
   	}

	response := shop_phone_response.StoreResponse{
		ID: int(shopPhoneId),
	}
	return response, nil
}

func (s *ShopServiceImpl) ShowShopPhone(phoneId uint) (interface{}, error) {
	shopPhone, err := s.ShopPhoneRepository.GetById(phoneId)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	showShopPhoneResponse := shop_phone_response.ShowResponse{
		ID:     int(shopPhone.ID),
		ShopID: shopPhone.ShopID,
		PhoneNumber: shopPhone.PhoneNumber,
	}

	return showShopPhoneResponse, nil
}

func (s *ShopServiceImpl) ListShopPhones(shopId uint) (interface{}, error) {
	data, err := s.ShopPhoneRepository.GetListByShopId(shopId)
	if err != nil {
		return nil, err
	}

	var response []shop_phone_response.ListResponse
	for _, item := range data {
		response = append(response, shop_phone_response.ListResponse{
			ID:     int(item.ID),
			ShopID: item.ShopID,
			PhoneNumber: item.PhoneNumber,
		})
	}

	return response, nil
}

func (s *ShopServiceImpl) EditShopPhone(data shop_phone_request.EditRequest) error {
	editDTO := dto.ShopPhoneDTO{
		ID: uint(data.PhoneID),
		ShopID: int(data.ShopID),
		PhoneNumber: data.PhoneNumber,
	}
	if _, err := s.ShopPhoneRepository.Edit(editDTO); err != nil {
		return err
	}
	
	return nil
}
