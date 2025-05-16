package repository

import (
	"github.com/hoopla/hoopla-api/internal/model"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	SubscriptionsList() (*[]model.SubscriptionModel, error)
	GetByID(id uint) (*model.SubscriptionModel, error)
}

type SubscriptionRepositoryImpl struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &SubscriptionRepositoryImpl{
		db: db,
	}
}

func (r *SubscriptionRepositoryImpl) SubscriptionsList() (*[]model.SubscriptionModel, error) {
	var subscriptions []model.SubscriptionModel
	err := r.db.Model(&model.SubscriptionModel{}).
		Preload("Features").
		Preload("WeekDays").
		Order("priority asc").
		Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}

	return &subscriptions, nil
}

func (r *SubscriptionRepositoryImpl) GetByID(id uint) (*model.SubscriptionModel, error) {
	var subscription model.SubscriptionModel

	query := r.db.First(&subscription, id)
	if query.Error != nil {
		return nil, query.Error
	}

	return &subscription, nil
}
