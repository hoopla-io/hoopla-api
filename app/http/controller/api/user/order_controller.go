package api_user

import (
	"github.com/gin-gonic/gin"
	user_orders_request "github.com/hoopla/hoopla-api/app/http/request/user/orders"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
	"github.com/hoopla/hoopla-api/utils"
)

type OrderController struct {
	userOrderService service.UserOrderService
}

func NewOrderController(userOrderService service.UserOrderService) *OrderController {
	return &OrderController{
		userOrderService: userOrderService,
	}
}

// @Tags User/Orders
// @Accept  json
// @Produce  json
// @Param data body user_orders_request.CreateRequest true "New order"
// @Router /v1/user/orders/create [post]
func (controller *OrderController) Create(ctx *gin.Context) {
	var request user_orders_request.CreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	var userHelper utils.UserHelper
	err := userHelper.Init(ctx)
	if err != nil {
		response.BadRequestResponse(ctx, "can not parse token")
		return
	}

	userOrder, code, err := controller.userOrderService.CreateOrder(request, &userHelper)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "OK!", userOrder, nil)
	return
}

// @Tags User/Orders
// @Accept  json
// @Produce  json
// @Param data query user_orders_request.OrdersRequest true "User-orders list"
// @Router /v1/user/orders/orders-list [get]
func (controller *OrderController) Orders(ctx *gin.Context) {
	var request user_orders_request.OrdersRequest
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

	orders, code, err := controller.userOrderService.GetOrders(request, userHelper.UserID)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", orders, nil)
	return
}

// @Tags User
// @Accept  json
// @Produce  json
// @Router /v1/user/orders/drinks-stat [get]
func (controller *OrderController) DrinksStat(ctx *gin.Context) {
	var userHelper utils.UserHelper
	err := userHelper.Init(ctx)
	if err != nil {
		response.BadRequestResponse(ctx, "can not parse token")
		return
	}

	stat, code, err := controller.userOrderService.GetDrinksStat(userHelper.UserID)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", stat, nil)
	return
}
