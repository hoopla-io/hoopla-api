package service

import (
	"context"
	"github.com/qahvazor/qahvazor/internal/model"
	"github.com/qahvazor/qahvazor/internal/repository"
)

type SubscriptionService interface {
	GetAllSubscriptions(ctx context.Context) ([]model.Subscription, error)
	GetSubscriptionByID(ctx context.Context, id uint) (*model.Subscription, error)
}

type subscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{repo: repo}
}

func (s *subscriptionService) GetAllSubscriptions(ctx context.Context) ([]model.Subscription, error) {
	return s.repo.GetAllSubscriptions(ctx)
}

func (s *subscriptionService) GetSubscriptionByID(ctx context.Context, id uint) (*model.Subscription, error) {
	return s.repo.GetSubscriptionByID(ctx, id)
}
