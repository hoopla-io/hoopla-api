package repository

import (
	"context"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

// SubscriptionRepository определяет интерфейс для работы с таблицей subscription
type SubscriptionRepository interface {
	GetAllSubscriptions(ctx context.Context) ([]model.Subscription, error)         // Получить все подписки
	GetSubscriptionByID(ctx context.Context, id uint) (*model.Subscription, error) // Получить подписку по ID
}

type subscriptionRepository struct {
	db *gorm.DB
}

// NewSubscriptionRepository создаёт новый экземпляр SubscriptionRepository
func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

// GetAllSubscriptions возвращает все подписки
func (r *subscriptionRepository) GetAllSubscriptions(ctx context.Context) ([]model.Subscription, error) {
	var subscriptions []model.Subscription
	err := r.db.WithContext(ctx).Find(&subscriptions).Error
	return subscriptions, err
}

// GetSubscriptionByID возвращает подписку по её ID
func (r *subscriptionRepository) GetSubscriptionByID(ctx context.Context, id uint) (*model.Subscription, error) {
	var subscription model.Subscription
	err := r.db.WithContext(ctx).First(&subscription, id).Error
	return &subscription, err
}
