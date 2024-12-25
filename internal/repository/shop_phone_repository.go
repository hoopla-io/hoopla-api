package repository

import (
	"errors"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type ShopPhoneRepositoryImpl struct {
	db *gorm.DB
}

func NewShopPhoneRepository(db *gorm.DB) ShopPhoneRepository {
	return &ShopPhoneRepositoryImpl{db: db}
}

func (r *ShopPhoneRepositoryImpl) Store(data dto.ShopPhoneDTO) (uint, error) {
	var shopPhone dto.ShopPhoneDTO

	query := r.db.Create(&model.ShopPhoneModel{
		ShopID: data.ShopID,
		PhoneNumber: data.PhoneNumber,
	}).Scan(&shopPhone)
	if query.Error != nil {
		return 0, query.Error
	}

	return shopPhone.ID, nil
}

func (r *ShopPhoneRepositoryImpl) GetById(phoneId uint) (dto.ShopPhoneDTO, error) {
	shopPhoneModel := model.ShopPhoneModel{}

	query := r.db.Where("id = ?", phoneId).
		First(&shopPhoneModel)
	if query.Error != nil {
		return dto.ShopPhoneDTO{}, query.Error
	}

	phones := dto.ShopPhoneDTO{
		ID:        shopPhoneModel.ID,
		ShopID:    shopPhoneModel.ShopID,
		PhoneNumber: shopPhoneModel.PhoneNumber,
	}

	return phones, nil
}

func (r *ShopPhoneRepositoryImpl) GetListByShopId(shopId uint) ([]dto.ShopPhoneDTO, error) {
	var shopPhones []dto.ShopPhoneDTO
	err := r.db.Model(&model.ShopPhoneModel{}).Where("shop_id = ?", shopId).Scan(&shopPhones).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return shopPhones, nil
}

func (r *ShopPhoneRepositoryImpl) Edit(data dto.ShopPhoneDTO) (uint, error){
	shopPhone := model.ShopPhoneModel{
		ShopID:   data.ShopID,
		PhoneNumber: data.PhoneNumber,
	}

	query := r.db.Model(&shopPhone).Where("id = ?", data.ID).Updates(shopPhone)
	if query.Error != nil {
		return 0, query.Error
	}

	return shopPhone.ID, nil
}