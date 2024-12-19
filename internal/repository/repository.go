package repository

import (
	"github.com/qahvazor/qahvazor/internal/dto"
	"gorm.io/gorm"
)

type Repository struct {
	AuthRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db),
	}
}

type AuthRepository interface {
	GetByPhoneNumber(phoneNumber string) (*dto.UserDTO, error)
}
