package repository

import (
	"context"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type SubscriptionRepositoryImpl struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &SubscriptionRepositoryImpl{db: db}
}

// GetAllSubscriptions возвращает все подписки
func (r *SubscriptionRepositoryImpl) GetAllSubscriptions(ctx context.Context) ([]model.SubscriptionModel, error) {
	var subscriptions []model.SubscriptionModel
	err := r.db.WithContext(ctx).Find(&subscriptions).Error
	return subscriptions, err
}

// GetSubscriptionByID возвращает подписку по её ID
func (r *SubscriptionRepositoryImpl) GetSubscriptionByID(ctx context.Context, id uint) (*model.SubscriptionModel, error) {
	var subscription model.SubscriptionModel
	err := r.db.WithContext(ctx).First(&subscription, id).Error
	return &subscription, err
}

func (r *SubscriptionRepositoryImpl) Store(data dto.SubscriptionDTO) (uint, error) {
	var subscription dto.SubscriptionDTO

	query := r.db.Create(&model.SubscriptionModel{
		Name: data.Name,
		CoffeeLimit: data.CoffeeLimit,
		Interval: data.Interval,
		Period: data.Period,
	}).Scan(&subscription)
	if query.Error != nil {
		return 0, query.Error
	}

	return subscription.ID, nil
}

func (r *SubscriptionRepositoryImpl) GetById(subscriptionId uint) (dto.SubscriptionDTO, error) {
	subscriptionModel := model.SubscriptionModel{}

	query := r.db.Where("id = ?", subscriptionId).
		First(&subscriptionModel)
	if query.Error != nil {
		return dto.SubscriptionDTO{}, query.Error
	}

	subscription := dto.SubscriptionDTO{
		ID:        subscriptionModel.ID,
		Name:      subscriptionModel.Name,
		CoffeeLimit: subscriptionModel.CoffeeLimit,
		Interval:    subscriptionModel.Interval,
		Period:      subscriptionModel.Period,
	}

	return subscription, nil
}

func (r *SubscriptionRepositoryImpl) List() ([]dto.SubscriptionDTO, error) {
	var subscriptions []dto.SubscriptionDTO

	query := r.db.Model(&model.SubscriptionModel{}).Find(&subscriptions)
	if query.Error != nil {
		return nil, query.Error
	}

	return subscriptions, nil
}

func (r *SubscriptionRepositoryImpl) Edit(data dto.SubscriptionDTO) (uint, error){
	subscription := model.SubscriptionModel{
		Name:     data.Name,
		CoffeeLimit: data.CoffeeLimit,
		Interval: data.Interval,
		Period: data.Period,
	}

	query := r.db.Model(&subscription).Where("id = ?", data.ID).Updates(subscription)
	if query.Error != nil {
		return 0, query.Error
	}

	return subscription.ID, nil
}