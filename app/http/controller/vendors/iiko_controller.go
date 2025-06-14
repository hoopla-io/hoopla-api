package vendor_controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hoopla/hoopla-api/app/http/response"
)

type IikoController struct{}

func NewIikoController() *IikoController {
	return &IikoController{}
}

// @Tags IIKO
// @Accept  json
// @Produce  json
// @Param data body vendors_poster_request.WebhookRequest true "Webhook for iiko"
// @Router /vendors/iiko/webhook [post]
func (c *IikoController) Webhook(ctx *gin.Context) {
	response.SuccessResponse(ctx, "OK!", nil, nil)
}
