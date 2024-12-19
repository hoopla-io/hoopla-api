package service

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/qahvazor/qahvazor/app/http/request"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/pkg"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/utils"
	"golang.org/x/exp/rand"
)


type AuthServiceImpl struct {
	repository *repository.Repository
	sessionCache *cache.Cache
}

func NewAuthService(repository *repository.Repository) AuthService {
	return &AuthServiceImpl{
		repository: repository,
		sessionCache: cache.New(10*time.Minute, 20*time.Minute),
	}
}

func (s *AuthServiceImpl) Login(data request.LoginRequest) (interface{}, error) {
	mobileProvider := utils.PredictProvider(data.PhoneNumber)

	uniqueId := fmt.Sprintf("%d", time.Now().UnixNano())
	sessionId := md5.Sum([]byte(data.PhoneNumber + uniqueId))
	sessionIdStr := fmt.Sprintf("%x", sessionId)
	session := dto.SessionDTO{
		PhoneNumber: data.PhoneNumber,
		MobileProvider: mobileProvider,
		Session: dto.Session{
			Attempt: 0,
			Code: rand.Intn(90000) + 10000,
			ExpiresAt: int(time.Now().Unix()) + 90,
		},
	}
    s.sessionCache.Set(sessionIdStr, session, 10*time.Minute)

	err := pkg.SendCode(mobileProvider, data.PhoneNumber, session.Session.Code)
	if err != nil {
		log.Printf("Error sending SMS: %v", err)
		return response.NewErrorResponse(500, "Try later! Error sending SMS"), nil
	}

	response := response.LoginResponse{
		PhoneNumber:      data.PhoneNumber,
		SessionID:        sessionIdStr,
		SessionExpiresAt: int64(session.Session.ExpiresAt),
	}
	return response, nil
}

func (s *AuthServiceImpl) ConfirmSms(data request.ConfirmSmsRequest) (interface{}, error){
	session, found := s.sessionCache.Get(data.SessionID)
	sessionData, ok := session.(dto.SessionDTO)
	if !ok {
		return response.NewErrorResponse(422, "Invalid session data!"), nil
	}
	if !found || session == nil {
		return response.NewErrorResponse(422, "Session Expired!"), nil
	}
	if sessionData.Session.Code != data.Code {
		return response.NewErrorResponse(422, "Invalid code!"), nil
	}
	if int64(sessionData.Session.ExpiresAt) < time.Now().Unix() {
		return response.NewErrorResponse(422, "Session Expired!"), nil
	}

	user, err := s.repository.GetByPhoneNumber(sessionData.PhoneNumber)
	if err != nil {
		return response.NewErrorResponse(500, "Try later!"), nil
	}

	if user == nil {
		createUserData := dto.UserDTO{
       		PhoneNumber: 	sessionData.PhoneNumber,               
			MobileProvider: sessionData.MobileProvider,
   		}
		user, err = s.repository.CreateUser(createUserData)
		if err != nil {
			return response.NewErrorResponse(500, "Try later!"), nil
    	}
	}

	expireAt := time.Now().Add(10 * time.Minute).Unix()
	// Generate the JWT token
	accessToken, err := utils.EncodeJWT(user.PhoneNumber, expireAt)
	if err != nil {
		log.Printf("Error encoding JWT: %v", err)
		return response.NewErrorResponse(500, "Failed to generate access token!"), nil
	}

	uid := uuid.New().String()
	hash := sha256.New()
	hash.Write([]byte(uid))
	refreshToken := hex.EncodeToString(hash.Sum(nil))

	response := response.ConfirmSmsResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireAt:  expireAt * 1000, 
		PhoneNumber: user.PhoneNumber,
	}
	return response, nil
}

func (s *AuthServiceImpl) ResendSms(data request.ResendSmsRequest) (interface{}, error) {
	session, found := s.sessionCache.Get(data.SessionID)
	if !found {
		return response.NewErrorResponse(422, "Session Expired!"), nil
	}
	sessionData, ok := session.(dto.SessionDTO)
	if !ok {
		return response.NewErrorResponse(422, "Invalid session data!"), nil
	}

	// Update session with new attempt, verification code, and expiration time
	sessionData.Session.Attempt = sessionData.Session.Attempt + 1
	sessionData.Session.Code = rand.Intn(90000) + 10000 
	sessionData.Session.ExpiresAt = int(time.Now().Unix()) + 90 

	sessionExpiration := time.Duration(90 * time.Second)
	s.sessionCache.Set(data.SessionID, sessionData, sessionExpiration)

	err := pkg.SendCode(sessionData.MobileProvider, sessionData.PhoneNumber, sessionData.Session.Code)
	if err != nil {
		log.Printf("Error sending SMS: %v", err)
		return response.NewErrorResponse(500, "Try later! Error sending SMS"), nil
	}

	response := response.ResendSmsResponse{
		PhoneNumber:      sessionData.PhoneNumber,
		SessionID:        data.SessionID,
		SessionExpiresAt: int64(sessionData.Session.ExpiresAt),
	}
	return response, nil
} 