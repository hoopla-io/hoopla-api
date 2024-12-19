package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qahvazor/qahvazor/app/http/request"
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

func (ctr *AuthController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.Login(loginRequest)
	if err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *AuthController) ConfirmSms(ctx *gin.Context) {
	confirmSmsRequest := request.ConfirmSmsRequest{}
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

func (ctr *AuthController) ResendSms(ctx *gin.Context) {
	resendSmsRequest := request.ResendSmsRequest{}
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