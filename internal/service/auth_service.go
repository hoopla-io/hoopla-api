package service

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	auth_resource "github.com/qahvazor/qahvazor/app/http/resource/auth"
	"github.com/qahvazor/qahvazor/pkg/itvmsq"
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	auth_request "github.com/qahvazor/qahvazor/app/http/request/auth"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/utils"
	"golang.org/x/exp/rand"
)

type AuthService interface {
	Login(data auth_request.LoginRequest) (*auth_resource.SessionResource, int, error)
	ConfirmSms(data auth_request.ConfirmSmsRequest) (*auth_resource.LoginResource, int, error)
	ResendSms(data auth_request.ResendSmsRequest) (*auth_resource.SessionResource, int, error)
}

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	sessionCache   *cache.Cache
}

func NewAuthService(AuthRepository repository.AuthRepository) AuthService {
	return &AuthServiceImpl{
		AuthRepository: AuthRepository,
		sessionCache:   cache.New(10*time.Minute, 20*time.Minute),
	}
}

func (s *AuthServiceImpl) Login(data auth_request.LoginRequest) (*auth_resource.SessionResource, int, error) {
	mobileProvider := utils.PredictProvider(data.PhoneNumber)

	uniqueId := fmt.Sprintf("%d", time.Now().UnixNano())
	sessionId := md5.Sum([]byte(data.PhoneNumber + uniqueId))
	sessionIdStr := fmt.Sprintf("%x", sessionId)
	session := dto.SessionDTO{
		PhoneNumber:    data.PhoneNumber,
		MobileProvider: mobileProvider,
		Session: dto.Session{
			Code:      rand.Intn(90000) + 10000,
			ExpiresAt: time.Now().Unix() + 90,
		},
	}
	s.sessionCache.Set(sessionIdStr, session, 10*time.Minute)

	err := itvmsq.SendCode(mobileProvider, data.PhoneNumber, session.Session.Code)
	if err != nil {
		return nil, 500, err
	}

	sessionResource := &auth_resource.SessionResource{
		PhoneNumber:      data.PhoneNumber,
		SessionID:        sessionIdStr,
		SessionExpiresAt: session.Session.ExpiresAt,
	}
	return sessionResource, 200, nil
}

func (s *AuthServiceImpl) ConfirmSms(data auth_request.ConfirmSmsRequest) (*auth_resource.LoginResource, int, error) {
	session, callback := s.sessionCache.Get(data.SessionID)
	if !callback {
		return nil, 404, errors.New("session not found")
	}
	sessionData, callback := session.(dto.SessionDTO)
	if !callback {
		return nil, 400, errors.New("invalid session data")
	}

	if sessionData.Session.ExpiresAt < time.Now().Unix() {
		return nil, 422, errors.New("session expired")
	}

	if sessionData.Session.Code != data.Code {
		return nil, 422, errors.New("invalid session code")
	}

	isNewUser := false
	user, err := s.AuthRepository.GetByPhoneNumber(sessionData.PhoneNumber)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) { // new user
		createUserData := dto.UserDTO{
			PhoneNumber:    sessionData.PhoneNumber,
			MobileProvider: sessionData.MobileProvider,
		}
		user, err = s.AuthRepository.CreateUser(createUserData)
		isNewUser = true
	}
	if err != nil {
		return nil, 500, err
	}

	// Generate the JWT token
	expireAt := time.Now().Add(10 * time.Minute).Unix()
	accessToken, err := utils.EncodeJWT(user.ID, user.PhoneNumber, expireAt)
	if err != nil {
		return nil, 500, err
	}

	uid := uuid.New().String()
	hash := sha256.New()
	hash.Write([]byte(uid))
	refreshToken := hex.EncodeToString(hash.Sum(nil))

	loginResource := &auth_resource.LoginResource{
		UserID:       user.ID,
		PhoneNumber:  user.PhoneNumber,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireAt:     expireAt * 1000,
		IsNewUser:    isNewUser,
	}
	return loginResource, 200, nil
}

func (s *AuthServiceImpl) ResendSms(data auth_request.ResendSmsRequest) (*auth_resource.SessionResource, int, error) {
	session, callback := s.sessionCache.Get(data.SessionID)
	if !callback {
		return nil, 404, errors.New("session not found")
	}
	sessionData, callback := session.(dto.SessionDTO)
	if !callback {
		return nil, 400, errors.New("invalid session data")
	}

	sessionData.Session.Code = rand.Intn(90000) + 10000
	sessionData.Session.ExpiresAt = time.Now().Unix() + 90

	s.sessionCache.Delete(data.SessionID)

	uniqueId := fmt.Sprintf("%d", time.Now().UnixNano())
	sessionId := md5.Sum([]byte(sessionData.PhoneNumber + uniqueId))
	sessionIdStr := fmt.Sprintf("%x", sessionId)

	s.sessionCache.Set(sessionIdStr, sessionData, 10*time.Minute)

	err := itvmsq.SendCode(sessionData.MobileProvider, sessionData.PhoneNumber, sessionData.Session.Code)
	if err != nil {
		return nil, 500, err
	}

	sessionResource := &auth_resource.SessionResource{
		PhoneNumber:      sessionData.PhoneNumber,
		SessionID:        sessionIdStr,
		SessionExpiresAt: sessionData.Session.ExpiresAt,
	}
	return sessionResource, 200, nil
}
