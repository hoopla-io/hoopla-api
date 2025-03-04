package service

import (
	"errors"
	subscriptions_request "github.com/hoopla/hoopla-api/app/http/request/subscriptions"
	subscription_resource "github.com/hoopla/hoopla-api/app/http/resource/subscription"
	"github.com/hoopla/hoopla-api/internal/dto"
	"github.com/hoopla/hoopla-api/internal/repository"
	"gorm.io/gorm"
	"time"
)

type SubscriptionService interface {
	GetSubscriptions(data subscriptions_request.SubscriptionsRequest) (*[]subscription_resource.SubscriptionsCollection, int, error)
	BuySubscription(data subscriptions_request.BuySubscriptionRequest, userId uint) (int, error)
}

type SubscriptionServiceImpl struct {
	subscriptionRepository     repository.SubscriptionRepository
	userRepository             repository.UserRepository
	userSubscriptionRepository repository.UserSubscriptionRepository
}

func NewSubscriptionService(
	subscriptionRepository repository.SubscriptionRepository,
	userRepository repository.UserRepository,
	userSubscriptionRepository repository.UserSubscriptionRepository,
) SubscriptionService {
	return &SubscriptionServiceImpl{
		subscriptionRepository:     subscriptionRepository,
		userRepository:             userRepository,
		userSubscriptionRepository: userSubscriptionRepository,
	}
}

func (s *SubscriptionServiceImpl) GetSubscriptions(data subscriptions_request.SubscriptionsRequest) (*[]subscription_resource.SubscriptionsCollection, int, error) {
	subscriptions, err := s.subscriptionRepository.SubscriptionsList()
	if err != nil {
		return nil, 500, err
	}
	var subscriptionsCollection []subscription_resource.SubscriptionsCollection
	for _, item := range *subscriptions {
		subscription := subscription_resource.SubscriptionsCollection{
			ID:       item.ID,
			Name:     item.Name,
			Days:     item.Days,
			Price:    item.Price,
			Currency: item.Currency,
		}
		var features []subscription_resource.FeaturesCollection
		for _, feature := range item.Features {
			features = append(features, subscription_resource.FeaturesCollection{
				ID:      feature.ID,
				Feature: feature.Feature,
			})
		}
		subscription.Features = &features
		subscriptionsCollection = append(subscriptionsCollection, subscription)
	}

	return &subscriptionsCollection, 200, nil
}

func (s *SubscriptionServiceImpl) BuySubscription(data subscriptions_request.BuySubscriptionRequest, userId uint) (int, error) {
	subscription, err := s.subscriptionRepository.GetByID(data.SubscriptionID)
	if err != nil {
		return 500, err
	}

	user, err := s.userRepository.GetByID(userId)
	if err != nil {
		return 500, err
	}

	if user.GetBalance() < subscription.Price {
		return 422, errors.New("Insufficient balance!")
	}

	currentTimeUnix := time.Now().Unix()

	oldSubscription, err := s.userSubscriptionRepository.GetByUserID(user.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 500, err
	}

	if oldSubscription.EndDate > currentTimeUnix {
		return 428, errors.New("You currently have an active subscription!")
	}

	subscriptionEndDateUnix := currentTimeUnix + int64(86400*subscription.Days)
	createSubscription := dto.UserSubscriptionDTO{
		SubscriptionID: data.SubscriptionID,
		UserID:         user.ID,
		StartDate:      currentTimeUnix,
		EndDate:        subscriptionEndDateUnix,
	}
	if err = s.userSubscriptionRepository.Create(createSubscription); err != nil {
		return 500, err
	}

	addCredit := dto.AddCreditDTO{
		UserID: user.ID,
		Amount: subscription.Price,
	}
	if err = s.userRepository.AddCredit(addCredit); err != nil {
		return 500, err
	}

	return 500, nil
}
