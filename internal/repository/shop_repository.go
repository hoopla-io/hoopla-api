package repository

import (
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type ShopRepositoryImpl struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &ShopRepositoryImpl{db: db}
}

func (r *ShopRepositoryImpl) GetByCompanyId(companyId uint) ([]dto.ShopDTO, error) {
	var shops []dto.ShopDTO

	query := r.db.Model(&model.ShopModel{}).
		Where("company_id = ?", companyId).
		Find(&shops)
	if query.Error != nil {
		return nil, query.Error
	}

	return shops, nil
}

func (r *ShopRepositoryImpl) Store(data dto.ShopDTO) (uint, error) {
	var shop dto.ShopDTO

	query := r.db.Create(&model.ShopModel{
		Name: data.Name,
		Location: data.Location,
		ImageID: data.ImageID,
		CompanyID: data.CompanyID,
	}).Scan(&shop)
	if query.Error != nil {
		return 0, query.Error
	}

	return shop.ID, nil
}

func (r *ShopRepositoryImpl) GetById(shopId uint) (dto.ShopDTO, error) {
	shopModel := model.ShopModel{}

	query := r.db.Where("id = ?", shopId).
		First(&shopModel)
	if query.Error != nil {
		return dto.ShopDTO{}, query.Error
	}

	shop := dto.ShopDTO{
		ID:        shopModel.ID,
		Name:      shopModel.Name,
		Location:  shopModel.Location,
		CompanyID: shopModel.CompanyID,
		ImageID:   shopModel.ImageID,
	}

	return shop, nil
}

func (r *ShopRepositoryImpl) List() ([]dto.ShopDTO, error) {
	var shops []dto.ShopDTO

	query := r.db.Model(&model.ShopModel{}).Find(&shops)
	if query.Error != nil {
		return nil, query.Error
	}

	return shops, nil
}

func (r *ShopRepositoryImpl) Edit(data dto.ShopDTO) (uint, error){
	shop := model.ShopModel{
		Name:     data.Name,
		Location: data.Location,
		ImageID: data.ImageID,
	}

	query := r.db.Model(&shop).Where("id = ?", data.ID).Updates(shop)
	if query.Error != nil {
		return 0, query.Error
	}

	return shop.ID, nil
}