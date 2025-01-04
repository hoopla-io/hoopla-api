package service

import (
	"fmt"

	company_request "github.com/qahvazor/qahvazor/app/http/request/company"
	company_social_request "github.com/qahvazor/qahvazor/app/http/request/company/social"
	"github.com/qahvazor/qahvazor/app/http/response"
	company_response "github.com/qahvazor/qahvazor/app/http/response/company"
	company_social_response "github.com/qahvazor/qahvazor/app/http/response/company/social"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/utils"
)

type CompanyService interface {
	GetCompanyShopsList(data company_request.GetCompanyShopsRequest) (interface{}, error)
	Store(data company_request.StoreRequest) (interface{}, error)
	Show(companyId uint) (interface{}, error)
	List() (interface{}, error)
	Edit(data company_request.EditRequest) error
	StoreCompanySocial(data company_social_request.StoreRequest) (interface{}, error)
	ShowCompanySocial(socialId uint) (interface{}, error)
	ListCompanySocials(companyId uint) (interface{}, error)
	EditCompanySocial(data company_social_request.EditRequest) error
}

type CompanyServiceImpl struct {
	CompanyRepository       repository.CompanyRepository
	ImageRepository         repository.ImageRepository
	CompanySocialRepository repository.CompanySocialRepository
	ShopRepository          repository.ShopRepository
}

func NewCompanyService(
	CompanyRepository repository.CompanyRepository,
	ImageRepository repository.ImageRepository,
	CompanySocialRepository repository.CompanySocialRepository,
	ShopRepository repository.ShopRepository,
) CompanyService {
	return &CompanyServiceImpl{
		CompanyRepository:       CompanyRepository,
		ImageRepository:         ImageRepository,
		CompanySocialRepository: CompanySocialRepository,
		ShopRepository:          ShopRepository,
	}
}

func (s *CompanyServiceImpl) GetCompanyShopsList(data company_request.GetCompanyShopsRequest) (interface{}, error) {
	company, err := s.CompanyRepository.GetById(data.CompanyID)
	if err != nil {
		return response.NewErrorResponse(404, "Company not found"), nil
	}

	shops, err := s.ShopRepository.GetByCompanyId(data.CompanyID)
	if err != nil {
		return response.NewErrorResponse(404, "Shops not found"), nil
	}

	image, _ := s.ImageRepository.GetImageById(uint(company.ImageID))
	companyImageUrl := fmt.Sprintf("http://192.168.31.72:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)

	var shopsResponse []company_response.Shop
	for _, shop := range shops {
		shopImage, _ := s.ImageRepository.GetImageById(uint(shop.ImageID))
		shopImageUrl := fmt.Sprintf("http://192.168.31.72:8000/%s/%s.%s", shopImage.FilePath, shopImage.FileName, shopImage.FileExt)
		shopsResponse = append(shopsResponse, company_response.Shop{
			ID:       int(shop.ID),
			Name:     shop.Name,
			Location: shop.Location,
			ImageUrl: shopImageUrl,
		})
	}

	getCompanyShopsResponse := company_response.GetCompanyShopsResponse{
		ID:          int(company.ID),
		Name:        company.Name,
		Description: company.Description,
		ImageUrl:    companyImageUrl,
		Shops:       shopsResponse,
	}

	return getCompanyShopsResponse, nil
}

func (s *CompanyServiceImpl) Store(data company_request.StoreRequest) (interface{}, error) {
	fileName, filePath, fileExt, err := utils.ConvertAndSaveImage(data.File)
	if err != nil {
		return response.NewErrorResponse(500, "Failed to process the image. Please try again later."), nil
	}

	createImageDTO := dto.ImageDTO{
		FileName: fileName,
		FilePath: filePath,
		FileExt:  fileExt[1:],
	}
	imageId, err := s.ImageRepository.CreateImage(createImageDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Failed to save the image. Please try again later."), nil
	}

	createCompanyDTO := dto.CompanyDTO{
		Name:        data.Name,
		Description: data.Description,
		ImageID:     int(imageId),
	}
	companyId, err := s.CompanyRepository.Store(createCompanyDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Failed to save the company details."), nil
	}

	response := company_response.StoreResponse{
		CompanyID: int(companyId),
	}
	return response, nil
}

func (s *CompanyServiceImpl) Show(companyId uint) (interface{}, error) {
	company, err := s.CompanyRepository.GetById(companyId)
	if err != nil {
		return response.NewErrorResponse(404, "Company not found"), nil
	}

	image, err := s.ImageRepository.GetImageById(uint(company.ImageID))
	if err != nil {
		return response.NewErrorResponse(404, "Company not found"), nil
	}
	companyImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)

	showResponse := company_response.ShowResponse{
		ID:          int(company.ID),
		Name:        company.Name,
		Description: company.Description,
		ImageUrl:    companyImageUrl,
	}

	return showResponse, nil
}

func (s *CompanyServiceImpl) List() (interface{}, error) {
	data, err := s.CompanyRepository.List()
	if err != nil {
		return response.NewErrorResponse(500, "Failed to fetch companies. Please try again later."), nil
	}

	var response []company_response.ListResponse
	for _, item := range data {
		image, _ := s.ImageRepository.GetImageById(uint(item.ImageID))
		companyImageUrl := fmt.Sprintf("http://192.168.31.72:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)
		response = append(response, company_response.ListResponse{
			ID:          int(item.ID),
			Name:        item.Name,
			Description: item.Description,
			ImageUrl:    companyImageUrl,
		})
	}

	return response, nil
}

func (s *CompanyServiceImpl) Edit(data company_request.EditRequest) error {
	if data.File != nil {
		fileName, filePath, fileExt, err := utils.ConvertAndSaveImage(data.File)
		if err != nil {
			return fmt.Errorf("failed to save image: %w", err)
		}

		createImageDTO := dto.ImageDTO{
			FileName: fileName,
			FilePath: filePath,
			FileExt:  fileExt[1:],
		}
		imageId, err := s.ImageRepository.CreateImage(createImageDTO)
		if err != nil {
			return fmt.Errorf("failed to create image record: %w", err)
		}
		data.ImageId = imageId
	}

	editDTO := dto.CompanyDTO{
		ID:          uint(data.CompanyID),
		Name:        data.Name,
		Description: data.Description,
		ImageID:     data.ImageId,
	}

	if _, err := s.CompanyRepository.Edit(editDTO); err != nil {
		return fmt.Errorf("failed to edit company record: %w", err)
	}
	return nil
}

func (s *CompanyServiceImpl) StoreCompanySocial(data company_social_request.StoreRequest) (interface{}, error) {
	storeCompanySocialDTO := dto.CompanySocialDTO{
		CompanyID: data.CompanyID,
		Platform:  data.Platform,
		Url:       data.Url,
	}
	companySocialId, err := s.CompanySocialRepository.Store(storeCompanySocialDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	response := company_social_response.StoreResponse{
		SocialID: int(companySocialId),
	}
	return response, nil
}

func (s *CompanyServiceImpl) ShowCompanySocial(socialId uint) (interface{}, error) {
	companySocial, err := s.CompanySocialRepository.GetById(socialId)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	showCompanySocialResponse := company_social_response.ShowResponse{
		ID:        int(companySocial.ID),
		CompanyID: companySocial.CompanyID,
		Platform:  companySocial.Platform,
		Url:       companySocial.Url,
	}

	return showCompanySocialResponse, nil
}

func (s *CompanyServiceImpl) ListCompanySocials(companyId uint) (interface{}, error) {
	data, err := s.CompanySocialRepository.GetListByCompanyId(companyId)
	if err != nil {
		return nil, err
	}

	var response []company_social_response.ListResponse
	for _, item := range data {
		response = append(response, company_social_response.ListResponse{
			ID:        int(item.ID),
			CompanyID: item.CompanyID,
			Platform:  item.Platform,
			Url:       item.Url,
		})
	}

	return response, nil
}

func (s *CompanyServiceImpl) EditCompanySocial(data company_social_request.EditRequest) error {
	editDTO := dto.CompanySocialDTO{
		ID:        uint(data.SocialID),
		CompanyID: data.CompanyID,
		Platform:  data.Platform,
		Url:       data.Url,
	}
	if _, err := s.CompanySocialRepository.Edit(editDTO); err != nil {
		return err
	}

	return nil
}
