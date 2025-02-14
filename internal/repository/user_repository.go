package repository

import (
	"fmt"
	"time"

	"github.com/hoopla/hoopla-api/internal/dto"
	"github.com/hoopla/hoopla-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(ID uint) (*model.UserModel, error)
	GetByPhoneNumber(phoneNumber string) (*model.UserModel, error)
	GetByRefreshToken(refreshToken string) (*model.UserModel, error)
	CreateUser(data dto.UserDTO) (*model.UserModel, error)
	UpdateToken(uuid string, user *model.UserModel) error
	RemoveToken(*model.UserModel) error
	AddCredit(creditDto dto.AddCreditDTO) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetByID(id uint) (*model.UserModel, error) {
	var userModel model.UserModel
	if err := r.db.Where("id = ?", id).First(&userModel).Error; err != nil {
		return nil, err
	}

	return &userModel, nil
}

func (r *UserRepositoryImpl) GetByPhoneNumber(phoneNumber string) (*model.UserModel, error) {
	var userModel model.UserModel
	if err := r.db.Where("phone_number = ?", phoneNumber).First(&userModel).Error; err != nil {
		return nil, err
	}

	return &userModel, nil
}

func (r *UserRepositoryImpl) GetByRefreshToken(refreshToken string) (*model.UserModel, error) {
	var userModel model.UserModel
	if err := r.db.Where("refresh_token = ?", refreshToken).First(&userModel).Error; err != nil {
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

func (r *UserRepositoryImpl) RemoveToken(user *model.UserModel) error {
	if err := r.db.Model(user).Update("refresh_token", fmt.Sprintf("%d", time.Now().UnixMicro())).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) AddCredit(creditDto dto.AddCreditDTO) error {
	amountCents := int64(creditDto.Amount * 100)
	query := r.db.Model(&model.UserModel{}).
		Where("id = ?", creditDto.UserID).
		Update("credit", gorm.Expr("credit + ?", amountCents))

	if query.Error != nil {
		return query.Error
	}

	return nil
}
