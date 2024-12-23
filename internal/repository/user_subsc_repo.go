package repository

import (
	"context"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
	"time"
)

// UserSubscriptionRepository определяет интерфейс для работы с таблицей user_subscriptions
type UserSubscriptionRepository interface {
	GetUserActiveSubscription(ctx context.Context, userID uint) (*model.UserSubscription, error)  // Получить активную подписку пользователя
	AssignSubscriptionToUser(ctx context.Context, userSubscription *model.UserSubscription) error // Назначить подписку пользователю
}

type userSubscriptionRepository struct {
	db *gorm.DB
}

// NewUserSubscriptionRepository создаёт новый экземпляр UserSubscriptionRepository
func NewUserSubscriptionRepository(db *gorm.DB) UserSubscriptionRepository {
	return &userSubscriptionRepository{db: db}
}

// GetUserActiveSubscription возвращает активную подписку пользователя
func (r *userSubscriptionRepository) GetUserActiveSubscription(ctx context.Context, userID uint) (*model.UserSubscription, error) {
	var userSubscription model.UserSubscription
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND end_date > ?", userID, time.Now()).
		Preload("Subscription"). // Загрузить связанную подписку
		First(&userSubscription).Error
	return &userSubscription, err
}

// AssignSubscriptionToUser добавляет новую подписку пользователю
func (r *userSubscriptionRepository) AssignSubscriptionToUser(ctx context.Context, userSubscription *model.UserSubscription) error {
	return r.db.WithContext(ctx).Create(userSubscription).Error
}
