package model

type PartnerDrinkAddOnModel struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	PartnerID     uint   `gorm:"not null;index"`
	DrinkID       uint   `gorm:"not null;index"`
	Addon         string `gorm:"not null;index"`
	VendorAddonID string `gorm:"index"`
}

func (PartnerDrinkAddOnModel) TableName() string {
	return "partner_drink_addons"
}
