package api

import (
	"github.com/gin-gonic/gin"
	subscriptions_request "github.com/hoopla/hoopla-api/app/http/request/subscriptions"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
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
// @Router /subscriptions [get]
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
