package api

import (
	"github.com/gin-gonic/gin"
	user_request "github.com/qahvazor/qahvazor/app/http/request/user"
	user_resource "github.com/qahvazor/qahvazor/app/http/resource/user"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/service"
	"github.com/qahvazor/qahvazor/utils"
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
// @Router /user/get-me [get]
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

	response.SuccessResponse(
		ctx, "ok!",
		user_resource.UserBaseResource{
			UserID:      userHelper.UserID,
			PhoneNumber: userHelper.PhoneNumber,
			Name:        "test",
		},
		nil)
	return
}

// @Tags User
// @Accept  json
// @Produce  json
// @Param data query user_request.RefreshTokenRequest true "Get me"
// @Router /user/refresh-token [patch]
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
// @Router /user/logout [post]
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
