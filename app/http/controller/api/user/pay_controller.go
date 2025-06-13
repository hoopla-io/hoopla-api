package api_user

import (
	"github.com/gin-gonic/gin"
	user_pay_request "github.com/hoopla/hoopla-api/app/http/request/user/pay"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
	"github.com/hoopla/hoopla-api/utils"
)

type PayController struct {
	payService service.PayService
}

func NewPayController(payService service.PayService) *PayController {
	return &PayController{
		payService: payService,
	}
}

// @Tags User/Pay
// @Accept  json
// @Produce  json
// @Param data query user_pay_request.ServicesRequest true "Pay services list"
// @Router /v1/user/pay/services [get]
func (pc *PayController) Services(ctx *gin.Context) {
	var request user_pay_request.ServicesRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	services, code, err := pc.payService.Services(request)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", services, nil)
	return
}

// @Tags User/Pay
// @Accept  json
// @Produce  json
// @Param data query user_pay_request.TopUpRequest true "Top up balance"
// @Router /v1/user/pay/top-up [get]
func (pc *PayController) TopUp(ctx *gin.Context) {
	var request user_pay_request.TopUpRequest
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

	res, code, err := pc.payService.TopUp(request, &userHelper)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", res, nil)
	return
}
