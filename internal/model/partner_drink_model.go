package model

type PartnerDrinkModel struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	ImageID   uint `gorm:"index"`
	PartnerID uint `gorm:"not null;index"`
	DrinkID   uint `gorm:"not null;index"`

	VendorProductID    string  `gorm:"index"`
	VendorProductPrice float64 `gorm:"not null;default:0.0"`
	ProductPrice       float64 `gorm:"not null;default:0.0"`

	Partner *PartnerModel `gorm:"foreignKey:id;references:partner_id"`
	Drink   *DrinkModel   `gorm:"foreignKey:id;references:drink_id"`
}

func (PartnerDrinkModel) TableName() string {
	return "partner_drinks"
}
