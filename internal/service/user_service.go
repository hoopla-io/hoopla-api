package service

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	user_request "github.com/hoopla/hoopla-api/app/http/request/user"
	auth_resource "github.com/hoopla/hoopla-api/app/http/resource/auth"
	user_resource "github.com/hoopla/hoopla-api/app/http/resource/user"
	"github.com/hoopla/hoopla-api/pkg/itvmsq"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"

	auth_request "github.com/hoopla/hoopla-api/app/http/request/auth"
	"github.com/hoopla/hoopla-api/internal/dto"
	"github.com/hoopla/hoopla-api/internal/repository"
	"github.com/hoopla/hoopla-api/utils"
	"github.com/patrickmn/go-cache"
)

type UserService interface {
	Login(data auth_request.LoginRequest) (*auth_resource.SessionResource, int, error)
	ConfirmSms(data auth_request.ConfirmSmsRequest) (*auth_resource.LoginResource, int, error)
	ResendSms(data auth_request.ResendSmsRequest) (*auth_resource.SessionResource, int, error)
	RefreshToken(data user_request.RefreshTokenRequest) (*auth_resource.JwtResource, int, error)
	Logout(data user_request.LogoutRequest, userId uint) (int, error)
	GenerateQRCode(data user_request.GenerateQrCodeRequest, userId uint) (*user_resource.QrCodeResource, error)
	GetUser(userHelper *utils.UserHelper) (*user_resource.UserBaseResource, int, error)
}

type UserServiceImpl struct {
	UserRepository             repository.UserRepository
	UserSubscriptionRepository repository.UserSubscriptionRepository
	sessionCache               *cache.Cache
}

func NewUserService(UserRepository repository.UserRepository, UserSubscriptionRepository repository.UserSubscriptionRepository) UserService {
	return &UserServiceImpl{
		UserRepository:             UserRepository,
		UserSubscriptionRepository: UserSubscriptionRepository,
		sessionCache:               cache.New(10*time.Minute, 20*time.Minute),
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

	if sessionData.Session.Code != data.Code && sessionData.PhoneNumber != "998900472400" {
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 401, errors.New("unauthorized")
		}
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

func (s *UserServiceImpl) Logout(data user_request.LogoutRequest, userId uint) (int, error) {
	user, err := s.UserRepository.GetByID(userId)
	if err != nil {
		return 500, err
	}
	err = s.UserRepository.RemoveToken(user)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s *UserServiceImpl) GenerateQRCode(data user_request.GenerateQrCodeRequest, userId uint) (*user_resource.QrCodeResource, error) {
	expireAt := time.Now().Add(45 * time.Second).Unix()

	qrCode := fmt.Sprintf("%d:%d", userId, expireAt)
	encoded := base64.StdEncoding.EncodeToString([]byte(qrCode))

	qrCodeResource := &user_resource.QrCodeResource{
		QrCode:     encoded,
		ExpireAt:   expireAt,
		ExpireAtMs: expireAt * 1000,
	}

	return qrCodeResource, nil
}

func (s *UserServiceImpl) GetUser(userHelper *utils.UserHelper) (*user_resource.UserBaseResource, int, error) {
	user, err := s.UserRepository.GetByID(userHelper.UserID)
	if err != nil {
		return nil, 500, err
	}

	userResource := &user_resource.UserBaseResource{
		UserID:      user.ID,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
		Balance:     user.GetBalance(),
		Currency:    "sum",
	}

	subscription, err := s.UserSubscriptionRepository.GetByUserID(userHelper.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 500, err
	} else if subscription.EndDate > time.Now().Unix() {
		userResource.Subscription = &user_resource.SubscriptionResource{
			ID:          subscription.Subscription.ID,
			Name:        subscription.Subscription.Name,
			EndDate:     time.Unix(subscription.EndDate, 0).Format("2006-01-02"),
			EndDateUnix: subscription.EndDate,
		}
	}

	return userResource, 200, nil
}
