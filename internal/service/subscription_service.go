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
	subscription_request "github.com/qahvazor/qahvazor/app/http/request/subscription"
	"github.com/qahvazor/qahvazor/app/http/response"
	subscription_response "github.com/qahvazor/qahvazor/app/http/response/subscription"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/repository"
)


type SubscriptionServiceImpl struct {
	SubscriptionRepository repository.SubscriptionRepository
}

func NewSubscriptionService(
	SubscriptionRepository repository.SubscriptionRepository,
	) SubscriptionService {
	return &SubscriptionServiceImpl{
		SubscriptionRepository: SubscriptionRepository,
	}
}

func (s *SubscriptionServiceImpl) Store(data subscription_request.StoreRequest) (interface{}, error) {
	createSubscriptionDTO := dto.SubscriptionDTO{
		Name: data.Name,
		CoffeeLimit: data.CoffeeLimit,
		Interval: data.Interval,
		Period: data.Period,
	}
	subscriptionId, err := s.SubscriptionRepository.Store(createSubscriptionDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
   	}

	response := subscription_response.StoreResponse{
		SubscriptionID: int(subscriptionId),
	}
	return response, nil
}

func (s *SubscriptionServiceImpl) Show(subscriptionId uint) (interface{}, error) {
	subscription, err := s.SubscriptionRepository.GetById(subscriptionId)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	showResponse := subscription_response.ShowResponse{
		ID:       int(subscription.ID),
		Name:     subscription.Name,
		CoffeeLimit: subscription.CoffeeLimit,
		Interval: subscription.Interval,
		Period: subscription.Period,
	}

	return showResponse, nil
}

func (s *SubscriptionServiceImpl) List() (interface{}, error) {
	data, err := s.SubscriptionRepository.List()
	if err != nil {
		return nil, err
	}

	var response []subscription_response.ListResponse
	for _, item := range data {
		response = append(response, subscription_response.ListResponse{
			ID:          int(item.ID),
			Name:        item.Name,
			CoffeeLimit: item.CoffeeLimit,
			Interval: item.Interval,
			Period: item.Period,
		})
	}

	return response, nil
}

func (s *SubscriptionServiceImpl) Edit(data subscription_request.EditRequest) error {
	editDTO := dto.SubscriptionDTO{
		ID: uint(data.SubscriptionID),
		Name: data.Name,
		CoffeeLimit: data.CoffeeLimit,
		Interval: data.Interval,
		Period: data.Period,
	}
	if _, err := s.SubscriptionRepository.Edit(editDTO); err != nil {
		return err
	}
	return nil
}
