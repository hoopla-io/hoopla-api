package service

import (
	"context"
	"time"

	"github.com/qahvazor/qahvazor/internal/model"
	"github.com/qahvazor/qahvazor/internal/repository"
)

type UserSubscriptionService interface {
	GetUserActiveSubscription(ctx context.Context, userID uint) (*model.UserSubscription, error)
	AssignSubscriptionToUser(ctx context.Context, userID, subscriptionID uint) error
}

type userSubscriptionService struct {
	repo repository.UserSubscriptionRepository
}

func NewUserSubscriptionService(repo repository.UserSubscriptionRepository) UserSubscriptionService {
	return &userSubscriptionService{repo: repo}
}

func (s *userSubscriptionService) GetUserActiveSubscription(ctx context.Context, userID uint) (*model.UserSubscription, error) {
	return s.repo.GetUserActiveSubscription(ctx, userID)
}

func (s *userSubscriptionService) AssignSubscriptionToUser(ctx context.Context, userID, subscriptionID uint) error {
	userSubscription := &model.UserSubscription{
		UserID:         userID,
		SubscriptionID: subscriptionID,
		EndDate:        time.Now().AddDate(0, 0, 30),
	}

	return s.repo.AssignSubscriptionToUser(ctx, userSubscription)
}
