package service

import (
	"context"
	subscription_request "github.com/qahvazor/qahvazor/app/http/request/subscription"
	"github.com/qahvazor/qahvazor/app/http/response"
	subscription_response "github.com/qahvazor/qahvazor/app/http/response/subscription"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"github.com/qahvazor/qahvazor/internal/repository"
)

// SubscriptionService описывает операции с подписками
type SubscriptionService interface {
	Store(data subscription_request.StoreRequest) (interface{}, error)
	Show(subscriptionID uint) (interface{}, error)
	List() (interface{}, error)
	Edit(data subscription_request.EditRequest) error
	GetAllSubscriptions(ctx context.Context) ([]model.SubscriptionModel, error)
	GetSubscriptionByID(ctx context.Context, id uint) (*model.SubscriptionModel, error)
}

// subscriptionService реализует интерфейс SubscriptionService
type subscriptionService struct {
	repo repository.SubscriptionRepository
}

// NewSubscriptionService создаёт новый экземпляр SubscriptionService
func NewSubscriptionService(repo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{repo: repo}
}

// Store создаёт новую подписку
func (s *subscriptionService) Store(data subscription_request.StoreRequest) (interface{}, error) {
	createSubscriptionDTO := dto.SubscriptionDTO{
		Name:        data.Name,
		CoffeeLimit: data.CoffeeLimit,
		Interval:    data.Interval,
		Period:      data.Period,
	}

	subscriptionID, err := s.repo.Store(createSubscriptionDTO)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), err
	}

	return subscription_response.StoreResponse{
		SubscriptionID: int(subscriptionID),
	}, nil
}

// Show возвращает подробную информацию о подписке
func (s *subscriptionService) Show(subscriptionID uint) (interface{}, error) {
	subscription, err := s.repo.GetById(subscriptionID)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), err
	}

	return subscription_response.ShowResponse{
		ID:          int(subscription.ID),
		Name:        subscription.Name,
		CoffeeLimit: subscription.CoffeeLimit,
		Interval:    subscription.Interval,
		Period:      subscription.Period,
	}, nil
}

// List возвращает список подписок
func (s *subscriptionService) List() (interface{}, error) {
	data, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	var responseList []subscription_response.ListResponse
	for _, item := range data {
		responseList = append(responseList, subscription_response.ListResponse{
			ID:          int(item.ID),
			Name:        item.Name,
			CoffeeLimit: item.CoffeeLimit,
			Interval:    item.Interval,
			Period:      item.Period,
		})
	}

	return responseList, nil
}

// Edit обновляет существующую подписку
func (s *subscriptionService) Edit(data subscription_request.EditRequest) error {
	editDTO := dto.SubscriptionDTO{
		ID:          uint(data.SubscriptionID),
		Name:        data.Name,
		CoffeeLimit: data.CoffeeLimit,
		Interval:    data.Interval,
		Period:      data.Period,
	}
	return s.repo.Edit(editDTO)
}

// Реализация методов в subscriptionService
func (s *subscriptionService) GetAllSubscriptions(ctx context.Context) ([]model.SubscriptionModel, error) {
	return s.repo.GetAllSubscriptions(ctx)
}

func (s *subscriptionService) GetSubscriptionByID(ctx context.Context, id uint) (*model.SubscriptionModel, error) {
	return s.repo.GetSubscriptionByID(ctx, id)
}
