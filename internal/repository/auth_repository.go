package repository

import (
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	GetByPhoneNumber(phoneNumber string) (*model.UserModel, error)
	CreateUser(data dto.UserDTO) (*model.UserModel, error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) GetByPhoneNumber(phoneNumber string) (*model.UserModel, error) {
	var userModel model.UserModel
	if err := r.db.Where("phone_number = ?", phoneNumber).First(&userModel).Error; err != nil {
		return nil, err
	}

	return &userModel, nil
}

func (r *AuthRepositoryImpl) CreateUser(data dto.UserDTO) (*model.UserModel, error) {
	userModel := &model.UserModel{
		Name:           "qahvazor",
		PhoneNumber:    data.PhoneNumber,
		MobileProvider: data.MobileProvider,
	}
	query := r.db.Create(userModel)
	if query.Error != nil {
		return nil, query.Error
	}

	return userModel, nil
}
