package model

type UserModel struct {
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	Name           string `gorm:"column:name"`
	PhoneNumber    string `gorm:"not null"`
	MobileProvider string `gorm:"not null"`
	CreatedAt      int    `gorm:"not null;default:0"` 
	UpdatedAt      int    `gorm:"column:updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}
