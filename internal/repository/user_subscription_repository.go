package repository

import (
	"github.com/hoopla/hoopla-api/internal/dto"
	"github.com/hoopla/hoopla-api/internal/model"
	"gorm.io/gorm"
)

type UserSubscriptionRepository interface {
	Create(dto dto.UserSubscriptionDTO) error
	GetLastSubscriptionByUserID(userID uint) (*model.UserSubscriptionModel, error)
}

type userSubscriptionRepository struct {
	db *gorm.DB
}

func NewUserSubscriptionRepository(db *gorm.DB) UserSubscriptionRepository {
	return &userSubscriptionRepository{
		db: db,
	}
}

func (r *userSubscriptionRepository) Create(dto dto.UserSubscriptionDTO) error {
	query := r.db.Create(&model.UserSubscriptionModel{
		UserID:         dto.UserID,
		SubscriptionID: dto.SubscriptionID,
		StartDate:      dto.StartDate,
		EndDate:        dto.EndDate,
	})

	return query.Error

}

func (r *userSubscriptionRepository) GetLastSubscriptionByUserID(userID uint) (*model.UserSubscriptionModel, error) {
	var subscription model.UserSubscriptionModel

	err := r.db.Model(&model.UserSubscriptionModel{}).
		Preload("Subscription").
		Preload("SubscriptionDays").
		Last(&subscription, "user_id = ?", userID).Error

	if err != nil {
		return nil, err
	}

	return &subscription, nil
}
