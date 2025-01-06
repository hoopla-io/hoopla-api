package service

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	user_request "github.com/qahvazor/qahvazor/app/http/request/user"
	auth_resource "github.com/qahvazor/qahvazor/app/http/resource/auth"
	"github.com/qahvazor/qahvazor/pkg/itvmsq"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"time"

	"github.com/patrickmn/go-cache"
	auth_request "github.com/qahvazor/qahvazor/app/http/request/auth"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/repository"
	"github.com/qahvazor/qahvazor/utils"
)

type UserService interface {
	Login(data auth_request.LoginRequest) (*auth_resource.SessionResource, int, error)
	ConfirmSms(data auth_request.ConfirmSmsRequest) (*auth_resource.LoginResource, int, error)
	ResendSms(data auth_request.ResendSmsRequest) (*auth_resource.SessionResource, int, error)
	RefreshToken(data user_request.RefreshTokenRequest) (*auth_resource.JwtResource, int, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	sessionCache   *cache.Cache
}

func NewUserService(UserRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		sessionCache:   cache.New(10*time.Minute, 20*time.Minute),
	}
}

func (s *UserServiceImpl) Login(data auth_request.LoginRequest) (*auth_resource.SessionResource, int, error) {
	mobileProvider := utils.PredictProvider(data.PhoneNumber)

	uniqueId := fmt.Sprintf("%d", time.Now().UnixNano())
	sessionId := md5.Sum([]byte(data.PhoneNumber + uniqueId))
	sessionIdStr := fmt.Sprintf("%x", sessionId)
	session := dto.SessionDTO{
		PhoneNumber:    data.PhoneNumber,
		MobileProvider: mobileProvider,
		Session: dto.Session{
			Code:      rand.Intn(89999) + 10000,
			ExpiresAt: time.Now().Unix() + 90,
		},
	}
	s.sessionCache.Set(sessionIdStr, session, 10*time.Minute)

	err := itvmsq.SendCode(mobileProvider, data.PhoneNumber, session.Session.Code)
	if err != nil {
		return nil, 500, err
	}

	sessionResource := &auth_resource.SessionResource{
		PhoneNumber:       data.PhoneNumber,
		SessionID:         sessionIdStr,
		SessionExpiresAt:  session.Session.ExpiresAt,
		SessionExpireAtMs: session.Session.ExpiresAt * 1000,
	}
	return sessionResource, 200, nil
}

func (s *UserServiceImpl) ConfirmSms(data auth_request.ConfirmSmsRequest) (*auth_resource.LoginResource, int, error) {
	session, callback := s.sessionCache.Get(data.SessionID)
	if !callback {
		return nil, 422, errors.New("session not found")
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

	uid := uuid.New().String()
	hash := sha256.New()
	hash.Write([]byte(uid))
	refreshToken := hex.EncodeToString(hash.Sum(nil))

	isNewUser := false
	user, err := s.UserRepository.GetByPhoneNumber(sessionData.PhoneNumber)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) { // new user
		createUserData := dto.UserDTO{
			PhoneNumber:    sessionData.PhoneNumber,
			MobileProvider: sessionData.MobileProvider,
			RefreshToken:   refreshToken,
		}
		user, err = s.UserRepository.CreateUser(createUserData)
		isNewUser = true
	}
	if err != nil {
		return nil, 500, err
	}

	if isNewUser == false { // updating token for existing user
		err = s.UserRepository.UpdateToken(refreshToken, user)
		if err != nil {
			return nil, 500, err
		}
	}

	// Generate the JWT token
	expireAt := time.Now().Add(1 * time.Hour).Unix()
	accessToken, err := utils.EncodeJWT(user.ID, user.PhoneNumber, expireAt)
	if err != nil {
		return nil, 500, err
	}

	loginResource := &auth_resource.LoginResource{
		UserID:      user.ID,
		PhoneNumber: user.PhoneNumber,
		IsNewUser:   isNewUser,
		Jwt: auth_resource.JwtResource{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpireAt:     expireAt,
			ExpireAtMs:   expireAt * 1000,
		},
	}
	return loginResource, 200, nil
}

func (s *UserServiceImpl) ResendSms(data auth_request.ResendSmsRequest) (*auth_resource.SessionResource, int, error) {
	session, callback := s.sessionCache.Get(data.SessionID)
	if !callback {
		return nil, 422, errors.New("session not found")
	}
	sessionData, callback := session.(dto.SessionDTO)
	if !callback {
		return nil, 400, errors.New("invalid session data")
	}

	sessionData.Session.Code = rand.Intn(89999) + 10000
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
		PhoneNumber:       sessionData.PhoneNumber,
		SessionID:         sessionIdStr,
		SessionExpiresAt:  sessionData.Session.ExpiresAt,
		SessionExpireAtMs: sessionData.Session.ExpiresAt * 1000,
	}
	return sessionResource, 200, nil
}

func (s *UserServiceImpl) RefreshToken(data user_request.RefreshTokenRequest) (*auth_resource.JwtResource, int, error) {
	user, err := s.UserRepository.GetByRefreshToken(data.RefreshToken)
	if err != nil {
		return nil, 500, err
	}

	uid := uuid.New().String()
	hash := sha256.New()
	hash.Write([]byte(uid))
	refreshToken := hex.EncodeToString(hash.Sum(nil))

	err = s.UserRepository.UpdateToken(refreshToken, user)
	if err != nil {
		return nil, 500, err
	}

	// Generate the JWT token
	expireAt := time.Now().Add(1 * time.Hour).Unix()
	accessToken, err := utils.EncodeJWT(user.ID, user.PhoneNumber, expireAt)
	if err != nil {
		return nil, 500, err
	}

	jwtResource := &auth_resource.JwtResource{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireAt:     expireAt,
		ExpireAtMs:   expireAt * 1000,
	}
	return jwtResource, 200, nil
}
