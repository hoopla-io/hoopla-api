package repository

import (
	"errors"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type CompanyRepositoryImpl struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &CompanyRepositoryImpl{db: db}
}

func (r *CompanyRepositoryImpl) GetCompanyById(id uint) (*dto.CompanyDTO, error) {
	var companyModel model.CompanyModel
	if err := r.db.Where("id = ?", id).First(&companyModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &dto.CompanyDTO{
		ID: companyModel.ID,
		Name: companyModel.Name,
		Description: companyModel.Description,
		ImageID: companyModel.ImageID,
	}, nil
} 

func (r *CompanyRepositoryImpl) CreateCompany(data dto.CompanyDTO) (*dto.CompanyDTO, error) {
	var createCompany dto.CompanyDTO

	query := r.db.Create(&model.CompanyModel{
		Name: data.Name,
		Description: data.Description,
		ImageID: data.ImageID,
	}).Scan(&createCompany)
	if query.Error != nil {
		return nil, query.Error
	}

	return &createCompany, nil
}

func (r *CompanyRepositoryImpl) GetList() ([]dto.CompanyDTO, error) {
	var categories []dto.CompanyDTO

	query := r.db.Model(&model.CompanyModel{}).
		Order("priority").
		Find(&categories)
	if query.Error != nil {
		return nil, query.Error
	}

	return categories, nil
}


