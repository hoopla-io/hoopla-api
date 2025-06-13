package api

import (
	"github.com/gin-gonic/gin"
	subscriptions_request "github.com/hoopla/hoopla-api/app/http/request/subscriptions"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
	"github.com/hoopla/hoopla-api/utils"
)

type SubscriptionController struct {
	subscriptionService service.SubscriptionService
}

func NewSubscriptionController(subscriptionService service.SubscriptionService) *SubscriptionController {
	return &SubscriptionController{
		subscriptionService: subscriptionService,
	}
}

// @Tags Subscriptions
// @Accept  json
// @Produce  json
// @Param data query subscriptions_request.SubscriptionsRequest true "Subscriptions list"
// @Router /v1/subscriptions [get]
func (controller *SubscriptionController) Subscriptions(ctx *gin.Context) {
	var request subscriptions_request.SubscriptionsRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	subscriptions, code, err := controller.subscriptionService.GetSubscriptions(request)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", subscriptions, nil)
	return
}

// @Tags Subscriptions
// @Accept  json
// @Produce  json
// @Param data body subscriptions_request.BuySubscriptionRequest true "Buy Subscription"
// @Router /v1/subscriptions/buy [post]
func (controller *SubscriptionController) BuySubscription(ctx *gin.Context) {
	var request subscriptions_request.BuySubscriptionRequest
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

	statusCode, err := controller.subscriptionService.BuySubscription(request, userHelper.UserID)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", nil, nil)
	return
}
