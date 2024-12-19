package service

import (
	"github.com/qahvazor/qahvazor/app/http/request"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/repository"
)

type Service struct {
	AuthService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repo),
	}
}

type AuthService interface {
	Login(data request.LoginRequest) (response.LoginResponse, error)
}
