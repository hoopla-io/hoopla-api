package repository

import (
	"time"

	"github.com/hoopla/hoopla-api/internal/model"
	"gorm.io/gorm"
)

type UserOrderRepository interface {
	GetAllByUserId(userId uint) (*[]model.UserOrderModel, error)
	GetTodaysByUserId(userId uint) ([]model.UserOrderModel, error)
}

type UserOrderRepositoryImpl struct {
	db *gorm.DB
}

func NewUserOrderRepository(db *gorm.DB) UserOrderRepository {
	return &UserOrderRepositoryImpl{
		db: db,
	}
}

func (r *UserOrderRepositoryImpl) GetAllByUserId(userId uint) (*[]model.UserOrderModel, error) {
	var userOrders []model.UserOrderModel
	err := r.db.Model(&model.UserOrderModel{}).
		Where("user_id = ?", userId).
		Order("id desc").
		Preload("Partner").
		Preload("Shop").
		Find(&userOrders).Error

	if err != nil {
		return nil, err
	}

	return &userOrders, nil
}

func (r *UserOrderRepositoryImpl) GetTodaysByUserId(userId uint) ([]model.UserOrderModel, error) {
	var userOrders []model.UserOrderModel

	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	err := r.db.Model(&model.UserOrderModel{}).
		Where("user_id = ? AND created_at >= ? AND created_at < ?", userId, today, tomorrow).
		Order("id desc").
		Preload("Partner").
		Preload("Shop").
		Find(&userOrders).Error

	if err != nil {
		return nil, err
	}

	return userOrders, nil
}
