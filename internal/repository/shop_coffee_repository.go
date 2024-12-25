package repository

import (
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type ShopCoffeeRepositoryImpl struct {
	db *gorm.DB
}

func NewShopCoffeeRepository(db *gorm.DB) ShopCoffeeRepository {
	return &ShopCoffeeRepositoryImpl{db: db}
}

func (r *ShopCoffeeRepositoryImpl) Store(shopId uint, coffeeId uint) error {
	newShopCoffee := model.ShopCoffeeModel{
		ShopID: int(shopId),
		CoffeeID:  int(coffeeId),
	}

	query := r.db.Create(&newShopCoffee)
	if query.Error != nil {
		return query.Error
	}

	return nil
}


func (r *ShopCoffeeRepositoryImpl) GetListByShopId(shopId uint) ([]dto.ShopCoffeeDTO, error) {
	var coffeeIDs []uint

	query := r.db.Model(&model.ShopCoffeeModel{}).
		Select(
			"coffee_id",
		).
		Where("shop_id = ?", shopId).
		Pluck("coffee_id", &coffeeIDs)
	if query.Error != nil {
		return nil, query.Error
	}

	var results []dto.ShopCoffeeDTO

	query = r.db.Model(&model.CoffeeModel{}).
		Select(
			"id",
			"name",
			"image_id",
		).
		Where("id IN ?", coffeeIDs).Scan(&results)
	if query.Error != nil {
		return nil, query.Error
	}

	return results, nil
}

func (r *ShopCoffeeRepositoryImpl) DeleteByShopId(shopId uint) error {
	query := r.db.Where("shop_id = ?", shopId).
		Delete(&model.ShopCoffeeModel{})
	if query.Error != nil {
		return query.Error
	}

	return nil
}
