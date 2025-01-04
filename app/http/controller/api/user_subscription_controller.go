package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/qahvazor/qahvazor/internal/service"
)

type UserSubscriptionController struct {
	userSubscriptionService service.UserSubscriptionService
}

func NewUserSubscriptionController(userSubscriptionService service.UserSubscriptionService) *UserSubscriptionController {
	return &UserSubscriptionController{
		userSubscriptionService: userSubscriptionService,
	}
}

func (c *UserSubscriptionController) GetUserActiveSubscription(ctx *gin.Context) {
	userIDParam := ctx.Param("user_id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	subscription, err := c.userSubscriptionService.GetUserActiveSubscription(ctx.Request.Context(), uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No active subscription found"})
		return
	}
	ctx.JSON(http.StatusOK, subscription)
}

func (c *UserSubscriptionController) AssignSubscriptionToUser(ctx *gin.Context) {
	var request struct {
		UserID         uint `json:"user_id"`
		SubscriptionID uint `json:"subscription_id"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.userSubscriptionService.AssignSubscriptionToUser(ctx.Request.Context(), request.UserID, request.SubscriptionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Subscription assigned successfully"})
}
