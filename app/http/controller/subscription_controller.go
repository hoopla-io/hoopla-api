package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/qahvazor/qahvazor/internal/service"
)

type SubscriptionController struct {
	subscriptionService service.SubscriptionService
}

func NewSubscriptionController(subscriptionService service.SubscriptionService) *SubscriptionController {
	return &SubscriptionController{
		subscriptionService: subscriptionService,
	}
}

func (c *SubscriptionController) GetAllSubscriptions(ctx *gin.Context) {
	subscriptions, err := c.subscriptionService.GetAllSubscriptions(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, subscriptions)
}

func (c *SubscriptionController) GetSubscriptionByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	subscription, err := c.subscriptionService.GetSubscriptionByID(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}
	ctx.JSON(http.StatusOK, subscription)
}
