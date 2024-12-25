package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth_request "github.com/qahvazor/qahvazor/app/http/request/auth"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/service"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body auth_request.LoginRequest true "Login Request"
// @Success 200 {array} auth_response.LoginResponse "Login Response"
// @Router /auth/login [post]
func (ctr *AuthController) Login(ctx *gin.Context) {
	loginRequest := auth_request.LoginRequest{}
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.Login(loginRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body auth_request.ConfirmSmsRequest true "Confirm Sms Request"
// @Success 200 {array} auth_response.ConfirmSmsResponse "Confirm Sms Response"
// @Router /auth/confirm-sms [post]
func (ctr *AuthController) ConfirmSms(ctx *gin.Context) {
	confirmSmsRequest := auth_request.ConfirmSmsRequest{}
	if err := ctx.ShouldBindJSON(&confirmSmsRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.ConfirmSms(confirmSmsRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body auth_request.ResendSmsRequest true "Resend Sms Request"
// @Success 200 {array} auth_response.ResendSmsResponse "Resend Sms Response"
// @Router /auth/resend-sms [post]
func (ctr *AuthController) ResendSms(ctx *gin.Context) {
	resendSmsRequest := auth_request.ResendSmsRequest{}
	if err := ctx.ShouldBindJSON(&resendSmsRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.ResendSms(resendSmsRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}