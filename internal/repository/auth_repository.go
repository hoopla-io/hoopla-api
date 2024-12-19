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

// func (r *AuthRepositoryImpl) Login(data dto.CreateTransactionDTO) (*dto.TransactionDTO, error) {
// 	var createdTransaction dto.TransactionDTO

// 	query := r.db.Create(&model.TransactionModel{
// 		TransID:       data.TransID,
// 		ServiceID: int(data.ServiceID),
// 		Status:        "CREATED",
// 		AccountNumber: data.Params.Account,
// 		Amount:        data.Amount / 100,
// 		CreatedAt:     int64(data.Timestamp),
// 	}).Scan(&createdTransaction)
// 	if query.Error != nil {
// 		return nil, query.Error
// 	}

// 	createdTransaction.TransTime = int64(data.Timestamp)

// 	return &createdTransaction, nil
// }