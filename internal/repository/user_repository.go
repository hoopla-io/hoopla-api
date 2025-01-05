package repository

import (
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByPhoneNumber(phoneNumber string) (*model.UserModel, error)
	CreateUser(data dto.UserDTO) (*model.UserModel, error)
	UpdateToken(uuid string, user *model.UserModel) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetByPhoneNumber(phoneNumber string) (*model.UserModel, error) {
	var userModel model.UserModel
	if err := r.db.Where("phone_number = ?", phoneNumber).First(&userModel).Error; err != nil {
		return nil, err
	}

	return &userModel, nil
}

func (r *UserRepositoryImpl) CreateUser(data dto.UserDTO) (*model.UserModel, error) {
	userModel := &model.UserModel{
		Name:           "qahvazor",
		PhoneNumber:    data.PhoneNumber,
		MobileProvider: data.MobileProvider,
		RefreshToken:   data.RefreshToken,
	}
	query := r.db.Create(userModel)
	if query.Error != nil {
		return nil, query.Error
	}

	return userModel, nil
}

func (r *UserRepositoryImpl) UpdateToken(uuid string, user *model.UserModel) error {
	if err := r.db.Model(user).Update("refresh_token", uuid).Error; err != nil {
		return err
	}

	return nil
}
