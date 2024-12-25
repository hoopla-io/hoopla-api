package repository

import (
	"errors"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type CompanySocialRepositoryImpl struct {
	db *gorm.DB
}

func NewCompanySocialRepository(db *gorm.DB) CompanySocialRepository {
	return &CompanySocialRepositoryImpl{db: db}
}

func (r *CompanySocialRepositoryImpl) Store(data dto.CompanySocialDTO) (uint, error) {
	var companySocial dto.CompanySocialDTO

	query := r.db.Create(&model.CompanySocialModel{
		CompanyID: data.CompanyID,
		Platform: data.Platform,
		Url: data.Url,
	}).Scan(&companySocial)
	if query.Error != nil {
		return 0, query.Error
	}

	return companySocial.ID, nil
}

func (r *CompanySocialRepositoryImpl) GetById(socialId uint) (dto.CompanySocialDTO, error) {
	companySocialModel := model.CompanySocialModel{}

	query := r.db.Where("id = ?", socialId).
		First(&companySocialModel)
	if query.Error != nil {
		return dto.CompanySocialDTO{}, query.Error
	}

	worktimes := dto.CompanySocialDTO{
		ID:        companySocialModel.ID,
		CompanyID: companySocialModel.CompanyID,
		Platform:  companySocialModel.Platform,
		Url:       companySocialModel.Url,
	}

	return worktimes, nil
}

func (r *CompanySocialRepositoryImpl) GetListByCompanyId(companyId uint) ([]dto.CompanySocialDTO, error) {
	var companySocials []dto.CompanySocialDTO
	err := r.db.Model(&model.CompanySocialModel{}).Where("company_id = ?", companyId).Scan(&companySocials).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return companySocials, nil
}

func (r *CompanySocialRepositoryImpl) Edit(data dto.CompanySocialDTO) (uint, error){
	companySocial := model.CompanySocialModel{
		CompanyID: data.CompanyID,
		Platform:  data.Platform,
		Url:       data.Url,
	}

	query :=  r.db.Model(&companySocial).Where("id = ?", data.ID).Updates(companySocial)
	if query.Error != nil {
		return 0, query.Error
	}

	return companySocial.ID, nil
}
