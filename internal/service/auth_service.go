package service

import (
	"crypto/md5"
	"fmt"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/qahvazor/qahvazor/app/http/request"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/utils"
	"golang.org/x/exp/rand"
)


type AuthServiceImpl struct {
	repository *repository.Repository
}

func NewAuthService(repository *repository.Repository) AuthService {
	return &AuthServiceImpl{
		repository: repository,
	}
}

func (s *AuthServiceImpl) Login(data request.LoginRequest) (response.LoginResponse, error) {
	user, err := s.repository.GetByPhoneNumber(data.PhoneNumber)
	if err != nil {
		log.Printf("error getting user by phone number: %s", err.Error())
		return response.LoginResponse{}, nil
	}

	mobileProvider := utils.PredictProvider(data.PhoneNumber)

	sessionId := md5.Sum([]byte(data.PhoneNumber))
	sessionIdStr := fmt.Sprintf("%x", sessionId)
	session := map[string]interface{}{
		"phone_number": data.PhoneNumber,
		"session": map[string]interface{}{
			"attempt":          0,
			"code":             rand.Intn(90000) + 10000,
			"expires_at":       time.Now().Unix() + 90, // expires in 90 seconds
		},
	}

	
	var sessionCache = cache.New(10*time.Minute, 20*time.Minute)
	sessionCache.Add(sessionIdStr, session, 600*time.Second)

	log.Println(sessionId)
	log.Println(mobileProvider)
	log.Println(user)

	response := response.LoginResponse{
		PhoneNumber:      data.PhoneNumber,
		SessionID:        sessionIdStr,
		SessionExpiresAt: session["session"].(map[string]interface{})["expires_at"].(int64),
	}
	return response, nil
}