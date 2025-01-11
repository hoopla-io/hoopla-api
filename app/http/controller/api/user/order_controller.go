package api_user

import (
	"github.com/gin-gonic/gin"
	user_orders_request "github.com/qahvazor/qahvazor/app/http/request/user/orders"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/service"
	"github.com/qahvazor/qahvazor/utils"
)

type OrderController struct {
	userOrderService service.UserOrderService
}

func NewOrderController(userOrderService service.UserOrderService) *OrderController {
	return &OrderController{
		userOrderService: userOrderService,
	}
}

// @Tags User
// @Accept  json
// @Produce  json
// @Param data query user_orders_request.OrdersRequest true "User-orders list"
// @Router /user/orders [get]
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
