package service

import (
	"errors"
	user_orders_request "github.com/hoopla/hoopla-api/app/http/request/user/orders"
	user_order_resource "github.com/hoopla/hoopla-api/app/http/resource/user/order"
	"github.com/hoopla/hoopla-api/internal/model"
	"github.com/hoopla/hoopla-api/internal/repository"
	"github.com/hoopla/hoopla-api/utils"
	"gorm.io/gorm"
)

type UserOrderService interface {
	GetOrders(data user_orders_request.OrdersRequest, userId uint) (*[]user_order_resource.UserOrdersCollection, int, error)
	GetDrinksStat(userId uint) (*user_order_resource.DrinksStatCollection, int, error)
	GetOrderByVendorOrderID(partnerID uint, vendor string, vendorID string) (*model.UserOrderModel, int, error)
	UpdateOrderStatus(userOrder *model.UserOrderModel, status string) (*model.UserOrderModel, int, error)
	CreateOrder(data user_orders_request.CreateRequest, userHelper *utils.UserHelper) (*model.UserOrderModel, int, error)
}

type UserOrderServiceImpl struct {
	userOrderRepository        repository.UserOrderRepository
	userSubscriptionRepository repository.UserSubscriptionRepository
	partnerTokenService        PartnerTokenService
	partnerDrinkRepository     repository.PartnerDrinkRepository
	shopRepository             repository.ShopRepository
}

func NewUserOrderService(
	userOrderRepository repository.UserOrderRepository,
	UserSubscriptionRepository repository.UserSubscriptionRepository,
	partnerTokenService PartnerTokenService,
	partnerDrinkRepository repository.PartnerDrinkRepository,
	shopRepository repository.ShopRepository,
) UserOrderService {
	return &UserOrderServiceImpl{
		userOrderRepository:        userOrderRepository,
		userSubscriptionRepository: UserSubscriptionRepository,
		partnerTokenService:        partnerTokenService,
		partnerDrinkRepository:     partnerDrinkRepository,
		shopRepository:             shopRepository,
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
			DrinkName:       order.Drink.Name,
			OrderStatus:     order.Status,
		})
	}

	return &userOrdersCollection, 200, nil
}

func (s *UserOrderServiceImpl) GetDrinksStat(userId uint) (*user_order_resource.DrinksStatCollection, int, error) {
	var available uint

	userSubscription, err := s.userSubscriptionRepository.GetByUserID(userId)
	if err != nil {
		return nil, 500, err
	}

	available = userSubscription.Subscription.CupsDay

	orders, err := s.userOrderRepository.GetTodaysByUserId(userId)
	if err != nil {
		return nil, 500, err
	}

	return &user_order_resource.DrinksStatCollection{
		Available: available,
		Left:      uint(len(orders)),
	}, 200, nil
}

func (s *UserOrderServiceImpl) GetOrderByVendorOrderID(partnerID uint, vendor string, vendorOrderID string) (*model.UserOrderModel, int, error) {
	userOder, err := s.userOrderRepository.GetOrderByVendorOrderID(partnerID, vendor, vendorOrderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 404, err
		}
		return nil, 500, err
	}

	return userOder, 200, nil
}

func (s *UserOrderServiceImpl) UpdateOrderStatus(userOrder *model.UserOrderModel, status string) (*model.UserOrderModel, int, error) {
	userOrder.Status = status
	err := s.userOrderRepository.UpdateOrder(userOrder)
	if err != nil {
		return nil, 500, err
	}

	return userOrder, 200, nil
}

func (s *UserOrderServiceImpl) CreateOrder(data user_orders_request.CreateRequest, userHelper *utils.UserHelper) (*model.UserOrderModel, int, error) {
	shop, err := s.shopRepository.ShopBasicDetailById(data.ShopID)
	if err != nil {
		return nil, 500, err
	}
	partner := shop.Partner

	partnerDrink, err := s.partnerDrinkRepository.PartnerDrinkByDrinkId(partner.ID, data.DrinkID)
	if err != nil {
		return nil, 500, err
	}

	vendor := utils.Vendor{}
	vendor.Init(partner.Vendor, partner.VendorID, partner.VendorKey)
	accessToken, err := s.partnerTokenService.GetAccessToken(partner)
	if err != nil {
		return nil, 500, err
	}
	vendor.VendorInterface.SetAccessToken(accessToken)

	userOrder := model.UserOrderModel{
		PartnerID: partner.ID,
		ShopID:    shop.ID,
		UserID:    userHelper.UserID,
		DrinkID:   data.DrinkID,

		Status:        "created",
		Vendor:        partner.Vendor,
		VendorOrderID: "",
		ProductPrice:  partnerDrink.ProductPrice,
	}

	err = s.userOrderRepository.CreateOrder(&userOrder)
	if err != nil {
		return nil, 500, err
	}

	orderStatus, vendorOrderID, err := vendor.VendorInterface.CreateOrder(
		partnerDrink,
		shop,
		&userOrder,
		userHelper.PhoneNumber,
	)
	if err != nil {
		return nil, 500, err
	}

	userOrder.VendorOrderID = vendorOrderID
	userOrder.Status = orderStatus
	err = s.userOrderRepository.UpdateOrder(&userOrder)
	if err != nil {
		return nil, 500, err
	}

	return &userOrder, 200, nil
}
