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

func (r *ShopPhoneRepositoryImpl) CreateShopPhone(data dto.ShopPhoneDTO) (*dto.ShopPhoneDTO, error) {
	var createShopPhone dto.ShopPhoneDTO

	query := r.db.Create(&model.ShopPhoneModel{
		ShopID: data.ShopID,
		PhoneNumber: data.PhoneNumber,
	}).Scan(&createShopPhone)
	if query.Error != nil {
		return nil, query.Error
	}

	return &createShopPhone, nil
}

func (r *ShopPhoneRepositoryImpl) GetShopPhoneByShopId(shopId uint) ([]dto.ShopPhoneDTO, error) {
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
