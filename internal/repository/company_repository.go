package repository

import (
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

func (r *CompanyRepositoryImpl) Store(data dto.CompanyDTO) (uint, error) {
	var company dto.CompanyDTO

	query := r.db.Create(&model.CompanyModel{
		Name: data.Name,
		Description: data.Description,
		ImageID: data.ImageID,
	}).Scan(&company)
	if query.Error != nil {
		return 0, query.Error
	}

	return company.ID, nil
}

func (r *CompanyRepositoryImpl) GetById(categoryId uint) (dto.CompanyDTO, error) {
	companyModel := model.CompanyModel{}

	query := r.db.Where("id = ?", categoryId).
		First(&companyModel)
	if query.Error != nil {
		return dto.CompanyDTO{}, query.Error
	}

	companies := dto.CompanyDTO{
		ID:        companyModel.ID,
		Name:      companyModel.Name,
		Description: companyModel.Description,
	}

	return companies, nil
}

func (r *CompanyRepositoryImpl) List() ([]dto.CompanyDTO, error) {
	var companies []dto.CompanyDTO

	query := r.db.Model(&model.CompanyModel{}).Find(&companies)
	if query.Error != nil {
		return nil, query.Error
	}

	return companies, nil
}

func (r *CompanyRepositoryImpl) Edit(data dto.CompanyDTO) (uint, error){
	company := model.CompanyModel{
		Model:    gorm.Model{ID: data.ID},
		Name:     data.Name,
		Description: data.Description,
		ImageID: data.ImageID,
	}

	query := r.db.Updates(&company)
	if query.Error != nil {
		return 0, query.Error
	}

	return company.ID, nil
}