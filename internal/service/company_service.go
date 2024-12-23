package service

import (
	"fmt"

	company_request "github.com/qahvazor/qahvazor/app/http/request/company"
	"github.com/qahvazor/qahvazor/app/http/response"
	company_response "github.com/qahvazor/qahvazor/app/http/response/company"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/utils"
)


type CompanyServiceImpl struct {
	repository *repository.Repository
}

func NewCompanyService(repository *repository.Repository) CompanyService {
	return &CompanyServiceImpl{
		repository: repository,
	}
}

func (s *CompanyServiceImpl) CreateCompany(data company_request.CreateCompanyRequest) (interface{}, error) {
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

	createCompanyDTO := dto.CompanyDTO{
		Name: data.Name,
		Description: data.Description,
		ImageID: int(imageId),
	}
	company, err := s.repository.CreateCompany(createCompanyDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
   	}

	response := company_response.CreateCompanyResponse{
		ID: int(company.ID),
		Name: company.Name,
		Description: company.Description,
	}
	return response, nil
}

func (s *CompanyServiceImpl) GetCompany(data company_request.GetCompanyRequest) (interface{}, error){
	company, err := s.repository.GetCompanyById(uint(data.CompanyID))
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	image, _ := s.repository.GetImageById(uint(company.ImageID))
	companyImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)
	
	shops, err := s.repository.GetShopsByCompanyId(uint(data.CompanyID))
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	var shopList []company_response.Shops
	for _, shop := range shops {
		image, _ := s.repository.GetImageById(uint(shop.ImageID))
		shopImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)
   		shopList = append(shopList, company_response.Shops{
       		ShopID:    int(shop.ID),
        	CompanyID: shop.CompanyID,
        	Name:      shop.Name,
        	Location:  shop.Location,
			ImageUrl:  shopImageUrl,
    	})
	}

	getCompanyResponse := company_response.GetCompanyResponse{
		ID:     int(company.ID),
		Name: company.Name,
		Description: company.Description,
		ImageUrl: companyImageUrl,
		Shops: shopList,
	}

	return getCompanyResponse, nil
}

func (s *CompanyServiceImpl) GetList() (interface{}, error){
	data, err := s.repository.GetList()
	if err != nil {
		return nil, err
	}

	var response []company_response.GetListResponse
	for _, item := range data {
		image, _ := s.repository.GetImageById(uint(item.ImageID))
		companyImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)
		response = append(response, company_response.GetListResponse{
			ID:          int(item.ID),
			Name:        item.Name,
			Description: item.Description,
			ImageUrl:    companyImageUrl,
		})
	}

	return response, nil
}