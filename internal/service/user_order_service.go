package service

import (
	user_orders_request "github.com/hoopla/hoopla-api/app/http/request/user/orders"
	user_order_resource "github.com/hoopla/hoopla-api/app/http/resource/user/order"
	"github.com/hoopla/hoopla-api/internal/repository"
)

type UserOrderService interface {
	GetOrders(data user_orders_request.OrdersRequest, userId uint) (*[]user_order_resource.UserOrdersCollection, int, error)
}

type UserOrderServiceImpl struct {
	userOrderRepository repository.UserOrderRepository
}

func NewUserOrderService(userOrderRepository repository.UserOrderRepository) UserOrderService {
	return &UserOrderServiceImpl{
		userOrderRepository: userOrderRepository,
	}
}

func (s *UserOrderServiceImpl) GetOrders(data user_orders_request.OrdersRequest, userId uint) (*[]user_order_resource.UserOrdersCollection, int, error) {
	orders, err := s.userOrderRepository.GetAllByUserId(userId)
	if err != nil {
		return nil, 500, err
	}
	var userOrdersCollection []user_order_resource.UserOrdersCollection
	for _, order := range *orders {
		userOrdersCollection = append(userOrdersCollection, user_order_resource.UserOrdersCollection{
			ID:              order.ID,
			PartnerName:     order.Partner.Name,
			ShopName:        order.Shop.Name,
			PurchasedAt:     order.CreatedAt,
			PurchasedAtUnix: order.CreatedAt.Unix(),
		})
	}

	return &userOrdersCollection, 200, nil
}
