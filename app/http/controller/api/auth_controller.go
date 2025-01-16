package api

import (
	"github.com/gin-gonic/gin"
	auth_request "github.com/hoopla/hoopla-api/app/http/request/auth"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
)

type AuthController struct {
	userService service.UserService
}

func NewAuthController(userService service.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body auth_request.LoginRequest true "Login Request"
// @Router /auth/login [post]
func (ctr *AuthController) Login(ctx *gin.Context) {
	loginRequest := auth_request.LoginRequest{}
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	loginResponse, code, err := ctr.userService.Login(loginRequest)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "OK!", loginResponse, nil)
	return
}

// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body auth_request.ConfirmSmsRequest true "Confirm Sms Request"
// @Router /auth/confirm-sms [post]
func (ctr *AuthController) ConfirmSms(ctx *gin.Context) {
	confirmSmsRequest := auth_request.ConfirmSmsRequest{}
	if err := ctx.ShouldBindJSON(&confirmSmsRequest); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	confirmSmsResponse, code, err := ctr.userService.ConfirmSms(confirmSmsRequest)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "OK!", confirmSmsResponse, nil)
}

// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body auth_request.ResendSmsRequest true "Resend Sms Request"
// @Router /auth/resend-sms [post]
func (ctr *AuthController) ResendSms(ctx *gin.Context) {
	resendSmsRequest := auth_request.ResendSmsRequest{}
	if err := ctx.ShouldBindJSON(&resendSmsRequest); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	resendSmsResponse, code, err := ctr.userService.ResendSms(resendSmsRequest)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "OK!", resendSmsResponse, nil)
}
