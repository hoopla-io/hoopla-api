package repository

import (
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type CoffeeRepositoryImpl struct {
	db *gorm.DB
}

func NewCoffeeRepository(db *gorm.DB) CoffeeRepository {
	return &CoffeeRepositoryImpl{db: db}
}

func (r *CoffeeRepositoryImpl) Store(data dto.CoffeeDTO) (uint, error) {
	var coffee dto.CoffeeDTO

	query := r.db.Create(&model.CoffeeModel{
		Name: data.Name,
		ImageID: data.ImageID,
	}).Scan(&coffee)
	if query.Error != nil {
		return 0, query.Error
	}

	return coffee.ID, nil
}

func (r *CoffeeRepositoryImpl) GetById(coffeeId uint) (dto.CoffeeDTO, error) {
	coffeeModel := model.CoffeeModel{}

	query := r.db.Where("id = ?", coffeeId).
		First(&coffeeModel)
	if query.Error != nil {
		return dto.CoffeeDTO{}, query.Error
	}

	coffee := dto.CoffeeDTO{
		ID:        coffeeModel.ID,
		Name:      coffeeModel.Name,
		ImageID:   coffeeModel.ImageID,
	}

	return coffee, nil
}

func (r *CoffeeRepositoryImpl) List() ([]dto.CoffeeDTO, error) {
	var coffees []dto.CoffeeDTO

	query := r.db.Model(&model.CoffeeModel{}).Find(&coffees)
	if query.Error != nil {
		return nil, query.Error
	}

	return coffees, nil
}

func (r *CoffeeRepositoryImpl) Edit(data dto.CoffeeDTO) (uint, error){
	coffee := model.CoffeeModel{
		Model:    gorm.Model{ID: data.ID},
		Name:     data.Name,
		ImageID: data.ImageID,
	}

	query := r.db.Updates(&coffee)
	if query.Error != nil {
		return 0, query.Error
	}

	return coffee.ID, nil
}