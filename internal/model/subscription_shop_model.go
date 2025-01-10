package model

type SubscriptionShopModel struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	PartnerID      uint `gorm:"not null;index"`
	ShopID         uint `gorm:"not null;index"`
	SubscriptionID uint `gorm:"not null;index"`
}

func (SubscriptionShopModel) TableName() string {
	return "subscription_shops"
}
