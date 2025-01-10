package service

import (
	subscriptions_request "github.com/qahvazor/qahvazor/app/http/request/subscriptions"
	subscription_resource "github.com/qahvazor/qahvazor/app/http/resource/subscription"
	"github.com/qahvazor/qahvazor/internal/repository"
)

type SubscriptionService interface {
	GetSubscriptions(data subscriptions_request.SubscriptionsRequest) (*[]subscription_resource.SubscriptionsCollection, int, error)
}

type SubscriptionServiceImpl struct {
	subscriptionRepository repository.SubscriptionRepository
}

func NewSubscriptionService(subscriptionRepository repository.SubscriptionRepository) SubscriptionService {
	return &SubscriptionServiceImpl{
		subscriptionRepository: subscriptionRepository,
	}
}

func (s *SubscriptionServiceImpl) GetSubscriptions(data subscriptions_request.SubscriptionsRequest) (*[]subscription_resource.SubscriptionsCollection, int, error) {
	subscriptions, err := s.subscriptionRepository.SubscriptionsList()
	if err != nil {
		return nil, 500, err
	}
	var subscriptionsCollection []subscription_resource.SubscriptionsCollection
	for _, item := range *subscriptions {
		subscriptionsCollection = append(subscriptionsCollection, subscription_resource.SubscriptionsCollection{
			ID:       item.ID,
			Name:     item.Name,
			Days:     item.Days,
			Price:    item.Price,
			Currency: item.Currency,
		})
	}

	return &subscriptionsCollection, 200, nil
}
