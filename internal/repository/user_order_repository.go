package repository

import (
	"time"

	"github.com/hoopla/hoopla-api/internal/model"
	"gorm.io/gorm"
)

type UserOrderRepository interface {
	GetAllByUserId(userId uint) (*[]model.UserOrderModel, error)
	GetOrdersNumberForToday(userId uint) (int64, error)
	GetOrderByVendorOrderID(partnerID uint, vendor string, vendorOrderID string) (*model.UserOrderModel, error)
	UpdateOrder(userOrder *model.UserOrderModel) error
	CreateOrder(userOrder *model.UserOrderModel) error
}

type UserOrderRepositoryImpl struct {
	db *gorm.DB
}

func NewUserOrderRepository(db *gorm.DB) UserOrderRepository {
	return &UserOrderRepositoryImpl{
		db: db,
	}
}

func (r *UserOrderRepositoryImpl) GetAllByUserId(userId uint) (*[]model.UserOrderModel, error) {
	var userOrders []model.UserOrderModel
	err := r.db.Model(&model.UserOrderModel{}).
		Where("user_id = ?", userId).
		Order("id desc").
		Preload("Partner", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Preload("Shop", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Preload("Drink", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Limit(10).
		Find(&userOrders).Error

	if err != nil {
		return nil, err
	}

	return &userOrders, nil
}

func (r *UserOrderRepositoryImpl) GetOrdersNumberForToday(userId uint) (int64, error) {
	var ttlUserOrders int64

	today := time.Now().Format("2006-01-02 00:00:00")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02 00:00:00")

	err := r.db.Model(&model.UserOrderModel{}).
		Where("user_id = ? AND created_at >= ? AND created_at < ?", userId, today, tomorrow).
		Where("status != ?", "canceled").
		Order("id desc").
		Count(&ttlUserOrders).Error

	if err != nil {
		return 0, err
	}

	return ttlUserOrders, nil
}

func (r *UserOrderRepositoryImpl) GetOrderByVendorOrderID(partnerID uint, vendor string, vendorOrderID string) (*model.UserOrderModel, error) {
	var userOrder model.UserOrderModel
	err := r.db.Model(&model.UserOrderModel{}).
		Where("partner_id = ?", partnerID).
		Where("vendor = ?", vendor).
		Where("vendor_order_id = ?", vendorOrderID).
		Find(&userOrder).Error

	if err != nil {
		return nil, err
	}

	return &userOrder, nil
}

func (r *UserOrderRepositoryImpl) UpdateOrder(userOrder *model.UserOrderModel) error {
	err := r.db.Save(userOrder).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserOrderRepositoryImpl) CreateOrder(userOrder *model.UserOrderModel) error {
	err := r.db.Create(userOrder).Error
	if err != nil {
		return err
	}

	return nil
}
