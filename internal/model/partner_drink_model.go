package model

type PartnerDrinkModel struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	ImageID   uint `gorm:"index"`
	PartnerID uint `gorm:"not null;index"`
	DrinkID   uint `gorm:"not null;index"`

	VendorProductID string `gorm:"index"`

	Drink DrinkModel `gorm:"foreignKey:id;references:drink_id"`
}

func (PartnerDrinkModel) TableName() string {
	return "partner_drinks"
}
