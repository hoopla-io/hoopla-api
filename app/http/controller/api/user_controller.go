package api

import (
	"github.com/gin-gonic/gin"
	user_request "github.com/hoopla/hoopla-api/app/http/request/user"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
	"github.com/hoopla/hoopla-api/utils"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// @Tags User
// @Accept  json
// @Produce  json
// @Param data query user_request.GetMeRequest true "Get me"
// @Router /v1/user/get-me [get]
func (uc *UserController) GetMe(ctx *gin.Context) {
	var request user_request.GetMeRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	var userHelper utils.UserHelper
	err := userHelper.Init(ctx)
	if err != nil {
		response.BadRequestResponse(ctx, "can not parse token")
		return
	}

	userResource, code, err := uc.userService.GetUser(&userHelper)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", userResource, nil)
	return
}

// @Tags User
// @Accept  json
// @Produce  json
// @Param data query user_request.RefreshTokenRequest true "Get me"
// @Router /v1/user/refresh-token [patch]
func (uc *UserController) RefreshToken(ctx *gin.Context) {
	var request user_request.RefreshTokenRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	jwtResource, code, err := uc.userService.RefreshToken(request)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "OK!", jwtResource, nil)
	return
}

// @Tags User
// @Accept  json
// @Produce  json
// @Param data body user_request.LogoutRequest true "Logout from an account"
// @Router /v1/user/logout [post]
func (uc *UserController) Logout(ctx *gin.Context) {
	var logoutRequest user_request.LogoutRequest
	if err := ctx.ShouldBindQuery(&logoutRequest); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	var userHelper utils.UserHelper
	err := userHelper.Init(ctx)
	if err != nil {
		response.BadRequestResponse(ctx, "can not parse token")
		return
	}

	code, err := uc.userService.Logout(logoutRequest, userHelper.UserID)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "OK!", nil, nil)
	return
}

// @Tags QR
// @Accept  json
// @Produce  json
// @Param data query user_request.GenerateQrCodeRequest true "Request for new QR Code"
// @Router /v1/user/generate-qr-code [get]
func (uc *UserController) GenerateQRCode(ctx *gin.Context) {
	var request user_request.GenerateQrCodeRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	var userHelper utils.UserHelper
	err := userHelper.Init(ctx)
	if err != nil {
		response.BadRequestResponse(ctx, "can not parse token")
		return
	}

	qrCodeResource, err := uc.userService.GenerateQRCode(request, userHelper.UserID)
	if err != nil {
		response.BadRequestResponse(ctx, "can not generate qrcode")
		return
	}

	response.SuccessResponse(ctx, "OK!", qrCodeResource, nil)
	return
}

// @Tags User
// @Accept  json
// @Produce  json
// @Param data query user_request.DeactivateRequest true "Deactivate an account"
// @Router /v1/user/deactivate [delete]
func (uc *UserController) Deactivate(ctx *gin.Context) {
	var request user_request.DeactivateRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	var userHelper utils.UserHelper
	err := userHelper.Init(ctx)
	if err != nil {
		response.BadRequestResponse(ctx, "can not parse token")
		return
	}
	deactivated, err := uc.userService.DeactivateUser(&userHelper)
	if err != nil {
		response.BadRequestResponse(ctx, "can not deactivate")
		return
	}

	response.SuccessResponse(ctx, "OK!", deactivated, nil)
	return
}
