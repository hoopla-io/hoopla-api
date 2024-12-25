package dashboard

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	subscription_request "github.com/qahvazor/qahvazor/app/http/request/subscription"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/service"
)

type SubscriptionController struct {
	service service.SubscriptionService
}

func NewSubscriptionController(service service.SubscriptionService) *SubscriptionController {
	return &SubscriptionController{
		service: service,
	}
}

func (ctr *SubscriptionController) Store(ctx *gin.Context) {
	storeRequest := subscription_request.StoreRequest{}
	if err := ctx.ShouldBind(&storeRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.Store(storeRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *SubscriptionController) Show(ctx *gin.Context) {
	subscriptionParamID := ctx.Param("subscription_id")
	subscriptionId, err := strconv.Atoi(subscriptionParamID)
	if err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	result, err := ctr.service.Show(uint(subscriptionId))
	if err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(
		ctx,
		http.StatusOK,
		"OK!",
		result,
		nil,
	)
}

func (ctr *SubscriptionController) List(ctx *gin.Context) {
	results, err := ctr.service.List()
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *SubscriptionController) Edit(ctx *gin.Context) {
	editSubscriptionRequest := subscription_request.EditRequest{}
	if err := ctx.ShouldBind(&editSubscriptionRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}
	
	if err := ctr.service.Edit(editSubscriptionRequest); err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", nil, nil)
}