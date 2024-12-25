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


type CompanyServiceImpl struct {
	CompanyRepository repository.CompanyRepository
	ImageRepository   repository.ImageRepository
	CompanySocialRepository repository.CompanySocialRepository
}

func NewCompanyService(
	CompanyRepository repository.CompanyRepository,
	ImageRepository repository.ImageRepository,
	CompanySocialRepository repository.CompanySocialRepository,
	) CompanyService {
	return &CompanyServiceImpl{
		CompanyRepository: CompanyRepository,
		ImageRepository: ImageRepository,
		CompanySocialRepository: CompanySocialRepository,
	}
}

func (s *CompanyServiceImpl) Store(data company_request.StoreRequest) (interface{}, error) {
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

	createCompanyDTO := dto.CompanyDTO{
		Name: data.Name,
		Description: data.Description,
		ImageID: int(imageId),
	}
	companyId, err := s.CompanyRepository.Store(createCompanyDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
   	}

	response := company_response.StoreResponse{
		CompanyID: int(companyId),
	}
	return response, nil
}

func (s *CompanyServiceImpl) Show(companyId uint) (interface{}, error) {
	company, err := s.CompanyRepository.GetById(companyId)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	image, _ := s.ImageRepository.GetImageById(uint(company.ImageID))
	companyImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)

	showResponse := company_response.ShowResponse{
		ID:     int(company.ID),
		Name: company.Name,
		Description: company.Description,
		ImageUrl: companyImageUrl,
	}

	return showResponse, nil
}

func (s *CompanyServiceImpl) List() (interface{}, error) {
	data, err := s.CompanyRepository.List()
	if err != nil {
		return nil, err
	}

	var response []company_response.ListResponse
	for _, item := range data {
		image, _ := s.ImageRepository.GetImageById(uint(item.ImageID))
		companyImageUrl := fmt.Sprintf("http://127.0.0.1:8000/%s/%s.%s", image.FilePath, image.FileName, image.FileExt)
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
		data.ImageId = imageId
	}

	editDTO := dto.CompanyDTO{
		ID: uint(data.CompanyID),
		Name: data.Name,
		Description: data.Description,
		ImageID: data.ImageId,
	}
	
	if _, err := s.CompanyRepository.Edit(editDTO); err != nil {
		return err
	}
	return nil
}


func (s *CompanyServiceImpl) StoreCompanySocial(data company_social_request.StoreRequest) (interface{}, error) {
	storeCompanySocialDTO := dto.CompanySocialDTO{
		CompanyID: data.CompanyID,
		Platform: data.Platform,
		Url:      data.Url,
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
		ID:     int(companySocial.ID),
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
			ID:     int(item.ID),
			CompanyID: item.CompanyID,
			Platform:  item.Platform,
			Url:       item.Url,
		})
	}

	return response, nil
}

func (s *CompanyServiceImpl) EditCompanySocial(data company_social_request.EditRequest) error {
	editDTO := dto.CompanySocialDTO{
		ID: uint(data.SocialID),
		CompanyID: data.CompanyID,
		Platform: data.Platform,
		Url: data.Url,
	}
	if _, err := s.CompanySocialRepository.Edit(editDTO); err != nil {
		return err
	}
	
	return nil
}
