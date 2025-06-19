package vendor_controllers

import (
	"fmt"
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
	fmt.Println("IIIKOOO__------------------------request")
	var r map[string]interface{}
	if err := ctx.ShouldBindBodyWithJSON(&r); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	fmt.Println("IIIKOOO__------------------------")
	fmt.Println(r)
	fmt.Println("IIIKOOO__------------------------")

	response.SuccessResponse(ctx, "OK!", nil, nil)
}
