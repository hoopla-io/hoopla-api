package repository

import (
	"github.com/hoopla/hoopla-api/internal/model"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	SubscriptionsList() (*[]model.SubscriptionModel, error)
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
		Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}

	return &subscriptions, nil
}
