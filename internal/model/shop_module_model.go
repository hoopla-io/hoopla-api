package model

type ShopModuleModel struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	PartnerID uint `gorm:"not null;index"`
	ShopID    uint `gorm:"not null;index"`
	ModuleID  uint `gorm:"not null;index"`

	Module *ModuleModel `gorm:"foreignKey:module_id;references:id"`
}

func (ShopModuleModel) TableName() string {
	return "shop_modules"
}
