package repository

import (
	"errors"

	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) GetByPhoneNumber(phoneNumber string) (*dto.UserDTO, error) {
	var userModel model.UserModel
	if err := r.db.Where("phone_number = ?", phoneNumber).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &dto.UserDTO{
		ID: int(userModel.ID),
		Name: userModel.Name,
		PhoneNumber: userModel.PhoneNumber,
		MobileProvider: userModel.MobileProvider,
	}, nil
}

func (r *AuthRepositoryImpl) CreateUser(data dto.UserDTO) (*dto.UserDTO, error) {
	var createUser dto.UserDTO

	query := r.db.Create(&model.UserModel{
		PhoneNumber: data.PhoneNumber,
		MobileProvider: data.MobileProvider,
	}).Scan(&createUser)
	if query.Error != nil {
		return nil, query.Error
	}

	return &createUser, nil
}