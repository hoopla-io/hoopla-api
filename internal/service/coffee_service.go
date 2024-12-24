package service

import (
	"fmt"

	coffee_request "github.com/qahvazor/qahvazor/app/http/request/coffee"
	"github.com/qahvazor/qahvazor/app/http/response"
	coffee_response "github.com/qahvazor/qahvazor/app/http/response/coffee"
	company_response "github.com/qahvazor/qahvazor/app/http/response/company"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/utils"
)


type CoffeeServiceImpl struct {
	CoffeeRepository repository.CoffeeRepository
	ImageRepository   repository.ImageRepository
}

func NewCoffeeService(
	CoffeeRepository repository.CoffeeRepository,
	ImageRepository  repository.ImageRepository,
	) CoffeeService {
	return &CoffeeServiceImpl{
		CoffeeRepository: CoffeeRepository,
		ImageRepository:  ImageRepository,
	}
}

func (s *CoffeeServiceImpl) Store(data coffee_request.StoreRequest) (interface{}, error) {
	fileName, filePath, fileExt, err := utils.ConvertAndSaveImage(data.File)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	createImageDTO := dto.ImageDTO{
		FileName: fileName,
		FilePath: filePath,
		FileExt: fileExt[1:],
	}
	imageId, err := s.ImageRepository.CreateImage(createImageDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	createCoffeeDTO := dto.CoffeeDTO{
		Name: data.Name,
		ImageID: int(imageId),
	}
	coffeeId, err := s.CoffeeRepository.Store(createCoffeeDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
   	}

	response := coffee_response.StoreResponse{
		CoffeeID: int(coffeeId),
	}
	return response, nil
}

func (s *CoffeeServiceImpl) Show(coffeeId uint) (interface{}, error) {
	coffee, err := s.CoffeeRepository.GetById(coffeeId)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	image, _ := s.ImageRepository.GetImageById(uint(coffee.ImageID))
	coffeeImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)

	showResponse := company_response.ShowResponse{
		ID:       int(coffee.ID),
		Name:     coffee.Name,
		ImageUrl: coffeeImageUrl,
	}

	return showResponse, nil
}

func (s *CoffeeServiceImpl) List() (interface{}, error) {
	data, err := s.CoffeeRepository.List()
	if err != nil {
		return nil, err
	}

	var response []coffee_response.ListResponse
	for _, item := range data {
		image, _ := s.ImageRepository.GetImageById(uint(item.ImageID))
		coffeeImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)
		response = append(response, coffee_response.ListResponse{
			ID:          int(item.ID),
			Name:        item.Name,
			ImageUrl:    coffeeImageUrl,
		})
	}

	return response, nil
}

func (s *CoffeeServiceImpl) Edit(data coffee_request.EditRequest) error {
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

	editDTO := dto.CoffeeDTO{
		ID: uint(data.CoffeeID),
		Name: data.Name,
		ImageID: int(imageId),
	}
	if _, err := s.CoffeeRepository.Edit(editDTO); err != nil {
		return err
	}
	return nil
}
